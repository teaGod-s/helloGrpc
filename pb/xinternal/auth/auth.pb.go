// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: xinternal/auth/auth.proto

package authipb

import (
	xcommon "gomod.pri/helloGrpc/pb/xcommon"
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

// ListUserGroupByRoleIDReq 列出一个角色下面所有的用户组列表请求包
type ListUserGroupByRoleIDReq struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	CommonIn *xcommon.CommonIn      `protobuf:"bytes,1,opt,name=CommonIn,proto3" json:"CommonIn,omitempty"`
	// 角色ID
	RoleId        int64 `protobuf:"varint,2,opt,name=RoleId,proto3" json:"RoleId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListUserGroupByRoleIDReq) Reset() {
	*x = ListUserGroupByRoleIDReq{}
	mi := &file_xinternal_auth_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUserGroupByRoleIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserGroupByRoleIDReq) ProtoMessage() {}

func (x *ListUserGroupByRoleIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_xinternal_auth_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserGroupByRoleIDReq.ProtoReflect.Descriptor instead.
func (*ListUserGroupByRoleIDReq) Descriptor() ([]byte, []int) {
	return file_xinternal_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *ListUserGroupByRoleIDReq) GetCommonIn() *xcommon.CommonIn {
	if x != nil {
		return x.CommonIn
	}
	return nil
}

func (x *ListUserGroupByRoleIDReq) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

// ListUserGroupByRoleIDResp 列出一个角色下面所有的用户组列表回包
type ListUserGroupByRoleIDResp struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	CommonOut      *xcommon.CommonOut     `protobuf:"bytes,1,opt,name=CommonOut,proto3" json:"CommonOut,omitempty"`
	UserGroupInfos []*xcommon.MarkMeta    `protobuf:"bytes,2,rep,name=UserGroupInfos,proto3" json:"UserGroupInfos,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ListUserGroupByRoleIDResp) Reset() {
	*x = ListUserGroupByRoleIDResp{}
	mi := &file_xinternal_auth_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUserGroupByRoleIDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserGroupByRoleIDResp) ProtoMessage() {}

func (x *ListUserGroupByRoleIDResp) ProtoReflect() protoreflect.Message {
	mi := &file_xinternal_auth_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserGroupByRoleIDResp.ProtoReflect.Descriptor instead.
func (*ListUserGroupByRoleIDResp) Descriptor() ([]byte, []int) {
	return file_xinternal_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *ListUserGroupByRoleIDResp) GetCommonOut() *xcommon.CommonOut {
	if x != nil {
		return x.CommonOut
	}
	return nil
}

func (x *ListUserGroupByRoleIDResp) GetUserGroupInfos() []*xcommon.MarkMeta {
	if x != nil {
		return x.UserGroupInfos
	}
	return nil
}

var File_xinternal_auth_auth_proto protoreflect.FileDescriptor

var file_xinternal_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x19, 0x78, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x74,
	0x68, 0x69, 0x70, 0x62, 0x1a, 0x14, 0x78, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x78, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x78, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x12, 0x2d, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x49, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x78, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x49, 0x6e, 0x52, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x49, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x88, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4f,
	0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x78, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x09, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x12, 0x39, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x78, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x32, 0x6f, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x60, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x21, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x69, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x69, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x42, 0x15, 0x5a, 0x13, 0x70, 0x62, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x69, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_xinternal_auth_auth_proto_rawDescOnce sync.Once
	file_xinternal_auth_auth_proto_rawDescData = file_xinternal_auth_auth_proto_rawDesc
)

func file_xinternal_auth_auth_proto_rawDescGZIP() []byte {
	file_xinternal_auth_auth_proto_rawDescOnce.Do(func() {
		file_xinternal_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_xinternal_auth_auth_proto_rawDescData)
	})
	return file_xinternal_auth_auth_proto_rawDescData
}

var file_xinternal_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_xinternal_auth_auth_proto_goTypes = []any{
	(*ListUserGroupByRoleIDReq)(nil),  // 0: authipb.ListUserGroupByRoleIDReq
	(*ListUserGroupByRoleIDResp)(nil), // 1: authipb.ListUserGroupByRoleIDResp
	(*xcommon.CommonIn)(nil),          // 2: xcommon.CommonIn
	(*xcommon.CommonOut)(nil),         // 3: xcommon.CommonOut
	(*xcommon.MarkMeta)(nil),          // 4: xcommon.MarkMeta
}
var file_xinternal_auth_auth_proto_depIdxs = []int32{
	2, // 0: authipb.ListUserGroupByRoleIDReq.CommonIn:type_name -> xcommon.CommonIn
	3, // 1: authipb.ListUserGroupByRoleIDResp.CommonOut:type_name -> xcommon.CommonOut
	4, // 2: authipb.ListUserGroupByRoleIDResp.UserGroupInfos:type_name -> xcommon.MarkMeta
	0, // 3: authipb.AuthService.ListUserGroupByRoleID:input_type -> authipb.ListUserGroupByRoleIDReq
	1, // 4: authipb.AuthService.ListUserGroupByRoleID:output_type -> authipb.ListUserGroupByRoleIDResp
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_xinternal_auth_auth_proto_init() }
func file_xinternal_auth_auth_proto_init() {
	if File_xinternal_auth_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_xinternal_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_xinternal_auth_auth_proto_goTypes,
		DependencyIndexes: file_xinternal_auth_auth_proto_depIdxs,
		MessageInfos:      file_xinternal_auth_auth_proto_msgTypes,
	}.Build()
	File_xinternal_auth_auth_proto = out.File
	file_xinternal_auth_auth_proto_rawDesc = nil
	file_xinternal_auth_auth_proto_goTypes = nil
	file_xinternal_auth_auth_proto_depIdxs = nil
}