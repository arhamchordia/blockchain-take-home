// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blog/blog/post.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Post struct {
	Title         string    `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body          string    `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Creator       string    `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	Id            uint64    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt     time.Time `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at"`
	LastUpdatedAt time.Time `protobuf:"bytes,6,opt,name=last_updated_at,json=lastUpdatedAt,proto3,stdtime" json:"last_updated_at"`
	Editors       []string  `protobuf:"bytes,7,rep,name=editors,proto3" json:"editors,omitempty"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f060607f92e3b72, []int{0}
}
func (m *Post) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Post.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return m.Size()
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Post) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Post) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Post) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *Post) GetLastUpdatedAt() time.Time {
	if m != nil {
		return m.LastUpdatedAt
	}
	return time.Time{}
}

func (m *Post) GetEditors() []string {
	if m != nil {
		return m.Editors
	}
	return nil
}

func init() {
	proto.RegisterType((*Post)(nil), "blog.blog.Post")
}

func init() { proto.RegisterFile("blog/blog/post.proto", fileDescriptor_8f060607f92e3b72) }

var fileDescriptor_8f060607f92e3b72 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x8c, 0xd3, 0xb4, 0xa5, 0x46, 0x80, 0x6a, 0x75, 0xb0, 0x3a, 0xb8, 0x11, 0x53, 0x25, 0xa4,
	0x44, 0x82, 0x2f, 0x68, 0x59, 0x19, 0x50, 0x05, 0x0b, 0x4b, 0x95, 0x60, 0x13, 0x59, 0x4a, 0x79,
	0x51, 0xfc, 0x2a, 0xd1, 0x2f, 0x60, 0xed, 0x67, 0x75, 0xec, 0xc8, 0x04, 0xa8, 0xf9, 0x11, 0x64,
	0x3b, 0xf9, 0x00, 0x96, 0xd3, 0xdd, 0x3d, 0xdf, 0xf3, 0x49, 0x8f, 0x4e, 0xf2, 0x12, 0x8a, 0xd4,
	0x41, 0x05, 0x06, 0x93, 0xaa, 0x06, 0x04, 0x36, 0xb2, 0x46, 0x62, 0x61, 0x3a, 0x2b, 0x00, 0x8a,
	0x52, 0xa5, 0x6e, 0x90, 0x6f, 0xdf, 0x52, 0xd4, 0x1b, 0x65, 0x30, 0xdb, 0x54, 0xfe, 0xed, 0x74,
	0x52, 0x40, 0x01, 0x8e, 0xa6, 0x96, 0xb5, 0xee, 0x38, 0xdb, 0xe8, 0x77, 0x48, 0x1d, 0x7a, 0xeb,
	0xfa, 0x33, 0xa4, 0xd1, 0x23, 0x18, 0x64, 0x13, 0xda, 0x47, 0x8d, 0xa5, 0xe2, 0x24, 0x26, 0xf3,
	0xd1, 0xca, 0x0b, 0xc6, 0x68, 0x94, 0x83, 0xdc, 0xf1, 0xd0, 0x99, 0x8e, 0x33, 0x4e, 0x87, 0xaf,
	0xb5, 0xca, 0x10, 0x6a, 0xde, 0x73, 0x76, 0x27, 0xd9, 0x25, 0x0d, 0xb5, 0xe4, 0x51, 0x4c, 0xe6,
	0xd1, 0x2a, 0xd4, 0x92, 0xdd, 0x53, 0xea, 0x46, 0x4a, 0xae, 0x33, 0xe4, 0xfd, 0x98, 0xcc, 0xcf,
	0x6f, 0xa7, 0x89, 0xef, 0x9e, 0x74, 0xdd, 0x93, 0xa7, 0xae, 0xfb, 0xf2, 0xec, 0xf0, 0x3d, 0x0b,
	0xf6, 0x3f, 0x33, 0xb2, 0x1a, 0xb5, 0xb9, 0x05, 0xb2, 0x07, 0x7a, 0x55, 0x66, 0x06, 0xd7, 0xdb,
	0x4a, 0x76, 0x9b, 0x06, 0xff, 0xd8, 0x74, 0x61, 0xc3, 0xcf, 0x3e, 0xbb, 0x40, 0x5b, 0x5e, 0x49,
	0x8d, 0x50, 0x1b, 0x3e, 0x8c, 0x7b, 0xb6, 0x7c, 0x2b, 0x97, 0x37, 0x87, 0x93, 0x20, 0xc7, 0x93,
	0x20, 0xbf, 0x27, 0x41, 0xf6, 0x8d, 0x08, 0x8e, 0x8d, 0x08, 0xbe, 0x1a, 0x11, 0xbc, 0x8c, 0xdd,
	0x25, 0x3e, 0xfc, 0x41, 0x70, 0x57, 0x29, 0x93, 0x0f, 0xdc, 0x9f, 0x77, 0x7f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x92, 0x62, 0x54, 0x27, 0xaa, 0x01, 0x00, 0x00,
}

func (m *Post) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Post) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Post) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Editors) > 0 {
		for iNdEx := len(m.Editors) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Editors[iNdEx])
			copy(dAtA[i:], m.Editors[iNdEx])
			i = encodeVarintPost(dAtA, i, uint64(len(m.Editors[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.LastUpdatedAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.LastUpdatedAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintPost(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CreatedAt):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintPost(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x2a
	if m.Id != 0 {
		i = encodeVarintPost(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPost(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintPost(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintPost(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPost(dAtA []byte, offset int, v uint64) int {
	offset -= sovPost(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Post) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovPost(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovPost(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPost(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovPost(uint64(m.Id))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovPost(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.LastUpdatedAt)
	n += 1 + l + sovPost(uint64(l))
	if len(m.Editors) > 0 {
		for _, s := range m.Editors {
			l = len(s)
			n += 1 + l + sovPost(uint64(l))
		}
	}
	return n
}

func sovPost(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPost(x uint64) (n int) {
	return sovPost(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Post) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPost
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Post: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Post: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.LastUpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Editors", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPost
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Editors = append(m.Editors, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPost
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
func skipPost(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPost
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
					return 0, ErrIntOverflowPost
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPost
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
			if length < 0 {
				return 0, ErrInvalidLengthPost
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPost
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPost
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPost        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPost          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPost = fmt.Errorf("proto: unexpected end of group")
)
