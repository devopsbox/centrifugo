// Code generated by protoc-gen-gogo.
// source: message.proto
// DO NOT EDIT!

/*
	Package libcentrifugo is a generated protocol buffer package.

	It is generated from these files:
		message.proto

	It has these top-level messages:
		ClientInfo
		Message
		JoinMessage
		LeaveMessage
*/
package libcentrifugo

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import github_com_centrifugal_centrifugo_libcentrifugo_raw "github.com/centrifugal/centrifugo/libcentrifugo/raw"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

type ClientInfo struct {
	User        string                                                   `protobuf:"bytes,1,opt,name=User" json:"user"`
	Client      string                                                   `protobuf:"bytes,2,opt,name=Client" json:"client"`
	DefaultInfo *github_com_centrifugal_centrifugo_libcentrifugo_raw.Raw `protobuf:"bytes,3,opt,name=DefaultInfo,customtype=github.com/centrifugal/centrifugo/libcentrifugo/raw.Raw" json:"default_info,omitempty"`
	ChannelInfo *github_com_centrifugal_centrifugo_libcentrifugo_raw.Raw `protobuf:"bytes,4,opt,name=ChannelInfo,customtype=github.com/centrifugal/centrifugo/libcentrifugo/raw.Raw" json:"channel_info,omitempty"`
}

func (m *ClientInfo) Reset()                    { *m = ClientInfo{} }
func (m *ClientInfo) String() string            { return proto.CompactTextString(m) }
func (*ClientInfo) ProtoMessage()               {}
func (*ClientInfo) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{0} }

func (m *ClientInfo) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ClientInfo) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

type Message struct {
	UID       string                                                   `protobuf:"bytes,1,opt,name=UID" json:"uid"`
	Timestamp string                                                   `protobuf:"bytes,2,opt,name=Timestamp" json:"timestamp"`
	Channel   string                                                   `protobuf:"bytes,3,opt,name=Channel" json:"channel"`
	Data      *github_com_centrifugal_centrifugo_libcentrifugo_raw.Raw `protobuf:"bytes,4,opt,name=Data,customtype=github.com/centrifugal/centrifugo/libcentrifugo/raw.Raw" json:"data"`
	Client    string                                                   `protobuf:"bytes,5,opt,name=Client" json:"client,omitempty"`
	Info      *ClientInfo                                              `protobuf:"bytes,6,opt,name=Info" json:"info,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{1} }

func (m *Message) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *Message) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Message) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *Message) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

func (m *Message) GetInfo() *ClientInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type JoinMessage struct {
	Channel string     `protobuf:"bytes,1,opt,name=Channel" json:"channel"`
	Data    ClientInfo `protobuf:"bytes,2,opt,name=Data" json:"data"`
}

func (m *JoinMessage) Reset()                    { *m = JoinMessage{} }
func (m *JoinMessage) String() string            { return proto.CompactTextString(m) }
func (*JoinMessage) ProtoMessage()               {}
func (*JoinMessage) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{2} }

func (m *JoinMessage) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *JoinMessage) GetData() ClientInfo {
	if m != nil {
		return m.Data
	}
	return ClientInfo{}
}

type LeaveMessage struct {
	Channel string     `protobuf:"bytes,1,opt,name=Channel" json:"channel"`
	Data    ClientInfo `protobuf:"bytes,2,opt,name=Data" json:"data"`
}

func (m *LeaveMessage) Reset()                    { *m = LeaveMessage{} }
func (m *LeaveMessage) String() string            { return proto.CompactTextString(m) }
func (*LeaveMessage) ProtoMessage()               {}
func (*LeaveMessage) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{3} }

func (m *LeaveMessage) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *LeaveMessage) GetData() ClientInfo {
	if m != nil {
		return m.Data
	}
	return ClientInfo{}
}

func init() {
	proto.RegisterType((*ClientInfo)(nil), "libcentrifugo.ClientInfo")
	proto.RegisterType((*Message)(nil), "libcentrifugo.Message")
	proto.RegisterType((*JoinMessage)(nil), "libcentrifugo.JoinMessage")
	proto.RegisterType((*LeaveMessage)(nil), "libcentrifugo.LeaveMessage")
}
func (this *ClientInfo) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ClientInfo)
	if !ok {
		that2, ok := that.(ClientInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.User != that1.User {
		return false
	}
	if this.Client != that1.Client {
		return false
	}
	if that1.DefaultInfo == nil {
		if this.DefaultInfo != nil {
			return false
		}
	} else if !this.DefaultInfo.Equal(*that1.DefaultInfo) {
		return false
	}
	if that1.ChannelInfo == nil {
		if this.ChannelInfo != nil {
			return false
		}
	} else if !this.ChannelInfo.Equal(*that1.ChannelInfo) {
		return false
	}
	return true
}
func (this *Message) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Message)
	if !ok {
		that2, ok := that.(Message)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.UID != that1.UID {
		return false
	}
	if this.Timestamp != that1.Timestamp {
		return false
	}
	if this.Channel != that1.Channel {
		return false
	}
	if that1.Data == nil {
		if this.Data != nil {
			return false
		}
	} else if !this.Data.Equal(*that1.Data) {
		return false
	}
	if this.Client != that1.Client {
		return false
	}
	if !this.Info.Equal(that1.Info) {
		return false
	}
	return true
}
func (this *JoinMessage) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*JoinMessage)
	if !ok {
		that2, ok := that.(JoinMessage)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Channel != that1.Channel {
		return false
	}
	if !this.Data.Equal(&that1.Data) {
		return false
	}
	return true
}
func (this *LeaveMessage) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*LeaveMessage)
	if !ok {
		that2, ok := that.(LeaveMessage)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Channel != that1.Channel {
		return false
	}
	if !this.Data.Equal(&that1.Data) {
		return false
	}
	return true
}
func (m *ClientInfo) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ClientInfo) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintMessage(data, i, uint64(len(m.User)))
	i += copy(data[i:], m.User)
	data[i] = 0x12
	i++
	i = encodeVarintMessage(data, i, uint64(len(m.Client)))
	i += copy(data[i:], m.Client)
	if m.DefaultInfo != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintMessage(data, i, uint64(m.DefaultInfo.Size()))
		n1, err := m.DefaultInfo.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.ChannelInfo != nil {
		data[i] = 0x22
		i++
		i = encodeVarintMessage(data, i, uint64(m.ChannelInfo.Size()))
		n2, err := m.ChannelInfo.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *Message) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Message) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintMessage(data, i, uint64(len(m.UID)))
	i += copy(data[i:], m.UID)
	data[i] = 0x12
	i++
	i = encodeVarintMessage(data, i, uint64(len(m.Timestamp)))
	i += copy(data[i:], m.Timestamp)
	data[i] = 0x1a
	i++
	i = encodeVarintMessage(data, i, uint64(len(m.Channel)))
	i += copy(data[i:], m.Channel)
	if m.Data != nil {
		data[i] = 0x22
		i++
		i = encodeVarintMessage(data, i, uint64(m.Data.Size()))
		n3, err := m.Data.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	data[i] = 0x2a
	i++
	i = encodeVarintMessage(data, i, uint64(len(m.Client)))
	i += copy(data[i:], m.Client)
	if m.Info != nil {
		data[i] = 0x32
		i++
		i = encodeVarintMessage(data, i, uint64(m.Info.Size()))
		n4, err := m.Info.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *JoinMessage) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *JoinMessage) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintMessage(data, i, uint64(len(m.Channel)))
	i += copy(data[i:], m.Channel)
	data[i] = 0x12
	i++
	i = encodeVarintMessage(data, i, uint64(m.Data.Size()))
	n5, err := m.Data.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	return i, nil
}

func (m *LeaveMessage) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *LeaveMessage) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintMessage(data, i, uint64(len(m.Channel)))
	i += copy(data[i:], m.Channel)
	data[i] = 0x12
	i++
	i = encodeVarintMessage(data, i, uint64(m.Data.Size()))
	n6, err := m.Data.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	return i, nil
}

func encodeFixed64Message(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Message(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMessage(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedClientInfo(r randyMessage, easy bool) *ClientInfo {
	this := &ClientInfo{}
	this.User = randStringMessage(r)
	this.Client = randStringMessage(r)
	if r.Intn(10) != 0 {
		this.DefaultInfo = github_com_centrifugal_centrifugo_libcentrifugo_raw.NewPopulatedRaw(r)
	}
	if r.Intn(10) != 0 {
		this.ChannelInfo = github_com_centrifugal_centrifugo_libcentrifugo_raw.NewPopulatedRaw(r)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedMessage(r randyMessage, easy bool) *Message {
	this := &Message{}
	this.UID = randStringMessage(r)
	this.Timestamp = randStringMessage(r)
	this.Channel = randStringMessage(r)
	if r.Intn(10) != 0 {
		this.Data = github_com_centrifugal_centrifugo_libcentrifugo_raw.NewPopulatedRaw(r)
	}
	this.Client = randStringMessage(r)
	if r.Intn(10) != 0 {
		this.Info = NewPopulatedClientInfo(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedJoinMessage(r randyMessage, easy bool) *JoinMessage {
	this := &JoinMessage{}
	this.Channel = randStringMessage(r)
	v1 := NewPopulatedClientInfo(r, easy)
	this.Data = *v1
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedLeaveMessage(r randyMessage, easy bool) *LeaveMessage {
	this := &LeaveMessage{}
	this.Channel = randStringMessage(r)
	v2 := NewPopulatedClientInfo(r, easy)
	this.Data = *v2
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyMessage interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMessage(r randyMessage) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringMessage(r randyMessage) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneMessage(r)
	}
	return string(tmps)
}
func randUnrecognizedMessage(r randyMessage, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldMessage(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldMessage(data []byte, r randyMessage, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateMessage(data, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		data = encodeVarintPopulateMessage(data, uint64(v4))
	case 1:
		data = encodeVarintPopulateMessage(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateMessage(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateMessage(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateMessage(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateMessage(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *ClientInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.User)
	n += 1 + l + sovMessage(uint64(l))
	l = len(m.Client)
	n += 1 + l + sovMessage(uint64(l))
	if m.DefaultInfo != nil {
		l = m.DefaultInfo.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.ChannelInfo != nil {
		l = m.ChannelInfo.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func (m *Message) Size() (n int) {
	var l int
	_ = l
	l = len(m.UID)
	n += 1 + l + sovMessage(uint64(l))
	l = len(m.Timestamp)
	n += 1 + l + sovMessage(uint64(l))
	l = len(m.Channel)
	n += 1 + l + sovMessage(uint64(l))
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Client)
	n += 1 + l + sovMessage(uint64(l))
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func (m *JoinMessage) Size() (n int) {
	var l int
	_ = l
	l = len(m.Channel)
	n += 1 + l + sovMessage(uint64(l))
	l = m.Data.Size()
	n += 1 + l + sovMessage(uint64(l))
	return n
}

func (m *LeaveMessage) Size() (n int) {
	var l int
	_ = l
	l = len(m.Channel)
	n += 1 + l + sovMessage(uint64(l))
	l = m.Data.Size()
	n += 1 + l + sovMessage(uint64(l))
	return n
}

func sovMessage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientInfo) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClientInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Client", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Client = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultInfo", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_centrifugal_centrifugo_libcentrifugo_raw.Raw
			m.DefaultInfo = &v
			if err := m.DefaultInfo.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelInfo", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_centrifugal_centrifugo_libcentrifugo_raw.Raw
			m.ChannelInfo = &v
			if err := m.ChannelInfo.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *Message) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timestamp = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_centrifugal_centrifugo_libcentrifugo_raw.Raw
			m.Data = &v
			if err := m.Data.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Client", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Client = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Info == nil {
				m.Info = &ClientInfo{}
			}
			if err := m.Info.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *JoinMessage) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: JoinMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Data.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *LeaveMessage) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LeaveMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaveMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Data.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMessage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMessage
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipMessage(data[start:])
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
	ErrInvalidLengthMessage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorMessage = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcd, 0xc9, 0x4c, 0x4a, 0x4e,
	0xcd, 0x2b, 0x29, 0xca, 0x4c, 0x2b, 0x4d, 0xcf, 0x97, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d,
	0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xab, 0x4a, 0x2a, 0x4d, 0x03,
	0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2, 0x5b, 0x69, 0x05, 0x13, 0x17, 0x97, 0x73, 0x4e, 0x26, 0x50,
	0xbf, 0x67, 0x5e, 0x5a, 0xbe, 0x90, 0x14, 0x17, 0x4b, 0x68, 0x71, 0x6a, 0x91, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0xa7, 0x13, 0xcf, 0x89, 0x7b, 0xf2, 0x0c, 0xaf, 0xee, 0xc9, 0xb3, 0x94, 0x02, 0xc5,
	0x84, 0xe4, 0xb8, 0xd8, 0x20, 0x2a, 0x25, 0x98, 0xc0, 0xb2, 0x7c, 0x50, 0x59, 0xb6, 0x64, 0xb0,
	0xa8, 0x50, 0x1e, 0x17, 0xb7, 0x4b, 0x6a, 0x5a, 0x62, 0x69, 0x0e, 0xd8, 0x28, 0x09, 0x66, 0xa0,
	0x22, 0x1e, 0xa7, 0x48, 0xa0, 0x22, 0xc6, 0x5b, 0xf7, 0xe4, 0xcd, 0x91, 0x9c, 0x05, 0x77, 0x6d,
	0x62, 0x0e, 0x82, 0x9d, 0xaf, 0x8f, 0xe2, 0x0f, 0xfd, 0xa2, 0xc4, 0x72, 0xbd, 0xa0, 0xc4, 0x72,
	0xa0, 0xf9, 0x62, 0x29, 0x10, 0x53, 0xe3, 0x33, 0x81, 0xc6, 0xea, 0xe4, 0xe7, 0x66, 0x96, 0xa4,
	0xe6, 0x16, 0x94, 0x54, 0x82, 0xec, 0x73, 0xce, 0x48, 0xcc, 0xcb, 0x4b, 0xcd, 0x01, 0xdb, 0xc7,
	0x42, 0x35, 0xfb, 0x92, 0x21, 0xa6, 0xa2, 0xd9, 0xa7, 0xb4, 0x92, 0x89, 0x8b, 0xdd, 0x17, 0x12,
	0xf4, 0x42, 0x12, 0x5c, 0xcc, 0xa1, 0x9e, 0x2e, 0xd0, 0x60, 0xe2, 0x86, 0x06, 0x04, 0x73, 0x69,
	0x66, 0x8a, 0x90, 0x0a, 0x17, 0x67, 0x48, 0x26, 0x30, 0x86, 0x4a, 0x12, 0x73, 0x0b, 0xa0, 0x01,
	0x25, 0x08, 0x95, 0xe7, 0x2c, 0x81, 0x49, 0x08, 0x29, 0x70, 0xb1, 0x43, 0xdd, 0x0e, 0x0e, 0x27,
	0x4e, 0x27, 0x7e, 0xa8, 0x1a, 0x76, 0xa8, 0xe5, 0x42, 0xa1, 0x5c, 0x2c, 0x2e, 0x89, 0x25, 0x89,
	0x50, 0x6f, 0xb9, 0x53, 0xee, 0x2d, 0x96, 0x14, 0xa0, 0x71, 0x42, 0x1a, 0xf0, 0x48, 0x64, 0x05,
	0xdb, 0x2b, 0x01, 0xb5, 0x57, 0x00, 0x12, 0x89, 0x48, 0xc1, 0x6b, 0xcd, 0xc5, 0x02, 0x0e, 0x57,
	0x36, 0xa0, 0x3a, 0x6e, 0x23, 0x49, 0x3d, 0x14, 0x73, 0xf5, 0x10, 0x69, 0xc6, 0x49, 0x08, 0xa8,
	0x9d, 0x0f, 0x2d, 0xac, 0xd2, 0xb8, 0xb8, 0xbd, 0xf2, 0x33, 0xf3, 0x60, 0xc1, 0x85, 0xe4, 0x5d,
	0x46, 0xec, 0xde, 0x35, 0x85, 0x7a, 0x97, 0x89, 0x90, 0x6d, 0xf0, 0x34, 0x09, 0xf2, 0x8e, 0x52,
	0x3a, 0x17, 0x8f, 0x4f, 0x6a, 0x62, 0x59, 0x2a, 0xad, 0x2d, 0x72, 0x92, 0xf9, 0xf1, 0x50, 0x8e,
	0x71, 0xc5, 0x23, 0x39, 0xc6, 0x1d, 0x40, 0x7c, 0x02, 0x88, 0x2f, 0x00, 0xf1, 0x03, 0x20, 0x9e,
	0xf0, 0x58, 0x8e, 0x01, 0x10, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x68, 0x36, 0x6d, 0x93, 0x03, 0x00,
	0x00,
}
