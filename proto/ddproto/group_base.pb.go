// Code generated by protoc-gen-go. DO NOT EDIT.
// source: group_base.proto

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

// Ignoring public import of common_enum_reg from common_client.proto

// Ignoring public import of common_enum_os_type from common_client.proto

// Ignoring public import of common_enum_pokerColor from common_client.proto

// 消息类型
type GroupMsgType int32

const (
	GroupMsgType_group_msg_type_TXT        GroupMsgType = 1
	GroupMsgType_group_msg_type_MemberIn   GroupMsgType = 2
	GroupMsgType_group_msg_type_MemberOut  GroupMsgType = 3
	GroupMsgType_group_msg_type_CreateDesk GroupMsgType = 4
	GroupMsgType_group_msg_type_SysSuccess GroupMsgType = 5
	GroupMsgType_group_msg_type_SysFail    GroupMsgType = 6
	GroupMsgType_group_msg_type_Voice      GroupMsgType = 7
	GroupMsgType_group_msg_type_Image      GroupMsgType = 8
)

var GroupMsgType_name = map[int32]string{
	1: "group_msg_type_TXT",
	2: "group_msg_type_MemberIn",
	3: "group_msg_type_MemberOut",
	4: "group_msg_type_CreateDesk",
	5: "group_msg_type_SysSuccess",
	6: "group_msg_type_SysFail",
	7: "group_msg_type_Voice",
	8: "group_msg_type_Image",
}
var GroupMsgType_value = map[string]int32{
	"group_msg_type_TXT":        1,
	"group_msg_type_MemberIn":   2,
	"group_msg_type_MemberOut":  3,
	"group_msg_type_CreateDesk": 4,
	"group_msg_type_SysSuccess": 5,
	"group_msg_type_SysFail":    6,
	"group_msg_type_Voice":      7,
	"group_msg_type_Image":      8,
}

func (x GroupMsgType) Enum() *GroupMsgType {
	p := new(GroupMsgType)
	*p = x
	return p
}
func (x GroupMsgType) String() string {
	return proto.EnumName(GroupMsgType_name, int32(x))
}
func (x *GroupMsgType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(GroupMsgType_value, data, "GroupMsgType")
	if err != nil {
		return err
	}
	*x = GroupMsgType(value)
	return nil
}
func (GroupMsgType) EnumDescriptor() ([]byte, []int) { return fileDescriptor26, []int{0} }

// 俱乐部信息
type GroupInfo struct {
	Id               *int32             `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name             *string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Info             *string            `protobuf:"bytes,3,opt,name=info" json:"info,omitempty"`
	Owner            *uint32            `protobuf:"varint,4,opt,name=owner" json:"owner,omitempty"`
	Members          []*GroupMemberInfo `protobuf:"bytes,6,rep,name=members" json:"members,omitempty"`
	GameOpts         []*GameOption      `protobuf:"bytes,7,rep,name=gameOpts" json:"gameOpts,omitempty"`
	SyncId           *int32             `protobuf:"varint,8,opt,name=syncId" json:"syncId,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *GroupInfo) Reset()                    { *m = GroupInfo{} }
func (m *GroupInfo) String() string            { return proto.CompactTextString(m) }
func (*GroupInfo) ProtoMessage()               {}
func (*GroupInfo) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{0} }

func (m *GroupInfo) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *GroupInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *GroupInfo) GetInfo() string {
	if m != nil && m.Info != nil {
		return *m.Info
	}
	return ""
}

func (m *GroupInfo) GetOwner() uint32 {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return 0
}

func (m *GroupInfo) GetMembers() []*GroupMemberInfo {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *GroupInfo) GetGameOpts() []*GameOption {
	if m != nil {
		return m.GameOpts
	}
	return nil
}

func (m *GroupInfo) GetSyncId() int32 {
	if m != nil && m.SyncId != nil {
		return *m.SyncId
	}
	return 0
}

// 成员信息
type GroupMemberInfo struct {
	Uid              *uint32 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	NickName         *string `protobuf:"bytes,2,opt,name=nickName" json:"nickName,omitempty"`
	Remark           *string `protobuf:"bytes,3,opt,name=remark" json:"remark,omitempty"`
	HeadImg          *string `protobuf:"bytes,4,opt,name=headImg" json:"headImg,omitempty"`
	OpenId           *string `protobuf:"bytes,5,opt,name=openId" json:"openId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GroupMemberInfo) Reset()                    { *m = GroupMemberInfo{} }
func (m *GroupMemberInfo) String() string            { return proto.CompactTextString(m) }
func (*GroupMemberInfo) ProtoMessage()               {}
func (*GroupMemberInfo) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{1} }

func (m *GroupMemberInfo) GetUid() uint32 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *GroupMemberInfo) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *GroupMemberInfo) GetRemark() string {
	if m != nil && m.Remark != nil {
		return *m.Remark
	}
	return ""
}

func (m *GroupMemberInfo) GetHeadImg() string {
	if m != nil && m.HeadImg != nil {
		return *m.HeadImg
	}
	return ""
}

func (m *GroupMemberInfo) GetOpenId() string {
	if m != nil && m.OpenId != nil {
		return *m.OpenId
	}
	return ""
}

// 游戏配置
type GameOption struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	GameId           *int32  `protobuf:"varint,2,opt,name=gameId" json:"gameId,omitempty"`
	Remark           *string `protobuf:"bytes,3,opt,name=remark" json:"remark,omitempty"`
	Option           *string `protobuf:"bytes,4,opt,name=option" json:"option,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GameOption) Reset()                    { *m = GameOption{} }
func (m *GameOption) String() string            { return proto.CompactTextString(m) }
func (*GameOption) ProtoMessage()               {}
func (*GameOption) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{2} }

func (m *GameOption) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *GameOption) GetGameId() int32 {
	if m != nil && m.GameId != nil {
		return *m.GameId
	}
	return 0
}

func (m *GameOption) GetRemark() string {
	if m != nil && m.Remark != nil {
		return *m.Remark
	}
	return ""
}

func (m *GameOption) GetOption() string {
	if m != nil && m.Option != nil {
		return *m.Option
	}
	return ""
}

// 群消息结构
type GroupMsgItem struct {
	FromUid          *uint32            `protobuf:"varint,1,opt,name=fromUid" json:"fromUid,omitempty"`
	ToGroup          *int32             `protobuf:"varint,2,opt,name=toGroup" json:"toGroup,omitempty"`
	SendTime         *int64             `protobuf:"varint,3,opt,name=sendTime" json:"sendTime,omitempty"`
	MsgType          *GroupMsgType      `protobuf:"varint,4,opt,name=msgType,enum=ddproto.GroupMsgType" json:"msgType,omitempty"`
	TxtContent       *string            `protobuf:"bytes,5,opt,name=txtContent" json:"txtContent,omitempty"`
	DeskContent      *CommonDeskByAgent `protobuf:"bytes,7,opt,name=deskContent" json:"deskContent,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *GroupMsgItem) Reset()                    { *m = GroupMsgItem{} }
func (m *GroupMsgItem) String() string            { return proto.CompactTextString(m) }
func (*GroupMsgItem) ProtoMessage()               {}
func (*GroupMsgItem) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{3} }

func (m *GroupMsgItem) GetFromUid() uint32 {
	if m != nil && m.FromUid != nil {
		return *m.FromUid
	}
	return 0
}

func (m *GroupMsgItem) GetToGroup() int32 {
	if m != nil && m.ToGroup != nil {
		return *m.ToGroup
	}
	return 0
}

func (m *GroupMsgItem) GetSendTime() int64 {
	if m != nil && m.SendTime != nil {
		return *m.SendTime
	}
	return 0
}

func (m *GroupMsgItem) GetMsgType() GroupMsgType {
	if m != nil && m.MsgType != nil {
		return *m.MsgType
	}
	return GroupMsgType_group_msg_type_TXT
}

func (m *GroupMsgItem) GetTxtContent() string {
	if m != nil && m.TxtContent != nil {
		return *m.TxtContent
	}
	return ""
}

func (m *GroupMsgItem) GetDeskContent() *CommonDeskByAgent {
	if m != nil {
		return m.DeskContent
	}
	return nil
}

func init() {
	proto.RegisterType((*GroupInfo)(nil), "ddproto.group_info")
	proto.RegisterType((*GroupMemberInfo)(nil), "ddproto.group_member_info")
	proto.RegisterType((*GameOption)(nil), "ddproto.game_option")
	proto.RegisterType((*GroupMsgItem)(nil), "ddproto.group_msg_item")
	proto.RegisterEnum("ddproto.GroupMsgType", GroupMsgType_name, GroupMsgType_value)
}

func init() { proto.RegisterFile("group_base.proto", fileDescriptor26) }

var fileDescriptor26 = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcd, 0x8e, 0xd3, 0x30,
	0x18, 0x24, 0xfd, 0x4b, 0xf7, 0xab, 0xb6, 0x0a, 0xa6, 0xea, 0x9a, 0xc2, 0xa2, 0xaa, 0xa7, 0x8a,
	0x43, 0x05, 0x15, 0x77, 0x0e, 0x8b, 0x40, 0x39, 0xc0, 0x22, 0x6f, 0x41, 0xdc, 0xa2, 0x34, 0xf9,
	0x36, 0x44, 0x5d, 0xdb, 0x51, 0xec, 0x0a, 0xfa, 0x0c, 0xbc, 0x21, 0x6f, 0x81, 0xc4, 0x03, 0x20,
	0x3b, 0x4e, 0xb6, 0xdb, 0x96, 0x9b, 0xe7, 0x9b, 0x19, 0xcf, 0x37, 0x4e, 0x20, 0xc8, 0x4a, 0xb9,
	0x2d, 0xa2, 0x75, 0xac, 0x70, 0x51, 0x94, 0x52, 0x4b, 0xe2, 0xa7, 0xa9, 0x3d, 0x4c, 0x9e, 0x24,
	0x92, 0x73, 0x29, 0xa2, 0xe4, 0x2e, 0x47, 0xa1, 0x2b, 0x76, 0xf6, 0xdb, 0x03, 0xa8, 0x2c, 0xb9,
	0xb8, 0x95, 0x64, 0x08, 0xad, 0x3c, 0xa5, 0xde, 0xd4, 0x9b, 0x77, 0x59, 0x2b, 0x4f, 0x09, 0x81,
	0x8e, 0x88, 0x39, 0xd2, 0xd6, 0xd4, 0x9b, 0x9f, 0x31, 0x7b, 0x36, 0x33, 0xa3, 0xa5, 0xed, 0x6a,
	0x66, 0x7d, 0x23, 0xe8, 0xca, 0x1f, 0x02, 0x4b, 0xda, 0x99, 0x7a, 0xf3, 0x73, 0x56, 0x01, 0xf2,
	0x06, 0x7c, 0x8e, 0x7c, 0x8d, 0xa5, 0xa2, 0xbd, 0x69, 0x7b, 0x3e, 0x58, 0x4e, 0x16, 0x6e, 0x99,
	0x45, 0x95, 0x59, 0xb1, 0x36, 0x9a, 0xd5, 0x52, 0xf2, 0x0a, 0xfa, 0x59, 0xcc, 0xf1, 0xba, 0xd0,
	0x8a, 0xfa, 0xd6, 0x36, 0xba, 0xb7, 0xc5, 0x1c, 0x23, 0x59, 0xe8, 0x5c, 0x0a, 0xd6, 0xa8, 0xc8,
	0x18, 0x7a, 0x6a, 0x27, 0x92, 0x30, 0xa5, 0x7d, 0xbb, 0xb9, 0x43, 0xb3, 0x5f, 0x1e, 0x3c, 0x3e,
	0x0a, 0x22, 0x01, 0xb4, 0xb7, 0xae, 0xe4, 0x39, 0x33, 0x47, 0x32, 0x81, 0xbe, 0xc8, 0x93, 0xcd,
	0xa7, 0xfb, 0xa6, 0x0d, 0x36, 0x77, 0x97, 0xc8, 0xe3, 0x72, 0xe3, 0xfa, 0x3a, 0x44, 0x28, 0xf8,
	0xdf, 0x31, 0x4e, 0x43, 0x9e, 0xd9, 0xce, 0x67, 0xac, 0x86, 0xc6, 0x21, 0x0b, 0x14, 0x61, 0x4a,
	0xbb, 0x95, 0xa3, 0x42, 0x33, 0x84, 0xc1, 0xde, 0xfa, 0x47, 0x4f, 0x3d, 0x86, 0x9e, 0xa1, 0xc3,
	0xd4, 0xae, 0xd0, 0x65, 0x0e, 0xfd, 0x77, 0x01, 0x1b, 0x63, 0x6e, 0x72, 0xf9, 0x0e, 0xcd, 0xfe,
	0x7a, 0x30, 0x74, 0xa5, 0x55, 0x16, 0xe5, 0x1a, 0xb9, 0xd9, 0xf5, 0xb6, 0x94, 0xfc, 0x4b, 0xd3,
	0xba, 0x86, 0x86, 0xd1, 0xf2, 0x83, 0x51, 0xbb, 0xd4, 0x1a, 0x9a, 0x37, 0x51, 0x28, 0xd2, 0x55,
	0xce, 0xd1, 0x06, 0xb7, 0x59, 0x83, 0xc9, 0x6b, 0xf0, 0xb9, 0xca, 0x56, 0xbb, 0x02, 0x6d, 0xf6,
	0x70, 0x79, 0x71, 0xf8, 0x5d, 0x55, 0x16, 0xe9, 0x5d, 0x81, 0xac, 0xd6, 0x91, 0x17, 0x00, 0xfa,
	0xa7, 0xbe, 0x92, 0x42, 0xa3, 0xd0, 0xee, 0x61, 0xf6, 0x26, 0xe4, 0x2d, 0x0c, 0x52, 0x54, 0x9b,
	0x5a, 0xe0, 0x4f, 0xbd, 0xf9, 0x60, 0x79, 0xd9, 0x5c, 0xeb, 0x7e, 0x5d, 0x23, 0x89, 0xd6, 0xbb,
	0x28, 0xce, 0x50, 0x68, 0xb6, 0xef, 0x78, 0xf9, 0xe7, 0x41, 0x6d, 0x13, 0x4e, 0xc6, 0x40, 0x1e,
	0x4e, 0xa2, 0xd5, 0xb7, 0x55, 0xe0, 0x91, 0x67, 0x70, 0x71, 0x30, 0xff, 0x68, 0x7f, 0x8f, 0x50,
	0x04, 0x2d, 0xf2, 0x1c, 0xe8, 0x49, 0xf2, 0x7a, 0xab, 0x83, 0x36, 0xb9, 0x84, 0xa7, 0x07, 0xec,
	0x55, 0x89, 0xb1, 0xc6, 0x77, 0xa8, 0x36, 0x41, 0xe7, 0x04, 0x7d, 0xb3, 0x53, 0x37, 0xdb, 0x24,
	0x41, 0xa5, 0x82, 0x2e, 0x99, 0xc0, 0xf8, 0x98, 0x7e, 0x1f, 0xe7, 0x77, 0x41, 0x8f, 0x50, 0x18,
	0x1d, 0x70, 0x5f, 0x65, 0x9e, 0x60, 0xe0, 0x9f, 0x60, 0x42, 0x1e, 0x67, 0x18, 0xf4, 0x3f, 0x3f,
	0xfa, 0x17, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x55, 0x67, 0xe4, 0xef, 0x03, 0x00, 0x00,
}
