// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: mwebd.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mwebd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mwebd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_mwebd_proto_rawDescGZIP(), []int{0}
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The height of the latest block.
	BlockHeaderHeight int32 `protobuf:"varint,1,opt,name=block_header_height,json=blockHeaderHeight,proto3" json:"block_header_height,omitempty"`
	// The height of the latest MWEB header.
	MwebHeaderHeight int32 `protobuf:"varint,2,opt,name=mweb_header_height,json=mwebHeaderHeight,proto3" json:"mweb_header_height,omitempty"`
	// The height at which the MWEB utxo set is synced to.
	MwebUtxosHeight int32 `protobuf:"varint,3,opt,name=mweb_utxos_height,json=mwebUtxosHeight,proto3" json:"mweb_utxos_height,omitempty"`
	// The timestamp of the latest block.
	BlockTime uint32 `protobuf:"varint,4,opt,name=block_time,json=blockTime,proto3" json:"block_time,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mwebd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mwebd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_mwebd_proto_rawDescGZIP(), []int{1}
}

func (x *StatusResponse) GetBlockHeaderHeight() int32 {
	if x != nil {
		return x.BlockHeaderHeight
	}
	return 0
}

func (x *StatusResponse) GetMwebHeaderHeight() int32 {
	if x != nil {
		return x.MwebHeaderHeight
	}
	return 0
}

func (x *StatusResponse) GetMwebUtxosHeight() int32 {
	if x != nil {
		return x.MwebUtxosHeight
	}
	return 0
}

func (x *StatusResponse) GetBlockTime() uint32 {
	if x != nil {
		return x.BlockTime
	}
	return 0
}

type UtxosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The block height from which to start fetching utxos from.
	// After all mined utxos have been streamed, unconfirmed and
	// newly confirmed utxos will also be streamed. If this is set
	// to 0 then all utxos belonging to the account will be fetched.
	FromHeight int32 `protobuf:"varint,1,opt,name=from_height,json=fromHeight,proto3" json:"from_height,omitempty"`
	// The scan secret or view key represents the account for
	// which utxos should be streamed.
	ScanSecret []byte `protobuf:"bytes,2,opt,name=scan_secret,json=scanSecret,proto3" json:"scan_secret,omitempty"`
}

func (x *UtxosRequest) Reset() {
	*x = UtxosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mwebd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UtxosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UtxosRequest) ProtoMessage() {}

func (x *UtxosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mwebd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UtxosRequest.ProtoReflect.Descriptor instead.
func (*UtxosRequest) Descriptor() ([]byte, []int) {
	return file_mwebd_proto_rawDescGZIP(), []int{2}
}

func (x *UtxosRequest) GetFromHeight() int32 {
	if x != nil {
		return x.FromHeight
	}
	return 0
}

func (x *UtxosRequest) GetScanSecret() []byte {
	if x != nil {
		return x.ScanSecret
	}
	return nil
}

type Utxo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The block height of the utxo, or 0 for unconfirmed.
	Height int32 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// The value of the utxo in litoshis.
	Value uint64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	// The MWEB address that the utxo was received on.
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	// The output ID. This functions like a transaction hash,
	// but is unique to every utxo.
	OutputId string `protobuf:"bytes,4,opt,name=output_id,json=outputId,proto3" json:"output_id,omitempty"`
	// The timestamp of the block the utxo was mined in.
	BlockTime uint32 `protobuf:"varint,5,opt,name=block_time,json=blockTime,proto3" json:"block_time,omitempty"`
}

func (x *Utxo) Reset() {
	*x = Utxo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mwebd_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Utxo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Utxo) ProtoMessage() {}

func (x *Utxo) ProtoReflect() protoreflect.Message {
	mi := &file_mwebd_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Utxo.ProtoReflect.Descriptor instead.
func (*Utxo) Descriptor() ([]byte, []int) {
	return file_mwebd_proto_rawDescGZIP(), []int{3}
}

func (x *Utxo) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Utxo) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Utxo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Utxo) GetOutputId() string {
	if x != nil {
		return x.OutputId
	}
	return ""
}

func (x *Utxo) GetBlockTime() uint32 {
	if x != nil {
		return x.BlockTime
	}
	return 0
}

type AddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The starting index of the range.
	FromIndex uint32 `protobuf:"varint,1,opt,name=from_index,json=fromIndex,proto3" json:"from_index,omitempty"`
	// The ending index of the range. The result will contain all
	// addresses up to but not including this index.
	ToIndex uint32 `protobuf:"varint,2,opt,name=to_index,json=toIndex,proto3" json:"to_index,omitempty"`
	// The scan secret or view key represents the account for
	// which addresses should be returned.
	ScanSecret []byte `protobuf:"bytes,3,opt,name=scan_secret,json=scanSecret,proto3" json:"scan_secret,omitempty"`
	// The public key of the spend secret for the account. The spend
	// key is required for spending utxos but is also required
	// for generating addresses.
	SpendPubkey []byte `protobuf:"bytes,4,opt,name=spend_pubkey,json=spendPubkey,proto3" json:"spend_pubkey,omitempty"`
}

func (x *AddressRequest) Reset() {
	*x = AddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mwebd_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressRequest) ProtoMessage() {}

func (x *AddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mwebd_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressRequest.ProtoReflect.Descriptor instead.
func (*AddressRequest) Descriptor() ([]byte, []int) {
	return file_mwebd_proto_rawDescGZIP(), []int{4}
}

func (x *AddressRequest) GetFromIndex() uint32 {
	if x != nil {
		return x.FromIndex
	}
	return 0
}

func (x *AddressRequest) GetToIndex() uint32 {
	if x != nil {
		return x.ToIndex
	}
	return 0
}

func (x *AddressRequest) GetScanSecret() []byte {
	if x != nil {
		return x.ScanSecret
	}
	return nil
}

func (x *AddressRequest) GetSpendPubkey() []byte {
	if x != nil {
		return x.SpendPubkey
	}
	return nil
}

type AddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An array of MWEB addresses within the requested range.
	Address []string `protobuf:"bytes,1,rep,name=address,proto3" json:"address,omitempty"`
}

func (x *AddressResponse) Reset() {
	*x = AddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mwebd_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressResponse) ProtoMessage() {}

func (x *AddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mwebd_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressResponse.ProtoReflect.Descriptor instead.
func (*AddressResponse) Descriptor() ([]byte, []int) {
	return file_mwebd_proto_rawDescGZIP(), []int{5}
}

func (x *AddressResponse) GetAddress() []string {
	if x != nil {
		return x.Address
	}
	return nil
}

type LedgerApdu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *LedgerApdu) Reset() {
	*x = LedgerApdu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mwebd_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LedgerApdu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LedgerApdu) ProtoMessage() {}

func (x *LedgerApdu) ProtoReflect() protoreflect.Message {
	mi := &file_mwebd_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LedgerApdu.ProtoReflect.Descriptor instead.
func (*LedgerApdu) Descriptor() ([]byte, []int) {
	return file_mwebd_proto_rawDescGZIP(), []int{6}
}

func (x *LedgerApdu) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SpentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An array of output IDs to perform checks for.
	OutputId []string `protobuf:"bytes,1,rep,name=output_id,json=outputId,proto3" json:"output_id,omitempty"`
}

func (x *SpentRequest) Reset() {
	*x = SpentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mwebd_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpentRequest) ProtoMessage() {}

func (x *SpentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mwebd_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpentRequest.ProtoReflect.Descriptor instead.
func (*SpentRequest) Descriptor() ([]byte, []int) {
	return file_mwebd_proto_rawDescGZIP(), []int{7}
}

func (x *SpentRequest) GetOutputId() []string {
	if x != nil {
		return x.OutputId
	}
	return nil
}

type SpentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An array of the output IDs that were not found in the
	// unspent set. This means that the outputs are either
	// unconfirmed or were spent.
	OutputId []string `protobuf:"bytes,1,rep,name=output_id,json=outputId,proto3" json:"output_id,omitempty"`
}

func (x *SpentResponse) Reset() {
	*x = SpentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mwebd_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpentResponse) ProtoMessage() {}

func (x *SpentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mwebd_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpentResponse.ProtoReflect.Descriptor instead.
func (*SpentResponse) Descriptor() ([]byte, []int) {
	return file_mwebd_proto_rawDescGZIP(), []int{8}
}

func (x *SpentResponse) GetOutputId() []string {
	if x != nil {
		return x.OutputId
	}
	return nil
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The raw bytes of the serialized transaction. This will be
	// a template where the non-MWEB inputs will remain unchanged,
	// and the MWEB inputs are specified by TxIns with the outpoint
	// hash set to the output ID of the utxo being spent, and the
	// outpoint index set to the index of the address that the utxo
	// was received on. MWEB outputs are specified by TxOuts with
	// the script pubkey set to the serialized scan and spend pubkeys
	// of the destination MWEB address. Any non-MWEB outputs will be
	// transformed into MWEB peg-outs. If the transaction doesn't
	// contain any MWEB i/o then the result will be unchanged.
	RawTx []byte `protobuf:"bytes,1,opt,name=raw_tx,json=rawTx,proto3" json:"raw_tx,omitempty"`
	// The scan secret or view key represents the account that
	// the utxos being spent belong to.
	ScanSecret []byte `protobuf:"bytes,2,opt,name=scan_secret,json=scanSecret,proto3" json:"scan_secret,omitempty"`
	// The spend secret is the private key necessary for spending
	// the utxos belonging to the account.
	SpendSecret []byte `protobuf:"bytes,3,opt,name=spend_secret,json=spendSecret,proto3" json:"spend_secret,omitempty"`
	// The fee rate per KB in litoshis.
	FeeRatePerKb uint64 `protobuf:"varint,4,opt,name=fee_rate_per_kb,json=feeRatePerKb,proto3" json:"fee_rate_per_kb,omitempty"`
	// Whether to skip MWEB transaction creation. This is useful
	// for fee estimation.
	DryRun bool `protobuf:"varint,5,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mwebd_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mwebd_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_mwebd_proto_rawDescGZIP(), []int{9}
}

func (x *CreateRequest) GetRawTx() []byte {
	if x != nil {
		return x.RawTx
	}
	return nil
}

func (x *CreateRequest) GetScanSecret() []byte {
	if x != nil {
		return x.ScanSecret
	}
	return nil
}

func (x *CreateRequest) GetSpendSecret() []byte {
	if x != nil {
		return x.SpendSecret
	}
	return nil
}

func (x *CreateRequest) GetFeeRatePerKb() uint64 {
	if x != nil {
		return x.FeeRatePerKb
	}
	return 0
}

func (x *CreateRequest) GetDryRun() bool {
	if x != nil {
		return x.DryRun
	}
	return false
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The raw bytes of the serialized transaction. It will contain
	// a single TxOut representing the peg-in required. From this
	// it is possible to determine the addtional fee that was added
	// by the MWEB transaction.
	RawTx []byte `protobuf:"bytes,1,opt,name=raw_tx,json=rawTx,proto3" json:"raw_tx,omitempty"`
	// The output IDs of any utxos created by the transaction,
	// in the same order as in the template.
	OutputId []string `protobuf:"bytes,2,rep,name=output_id,json=outputId,proto3" json:"output_id,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mwebd_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mwebd_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_mwebd_proto_rawDescGZIP(), []int{10}
}

func (x *CreateResponse) GetRawTx() []byte {
	if x != nil {
		return x.RawTx
	}
	return nil
}

func (x *CreateResponse) GetOutputId() []string {
	if x != nil {
		return x.OutputId
	}
	return nil
}

type BroadcastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The raw bytes of the serialized transaction.
	RawTx []byte `protobuf:"bytes,1,opt,name=raw_tx,json=rawTx,proto3" json:"raw_tx,omitempty"`
}

func (x *BroadcastRequest) Reset() {
	*x = BroadcastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mwebd_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastRequest) ProtoMessage() {}

func (x *BroadcastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mwebd_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastRequest.ProtoReflect.Descriptor instead.
func (*BroadcastRequest) Descriptor() ([]byte, []int) {
	return file_mwebd_proto_rawDescGZIP(), []int{11}
}

func (x *BroadcastRequest) GetRawTx() []byte {
	if x != nil {
		return x.RawTx
	}
	return nil
}

type BroadcastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The transaction ID.
	Txid string `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
}

func (x *BroadcastResponse) Reset() {
	*x = BroadcastResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mwebd_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastResponse) ProtoMessage() {}

func (x *BroadcastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mwebd_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastResponse.ProtoReflect.Descriptor instead.
func (*BroadcastResponse) Descriptor() ([]byte, []int) {
	return file_mwebd_proto_rawDescGZIP(), []int{12}
}

func (x *BroadcastResponse) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

var File_mwebd_proto protoreflect.FileDescriptor

var file_mwebd_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x77, 0x65, 0x62, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f, 0x0a,
	0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xb9,
	0x01, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x13, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x77, 0x65, 0x62, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6d,
	0x77, 0x65, 0x62, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x2a, 0x0a, 0x11, 0x6d, 0x77, 0x65, 0x62, 0x5f, 0x75, 0x74, 0x78, 0x6f, 0x73, 0x5f, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6d, 0x77, 0x65, 0x62,
	0x55, 0x74, 0x78, 0x6f, 0x73, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x50, 0x0a, 0x0c, 0x55, 0x74,
	0x78, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x66, 0x72, 0x6f, 0x6d, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x63, 0x61, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x73, 0x63, 0x61, 0x6e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x8a, 0x01, 0x0a,
	0x04, 0x55, 0x74, 0x78, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x8e, 0x01, 0x0a, 0x0e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x6f, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74,
	0x6f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x63, 0x61,
	0x6e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x70, 0x65, 0x6e, 0x64,
	0x5f, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x73,
	0x70, 0x65, 0x6e, 0x64, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x22, 0x2b, 0x0a, 0x0f, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x20, 0x0a, 0x0a, 0x4c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x41, 0x70, 0x64, 0x75, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2b, 0x0a, 0x0c, 0x53, 0x70, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x0d, 0x53, 0x70, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x49, 0x64, 0x22, 0xaa, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x61, 0x77, 0x5f, 0x74, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x72, 0x61, 0x77, 0x54, 0x78, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0a, 0x73, 0x63, 0x61, 0x6e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x12, 0x25, 0x0a, 0x0f, 0x66, 0x65, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x65,
	0x72, 0x5f, 0x6b, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x66, 0x65, 0x65, 0x52,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x4b, 0x62, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x72, 0x79, 0x5f,
	0x72, 0x75, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64, 0x72, 0x79, 0x52, 0x75,
	0x6e, 0x22, 0x44, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x61, 0x77, 0x5f, 0x74, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x72, 0x61, 0x77, 0x54, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x72,
	0x61, 0x77, 0x5f, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x72, 0x61, 0x77,
	0x54, 0x78, 0x22, 0x27, 0x0a, 0x11, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x69, 0x64, 0x32, 0xb4, 0x02, 0x0a, 0x03,
	0x52, 0x70, 0x63, 0x12, 0x29, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x05, 0x55, 0x74, 0x78, 0x6f, 0x73, 0x12, 0x0d, 0x2e, 0x55, 0x74, 0x78, 0x6f, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x55, 0x74, 0x78, 0x6f, 0x30, 0x01, 0x12,
	0x2e, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x05, 0x53, 0x70, 0x65, 0x6e, 0x74, 0x12, 0x0d, 0x2e, 0x53, 0x70, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x53, 0x70, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x0e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x0e, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x0b, 0x2e, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x41, 0x70, 0x64,
	0x75, 0x1a, 0x0b, 0x2e, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x41, 0x70, 0x64, 0x75, 0x12, 0x32,
	0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6c, 0x74, 0x63, 0x6d, 0x77, 0x65, 0x62, 0x2f, 0x6d, 0x77, 0x65, 0x62, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mwebd_proto_rawDescOnce sync.Once
	file_mwebd_proto_rawDescData = file_mwebd_proto_rawDesc
)

func file_mwebd_proto_rawDescGZIP() []byte {
	file_mwebd_proto_rawDescOnce.Do(func() {
		file_mwebd_proto_rawDescData = protoimpl.X.CompressGZIP(file_mwebd_proto_rawDescData)
	})
	return file_mwebd_proto_rawDescData
}

var file_mwebd_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_mwebd_proto_goTypes = []any{
	(*StatusRequest)(nil),     // 0: StatusRequest
	(*StatusResponse)(nil),    // 1: StatusResponse
	(*UtxosRequest)(nil),      // 2: UtxosRequest
	(*Utxo)(nil),              // 3: Utxo
	(*AddressRequest)(nil),    // 4: AddressRequest
	(*AddressResponse)(nil),   // 5: AddressResponse
	(*LedgerApdu)(nil),        // 6: LedgerApdu
	(*SpentRequest)(nil),      // 7: SpentRequest
	(*SpentResponse)(nil),     // 8: SpentResponse
	(*CreateRequest)(nil),     // 9: CreateRequest
	(*CreateResponse)(nil),    // 10: CreateResponse
	(*BroadcastRequest)(nil),  // 11: BroadcastRequest
	(*BroadcastResponse)(nil), // 12: BroadcastResponse
}
var file_mwebd_proto_depIdxs = []int32{
	0,  // 0: Rpc.Status:input_type -> StatusRequest
	2,  // 1: Rpc.Utxos:input_type -> UtxosRequest
	4,  // 2: Rpc.Addresses:input_type -> AddressRequest
	7,  // 3: Rpc.Spent:input_type -> SpentRequest
	9,  // 4: Rpc.Create:input_type -> CreateRequest
	6,  // 5: Rpc.LedgerExchange:input_type -> LedgerApdu
	11, // 6: Rpc.Broadcast:input_type -> BroadcastRequest
	1,  // 7: Rpc.Status:output_type -> StatusResponse
	3,  // 8: Rpc.Utxos:output_type -> Utxo
	5,  // 9: Rpc.Addresses:output_type -> AddressResponse
	8,  // 10: Rpc.Spent:output_type -> SpentResponse
	10, // 11: Rpc.Create:output_type -> CreateResponse
	6,  // 12: Rpc.LedgerExchange:output_type -> LedgerApdu
	12, // 13: Rpc.Broadcast:output_type -> BroadcastResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_mwebd_proto_init() }
func file_mwebd_proto_init() {
	if File_mwebd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mwebd_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*StatusRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mwebd_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*StatusResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mwebd_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UtxosRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mwebd_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Utxo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mwebd_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AddressRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mwebd_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*AddressResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mwebd_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*LedgerApdu); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mwebd_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*SpentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mwebd_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*SpentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mwebd_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mwebd_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*CreateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mwebd_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*BroadcastRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mwebd_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*BroadcastResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mwebd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mwebd_proto_goTypes,
		DependencyIndexes: file_mwebd_proto_depIdxs,
		MessageInfos:      file_mwebd_proto_msgTypes,
	}.Build()
	File_mwebd_proto = out.File
	file_mwebd_proto_rawDesc = nil
	file_mwebd_proto_goTypes = nil
	file_mwebd_proto_depIdxs = nil
}
