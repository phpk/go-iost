// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: core/tx/tx.proto

package tx

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import crypto "github.com/iost-official/go-iost/crypto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ActionRaw struct {
	Contract             string   `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	ActionName           string   `protobuf:"bytes,2,opt,name=actionName,proto3" json:"actionName,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActionRaw) Reset()         { *m = ActionRaw{} }
func (m *ActionRaw) String() string { return proto.CompactTextString(m) }
func (*ActionRaw) ProtoMessage()    {}
func (*ActionRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_tx_8a326170601d5715, []int{0}
}
func (m *ActionRaw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActionRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActionRaw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ActionRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionRaw.Merge(dst, src)
}
func (m *ActionRaw) XXX_Size() int {
	return m.Size()
}
func (m *ActionRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionRaw.DiscardUnknown(m)
}

var xxx_messageInfo_ActionRaw proto.InternalMessageInfo

func (m *ActionRaw) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *ActionRaw) GetActionName() string {
	if m != nil {
		return m.ActionName
	}
	return ""
}

func (m *ActionRaw) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type TxRaw struct {
	Time                 int64                  `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Expiration           int64                  `protobuf:"varint,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
	GasLimit             int64                  `protobuf:"varint,3,opt,name=gasLimit,proto3" json:"gasLimit,omitempty"`
	GasPrice             int64                  `protobuf:"varint,4,opt,name=gasPrice,proto3" json:"gasPrice,omitempty"`
	Actions              []*ActionRaw           `protobuf:"bytes,5,rep,name=actions" json:"actions,omitempty"`
	Signers              [][]byte               `protobuf:"bytes,6,rep,name=signers" json:"signers,omitempty"`
	Signs                []*crypto.SignatureRaw `protobuf:"bytes,7,rep,name=signs" json:"signs,omitempty"`
	Publisher            *crypto.SignatureRaw   `protobuf:"bytes,8,opt,name=publisher" json:"publisher,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TxRaw) Reset()         { *m = TxRaw{} }
func (m *TxRaw) String() string { return proto.CompactTextString(m) }
func (*TxRaw) ProtoMessage()    {}
func (*TxRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_tx_8a326170601d5715, []int{1}
}
func (m *TxRaw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxRaw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TxRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxRaw.Merge(dst, src)
}
func (m *TxRaw) XXX_Size() int {
	return m.Size()
}
func (m *TxRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_TxRaw.DiscardUnknown(m)
}

var xxx_messageInfo_TxRaw proto.InternalMessageInfo

func (m *TxRaw) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *TxRaw) GetExpiration() int64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

func (m *TxRaw) GetGasLimit() int64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *TxRaw) GetGasPrice() int64 {
	if m != nil {
		return m.GasPrice
	}
	return 0
}

func (m *TxRaw) GetActions() []*ActionRaw {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *TxRaw) GetSigners() [][]byte {
	if m != nil {
		return m.Signers
	}
	return nil
}

func (m *TxRaw) GetSigns() []*crypto.SignatureRaw {
	if m != nil {
		return m.Signs
	}
	return nil
}

func (m *TxRaw) GetPublisher() *crypto.SignatureRaw {
	if m != nil {
		return m.Publisher
	}
	return nil
}

type ReceiptRaw struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiptRaw) Reset()         { *m = ReceiptRaw{} }
func (m *ReceiptRaw) String() string { return proto.CompactTextString(m) }
func (*ReceiptRaw) ProtoMessage()    {}
func (*ReceiptRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_tx_8a326170601d5715, []int{2}
}
func (m *ReceiptRaw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReceiptRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReceiptRaw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ReceiptRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiptRaw.Merge(dst, src)
}
func (m *ReceiptRaw) XXX_Size() int {
	return m.Size()
}
func (m *ReceiptRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiptRaw.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiptRaw proto.InternalMessageInfo

func (m *ReceiptRaw) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReceiptRaw) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type StatusRaw struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRaw) Reset()         { *m = StatusRaw{} }
func (m *StatusRaw) String() string { return proto.CompactTextString(m) }
func (*StatusRaw) ProtoMessage()    {}
func (*StatusRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_tx_8a326170601d5715, []int{3}
}
func (m *StatusRaw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StatusRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StatusRaw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *StatusRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRaw.Merge(dst, src)
}
func (m *StatusRaw) XXX_Size() int {
	return m.Size()
}
func (m *StatusRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRaw.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRaw proto.InternalMessageInfo

func (m *StatusRaw) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *StatusRaw) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type TxReceiptRaw struct {
	TxHash               []byte        `protobuf:"bytes,1,opt,name=txHash,proto3" json:"txHash,omitempty"`
	GasUsage             int64         `protobuf:"varint,2,opt,name=gasUsage,proto3" json:"gasUsage,omitempty"`
	Status               *StatusRaw    `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	SuccActionNum        int32         `protobuf:"varint,4,opt,name=succActionNum,proto3" json:"succActionNum,omitempty"`
	Receipts             []*ReceiptRaw `protobuf:"bytes,5,rep,name=receipts" json:"receipts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TxReceiptRaw) Reset()         { *m = TxReceiptRaw{} }
func (m *TxReceiptRaw) String() string { return proto.CompactTextString(m) }
func (*TxReceiptRaw) ProtoMessage()    {}
func (*TxReceiptRaw) Descriptor() ([]byte, []int) {
	return fileDescriptor_tx_8a326170601d5715, []int{4}
}
func (m *TxReceiptRaw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxReceiptRaw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxReceiptRaw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TxReceiptRaw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxReceiptRaw.Merge(dst, src)
}
func (m *TxReceiptRaw) XXX_Size() int {
	return m.Size()
}
func (m *TxReceiptRaw) XXX_DiscardUnknown() {
	xxx_messageInfo_TxReceiptRaw.DiscardUnknown(m)
}

var xxx_messageInfo_TxReceiptRaw proto.InternalMessageInfo

func (m *TxReceiptRaw) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *TxReceiptRaw) GetGasUsage() int64 {
	if m != nil {
		return m.GasUsage
	}
	return 0
}

func (m *TxReceiptRaw) GetStatus() *StatusRaw {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *TxReceiptRaw) GetSuccActionNum() int32 {
	if m != nil {
		return m.SuccActionNum
	}
	return 0
}

func (m *TxReceiptRaw) GetReceipts() []*ReceiptRaw {
	if m != nil {
		return m.Receipts
	}
	return nil
}

func init() {
	proto.RegisterType((*ActionRaw)(nil), "tx.ActionRaw")
	proto.RegisterType((*TxRaw)(nil), "tx.TxRaw")
	proto.RegisterType((*ReceiptRaw)(nil), "tx.ReceiptRaw")
	proto.RegisterType((*StatusRaw)(nil), "tx.StatusRaw")
	proto.RegisterType((*TxReceiptRaw)(nil), "tx.TxReceiptRaw")
}
func (m *ActionRaw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActionRaw) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Contract) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTx(dAtA, i, uint64(len(m.Contract)))
		i += copy(dAtA[i:], m.Contract)
	}
	if len(m.ActionName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTx(dAtA, i, uint64(len(m.ActionName)))
		i += copy(dAtA[i:], m.ActionName)
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TxRaw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxRaw) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Time != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTx(dAtA, i, uint64(m.Time))
	}
	if m.Expiration != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTx(dAtA, i, uint64(m.Expiration))
	}
	if m.GasLimit != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintTx(dAtA, i, uint64(m.GasLimit))
	}
	if m.GasPrice != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintTx(dAtA, i, uint64(m.GasPrice))
	}
	if len(m.Actions) > 0 {
		for _, msg := range m.Actions {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintTx(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Signers) > 0 {
		for _, b := range m.Signers {
			dAtA[i] = 0x32
			i++
			i = encodeVarintTx(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	if len(m.Signs) > 0 {
		for _, msg := range m.Signs {
			dAtA[i] = 0x3a
			i++
			i = encodeVarintTx(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Publisher != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintTx(dAtA, i, uint64(m.Publisher.Size()))
		n1, err := m.Publisher.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ReceiptRaw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReceiptRaw) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTx(dAtA, i, uint64(m.Type))
	}
	if len(m.Content) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTx(dAtA, i, uint64(len(m.Content)))
		i += copy(dAtA[i:], m.Content)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *StatusRaw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatusRaw) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTx(dAtA, i, uint64(m.Code))
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTx(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TxReceiptRaw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxReceiptRaw) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.TxHash) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTx(dAtA, i, uint64(len(m.TxHash)))
		i += copy(dAtA[i:], m.TxHash)
	}
	if m.GasUsage != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintTx(dAtA, i, uint64(m.GasUsage))
	}
	if m.Status != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTx(dAtA, i, uint64(m.Status.Size()))
		n2, err := m.Status.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.SuccActionNum != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintTx(dAtA, i, uint64(m.SuccActionNum))
	}
	if len(m.Receipts) > 0 {
		for _, msg := range m.Receipts {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintTx(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ActionRaw) Size() (n int) {
	var l int
	_ = l
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ActionName)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TxRaw) Size() (n int) {
	var l int
	_ = l
	if m.Time != 0 {
		n += 1 + sovTx(uint64(m.Time))
	}
	if m.Expiration != 0 {
		n += 1 + sovTx(uint64(m.Expiration))
	}
	if m.GasLimit != 0 {
		n += 1 + sovTx(uint64(m.GasLimit))
	}
	if m.GasPrice != 0 {
		n += 1 + sovTx(uint64(m.GasPrice))
	}
	if len(m.Actions) > 0 {
		for _, e := range m.Actions {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.Signers) > 0 {
		for _, b := range m.Signers {
			l = len(b)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.Signs) > 0 {
		for _, e := range m.Signs {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.Publisher != nil {
		l = m.Publisher.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ReceiptRaw) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovTx(uint64(m.Type))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StatusRaw) Size() (n int) {
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovTx(uint64(m.Code))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TxReceiptRaw) Size() (n int) {
	var l int
	_ = l
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.GasUsage != 0 {
		n += 1 + sovTx(uint64(m.GasUsage))
	}
	if m.Status != nil {
		l = m.Status.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.SuccActionNum != 0 {
		n += 1 + sovTx(uint64(m.SuccActionNum))
	}
	if len(m.Receipts) > 0 {
		for _, e := range m.Receipts {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTx(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ActionRaw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ActionRaw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActionRaw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ActionName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TxRaw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TxRaw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxRaw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			m.Expiration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expiration |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLimit", wireType)
			}
			m.GasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasLimit |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrice", wireType)
			}
			m.GasPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasPrice |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Actions = append(m.Actions, &ActionRaw{})
			if err := m.Actions[len(m.Actions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, make([]byte, postIndex-iNdEx))
			copy(m.Signers[len(m.Signers)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signs = append(m.Signs, &crypto.SignatureRaw{})
			if err := m.Signs[len(m.Signs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Publisher", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Publisher == nil {
				m.Publisher = &crypto.SignatureRaw{}
			}
			if err := m.Publisher.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReceiptRaw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReceiptRaw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReceiptRaw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatusRaw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatusRaw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusRaw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TxReceiptRaw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TxReceiptRaw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxReceiptRaw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = append(m.TxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TxHash == nil {
				m.TxHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsage", wireType)
			}
			m.GasUsage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsage |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Status == nil {
				m.Status = &StatusRaw{}
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuccActionNum", wireType)
			}
			m.SuccActionNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SuccActionNum |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receipts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receipts = append(m.Receipts, &ReceiptRaw{})
			if err := m.Receipts[len(m.Receipts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTx
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTx(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTx = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("core/tx/tx.proto", fileDescriptor_tx_8a326170601d5715) }

var fileDescriptor_tx_8a326170601d5715 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0x26, 0xcd, 0x66, 0x7f, 0x86, 0x2d, 0xaa, 0x2c, 0x84, 0xac, 0x1e, 0xa2, 0x55, 0x04, 0x62,
	0x55, 0xa9, 0x89, 0x54, 0x0e, 0x08, 0x6e, 0x70, 0xe2, 0x80, 0x2a, 0xe4, 0x96, 0x13, 0x27, 0xaf,
	0xeb, 0x66, 0x2d, 0x35, 0xeb, 0xc8, 0x9e, 0x88, 0xf4, 0x45, 0x10, 0xef, 0xc2, 0x0b, 0x70, 0xe4,
	0x11, 0xd0, 0xf2, 0x22, 0xc8, 0x93, 0x9f, 0x2e, 0x07, 0x6e, 0xf3, 0xcd, 0x37, 0xf3, 0xf9, 0xcb,
	0x37, 0x81, 0x13, 0x65, 0x9d, 0x2e, 0xb0, 0x2d, 0xb0, 0xcd, 0x6b, 0x67, 0xd1, 0xb2, 0x23, 0x6c,
	0x4f, 0x5f, 0x97, 0x06, 0xb7, 0xcd, 0x26, 0x57, 0xb6, 0x2a, 0x8c, 0xf5, 0x78, 0x6e, 0x6f, 0x6f,
	0x8d, 0x32, 0xf2, 0xae, 0x28, 0xed, 0x79, 0x68, 0x14, 0xca, 0xdd, 0xd7, 0x68, 0x0b, 0x6f, 0xca,
	0x9d, 0xc4, 0xc6, 0xe9, 0x6e, 0x39, 0xfb, 0x02, 0x8b, 0x77, 0x0a, 0x8d, 0xdd, 0x09, 0xf9, 0x95,
	0x9d, 0xc2, 0x5c, 0xd9, 0x1d, 0x3a, 0xa9, 0x90, 0x47, 0xab, 0x68, 0xbd, 0x10, 0x23, 0x66, 0x29,
	0x80, 0xa4, 0xc1, 0x4b, 0x59, 0x69, 0x7e, 0x44, 0xec, 0x41, 0x87, 0x31, 0x98, 0xdc, 0x48, 0x94,
	0x3c, 0x26, 0x86, 0xea, 0xec, 0xdb, 0x11, 0x24, 0xd7, 0x6d, 0x50, 0x66, 0x30, 0x41, 0x53, 0x69,
	0x52, 0x8d, 0x05, 0xd5, 0x41, 0x51, 0xb7, 0xb5, 0x71, 0x32, 0x68, 0x90, 0x62, 0x2c, 0x0e, 0x3a,
	0xc1, 0x4d, 0x29, 0xfd, 0x47, 0x53, 0x19, 0x24, 0xd5, 0x58, 0x8c, 0xb8, 0xe7, 0x3e, 0x39, 0xa3,
	0x34, 0x9f, 0x8c, 0x1c, 0x61, 0xf6, 0x12, 0x66, 0x9d, 0x2f, 0xcf, 0x93, 0x55, 0xbc, 0x7e, 0x7c,
	0x71, 0x9c, 0x63, 0x9b, 0x8f, 0x5f, 0x29, 0x06, 0x96, 0x71, 0x98, 0x85, 0x38, 0xb4, 0xf3, 0x7c,
	0xba, 0x8a, 0xd7, 0x4b, 0x31, 0x40, 0x76, 0x06, 0x49, 0x28, 0x3d, 0x9f, 0x91, 0xc0, 0xd3, 0xbc,
	0x4b, 0x2f, 0xbf, 0x1a, 0xd2, 0x0b, 0x3a, 0xdd, 0x08, 0xbb, 0x80, 0x45, 0xdd, 0x6c, 0xee, 0x8c,
	0xdf, 0x6a, 0xc7, 0xe7, 0xab, 0xe8, 0xbf, 0xf3, 0x0f, 0x63, 0xd9, 0x5b, 0x00, 0xa1, 0x95, 0x36,
	0x35, 0x0e, 0xe1, 0xdc, 0xd7, 0x5d, 0x38, 0x89, 0xa0, 0x3a, 0x78, 0x0b, 0xd1, 0xeb, 0x1d, 0xf6,
	0x59, 0x0f, 0x30, 0x7b, 0x03, 0x8b, 0x2b, 0x94, 0xd8, 0xf8, 0x7e, 0x55, 0xd9, 0x9b, 0x71, 0x35,
	0xd4, 0x61, 0xb5, 0xd2, 0xde, 0xcb, 0x72, 0x38, 0xd3, 0x00, 0xb3, 0x1f, 0x11, 0x2c, 0xaf, 0xdb,
	0x83, 0x97, 0x9f, 0xc1, 0x14, 0xdb, 0x0f, 0xd2, 0x6f, 0x49, 0x60, 0x29, 0x7a, 0xd4, 0xc7, 0xfb,
	0x79, 0xd4, 0xe8, 0xe2, 0x25, 0xcc, 0x5e, 0xc0, 0xd4, 0xd3, 0xfb, 0x74, 0x94, 0x3e, 0xdd, 0xd1,
	0x91, 0xe8, 0x49, 0xf6, 0x1c, 0x8e, 0x7d, 0xa3, 0x54, 0x17, 0xfb, 0x65, 0x53, 0xd1, 0x99, 0x12,
	0xf1, 0x6f, 0x93, 0x9d, 0xc1, 0xdc, 0x75, 0x76, 0x86, 0x63, 0x3d, 0x09, 0x72, 0x0f, 0x16, 0xc5,
	0xc8, 0xbf, 0x3f, 0xf9, 0xb9, 0x4f, 0xa3, 0x5f, 0xfb, 0x34, 0xfa, 0xbd, 0x4f, 0xa3, 0xef, 0x7f,
	0xd2, 0x47, 0x9b, 0x29, 0xfd, 0xc3, 0xaf, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xb3, 0xc0,
	0xfe, 0x14, 0x03, 0x00, 0x00,
}
