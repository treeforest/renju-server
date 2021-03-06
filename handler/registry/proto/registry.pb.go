// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: registry.proto

package registry

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

//! 注册的事件端口（接口ID）
type Event int32

const (
	Event_REGISTRY            Event = 1 // 注册账号 RegisterRequest RegisterResponse
	Event_GET_SMS_CODE        Event = 2 // 获取手机验证码 GetSmsCodeRequest GetSmsCodeResponse
	Event_PRE_CHECK_CODE      Event = 3 // 验证手机验证码 PreCheckCSmsCodeRequest PreCheckCSmsCodeResponse
	Event_GET_MAIL_CODE       Event = 8 // 获取邮箱验证码 GetMailCodeRequest GetMailCodeResponse
	Event_PRE_CHECK_MAIL_CODE Event = 9 // 验证邮箱验证码 PreCheckCMailCodeRequest PreCheckCMailCodeResponse
)

// Enum value maps for Event.
var (
	Event_name = map[int32]string{
		1: "REGISTRY",
		2: "GET_SMS_CODE",
		3: "PRE_CHECK_CODE",
		8: "GET_MAIL_CODE",
		9: "PRE_CHECK_MAIL_CODE",
	}
	Event_value = map[string]int32{
		"REGISTRY":            1,
		"GET_SMS_CODE":        2,
		"PRE_CHECK_CODE":      3,
		"GET_MAIL_CODE":       8,
		"PRE_CHECK_MAIL_CODE": 9,
	}
)

func (x Event) Enum() *Event {
	p := new(Event)
	*p = x
	return p
}

func (x Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event) Descriptor() protoreflect.EnumDescriptor {
	return file_registry_proto_enumTypes[0].Descriptor()
}

func (Event) Type() protoreflect.EnumType {
	return &file_registry_proto_enumTypes[0]
}

func (x Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Event) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Event(num)
	return nil
}

// Deprecated: Use Event.Descriptor instead.
func (Event) EnumDescriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{0}
}

//! 返回码
type Code int32

const (
	Code_SUCCESS                 Code = 0      //! 成功
	Code_ERR_ACCT_EXISTS         Code = 196609 //注册时账号已存在
	Code_ERR_UNMARSHAL           Code = 196610 //反序列化失败
	Code_ERR_REDIS_FATAL         Code = 196611 //redis异常
	Code_ERR_PHONENUM_FORMAT     Code = 196612 //手机号格式不对
	Code_ERR_CODE_VERIFY         Code = 196613 //验证码不对
	Code_ERR_CODE_TOOFAST        Code = 196614 //请求验证码频率太快
	Code_ERR_EMAIL_FORMAT        Code = 196615 //邮箱号格式不对
	Code_ERR_CODE_EXPIRE         Code = 196616 //验证码过期
	Code_ERR_CODE_OUTTIME        Code = 196617 //验证码错误次数过多
	Code_ERR_THIRDPART_FILL_INFO Code = 196618 //第三方首次登录需要完善信息
	Code_ERR_CODE_UNEXPIRE       Code = 196619 //验证码还没过期
	Code_ERR_EMAIL_BIND_ANOTHER  Code = 196620 //邮箱号绑定了其他账号
	Code_ERR_PWD_EXISTS          Code = 196621 //密码已存在
	Code_ERR_NOT_WHITE_LIST      Code = 196622 //非白名单
	Code_ERR_ACCOUNT_LOCKOUT     Code = 196623 //账号被封
	Code_ERR_SVR_FATAL           Code = 196624 //svr异常
	Code_ERR_BLACK_LIST          Code = 196625 //黑名单
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0:      "SUCCESS",
		196609: "ERR_ACCT_EXISTS",
		196610: "ERR_UNMARSHAL",
		196611: "ERR_REDIS_FATAL",
		196612: "ERR_PHONENUM_FORMAT",
		196613: "ERR_CODE_VERIFY",
		196614: "ERR_CODE_TOOFAST",
		196615: "ERR_EMAIL_FORMAT",
		196616: "ERR_CODE_EXPIRE",
		196617: "ERR_CODE_OUTTIME",
		196618: "ERR_THIRDPART_FILL_INFO",
		196619: "ERR_CODE_UNEXPIRE",
		196620: "ERR_EMAIL_BIND_ANOTHER",
		196621: "ERR_PWD_EXISTS",
		196622: "ERR_NOT_WHITE_LIST",
		196623: "ERR_ACCOUNT_LOCKOUT",
		196624: "ERR_SVR_FATAL",
		196625: "ERR_BLACK_LIST",
	}
	Code_value = map[string]int32{
		"SUCCESS":                 0,
		"ERR_ACCT_EXISTS":         196609,
		"ERR_UNMARSHAL":           196610,
		"ERR_REDIS_FATAL":         196611,
		"ERR_PHONENUM_FORMAT":     196612,
		"ERR_CODE_VERIFY":         196613,
		"ERR_CODE_TOOFAST":        196614,
		"ERR_EMAIL_FORMAT":        196615,
		"ERR_CODE_EXPIRE":         196616,
		"ERR_CODE_OUTTIME":        196617,
		"ERR_THIRDPART_FILL_INFO": 196618,
		"ERR_CODE_UNEXPIRE":       196619,
		"ERR_EMAIL_BIND_ANOTHER":  196620,
		"ERR_PWD_EXISTS":          196621,
		"ERR_NOT_WHITE_LIST":      196622,
		"ERR_ACCOUNT_LOCKOUT":     196623,
		"ERR_SVR_FATAL":           196624,
		"ERR_BLACK_LIST":          196625,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_registry_proto_enumTypes[1].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_registry_proto_enumTypes[1]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Code) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Code(num)
	return nil
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{1}
}

//! 账号类型
type AccountType int32

const (
	AccountType_PHONE   AccountType = 1 //! 电话
	AccountType_EMAIL   AccountType = 2 //! 邮箱
	AccountType_VISITOR AccountType = 3 //! 游客(GUID)
)

// Enum value maps for AccountType.
var (
	AccountType_name = map[int32]string{
		1: "PHONE",
		2: "EMAIL",
		3: "VISITOR",
	}
	AccountType_value = map[string]int32{
		"PHONE":   1,
		"EMAIL":   2,
		"VISITOR": 3,
	}
)

func (x AccountType) Enum() *AccountType {
	p := new(AccountType)
	*p = x
	return p
}

func (x AccountType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountType) Descriptor() protoreflect.EnumDescriptor {
	return file_registry_proto_enumTypes[2].Descriptor()
}

func (AccountType) Type() protoreflect.EnumType {
	return &file_registry_proto_enumTypes[2]
}

func (x AccountType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *AccountType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = AccountType(num)
	return nil
}

// Deprecated: Use AccountType.Descriptor instead.
func (AccountType) EnumDescriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{2}
}

//! 注册
type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountType *AccountType `protobuf:"varint,1,req,name=account_type,json=accountType,enum=registry.AccountType" json:"account_type,omitempty"` // 注册的账号类型
	UserName    *string      `protobuf:"bytes,2,opt,name=userName" json:"userName,omitempty"`                                                     // 用户名（要求唯一）
	Password    *string      `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`                                                     // 密码（当前仅前端判断格式要求）
	Account     *string      `protobuf:"bytes,4,opt,name=account" json:"account,omitempty"`                                                       // 注册的账号（填写手机号/邮箱号，与AccountType匹配）
	Code        *string      `protobuf:"bytes,5,opt,name=code" json:"code,omitempty"`                                                             // 验证码
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetAccountType() AccountType {
	if x != nil && x.AccountType != nil {
		return *x.AccountType
	}
	return AccountType_PHONE
}

func (x *RegisterRequest) GetUserName() string {
	if x != nil && x.UserName != nil {
		return *x.UserName
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *RegisterRequest) GetAccount() string {
	if x != nil && x.Account != nil {
		return *x.Account
	}
	return ""
}

func (x *RegisterRequest) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	UserId *int64 `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"` //! 注册成功后，服务端返回的唯一ID
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetResult() int32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *RegisterResponse) GetUserId() int64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

//! 获取验证码手机
type GetSmsCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNum *string `protobuf:"bytes,1,opt,name=phone_num,json=phoneNum" json:"phone_num,omitempty"` // 手机号
}

func (x *GetSmsCodeRequest) Reset() {
	*x = GetSmsCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSmsCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSmsCodeRequest) ProtoMessage() {}

func (x *GetSmsCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSmsCodeRequest.ProtoReflect.Descriptor instead.
func (*GetSmsCodeRequest) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{2}
}

func (x *GetSmsCodeRequest) GetPhoneNum() string {
	if x != nil && x.PhoneNum != nil {
		return *x.PhoneNum
	}
	return ""
}

type GetSmsCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result     *int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	ExpireTime *int64 `protobuf:"varint,2,opt,name=ExpireTime" json:"ExpireTime,omitempty"` // 验证码过期时间，目前返回300s
}

func (x *GetSmsCodeResponse) Reset() {
	*x = GetSmsCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSmsCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSmsCodeResponse) ProtoMessage() {}

func (x *GetSmsCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSmsCodeResponse.ProtoReflect.Descriptor instead.
func (*GetSmsCodeResponse) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{3}
}

func (x *GetSmsCodeResponse) GetResult() int32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *GetSmsCodeResponse) GetExpireTime() int64 {
	if x != nil && x.ExpireTime != nil {
		return *x.ExpireTime
	}
	return 0
}

// 提前验证短信密码是否正确，正确之后再进行注册请求
type PreCheckSmsCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNum *string `protobuf:"bytes,1,opt,name=phone_num,json=phoneNum" json:"phone_num,omitempty"` // 手机号
	SmsCode  *string `protobuf:"bytes,2,opt,name=sms_code,json=smsCode" json:"sms_code,omitempty"`    // 短信验证码
}

func (x *PreCheckSmsCodeRequest) Reset() {
	*x = PreCheckSmsCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreCheckSmsCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreCheckSmsCodeRequest) ProtoMessage() {}

func (x *PreCheckSmsCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreCheckSmsCodeRequest.ProtoReflect.Descriptor instead.
func (*PreCheckSmsCodeRequest) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{4}
}

func (x *PreCheckSmsCodeRequest) GetPhoneNum() string {
	if x != nil && x.PhoneNum != nil {
		return *x.PhoneNum
	}
	return ""
}

func (x *PreCheckSmsCodeRequest) GetSmsCode() string {
	if x != nil && x.SmsCode != nil {
		return *x.SmsCode
	}
	return ""
}

type PreCheckSmsCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (x *PreCheckSmsCodeResponse) Reset() {
	*x = PreCheckSmsCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreCheckSmsCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreCheckSmsCodeResponse) ProtoMessage() {}

func (x *PreCheckSmsCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreCheckSmsCodeResponse.ProtoReflect.Descriptor instead.
func (*PreCheckSmsCodeResponse) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{5}
}

func (x *PreCheckSmsCodeResponse) GetResult() int32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

//! 获取验证码邮箱
type GetMailCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email *string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"` // 邮箱
}

func (x *GetMailCodeRequest) Reset() {
	*x = GetMailCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMailCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMailCodeRequest) ProtoMessage() {}

func (x *GetMailCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMailCodeRequest.ProtoReflect.Descriptor instead.
func (*GetMailCodeRequest) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{6}
}

func (x *GetMailCodeRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

type GetMailCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result     *int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	ExpireTime *int64 `protobuf:"varint,2,opt,name=ExpireTime" json:"ExpireTime,omitempty"`
}

func (x *GetMailCodeResponse) Reset() {
	*x = GetMailCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMailCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMailCodeResponse) ProtoMessage() {}

func (x *GetMailCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMailCodeResponse.ProtoReflect.Descriptor instead.
func (*GetMailCodeResponse) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{7}
}

func (x *GetMailCodeResponse) GetResult() int32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *GetMailCodeResponse) GetExpireTime() int64 {
	if x != nil && x.ExpireTime != nil {
		return *x.ExpireTime
	}
	return 0
}

// 提前验证邮箱密码是否正确，正确之后再进行注册请求
type PreCheckCMailCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email     *string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`                          // 邮箱
	EmailCode *string `protobuf:"bytes,2,opt,name=email_code,json=emailCode" json:"email_code,omitempty"` // 邮箱验证码
}

func (x *PreCheckCMailCodeRequest) Reset() {
	*x = PreCheckCMailCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreCheckCMailCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreCheckCMailCodeRequest) ProtoMessage() {}

func (x *PreCheckCMailCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreCheckCMailCodeRequest.ProtoReflect.Descriptor instead.
func (*PreCheckCMailCodeRequest) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{8}
}

func (x *PreCheckCMailCodeRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *PreCheckCMailCodeRequest) GetEmailCode() string {
	if x != nil && x.EmailCode != nil {
		return *x.EmailCode
	}
	return ""
}

type PreCheckCMailCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (x *PreCheckCMailCodeResponse) Reset() {
	*x = PreCheckCMailCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreCheckCMailCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreCheckCMailCodeResponse) ProtoMessage() {}

func (x *PreCheckCMailCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_registry_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreCheckCMailCodeResponse.ProtoReflect.Descriptor instead.
func (*PreCheckCMailCodeResponse) Descriptor() ([]byte, []int) {
	return file_registry_proto_rawDescGZIP(), []int{9}
}

func (x *PreCheckCMailCodeResponse) GetResult() int32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

var File_registry_proto protoreflect.FileDescriptor

var file_registry_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x22, 0xb1, 0x01, 0x0a, 0x0f, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38,
	0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x42,
	0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x30, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x22, 0x4c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x6d, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x50, 0x0a, 0x16, 0x50, 0x72, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x6d,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6d, 0x73,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6d, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0x31, 0x0a, 0x17, 0x50, 0x72, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x53, 0x6d, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x69, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0x4d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x18, 0x50, 0x72, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x4d,
	0x61, 0x69, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0x33, 0x0a, 0x19, 0x50, 0x72, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43,
	0x4d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2a, 0x67, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x52, 0x59, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x47, 0x45, 0x54, 0x5f, 0x53, 0x4d, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10,
	0x02, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x52, 0x45, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x5f, 0x43,
	0x4f, 0x44, 0x45, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x45, 0x54, 0x5f, 0x4d, 0x41, 0x49,
	0x4c, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x08, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x52, 0x45, 0x5f,
	0x43, 0x48, 0x45, 0x43, 0x4b, 0x5f, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10,
	0x09, 0x2a, 0xb3, 0x03, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x0f, 0x45, 0x52, 0x52, 0x5f, 0x41,
	0x43, 0x43, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x81, 0x80, 0x0c, 0x12, 0x13,
	0x0a, 0x0d, 0x45, 0x52, 0x52, 0x5f, 0x55, 0x4e, 0x4d, 0x41, 0x52, 0x53, 0x48, 0x41, 0x4c, 0x10,
	0x82, 0x80, 0x0c, 0x12, 0x15, 0x0a, 0x0f, 0x45, 0x52, 0x52, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x53,
	0x5f, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x83, 0x80, 0x0c, 0x12, 0x19, 0x0a, 0x13, 0x45, 0x52,
	0x52, 0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41,
	0x54, 0x10, 0x84, 0x80, 0x0c, 0x12, 0x15, 0x0a, 0x0f, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x10, 0x85, 0x80, 0x0c, 0x12, 0x16, 0x0a, 0x10,
	0x45, 0x52, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x4f, 0x4f, 0x46, 0x41, 0x53, 0x54,
	0x10, 0x86, 0x80, 0x0c, 0x12, 0x16, 0x0a, 0x10, 0x45, 0x52, 0x52, 0x5f, 0x45, 0x4d, 0x41, 0x49,
	0x4c, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x10, 0x87, 0x80, 0x0c, 0x12, 0x15, 0x0a, 0x0f,
	0x45, 0x52, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x10,
	0x88, 0x80, 0x0c, 0x12, 0x16, 0x0a, 0x10, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x4f, 0x55, 0x54, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x89, 0x80, 0x0c, 0x12, 0x1d, 0x0a, 0x17, 0x45,
	0x52, 0x52, 0x5f, 0x54, 0x48, 0x49, 0x52, 0x44, 0x50, 0x41, 0x52, 0x54, 0x5f, 0x46, 0x49, 0x4c,
	0x4c, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x8a, 0x80, 0x0c, 0x12, 0x17, 0x0a, 0x11, 0x45, 0x52,
	0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x10,
	0x8b, 0x80, 0x0c, 0x12, 0x1c, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c,
	0x5f, 0x42, 0x49, 0x4e, 0x44, 0x5f, 0x41, 0x4e, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x8c, 0x80,
	0x0c, 0x12, 0x14, 0x0a, 0x0e, 0x45, 0x52, 0x52, 0x5f, 0x50, 0x57, 0x44, 0x5f, 0x45, 0x58, 0x49,
	0x53, 0x54, 0x53, 0x10, 0x8d, 0x80, 0x0c, 0x12, 0x18, 0x0a, 0x12, 0x45, 0x52, 0x52, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x57, 0x48, 0x49, 0x54, 0x45, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x8e, 0x80,
	0x0c, 0x12, 0x19, 0x0a, 0x13, 0x45, 0x52, 0x52, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54,
	0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x4f, 0x55, 0x54, 0x10, 0x8f, 0x80, 0x0c, 0x12, 0x13, 0x0a, 0x0d,
	0x45, 0x52, 0x52, 0x5f, 0x53, 0x56, 0x52, 0x5f, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x90, 0x80,
	0x0c, 0x12, 0x14, 0x0a, 0x0e, 0x45, 0x52, 0x52, 0x5f, 0x42, 0x4c, 0x41, 0x43, 0x4b, 0x5f, 0x4c,
	0x49, 0x53, 0x54, 0x10, 0x91, 0x80, 0x0c, 0x2a, 0x30, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x10,
	0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07,
	0x56, 0x49, 0x53, 0x49, 0x54, 0x4f, 0x52, 0x10, 0x03,
}

var (
	file_registry_proto_rawDescOnce sync.Once
	file_registry_proto_rawDescData = file_registry_proto_rawDesc
)

func file_registry_proto_rawDescGZIP() []byte {
	file_registry_proto_rawDescOnce.Do(func() {
		file_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_registry_proto_rawDescData)
	})
	return file_registry_proto_rawDescData
}

var file_registry_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_registry_proto_goTypes = []interface{}{
	(Event)(0),                        // 0: registry.Event
	(Code)(0),                         // 1: registry.Code
	(AccountType)(0),                  // 2: registry.AccountType
	(*RegisterRequest)(nil),           // 3: registry.RegisterRequest
	(*RegisterResponse)(nil),          // 4: registry.RegisterResponse
	(*GetSmsCodeRequest)(nil),         // 5: registry.GetSmsCodeRequest
	(*GetSmsCodeResponse)(nil),        // 6: registry.GetSmsCodeResponse
	(*PreCheckSmsCodeRequest)(nil),    // 7: registry.PreCheckSmsCodeRequest
	(*PreCheckSmsCodeResponse)(nil),   // 8: registry.PreCheckSmsCodeResponse
	(*GetMailCodeRequest)(nil),        // 9: registry.GetMailCodeRequest
	(*GetMailCodeResponse)(nil),       // 10: registry.GetMailCodeResponse
	(*PreCheckCMailCodeRequest)(nil),  // 11: registry.PreCheckCMailCodeRequest
	(*PreCheckCMailCodeResponse)(nil), // 12: registry.PreCheckCMailCodeResponse
}
var file_registry_proto_depIdxs = []int32{
	2, // 0: registry.RegisterRequest.account_type:type_name -> registry.AccountType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_registry_proto_init() }
func file_registry_proto_init() {
	if File_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_registry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSmsCodeRequest); i {
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
		file_registry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSmsCodeResponse); i {
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
		file_registry_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreCheckSmsCodeRequest); i {
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
		file_registry_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreCheckSmsCodeResponse); i {
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
		file_registry_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMailCodeRequest); i {
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
		file_registry_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMailCodeResponse); i {
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
		file_registry_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreCheckCMailCodeRequest); i {
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
		file_registry_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreCheckCMailCodeResponse); i {
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
			RawDescriptor: file_registry_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_registry_proto_goTypes,
		DependencyIndexes: file_registry_proto_depIdxs,
		EnumInfos:         file_registry_proto_enumTypes,
		MessageInfos:      file_registry_proto_msgTypes,
	}.Build()
	File_registry_proto = out.File
	file_registry_proto_rawDesc = nil
	file_registry_proto_goTypes = nil
	file_registry_proto_depIdxs = nil
}
