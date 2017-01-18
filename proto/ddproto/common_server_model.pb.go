// Code generated by protoc-gen-go.
// source: common_server_model.proto
// DO NOT EDIT!

package ddproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 用户的资
type User struct {
	RoomCard            *int64  `protobuf:"varint,1,opt,name=RoomCard" json:"RoomCard,omitempty"`
	Pwd                 *string `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
	Coin                *int64  `protobuf:"varint,3,opt,name=coin" json:"coin,omitempty"`
	Id                  *uint32 `protobuf:"varint,4,opt,name=id" json:"id,omitempty"`
	NickName            *string `protobuf:"bytes,5,opt,name=nickName" json:"nickName,omitempty"`
	Scores              *int32  `protobuf:"varint,6,opt,name=scores" json:"scores,omitempty"`
	LastDrawLotteryTime *string `protobuf:"bytes,7,opt,name=lastDrawLotteryTime" json:"lastDrawLotteryTime,omitempty"`
	LastSignTime        *string `protobuf:"bytes,8,opt,name=lastSignTime" json:"lastSignTime,omitempty"`
	SignTotalDays       *int32  `protobuf:"varint,9,opt,name=signTotalDays" json:"signTotalDays,omitempty"`
	SignContinuousDays  *int32  `protobuf:"varint,10,opt,name=signContinuousDays" json:"signContinuousDays,omitempty"`
	Diamond             *int64  `protobuf:"varint,11,opt,name=Diamond" json:"Diamond,omitempty"`
	Diamond2            *int64  `protobuf:"varint,12,opt,name=Diamond2" json:"Diamond2,omitempty"`
	OpenId              *string `protobuf:"bytes,13,opt,name=openId" json:"openId,omitempty"`
	UnionId             *string `protobuf:"bytes,14,opt,name=UnionId" json:"UnionId,omitempty"`
	HeadUrl             *string `protobuf:"bytes,15,opt,name=headUrl" json:"headUrl,omitempty"`
	City                *string `protobuf:"bytes,16,opt,name=city" json:"city,omitempty"`
	Sex                 *int32  `protobuf:"varint,17,opt,name=sex" json:"sex,omitempty"`
	RobotType           *int32  `protobuf:"varint,18,opt,name=robotType" json:"robotType,omitempty"`
	XXX_unrecognized    []byte  `json:"-"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *User) GetRoomCard() int64 {
	if m != nil && m.RoomCard != nil {
		return *m.RoomCard
	}
	return 0
}

func (m *User) GetPwd() string {
	if m != nil && m.Pwd != nil {
		return *m.Pwd
	}
	return ""
}

func (m *User) GetCoin() int64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *User) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *User) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *User) GetScores() int32 {
	if m != nil && m.Scores != nil {
		return *m.Scores
	}
	return 0
}

func (m *User) GetLastDrawLotteryTime() string {
	if m != nil && m.LastDrawLotteryTime != nil {
		return *m.LastDrawLotteryTime
	}
	return ""
}

func (m *User) GetLastSignTime() string {
	if m != nil && m.LastSignTime != nil {
		return *m.LastSignTime
	}
	return ""
}

func (m *User) GetSignTotalDays() int32 {
	if m != nil && m.SignTotalDays != nil {
		return *m.SignTotalDays
	}
	return 0
}

func (m *User) GetSignContinuousDays() int32 {
	if m != nil && m.SignContinuousDays != nil {
		return *m.SignContinuousDays
	}
	return 0
}

func (m *User) GetDiamond() int64 {
	if m != nil && m.Diamond != nil {
		return *m.Diamond
	}
	return 0
}

func (m *User) GetDiamond2() int64 {
	if m != nil && m.Diamond2 != nil {
		return *m.Diamond2
	}
	return 0
}

func (m *User) GetOpenId() string {
	if m != nil && m.OpenId != nil {
		return *m.OpenId
	}
	return ""
}

func (m *User) GetUnionId() string {
	if m != nil && m.UnionId != nil {
		return *m.UnionId
	}
	return ""
}

func (m *User) GetHeadUrl() string {
	if m != nil && m.HeadUrl != nil {
		return *m.HeadUrl
	}
	return ""
}

func (m *User) GetCity() string {
	if m != nil && m.City != nil {
		return *m.City
	}
	return ""
}

func (m *User) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *User) GetRobotType() int32 {
	if m != nil && m.RobotType != nil {
		return *m.RobotType
	}
	return 0
}

// 通知公告
type TNotice struct {
	Id               *int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	NoticeType       *int32   `protobuf:"varint,2,opt,name=noticeType" json:"noticeType,omitempty"`
	NoticeTitle      *string  `protobuf:"bytes,3,opt,name=noticeTitle" json:"noticeTitle,omitempty"`
	NoticeContent    *string  `protobuf:"bytes,4,opt,name=noticeContent" json:"noticeContent,omitempty"`
	NoticeMemo       *string  `protobuf:"bytes,5,opt,name=noticeMemo" json:"noticeMemo,omitempty"`
	Noticefileds     []string `protobuf:"bytes,6,rep,name=noticefileds" json:"noticefileds,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TNotice) Reset()                    { *m = TNotice{} }
func (m *TNotice) String() string            { return proto.CompactTextString(m) }
func (*TNotice) ProtoMessage()               {}
func (*TNotice) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *TNotice) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TNotice) GetNoticeType() int32 {
	if m != nil && m.NoticeType != nil {
		return *m.NoticeType
	}
	return 0
}

func (m *TNotice) GetNoticeTitle() string {
	if m != nil && m.NoticeTitle != nil {
		return *m.NoticeTitle
	}
	return ""
}

func (m *TNotice) GetNoticeContent() string {
	if m != nil && m.NoticeContent != nil {
		return *m.NoticeContent
	}
	return ""
}

func (m *TNotice) GetNoticeMemo() string {
	if m != nil && m.NoticeMemo != nil {
		return *m.NoticeMemo
	}
	return ""
}

func (m *TNotice) GetNoticefileds() []string {
	if m != nil {
		return m.Noticefileds
	}
	return nil
}

// 玩家游戏场次统计
type TGameCounts struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UserId           *uint32 `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	DdzCount         *int32  `protobuf:"varint,3,opt,name=ddzCount" json:"ddzCount,omitempty"`
	DdzWinCount      *int32  `protobuf:"varint,4,opt,name=ddzWinCount" json:"ddzWinCount,omitempty"`
	MjCount          *int32  `protobuf:"varint,5,opt,name=mjCount" json:"mjCount,omitempty"`
	MjWinCount       *int32  `protobuf:"varint,6,opt,name=mjWinCount" json:"mjWinCount,omitempty"`
	ThCount          *int32  `protobuf:"varint,7,opt,name=thCount" json:"thCount,omitempty"`
	ThWinWinCount    *int32  `protobuf:"varint,8,opt,name=thWinWinCount" json:"thWinWinCount,omitempty"`
	ZjhCount         *int32  `protobuf:"varint,9,opt,name=zjhCount" json:"zjhCount,omitempty"`
	ZjhWinCount      *int32  `protobuf:"varint,10,opt,name=zjhWinCount" json:"zjhWinCount,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TGameCounts) Reset()                    { *m = TGameCounts{} }
func (m *TGameCounts) String() string            { return proto.CompactTextString(m) }
func (*TGameCounts) ProtoMessage()               {}
func (*TGameCounts) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *TGameCounts) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TGameCounts) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *TGameCounts) GetDdzCount() int32 {
	if m != nil && m.DdzCount != nil {
		return *m.DdzCount
	}
	return 0
}

func (m *TGameCounts) GetDdzWinCount() int32 {
	if m != nil && m.DdzWinCount != nil {
		return *m.DdzWinCount
	}
	return 0
}

func (m *TGameCounts) GetMjCount() int32 {
	if m != nil && m.MjCount != nil {
		return *m.MjCount
	}
	return 0
}

func (m *TGameCounts) GetMjWinCount() int32 {
	if m != nil && m.MjWinCount != nil {
		return *m.MjWinCount
	}
	return 0
}

func (m *TGameCounts) GetThCount() int32 {
	if m != nil && m.ThCount != nil {
		return *m.ThCount
	}
	return 0
}

func (m *TGameCounts) GetThWinWinCount() int32 {
	if m != nil && m.ThWinWinCount != nil {
		return *m.ThWinWinCount
	}
	return 0
}

func (m *TGameCounts) GetZjhCount() int32 {
	if m != nil && m.ZjhCount != nil {
		return *m.ZjhCount
	}
	return 0
}

func (m *TGameCounts) GetZjhWinCount() int32 {
	if m != nil && m.ZjhWinCount != nil {
		return *m.ZjhWinCount
	}
	return 0
}

// 玩家的任务
type TUserTask struct {
	Id               *int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UserId           *uint32 `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	TaskId           *int32  `protobuf:"varint,3,opt,name=taskId" json:"taskId,omitempty"`
	AwardId          *int32  `protobuf:"varint,4,opt,name=awardId" json:"awardId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TUserTask) Reset()                    { *m = TUserTask{} }
func (m *TUserTask) String() string            { return proto.CompactTextString(m) }
func (*TUserTask) ProtoMessage()               {}
func (*TUserTask) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *TUserTask) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *TUserTask) GetUserId() uint32 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *TUserTask) GetTaskId() int32 {
	if m != nil && m.TaskId != nil {
		return *m.TaskId
	}
	return 0
}

func (m *TUserTask) GetAwardId() int32 {
	if m != nil && m.AwardId != nil {
		return *m.AwardId
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "ddproto.User")
	proto.RegisterType((*TNotice)(nil), "ddproto.TNotice")
	proto.RegisterType((*TGameCounts)(nil), "ddproto.TGameCounts")
	proto.RegisterType((*TUserTask)(nil), "ddproto.TUserTask")
}


var fileDescriptor7 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0x13, 0x3d,
	0x14, 0xd5, 0xe4, 0xb7, 0xe3, 0x34, 0xfd, 0xfa, 0x19, 0x29, 0x32, 0x08, 0xa1, 0x28, 0x62, 0x91,
	0x55, 0x85, 0x78, 0x85, 0x44, 0x42, 0x91, 0xa0, 0x8b, 0x61, 0xa2, 0x2e, 0x2b, 0x13, 0x1b, 0xe2,
	0x64, 0x6c, 0x47, 0xb6, 0x43, 0x48, 0x5e, 0x89, 0xe7, 0xe0, 0xb5, 0x10, 0xba, 0xd7, 0xce, 0x74,
	0x22, 0xba, 0x60, 0x77, 0xcf, 0x39, 0xf7, 0x7a, 0xee, 0xcf, 0x19, 0xf2, 0x72, 0x65, 0xb5, 0xb6,
	0xe6, 0xd1, 0x4b, 0xf7, 0x5d, 0xba, 0x47, 0x6d, 0x85, 0xac, 0xee, 0x76, 0xce, 0x06, 0x4b, 0xfb,
	0x42, 0x60, 0x30, 0xf9, 0xdd, 0x26, 0x9d, 0xa5, 0x97, 0x8e, 0xbe, 0x22, 0x57, 0x85, 0xb5, 0x7a,
	0xc6, 0x9d, 0x60, 0xd9, 0x38, 0x9b, 0xb6, 0x8b, 0x1a, 0xd3, 0x5b, 0xd2, 0xde, 0x1d, 0x04, 0x6b,
	0x8d, 0xb3, 0x69, 0x5e, 0x40, 0x48, 0x29, 0xe9, 0xac, 0xac, 0x32, 0xac, 0x8d, 0x99, 0x18, 0xd3,
	0x1b, 0xd2, 0x52, 0x82, 0x75, 0xc6, 0xd9, 0x74, 0x58, 0xb4, 0x94, 0x80, 0x17, 0x8d, 0x5a, 0x6d,
	0xef, 0xb9, 0x96, 0xac, 0x8b, 0xa5, 0x35, 0xa6, 0x23, 0xd2, 0xf3, 0x2b, 0xeb, 0xa4, 0x67, 0xbd,
	0x71, 0x36, 0xed, 0x16, 0x09, 0xd1, 0x77, 0xe4, 0x45, 0xc5, 0x7d, 0x98, 0x3b, 0x7e, 0xf8, 0x68,
	0x43, 0x90, 0xee, 0x58, 0x2a, 0x2d, 0x59, 0x1f, 0xcb, 0x9f, 0x93, 0xe8, 0x84, 0x5c, 0x03, 0xfd,
	0x59, 0x7d, 0x33, 0x98, 0x7a, 0x85, 0xa9, 0x17, 0x1c, 0x7d, 0x4b, 0x86, 0x1e, 0x62, 0x1b, 0x78,
	0x35, 0xe7, 0x47, 0xcf, 0x72, 0xfc, 0xe8, 0x25, 0x49, 0xef, 0x08, 0x05, 0x62, 0x66, 0x4d, 0x50,
	0x66, 0x6f, 0xf7, 0x1e, 0x53, 0x09, 0xa6, 0x3e, 0xa3, 0x50, 0x46, 0xfa, 0x73, 0xc5, 0xb5, 0x35,
	0x82, 0x0d, 0x70, 0x0d, 0x67, 0x08, 0x93, 0xa7, 0xf0, 0x3d, 0xbb, 0x8e, 0xbb, 0x3c, 0x63, 0x98,
	0xdc, 0xee, 0xa4, 0x59, 0x08, 0x36, 0xc4, 0x4e, 0x13, 0x82, 0xd7, 0x96, 0x46, 0x59, 0x10, 0x6e,
	0x50, 0x38, 0x43, 0x50, 0xd6, 0x92, 0x8b, 0xa5, 0xab, 0xd8, 0x7f, 0x51, 0x49, 0x10, 0xaf, 0xa0,
	0xc2, 0x91, 0xdd, 0x22, 0x8d, 0x31, 0xdc, 0xca, 0xcb, 0x1f, 0xec, 0x7f, 0x6c, 0x1b, 0x42, 0xfa,
	0x9a, 0xe4, 0xce, 0x7e, 0xb1, 0xa1, 0x3c, 0xee, 0x24, 0xa3, 0xc8, 0x3f, 0x11, 0x93, 0x5f, 0x19,
	0xe9, 0x97, 0xf7, 0x36, 0xa8, 0x95, 0x4c, 0x17, 0xcc, 0x30, 0x05, 0x2e, 0xf8, 0x86, 0x10, 0x83,
	0x0a, 0x96, 0xb6, 0x90, 0x6f, 0x30, 0x74, 0x4c, 0x06, 0x09, 0xa9, 0x50, 0x49, 0x34, 0x43, 0x5e,
	0x34, 0x29, 0xd8, 0x7c, 0x84, 0xb0, 0x3b, 0x69, 0x02, 0xda, 0x23, 0x2f, 0x2e, 0xc9, 0xa7, 0xef,
	0x7c, 0x92, 0xda, 0x26, 0xaf, 0x34, 0x18, 0xb8, 0x71, 0x44, 0x5f, 0x55, 0x25, 0x05, 0x78, 0xa6,
	0x0d, 0x37, 0x6e, 0x72, 0x93, 0x9f, 0x2d, 0x32, 0x28, 0x3f, 0x70, 0x2d, 0x67, 0x76, 0x6f, 0x82,
	0xff, 0x6b, 0x96, 0x11, 0xe9, 0xed, 0xbd, 0x74, 0x8b, 0x68, 0xe3, 0x61, 0x91, 0x10, 0xdc, 0x4a,
	0x88, 0x13, 0x16, 0xe1, 0x00, 0xdd, 0xa2, 0xc6, 0x30, 0x9f, 0x10, 0xa7, 0x07, 0x65, 0xa2, 0xdc,
	0x41, 0xb9, 0x49, 0xc1, 0x6d, 0xf4, 0x26, 0xaa, 0x5d, 0x54, 0xcf, 0x10, 0x66, 0xd2, 0x9b, 0xba,
	0x34, 0xba, 0xbc, 0xc1, 0x40, 0x65, 0x58, 0x47, 0xb1, 0x1f, 0x2b, 0x13, 0x84, 0x9d, 0x85, 0xf5,
	0x83, 0x32, 0x75, 0xf1, 0x55, 0x74, 0xeb, 0x05, 0x09, 0x7d, 0x9f, 0x36, 0xe9, 0x81, 0x68, 0xe7,
	0x1a, 0x43, 0xdf, 0xa7, 0xcd, 0xba, 0xae, 0x8f, 0x16, 0x6e, 0x52, 0x13, 0x49, 0xf2, 0x12, 0x7e,
	0xfb, 0x92, 0xfb, 0xed, 0x3f, 0xaf, 0x6a, 0x44, 0x7a, 0x81, 0xfb, 0xed, 0x42, 0xa4, 0x45, 0x25,
	0x04, 0xa3, 0xf0, 0x03, 0x77, 0x62, 0x21, 0xd2, 0x8a, 0xce, 0xf0, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xa6, 0x0b, 0x89, 0x0c, 0x82, 0x04, 0x00, 0x00,
}
