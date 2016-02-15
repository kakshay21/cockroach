// Code generated by protoc-gen-gogo.
// source: cockroach/storage/raft.proto
// DO NOT EDIT!

/*
	Package storage is a generated protocol buffer package.

	It is generated from these files:
		cockroach/storage/raft.proto
		cockroach/storage/status.proto

	It has these top-level messages:
		RaftMessageRequest
		RaftMessageResponse
		ConfChangeContext
		StoreStatus
*/
package storage

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_roachpb "github.com/cockroachdb/cockroach/roachpb"
import raftpb "github.com/coreos/etcd/raft/raftpb"

// skipping weak import gogoproto "github.com/cockroachdb/gogoproto"

import github_com_cockroachdb_cockroach_roachpb "github.com/cockroachdb/cockroach/roachpb"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// RaftMessageRequest is the request used to send raft messages using our
// protobuf-based RPC codec.
type RaftMessageRequest struct {
	GroupID     github_com_cockroachdb_cockroach_roachpb.RangeID `protobuf:"varint,1,opt,name=group_id,casttype=github.com/cockroachdb/cockroach/roachpb.RangeID" json:"group_id"`
	FromReplica cockroach_roachpb.ReplicaDescriptor              `protobuf:"bytes,2,opt,name=from_replica" json:"from_replica"`
	ToReplica   cockroach_roachpb.ReplicaDescriptor              `protobuf:"bytes,3,opt,name=to_replica" json:"to_replica"`
	Message     raftpb.Message                                   `protobuf:"bytes,4,opt,name=message" json:"message"`
}

func (m *RaftMessageRequest) Reset()         { *m = RaftMessageRequest{} }
func (m *RaftMessageRequest) String() string { return proto.CompactTextString(m) }
func (*RaftMessageRequest) ProtoMessage()    {}

// RaftMessageResponse is an empty message returned by raft RPCs. If a
// response is needed it will be sent as a separate message.
type RaftMessageResponse struct {
}

func (m *RaftMessageResponse) Reset()         { *m = RaftMessageResponse{} }
func (m *RaftMessageResponse) String() string { return proto.CompactTextString(m) }
func (*RaftMessageResponse) ProtoMessage()    {}

// ConfChangeContext is encoded in the raftpb.ConfChange.Context field.
type ConfChangeContext struct {
	CommandID string `protobuf:"bytes,1,opt,name=command_id" json:"command_id"`
	// Payload is the application-level command (i.e. an encoded
	// roachpb.EndTransactionRequest).
	Payload []byte `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
	// Replica contains full details about the replica being added or removed.
	Replica cockroach_roachpb.ReplicaDescriptor `protobuf:"bytes,3,opt,name=replica" json:"replica"`
}

func (m *ConfChangeContext) Reset()         { *m = ConfChangeContext{} }
func (m *ConfChangeContext) String() string { return proto.CompactTextString(m) }
func (*ConfChangeContext) ProtoMessage()    {}

func init() {
	proto.RegisterType((*RaftMessageRequest)(nil), "cockroach.storage.RaftMessageRequest")
	proto.RegisterType((*RaftMessageResponse)(nil), "cockroach.storage.RaftMessageResponse")
	proto.RegisterType((*ConfChangeContext)(nil), "cockroach.storage.ConfChangeContext")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for MultiRaft service

type MultiRaftClient interface {
	RaftMessage(ctx context.Context, opts ...grpc.CallOption) (MultiRaft_RaftMessageClient, error)
}

type multiRaftClient struct {
	cc *grpc.ClientConn
}

func NewMultiRaftClient(cc *grpc.ClientConn) MultiRaftClient {
	return &multiRaftClient{cc}
}

func (c *multiRaftClient) RaftMessage(ctx context.Context, opts ...grpc.CallOption) (MultiRaft_RaftMessageClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MultiRaft_serviceDesc.Streams[0], c.cc, "/cockroach.storage.MultiRaft/RaftMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &multiRaftRaftMessageClient{stream}
	return x, nil
}

type MultiRaft_RaftMessageClient interface {
	Send(*RaftMessageRequest) error
	CloseAndRecv() (*RaftMessageResponse, error)
	grpc.ClientStream
}

type multiRaftRaftMessageClient struct {
	grpc.ClientStream
}

func (x *multiRaftRaftMessageClient) Send(m *RaftMessageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *multiRaftRaftMessageClient) CloseAndRecv() (*RaftMessageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RaftMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MultiRaft service

type MultiRaftServer interface {
	RaftMessage(MultiRaft_RaftMessageServer) error
}

func RegisterMultiRaftServer(s *grpc.Server, srv MultiRaftServer) {
	s.RegisterService(&_MultiRaft_serviceDesc, srv)
}

func _MultiRaft_RaftMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MultiRaftServer).RaftMessage(&multiRaftRaftMessageServer{stream})
}

type MultiRaft_RaftMessageServer interface {
	SendAndClose(*RaftMessageResponse) error
	Recv() (*RaftMessageRequest, error)
	grpc.ServerStream
}

type multiRaftRaftMessageServer struct {
	grpc.ServerStream
}

func (x *multiRaftRaftMessageServer) SendAndClose(m *RaftMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *multiRaftRaftMessageServer) Recv() (*RaftMessageRequest, error) {
	m := new(RaftMessageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MultiRaft_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cockroach.storage.MultiRaft",
	HandlerType: (*MultiRaftServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RaftMessage",
			Handler:       _MultiRaft_RaftMessage_Handler,
			ClientStreams: true,
		},
	},
}

func (m *RaftMessageRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RaftMessageRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintRaft(data, i, uint64(m.GroupID))
	data[i] = 0x12
	i++
	i = encodeVarintRaft(data, i, uint64(m.FromReplica.Size()))
	n1, err := m.FromReplica.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x1a
	i++
	i = encodeVarintRaft(data, i, uint64(m.ToReplica.Size()))
	n2, err := m.ToReplica.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	data[i] = 0x22
	i++
	i = encodeVarintRaft(data, i, uint64(m.Message.Size()))
	n3, err := m.Message.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *RaftMessageResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RaftMessageResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *ConfChangeContext) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ConfChangeContext) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintRaft(data, i, uint64(len(m.CommandID)))
	i += copy(data[i:], m.CommandID)
	if m.Payload != nil {
		data[i] = 0x12
		i++
		i = encodeVarintRaft(data, i, uint64(len(m.Payload)))
		i += copy(data[i:], m.Payload)
	}
	data[i] = 0x1a
	i++
	i = encodeVarintRaft(data, i, uint64(m.Replica.Size()))
	n4, err := m.Replica.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func encodeFixed64Raft(data []byte, offset int, v uint64) int {
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
func encodeFixed32Raft(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRaft(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *RaftMessageRequest) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovRaft(uint64(m.GroupID))
	l = m.FromReplica.Size()
	n += 1 + l + sovRaft(uint64(l))
	l = m.ToReplica.Size()
	n += 1 + l + sovRaft(uint64(l))
	l = m.Message.Size()
	n += 1 + l + sovRaft(uint64(l))
	return n
}

func (m *RaftMessageResponse) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *ConfChangeContext) Size() (n int) {
	var l int
	_ = l
	l = len(m.CommandID)
	n += 1 + l + sovRaft(uint64(l))
	if m.Payload != nil {
		l = len(m.Payload)
		n += 1 + l + sovRaft(uint64(l))
	}
	l = m.Replica.Size()
	n += 1 + l + sovRaft(uint64(l))
	return n
}

func sovRaft(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRaft(x uint64) (n int) {
	return sovRaft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RaftMessageRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
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
			return fmt.Errorf("proto: RaftMessageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftMessageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupID", wireType)
			}
			m.GroupID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.GroupID |= (github_com_cockroachdb_cockroach_roachpb.RangeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromReplica", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
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
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FromReplica.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToReplica", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
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
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ToReplica.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
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
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Message.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
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
func (m *RaftMessageResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
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
			return fmt.Errorf("proto: RaftMessageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftMessageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
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
func (m *ConfChangeContext) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
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
			return fmt.Errorf("proto: ConfChangeContext: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfChangeContext: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommandID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
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
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommandID = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
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
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], data[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replica", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
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
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Replica.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
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
func skipRaft(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRaft
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
					return 0, ErrIntOverflowRaft
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
					return 0, ErrIntOverflowRaft
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
				return 0, ErrInvalidLengthRaft
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRaft
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
				next, err := skipRaft(data[start:])
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
	ErrInvalidLengthRaft = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRaft   = fmt.Errorf("proto: integer overflow")
)
