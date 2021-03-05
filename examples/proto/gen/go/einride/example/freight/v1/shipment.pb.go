// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: einride/example/freight/v1/shipment.proto

package examplefreightv1

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A shipment represents transportation of goods between an origin
// [site][einride.example.freight.v1.Site] and a destination
// [site][einride.example.freight.v1.Site].
type Shipment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the shipment.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The creation timestamp of the shipment.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last update timestamp of the shipment.
	//
	// Updated when create/update/delete operation is shipment.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The deletion timestamp of the shipment.
	DeleteTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// The resource name of the origin site of the shipment.
	// Format: shippers/{shipper}/sites/{site}
	OriginSite string `protobuf:"bytes,5,opt,name=origin_site,json=originSite,proto3" json:"origin_site,omitempty"`
	// The resource name of the destination site of the shipment.
	// Format: shippers/{shipper}/sites/{site}
	DestinationSite string `protobuf:"bytes,6,opt,name=destination_site,json=destinationSite,proto3" json:"destination_site,omitempty"`
	// The earliest pickup time of the shipment at the origin site.
	PickupEarliestTime *timestamp.Timestamp `protobuf:"bytes,7,opt,name=pickup_earliest_time,json=pickupEarliestTime,proto3" json:"pickup_earliest_time,omitempty"`
	// The latest pickup time of the shipment at the origin site.
	PickupLatestTime *timestamp.Timestamp `protobuf:"bytes,8,opt,name=pickup_latest_time,json=pickupLatestTime,proto3" json:"pickup_latest_time,omitempty"`
	// The earliest delivery time of the shipment at the destination site.
	DeliveryEarliestTime *timestamp.Timestamp `protobuf:"bytes,9,opt,name=delivery_earliest_time,json=deliveryEarliestTime,proto3" json:"delivery_earliest_time,omitempty"`
	// The latest delivery time of the shipment at the destination site.
	DeliveryLatestTime *timestamp.Timestamp `protobuf:"bytes,10,opt,name=delivery_latest_time,json=deliveryLatestTime,proto3" json:"delivery_latest_time,omitempty"`
	// The line items of the shipment.
	LineItems []*LineItem `protobuf:"bytes,11,rep,name=line_items,json=lineItems,proto3" json:"line_items,omitempty"`
	// Annotations of the shipment.
	Annotations map[string]string `protobuf:"bytes,12,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Shipment) Reset() {
	*x = Shipment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_example_freight_v1_shipment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shipment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shipment) ProtoMessage() {}

func (x *Shipment) ProtoReflect() protoreflect.Message {
	mi := &file_einride_example_freight_v1_shipment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shipment.ProtoReflect.Descriptor instead.
func (*Shipment) Descriptor() ([]byte, []int) {
	return file_einride_example_freight_v1_shipment_proto_rawDescGZIP(), []int{0}
}

func (x *Shipment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Shipment) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Shipment) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Shipment) GetDeleteTime() *timestamp.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

func (x *Shipment) GetOriginSite() string {
	if x != nil {
		return x.OriginSite
	}
	return ""
}

func (x *Shipment) GetDestinationSite() string {
	if x != nil {
		return x.DestinationSite
	}
	return ""
}

func (x *Shipment) GetPickupEarliestTime() *timestamp.Timestamp {
	if x != nil {
		return x.PickupEarliestTime
	}
	return nil
}

func (x *Shipment) GetPickupLatestTime() *timestamp.Timestamp {
	if x != nil {
		return x.PickupLatestTime
	}
	return nil
}

func (x *Shipment) GetDeliveryEarliestTime() *timestamp.Timestamp {
	if x != nil {
		return x.DeliveryEarliestTime
	}
	return nil
}

func (x *Shipment) GetDeliveryLatestTime() *timestamp.Timestamp {
	if x != nil {
		return x.DeliveryLatestTime
	}
	return nil
}

func (x *Shipment) GetLineItems() []*LineItem {
	if x != nil {
		return x.LineItems
	}
	return nil
}

func (x *Shipment) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

// A shipment line item.
type LineItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The title of the line item.
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// The quantity of the line item.
	Quantity float32 `protobuf:"fixed32,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// The weight of the line item in kilograms.
	WeightKg float32 `protobuf:"fixed32,3,opt,name=weight_kg,json=weightKg,proto3" json:"weight_kg,omitempty"`
	// The volume of the line item in cubic meters.
	VolumeM3 float32 `protobuf:"fixed32,4,opt,name=volume_m3,json=volumeM3,proto3" json:"volume_m3,omitempty"`
}

func (x *LineItem) Reset() {
	*x = LineItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_example_freight_v1_shipment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineItem) ProtoMessage() {}

func (x *LineItem) ProtoReflect() protoreflect.Message {
	mi := &file_einride_example_freight_v1_shipment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineItem.ProtoReflect.Descriptor instead.
func (*LineItem) Descriptor() ([]byte, []int) {
	return file_einride_example_freight_v1_shipment_proto_rawDescGZIP(), []int{1}
}

func (x *LineItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *LineItem) GetQuantity() float32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *LineItem) GetWeightKg() float32 {
	if x != nil {
		return x.WeightKg
	}
	return 0
}

func (x *LineItem) GetVolumeM3() float32 {
	if x != nil {
		return x.VolumeM3
	}
	return 0
}

var File_einride_example_freight_v1_shipment_proto protoreflect.FileDescriptor

var file_einride_example_freight_v1_shipment_proto_rawDesc = []byte{
	0x0a, 0x29, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x65, 0x69, 0x6e,
	0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x66, 0x72, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa3, 0x08, 0x0a, 0x08, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03,
	0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x0b,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x69, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x2a, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x66, 0x72, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x65, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x53, 0x69, 0x74, 0x65, 0x52, 0x0a, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x53, 0x69, 0x74, 0x65, 0x12, 0x55, 0x0a, 0x10, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2a, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41, 0x23, 0x0a, 0x21, 0x66, 0x72,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x65, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x53, 0x69, 0x74, 0x65, 0x52,
	0x0f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x74, 0x65,
	0x12, 0x52, 0x0a, 0x14, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x65, 0x61, 0x72, 0x6c, 0x69,
	0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02,
	0x52, 0x12, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x45, 0x61, 0x72, 0x6c, 0x69, 0x65, 0x73, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x12, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41,
	0x01, 0x02, 0x52, 0x10, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x56, 0x0a, 0x16, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x5f, 0x65, 0x61, 0x72, 0x6c, 0x69, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x14, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x45, 0x61, 0x72, 0x6c, 0x69, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x52, 0x0a, 0x14,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x12, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x43, 0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x0b,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x6c, 0x69, 0x6e, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x57, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x65, 0x69, 0x6e,
	0x72, 0x69, 0x64, 0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x66, 0x72, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3e,
	0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x68,
	0xea, 0x41, 0x65, 0x0a, 0x25, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x2d, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63,
	0x68, 0x2f, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x73, 0x68, 0x69, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x7d, 0x2f, 0x73,
	0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x7d, 0x2a, 0x09, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x08,
	0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x76, 0x0a, 0x08, 0x4c, 0x69, 0x6e, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x5f, 0x6b, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x4b, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x6d, 0x33,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4d, 0x33,
	0x42, 0xa3, 0x01, 0x0a, 0x1f, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64,
	0x65, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2d,
	0x68, 0x74, 0x74, 0x70, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x66, 0x72, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x66, 0x72, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_example_freight_v1_shipment_proto_rawDescOnce sync.Once
	file_einride_example_freight_v1_shipment_proto_rawDescData = file_einride_example_freight_v1_shipment_proto_rawDesc
)

func file_einride_example_freight_v1_shipment_proto_rawDescGZIP() []byte {
	file_einride_example_freight_v1_shipment_proto_rawDescOnce.Do(func() {
		file_einride_example_freight_v1_shipment_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_example_freight_v1_shipment_proto_rawDescData)
	})
	return file_einride_example_freight_v1_shipment_proto_rawDescData
}

var file_einride_example_freight_v1_shipment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_einride_example_freight_v1_shipment_proto_goTypes = []interface{}{
	(*Shipment)(nil),            // 0: einride.example.freight.v1.Shipment
	(*LineItem)(nil),            // 1: einride.example.freight.v1.LineItem
	nil,                         // 2: einride.example.freight.v1.Shipment.AnnotationsEntry
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_einride_example_freight_v1_shipment_proto_depIdxs = []int32{
	3, // 0: einride.example.freight.v1.Shipment.create_time:type_name -> google.protobuf.Timestamp
	3, // 1: einride.example.freight.v1.Shipment.update_time:type_name -> google.protobuf.Timestamp
	3, // 2: einride.example.freight.v1.Shipment.delete_time:type_name -> google.protobuf.Timestamp
	3, // 3: einride.example.freight.v1.Shipment.pickup_earliest_time:type_name -> google.protobuf.Timestamp
	3, // 4: einride.example.freight.v1.Shipment.pickup_latest_time:type_name -> google.protobuf.Timestamp
	3, // 5: einride.example.freight.v1.Shipment.delivery_earliest_time:type_name -> google.protobuf.Timestamp
	3, // 6: einride.example.freight.v1.Shipment.delivery_latest_time:type_name -> google.protobuf.Timestamp
	1, // 7: einride.example.freight.v1.Shipment.line_items:type_name -> einride.example.freight.v1.LineItem
	2, // 8: einride.example.freight.v1.Shipment.annotations:type_name -> einride.example.freight.v1.Shipment.AnnotationsEntry
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_einride_example_freight_v1_shipment_proto_init() }
func file_einride_example_freight_v1_shipment_proto_init() {
	if File_einride_example_freight_v1_shipment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_einride_example_freight_v1_shipment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shipment); i {
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
		file_einride_example_freight_v1_shipment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineItem); i {
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
			RawDescriptor: file_einride_example_freight_v1_shipment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_example_freight_v1_shipment_proto_goTypes,
		DependencyIndexes: file_einride_example_freight_v1_shipment_proto_depIdxs,
		MessageInfos:      file_einride_example_freight_v1_shipment_proto_msgTypes,
	}.Build()
	File_einride_example_freight_v1_shipment_proto = out.File
	file_einride_example_freight_v1_shipment_proto_rawDesc = nil
	file_einride_example_freight_v1_shipment_proto_goTypes = nil
	file_einride_example_freight_v1_shipment_proto_depIdxs = nil
}
