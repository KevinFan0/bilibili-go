// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v1/BannerMng.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type BannerMngGetNewBannerReq struct {
	// 业务id，首页为0，web主站hove为1
	Platform int64 `protobuf:"varint,1,opt,name=platform,proto3" json:"platform"`
	// 第几帧，0表示取全部
	Position int64 `protobuf:"varint,2,opt,name=position,proto3" json:"position"`
	// 平台
	UserPlatform string `protobuf:"bytes,3,opt,name=userPlatform,proto3" json:"userPlatform"`
	// 设备
	UserDevice string `protobuf:"bytes,4,opt,name=userDevice,proto3" json:"userDevice"`
	// 版本号
	Build int64 `protobuf:"varint,5,opt,name=build,proto3" json:"build"`
	// 是否返回所有版本配置：是，则不论build传的值，全返回banner和版本配置
	ReturnBuilds int64 `protobuf:"varint,6,opt,name=returnBuilds,proto3" json:"returnBuilds"`
}

func (m *BannerMngGetNewBannerReq) Reset()         { *m = BannerMngGetNewBannerReq{} }
func (m *BannerMngGetNewBannerReq) String() string { return proto.CompactTextString(m) }
func (*BannerMngGetNewBannerReq) ProtoMessage()    {}
func (*BannerMngGetNewBannerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_BannerMng_9505b5c7b0530020, []int{0}
}
func (m *BannerMngGetNewBannerReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BannerMngGetNewBannerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BannerMngGetNewBannerReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BannerMngGetNewBannerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BannerMngGetNewBannerReq.Merge(dst, src)
}
func (m *BannerMngGetNewBannerReq) XXX_Size() int {
	return m.Size()
}
func (m *BannerMngGetNewBannerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BannerMngGetNewBannerReq.DiscardUnknown(m)
}

var xxx_messageInfo_BannerMngGetNewBannerReq proto.InternalMessageInfo

func (m *BannerMngGetNewBannerReq) GetPlatform() int64 {
	if m != nil {
		return m.Platform
	}
	return 0
}

func (m *BannerMngGetNewBannerReq) GetPosition() int64 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *BannerMngGetNewBannerReq) GetUserPlatform() string {
	if m != nil {
		return m.UserPlatform
	}
	return ""
}

func (m *BannerMngGetNewBannerReq) GetUserDevice() string {
	if m != nil {
		return m.UserDevice
	}
	return ""
}

func (m *BannerMngGetNewBannerReq) GetBuild() int64 {
	if m != nil {
		return m.Build
	}
	return 0
}

func (m *BannerMngGetNewBannerReq) GetReturnBuilds() int64 {
	if m != nil {
		return m.ReturnBuilds
	}
	return 0
}

type BannerMngGetNewBannerResp struct {
	// code
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	// msg
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
	//
	Data []*BannerMngGetNewBannerResp_NewBanner `protobuf:"bytes,3,rep,name=data" json:"data"`
}

func (m *BannerMngGetNewBannerResp) Reset()         { *m = BannerMngGetNewBannerResp{} }
func (m *BannerMngGetNewBannerResp) String() string { return proto.CompactTextString(m) }
func (*BannerMngGetNewBannerResp) ProtoMessage()    {}
func (*BannerMngGetNewBannerResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_BannerMng_9505b5c7b0530020, []int{1}
}
func (m *BannerMngGetNewBannerResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BannerMngGetNewBannerResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BannerMngGetNewBannerResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BannerMngGetNewBannerResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BannerMngGetNewBannerResp.Merge(dst, src)
}
func (m *BannerMngGetNewBannerResp) XXX_Size() int {
	return m.Size()
}
func (m *BannerMngGetNewBannerResp) XXX_DiscardUnknown() {
	xxx_messageInfo_BannerMngGetNewBannerResp.DiscardUnknown(m)
}

var xxx_messageInfo_BannerMngGetNewBannerResp proto.InternalMessageInfo

func (m *BannerMngGetNewBannerResp) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BannerMngGetNewBannerResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *BannerMngGetNewBannerResp) GetData() []*BannerMngGetNewBannerResp_NewBanner {
	if m != nil {
		return m.Data
	}
	return nil
}

type BannerMngGetNewBannerResp_NewBanner struct {
	// banner id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 图片地址
	Pic string `protobuf:"bytes,2,opt,name=pic,proto3" json:"pic"`
	// 图片地址
	Img string `protobuf:"bytes,3,opt,name=img,proto3" json:"img"`
	// 跳转链接
	Link string `protobuf:"bytes,4,opt,name=link,proto3" json:"link"`
	// 标题
	Title string `protobuf:"bytes,5,opt,name=title,proto3" json:"title"`
	// 第几帧
	Position string `protobuf:"bytes,6,opt,name=position,proto3" json:"position"`
	// 权重
	SortNum string `protobuf:"bytes,7,opt,name=sort_num,json=sortNum,proto3" json:"sort_num"`
	// 注释
	Remark string `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark"`
}

func (m *BannerMngGetNewBannerResp_NewBanner) Reset()         { *m = BannerMngGetNewBannerResp_NewBanner{} }
func (m *BannerMngGetNewBannerResp_NewBanner) String() string { return proto.CompactTextString(m) }
func (*BannerMngGetNewBannerResp_NewBanner) ProtoMessage()    {}
func (*BannerMngGetNewBannerResp_NewBanner) Descriptor() ([]byte, []int) {
	return fileDescriptor_BannerMng_9505b5c7b0530020, []int{1, 0}
}
func (m *BannerMngGetNewBannerResp_NewBanner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BannerMngGetNewBannerResp_NewBanner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BannerMngGetNewBannerResp_NewBanner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BannerMngGetNewBannerResp_NewBanner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BannerMngGetNewBannerResp_NewBanner.Merge(dst, src)
}
func (m *BannerMngGetNewBannerResp_NewBanner) XXX_Size() int {
	return m.Size()
}
func (m *BannerMngGetNewBannerResp_NewBanner) XXX_DiscardUnknown() {
	xxx_messageInfo_BannerMngGetNewBannerResp_NewBanner.DiscardUnknown(m)
}

var xxx_messageInfo_BannerMngGetNewBannerResp_NewBanner proto.InternalMessageInfo

func (m *BannerMngGetNewBannerResp_NewBanner) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BannerMngGetNewBannerResp_NewBanner) GetPic() string {
	if m != nil {
		return m.Pic
	}
	return ""
}

func (m *BannerMngGetNewBannerResp_NewBanner) GetImg() string {
	if m != nil {
		return m.Img
	}
	return ""
}

func (m *BannerMngGetNewBannerResp_NewBanner) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *BannerMngGetNewBannerResp_NewBanner) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *BannerMngGetNewBannerResp_NewBanner) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *BannerMngGetNewBannerResp_NewBanner) GetSortNum() string {
	if m != nil {
		return m.SortNum
	}
	return ""
}

func (m *BannerMngGetNewBannerResp_NewBanner) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func init() {
	proto.RegisterType((*BannerMngGetNewBannerReq)(nil), "room_ex.v1.BannerMngGetNewBannerReq")
	proto.RegisterType((*BannerMngGetNewBannerResp)(nil), "room_ex.v1.BannerMngGetNewBannerResp")
	proto.RegisterType((*BannerMngGetNewBannerResp_NewBanner)(nil), "room_ex.v1.BannerMngGetNewBannerResp.NewBanner")
}
func (m *BannerMngGetNewBannerReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BannerMngGetNewBannerReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Platform != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintBannerMng(dAtA, i, uint64(m.Platform))
	}
	if m.Position != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintBannerMng(dAtA, i, uint64(m.Position))
	}
	if len(m.UserPlatform) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintBannerMng(dAtA, i, uint64(len(m.UserPlatform)))
		i += copy(dAtA[i:], m.UserPlatform)
	}
	if len(m.UserDevice) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintBannerMng(dAtA, i, uint64(len(m.UserDevice)))
		i += copy(dAtA[i:], m.UserDevice)
	}
	if m.Build != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintBannerMng(dAtA, i, uint64(m.Build))
	}
	if m.ReturnBuilds != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintBannerMng(dAtA, i, uint64(m.ReturnBuilds))
	}
	return i, nil
}

func (m *BannerMngGetNewBannerResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BannerMngGetNewBannerResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintBannerMng(dAtA, i, uint64(m.Code))
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBannerMng(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	if len(m.Data) > 0 {
		for _, msg := range m.Data {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintBannerMng(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *BannerMngGetNewBannerResp_NewBanner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BannerMngGetNewBannerResp_NewBanner) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBannerMng(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Pic) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBannerMng(dAtA, i, uint64(len(m.Pic)))
		i += copy(dAtA[i:], m.Pic)
	}
	if len(m.Img) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintBannerMng(dAtA, i, uint64(len(m.Img)))
		i += copy(dAtA[i:], m.Img)
	}
	if len(m.Link) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintBannerMng(dAtA, i, uint64(len(m.Link)))
		i += copy(dAtA[i:], m.Link)
	}
	if len(m.Title) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintBannerMng(dAtA, i, uint64(len(m.Title)))
		i += copy(dAtA[i:], m.Title)
	}
	if len(m.Position) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintBannerMng(dAtA, i, uint64(len(m.Position)))
		i += copy(dAtA[i:], m.Position)
	}
	if len(m.SortNum) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintBannerMng(dAtA, i, uint64(len(m.SortNum)))
		i += copy(dAtA[i:], m.SortNum)
	}
	if len(m.Remark) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintBannerMng(dAtA, i, uint64(len(m.Remark)))
		i += copy(dAtA[i:], m.Remark)
	}
	return i, nil
}

func encodeVarintBannerMng(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BannerMngGetNewBannerReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Platform != 0 {
		n += 1 + sovBannerMng(uint64(m.Platform))
	}
	if m.Position != 0 {
		n += 1 + sovBannerMng(uint64(m.Position))
	}
	l = len(m.UserPlatform)
	if l > 0 {
		n += 1 + l + sovBannerMng(uint64(l))
	}
	l = len(m.UserDevice)
	if l > 0 {
		n += 1 + l + sovBannerMng(uint64(l))
	}
	if m.Build != 0 {
		n += 1 + sovBannerMng(uint64(m.Build))
	}
	if m.ReturnBuilds != 0 {
		n += 1 + sovBannerMng(uint64(m.ReturnBuilds))
	}
	return n
}

func (m *BannerMngGetNewBannerResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovBannerMng(uint64(m.Code))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovBannerMng(uint64(l))
	}
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovBannerMng(uint64(l))
		}
	}
	return n
}

func (m *BannerMngGetNewBannerResp_NewBanner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovBannerMng(uint64(l))
	}
	l = len(m.Pic)
	if l > 0 {
		n += 1 + l + sovBannerMng(uint64(l))
	}
	l = len(m.Img)
	if l > 0 {
		n += 1 + l + sovBannerMng(uint64(l))
	}
	l = len(m.Link)
	if l > 0 {
		n += 1 + l + sovBannerMng(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovBannerMng(uint64(l))
	}
	l = len(m.Position)
	if l > 0 {
		n += 1 + l + sovBannerMng(uint64(l))
	}
	l = len(m.SortNum)
	if l > 0 {
		n += 1 + l + sovBannerMng(uint64(l))
	}
	l = len(m.Remark)
	if l > 0 {
		n += 1 + l + sovBannerMng(uint64(l))
	}
	return n
}

func sovBannerMng(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBannerMng(x uint64) (n int) {
	return sovBannerMng(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BannerMngGetNewBannerReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBannerMng
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
			return fmt.Errorf("proto: BannerMngGetNewBannerReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BannerMngGetNewBannerReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Platform", wireType)
			}
			m.Platform = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBannerMng
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Platform |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			m.Position = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBannerMng
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Position |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserPlatform", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBannerMng
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
				return ErrInvalidLengthBannerMng
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserPlatform = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserDevice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBannerMng
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
				return ErrInvalidLengthBannerMng
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserDevice = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Build", wireType)
			}
			m.Build = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBannerMng
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Build |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReturnBuilds", wireType)
			}
			m.ReturnBuilds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBannerMng
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReturnBuilds |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBannerMng(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBannerMng
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
func (m *BannerMngGetNewBannerResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBannerMng
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
			return fmt.Errorf("proto: BannerMngGetNewBannerResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BannerMngGetNewBannerResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBannerMng
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBannerMng
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
				return ErrInvalidLengthBannerMng
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBannerMng
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBannerMng
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &BannerMngGetNewBannerResp_NewBanner{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBannerMng(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBannerMng
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
func (m *BannerMngGetNewBannerResp_NewBanner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBannerMng
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
			return fmt.Errorf("proto: NewBanner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NewBanner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBannerMng
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
				return ErrInvalidLengthBannerMng
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBannerMng
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
				return ErrInvalidLengthBannerMng
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Img", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBannerMng
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
				return ErrInvalidLengthBannerMng
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Img = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Link", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBannerMng
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
				return ErrInvalidLengthBannerMng
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Link = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBannerMng
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
				return ErrInvalidLengthBannerMng
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBannerMng
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
				return ErrInvalidLengthBannerMng
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Position = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SortNum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBannerMng
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
				return ErrInvalidLengthBannerMng
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SortNum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Remark", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBannerMng
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
				return ErrInvalidLengthBannerMng
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Remark = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBannerMng(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBannerMng
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
func skipBannerMng(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBannerMng
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
					return 0, ErrIntOverflowBannerMng
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
					return 0, ErrIntOverflowBannerMng
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
				return 0, ErrInvalidLengthBannerMng
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBannerMng
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
				next, err := skipBannerMng(dAtA[start:])
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
	ErrInvalidLengthBannerMng = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBannerMng   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("v1/BannerMng.proto", fileDescriptor_BannerMng_9505b5c7b0530020) }

var fileDescriptor_BannerMng_9505b5c7b0530020 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x6e, 0x9b, 0x40,
	0x10, 0x35, 0xe0, 0xd8, 0x66, 0x6a, 0x55, 0xd5, 0x1e, 0x2a, 0x62, 0x45, 0x60, 0x59, 0xad, 0xea,
	0x4b, 0xb1, 0x9c, 0xf6, 0x0b, 0x50, 0xa5, 0x9e, 0x12, 0x55, 0x7b, 0x6c, 0x0f, 0x91, 0x0d, 0x1b,
	0xb2, 0x8a, 0x61, 0xc9, 0xb2, 0xb8, 0xfd, 0x8b, 0xf6, 0x3f, 0x7a, 0xee, 0x3f, 0xf4, 0x98, 0x63,
	0x4f, 0xa8, 0xb2, 0x6f, 0x7c, 0x45, 0xb4, 0x03, 0x21, 0xf8, 0x60, 0x29, 0x97, 0xd9, 0x79, 0x6f,
	0xdf, 0xee, 0x0c, 0x6f, 0x16, 0x20, 0xdb, 0xe5, 0x22, 0x58, 0xa5, 0x29, 0x93, 0x17, 0x69, 0xec,
	0x67, 0x52, 0x28, 0x41, 0x40, 0x0a, 0x91, 0x5c, 0xb1, 0x1f, 0xfe, 0x76, 0x39, 0x79, 0x1f, 0x73,
	0x75, 0x53, 0xac, 0xfd, 0x50, 0x24, 0x8b, 0x58, 0xc4, 0x62, 0x81, 0x92, 0x75, 0x71, 0x8d, 0x08,
	0x01, 0x66, 0xf5, 0xd1, 0xd9, 0x6f, 0x13, 0x9c, 0xf6, 0xba, 0xcf, 0x4c, 0x5d, 0xb2, 0xef, 0x35,
	0xa4, 0xec, 0x8e, 0xcc, 0x61, 0x94, 0x6d, 0x56, 0xea, 0x5a, 0xc8, 0xc4, 0x31, 0xa6, 0xc6, 0xdc,
	0x0a, 0xc6, 0x55, 0xe9, 0xb5, 0x1c, 0x6d, 0x33, 0x54, 0x8a, 0x9c, 0x2b, 0x2e, 0x52, 0xc7, 0xec,
	0x28, 0x1b, 0x8e, 0xb6, 0x19, 0xf9, 0x08, 0xe3, 0x22, 0x67, 0xf2, 0xcb, 0xe3, 0xbd, 0xd6, 0xd4,
	0x98, 0xdb, 0xc1, 0xab, 0xaa, 0xf4, 0x0e, 0x78, 0x7a, 0x80, 0x88, 0x0f, 0xa0, 0xf1, 0x27, 0xb6,
	0xe5, 0x21, 0x73, 0xfa, 0x78, 0xe6, 0x65, 0x55, 0x7a, 0x1d, 0x96, 0x76, 0x72, 0xe2, 0xc1, 0xc9,
	0xba, 0xe0, 0x9b, 0xc8, 0x39, 0xc1, 0x66, 0xec, 0xaa, 0xf4, 0x6a, 0x82, 0xd6, 0x8b, 0x6e, 0x43,
	0x32, 0x55, 0xc8, 0x34, 0xd0, 0x30, 0x77, 0x06, 0xa8, 0xc3, 0x36, 0xba, 0x3c, 0x3d, 0x40, 0xb3,
	0x3f, 0x16, 0x9c, 0x1e, 0x71, 0x2b, 0xcf, 0xc8, 0x19, 0xf4, 0x43, 0x11, 0xb1, 0xc6, 0xaa, 0x51,
	0x55, 0x7a, 0x88, 0x29, 0x46, 0x72, 0x0a, 0x56, 0x92, 0xc7, 0xe8, 0x8e, 0x1d, 0x0c, 0xab, 0xd2,
	0xd3, 0x90, 0xea, 0x40, 0x2e, 0xa0, 0x1f, 0xad, 0xd4, 0xca, 0xb1, 0xa6, 0xd6, 0xfc, 0xc5, 0xf9,
	0xc2, 0x7f, 0x1a, 0xa7, 0x7f, 0xb4, 0x9a, 0xdf, 0xa2, 0xba, 0x92, 0xbe, 0x80, 0x62, 0x9c, 0xfc,
	0x34, 0xc1, 0x6e, 0x77, 0xc9, 0x6b, 0x30, 0x79, 0x84, 0x3d, 0xd9, 0xc1, 0xa0, 0x2a, 0x3d, 0x93,
	0x47, 0xd4, 0xe4, 0x91, 0xee, 0x27, 0xe3, 0x61, 0xb7, 0x9f, 0x8c, 0x87, 0x54, 0x07, 0xbd, 0xc5,
	0x93, 0xb8, 0x19, 0x0d, 0x6e, 0xf1, 0x24, 0xa6, 0x3a, 0xe8, 0x6f, 0xdc, 0xf0, 0xf4, 0xb6, 0x19,
	0x01, 0x56, 0xd6, 0x98, 0x62, 0xd4, 0xb6, 0x2b, 0xae, 0x36, 0x0c, 0x6d, 0xb7, 0x6b, 0xdb, 0x91,
	0xa0, 0xf5, 0x72, 0xf0, 0x4e, 0x06, 0xa8, 0x39, 0xf6, 0x4e, 0xde, 0xc1, 0x28, 0x17, 0x52, 0x5d,
	0xa5, 0x45, 0xe2, 0x0c, 0x9f, 0x94, 0x8f, 0x1c, 0x1d, 0xea, 0xec, 0xb2, 0x48, 0xc8, 0x0c, 0x06,
	0x92, 0x25, 0x2b, 0x79, 0xeb, 0x8c, 0x50, 0x06, 0x55, 0xe9, 0x35, 0x0c, 0x6d, 0xd6, 0xf3, 0x1b,
	0xb0, 0x5b, 0x23, 0xc9, 0x37, 0x18, 0xc7, 0x1d, 0x33, 0xc9, 0x9b, 0x67, 0xf8, 0x7d, 0x37, 0x79,
	0xfb, 0xac, 0xa9, 0x04, 0x67, 0x7f, 0x77, 0xae, 0x71, 0xbf, 0x73, 0x8d, 0xff, 0x3b, 0xd7, 0xf8,
	0xb5, 0x77, 0x7b, 0xf7, 0x7b, 0xb7, 0xf7, 0x6f, 0xef, 0xf6, 0xbe, 0x9a, 0xdb, 0xe5, 0x7a, 0x80,
	0x3f, 0xdd, 0x87, 0x87, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0xa9, 0xe6, 0x47, 0xc5, 0x03, 0x00,
	0x00,
}
