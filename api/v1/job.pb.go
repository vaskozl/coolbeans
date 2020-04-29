// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: job.proto

package coolbeans_api_v1

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

type JobStateProto int32

const (
	JobStateProto_INITIAL  JobStateProto = 0
	JobStateProto_READY    JobStateProto = 1
	JobStateProto_RESERVED JobStateProto = 2
	JobStateProto_BURIED   JobStateProto = 3
	JobStateProto_DELAYED  JobStateProto = 4
	JobStateProto_DELETED  JobStateProto = 5
)

// Enum value maps for JobStateProto.
var (
	JobStateProto_name = map[int32]string{
		0: "INITIAL",
		1: "READY",
		2: "RESERVED",
		3: "BURIED",
		4: "DELAYED",
		5: "DELETED",
	}
	JobStateProto_value = map[string]int32{
		"INITIAL":  0,
		"READY":    1,
		"RESERVED": 2,
		"BURIED":   3,
		"DELAYED":  4,
		"DELETED":  5,
	}
)

func (x JobStateProto) Enum() *JobStateProto {
	p := new(JobStateProto)
	*p = x
	return p
}

func (x JobStateProto) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobStateProto) Descriptor() protoreflect.EnumDescriptor {
	return file_job_proto_enumTypes[0].Descriptor()
}

func (JobStateProto) Type() protoreflect.EnumType {
	return &file_job_proto_enumTypes[0]
}

func (x JobStateProto) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobStateProto.Descriptor instead.
func (JobStateProto) EnumDescriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{0}
}

type JobProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for job
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Priority is an integer < 2**32. Jobs with smaller priority values will be
	// scheduled before jobs with larger priorities. The most urgent priority is 0;
	// the least urgent priority is 4,294,967,295.
	Priority uint32 `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	// delay is an integer number of seconds to wait before putting the job in
	// the ready queue. The job will be in the "delayed" state during this time.
	// Maximum delay is 2**32-1.
	Delay int64 `protobuf:"varint,3,opt,name=delay,proto3" json:"delay,omitempty"`
	// TTR/time to run -- is an integer number of seconds to allow a worker
	// to run this job. This time is counted from the moment a worker reserves
	// this job. If the worker does not delete, release, or bury the job within
	// <ttr> seconds, the job will time out and the server will release the job.
	// The minimum ttr is 1. If the client sends 0, the server will silently
	// increase the ttr to 1. Maximum ttr is 2**32-1.
	Ttr int32 `protobuf:"varint,4,opt,name=ttr,proto3" json:"ttr,omitempty"`
	// tube_name is the name of the tube associated with this job
	TubeName string `protobuf:"bytes,5,opt,name=tube_name,json=tubeName,proto3" json:"tube_name,omitempty"`
	// created is the time in UTC the job is created
	CreatedAt int64 `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// ready_at is the time in UTC the job is ready
	ReadyAt int64 `protobuf:"varint,7,opt,name=ready_at,json=readyAt,proto3" json:"ready_at,omitempty"`
	// expires_at is the time in UTC, when current reservation expires
	ExpiresAt int64 `protobuf:"varint,8,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	// state is the current state of this job
	State JobStateProto `protobuf:"varint,9,opt,name=state,proto3,enum=coolbeans.api.v1.JobStateProto" json:"state,omitempty"`
	// reserved_by is the identifier of the client which has
	// reserved this job, the value is empty if un-reserved
	ReservedBy string `protobuf:"bytes,10,opt,name=reserved_by,json=reservedBy,proto3" json:"reserved_by,omitempty"`
	// body_size is an integer indicating the size of the job body, not including the
	// trailing "\r\n". This value must be less than max-job-size (default: 2**16)
	BodySize int32 `protobuf:"varint,11,opt,name=body_size,json=bodySize,proto3" json:"body_size,omitempty"`
	// body is the job body -- a sequence of bytes of length BodySize
	Body []byte `protobuf:"bytes,12,opt,name=body,proto3" json:"body,omitempty"`
	// buried_at the clock time when the job is buried
	BuriedAt int64 `protobuf:"varint,13,opt,name=buried_at,json=buriedAt,proto3" json:"buried_at,omitempty"`
}

func (x *JobProto) Reset() {
	*x = JobProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobProto) ProtoMessage() {}

func (x *JobProto) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobProto.ProtoReflect.Descriptor instead.
func (*JobProto) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{0}
}

func (x *JobProto) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *JobProto) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *JobProto) GetDelay() int64 {
	if x != nil {
		return x.Delay
	}
	return 0
}

func (x *JobProto) GetTtr() int32 {
	if x != nil {
		return x.Ttr
	}
	return 0
}

func (x *JobProto) GetTubeName() string {
	if x != nil {
		return x.TubeName
	}
	return ""
}

func (x *JobProto) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *JobProto) GetReadyAt() int64 {
	if x != nil {
		return x.ReadyAt
	}
	return 0
}

func (x *JobProto) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *JobProto) GetState() JobStateProto {
	if x != nil {
		return x.State
	}
	return JobStateProto_INITIAL
}

func (x *JobProto) GetReservedBy() string {
	if x != nil {
		return x.ReservedBy
	}
	return ""
}

func (x *JobProto) GetBodySize() int32 {
	if x != nil {
		return x.BodySize
	}
	return 0
}

func (x *JobProto) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *JobProto) GetBuriedAt() int64 {
	if x != nil {
		return x.BuriedAt
	}
	return 0
}

var File_job_proto protoreflect.FileDescriptor

var file_job_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x6f, 0x6f,
	0x6c, 0x62, 0x65, 0x61, 0x6e, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22, 0xfa, 0x02,
	0x0a, 0x08, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x74, 0x74, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74, 0x72, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x75, 0x62, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x75, 0x62, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65,
	0x61, 0x64, 0x79, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65,
	0x61, 0x64, 0x79, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6f, 0x6c, 0x62, 0x65, 0x61, 0x6e, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x6f, 0x64, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x62, 0x6f, 0x64, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x75, 0x72, 0x69, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x62, 0x75, 0x72, 0x69, 0x65, 0x64, 0x41, 0x74, 0x2a, 0x5b, 0x0a, 0x0d, 0x4a, 0x6f,
	0x62, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x0a, 0x07, 0x49,
	0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x55, 0x52, 0x49, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a,
	0x07, 0x44, 0x45, 0x4c, 0x41, 0x59, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x05, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x3b, 0x63, 0x6f, 0x6f,
	0x6c, 0x62, 0x65, 0x61, 0x6e, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_job_proto_rawDescOnce sync.Once
	file_job_proto_rawDescData = file_job_proto_rawDesc
)

func file_job_proto_rawDescGZIP() []byte {
	file_job_proto_rawDescOnce.Do(func() {
		file_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_job_proto_rawDescData)
	})
	return file_job_proto_rawDescData
}

var file_job_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_job_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_job_proto_goTypes = []interface{}{
	(JobStateProto)(0), // 0: coolbeans.api.v1.JobStateProto
	(*JobProto)(nil),   // 1: coolbeans.api.v1.JobProto
}
var file_job_proto_depIdxs = []int32{
	0, // 0: coolbeans.api.v1.JobProto.state:type_name -> coolbeans.api.v1.JobStateProto
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_job_proto_init() }
func file_job_proto_init() {
	if File_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobProto); i {
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
			RawDescriptor: file_job_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_job_proto_goTypes,
		DependencyIndexes: file_job_proto_depIdxs,
		EnumInfos:         file_job_proto_enumTypes,
		MessageInfos:      file_job_proto_msgTypes,
	}.Build()
	File_job_proto = out.File
	file_job_proto_rawDesc = nil
	file_job_proto_goTypes = nil
	file_job_proto_depIdxs = nil
}
