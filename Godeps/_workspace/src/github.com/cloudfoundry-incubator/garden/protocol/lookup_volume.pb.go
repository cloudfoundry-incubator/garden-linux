// Code generated by protoc-gen-gogo.
// source: lookup_volume.proto
// DO NOT EDIT!

package garden

import proto "code.google.com/p/gogoprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type LookupVolumeRequest struct {
	Handle           *string `protobuf:"bytes,1,req,name=handle" json:"handle,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LookupVolumeRequest) Reset()         { *m = LookupVolumeRequest{} }
func (m *LookupVolumeRequest) String() string { return proto.CompactTextString(m) }
func (*LookupVolumeRequest) ProtoMessage()    {}

func (m *LookupVolumeRequest) GetHandle() string {
	if m != nil && m.Handle != nil {
		return *m.Handle
	}
	return ""
}

type LookupVolumeResponse struct {
	Handle           *string `protobuf:"bytes,1,req,name=handle" json:"handle,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LookupVolumeResponse) Reset()         { *m = LookupVolumeResponse{} }
func (m *LookupVolumeResponse) String() string { return proto.CompactTextString(m) }
func (*LookupVolumeResponse) ProtoMessage()    {}

func (m *LookupVolumeResponse) GetHandle() string {
	if m != nil && m.Handle != nil {
		return *m.Handle
	}
	return ""
}

func init() {
}
