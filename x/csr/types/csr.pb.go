// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: canto/csr/v1/csr.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// The CSR struct is a wrapper to all of the metadata associated with a given Contract Secured Revenue registration. It maintains the pool
// where all of the fees are being sent to, the deployer that is responsible for all deployments, and the set of dApps (smart contracts)
// that are registered with this pool.
type CSR struct {
	// Contracts is the list of all EVM address that are registered to this pool (EVM addresses)
	Contracts []string `protobuf:"bytes,1,rep,name=contracts,proto3" json:"contracts,omitempty"`
	// The NFT id which this CSR corresponds to
	Id uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// The account which will be accumulating rewards for this CSR (bech32 Canto address)
	Beneficiary string `protobuf:"bytes,3,opt,name=beneficiary,proto3" json:"beneficiary,omitempty"`
}

func (m *CSR) Reset()         { *m = CSR{} }
func (m *CSR) String() string { return proto.CompactTextString(m) }
func (*CSR) ProtoMessage()    {}
func (*CSR) Descriptor() ([]byte, []int) {
	return fileDescriptor_57c53cea3d443afa, []int{0}
}
func (m *CSR) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CSR) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CSR.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CSR) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CSR.Merge(m, src)
}
func (m *CSR) XXX_Size() int {
	return m.Size()
}
func (m *CSR) XXX_DiscardUnknown() {
	xxx_messageInfo_CSR.DiscardUnknown(m)
}

var xxx_messageInfo_CSR proto.InternalMessageInfo

func (m *CSR) GetContracts() []string {
	if m != nil {
		return m.Contracts
	}
	return nil
}

func (m *CSR) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CSR) GetBeneficiary() string {
	if m != nil {
		return m.Beneficiary
	}
	return ""
}

// Keeps track of all of the CSRs, primarily used to (un)marshal data
type CSRs struct {
	Csrs []*CSR `protobuf:"bytes,1,rep,name=csrs,proto3" json:"csrs,omitempty"`
}

func (m *CSRs) Reset()         { *m = CSRs{} }
func (m *CSRs) String() string { return proto.CompactTextString(m) }
func (*CSRs) ProtoMessage()    {}
func (*CSRs) Descriptor() ([]byte, []int) {
	return fileDescriptor_57c53cea3d443afa, []int{1}
}
func (m *CSRs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CSRs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CSRs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CSRs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CSRs.Merge(m, src)
}
func (m *CSRs) XXX_Size() int {
	return m.Size()
}
func (m *CSRs) XXX_DiscardUnknown() {
	xxx_messageInfo_CSRs.DiscardUnknown(m)
}

var xxx_messageInfo_CSRs proto.InternalMessageInfo

func (m *CSRs) GetCsrs() []*CSR {
	if m != nil {
		return m.Csrs
	}
	return nil
}

func init() {
	proto.RegisterType((*CSR)(nil), "canto.csr.v1.CSR")
	proto.RegisterType((*CSRs)(nil), "canto.csr.v1.CSRs")
}

func init() { proto.RegisterFile("canto/csr/v1/csr.proto", fileDescriptor_57c53cea3d443afa) }

var fileDescriptor_57c53cea3d443afa = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x4e, 0xcc, 0x2b,
	0xc9, 0xd7, 0x4f, 0x2e, 0x2e, 0xd2, 0x2f, 0x33, 0x04, 0x51, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0x3c, 0x60, 0x71, 0x3d, 0x90, 0x40, 0x99, 0xa1, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x58,
	0x42, 0x1f, 0xc4, 0x82, 0xa8, 0x51, 0x0a, 0xe5, 0x62, 0x76, 0x0e, 0x0e, 0x12, 0x92, 0xe1, 0xe2,
	0x4c, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0x29, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x0c,
	0x42, 0x08, 0x08, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x29, 0x30, 0x6a, 0xb0, 0x04, 0x31,
	0x65, 0xa6, 0x08, 0x29, 0x70, 0x71, 0x27, 0xa5, 0xe6, 0xa5, 0xa6, 0x65, 0x26, 0x67, 0x26, 0x16,
	0x55, 0x4a, 0x30, 0x2b, 0x30, 0x6a, 0x70, 0x06, 0x21, 0x0b, 0x29, 0xe9, 0x72, 0xb1, 0x38, 0x07,
	0x07, 0x15, 0x0b, 0xa9, 0x72, 0xb1, 0x24, 0x17, 0x17, 0x41, 0x8c, 0xe4, 0x36, 0x12, 0xd4, 0x43,
	0x76, 0x91, 0x9e, 0x73, 0x70, 0x50, 0x10, 0x58, 0xda, 0xc9, 0xfd, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86,
	0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x74, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73,
	0xf5, 0x9d, 0x41, 0x9a, 0x75, 0xfd, 0x52, 0x4b, 0xca, 0xf3, 0x8b, 0xb2, 0x21, 0x3c, 0xfd, 0x32,
	0x23, 0xfd, 0x0a, 0xb0, 0xcf, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xbe, 0x32, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x33, 0x0a, 0x6d, 0x13, 0x01, 0x00, 0x00,
}

func (m *CSR) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CSR) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CSR) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Beneficiary) > 0 {
		i -= len(m.Beneficiary)
		copy(dAtA[i:], m.Beneficiary)
		i = encodeVarintCsr(dAtA, i, uint64(len(m.Beneficiary)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintCsr(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Contracts) > 0 {
		for iNdEx := len(m.Contracts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Contracts[iNdEx])
			copy(dAtA[i:], m.Contracts[iNdEx])
			i = encodeVarintCsr(dAtA, i, uint64(len(m.Contracts[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CSRs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CSRs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CSRs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Csrs) > 0 {
		for iNdEx := len(m.Csrs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Csrs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCsr(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintCsr(dAtA []byte, offset int, v uint64) int {
	offset -= sovCsr(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CSR) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Contracts) > 0 {
		for _, s := range m.Contracts {
			l = len(s)
			n += 1 + l + sovCsr(uint64(l))
		}
	}
	if m.Id != 0 {
		n += 1 + sovCsr(uint64(m.Id))
	}
	l = len(m.Beneficiary)
	if l > 0 {
		n += 1 + l + sovCsr(uint64(l))
	}
	return n
}

func (m *CSRs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Csrs) > 0 {
		for _, e := range m.Csrs {
			l = e.Size()
			n += 1 + l + sovCsr(uint64(l))
		}
	}
	return n
}

func sovCsr(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCsr(x uint64) (n int) {
	return sovCsr(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CSR) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCsr
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
			return fmt.Errorf("proto: CSR: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CSR: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contracts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCsr
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
				return ErrInvalidLengthCsr
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCsr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contracts = append(m.Contracts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCsr
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Beneficiary", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCsr
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
				return ErrInvalidLengthCsr
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCsr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Beneficiary = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCsr(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCsr
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
func (m *CSRs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCsr
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
			return fmt.Errorf("proto: CSRs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CSRs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Csrs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCsr
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
				return ErrInvalidLengthCsr
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCsr
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Csrs = append(m.Csrs, &CSR{})
			if err := m.Csrs[len(m.Csrs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCsr(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCsr
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
func skipCsr(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCsr
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
					return 0, ErrIntOverflowCsr
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
					return 0, ErrIntOverflowCsr
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
				return 0, ErrInvalidLengthCsr
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCsr
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCsr
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCsr        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCsr          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCsr = fmt.Errorf("proto: unexpected end of group")
)
