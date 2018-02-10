// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hot_update.proto

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Ignoring public import of ProtoHeader from common_client.proto

// Ignoring public import of Heartbeat from common_client.proto

// Ignoring public import of ServerInfo from common_client.proto

// Ignoring public import of QuickConn from common_client.proto

// Ignoring public import of AckQuickConn from common_client.proto

// Ignoring public import of WeixinInfo from common_client.proto

// Ignoring public import of common_req_reg from common_client.proto

// Ignoring public import of common_req_reg_via_input from common_client.proto

// Ignoring public import of common_ack_reg from common_client.proto

// Ignoring public import of common_req_gameLogin from common_client.proto

// Ignoring public import of common_req_gameLogin_via_input from common_client.proto

// Ignoring public import of common_ack_gameLogin from common_client.proto

// Ignoring public import of common_req_qrLogin from common_client.proto

// Ignoring public import of common_ack_qrLogin from common_client.proto

// Ignoring public import of common_req_qrWxInfo from common_client.proto

// Ignoring public import of common_ack_qrWxInfo from common_client.proto

// Ignoring public import of common_ack_reconnect from common_client.proto

// Ignoring public import of common_req_reconnect from common_client.proto

// Ignoring public import of common_req_gameState from common_client.proto

// Ignoring public import of common_ack_gameState from common_client.proto

// Ignoring public import of common_req_logout from common_client.proto

// Ignoring public import of common_ack_logout from common_client.proto

// Ignoring public import of common_req_feedback from common_client.proto

// Ignoring public import of client_base_poker from common_client.proto

// Ignoring public import of common_req_message from common_client.proto

// Ignoring public import of common_bc_message from common_client.proto

// Ignoring public import of common_req_notice from common_client.proto

// Ignoring public import of common_ack_notice from common_client.proto

// Ignoring public import of common_req_enterAgentMode from common_client.proto

// Ignoring public import of common_ack_enterAgentMode from common_client.proto

// Ignoring public import of common_req_quitAgentMode from common_client.proto

// Ignoring public import of common_ack_quitAgentMode from common_client.proto

// Ignoring public import of common_req_leaveDesk from common_client.proto

// Ignoring public import of common_ack_leaveDesk from common_client.proto

// Ignoring public import of common_req_kickout from common_client.proto

// Ignoring public import of common_bc_kickout from common_client.proto

// Ignoring public import of common_req_allowance from common_client.proto

// Ignoring public import of common_ack_allowance from common_client.proto

// Ignoring public import of common_req_applyDissolve from common_client.proto

// Ignoring public import of common_bc_applyDissolve from common_client.proto

// Ignoring public import of common_req_applyDissolveBack from common_client.proto

// Ignoring public import of common_ack_applyDissolveBack from common_client.proto

// Ignoring public import of common_ack_timeout from common_client.proto

// Ignoring public import of common_bc_userBreak from common_client.proto

// Ignoring public import of common_req_clickStatistic from common_client.proto

// Ignoring public import of common_req_offline from common_client.proto

// Ignoring public import of common_req_upload_location from common_client.proto

// Ignoring public import of common_bc_leaveTimeout from common_client.proto

// Ignoring public import of common_desk_by_agent from common_client.proto

// Ignoring public import of common_req_list_coin_desk from common_client.proto

// Ignoring public import of common_ack_list_coin_desk from common_client.proto

// Ignoring public import of CommonCoinDeskInfo from common_client.proto

// Ignoring public import of common_req_list_coin_info from common_client.proto

// Ignoring public import of common_ack_list_coin_info from common_client.proto

// Ignoring public import of CommonCoinLevelInfo from common_client.proto

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// Ignoring public import of common_enum_game from common_enum.proto

// Ignoring public import of COMMON_ENUM_ROOMTYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_GAMESTATUS from common_enum.proto

// Ignoring public import of COMMON_ENUM_RELEASETAG from common_enum.proto

// Ignoring public import of COMMON_ENUM_KICKOUT from common_enum.proto

// Ignoring public import of COMMON_ENUM_APPLYDISSOLVE from common_enum.proto

// Ignoring public import of BTN_TYPE from common_enum.proto

// Ignoring public import of COMMON_ENUM_ERROR_TYPE_ENTERCOINROOM from common_enum.proto

// Ignoring public import of COMMON_ENUM_ROOMCARD_BILL_TYPE from common_enum.proto

type VersionInfo struct {
	FileId           *int32 `protobuf:"varint,1,opt,name=fileId" json:"fileId,omitempty"`
	FileVer          *int32 `protobuf:"varint,2,opt,name=fileVer" json:"fileVer,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *VersionInfo) Reset()                    { *m = VersionInfo{} }
func (m *VersionInfo) String() string            { return proto.CompactTextString(m) }
func (*VersionInfo) ProtoMessage()               {}
func (*VersionInfo) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{0} }

func (m *VersionInfo) GetFileId() int32 {
	if m != nil && m.FileId != nil {
		return *m.FileId
	}
	return 0
}

func (m *VersionInfo) GetFileVer() int32 {
	if m != nil && m.FileVer != nil {
		return *m.FileVer
	}
	return 0
}

type AssetInfo struct {
	FileId           *int32          `protobuf:"varint,1,opt,name=fileId" json:"fileId,omitempty"`
	FileVer          *int32          `protobuf:"varint,2,opt,name=fileVer" json:"fileVer,omitempty"`
	FilePath         *string         `protobuf:"bytes,3,opt,name=filePath" json:"filePath,omitempty"`
	FileSize         *int32          `protobuf:"varint,4,opt,name=fileSize" json:"fileSize,omitempty"`
	Md5              *string         `protobuf:"bytes,5,opt,name=md5" json:"md5,omitempty"`
	IsCompress       *bool           `protobuf:"varint,6,opt,name=isCompress" json:"isCompress,omitempty"`
	IsCode           *bool           `protobuf:"varint,7,opt,name=isCode" json:"isCode,omitempty"`
	GameId           *CommonEnumGame `protobuf:"varint,8,opt,name=gameId,enum=ddproto.CommonEnumGame" json:"gameId,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *AssetInfo) Reset()                    { *m = AssetInfo{} }
func (m *AssetInfo) String() string            { return proto.CompactTextString(m) }
func (*AssetInfo) ProtoMessage()               {}
func (*AssetInfo) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{1} }

func (m *AssetInfo) GetFileId() int32 {
	if m != nil && m.FileId != nil {
		return *m.FileId
	}
	return 0
}

func (m *AssetInfo) GetFileVer() int32 {
	if m != nil && m.FileVer != nil {
		return *m.FileVer
	}
	return 0
}

func (m *AssetInfo) GetFilePath() string {
	if m != nil && m.FilePath != nil {
		return *m.FilePath
	}
	return ""
}

func (m *AssetInfo) GetFileSize() int32 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *AssetInfo) GetMd5() string {
	if m != nil && m.Md5 != nil {
		return *m.Md5
	}
	return ""
}

func (m *AssetInfo) GetIsCompress() bool {
	if m != nil && m.IsCompress != nil {
		return *m.IsCompress
	}
	return false
}

func (m *AssetInfo) GetIsCode() bool {
	if m != nil && m.IsCode != nil {
		return *m.IsCode
	}
	return false
}

func (m *AssetInfo) GetGameId() CommonEnumGame {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return CommonEnumGame_GID_SRC
}

// 获取资源文件的最新版本信息
type HotupdateReqVersionInfo struct {
	Header              *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ClientAppId         *int32       `protobuf:"varint,2,opt,name=clientAppId" json:"clientAppId,omitempty"`
	DownloadGameId      []int32      `protobuf:"varint,3,rep,name=downloadGameId" json:"downloadGameId,omitempty"`
	ClientAssetsVersion *int32       `protobuf:"varint,4,opt,name=clientAssetsVersion" json:"clientAssetsVersion,omitempty"`
	ClientResolution    *int32       `protobuf:"varint,5,opt,name=clientResolution" json:"clientResolution,omitempty"`
	AppVersion          *int32       `protobuf:"varint,6,opt,name=AppVersion" json:"AppVersion,omitempty"`
	NativeVersion       *int32       `protobuf:"varint,7,opt,name=nativeVersion" json:"nativeVersion,omitempty"`
	ClientOSType        *int32       `protobuf:"varint,8,opt,name=clientOSType" json:"clientOSType,omitempty"`
	XXX_unrecognized    []byte       `json:"-"`
}

func (m *HotupdateReqVersionInfo) Reset()                    { *m = HotupdateReqVersionInfo{} }
func (m *HotupdateReqVersionInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateReqVersionInfo) ProtoMessage()               {}
func (*HotupdateReqVersionInfo) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{2} }

func (m *HotupdateReqVersionInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateReqVersionInfo) GetClientAppId() int32 {
	if m != nil && m.ClientAppId != nil {
		return *m.ClientAppId
	}
	return 0
}

func (m *HotupdateReqVersionInfo) GetDownloadGameId() []int32 {
	if m != nil {
		return m.DownloadGameId
	}
	return nil
}

func (m *HotupdateReqVersionInfo) GetClientAssetsVersion() int32 {
	if m != nil && m.ClientAssetsVersion != nil {
		return *m.ClientAssetsVersion
	}
	return 0
}

func (m *HotupdateReqVersionInfo) GetClientResolution() int32 {
	if m != nil && m.ClientResolution != nil {
		return *m.ClientResolution
	}
	return 0
}

func (m *HotupdateReqVersionInfo) GetAppVersion() int32 {
	if m != nil && m.AppVersion != nil {
		return *m.AppVersion
	}
	return 0
}

func (m *HotupdateReqVersionInfo) GetNativeVersion() int32 {
	if m != nil && m.NativeVersion != nil {
		return *m.NativeVersion
	}
	return 0
}

func (m *HotupdateReqVersionInfo) GetClientOSType() int32 {
	if m != nil && m.ClientOSType != nil {
		return *m.ClientOSType
	}
	return 0
}

type HotupdateAckVersionInfo struct {
	Header               *ProtoHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Version              []*VersionInfo `protobuf:"bytes,2,rep,name=version" json:"version,omitempty"`
	LastestAssetsVersion *int32         `protobuf:"varint,3,opt,name=lastestAssetsVersion" json:"lastestAssetsVersion,omitempty"`
	ForceUpdate          *bool          `protobuf:"varint,4,opt,name=forceUpdate" json:"forceUpdate,omitempty"`
	ReleaseTag           *int32         `protobuf:"varint,5,opt,name=releaseTag" json:"releaseTag,omitempty"`
	AppDownloadUrl       *string        `protobuf:"bytes,6,opt,name=appDownloadUrl" json:"appDownloadUrl,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *HotupdateAckVersionInfo) Reset()                    { *m = HotupdateAckVersionInfo{} }
func (m *HotupdateAckVersionInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateAckVersionInfo) ProtoMessage()               {}
func (*HotupdateAckVersionInfo) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{3} }

func (m *HotupdateAckVersionInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateAckVersionInfo) GetVersion() []*VersionInfo {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *HotupdateAckVersionInfo) GetLastestAssetsVersion() int32 {
	if m != nil && m.LastestAssetsVersion != nil {
		return *m.LastestAssetsVersion
	}
	return 0
}

func (m *HotupdateAckVersionInfo) GetForceUpdate() bool {
	if m != nil && m.ForceUpdate != nil {
		return *m.ForceUpdate
	}
	return false
}

func (m *HotupdateAckVersionInfo) GetReleaseTag() int32 {
	if m != nil && m.ReleaseTag != nil {
		return *m.ReleaseTag
	}
	return 0
}

func (m *HotupdateAckVersionInfo) GetAppDownloadUrl() string {
	if m != nil && m.AppDownloadUrl != nil {
		return *m.AppDownloadUrl
	}
	return ""
}

// 获取资源文件的详细信息
type HotupdateReqAssetsInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ClientAppId      *int32       `protobuf:"varint,2,opt,name=clientAppId" json:"clientAppId,omitempty"`
	FileIds          []int32      `protobuf:"varint,3,rep,name=fileIds" json:"fileIds,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HotupdateReqAssetsInfo) Reset()                    { *m = HotupdateReqAssetsInfo{} }
func (m *HotupdateReqAssetsInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateReqAssetsInfo) ProtoMessage()               {}
func (*HotupdateReqAssetsInfo) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{4} }

func (m *HotupdateReqAssetsInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateReqAssetsInfo) GetClientAppId() int32 {
	if m != nil && m.ClientAppId != nil {
		return *m.ClientAppId
	}
	return 0
}

func (m *HotupdateReqAssetsInfo) GetFileIds() []int32 {
	if m != nil {
		return m.FileIds
	}
	return nil
}

type HotupdateAckAssetsInfo struct {
	Header               *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	AssetHost            *string      `protobuf:"bytes,2,opt,name=assetHost" json:"assetHost,omitempty"`
	Assets               []*AssetInfo `protobuf:"bytes,3,rep,name=assets" json:"assets,omitempty"`
	LastestAssetsVersion *int32       `protobuf:"varint,4,opt,name=lastestAssetsVersion" json:"lastestAssetsVersion,omitempty"`
	ClientResolution     *int32       `protobuf:"varint,5,opt,name=clientResolution" json:"clientResolution,omitempty"`
	ClientAppId          *int32       `protobuf:"varint,6,opt,name=clientAppId" json:"clientAppId,omitempty"`
	XXX_unrecognized     []byte       `json:"-"`
}

func (m *HotupdateAckAssetsInfo) Reset()                    { *m = HotupdateAckAssetsInfo{} }
func (m *HotupdateAckAssetsInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateAckAssetsInfo) ProtoMessage()               {}
func (*HotupdateAckAssetsInfo) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{5} }

func (m *HotupdateAckAssetsInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateAckAssetsInfo) GetAssetHost() string {
	if m != nil && m.AssetHost != nil {
		return *m.AssetHost
	}
	return ""
}

func (m *HotupdateAckAssetsInfo) GetAssets() []*AssetInfo {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *HotupdateAckAssetsInfo) GetLastestAssetsVersion() int32 {
	if m != nil && m.LastestAssetsVersion != nil {
		return *m.LastestAssetsVersion
	}
	return 0
}

func (m *HotupdateAckAssetsInfo) GetClientResolution() int32 {
	if m != nil && m.ClientResolution != nil {
		return *m.ClientResolution
	}
	return 0
}

func (m *HotupdateAckAssetsInfo) GetClientAppId() int32 {
	if m != nil && m.ClientAppId != nil {
		return *m.ClientAppId
	}
	return 0
}

// 获取某个独立游戏的下载信息
type HotupdateReqGameAssetsInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ClientAppId      *int32       `protobuf:"varint,2,opt,name=clientAppId" json:"clientAppId,omitempty"`
	GameId           *int32       `protobuf:"varint,3,opt,name=gameId" json:"gameId,omitempty"`
	ClientResolution *int32       `protobuf:"varint,4,opt,name=clientResolution" json:"clientResolution,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HotupdateReqGameAssetsInfo) Reset()                    { *m = HotupdateReqGameAssetsInfo{} }
func (m *HotupdateReqGameAssetsInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateReqGameAssetsInfo) ProtoMessage()               {}
func (*HotupdateReqGameAssetsInfo) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{6} }

func (m *HotupdateReqGameAssetsInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateReqGameAssetsInfo) GetClientAppId() int32 {
	if m != nil && m.ClientAppId != nil {
		return *m.ClientAppId
	}
	return 0
}

func (m *HotupdateReqGameAssetsInfo) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *HotupdateReqGameAssetsInfo) GetClientResolution() int32 {
	if m != nil && m.ClientResolution != nil {
		return *m.ClientResolution
	}
	return 0
}

type HotupdateAckGameAssetsInfo struct {
	Header           *ProtoHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	AssetHost        *string      `protobuf:"bytes,2,opt,name=assetHost" json:"assetHost,omitempty"`
	Assets           []*AssetInfo `protobuf:"bytes,3,rep,name=assets" json:"assets,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *HotupdateAckGameAssetsInfo) Reset()                    { *m = HotupdateAckGameAssetsInfo{} }
func (m *HotupdateAckGameAssetsInfo) String() string            { return proto.CompactTextString(m) }
func (*HotupdateAckGameAssetsInfo) ProtoMessage()               {}
func (*HotupdateAckGameAssetsInfo) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{7} }

func (m *HotupdateAckGameAssetsInfo) GetHeader() *ProtoHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HotupdateAckGameAssetsInfo) GetAssetHost() string {
	if m != nil && m.AssetHost != nil {
		return *m.AssetHost
	}
	return ""
}

func (m *HotupdateAckGameAssetsInfo) GetAssets() []*AssetInfo {
	if m != nil {
		return m.Assets
	}
	return nil
}

func init() {
	proto.RegisterType((*VersionInfo)(nil), "ddproto.VersionInfo")
	proto.RegisterType((*AssetInfo)(nil), "ddproto.AssetInfo")
	proto.RegisterType((*HotupdateReqVersionInfo)(nil), "ddproto.hotupdate_req_versionInfo")
	proto.RegisterType((*HotupdateAckVersionInfo)(nil), "ddproto.hotupdate_ack_versionInfo")
	proto.RegisterType((*HotupdateReqAssetsInfo)(nil), "ddproto.hotupdate_req_assetsInfo")
	proto.RegisterType((*HotupdateAckAssetsInfo)(nil), "ddproto.hotupdate_ack_assetsInfo")
	proto.RegisterType((*HotupdateReqGameAssetsInfo)(nil), "ddproto.hotupdate_req_gameAssetsInfo")
	proto.RegisterType((*HotupdateAckGameAssetsInfo)(nil), "ddproto.hotupdate_ack_gameAssetsInfo")
}

func init() { proto.RegisterFile("hot_update.proto", fileDescriptor31) }

var fileDescriptor31 = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x54, 0xdf, 0x8a, 0xd3, 0x4e,
	0x14, 0xfe, 0xa5, 0x69, 0xd3, 0xf6, 0xf4, 0xe7, 0x52, 0x67, 0x97, 0x25, 0x5b, 0x16, 0x09, 0x41,
	0xa4, 0x2c, 0x52, 0xb4, 0xe0, 0xb5, 0x54, 0x05, 0xdb, 0x2b, 0xcb, 0xec, 0x9f, 0xdb, 0x32, 0x34,
	0xa7, 0xdb, 0x60, 0x92, 0x89, 0x99, 0xe9, 0x8a, 0x5e, 0xfb, 0x08, 0xa2, 0xef, 0xe1, 0xeb, 0xf8,
	0x16, 0x3e, 0x81, 0xcc, 0x64, 0x12, 0x93, 0x6e, 0x15, 0xdc, 0x75, 0x6f, 0xca, 0x9c, 0xef, 0x7c,
	0x67, 0xf8, 0xf2, 0x7d, 0xa7, 0x03, 0xfd, 0x35, 0x97, 0x8b, 0x4d, 0x1a, 0x30, 0x89, 0xa3, 0x34,
	0xe3, 0x92, 0x93, 0x76, 0x10, 0xe8, 0xc3, 0x60, 0x7f, 0xc9, 0xe3, 0x98, 0x27, 0x8b, 0x65, 0x14,
	0x62, 0x22, 0xf3, 0xee, 0xe0, 0xbe, 0x01, 0x31, 0xd9, 0xc4, 0x39, 0xe4, 0x3f, 0x87, 0xde, 0x05,
	0x66, 0x22, 0xe4, 0xc9, 0x2c, 0x59, 0x71, 0x72, 0x08, 0xce, 0x2a, 0x8c, 0x70, 0x16, 0xb8, 0x96,
	0x67, 0x0d, 0x5b, 0xd4, 0x54, 0xc4, 0x85, 0xb6, 0x3a, 0x5d, 0x60, 0xe6, 0x36, 0x74, 0xa3, 0x28,
	0xfd, 0x1f, 0x16, 0x74, 0x27, 0x42, 0xa0, 0xbc, 0xd9, 0x3c, 0x19, 0x40, 0x47, 0x1d, 0xe7, 0x4c,
	0xae, 0x5d, 0xdb, 0xb3, 0x86, 0x5d, 0x5a, 0xd6, 0x45, 0xef, 0x34, 0xfc, 0x88, 0x6e, 0x53, 0x8f,
	0x95, 0x35, 0xe9, 0x83, 0x1d, 0x07, 0xcf, 0xdc, 0x96, 0x1e, 0x51, 0x47, 0xf2, 0x00, 0x20, 0x14,
	0x2f, 0x79, 0x9c, 0x66, 0x28, 0x84, 0xeb, 0x78, 0xd6, 0xb0, 0x43, 0x2b, 0x88, 0xd2, 0xa6, 0xaa,
	0x00, 0xdd, 0xb6, 0xee, 0x99, 0x8a, 0x3c, 0x05, 0xe7, 0x92, 0xc5, 0x4a, 0x73, 0xc7, 0xb3, 0x86,
	0x7b, 0xe3, 0xa3, 0x91, 0x31, 0x71, 0x54, 0xb1, 0x6b, 0xa1, 0x28, 0xd4, 0x10, 0xfd, 0xef, 0x0d,
	0x38, 0x5a, 0x73, 0x99, 0x5b, 0xbf, 0xc8, 0xf0, 0xdd, 0xe2, 0xaa, 0x62, 0xe2, 0x63, 0x70, 0xd6,
	0xc8, 0x02, 0xcc, 0xb4, 0x09, 0xbd, 0xf1, 0x41, 0x79, 0xe1, 0x5c, 0xfd, 0x4e, 0x75, 0x8f, 0x1a,
	0x0e, 0xf1, 0xa0, 0x97, 0x87, 0x34, 0x49, 0xd3, 0x59, 0x60, 0xec, 0xa9, 0x42, 0xe4, 0x11, 0xec,
	0x05, 0xfc, 0x7d, 0x12, 0x71, 0x16, 0xbc, 0xce, 0x85, 0xda, 0x9e, 0x3d, 0x6c, 0xd1, 0x2d, 0x94,
	0x3c, 0x81, 0x7d, 0x33, 0xa6, 0xf2, 0x10, 0x26, 0x57, 0xe3, 0xdc, 0xae, 0x16, 0x39, 0x81, 0x7e,
	0x0e, 0x53, 0x14, 0x3c, 0xda, 0x48, 0x45, 0x6f, 0x69, 0xfa, 0x35, 0x5c, 0xd9, 0x3b, 0x49, 0xd3,
	0xe2, 0x52, 0x47, 0xb3, 0x2a, 0x08, 0x79, 0x08, 0xf7, 0x12, 0x26, 0xc3, 0x2b, 0x2c, 0x28, 0x6d,
	0x4d, 0xa9, 0x83, 0xc4, 0x87, 0xff, 0xf3, 0x9b, 0xdf, 0x9c, 0x9e, 0x7d, 0x48, 0x51, 0x5b, 0xde,
	0xa2, 0x35, 0xcc, 0xff, 0x5a, 0x73, 0x97, 0x2d, 0xdf, 0xde, 0xc2, 0xdd, 0x11, 0xb4, 0xcd, 0xb0,
	0xdb, 0xf0, 0xec, 0x1a, 0xbd, 0xb2, 0xf7, 0xb4, 0x20, 0x91, 0x31, 0x1c, 0x44, 0x4c, 0x48, 0x14,
	0x5b, 0x26, 0xda, 0x5a, 0xe7, 0xce, 0x9e, 0x4a, 0x70, 0xc5, 0xb3, 0x25, 0x9e, 0x6b, 0xc1, 0xda,
	0xef, 0x0e, 0xad, 0x42, 0xca, 0xbb, 0x0c, 0x23, 0x64, 0x02, 0xcf, 0xd8, 0xa5, 0x71, 0xb8, 0x82,
	0xa8, 0x84, 0x59, 0x9a, 0xbe, 0x32, 0x71, 0x9e, 0x67, 0x91, 0xf6, 0xb7, 0x4b, 0xb7, 0x50, 0xff,
	0x93, 0x05, 0x6e, 0x7d, 0xef, 0x98, 0x56, 0x72, 0x27, 0x6b, 0x67, 0xfe, 0xb3, 0xb3, 0x40, 0x98,
	0x7d, 0x2b, 0x4a, 0xff, 0x73, 0xa3, 0x2a, 0x43, 0x05, 0x74, 0x63, 0x19, 0xc7, 0xd0, 0xd5, 0xb3,
	0x53, 0x2e, 0xa4, 0x16, 0xd1, 0xa5, 0xbf, 0x00, 0x72, 0x02, 0x4e, 0x7e, 0xb3, 0x56, 0xd0, 0x1b,
	0x93, 0xf2, 0xae, 0xf2, 0xc9, 0xa1, 0x86, 0xf1, 0xdb, 0xe4, 0x9a, 0x7f, 0x48, 0xee, 0x6f, 0xf6,
	0x7f, 0xcb, 0x30, 0xe7, 0x9a, 0x61, 0xfe, 0x37, 0x0b, 0x8e, 0xeb, 0xe9, 0xa8, 0xe7, 0x62, 0x72,
	0x77, 0x09, 0x1d, 0x96, 0x2f, 0x57, 0xbe, 0x9e, 0xa6, 0xda, 0xf9, 0x59, 0xcd, 0xdd, 0x9f, 0xe5,
	0x7f, 0xa9, 0x89, 0x56, 0x59, 0xde, 0x4a, 0xf4, 0x3f, 0xcb, 0xf3, 0x45, 0x63, 0x6a, 0xcf, 0xff,
	0x9b, 0x5b, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xbc, 0x2f, 0x20, 0xe5, 0x06, 0x00, 0x00,
}
