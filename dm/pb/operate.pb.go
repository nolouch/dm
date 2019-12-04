// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: operate.proto

package pb

import (
	"fmt"

	proto "github.com/gogo/protobuf/proto"

	math "math"

	io "io"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type OperateType int32

const (
	OperateType_OperateUnknown          OperateType = 0
	OperateType_MigrateWorkerRelay      OperateType = 1
	OperateType_UpdateWorkerRelayConfig OperateType = 2
	OperateType_StartTask               OperateType = 3
	OperateType_UpdateMasterConfig      OperateType = 4
	OperateType_OperateTask             OperateType = 5
	OperateType_UpdateTask              OperateType = 6
	OperateType_QueryStatus             OperateType = 7
	OperateType_QueryError              OperateType = 8
	OperateType_ShowDDLLocks            OperateType = 9
	OperateType_UnlockDDLLock           OperateType = 10
	OperateType_BreakWorkerDDLLock      OperateType = 11
	OperateType_SwitchWorkerRelayMaster OperateType = 12
	OperateType_OperateWorkerRelay      OperateType = 13
	OperateType_RefreshWorkerTasks      OperateType = 14
	OperateType_HandleSQLs              OperateType = 15
	OperateType_PurgeWorkerRelay        OperateType = 16
	OperateType_CheckTask               OperateType = 17
)

var OperateType_name = map[int32]string{
	0:  "OperateUnknown",
	1:  "MigrateWorkerRelay",
	2:  "UpdateWorkerRelayConfig",
	3:  "StartTask",
	4:  "UpdateMasterConfig",
	5:  "OperateTask",
	6:  "UpdateTask",
	7:  "QueryStatus",
	8:  "QueryError",
	9:  "ShowDDLLocks",
	10: "UnlockDDLLock",
	11: "BreakWorkerDDLLock",
	12: "SwitchWorkerRelayMaster",
	13: "OperateWorkerRelay",
	14: "RefreshWorkerTasks",
	15: "HandleSQLs",
	16: "PurgeWorkerRelay",
	17: "CheckTask",
}
var OperateType_value = map[string]int32{
	"OperateUnknown":          0,
	"MigrateWorkerRelay":      1,
	"UpdateWorkerRelayConfig": 2,
	"StartTask":               3,
	"UpdateMasterConfig":      4,
	"OperateTask":             5,
	"UpdateTask":              6,
	"QueryStatus":             7,
	"QueryError":              8,
	"ShowDDLLocks":            9,
	"UnlockDDLLock":           10,
	"BreakWorkerDDLLock":      11,
	"SwitchWorkerRelayMaster": 12,
	"OperateWorkerRelay":      13,
	"RefreshWorkerTasks":      14,
	"HandleSQLs":              15,
	"PurgeWorkerRelay":        16,
	"CheckTask":               17,
}

func (x OperateType) String() string {
	return proto.EnumName(OperateType_name, int32(x))
}
func (OperateType) EnumDescriptor() ([]byte, []int) { return fileDescriptorOperate, []int{0} }

type OperateStage int32

const (
	OperateStage_StageUnknown OperateStage = 0
	OperateStage_StageNew     OperateStage = 1
	OperateStage_StageDoing   OperateStage = 2
	OperateStage_StageDone    OperateStage = 3
)

var OperateStage_name = map[int32]string{
	0: "StageUnknown",
	1: "StageNew",
	2: "StageDoing",
	3: "StageDone",
}
var OperateStage_value = map[string]int32{
	"StageUnknown": 0,
	"StageNew":     1,
	"StageDoing":   2,
	"StageDone":    3,
}

func (x OperateStage) String() string {
	return proto.EnumName(OperateStage_name, int32(x))
}
func (OperateStage) EnumDescriptor() ([]byte, []int) { return fileDescriptorOperate, []int{1} }

type Operate struct {
	Tp       OperateType  `protobuf:"varint,1,opt,name=tp,proto3,enum=pb.OperateType" json:"tp,omitempty"`
	Request  []byte       `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Response []byte       `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
	Err      string       `protobuf:"bytes,4,opt,name=err,proto3" json:"err,omitempty"`
	Stage    OperateStage `protobuf:"varint,5,opt,name=stage,proto3,enum=pb.OperateStage" json:"stage,omitempty"`
}

func (m *Operate) Reset()                    { *m = Operate{} }
func (m *Operate) String() string            { return proto.CompactTextString(m) }
func (*Operate) ProtoMessage()               {}
func (*Operate) Descriptor() ([]byte, []int) { return fileDescriptorOperate, []int{0} }

func (m *Operate) GetTp() OperateType {
	if m != nil {
		return m.Tp
	}
	return OperateType_OperateUnknown
}

func (m *Operate) GetRequest() []byte {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Operate) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *Operate) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func (m *Operate) GetStage() OperateStage {
	if m != nil {
		return m.Stage
	}
	return OperateStage_StageUnknown
}

func init() {
	proto.RegisterType((*Operate)(nil), "pb.Operate")
	proto.RegisterEnum("pb.OperateType", OperateType_name, OperateType_value)
	proto.RegisterEnum("pb.OperateStage", OperateStage_name, OperateStage_value)
}
func (m *Operate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Operate) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Tp != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintOperate(dAtA, i, uint64(m.Tp))
	}
	if len(m.Request) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOperate(dAtA, i, uint64(len(m.Request)))
		i += copy(dAtA[i:], m.Request)
	}
	if len(m.Response) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintOperate(dAtA, i, uint64(len(m.Response)))
		i += copy(dAtA[i:], m.Response)
	}
	if len(m.Err) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintOperate(dAtA, i, uint64(len(m.Err)))
		i += copy(dAtA[i:], m.Err)
	}
	if m.Stage != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintOperate(dAtA, i, uint64(m.Stage))
	}
	return i, nil
}

func encodeVarintOperate(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Operate) Size() (n int) {
	var l int
	_ = l
	if m.Tp != 0 {
		n += 1 + sovOperate(uint64(m.Tp))
	}
	l = len(m.Request)
	if l > 0 {
		n += 1 + l + sovOperate(uint64(l))
	}
	l = len(m.Response)
	if l > 0 {
		n += 1 + l + sovOperate(uint64(l))
	}
	l = len(m.Err)
	if l > 0 {
		n += 1 + l + sovOperate(uint64(l))
	}
	if m.Stage != 0 {
		n += 1 + sovOperate(uint64(m.Stage))
	}
	return n
}

func sovOperate(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozOperate(x uint64) (n int) {
	return sovOperate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Operate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOperate
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
			return fmt.Errorf("proto: Operate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Operate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tp", wireType)
			}
			m.Tp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tp |= (OperateType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperate
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
				return ErrInvalidLengthOperate
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Request = append(m.Request[:0], dAtA[iNdEx:postIndex]...)
			if m.Request == nil {
				m.Request = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperate
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
				return ErrInvalidLengthOperate
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Response = append(m.Response[:0], dAtA[iNdEx:postIndex]...)
			if m.Response == nil {
				m.Response = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Err", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperate
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
				return ErrInvalidLengthOperate
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Err = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stage", wireType)
			}
			m.Stage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOperate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Stage |= (OperateStage(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOperate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOperate
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipOperate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOperate
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
					return 0, ErrIntOverflowOperate
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
					return 0, ErrIntOverflowOperate
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
				return 0, ErrInvalidLengthOperate
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowOperate
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
				next, err := skipOperate(dAtA[start:])
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
	ErrInvalidLengthOperate = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOperate   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("operate.proto", fileDescriptorOperate) }

var fileDescriptorOperate = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0xc1, 0x8e, 0x12, 0x41,
	0x10, 0xdd, 0x1e, 0x96, 0x05, 0x6a, 0x07, 0xa8, 0xed, 0x18, 0x9d, 0x68, 0x82, 0xc4, 0x83, 0x21,
	0x7b, 0xe0, 0xa0, 0x7f, 0xb0, 0x8b, 0x89, 0x07, 0x50, 0x77, 0x66, 0x89, 0xe7, 0x86, 0xad, 0x05,
	0x32, 0xa4, 0xbb, 0xed, 0x6e, 0x32, 0xe1, 0x2f, 0x3c, 0xfb, 0x45, 0x1e, 0xfd, 0x04, 0x83, 0xff,
	0x61, 0x4c, 0xf7, 0xcc, 0xca, 0x70, 0xeb, 0xf7, 0xea, 0x15, 0xef, 0xbd, 0x62, 0xa0, 0xab, 0x34,
	0x19, 0xe1, 0x68, 0xac, 0x8d, 0x72, 0x8a, 0x47, 0x7a, 0xf1, 0xe6, 0x07, 0x83, 0xd6, 0xe7, 0x92,
	0xe5, 0xaf, 0x21, 0x72, 0x3a, 0x61, 0x43, 0x36, 0xea, 0xbd, 0xeb, 0x8f, 0xf5, 0x62, 0x5c, 0x0d,
	0xee, 0xf7, 0x9a, 0xd2, 0xc8, 0x69, 0x9e, 0x40, 0xcb, 0xd0, 0xb7, 0x1d, 0x59, 0x97, 0x44, 0x43,
	0x36, 0x8a, 0xd3, 0x27, 0xc8, 0x5f, 0x42, 0xdb, 0x90, 0xd5, 0x4a, 0x5a, 0x4a, 0x1a, 0x61, 0xf4,
	0x1f, 0x73, 0x84, 0x06, 0x19, 0x93, 0x9c, 0x0f, 0xd9, 0xa8, 0x93, 0xfa, 0x27, 0x7f, 0x0b, 0x4d,
	0xeb, 0xc4, 0x8a, 0x92, 0x66, 0xf0, 0xc2, 0x9a, 0x57, 0xe6, 0xf9, 0xb4, 0x1c, 0x5f, 0xff, 0x8d,
	0xe0, 0xb2, 0x96, 0x81, 0x73, 0xe8, 0x55, 0x70, 0x2e, 0x73, 0xa9, 0x0a, 0x89, 0x67, 0xfc, 0x39,
	0xf0, 0xd9, 0x66, 0xe5, 0xb9, 0xaf, 0xca, 0xe4, 0x64, 0x52, 0xda, 0x8a, 0x3d, 0x32, 0xfe, 0x0a,
	0x5e, 0xcc, 0xf5, 0xc3, 0x29, 0x7d, 0xab, 0xe4, 0xe3, 0x66, 0x85, 0x11, 0xef, 0x42, 0x27, 0x73,
	0xc2, 0xb8, 0x7b, 0x61, 0x73, 0x6c, 0xf8, 0xdf, 0x28, 0xb5, 0x33, 0x61, 0x1d, 0x99, 0x4a, 0x76,
	0xce, 0xfb, 0x47, 0x7b, 0x2f, 0x6c, 0xf2, 0x1e, 0x40, 0x29, 0x0c, 0xf8, 0xc2, 0x0b, 0xee, 0x76,
	0x64, 0xf6, 0x99, 0x13, 0x6e, 0x67, 0xb1, 0xe5, 0x05, 0x81, 0xf8, 0x60, 0x8c, 0x32, 0xd8, 0xe6,
	0x08, 0x71, 0xb6, 0x56, 0xc5, 0x64, 0x32, 0x9d, 0xaa, 0x65, 0x6e, 0xb1, 0xc3, 0xaf, 0xa0, 0x3b,
	0x97, 0x5b, 0xb5, 0xcc, 0x2b, 0x0e, 0xc1, 0xdb, 0xdf, 0x18, 0x12, 0x79, 0x99, 0xf4, 0x89, 0xbf,
	0xf4, 0x15, 0xb2, 0x62, 0xe3, 0x96, 0xeb, 0x5a, 0x85, 0x32, 0x21, 0xc6, 0x7e, 0xa9, 0xca, 0x56,
	0xef, 0xdd, 0xf5, 0x7c, 0x4a, 0x8f, 0x86, 0x6c, 0xb5, 0xe5, 0x93, 0x5a, 0xec, 0xf9, 0x64, 0x1f,
	0x85, 0x7c, 0xd8, 0x52, 0x76, 0x37, 0xb5, 0xd8, 0xe7, 0xcf, 0x00, 0xbf, 0xec, 0xcc, 0xea, 0x64,
	0x1b, 0xfd, 0x61, 0x6e, 0xd7, 0xb4, 0xcc, 0x43, 0xbf, 0xab, 0xeb, 0x19, 0xc4, 0xf5, 0xff, 0x25,
	0xd4, 0xf1, 0x8f, 0xe3, 0xf9, 0x63, 0x68, 0x07, 0xe6, 0x13, 0x15, 0xc8, 0xbc, 0x49, 0x40, 0x13,
	0xb5, 0x91, 0xc7, 0x3b, 0x7b, 0x2c, 0x09, 0x1b, 0x37, 0xf8, 0xf3, 0x30, 0x60, 0xbf, 0x0e, 0x03,
	0xf6, 0xfb, 0x30, 0x60, 0xdf, 0xff, 0x0c, 0xce, 0x16, 0x17, 0xe1, 0x4b, 0x7c, 0xff, 0x2f, 0x00,
	0x00, 0xff, 0xff, 0x5d, 0x88, 0xcc, 0x04, 0x9a, 0x02, 0x00, 0x00,
}
