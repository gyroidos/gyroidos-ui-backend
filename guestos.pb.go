//
// This file is part of GyroidOS
// Copyright(c) 2013 - 2017 Fraunhofer AISEC
// Fraunhofer-Gesellschaft zur Förderung der angewandten Forschung e.V.
//
// This program is free software; you can redistribute it and/or modify it
// under the terms and conditions of the GNU General Public License,
// version 2 (GPL 2), as published by the Free Software Foundation.
//
// This program is distributed in the hope it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE. See the GPL 2 license for more details.
//
// You should have received a copy of the GNU General Public License along with
// this program; if not, see <http://www.gnu.org/licenses/>
//
// The full GNU General Public License is included in this distribution in
// the file called "COPYING".
//
// Contact Information:
// Fraunhofer AISEC <gyroidos@aisec.fraunhofer.de>

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.14.0
// source: guestos.proto

package main

import (
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

type GuestOSMount_Type int32

const (
	GuestOSMount_SHARED       GuestOSMount_Type = 1 // image shared by all containers of same OS type
	GuestOSMount_DEVICE       GuestOSMount_Type = 2 // image file is copied from a device partition
	GuestOSMount_DEVICE_RW    GuestOSMount_Type = 3 // image file is copied from a device partition
	GuestOSMount_EMPTY        GuestOSMount_Type = 4 // empty, created when a corresponding container is first instantiated
	GuestOSMount_COPY         GuestOSMount_Type = 5 // deprecated
	GuestOSMount_FLASH        GuestOSMount_Type = 6 // image to be flashed to a partition (e.g. boot.img; base system updates only for now)
	GuestOSMount_OVERLAY_RO   GuestOSMount_Type = 7 // read only overlay images e.g. for system features
	GuestOSMount_SHARED_RW    GuestOSMount_Type = 8 // image shared by all containers of same OS type (writable tmpfs as overlay)
	GuestOSMount_OVERLAY_RW   GuestOSMount_Type = 9 // similar to EMPTY image, however overlayed on given mount_point (writable persitent fs as overlay)
	GuestOSMount_BIND_FILE    GuestOSMount_Type = 10
	GuestOSMount_BIND_FILE_RW GuestOSMount_Type = 11
)

// Enum value maps for GuestOSMount_Type.
var (
	GuestOSMount_Type_name = map[int32]string{
		1:  "SHARED",
		2:  "DEVICE",
		3:  "DEVICE_RW",
		4:  "EMPTY",
		5:  "COPY",
		6:  "FLASH",
		7:  "OVERLAY_RO",
		8:  "SHARED_RW",
		9:  "OVERLAY_RW",
		10: "BIND_FILE",
		11: "BIND_FILE_RW",
	}
	GuestOSMount_Type_value = map[string]int32{
		"SHARED":       1,
		"DEVICE":       2,
		"DEVICE_RW":    3,
		"EMPTY":        4,
		"COPY":         5,
		"FLASH":        6,
		"OVERLAY_RO":   7,
		"SHARED_RW":    8,
		"OVERLAY_RW":   9,
		"BIND_FILE":    10,
		"BIND_FILE_RW": 11,
	}
)

func (x GuestOSMount_Type) Enum() *GuestOSMount_Type {
	p := new(GuestOSMount_Type)
	*p = x
	return p
}

func (x GuestOSMount_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GuestOSMount_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_guestos_proto_enumTypes[0].Descriptor()
}

func (GuestOSMount_Type) Type() protoreflect.EnumType {
	return &file_guestos_proto_enumTypes[0]
}

func (x GuestOSMount_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *GuestOSMount_Type) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = GuestOSMount_Type(num)
	return nil
}

// Deprecated: Use GuestOSMount_Type.Descriptor instead.
func (GuestOSMount_Type) EnumDescriptor() ([]byte, []int) {
	return file_guestos_proto_rawDescGZIP(), []int{1, 0}
}

// (English) string with optional translations in other languages
type I18NString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	En *string `protobuf:"bytes,1,req,name=en" json:"en,omitempty"`
	De *string `protobuf:"bytes,2,opt,name=de" json:"de,omitempty"`
	Fr *string `protobuf:"bytes,3,opt,name=fr" json:"fr,omitempty"` // TODO: add other languages as necessary...
}

func (x *I18NString) Reset() {
	*x = I18NString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guestos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *I18NString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*I18NString) ProtoMessage() {}

func (x *I18NString) ProtoReflect() protoreflect.Message {
	mi := &file_guestos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use I18NString.ProtoReflect.Descriptor instead.
func (*I18NString) Descriptor() ([]byte, []int) {
	return file_guestos_proto_rawDescGZIP(), []int{0}
}

func (x *I18NString) GetEn() string {
	if x != nil && x.En != nil {
		return *x.En
	}
	return ""
}

func (x *I18NString) GetDe() string {
	if x != nil && x.De != nil {
		return *x.De
	}
	return ""
}

func (x *I18NString) GetFr() string {
	if x != nil && x.Fr != nil {
		return *x.Fr
	}
	return ""
}

// Represents a single mount point/image provided by the GuestOS
type GuestOSMount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path for the image files is derived from guestos_path/guestos_name.
	ImageFile  *string            `protobuf:"bytes,1,req,name=image_file,json=imageFile" json:"image_file,omitempty"`                              // name of the image file, e.g. system
	MountPoint *string            `protobuf:"bytes,2,req,name=mount_point,json=mountPoint" json:"mount_point,omitempty"`                           // mountpoint inside the container
	FsType     *string            `protobuf:"bytes,3,req,name=fs_type,json=fsType" json:"fs_type,omitempty"`                                       // file system type, e.g. "ext4"
	MountType  *GuestOSMount_Type `protobuf:"varint,4,req,name=mount_type,json=mountType,enum=main.GuestOSMount_Type" json:"mount_type,omitempty"` // type of the image file
	// The following three fields are only used for EMPTY mount types:
	MinSize *uint32 `protobuf:"varint,6,opt,name=min_size,json=minSize,def=10" json:"min_size,omitempty"`    // required minimum size (MBytes) for EMPTY partition
	MaxSize *uint32 `protobuf:"varint,7,opt,name=max_size,json=maxSize,def=16384" json:"max_size,omitempty"` // allowed maximum size (MBytes) for EMPTY partition
	DefSize *uint32 `protobuf:"varint,8,opt,name=def_size,json=defSize,def=1024" json:"def_size,omitempty"`  // default size (MBytes) for EMPTY partition
	// The following two fields are only used when an actual image file is provided:
	ImageSize *uint64 `protobuf:"varint,10,opt,name=image_size,json=imageSize" json:"image_size,omitempty"` // size (bytes) of the image
	// hash(es) of image file
	ImageSha1         *string `protobuf:"bytes,11,opt,name=image_sha1,json=imageSha1" json:"image_sha1,omitempty"`
	ImageSha2_256     *string `protobuf:"bytes,12,opt,name=image_sha2_256,json=imageSha2256" json:"image_sha2_256,omitempty"` // TODO add further hashes as necessary
	MountData         *string `protobuf:"bytes,13,opt,name=mount_data,json=mountData" json:"mount_data,omitempty"`            // mount_data used for mount syscall, e.g. "context=" for selinux
	ImageVeritySha256 *string `protobuf:"bytes,14,opt,name=image_verity_sha256,json=imageVeritySha256" json:"image_verity_sha256,omitempty"`
}

// Default values for GuestOSMount fields.
const (
	Default_GuestOSMount_MinSize = uint32(10)
	Default_GuestOSMount_MaxSize = uint32(16384)
	Default_GuestOSMount_DefSize = uint32(1024)
)

func (x *GuestOSMount) Reset() {
	*x = GuestOSMount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guestos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuestOSMount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuestOSMount) ProtoMessage() {}

func (x *GuestOSMount) ProtoReflect() protoreflect.Message {
	mi := &file_guestos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuestOSMount.ProtoReflect.Descriptor instead.
func (*GuestOSMount) Descriptor() ([]byte, []int) {
	return file_guestos_proto_rawDescGZIP(), []int{1}
}

func (x *GuestOSMount) GetImageFile() string {
	if x != nil && x.ImageFile != nil {
		return *x.ImageFile
	}
	return ""
}

func (x *GuestOSMount) GetMountPoint() string {
	if x != nil && x.MountPoint != nil {
		return *x.MountPoint
	}
	return ""
}

func (x *GuestOSMount) GetFsType() string {
	if x != nil && x.FsType != nil {
		return *x.FsType
	}
	return ""
}

func (x *GuestOSMount) GetMountType() GuestOSMount_Type {
	if x != nil && x.MountType != nil {
		return *x.MountType
	}
	return GuestOSMount_SHARED
}

func (x *GuestOSMount) GetMinSize() uint32 {
	if x != nil && x.MinSize != nil {
		return *x.MinSize
	}
	return Default_GuestOSMount_MinSize
}

func (x *GuestOSMount) GetMaxSize() uint32 {
	if x != nil && x.MaxSize != nil {
		return *x.MaxSize
	}
	return Default_GuestOSMount_MaxSize
}

func (x *GuestOSMount) GetDefSize() uint32 {
	if x != nil && x.DefSize != nil {
		return *x.DefSize
	}
	return Default_GuestOSMount_DefSize
}

func (x *GuestOSMount) GetImageSize() uint64 {
	if x != nil && x.ImageSize != nil {
		return *x.ImageSize
	}
	return 0
}

func (x *GuestOSMount) GetImageSha1() string {
	if x != nil && x.ImageSha1 != nil {
		return *x.ImageSha1
	}
	return ""
}

func (x *GuestOSMount) GetImageSha2_256() string {
	if x != nil && x.ImageSha2_256 != nil {
		return *x.ImageSha2_256
	}
	return ""
}

func (x *GuestOSMount) GetMountData() string {
	if x != nil && x.MountData != nil {
		return *x.MountData
	}
	return ""
}

func (x *GuestOSMount) GetImageVeritySha256() string {
	if x != nil && x.ImageVeritySha256 != nil {
		return *x.ImageVeritySha256
	}
	return ""
}

type GuestOSConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The following three fields together should UNIQUELY identify the actual guestos image files and config:
	Name     *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`         // unique name of the operating system
	Hardware *string `protobuf:"bytes,2,req,name=hardware" json:"hardware,omitempty"` // target hardware; must match hardware_get_name(), e.g. "i9505", etc.
	// TODO version string OR major.minor-suffix? ("RC1", "beta", etc.)
	Version     *uint64         `protobuf:"varint,3,req,name=version" json:"version,omitempty"`                            // version string for the guest OS
	Mounts      []*GuestOSMount `protobuf:"bytes,5,rep,name=mounts" json:"mounts,omitempty"`                               // list of mounts inside the container
	InitPath    *string         `protobuf:"bytes,6,opt,name=init_path,json=initPath,def=/init" json:"init_path,omitempty"` // path to the init binary, e.g. "/init"
	InitParam   []string        `protobuf:"bytes,7,rep,name=init_param,json=initParam" json:"init_param,omitempty"`        // parameters (argv) for init
	InitEnv     []string        `protobuf:"bytes,8,rep,name=init_env,json=initEnv" json:"init_env,omitempty"`              // environment variables
	MountsSetup []*GuestOSMount `protobuf:"bytes,9,rep,name=mounts_setup,json=mountsSetup" json:"mounts_setup,omitempty"`  // list of mounts for setup mode of a container
	// Flags indicating the features supported by the OS:
	FeatureBgBooting    *bool `protobuf:"varint,27,opt,name=feature_bg_booting,json=featureBgBooting,def=1" json:"feature_bg_booting,omitempty"`
	FeatureInstallGuest *bool `protobuf:"varint,40,opt,name=feature_install_guest,json=featureInstallGuest,def=0" json:"feature_install_guest,omitempty"` // TODO: add further features as necessary
	// TODO: Determine RAM limit policy...
	MinRamLimit *uint32 `protobuf:"varint,30,opt,name=min_ram_limit,json=minRamLimit,def=128" json:"min_ram_limit,omitempty"`  // required minimum RAM size (MBytes)
	DefRamLimit *uint32 `protobuf:"varint,32,opt,name=def_ram_limit,json=defRamLimit,def=1024" json:"def_ram_limit,omitempty"` // default RAM size (MBytes)
	// Descriptive information (for GUI):
	Description     *I18NString `protobuf:"bytes,10,req,name=description" json:"description,omitempty"`                                // description/full name
	DescriptionLong *I18NString `protobuf:"bytes,11,opt,name=description_long,json=descriptionLong" json:"description_long,omitempty"` // help text
	UpstreamVersion *string     `protobuf:"bytes,12,opt,name=upstream_version,json=upstreamVersion" json:"upstream_version,omitempty"` // upstream version
	Icon            []byte      `protobuf:"bytes,13,opt,name=icon" json:"icon,omitempty"`                                              // name of icon file
	BuildDate       *string     `protobuf:"bytes,14,opt,name=build_date,json=buildDate" json:"build_date,omitempty"`                   // build date
	UpdateBaseUrl   *string     `protobuf:"bytes,15,opt,name=update_base_url,json=updateBaseUrl" json:"update_base_url,omitempty"`     // provide url to file server which hosts the actual image data (overwrites device.conf)
}

// Default values for GuestOSConfig fields.
const (
	Default_GuestOSConfig_InitPath            = string("/init")
	Default_GuestOSConfig_FeatureBgBooting    = bool(true)
	Default_GuestOSConfig_FeatureInstallGuest = bool(false)
	Default_GuestOSConfig_MinRamLimit         = uint32(128)
	Default_GuestOSConfig_DefRamLimit         = uint32(1024)
)

func (x *GuestOSConfig) Reset() {
	*x = GuestOSConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_guestos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuestOSConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuestOSConfig) ProtoMessage() {}

func (x *GuestOSConfig) ProtoReflect() protoreflect.Message {
	mi := &file_guestos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuestOSConfig.ProtoReflect.Descriptor instead.
func (*GuestOSConfig) Descriptor() ([]byte, []int) {
	return file_guestos_proto_rawDescGZIP(), []int{2}
}

func (x *GuestOSConfig) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *GuestOSConfig) GetHardware() string {
	if x != nil && x.Hardware != nil {
		return *x.Hardware
	}
	return ""
}

func (x *GuestOSConfig) GetVersion() uint64 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

func (x *GuestOSConfig) GetMounts() []*GuestOSMount {
	if x != nil {
		return x.Mounts
	}
	return nil
}

func (x *GuestOSConfig) GetInitPath() string {
	if x != nil && x.InitPath != nil {
		return *x.InitPath
	}
	return Default_GuestOSConfig_InitPath
}

func (x *GuestOSConfig) GetInitParam() []string {
	if x != nil {
		return x.InitParam
	}
	return nil
}

func (x *GuestOSConfig) GetInitEnv() []string {
	if x != nil {
		return x.InitEnv
	}
	return nil
}

func (x *GuestOSConfig) GetMountsSetup() []*GuestOSMount {
	if x != nil {
		return x.MountsSetup
	}
	return nil
}

func (x *GuestOSConfig) GetFeatureBgBooting() bool {
	if x != nil && x.FeatureBgBooting != nil {
		return *x.FeatureBgBooting
	}
	return Default_GuestOSConfig_FeatureBgBooting
}

func (x *GuestOSConfig) GetFeatureInstallGuest() bool {
	if x != nil && x.FeatureInstallGuest != nil {
		return *x.FeatureInstallGuest
	}
	return Default_GuestOSConfig_FeatureInstallGuest
}

func (x *GuestOSConfig) GetMinRamLimit() uint32 {
	if x != nil && x.MinRamLimit != nil {
		return *x.MinRamLimit
	}
	return Default_GuestOSConfig_MinRamLimit
}

func (x *GuestOSConfig) GetDefRamLimit() uint32 {
	if x != nil && x.DefRamLimit != nil {
		return *x.DefRamLimit
	}
	return Default_GuestOSConfig_DefRamLimit
}

func (x *GuestOSConfig) GetDescription() *I18NString {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *GuestOSConfig) GetDescriptionLong() *I18NString {
	if x != nil {
		return x.DescriptionLong
	}
	return nil
}

func (x *GuestOSConfig) GetUpstreamVersion() string {
	if x != nil && x.UpstreamVersion != nil {
		return *x.UpstreamVersion
	}
	return ""
}

func (x *GuestOSConfig) GetIcon() []byte {
	if x != nil {
		return x.Icon
	}
	return nil
}

func (x *GuestOSConfig) GetBuildDate() string {
	if x != nil && x.BuildDate != nil {
		return *x.BuildDate
	}
	return ""
}

func (x *GuestOSConfig) GetUpdateBaseUrl() string {
	if x != nil && x.UpdateBaseUrl != nil {
		return *x.UpdateBaseUrl
	}
	return ""
}

var File_guestos_proto protoreflect.FileDescriptor

var file_guestos_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x75, 0x65, 0x73, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x3c, 0x0a, 0x0a, 0x49, 0x31, 0x38, 0x4e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52,
	0x02, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x66, 0x72, 0x22, 0xd4, 0x04, 0x0a, 0x0c, 0x47, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x53, 0x4d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x06, 0x66, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a,
	0x0a, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x02, 0x28,
	0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x53,
	0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x02, 0x31, 0x30, 0x52, 0x07, 0x6d, 0x69, 0x6e,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x05, 0x31, 0x36, 0x33, 0x38, 0x34, 0x52, 0x07, 0x6d,
	0x61, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x08, 0x64, 0x65, 0x66, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x04, 0x31, 0x30, 0x32, 0x34, 0x52, 0x07,
	0x64, 0x65, 0x66, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x68, 0x61, 0x31, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x53, 0x68, 0x61, 0x31, 0x12, 0x24, 0x0a, 0x0e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x68, 0x61, 0x32, 0x5f, 0x32, 0x35, 0x36, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x53, 0x68, 0x61, 0x32, 0x32, 0x35, 0x36, 0x12, 0x1d, 0x0a, 0x0a, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x68, 0x61, 0x32, 0x35,
	0x36, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x56, 0x65,
	0x72, 0x69, 0x74, 0x79, 0x53, 0x68, 0x61, 0x32, 0x35, 0x36, 0x22, 0x9d, 0x01, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x48, 0x41, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x44,
	0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x52, 0x57, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d,
	0x50, 0x54, 0x59, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x4f, 0x50, 0x59, 0x10, 0x05, 0x12,
	0x09, 0x0a, 0x05, 0x46, 0x4c, 0x41, 0x53, 0x48, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x56,
	0x45, 0x52, 0x4c, 0x41, 0x59, 0x5f, 0x52, 0x4f, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x48,
	0x41, 0x52, 0x45, 0x44, 0x5f, 0x52, 0x57, 0x10, 0x08, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x56, 0x45,
	0x52, 0x4c, 0x41, 0x59, 0x5f, 0x52, 0x57, 0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x49, 0x4e,
	0x44, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x49, 0x4e, 0x44,
	0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x52, 0x57, 0x10, 0x0b, 0x22, 0xe5, 0x05, 0x0a, 0x0d, 0x47,
	0x75, 0x65, 0x73, 0x74, 0x4f, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x08, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x02, 0x28, 0x04, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x06, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x75,
	0x65, 0x73, 0x74, 0x4f, 0x53, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x12, 0x22, 0x0a, 0x09, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x3a, 0x05, 0x2f, 0x69, 0x6e, 0x69, 0x74, 0x52, 0x08, 0x69, 0x6e,
	0x69, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x69, 0x74,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x65, 0x6e,
	0x76, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x69, 0x74, 0x45, 0x6e, 0x76,
	0x12, 0x35, 0x0a, 0x0c, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x74, 0x75, 0x70,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x75,
	0x65, 0x73, 0x74, 0x4f, 0x53, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0b, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x53, 0x65, 0x74, 0x75, 0x70, 0x12, 0x32, 0x0a, 0x12, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x62, 0x67, 0x5f, 0x62, 0x6f, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x1b, 0x20,
	0x01, 0x28, 0x08, 0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x10, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x42, 0x67, 0x42, 0x6f, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x39, 0x0a, 0x15, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x5f, 0x67,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73,
	0x65, 0x52, 0x13, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x47, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x61,
	0x6d, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x03, 0x31,
	0x32, 0x38, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x52, 0x61, 0x6d, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x28, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x5f, 0x72, 0x61, 0x6d, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x20, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x04, 0x31, 0x30, 0x32, 0x34, 0x52, 0x0b, 0x64, 0x65,
	0x66, 0x52, 0x61, 0x6d, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x32, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x49, 0x31, 0x38, 0x4e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a,
	0x10, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x6e,
	0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x49,
	0x31, 0x38, 0x4e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x6e, 0x67, 0x12, 0x29, 0x0a, 0x10, 0x75, 0x70,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c,
	0x4a, 0x04, 0x08, 0x14, 0x10, 0x1b, 0x4a, 0x04, 0x08, 0x1c, 0x10, 0x1d, 0x4a, 0x04, 0x08, 0x1d,
	0x10, 0x1e, 0x42, 0x24, 0x0a, 0x1b, 0x64, 0x65, 0x2e, 0x66, 0x72, 0x61, 0x75, 0x6e, 0x68, 0x6f,
	0x66, 0x65, 0x72, 0x2e, 0x61, 0x69, 0x73, 0x65, 0x63, 0x2e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x6d,
	0x65, 0x5a, 0x05, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
}

var (
	file_guestos_proto_rawDescOnce sync.Once
	file_guestos_proto_rawDescData = file_guestos_proto_rawDesc
)

func file_guestos_proto_rawDescGZIP() []byte {
	file_guestos_proto_rawDescOnce.Do(func() {
		file_guestos_proto_rawDescData = protoimpl.X.CompressGZIP(file_guestos_proto_rawDescData)
	})
	return file_guestos_proto_rawDescData
}

var file_guestos_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_guestos_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_guestos_proto_goTypes = []interface{}{
	(GuestOSMount_Type)(0), // 0: main.GuestOSMount.Type
	(*I18NString)(nil),     // 1: main.I18NString
	(*GuestOSMount)(nil),   // 2: main.GuestOSMount
	(*GuestOSConfig)(nil),  // 3: main.GuestOSConfig
}
var file_guestos_proto_depIdxs = []int32{
	0, // 0: main.GuestOSMount.mount_type:type_name -> main.GuestOSMount.Type
	2, // 1: main.GuestOSConfig.mounts:type_name -> main.GuestOSMount
	2, // 2: main.GuestOSConfig.mounts_setup:type_name -> main.GuestOSMount
	1, // 3: main.GuestOSConfig.description:type_name -> main.I18NString
	1, // 4: main.GuestOSConfig.description_long:type_name -> main.I18NString
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_guestos_proto_init() }
func file_guestos_proto_init() {
	if File_guestos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_guestos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*I18NString); i {
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
		file_guestos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuestOSMount); i {
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
		file_guestos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuestOSConfig); i {
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
			RawDescriptor: file_guestos_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_guestos_proto_goTypes,
		DependencyIndexes: file_guestos_proto_depIdxs,
		EnumInfos:         file_guestos_proto_enumTypes,
		MessageInfos:      file_guestos_proto_msgTypes,
	}.Build()
	File_guestos_proto = out.File
	file_guestos_proto_rawDesc = nil
	file_guestos_proto_goTypes = nil
	file_guestos_proto_depIdxs = nil
}
