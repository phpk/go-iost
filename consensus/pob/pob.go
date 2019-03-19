package pob

import (
	"errors"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/iost-official/go-iost/account"
	"github.com/iost-official/go-iost/common"
	"github.com/iost-official/go-iost/consensus/synchro"
	"github.com/iost-official/go-iost/consensus/synchro/pb"
	"github.com/iost-official/go-iost/core/block"
	"github.com/iost-official/go-iost/core/blockcache"
	"github.com/iost-official/go-iost/core/global"
	"github.com/iost-official/go-iost/core/txpool"
	"github.com/iost-official/go-iost/crypto"
	"github.com/iost-official/go-iost/db"
	"github.com/iost-official/go-iost/ilog"
	"github.com/iost-official/go-iost/metrics"
	"github.com/iost-official/go-iost/p2p"
)

var (
	metricsGeneratedBlockCount   = metrics.NewCounter("iost_pob_generated_block", nil)
	metricsVerifyBlockCount      = metrics.NewCounter("iost_pob_verify_block", nil)
	metricsConfirmedLength       = metrics.NewGauge("iost_pob_confirmed_length", nil)
	metricsMode                  = metrics.NewGauge("iost_node_mode", nil)
	metricsTimeCost              = metrics.NewGauge("iost_time_cost", nil)
	metricsTransferCost          = metrics.NewGauge("iost_transfer_cost", nil)
	metricsGenerateBlockTimeCost = metrics.NewGauge("iost_generate_block_time_cost", nil)
)

var (
	errSingle     = errors.New("single block")
	errDuplicate  = errors.New("duplicate block")
	errOutOfLimit = errors.New("block out of limit in one slot")
)

var (
	blockNumPerWitness = 6
	maxBlockNumber     = int64(10000)
	subSlotTime        = 500 * time.Millisecond
	genBlockTime       = 400 * time.Millisecond
	last2GenBlockTime  = 50 * time.Millisecond
)

//PoB is a struct that handles the consensus logic.
type PoB struct {
	account      *account.KeyPair
	baseVariable global.BaseVariable
	blockChain   block.Chain
	blockCache   blockcache.BlockCache
	txPool       txpool.TxPool
	p2pService   p2p.Service
	verifyDB     db.MVCCDB
	produceDB    db.MVCCDB
	sync         *synchro.Sync

	exitSignal       chan struct{}
	quitGenerateMode chan struct{}
	wg               *sync.WaitGroup
	mu               *sync.RWMutex
}

// New init a new PoB.
func New(baseVariable global.BaseVariable, blockCache blockcache.BlockCache, txPool txpool.TxPool, p2pService p2p.Service) *PoB {
	// TODO: Move the code to account struct.
	accSecKey := baseVariable.Config().ACC.SecKey
	accAlgo := baseVariable.Config().ACC.Algorithm
	account, err := account.NewKeyPair(common.Base58Decode(accSecKey), crypto.NewAlgorithm(accAlgo))
	if err != nil {
		ilog.Fatalf("NewKeyPair failed, stop the program! err:%v", err)
	}

	// TODO: Organize the owner and lifecycle of all metrics.
	metricsMode.Set(float64(2), nil)

	p := PoB{
		account:      account,
		baseVariable: baseVariable,
		blockChain:   baseVariable.BlockChain(),
		blockCache:   blockCache,
		txPool:       txPool,
		p2pService:   p2pService,
		verifyDB:     baseVariable.StateDB(),
		produceDB:    baseVariable.StateDB().Fork(),
		sync:         nil,

		exitSignal:       make(chan struct{}),
		quitGenerateMode: make(chan struct{}),
		wg:               new(sync.WaitGroup),
		mu:               new(sync.RWMutex),
	}

	p.recoverBlockcache()
	close(p.quitGenerateMode)

	return &p
}

func (p *PoB) recoverBlockcache() error {
	err := p.blockCache.Recover(p)
	if err != nil {
		ilog.Error("Failed to recover blockCache, err: ", err)
		ilog.Info("Don't Recover, Move old file to BlockCacheWALCorrupted")
		err = p.blockCache.NewWAL(p.baseVariable.Config())
		if err != nil {
			ilog.Error(" Failed to NewWAL, err: ", err)
		}
	}
	return err
}

// Start make the PoB run.
func (p *PoB) Start() error {
	p.sync = synchro.New(p.p2pService, p.blockCache, p.blockChain)

	p.wg.Add(2)
	go p.verifyLoop()
	go p.scheduleLoop()
	return nil
}

// Stop make the PoB stop.
func (p *PoB) Stop() {
	close(p.exitSignal)
	p.wg.Wait()

	p.sync.Close()
}

// Mode return the mode of pob.
func (p *PoB) Mode() string {
	if p.sync == nil {
		return "ModeInit"
	} else if p.sync.IsCatchingUp() {
		return "ModeSync"
	} else {
		return "ModeNormal"
	}
}

func (p *PoB) broadcastBlockHash(blk *block.Block) {
	blkInfo := &msgpb.BlockInfo{
		Number: blk.Head.Number,
		Hash:   blk.HeadHash(),
	}
	b, err := proto.Marshal(blkInfo)
	if err != nil {
		ilog.Errorf("fail to encode block hash, err=%v, blockHash=%+v", err, *blkInfo)
	} else {
		p.p2pService.Broadcast(b, p2p.NewBlockHash, p2p.UrgentMessage)
	}
}

func calculateTime(blk *block.Block) float64 {
	return float64((time.Now().UnixNano() - blk.Head.Time) / 1e6)
}

func (p *PoB) doVerifyBlock(blkMsg *synchro.BlockMessage) {
	blk := blkMsg.Blk

	switch blkMsg.P2PType {
	case p2p.NewBlock:
		t1 := calculateTime(blk)
		metricsTransferCost.Set(t1, nil)
		err := p.handleRecvBlock(blk)
		t2 := calculateTime(blk)
		metricsTimeCost.Set(t2, nil)
		if err == errSingle || err == nil {
			p.broadcastBlockHash(blk)
		}
		if err != nil && err != errSingle && err != errDuplicate {
			ilog.Warnf("received new block error, err:%v", err)
			return
		}
	case p2p.SyncBlockResponse:
		err := p.handleRecvBlock(blk)
		if err != nil && err != errSingle && err != errDuplicate {
			ilog.Warnf("received sync block error, err:%v", err)
			return
		}
	}
	metricsVerifyBlockCount.Add(1, nil)
}

func (p *PoB) verifyLoop() {
	defer p.wg.Done()
	for {
		select {
		case blkMsg := <-p.sync.IncomingBlock():
			select {
			case <-p.quitGenerateMode:
			}
			if p.blockCache.Head().Head.Number+maxBlockNumber < blkMsg.Blk.Head.Number {
				ilog.Debugf("block number is too large, block number:%v", blkMsg.Blk.Head.Number)
				continue
			}

			p.doVerifyBlock(blkMsg)

		case <-p.exitSignal:
			return
		}
	}
}

func (p *PoB) scheduleLoop() {
	defer p.wg.Done()
	nextSchedule := timeUntilNextSchedule(time.Now().UnixNano())
	ilog.Debugf("nextSchedule: %.2f", time.Duration(nextSchedule).Seconds())
	pubkey := p.account.ReadablePubkey()

	var slotFlag int64
	for {
		select {
		case <-time.After(time.Duration(nextSchedule)):
			time.Sleep(time.Millisecond)
			if p.sync.IsCatchingUp() {
				metricsMode.Set(float64(1), nil)
				continue
			} else {
				metricsMode.Set(float64(0), nil)
			}
			t := time.Now()
			pTx, head := p.txPool.PendingTx()
			witnessList := head.Active()
			if slotFlag != slotOfSec(t.Unix()) && witnessOfNanoSec(t.UnixNano(), witnessList) == pubkey {
				p.quitGenerateMode = make(chan struct{})
				slotFlag = slotOfSec(t.Unix())
				generateBlockTicker := time.NewTicker(subSlotTime)
				for num := 0; num < blockNumPerWitness; num++ {
					p.gen(num, pTx, head)
					if num == blockNumPerWitness-1 {
						break
					}
					select {
					case <-generateBlockTicker.C:
					}
					pTx, head = p.txPool.PendingTx()
					witnessList = head.Active()
					if witnessOfNanoSec(time.Now().UnixNano(), witnessList) != pubkey {
						break
					}
				}
				close(p.quitGenerateMode)
				generateBlockTicker.Stop()
			}
			nextSchedule = timeUntilNextSchedule(time.Now().UnixNano())
			ilog.Debugf("nextSchedule: %.2f", time.Duration(nextSchedule).Seconds())
		case <-p.exitSignal:
			return
		}
	}
}

func (p *PoB) gen(num int, pTx *txpool.SortedTxMap, head *blockcache.BlockCacheNode) {
	limitTime := genBlockTime
	if num >= blockNumPerWitness-2 {
		limitTime = last2GenBlockTime
	}
	p.txPool.Lock()
	blk, err := generateBlock(p.account, p.txPool, p.produceDB, limitTime, pTx, head)
	p.txPool.Release()
	if err != nil {
		ilog.Error(err)
		return
	}
	p.printStatistics(num, blk)
	blkByte, err := blk.Encode()
	if err != nil {
		ilog.Error(err)
		return
	}
	p.p2pService.Broadcast(blkByte, p2p.NewBlock, p2p.UrgentMessage)
	metricsGenerateBlockTimeCost.Set(calculateTime(blk), nil)
	err = p.handleRecvBlock(blk)
	if err != nil {
		ilog.Errorf("[pob] handle block from myself, err:%v", err)
		return
	}
}

func (p *PoB) printStatistics(num int, blk *block.Block) {
	ptx, _ := p.txPool.PendingTx()
	ilog.Infof("Gen block - @%v id:%v..., t:%v, num:%v, confirmed:%v, txs:%v, pendingtxs:%v, et:%vms",
		num,
		p.account.ReadablePubkey()[:10],
		blk.Head.Time,
		blk.Head.Number,
		p.blockCache.LinkedRoot().Head.Number,
		len(blk.Txs),
		ptx.Size(),
		calculateTime(blk),
	)
}

// RecoverBlock recover block from block cache wal
func (p *PoB) RecoverBlock(blk *block.Block) error {
	_, err := p.blockCache.Find(blk.HeadHash())
	if err == nil {
		return errDuplicate
	}
	err = verifyBasics(blk, blk.Sign)
	if err != nil {
		return err
	}
	parent, err := p.blockCache.Find(blk.Head.ParentHash)
	p.blockCache.Add(blk)
	if err == nil && parent.Type == blockcache.Linked {
		return p.addExistingBlock(blk, parent, true)
	}
	return errSingle
}

func (p *PoB) handleRecvBlock(blk *block.Block) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	_, err := p.blockCache.Find(blk.HeadHash())
	if err == nil {
		return errDuplicate
	}

	err = verifyBasics(blk, blk.Sign)
	if err != nil {
		return err
	}

	parent, err := p.blockCache.Find(blk.Head.ParentHash)
	p.blockCache.Add(blk)
	if err == nil && parent.Type == blockcache.Linked {
		return p.addExistingBlock(blk, parent, false)
	}
	return errSingle
}

func (p *PoB) addExistingBlock(blk *block.Block, parentNode *blockcache.BlockCacheNode, replay bool) error {
	node, _ := p.blockCache.Find(blk.HeadHash())

	if parentNode.Block.Head.Witness != blk.Head.Witness ||
		slotOfSec(parentNode.Block.Head.Time/1e9) != slotOfSec(blk.Head.Time/1e9) {
		node.SerialNum = 0
	} else {
		node.SerialNum = parentNode.SerialNum + 1
	}

	if node.SerialNum >= int64(blockNumPerWitness) {
		return errOutOfLimit
	}
	ok := p.verifyDB.Checkout(string(blk.HeadHash()))
	if !ok {
		p.verifyDB.Checkout(string(blk.Head.ParentHash))
		p.txPool.Lock()
		err := verifyBlock(blk, parentNode.Block, &node.GetParent().WitnessList, p.txPool, p.verifyDB, p.blockChain, replay)
		p.txPool.Release()
		if err != nil {
			ilog.Errorf("verify block failed, blockNum:%v, blockHash:%v. err=%v", blk.Head.Number, common.Base58Encode(blk.HeadHash()), err)
			p.blockCache.Del(node)
			return err
		}
		p.verifyDB.Commit(string(blk.HeadHash()))
	}
	p.blockCache.Link(node, replay)
	p.blockCache.UpdateLib(node)
	// After UpdateLib, the block head active witness list will be right
	// So AddLinkedNode need execute after UpdateLib
	p.txPool.AddLinkedNode(node)

	metricsConfirmedLength.Set(float64(p.blockCache.LinkedRoot().Head.Number), nil)

	if isWitness(p.account.ReadablePubkey(), p.blockCache.Head().Active()) {
		p.p2pService.ConnectBPs(p.blockCache.Head().NetID())
	} else {
		p.p2pService.ConnectBPs(nil)
	}

	if node.Head.Witness != p.account.ReadablePubkey() {
		ilog.Infof("Rec block - @%v id:%v..., num:%v, t:%v, txs:%v, confirmed:%v, et:%vms",
			node.SerialNum, node.Head.Witness[:10], node.Head.Number, node.Head.Time, len(node.Txs), p.blockCache.LinkedRoot().Head.Number, calculateTime(node.Block))
	}

	for child := range node.Children {
		p.addExistingBlock(child.Block, node, replay)
	}
	return nil
}
