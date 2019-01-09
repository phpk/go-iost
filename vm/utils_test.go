package vm

import (
	"os"
	"testing"

	"encoding/json"

	"github.com/golang/mock/gomock"
	"github.com/iost-official/go-iost/account"
	"github.com/iost-official/go-iost/common"
	"github.com/iost-official/go-iost/core/tx"
	"github.com/iost-official/go-iost/crypto"
	"github.com/iost-official/go-iost/db"
	"github.com/iost-official/go-iost/ilog"
	"github.com/iost-official/go-iost/vm/database"
	"github.com/iost-official/go-iost/vm/host"
	"github.com/iost-official/go-iost/vm/native"
)

var testKps = make([]*account.KeyPair, 0)
func init() {
	privKeys := []string{
		"EhNiaU4DzUmjCrvynV3gaUeuj2VjB1v2DCmbGD5U2nSE",
		"8dJ9YKovJ5E7hkebAQaScaG1BA8snRUHPUbVcArcTVq6",
		"7CnwT7BXkEFAVx6QZqC7gkDhQwbvC3d2CkMZvXHZdDMN",
		"Htarc5Sp4trjqY4WrTLtZ85CF6qx87v7CRwtV4RRGnbF",
		"Bk8bAyG4VLBcrsoRErPuQGhwCy4C1VxfKE4jjX9oLhv",
		"546aCDG9igGgZqVZeybajaorP5ZeF9ghLu2oLncXk3d6",
		"DXNYRwG7dRFkbWzMNEbKfBhuS8Yn51x9J6XuTdNwB11M",
		"AG8uECmAwFis8uxTdWqcgGD9tGDwoP6CxqhkhpuCdSeC",
		"GJt5WSSv5WZi1axd3qkb1vLEfxCEgKGupcXf45b5tERU",
		"7U3uwEeGc2TF3Xde2oT66eTx1Uw15qRqYuTnMd3NNjai",
	}
	for _, k := range privKeys {
		kp, err := account.NewKeyPair(common.Base58Decode(k), crypto.Secp256k1)
		if err != nil {
			panic(err)
		}
		testKps = append(testKps, kp)
	}
}

var systemContract = native.SystemABI()

func ininit(t *testing.T) (*database.Visitor, db.MVCCDB) {
	mvccdb, err := db.NewMVCCDB("mvcc")
	if err != nil {
		t.Fatal(err)
	}
	//mvccdb := replaceDB(t)
	vi := database.NewVisitor(0, mvccdb)
	vi.SetTokenBalance("iost", testKps[0].ReadablePubkey(), 1000000)
	vi.SetContract(systemContract)
	vi.Commit()
	return vi, mvccdb
}

func closeMVCCDB(m db.MVCCDB) {
	m.Close()
	os.RemoveAll("mvcc")
}

func TestCheckPublisher(t *testing.T) {
	tr := tx.NewTx([]*tx.Action{{
		"system.iost",
		"Transfer",
		"[]",
	}}, []string{}, 10000, 1, 10000, 0)

	kp := testKps[0]
	t2, err := tx.SignTx(tr, "a", []*account.KeyPair{kp})
	if err != nil {
		t.Fatal(err)
	}

	ctl := gomock.NewController(t)

	k0 := testKps[0].ReadablePubkey()
	a := account.NewInitAccount("a", k0, k0)
	ax, err := json.Marshal(a)
	if err != nil {
		t.Fatal(err)
	}
	mock := database.NewMockIMultiValue(ctl)
	mock.EXPECT().Get("state", "m-auth.iost-auth-a").Return("s"+string(ax), nil)

	k1 := testKps[1].ReadablePubkey()
	b := account.NewInitAccount("b", k1, k1)
	bx, err := json.Marshal(b)
	if err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().Get("state", "m-auth.iost-auth-b").Return("s"+string(bx), nil)

	authList := make(map[string]int)
	authList[kp.ReadablePubkey()] = 2
	h := host.NewHost(host.NewContext(nil), database.NewVisitor(0, mock), nil, ilog.DefaultLogger())
	h.Context().Set("auth_list", authList)
	err = h.CheckPublisher(t2)
	if err != nil {
		t.Fatal(err)
	}

	t2.Publisher = "b"
	err = h.CheckPublisher(t2)
	if err == nil {
		t.Fatal(err)
	}
}

func TestCheckSigners(t *testing.T) {
	ctl := gomock.NewController(t)
	mock := database.NewMockIMultiValue(ctl)


	k0 := testKps[0].ReadablePubkey()
	a := account.NewInitAccount("a", k0, k0)
	ax, err := json.Marshal(a)
	if err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().Get("state", "m-auth.iost-auth-a").AnyTimes().Return("s"+string(ax), nil)


	k1 := testKps[1].ReadablePubkey()
	b := account.NewInitAccount("b", k1, k1)
	bx, err := json.Marshal(b)
	if err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().Get("state", "m-auth.iost-auth-b").AnyTimes().Return("s"+string(bx), nil)

	kp := testKps[0]
	kp2 := testKps[1]

	tr := tx.NewTx([]*tx.Action{{
		"system.iost",
		"Transfer",
		"[]",
	}}, []string{"a@acitve", "b@acitve"}, 10000, 1, 10000, 0)

	sig1, err := tx.SignTxContent(tr, "a", kp)
	if err != nil {
		t.Fatal(err)
	}
	tr.Signs = append(tr.Signs, sig1)

	authList := make(map[string]int)
	authList[kp.ReadablePubkey()] = 1
	h := host.NewHost(host.NewContext(nil), database.NewVisitor(0, mock), nil, ilog.DefaultLogger())
	h.Context().Set("auth_list", authList)
	err = h.CheckSigners(tr)
	if err == nil {
		t.Fatal(err)
	}

	sig2, err := tx.SignTxContent(tr, "b", kp2)
	if err != nil {
		t.Fatal(err)
	}
	tr.Signs = append(tr.Signs, sig2)
	authList[kp2.ReadablePubkey()] = 1
	h.Context().Set("auth_list", authList)

	err = h.CheckSigners(tr)
	if err != nil {
		t.Fatal(err)
	}

}
