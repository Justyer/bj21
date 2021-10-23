// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: api/bj21/v1/bj21.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd  string `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Text []byte `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bj21_v1_bj21_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_api_bj21_v1_bj21_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_api_bj21_v1_bj21_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *Message) GetText() []byte {
	if x != nil {
		return x.Text
	}
	return nil
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bj21_v1_bj21_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_api_bj21_v1_bj21_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_api_bj21_v1_bj21_proto_rawDescGZIP(), []int{1}
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Table struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Seq  string  `protobuf:"bytes,2,opt,name=seq,proto3" json:"seq,omitempty"`
	P1   *Player `protobuf:"bytes,3,opt,name=p1,proto3" json:"p1,omitempty"`
	P2   *Player `protobuf:"bytes,4,opt,name=p2,proto3" json:"p2,omitempty"`
	Me   int32   `protobuf:"varint,5,opt,name=me,proto3" json:"me,omitempty"`
}

func (x *Table) Reset() {
	*x = Table{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bj21_v1_bj21_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Table) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Table) ProtoMessage() {}

func (x *Table) ProtoReflect() protoreflect.Message {
	mi := &file_api_bj21_v1_bj21_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Table.ProtoReflect.Descriptor instead.
func (*Table) Descriptor() ([]byte, []int) {
	return file_api_bj21_v1_bj21_proto_rawDescGZIP(), []int{2}
}

func (x *Table) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Table) GetSeq() string {
	if x != nil {
		return x.Seq
	}
	return ""
}

func (x *Table) GetP1() *Player {
	if x != nil {
		return x.P1
	}
	return nil
}

func (x *Table) GetP2() *Player {
	if x != nil {
		return x.P2
	}
	return nil
}

func (x *Table) GetMe() int32 {
	if x != nil {
		return x.Me
	}
	return 0
}

// login
type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bj21_v1_bj21_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_bj21_v1_bj21_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_api_bj21_v1_bj21_proto_rawDescGZIP(), []int{3}
}

func (x *LoginRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bj21_v1_bj21_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_bj21_v1_bj21_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_api_bj21_v1_bj21_proto_rawDescGZIP(), []int{4}
}

func (x *LoginReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoginReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// logout
type LogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bj21_v1_bj21_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_bj21_v1_bj21_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_api_bj21_v1_bj21_proto_rawDescGZIP(), []int{5}
}

func (x *LogoutRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LogoutReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err string `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *LogoutReply) Reset() {
	*x = LogoutReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bj21_v1_bj21_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutReply) ProtoMessage() {}

func (x *LogoutReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_bj21_v1_bj21_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutReply.ProtoReflect.Descriptor instead.
func (*LogoutReply) Descriptor() ([]byte, []int) {
	return file_api_bj21_v1_bj21_proto_rawDescGZIP(), []int{6}
}

func (x *LogoutReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

// table list
type TableListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *TableListRequest) Reset() {
	*x = TableListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bj21_v1_bj21_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableListRequest) ProtoMessage() {}

func (x *TableListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_bj21_v1_bj21_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableListRequest.ProtoReflect.Descriptor instead.
func (*TableListRequest) Descriptor() ([]byte, []int) {
	return file_api_bj21_v1_bj21_proto_rawDescGZIP(), []int{7}
}

func (x *TableListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type TableListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tables []*Table `protobuf:"bytes,1,rep,name=tables,proto3" json:"tables,omitempty"`
}

func (x *TableListReply) Reset() {
	*x = TableListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bj21_v1_bj21_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableListReply) ProtoMessage() {}

func (x *TableListReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_bj21_v1_bj21_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableListReply.ProtoReflect.Descriptor instead.
func (*TableListReply) Descriptor() ([]byte, []int) {
	return file_api_bj21_v1_bj21_proto_rawDescGZIP(), []int{8}
}

func (x *TableListReply) GetTables() []*Table {
	if x != nil {
		return x.Tables
	}
	return nil
}

// table info
type TableInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	TableSeq string `protobuf:"bytes,2,opt,name=table_seq,json=tableSeq,proto3" json:"table_seq,omitempty"`
}

func (x *TableInfoRequest) Reset() {
	*x = TableInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bj21_v1_bj21_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableInfoRequest) ProtoMessage() {}

func (x *TableInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_bj21_v1_bj21_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableInfoRequest.ProtoReflect.Descriptor instead.
func (*TableInfoRequest) Descriptor() ([]byte, []int) {
	return file_api_bj21_v1_bj21_proto_rawDescGZIP(), []int{9}
}

func (x *TableInfoRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *TableInfoRequest) GetTableSeq() string {
	if x != nil {
		return x.TableSeq
	}
	return ""
}

type TableInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table *Table `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
}

func (x *TableInfoReply) Reset() {
	*x = TableInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bj21_v1_bj21_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableInfoReply) ProtoMessage() {}

func (x *TableInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_bj21_v1_bj21_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableInfoReply.ProtoReflect.Descriptor instead.
func (*TableInfoReply) Descriptor() ([]byte, []int) {
	return file_api_bj21_v1_bj21_proto_rawDescGZIP(), []int{10}
}

func (x *TableInfoReply) GetTable() *Table {
	if x != nil {
		return x.Table
	}
	return nil
}

// sit down
type SitDownRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	TableSeq string `protobuf:"bytes,2,opt,name=table_seq,json=tableSeq,proto3" json:"table_seq,omitempty"`
}

func (x *SitDownRequest) Reset() {
	*x = SitDownRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bj21_v1_bj21_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SitDownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SitDownRequest) ProtoMessage() {}

func (x *SitDownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_bj21_v1_bj21_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SitDownRequest.ProtoReflect.Descriptor instead.
func (*SitDownRequest) Descriptor() ([]byte, []int) {
	return file_api_bj21_v1_bj21_proto_rawDescGZIP(), []int{11}
}

func (x *SitDownRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SitDownRequest) GetTableSeq() string {
	if x != nil {
		return x.TableSeq
	}
	return ""
}

type SitDownReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table *Table `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Err   string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *SitDownReply) Reset() {
	*x = SitDownReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bj21_v1_bj21_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SitDownReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SitDownReply) ProtoMessage() {}

func (x *SitDownReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_bj21_v1_bj21_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SitDownReply.ProtoReflect.Descriptor instead.
func (*SitDownReply) Descriptor() ([]byte, []int) {
	return file_api_bj21_v1_bj21_proto_rawDescGZIP(), []int{12}
}

func (x *SitDownReply) GetTable() *Table {
	if x != nil {
		return x.Table
	}
	return nil
}

func (x *SitDownReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

// stand up
type StandUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	TableSeq string `protobuf:"bytes,2,opt,name=table_seq,json=tableSeq,proto3" json:"table_seq,omitempty"`
}

func (x *StandUpRequest) Reset() {
	*x = StandUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bj21_v1_bj21_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StandUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StandUpRequest) ProtoMessage() {}

func (x *StandUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_bj21_v1_bj21_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StandUpRequest.ProtoReflect.Descriptor instead.
func (*StandUpRequest) Descriptor() ([]byte, []int) {
	return file_api_bj21_v1_bj21_proto_rawDescGZIP(), []int{13}
}

func (x *StandUpRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *StandUpRequest) GetTableSeq() string {
	if x != nil {
		return x.TableSeq
	}
	return ""
}

type StandUpReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err string `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *StandUpReply) Reset() {
	*x = StandUpReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bj21_v1_bj21_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StandUpReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StandUpReply) ProtoMessage() {}

func (x *StandUpReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_bj21_v1_bj21_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StandUpReply.ProtoReflect.Descriptor instead.
func (*StandUpReply) Descriptor() ([]byte, []int) {
	return file_api_bj21_v1_bj21_proto_rawDescGZIP(), []int{14}
}

func (x *StandUpReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_api_bj21_v1_bj21_proto protoreflect.FileDescriptor

var file_api_bj21_v1_bj21_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6a, 0x32, 0x31, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6a,
	0x32, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x6a, 0x32, 0x31, 0x2e, 0x76,
	0x31, 0x22, 0x2f, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x22, 0x1c, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x7f, 0x0a, 0x05, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12,
	0x1f, 0x0a, 0x02, 0x70, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x6a,
	0x32, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x02, 0x70, 0x31,
	0x12, 0x1f, 0x0a, 0x02, 0x70, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62,
	0x6a, 0x32, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x02, 0x70,
	0x32, 0x12, 0x0e, 0x0a, 0x02, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x6d,
	0x65, 0x22, 0x22, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x25, 0x0a,
	0x0d, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x1f, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x28, 0x0a, 0x10, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x38, 0x0a, 0x0e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x26, 0x0a, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x6a, 0x32, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x22, 0x45, 0x0a, 0x10, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x71,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x71,
	0x22, 0x36, 0x0a, 0x0e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x6a, 0x32, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x43, 0x0a, 0x0e, 0x53, 0x69, 0x74, 0x44,
	0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x71, 0x22, 0x46, 0x0a,
	0x0c, 0x53, 0x69, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x24, 0x0a,
	0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62,
	0x6a, 0x32, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x05, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x43, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x55, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x71, 0x22, 0x20, 0x0a, 0x0c, 0x53, 0x74,
	0x61, 0x6e, 0x64, 0x55, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x32, 0x42, 0x0a, 0x09,
	0x42, 0x6c, 0x61, 0x63, 0x6b, 0x4a, 0x61, 0x63, 0x6b, 0x12, 0x35, 0x0a, 0x09, 0x4c, 0x6f, 0x67,
	0x69, 0x63, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x10, 0x2e, 0x62, 0x6a, 0x32, 0x31, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x62, 0x6a, 0x32, 0x31, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x1f, 0x5a, 0x1d, 0x66, 0x78, 0x6b, 0x74, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x62, 0x6a,
	0x32, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6a, 0x32, 0x31, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_bj21_v1_bj21_proto_rawDescOnce sync.Once
	file_api_bj21_v1_bj21_proto_rawDescData = file_api_bj21_v1_bj21_proto_rawDesc
)

func file_api_bj21_v1_bj21_proto_rawDescGZIP() []byte {
	file_api_bj21_v1_bj21_proto_rawDescOnce.Do(func() {
		file_api_bj21_v1_bj21_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_bj21_v1_bj21_proto_rawDescData)
	})
	return file_api_bj21_v1_bj21_proto_rawDescData
}

var file_api_bj21_v1_bj21_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_api_bj21_v1_bj21_proto_goTypes = []interface{}{
	(*Message)(nil),          // 0: bj21.v1.Message
	(*Player)(nil),           // 1: bj21.v1.Player
	(*Table)(nil),            // 2: bj21.v1.Table
	(*LoginRequest)(nil),     // 3: bj21.v1.LoginRequest
	(*LoginReply)(nil),       // 4: bj21.v1.LoginReply
	(*LogoutRequest)(nil),    // 5: bj21.v1.LogoutRequest
	(*LogoutReply)(nil),      // 6: bj21.v1.LogoutReply
	(*TableListRequest)(nil), // 7: bj21.v1.TableListRequest
	(*TableListReply)(nil),   // 8: bj21.v1.TableListReply
	(*TableInfoRequest)(nil), // 9: bj21.v1.TableInfoRequest
	(*TableInfoReply)(nil),   // 10: bj21.v1.TableInfoReply
	(*SitDownRequest)(nil),   // 11: bj21.v1.SitDownRequest
	(*SitDownReply)(nil),     // 12: bj21.v1.SitDownReply
	(*StandUpRequest)(nil),   // 13: bj21.v1.StandUpRequest
	(*StandUpReply)(nil),     // 14: bj21.v1.StandUpReply
}
var file_api_bj21_v1_bj21_proto_depIdxs = []int32{
	1, // 0: bj21.v1.Table.p1:type_name -> bj21.v1.Player
	1, // 1: bj21.v1.Table.p2:type_name -> bj21.v1.Player
	2, // 2: bj21.v1.TableListReply.tables:type_name -> bj21.v1.Table
	2, // 3: bj21.v1.TableInfoReply.table:type_name -> bj21.v1.Table
	2, // 4: bj21.v1.SitDownReply.table:type_name -> bj21.v1.Table
	0, // 5: bj21.v1.BlackJack.LogicConn:input_type -> bj21.v1.Message
	0, // 6: bj21.v1.BlackJack.LogicConn:output_type -> bj21.v1.Message
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_bj21_v1_bj21_proto_init() }
func file_api_bj21_v1_bj21_proto_init() {
	if File_api_bj21_v1_bj21_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_bj21_v1_bj21_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_api_bj21_v1_bj21_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
		file_api_bj21_v1_bj21_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Table); i {
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
		file_api_bj21_v1_bj21_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_api_bj21_v1_bj21_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply); i {
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
		file_api_bj21_v1_bj21_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutRequest); i {
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
		file_api_bj21_v1_bj21_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutReply); i {
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
		file_api_bj21_v1_bj21_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableListRequest); i {
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
		file_api_bj21_v1_bj21_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableListReply); i {
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
		file_api_bj21_v1_bj21_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableInfoRequest); i {
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
		file_api_bj21_v1_bj21_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableInfoReply); i {
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
		file_api_bj21_v1_bj21_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SitDownRequest); i {
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
		file_api_bj21_v1_bj21_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SitDownReply); i {
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
		file_api_bj21_v1_bj21_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StandUpRequest); i {
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
		file_api_bj21_v1_bj21_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StandUpReply); i {
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
			RawDescriptor: file_api_bj21_v1_bj21_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_bj21_v1_bj21_proto_goTypes,
		DependencyIndexes: file_api_bj21_v1_bj21_proto_depIdxs,
		MessageInfos:      file_api_bj21_v1_bj21_proto_msgTypes,
	}.Build()
	File_api_bj21_v1_bj21_proto = out.File
	file_api_bj21_v1_bj21_proto_rawDesc = nil
	file_api_bj21_v1_bj21_proto_goTypes = nil
	file_api_bj21_v1_bj21_proto_depIdxs = nil
}
