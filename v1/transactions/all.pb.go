// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/openbank/openbank/v1/transactions/all.proto

package transactions

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	types "github.com/openbank/openbank/v1/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Type defines the type of a transaction.
type Type int32

const (
	Type__ Type = 0
	// Type_Credit is the value for a credit transaction.
	Type_Credit Type = 1
	// Type_Debit is the value for a debit transaction.
	Type_Debit Type = 2
)

var Type_name = map[int32]string{
	0: "_",
	1: "Credit",
	2: "Debit",
}

var Type_value = map[string]int32{
	"_":      0,
	"Credit": 1,
	"Debit":  2,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36b883a290e71724, []int{0}
}

// Status defines the status of a transaction.
type Status int32

const (
	Status__ Status = 0
	// Status_Success is the value for a successful transaction.
	Status_Success Status = 1
	// Status_Pending is the value for a pending transaction.
	Status_Pending Status = 2
	// Status_Rejected is the value for a rejected transaction.
	Status_Rejected Status = 3
)

var Status_name = map[int32]string{
	0: "_",
	1: "Success",
	2: "Pending",
	3: "Rejected",
}

var Status_value = map[string]int32{
	"_":        0,
	"Success":  1,
	"Pending":  2,
	"Rejected": 3,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36b883a290e71724, []int{1}
}

// Transaction holds all details about a transaction.
type Transaction struct {
	// TransactionID is the unique identifier of a transaction.
	TransactionID string `protobuf:"bytes,1,opt,name=TransactionID,json=transaction_id,proto3" json:"TransactionID,omitempty"`
	// SourceAccount is the account emitting the transaction.
	SourceAccount *BankAccountInfo `protobuf:"bytes,2,opt,name=SourceAccount,json=source_account_id,proto3" json:"SourceAccount,omitempty"`
	// DestinationAccount is the account receiving the transaction.
	DestinationAccount *BankAccountInfo `protobuf:"bytes,3,opt,name=DestinationAccount,json=destination_account_id,proto3" json:"DestinationAccount,omitempty"`
	// Date is the date of the transaction.
	Date string `protobuf:"bytes,4,opt,name=Date,json=date,proto3" json:"Date,omitempty"`
	// Type is the type of transaction.
	Type Type `protobuf:"varint,5,opt,name=Type,json=type,proto3,enum=transactions.Type" json:"Type,omitempty"`
	// Status is the status of the transaction.
	Status Status `protobuf:"varint,6,opt,name=Status,json=status,proto3,enum=transactions.Status" json:"Status,omitempty"`
	// Amount holds the amount value and currency of the transaction.
	Amount *types.Amount `protobuf:"bytes,7,opt,name=Amount,json=amount,proto3" json:"Amount,omitempty"`
	// Description describes the transaction.
	Description string `protobuf:"bytes,8,opt,name=Description,json=description,proto3" json:"Description,omitempty"`
	// UserID is the identifier of the issuer of the transaction.
	UserID string `protobuf:"bytes,9,opt,name=UserID,json=user_id,proto3" json:"UserID,omitempty"`
	// Remarks is an informational note about the transaction.
	Remarks              string   `protobuf:"bytes,10,opt,name=Remarks,json=remarks,proto3" json:"Remarks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_36b883a290e71724, []int{0}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetTransactionID() string {
	if m != nil {
		return m.TransactionID
	}
	return ""
}

func (m *Transaction) GetSourceAccount() *BankAccountInfo {
	if m != nil {
		return m.SourceAccount
	}
	return nil
}

func (m *Transaction) GetDestinationAccount() *BankAccountInfo {
	if m != nil {
		return m.DestinationAccount
	}
	return nil
}

func (m *Transaction) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *Transaction) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type__
}

func (m *Transaction) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status__
}

func (m *Transaction) GetAmount() *types.Amount {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *Transaction) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Transaction) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Transaction) GetRemarks() string {
	if m != nil {
		return m.Remarks
	}
	return ""
}

// BankAccountInfo holds a lightweight account information.
type BankAccountInfo struct {
	// AccountID is the identifier of the account.
	AccountID string `protobuf:"bytes,1,opt,name=AccountID,json=account_id,proto3" json:"AccountID,omitempty"`
	// BankCode is code of the bank the account belongs to.
	BankCode types.BankCode `protobuf:"varint,2,opt,name=BankCode,json=bank_code,proto3,enum=types.BankCode" json:"BankCode,omitempty"`
	// OwnerName is the name of the owner of the account.
	OwnerName string `protobuf:"bytes,3,opt,name=OwnerName,json=owner_name,proto3" json:"OwnerName,omitempty"`
	// MajorType is the type of account.
	MajorType            types.MajorType `protobuf:"varint,4,opt,name=MajorType,json=major_type,proto3,enum=types.MajorType" json:"MajorType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BankAccountInfo) Reset()         { *m = BankAccountInfo{} }
func (m *BankAccountInfo) String() string { return proto.CompactTextString(m) }
func (*BankAccountInfo) ProtoMessage()    {}
func (*BankAccountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_36b883a290e71724, []int{1}
}

func (m *BankAccountInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BankAccountInfo.Unmarshal(m, b)
}
func (m *BankAccountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BankAccountInfo.Marshal(b, m, deterministic)
}
func (m *BankAccountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BankAccountInfo.Merge(m, src)
}
func (m *BankAccountInfo) XXX_Size() int {
	return xxx_messageInfo_BankAccountInfo.Size(m)
}
func (m *BankAccountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BankAccountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BankAccountInfo proto.InternalMessageInfo

func (m *BankAccountInfo) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *BankAccountInfo) GetBankCode() types.BankCode {
	if m != nil {
		return m.BankCode
	}
	return types.BankCode__
}

func (m *BankAccountInfo) GetOwnerName() string {
	if m != nil {
		return m.OwnerName
	}
	return ""
}

func (m *BankAccountInfo) GetMajorType() types.MajorType {
	if m != nil {
		return m.MajorType
	}
	return types.MajorType__
}

// GetTransactionRequest is the request envelope to get an transaction by its identifier.
type GetTransactionRequest struct {
	// TransactionID is the unique identifier of a transaction.
	TransactionID        string   `protobuf:"bytes,1,opt,name=TransactionID,json=transaction_id,proto3" json:"TransactionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTransactionRequest) Reset()         { *m = GetTransactionRequest{} }
func (m *GetTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*GetTransactionRequest) ProtoMessage()    {}
func (*GetTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_36b883a290e71724, []int{2}
}

func (m *GetTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionRequest.Unmarshal(m, b)
}
func (m *GetTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionRequest.Marshal(b, m, deterministic)
}
func (m *GetTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionRequest.Merge(m, src)
}
func (m *GetTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_GetTransactionRequest.Size(m)
}
func (m *GetTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionRequest proto.InternalMessageInfo

func (m *GetTransactionRequest) GetTransactionID() string {
	if m != nil {
		return m.TransactionID
	}
	return ""
}

// GetTransactionsRequest is the request envelope to get a list of transactions.
type GetTransactionsRequest struct {
	// NextStartingIndex is a cursor for pagination. It's a TransactionID that defines
	//the current place in the total list of Transaction.
	NextStartingIndex    string   `protobuf:"bytes,1,opt,name=NextStartingIndex,json=next_starting_index,proto3" json:"NextStartingIndex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTransactionsRequest) Reset()         { *m = GetTransactionsRequest{} }
func (m *GetTransactionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetTransactionsRequest) ProtoMessage()    {}
func (*GetTransactionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_36b883a290e71724, []int{3}
}

func (m *GetTransactionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionsRequest.Unmarshal(m, b)
}
func (m *GetTransactionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionsRequest.Marshal(b, m, deterministic)
}
func (m *GetTransactionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionsRequest.Merge(m, src)
}
func (m *GetTransactionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetTransactionsRequest.Size(m)
}
func (m *GetTransactionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionsRequest proto.InternalMessageInfo

func (m *GetTransactionsRequest) GetNextStartingIndex() string {
	if m != nil {
		return m.NextStartingIndex
	}
	return ""
}

// GetTransactionsResponse wraps the list of transactions.
type GetTransactionsResponse struct {
	// Result is a list containing up to 20 transactions.
	Result []*Transaction `protobuf:"bytes,1,rep,name=Result,json=result,proto3" json:"Result,omitempty"`
	// HasMore indicates if there are more transactions available.
	HasMore              bool     `protobuf:"varint,2,opt,name=HasMore,json=has_more,proto3" json:"HasMore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTransactionsResponse) Reset()         { *m = GetTransactionsResponse{} }
func (m *GetTransactionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetTransactionsResponse) ProtoMessage()    {}
func (*GetTransactionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_36b883a290e71724, []int{4}
}

func (m *GetTransactionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionsResponse.Unmarshal(m, b)
}
func (m *GetTransactionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionsResponse.Marshal(b, m, deterministic)
}
func (m *GetTransactionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionsResponse.Merge(m, src)
}
func (m *GetTransactionsResponse) XXX_Size() int {
	return xxx_messageInfo_GetTransactionsResponse.Size(m)
}
func (m *GetTransactionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionsResponse proto.InternalMessageInfo

func (m *GetTransactionsResponse) GetResult() []*Transaction {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GetTransactionsResponse) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

// CreateTransactionRequest wraps all required field for a transaction creation.
type CreateTransactionRequest struct {
	// SourceAccountID is the identifier of the account emitting the transaction.
	SourceAccountID string `protobuf:"bytes,1,opt,name=SourceAccountID,json=source_account_id,proto3" json:"SourceAccountID,omitempty"`
	// DestinationAccountID is the identifier of the account receiving the transaction.
	DestinationAccountID string `protobuf:"bytes,2,opt,name=DestinationAccountID,json=destination_account_id,proto3" json:"DestinationAccountID,omitempty"`
	// Amount holds the amount value and currency of the transaction.
	Amount *types.Amount `protobuf:"bytes,3,opt,name=Amount,json=amount,proto3" json:"Amount,omitempty"`
	// Remarks is an informational note about the transaction.
	Remarks              string   `protobuf:"bytes,4,opt,name=Remarks,json=remarks,proto3" json:"Remarks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTransactionRequest) Reset()         { *m = CreateTransactionRequest{} }
func (m *CreateTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTransactionRequest) ProtoMessage()    {}
func (*CreateTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_36b883a290e71724, []int{5}
}

func (m *CreateTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTransactionRequest.Unmarshal(m, b)
}
func (m *CreateTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTransactionRequest.Marshal(b, m, deterministic)
}
func (m *CreateTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTransactionRequest.Merge(m, src)
}
func (m *CreateTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTransactionRequest.Size(m)
}
func (m *CreateTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTransactionRequest proto.InternalMessageInfo

func (m *CreateTransactionRequest) GetSourceAccountID() string {
	if m != nil {
		return m.SourceAccountID
	}
	return ""
}

func (m *CreateTransactionRequest) GetDestinationAccountID() string {
	if m != nil {
		return m.DestinationAccountID
	}
	return ""
}

func (m *CreateTransactionRequest) GetAmount() *types.Amount {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *CreateTransactionRequest) GetRemarks() string {
	if m != nil {
		return m.Remarks
	}
	return ""
}

// CreateTransactionResponse is the response envelope for a transaction creation.
type CreateTransactionResponse struct {
	// TransactionID is the unique identifier of a transaction.
	TransactionID        string   `protobuf:"bytes,1,opt,name=TransactionID,json=transaction_id,proto3" json:"TransactionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTransactionResponse) Reset()         { *m = CreateTransactionResponse{} }
func (m *CreateTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTransactionResponse) ProtoMessage()    {}
func (*CreateTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_36b883a290e71724, []int{6}
}

func (m *CreateTransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTransactionResponse.Unmarshal(m, b)
}
func (m *CreateTransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTransactionResponse.Marshal(b, m, deterministic)
}
func (m *CreateTransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTransactionResponse.Merge(m, src)
}
func (m *CreateTransactionResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTransactionResponse.Size(m)
}
func (m *CreateTransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTransactionResponse proto.InternalMessageInfo

func (m *CreateTransactionResponse) GetTransactionID() string {
	if m != nil {
		return m.TransactionID
	}
	return ""
}

func init() {
	proto.RegisterEnum("transactions.Type", Type_name, Type_value)
	proto.RegisterEnum("transactions.Status", Status_name, Status_value)
	proto.RegisterType((*Transaction)(nil), "transactions.Transaction")
	proto.RegisterType((*BankAccountInfo)(nil), "transactions.BankAccountInfo")
	proto.RegisterType((*GetTransactionRequest)(nil), "transactions.GetTransactionRequest")
	proto.RegisterType((*GetTransactionsRequest)(nil), "transactions.GetTransactionsRequest")
	proto.RegisterType((*GetTransactionsResponse)(nil), "transactions.GetTransactionsResponse")
	proto.RegisterType((*CreateTransactionRequest)(nil), "transactions.CreateTransactionRequest")
	proto.RegisterType((*CreateTransactionResponse)(nil), "transactions.CreateTransactionResponse")
}

func init() {
	proto.RegisterFile("github.com/openbank/openbank/v1/transactions/all.proto", fileDescriptor_36b883a290e71724)
}

var fileDescriptor_36b883a290e71724 = []byte{
	// 1548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x57, 0xdd, 0x6f, 0x1b, 0x4b,
	0x15, 0xdf, 0xb5, 0x1d, 0xc7, 0x9e, 0x90, 0xc4, 0x99, 0x7e, 0xe0, 0x6b, 0x5d, 0xae, 0x86, 0xf4,
	0x5e, 0x1a, 0xa2, 0x76, 0xfd, 0x71, 0xcb, 0xd5, 0x55, 0xd0, 0x15, 0x75, 0x12, 0x68, 0x53, 0xfa,
	0x11, 0x39, 0xa5, 0x82, 0x02, 0xb2, 0xc6, 0xbb, 0xc7, 0xf6, 0x36, 0xeb, 0x19, 0x33, 0x33, 0x9b,
	0x0f, 0x10, 0x12, 0xe2, 0x01, 0xfa, 0x58, 0xa5, 0xe2, 0x8d, 0x37, 0xfe, 0x07, 0xd4, 0xbf, 0x80,
	0x67, 0x24, 0x24, 0xc4, 0x0b, 0x3c, 0x21, 0x24, 0x10, 0x2f, 0x80, 0x90, 0xfa, 0x88, 0x66, 0x76,
	0x6d, 0xef, 0xda, 0x4e, 0x53, 0xaa, 0xfb, 0xd4, 0xce, 0x99, 0x73, 0x7e, 0x73, 0xce, 0xef, 0xfc,
	0xce, 0xc9, 0x1a, 0x7d, 0xd2, 0xf3, 0x55, 0x3f, 0xec, 0x38, 0x2e, 0x1f, 0x54, 0xf9, 0x10, 0x58,
	0x87, 0xb2, 0xc3, 0xc9, 0x7f, 0x8e, 0xea, 0x55, 0x25, 0x28, 0x93, 0xd4, 0x55, 0x3e, 0x67, 0xb2,
	0x4a, 0x83, 0xc0, 0x19, 0x0a, 0xae, 0x38, 0xfe, 0x42, 0xd2, 0x5e, 0x79, 0xbf, 0xc7, 0x79, 0x2f,
	0x80, 0x2a, 0x1d, 0xfa, 0x55, 0xca, 0x18, 0x57, 0xd4, 0xd8, 0x23, 0xdf, 0xca, 0x0d, 0xf3, 0x8f,
	0x7b, 0xb3, 0x07, 0xec, 0xa6, 0x3c, 0xa6, 0xbd, 0x1e, 0x88, 0x2a, 0x1f, 0xc6, 0x88, 0x33, 0xde,
	0xd5, 0x0b, 0x33, 0x3a, 0x1d, 0x42, 0x22, 0x95, 0xf5, 0xdf, 0xe6, 0xd0, 0xd2, 0xe3, 0x49, 0x36,
	0xb8, 0x8e, 0x96, 0x13, 0xc7, 0xbd, 0xdd, 0xb2, 0x4d, 0xec, 0x8d, 0xe2, 0x36, 0x2a, 0x58, 0x65,
	0x6b, 0xc3, 0xaa, 0x59, 0xfb, 0x56, 0x6b, 0x25, 0x91, 0x7d, 0xdb, 0xf7, 0xf0, 0x01, 0x5a, 0x3e,
	0xe0, 0xa1, 0x70, 0xa1, 0xe9, 0xba, 0x3c, 0x64, 0xaa, 0x9c, 0x21, 0xf6, 0xc6, 0x52, 0xe3, 0x4b,
	0x4e, 0xb2, 0x4a, 0x67, 0x9b, 0xb2, 0xc3, 0xd8, 0x61, 0x8f, 0x75, 0x79, 0x0a, 0x71, 0x4d, 0x9a,
	0xf8, 0x36, 0x8d, 0xee, 0x35, 0xe8, 0x0f, 0x11, 0xde, 0x05, 0xa9, 0x7c, 0x66, 0xca, 0x1b, 0x21,
	0x67, 0xff, 0x5f, 0xe4, 0xab, 0xde, 0x04, 0x24, 0x09, 0xff, 0x01, 0xca, 0xed, 0x52, 0x05, 0xe5,
	0xdc, 0x4c, 0x75, 0x39, 0x8f, 0x2a, 0xc0, 0x0d, 0x94, 0x7b, 0x7c, 0x3a, 0x84, 0xf2, 0x02, 0xb1,
	0x37, 0x56, 0x1a, 0x38, 0xfd, 0xa0, 0xbe, 0x49, 0xc7, 0x68, 0x56, 0xf1, 0xa7, 0x28, 0x7f, 0xa0,
	0xa8, 0x0a, 0x65, 0x39, 0x6f, 0xa2, 0x2e, 0xa7, 0xa3, 0xa2, 0xbb, 0x54, 0x5c, 0x5e, 0x1a, 0x1b,
	0xae, 0xa3, 0x7c, 0x73, 0x60, 0x0a, 0x5c, 0x34, 0x05, 0x2e, 0x3b, 0xa6, 0x4d, 0x4e, 0x64, 0x4c,
	0x87, 0x50, 0x63, 0xc3, 0x37, 0xd0, 0xd2, 0x2e, 0x48, 0x57, 0xf8, 0x46, 0x0a, 0xe5, 0xc2, 0x4c,
	0x1d, 0x4b, 0xde, 0xe4, 0x1a, 0x5f, 0x43, 0xf9, 0xef, 0x48, 0x10, 0x7b, 0xbb, 0xe5, 0xe2, 0x8c,
	0xe3, 0x62, 0x28, 0x41, 0x68, 0x4e, 0x3e, 0x44, 0x8b, 0x2d, 0x18, 0x50, 0x71, 0x28, 0xcb, 0x68,
	0xd6, 0x4b, 0x44, 0x57, 0x5b, 0xf9, 0x82, 0x55, 0xb2, 0xca, 0xd6, 0xfa, 0xdf, 0x6c, 0xb4, 0x3a,
	0xc5, 0x3c, 0xfe, 0x2a, 0x2a, 0x8e, 0x8e, 0xf3, 0x84, 0x83, 0x12, 0x0d, 0xf8, 0x14, 0x15, 0x74,
	0xf4, 0x0e, 0xf7, 0xc0, 0xe8, 0x65, 0xa5, 0xb1, 0x1a, 0x17, 0x3d, 0x32, 0xa7, 0x42, 0x8b, 0x5a,
	0xbf, 0x6d, 0x97, 0x7b, 0xa0, 0x1f, 0x79, 0x74, 0xcc, 0x40, 0x3c, 0xa4, 0x03, 0x30, 0x82, 0x98,
	0x7a, 0x84, 0xeb, 0xcb, 0x36, 0xa3, 0x03, 0xc0, 0x5f, 0x47, 0xc5, 0x07, 0xf4, 0x19, 0x17, 0xa6,
	0x95, 0x39, 0xf3, 0x4a, 0x29, 0x7e, 0x65, 0x6c, 0x4f, 0x07, 0x0f, 0xb4, 0xb9, 0xad, 0x5d, 0xc6,
	0x85, 0xb6, 0xd0, 0x95, 0x3b, 0xa0, 0x12, 0x43, 0xd1, 0x82, 0x1f, 0x85, 0x20, 0xd5, 0x3b, 0x8c,
	0xca, 0x18, 0xf3, 0x07, 0xe8, 0x6a, 0x1a, 0x53, 0x8e, 0x40, 0xb7, 0xd0, 0xda, 0x43, 0x38, 0x51,
	0x07, 0x8a, 0x0a, 0xe5, 0xb3, 0xde, 0x1e, 0xf3, 0xe0, 0x64, 0x0e, 0xf0, 0x25, 0x06, 0x27, 0xaa,
	0x2d, 0x63, 0xaf, 0xb6, 0xaf, 0xdd, 0xc6, 0xe8, 0xbf, 0xb4, 0xd1, 0x17, 0x67, 0xe0, 0xe5, 0x90,
	0x33, 0x09, 0xf8, 0x33, 0x94, 0x6f, 0x81, 0x0c, 0x03, 0x55, 0xb6, 0x49, 0x76, 0x63, 0xa9, 0xf1,
	0xde, 0x94, 0xb4, 0x27, 0x87, 0xb4, 0xec, 0x84, 0x09, 0xc2, 0x1f, 0xa1, 0xc5, 0xbb, 0x54, 0x3e,
	0xe0, 0x22, 0xea, 0x5a, 0x21, 0xe5, 0x54, 0xe8, 0x53, 0xd9, 0x1e, 0x70, 0x31, 0xe1, 0xee, 0xbf,
	0x36, 0x2a, 0xef, 0x08, 0xa0, 0x0a, 0xe6, 0xf0, 0xf7, 0x09, 0x5a, 0x4d, 0xed, 0x8d, 0xb9, 0x0c,
	0xce, 0x59, 0x0d, 0xb7, 0xd1, 0xe5, 0xd9, 0xd5, 0xb0, 0xb7, 0x6b, 0x12, 0x2a, 0xbe, 0xd5, 0xf4,
	0x4f, 0xe6, 0x2d, 0xfb, 0xb6, 0xf3, 0x96, 0x18, 0x8e, 0xdc, 0xc5, 0xc3, 0xf1, 0x04, 0xbd, 0x37,
	0xa7, 0xec, 0xb8, 0x05, 0xef, 0xae, 0x9b, 0xcd, 0x5b, 0xd1, 0x5a, 0xc2, 0x45, 0x64, 0xb7, 0x4b,
	0x56, 0x25, 0x53, 0xb0, 0xf0, 0x0a, 0xca, 0xef, 0x08, 0xf0, 0x7c, 0x55, 0xb2, 0xcd, 0x79, 0x19,
	0x2d, 0xec, 0x42, 0xc7, 0x57, 0xa5, 0x8c, 0x3e, 0x56, 0x32, 0x65, 0x6b, 0xf3, 0xdb, 0xa3, 0xc5,
	0x94, 0x8c, 0x5b, 0x45, 0x8b, 0x07, 0xa1, 0xeb, 0x82, 0x94, 0x71, 0xe0, 0x2a, 0x5a, 0xdc, 0x07,
	0xe6, 0xf9, 0xac, 0x17, 0x85, 0xe2, 0x12, 0x2a, 0xb4, 0xe0, 0x19, 0xb8, 0x0a, 0xbc, 0x52, 0x76,
	0x04, 0xd6, 0xf8, 0x55, 0x11, 0xe1, 0x44, 0xfa, 0x07, 0x20, 0x8e, 0x7c, 0x17, 0xf0, 0x7f, 0x32,
	0x68, 0x25, 0xad, 0x39, 0x7c, 0x2d, 0x2d, 0xad, 0xb9, 0x43, 0x54, 0x39, 0x5f, 0x7f, 0xeb, 0xbf,
	0xce, 0x9c, 0x35, 0xff, 0x61, 0xa7, 0xff, 0x3c, 0x5d, 0x6d, 0x81, 0x12, 0x3e, 0x1c, 0x01, 0xa1,
	0x24, 0x11, 0x58, 0x39, 0x18, 0xd9, 0x25, 0xa1, 0x41, 0x40, 0x3c, 0xaa, 0x28, 0xe9, 0x0a, 0x3e,
	0x48, 0xbb, 0xdd, 0x20, 0x12, 0x02, 0x53, 0x10, 0xe9, 0x9c, 0x12, 0xd5, 0x07, 0x92, 0xe6, 0x99,
	0x9c, 0xf2, 0x90, 0xc8, 0x70, 0x38, 0x0c, 0x7c, 0xf0, 0x9c, 0x7b, 0x75, 0x94, 0xbd, 0x55, 0xbb,
	0x85, 0x37, 0xd1, 0x46, 0x0b, 0x54, 0x28, 0x18, 0x78, 0xe4, 0xb8, 0x0f, 0xcc, 0x04, 0x0a, 0x88,
	0x84, 0x49, 0x7c, 0x49, 0x18, 0x57, 0xa4, 0xcb, 0x43, 0xe6, 0x39, 0xf7, 0xae, 0xa3, 0x6c, 0xa3,
	0x56, 0xc3, 0x04, 0x7d, 0x10, 0x17, 0x48, 0xe0, 0x04, 0xdc, 0x50, 0xbf, 0x28, 0x23, 0xb6, 0xbb,
	0x61, 0x10, 0x9c, 0x3a, 0x1d, 0x8c, 0x4a, 0x28, 0xff, 0xa8, 0x19, 0xaa, 0x7e, 0x03, 0xe7, 0x51,
	0x4e, 0x00, 0xf5, 0x7e, 0xfe, 0x87, 0xbf, 0xbe, 0xcc, 0xac, 0x63, 0x32, 0xf3, 0xe1, 0xf0, 0x93,
	0x94, 0x5c, 0x7e, 0xfa, 0x3c, 0x63, 0xbd, 0xc8, 0x18, 0xa5, 0xe0, 0x3f, 0x67, 0xd1, 0xea, 0xd4,
	0xa4, 0xe3, 0x0f, 0xdf, 0x44, 0xfb, 0x68, 0xcf, 0x54, 0x3e, 0xba, 0xc0, 0x2b, 0xd2, 0xea, 0xfa,
	0xeb, 0xcc, 0x59, 0xf3, 0x8f, 0x99, 0x74, 0x0f, 0xae, 0xdc, 0xf7, 0xa5, 0x32, 0x34, 0xa7, 0x3e,
	0x63, 0xfe, 0x6d, 0x47, 0x2c, 0x49, 0x42, 0x49, 0xa0, 0x3d, 0x5c, 0xce, 0x14, 0xf5, 0x99, 0xcf,
	0x7a, 0x24, 0x1c, 0x12, 0xc5, 0x49, 0xa3, 0x96, 0x8a, 0x70, 0xc8, 0xf7, 0x78, 0x48, 0x5c, 0xca,
	0xc8, 0x90, 0xf6, 0xf4, 0x94, 0x02, 0x51, 0x7d, 0xc1, 0xc3, 0x5e, 0x3f, 0xe5, 0xa6, 0x5b, 0x65,
	0x7a, 0x72, 0xaa, 0x91, 0xe6, 0xec, 0x3d, 0xe2, 0x33, 0xdd, 0x38, 0x41, 0x64, 0xd8, 0x91, 0xba,
	0x3c, 0xa6, 0x24, 0x71, 0x69, 0x10, 0x48, 0x67, 0xae, 0x7b, 0x9c, 0x97, 0x9c, 0x27, 0x00, 0xde,
	0x35, 0xd6, 0x80, 0x4a, 0x75, 0xce, 0x95, 0x1b, 0x0a, 0x01, 0x4c, 0xe9, 0x9c, 0xe1, 0x73, 0xe8,
	0x37, 0xc6, 0xa5, 0xe9, 0x7e, 0x27, 0xfa, 0xfb, 0x9b, 0x0c, 0x5a, 0x9b, 0x59, 0x24, 0xf8, 0x2b,
	0xe9, 0xde, 0x9d, 0xb7, 0x60, 0x2b, 0xd7, 0x2f, 0xf4, 0x8b, 0xbb, 0xfc, 0xca, 0x3e, 0x6b, 0xbe,
	0x9c, 0x9a, 0xb4, 0xcb, 0x91, 0xfb, 0xd4, 0x9c, 0xdd, 0x8c, 0xac, 0xba, 0xc7, 0x0c, 0x8e, 0x93,
	0x57, 0x84, 0x32, 0x8f, 0x88, 0x58, 0x01, 0xbe, 0x92, 0xc4, 0xf7, 0x9c, 0x7b, 0x9b, 0x9a, 0x9e,
	0x3a, 0xbe, 0x86, 0xbe, 0x9c, 0x80, 0x26, 0xae, 0xc1, 0x98, 0x66, 0xe8, 0x12, 0x5a, 0x1b, 0x33,
	0xb4, 0x88, 0x16, 0x8e, 0x85, 0xaf, 0xc0, 0x50, 0x74, 0x65, 0xcb, 0xde, 0x5c, 0x7f, 0x03, 0x4b,
	0x95, 0xec, 0xf3, 0x8c, 0xb5, 0xfd, 0xaf, 0x85, 0xb3, 0xe6, 0xef, 0x16, 0x50, 0xb6, 0xe1, 0xd4,
	0xf0, 0xf7, 0x51, 0x29, 0xa9, 0x66, 0xd2, 0xdc, 0xdf, 0xc3, 0xb7, 0xf7, 0x05, 0x3f, 0xf2, 0x3d,
	0x90, 0xf1, 0xf3, 0x71, 0xbe, 0xd4, 0x23, 0x7c, 0x08, 0x22, 0xfa, 0x88, 0x26, 0x9c, 0x4d, 0x8b,
	0x62, 0x3c, 0xe8, 0x4e, 0x63, 0xa1, 0xee, 0xd4, 0x9c, 0xda, 0x7a, 0xb6, 0x7a, 0x54, 0xdf, 0xb4,
	0x33, 0x8d, 0x12, 0xd5, 0x7b, 0xc2, 0x35, 0x91, 0xd5, 0x67, 0x92, 0xb3, 0xad, 0x19, 0x4b, 0xeb,
	0xbe, 0xde, 0x20, 0x75, 0xfc, 0x4d, 0xb4, 0x33, 0x6f, 0x83, 0x44, 0x82, 0xf1, 0x38, 0x44, 0x2b,
	0x24, 0xa5, 0x4b, 0xfd, 0x51, 0x76, 0xdd, 0xe4, 0xea, 0x01, 0x53, 0x3e, 0x0d, 0xa4, 0xd3, 0xda,
	0xd7, 0x68, 0x1f, 0xe3, 0x3d, 0x74, 0x67, 0x16, 0x4d, 0xfb, 0x4f, 0xa0, 0xfa, 0xf4, 0x08, 0xc8,
	0x10, 0xc4, 0xc0, 0x97, 0x52, 0x17, 0xa1, 0x38, 0xa1, 0x86, 0xeb, 0xd4, 0xee, 0x72, 0x5a, 0xdf,
	0x42, 0xd9, 0xaf, 0xd5, 0x6a, 0xf8, 0x1b, 0xe8, 0xb3, 0x34, 0x22, 0x65, 0x24, 0x64, 0x70, 0x32,
	0x8c, 0x76, 0x25, 0x08, 0xc1, 0x05, 0xe1, 0xae, 0x1b, 0x0a, 0xf0, 0x46, 0x1c, 0x49, 0x10, 0x47,
	0x20, 0x88, 0xf4, 0x3d, 0x70, 0x5a, 0x6d, 0x9d, 0x59, 0x0d, 0x7f, 0x17, 0x3d, 0x39, 0xbf, 0xce,
	0x0e, 0xf7, 0x4e, 0xf5, 0xb6, 0x1c, 0xd0, 0xa0, 0xcb, 0xc5, 0x80, 0x2a, 0x0d, 0xcd, 0x13, 0x49,
	0x0f, 0xa8, 0x72, 0xfb, 0x26, 0x64, 0xfc, 0x72, 0x1c, 0xeb, 0x3c, 0xfd, 0xa7, 0x8d, 0xfe, 0x6e,
	0x8f, 0xe5, 0xf1, 0x17, 0xbb, 0x90, 0xc5, 0xbf, 0xb0, 0x9b, 0x71, 0x49, 0x3c, 0xbd, 0x2f, 0x46,
	0xe5, 0x49, 0xfd, 0x9a, 0x00, 0xa9, 0x84, 0x6f, 0xc0, 0x34, 0x09, 0xa1, 0xea, 0x6b, 0x3a, 0x5d,
	0xa3, 0x3f, 0xcd, 0x99, 0x74, 0xc8, 0xe3, 0x3e, 0x24, 0x2f, 0x34, 0x5f, 0x43, 0xc1, 0x0d, 0x74,
	0x97, 0x07, 0x01, 0x3f, 0x8e, 0x58, 0xd3, 0x4f, 0x73, 0xe1, 0xff, 0x38, 0xf2, 0xd0, 0x5f, 0xa8,
	0xa4, 0x1b, 0xf0, 0x63, 0x67, 0x23, 0xd7, 0x28, 0x68, 0x71, 0x6a, 0x88, 0xad, 0xa2, 0x91, 0x29,
	0x3f, 0x04, 0xb6, 0xbd, 0x85, 0xde, 0x8f, 0x55, 0x8c, 0x2f, 0xdd, 0x11, 0x54, 0x2f, 0x25, 0x73,
	0x8a, 0x3b, 0x81, 0x2a, 0xd1, 0x16, 0xc0, 0x38, 0xbe, 0x34, 0x5a, 0x8c, 0xee, 0xee, 0xda, 0xfb,
	0xd6, 0xd3, 0xd4, 0x8f, 0xc4, 0x9f, 0xd9, 0xd6, 0x73, 0xdb, 0x7a, 0x61, 0x5b, 0xaf, 0x6c, 0xeb,
	0x4f, 0xb6, 0xf5, 0xda, 0xb6, 0x7e, 0x9f, 0xb1, 0x3a, 0x79, 0xf3, 0x1b, 0xee, 0xe3, 0xff, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x6f, 0x39, 0x97, 0x05, 0x88, 0x0e, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionServiceClient interface {
	// GetTransaction retrieves the detail of a transaction, selected by its id.
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*Transaction, error)
	// GetTransactions returns a list containing up to 20 transactions.
	GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error)
	// CreateTransaction creates a new transaction and returns its id.
	CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error)
}

type transactionServiceClient struct {
	cc *grpc.ClientConn
}

func NewTransactionServiceClient(cc *grpc.ClientConn) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := c.cc.Invoke(ctx, "/transactions.TransactionService/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error) {
	out := new(GetTransactionsResponse)
	err := c.cc.Invoke(ctx, "/transactions.TransactionService/GetTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) CreateTransaction(ctx context.Context, in *CreateTransactionRequest, opts ...grpc.CallOption) (*CreateTransactionResponse, error) {
	out := new(CreateTransactionResponse)
	err := c.cc.Invoke(ctx, "/transactions.TransactionService/CreateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
type TransactionServiceServer interface {
	// GetTransaction retrieves the detail of a transaction, selected by its id.
	GetTransaction(context.Context, *GetTransactionRequest) (*Transaction, error)
	// GetTransactions returns a list containing up to 20 transactions.
	GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error)
	// CreateTransaction creates a new transaction and returns its id.
	CreateTransaction(context.Context, *CreateTransactionRequest) (*CreateTransactionResponse, error)
}

func RegisterTransactionServiceServer(s *grpc.Server, srv TransactionServiceServer) {
	s.RegisterService(&_TransactionService_serviceDesc, srv)
}

func _TransactionService_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.TransactionService/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.TransactionService/GetTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetTransactions(ctx, req.(*GetTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transactions.TransactionService/CreateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).CreateTransaction(ctx, req.(*CreateTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TransactionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transactions.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTransaction",
			Handler:    _TransactionService_GetTransaction_Handler,
		},
		{
			MethodName: "GetTransactions",
			Handler:    _TransactionService_GetTransactions_Handler,
		},
		{
			MethodName: "CreateTransaction",
			Handler:    _TransactionService_CreateTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/openbank/openbank/v1/transactions/all.proto",
}
