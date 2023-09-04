// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: terrariumpb/service.proto

//go:generate mockery --name TerrariumServiceServer

package terrariumpb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Module struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TaxonomyId      string             `protobuf:"bytes,2,opt,name=taxonomyId,proto3" json:"taxonomyId,omitempty"`
	ModuleName      string             `protobuf:"bytes,3,opt,name=moduleName,proto3" json:"moduleName,omitempty"`
	Source          string             `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Version         string             `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Description     string             `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	InputAttributes []*ModuleAttribute `protobuf:"bytes,7,rep,name=inputAttributes,proto3" json:"inputAttributes,omitempty"`
	Namespace       string             `protobuf:"bytes,8,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *Module) Reset() {
	*x = Module{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terrariumpb_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Module) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Module) ProtoMessage() {}

func (x *Module) ProtoReflect() protoreflect.Message {
	mi := &file_terrariumpb_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Module.ProtoReflect.Descriptor instead.
func (*Module) Descriptor() ([]byte, []int) {
	return file_terrariumpb_service_proto_rawDescGZIP(), []int{0}
}

func (x *Module) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Module) GetTaxonomyId() string {
	if x != nil {
		return x.TaxonomyId
	}
	return ""
}

func (x *Module) GetModuleName() string {
	if x != nil {
		return x.ModuleName
	}
	return ""
}

func (x *Module) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Module) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Module) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Module) GetInputAttributes() []*ModuleAttribute {
	if x != nil {
		return x.InputAttributes
	}
	return nil
}

func (x *Module) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type CompletionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CodeContext string   `protobuf:"bytes,1,opt,name=codeContext,proto3" json:"codeContext,omitempty"`
	Modules     []string `protobuf:"bytes,2,rep,name=modules,proto3" json:"modules,omitempty"`
}

func (x *CompletionRequest) Reset() {
	*x = CompletionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terrariumpb_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletionRequest) ProtoMessage() {}

func (x *CompletionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_terrariumpb_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletionRequest.ProtoReflect.Descriptor instead.
func (*CompletionRequest) Descriptor() ([]byte, []int) {
	return file_terrariumpb_service_proto_rawDescGZIP(), []int{1}
}

func (x *CompletionRequest) GetCodeContext() string {
	if x != nil {
		return x.CodeContext
	}
	return ""
}

func (x *CompletionRequest) GetModules() []string {
	if x != nil {
		return x.Modules
	}
	return nil
}

type CompletionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Suggestions []string `protobuf:"bytes,1,rep,name=suggestions,proto3" json:"suggestions,omitempty"`
}

func (x *CompletionResponse) Reset() {
	*x = CompletionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terrariumpb_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletionResponse) ProtoMessage() {}

func (x *CompletionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_terrariumpb_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletionResponse.ProtoReflect.Descriptor instead.
func (*CompletionResponse) Descriptor() ([]byte, []int) {
	return file_terrariumpb_service_proto_rawDescGZIP(), []int{2}
}

func (x *CompletionResponse) GetSuggestions() []string {
	if x != nil {
		return x.Suggestions
	}
	return nil
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size  int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`   // default 100, maximum no limit
	Index int32 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"` // 0 indexed, default 0
	Total int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"` // READ-ONLY. Total number of pages available. ceil(record_count/page_size)
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terrariumpb_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_terrariumpb_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_terrariumpb_service_proto_rawDescGZIP(), []int{3}
}

func (x *Page) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Page) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Page) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type ListModulesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page             *Page    `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Search           string   `protobuf:"bytes,2,opt,name=search,proto3" json:"search,omitempty"`                      // optional search text
	PopulateMappings bool     `protobuf:"varint,3,opt,name=populateMappings,proto3" json:"populateMappings,omitempty"` // optional. include input-output attribute mappings of the module as well. page size cannot be higher then 100 in order to enable this option.
	Namespaces       []string `protobuf:"bytes,4,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
}

func (x *ListModulesRequest) Reset() {
	*x = ListModulesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terrariumpb_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListModulesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListModulesRequest) ProtoMessage() {}

func (x *ListModulesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_terrariumpb_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListModulesRequest.ProtoReflect.Descriptor instead.
func (*ListModulesRequest) Descriptor() ([]byte, []int) {
	return file_terrariumpb_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListModulesRequest) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListModulesRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *ListModulesRequest) GetPopulateMappings() bool {
	if x != nil {
		return x.PopulateMappings
	}
	return false
}

func (x *ListModulesRequest) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

type ListModulesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Modules []*Module `protobuf:"bytes,1,rep,name=modules,proto3" json:"modules,omitempty"`
	Page    *Page     `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListModulesResponse) Reset() {
	*x = ListModulesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terrariumpb_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListModulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListModulesResponse) ProtoMessage() {}

func (x *ListModulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_terrariumpb_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListModulesResponse.ProtoReflect.Descriptor instead.
func (*ListModulesResponse) Descriptor() ([]byte, []int) {
	return file_terrariumpb_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListModulesResponse) GetModules() []*Module {
	if x != nil {
		return x.Modules
	}
	return nil
}

func (x *ListModulesResponse) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type ListModuleAttributesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleId         string `protobuf:"bytes,1,opt,name=moduleId,proto3" json:"moduleId,omitempty"`
	Page             *Page  `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	Search           string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`                      // optional search text
	PopulateMappings bool   `protobuf:"varint,4,opt,name=populateMappings,proto3" json:"populateMappings,omitempty"` // optional. include input-output attribute mappings of the module as well. page size cannot be higher then 100 in order to enable this option.
}

func (x *ListModuleAttributesRequest) Reset() {
	*x = ListModuleAttributesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terrariumpb_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListModuleAttributesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListModuleAttributesRequest) ProtoMessage() {}

func (x *ListModuleAttributesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_terrariumpb_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListModuleAttributesRequest.ProtoReflect.Descriptor instead.
func (*ListModuleAttributesRequest) Descriptor() ([]byte, []int) {
	return file_terrariumpb_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListModuleAttributesRequest) GetModuleId() string {
	if x != nil {
		return x.ModuleId
	}
	return ""
}

func (x *ListModuleAttributesRequest) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *ListModuleAttributesRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *ListModuleAttributesRequest) GetPopulateMappings() bool {
	if x != nil {
		return x.PopulateMappings
	}
	return false
}

type ListModuleAttributesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attributes []*ModuleAttribute `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Page       *Page              `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListModuleAttributesResponse) Reset() {
	*x = ListModuleAttributesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terrariumpb_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListModuleAttributesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListModuleAttributesResponse) ProtoMessage() {}

func (x *ListModuleAttributesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_terrariumpb_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListModuleAttributesResponse.ProtoReflect.Descriptor instead.
func (*ListModuleAttributesResponse) Descriptor() ([]byte, []int) {
	return file_terrariumpb_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListModuleAttributesResponse) GetAttributes() []*ModuleAttribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *ListModuleAttributesResponse) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type ModuleAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                   string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description            string             `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ParentModule           *Module            `protobuf:"bytes,3,opt,name=parentModule,proto3" json:"parentModule,omitempty"`
	OutputModuleAttributes []*ModuleAttribute `protobuf:"bytes,4,rep,name=outputModuleAttributes,proto3" json:"outputModuleAttributes,omitempty"`
	Optional               bool               `protobuf:"varint,5,opt,name=optional,proto3" json:"optional,omitempty"`
}

func (x *ModuleAttribute) Reset() {
	*x = ModuleAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terrariumpb_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModuleAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuleAttribute) ProtoMessage() {}

func (x *ModuleAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_terrariumpb_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModuleAttribute.ProtoReflect.Descriptor instead.
func (*ModuleAttribute) Descriptor() ([]byte, []int) {
	return file_terrariumpb_service_proto_rawDescGZIP(), []int{8}
}

func (x *ModuleAttribute) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModuleAttribute) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ModuleAttribute) GetParentModule() *Module {
	if x != nil {
		return x.ParentModule
	}
	return nil
}

func (x *ModuleAttribute) GetOutputModuleAttributes() []*ModuleAttribute {
	if x != nil {
		return x.OutputModuleAttributes
	}
	return nil
}

func (x *ModuleAttribute) GetOptional() bool {
	if x != nil {
		return x.Optional
	}
	return false
}

var File_terrariumpb_service_proto protoreflect.FileDescriptor

var file_terrariumpb_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x70, 0x62, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x65, 0x72,
	0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x76, 0x30, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x62, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x27, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x93, 0x02, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x0f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65,
	0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x62, 0x0a, 0x11, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x2b, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x11, 0xfa, 0x42, 0x0e, 0x92, 0x01, 0x0b, 0x08, 0x01, 0x18, 0x01, 0x22, 0x05, 0x72,
	0x03, 0xb0, 0x01, 0x01, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x36, 0x0a,
	0x12, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x46, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xa0, 0x01,
	0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x76,
	0x30, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10,
	0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x22, 0x6d, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61,
	0x72, 0x69, 0x75, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x07,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75,
	0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22,
	0xaf, 0x01, 0x0a, 0x1b, 0x6c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x08, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e,
	0x76, 0x30, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x10, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x73, 0x22, 0x85, 0x01, 0x0a, 0x1c, 0x6c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69,
	0x75, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x26, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0xf4, 0x01, 0x0a, 0x0f, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x65, 0x72, 0x72,
	0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52,
	0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x55, 0x0a,
	0x16, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x16, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x32, 0xe0, 0x03, 0x0a, 0x10, 0x54, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76,
	0x30, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x67, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x74,
	0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x30, 0x2f, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x6e, 0x0a, 0x0e, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61,
	0x72, 0x69, 0x75, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x65, 0x72, 0x72,
	0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x30, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x9a, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12,
	0x29, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x6c,
	0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x74, 0x65, 0x72,
	0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x2e, 0x76, 0x30, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x12, 0x23,
	0x2f, 0x76, 0x30, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x49, 0x64, 0x3d, 0x2a, 0x7d, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x42, 0x56, 0x92, 0x41, 0x1f, 0x12, 0x1a, 0x0a, 0x15, 0x54, 0x65, 0x72, 0x72,
	0x61, 0x72, 0x69, 0x75, 0x6d, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x41, 0x50,
	0x49, 0x32, 0x01, 0x30, 0x2a, 0x01, 0x02, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x64, 0x63, 0x76, 0x72, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61,
	0x72, 0x69, 0x75, 0x6d, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f,
	0x74, 0x65, 0x72, 0x72, 0x61, 0x72, 0x69, 0x75, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_terrariumpb_service_proto_rawDescOnce sync.Once
	file_terrariumpb_service_proto_rawDescData = file_terrariumpb_service_proto_rawDesc
)

func file_terrariumpb_service_proto_rawDescGZIP() []byte {
	file_terrariumpb_service_proto_rawDescOnce.Do(func() {
		file_terrariumpb_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_terrariumpb_service_proto_rawDescData)
	})
	return file_terrariumpb_service_proto_rawDescData
}

var file_terrariumpb_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_terrariumpb_service_proto_goTypes = []interface{}{
	(*Module)(nil),                       // 0: terrarium.v0.Module
	(*CompletionRequest)(nil),            // 1: terrarium.v0.CompletionRequest
	(*CompletionResponse)(nil),           // 2: terrarium.v0.CompletionResponse
	(*Page)(nil),                         // 3: terrarium.v0.Page
	(*ListModulesRequest)(nil),           // 4: terrarium.v0.ListModulesRequest
	(*ListModulesResponse)(nil),          // 5: terrarium.v0.ListModulesResponse
	(*ListModuleAttributesRequest)(nil),  // 6: terrarium.v0.listModuleAttributesRequest
	(*ListModuleAttributesResponse)(nil), // 7: terrarium.v0.listModuleAttributesResponse
	(*ModuleAttribute)(nil),              // 8: terrarium.v0.ModuleAttribute
	(*emptypb.Empty)(nil),                // 9: google.protobuf.Empty
}
var file_terrariumpb_service_proto_depIdxs = []int32{
	8,  // 0: terrarium.v0.Module.inputAttributes:type_name -> terrarium.v0.ModuleAttribute
	3,  // 1: terrarium.v0.ListModulesRequest.page:type_name -> terrarium.v0.Page
	0,  // 2: terrarium.v0.ListModulesResponse.modules:type_name -> terrarium.v0.Module
	3,  // 3: terrarium.v0.ListModulesResponse.page:type_name -> terrarium.v0.Page
	3,  // 4: terrarium.v0.listModuleAttributesRequest.page:type_name -> terrarium.v0.Page
	8,  // 5: terrarium.v0.listModuleAttributesResponse.attributes:type_name -> terrarium.v0.ModuleAttribute
	3,  // 6: terrarium.v0.listModuleAttributesResponse.page:type_name -> terrarium.v0.Page
	0,  // 7: terrarium.v0.ModuleAttribute.parentModule:type_name -> terrarium.v0.Module
	8,  // 8: terrarium.v0.ModuleAttribute.outputModuleAttributes:type_name -> terrarium.v0.ModuleAttribute
	9,  // 9: terrarium.v0.TerrariumService.HealthCheck:input_type -> google.protobuf.Empty
	4,  // 10: terrarium.v0.TerrariumService.ListModules:input_type -> terrarium.v0.ListModulesRequest
	1,  // 11: terrarium.v0.TerrariumService.CodeCompletion:input_type -> terrarium.v0.CompletionRequest
	6,  // 12: terrarium.v0.TerrariumService.ListModuleAttributes:input_type -> terrarium.v0.listModuleAttributesRequest
	9,  // 13: terrarium.v0.TerrariumService.HealthCheck:output_type -> google.protobuf.Empty
	5,  // 14: terrarium.v0.TerrariumService.ListModules:output_type -> terrarium.v0.ListModulesResponse
	2,  // 15: terrarium.v0.TerrariumService.CodeCompletion:output_type -> terrarium.v0.CompletionResponse
	7,  // 16: terrarium.v0.TerrariumService.ListModuleAttributes:output_type -> terrarium.v0.listModuleAttributesResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_terrariumpb_service_proto_init() }
func file_terrariumpb_service_proto_init() {
	if File_terrariumpb_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_terrariumpb_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Module); i {
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
		file_terrariumpb_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompletionRequest); i {
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
		file_terrariumpb_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompletionResponse); i {
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
		file_terrariumpb_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_terrariumpb_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListModulesRequest); i {
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
		file_terrariumpb_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListModulesResponse); i {
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
		file_terrariumpb_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListModuleAttributesRequest); i {
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
		file_terrariumpb_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListModuleAttributesResponse); i {
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
		file_terrariumpb_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModuleAttribute); i {
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
			RawDescriptor: file_terrariumpb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_terrariumpb_service_proto_goTypes,
		DependencyIndexes: file_terrariumpb_service_proto_depIdxs,
		MessageInfos:      file_terrariumpb_service_proto_msgTypes,
	}.Build()
	File_terrariumpb_service_proto = out.File
	file_terrariumpb_service_proto_rawDesc = nil
	file_terrariumpb_service_proto_goTypes = nil
	file_terrariumpb_service_proto_depIdxs = nil
}
