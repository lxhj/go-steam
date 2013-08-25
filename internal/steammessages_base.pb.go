// Code generated by protoc-gen-go.
// source: steammessages_base.proto
// DO NOT EDIT!

package internal

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"
import google_protobuf "code.google.com/p/goprotobuf/protoc-gen-go/descriptor"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type CMsgProtoBufHeader struct {
	Steamid          *uint64 `protobuf:"fixed64,1,opt,name=steamid" json:"steamid,omitempty"`
	ClientSessionid  *int32  `protobuf:"varint,2,opt,name=client_sessionid" json:"client_sessionid,omitempty"`
	RoutingAppid     *uint32 `protobuf:"varint,3,opt,name=routing_appid" json:"routing_appid,omitempty"`
	JobidSource      *uint64 `protobuf:"fixed64,10,opt,name=jobid_source,def=18446744073709551615" json:"jobid_source,omitempty"`
	JobidTarget      *uint64 `protobuf:"fixed64,11,opt,name=jobid_target,def=18446744073709551615" json:"jobid_target,omitempty"`
	TargetJobName    *string `protobuf:"bytes,12,opt,name=target_job_name" json:"target_job_name,omitempty"`
	Eresult          *int32  `protobuf:"varint,13,opt,name=eresult,def=2" json:"eresult,omitempty"`
	ErrorMessage     *string `protobuf:"bytes,14,opt,name=error_message" json:"error_message,omitempty"`
	Ip               *uint32 `protobuf:"varint,15,opt,name=ip" json:"ip,omitempty"`
	AuthAccountFlags *uint32 `protobuf:"varint,16,opt,name=auth_account_flags" json:"auth_account_flags,omitempty"`
	TransportError   *int32  `protobuf:"varint,17,opt,name=transport_error,def=1" json:"transport_error,omitempty"`
	Messageid        *uint64 `protobuf:"varint,18,opt,name=messageid,def=18446744073709551615" json:"messageid,omitempty"`
	PublisherGroupId *uint32 `protobuf:"varint,19,opt,name=publisher_group_id" json:"publisher_group_id,omitempty"`
	Sysid            *uint32 `protobuf:"varint,20,opt,name=sysid" json:"sysid,omitempty"`
	TraceTag         *uint64 `protobuf:"varint,21,opt,name=trace_tag" json:"trace_tag,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgProtoBufHeader) Reset()         { *m = CMsgProtoBufHeader{} }
func (m *CMsgProtoBufHeader) String() string { return proto.CompactTextString(m) }
func (*CMsgProtoBufHeader) ProtoMessage()    {}

const Default_CMsgProtoBufHeader_JobidSource uint64 = 18446744073709551615
const Default_CMsgProtoBufHeader_JobidTarget uint64 = 18446744073709551615
const Default_CMsgProtoBufHeader_Eresult int32 = 2
const Default_CMsgProtoBufHeader_TransportError int32 = 1
const Default_CMsgProtoBufHeader_Messageid uint64 = 18446744073709551615

func (m *CMsgProtoBufHeader) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

func (m *CMsgProtoBufHeader) GetClientSessionid() int32 {
	if m != nil && m.ClientSessionid != nil {
		return *m.ClientSessionid
	}
	return 0
}

func (m *CMsgProtoBufHeader) GetRoutingAppid() uint32 {
	if m != nil && m.RoutingAppid != nil {
		return *m.RoutingAppid
	}
	return 0
}

func (m *CMsgProtoBufHeader) GetJobidSource() uint64 {
	if m != nil && m.JobidSource != nil {
		return *m.JobidSource
	}
	return Default_CMsgProtoBufHeader_JobidSource
}

func (m *CMsgProtoBufHeader) GetJobidTarget() uint64 {
	if m != nil && m.JobidTarget != nil {
		return *m.JobidTarget
	}
	return Default_CMsgProtoBufHeader_JobidTarget
}

func (m *CMsgProtoBufHeader) GetTargetJobName() string {
	if m != nil && m.TargetJobName != nil {
		return *m.TargetJobName
	}
	return ""
}

func (m *CMsgProtoBufHeader) GetEresult() int32 {
	if m != nil && m.Eresult != nil {
		return *m.Eresult
	}
	return Default_CMsgProtoBufHeader_Eresult
}

func (m *CMsgProtoBufHeader) GetErrorMessage() string {
	if m != nil && m.ErrorMessage != nil {
		return *m.ErrorMessage
	}
	return ""
}

func (m *CMsgProtoBufHeader) GetIp() uint32 {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return 0
}

func (m *CMsgProtoBufHeader) GetAuthAccountFlags() uint32 {
	if m != nil && m.AuthAccountFlags != nil {
		return *m.AuthAccountFlags
	}
	return 0
}

func (m *CMsgProtoBufHeader) GetTransportError() int32 {
	if m != nil && m.TransportError != nil {
		return *m.TransportError
	}
	return Default_CMsgProtoBufHeader_TransportError
}

func (m *CMsgProtoBufHeader) GetMessageid() uint64 {
	if m != nil && m.Messageid != nil {
		return *m.Messageid
	}
	return Default_CMsgProtoBufHeader_Messageid
}

func (m *CMsgProtoBufHeader) GetPublisherGroupId() uint32 {
	if m != nil && m.PublisherGroupId != nil {
		return *m.PublisherGroupId
	}
	return 0
}

func (m *CMsgProtoBufHeader) GetSysid() uint32 {
	if m != nil && m.Sysid != nil {
		return *m.Sysid
	}
	return 0
}

func (m *CMsgProtoBufHeader) GetTraceTag() uint64 {
	if m != nil && m.TraceTag != nil {
		return *m.TraceTag
	}
	return 0
}

type CMsgMulti struct {
	SizeUnzipped     *uint32 `protobuf:"varint,1,opt,name=size_unzipped" json:"size_unzipped,omitempty"`
	MessageBody      []byte  `protobuf:"bytes,2,opt,name=message_body" json:"message_body,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgMulti) Reset()         { *m = CMsgMulti{} }
func (m *CMsgMulti) String() string { return proto.CompactTextString(m) }
func (*CMsgMulti) ProtoMessage()    {}

func (m *CMsgMulti) GetSizeUnzipped() uint32 {
	if m != nil && m.SizeUnzipped != nil {
		return *m.SizeUnzipped
	}
	return 0
}

func (m *CMsgMulti) GetMessageBody() []byte {
	if m != nil {
		return m.MessageBody
	}
	return nil
}

type CMsgProtobufWrapped struct {
	MessageBody      []byte `protobuf:"bytes,1,opt,name=message_body" json:"message_body,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CMsgProtobufWrapped) Reset()         { *m = CMsgProtobufWrapped{} }
func (m *CMsgProtobufWrapped) String() string { return proto.CompactTextString(m) }
func (*CMsgProtobufWrapped) ProtoMessage()    {}

func (m *CMsgProtobufWrapped) GetMessageBody() []byte {
	if m != nil {
		return m.MessageBody
	}
	return nil
}

type CMsgAuthTicket struct {
	Estate           *uint32 `protobuf:"varint,1,opt,name=estate" json:"estate,omitempty"`
	Eresult          *uint32 `protobuf:"varint,2,opt,name=eresult,def=2" json:"eresult,omitempty"`
	Steamid          *uint64 `protobuf:"fixed64,3,opt,name=steamid" json:"steamid,omitempty"`
	Gameid           *uint64 `protobuf:"fixed64,4,opt,name=gameid" json:"gameid,omitempty"`
	HSteamPipe       *uint32 `protobuf:"varint,5,opt,name=h_steam_pipe" json:"h_steam_pipe,omitempty"`
	TicketCrc        *uint32 `protobuf:"varint,6,opt,name=ticket_crc" json:"ticket_crc,omitempty"`
	Ticket           []byte  `protobuf:"bytes,7,opt,name=ticket" json:"ticket,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CMsgAuthTicket) Reset()         { *m = CMsgAuthTicket{} }
func (m *CMsgAuthTicket) String() string { return proto.CompactTextString(m) }
func (*CMsgAuthTicket) ProtoMessage()    {}

const Default_CMsgAuthTicket_Eresult uint32 = 2

func (m *CMsgAuthTicket) GetEstate() uint32 {
	if m != nil && m.Estate != nil {
		return *m.Estate
	}
	return 0
}

func (m *CMsgAuthTicket) GetEresult() uint32 {
	if m != nil && m.Eresult != nil {
		return *m.Eresult
	}
	return Default_CMsgAuthTicket_Eresult
}

func (m *CMsgAuthTicket) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

func (m *CMsgAuthTicket) GetGameid() uint64 {
	if m != nil && m.Gameid != nil {
		return *m.Gameid
	}
	return 0
}

func (m *CMsgAuthTicket) GetHSteamPipe() uint32 {
	if m != nil && m.HSteamPipe != nil {
		return *m.HSteamPipe
	}
	return 0
}

func (m *CMsgAuthTicket) GetTicketCrc() uint32 {
	if m != nil && m.TicketCrc != nil {
		return *m.TicketCrc
	}
	return 0
}

func (m *CMsgAuthTicket) GetTicket() []byte {
	if m != nil {
		return m.Ticket
	}
	return nil
}

type CCDDBAppDetailCommon struct {
	Appid                 *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	Name                  *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Icon                  *string `protobuf:"bytes,3,opt,name=icon" json:"icon,omitempty"`
	Logo                  *string `protobuf:"bytes,4,opt,name=logo" json:"logo,omitempty"`
	LogoSmall             *string `protobuf:"bytes,5,opt,name=logo_small" json:"logo_small,omitempty"`
	Tool                  *bool   `protobuf:"varint,6,opt,name=tool" json:"tool,omitempty"`
	Demo                  *bool   `protobuf:"varint,7,opt,name=demo" json:"demo,omitempty"`
	Media                 *bool   `protobuf:"varint,8,opt,name=media" json:"media,omitempty"`
	CommunityVisibleStats *bool   `protobuf:"varint,9,opt,name=community_visible_stats" json:"community_visible_stats,omitempty"`
	FriendlyName          *string `protobuf:"bytes,10,opt,name=friendly_name" json:"friendly_name,omitempty"`
	Propagation           *string `protobuf:"bytes,11,opt,name=propagation" json:"propagation,omitempty"`
	XXX_unrecognized      []byte  `json:"-"`
}

func (m *CCDDBAppDetailCommon) Reset()         { *m = CCDDBAppDetailCommon{} }
func (m *CCDDBAppDetailCommon) String() string { return proto.CompactTextString(m) }
func (*CCDDBAppDetailCommon) ProtoMessage()    {}

func (m *CCDDBAppDetailCommon) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CCDDBAppDetailCommon) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CCDDBAppDetailCommon) GetIcon() string {
	if m != nil && m.Icon != nil {
		return *m.Icon
	}
	return ""
}

func (m *CCDDBAppDetailCommon) GetLogo() string {
	if m != nil && m.Logo != nil {
		return *m.Logo
	}
	return ""
}

func (m *CCDDBAppDetailCommon) GetLogoSmall() string {
	if m != nil && m.LogoSmall != nil {
		return *m.LogoSmall
	}
	return ""
}

func (m *CCDDBAppDetailCommon) GetTool() bool {
	if m != nil && m.Tool != nil {
		return *m.Tool
	}
	return false
}

func (m *CCDDBAppDetailCommon) GetDemo() bool {
	if m != nil && m.Demo != nil {
		return *m.Demo
	}
	return false
}

func (m *CCDDBAppDetailCommon) GetMedia() bool {
	if m != nil && m.Media != nil {
		return *m.Media
	}
	return false
}

func (m *CCDDBAppDetailCommon) GetCommunityVisibleStats() bool {
	if m != nil && m.CommunityVisibleStats != nil {
		return *m.CommunityVisibleStats
	}
	return false
}

func (m *CCDDBAppDetailCommon) GetFriendlyName() string {
	if m != nil && m.FriendlyName != nil {
		return *m.FriendlyName
	}
	return ""
}

func (m *CCDDBAppDetailCommon) GetPropagation() string {
	if m != nil && m.Propagation != nil {
		return *m.Propagation
	}
	return ""
}

type CMsgAppRights struct {
	EditInfo                 *bool  `protobuf:"varint,1,opt,name=edit_info" json:"edit_info,omitempty"`
	Publish                  *bool  `protobuf:"varint,2,opt,name=publish" json:"publish,omitempty"`
	ViewErrorData            *bool  `protobuf:"varint,3,opt,name=view_error_data" json:"view_error_data,omitempty"`
	Download                 *bool  `protobuf:"varint,4,opt,name=download" json:"download,omitempty"`
	UploadCdkeys             *bool  `protobuf:"varint,5,opt,name=upload_cdkeys" json:"upload_cdkeys,omitempty"`
	GenerateCdkeys           *bool  `protobuf:"varint,6,opt,name=generate_cdkeys" json:"generate_cdkeys,omitempty"`
	ViewFinancials           *bool  `protobuf:"varint,7,opt,name=view_financials" json:"view_financials,omitempty"`
	ManageCeg                *bool  `protobuf:"varint,8,opt,name=manage_ceg" json:"manage_ceg,omitempty"`
	ManageSigning            *bool  `protobuf:"varint,9,opt,name=manage_signing" json:"manage_signing,omitempty"`
	ManageCdkeys             *bool  `protobuf:"varint,10,opt,name=manage_cdkeys" json:"manage_cdkeys,omitempty"`
	EditMarketing            *bool  `protobuf:"varint,11,opt,name=edit_marketing" json:"edit_marketing,omitempty"`
	EconomySupport           *bool  `protobuf:"varint,12,opt,name=economy_support" json:"economy_support,omitempty"`
	EconomySupportSupervisor *bool  `protobuf:"varint,13,opt,name=economy_support_supervisor" json:"economy_support_supervisor,omitempty"`
	XXX_unrecognized         []byte `json:"-"`
}

func (m *CMsgAppRights) Reset()         { *m = CMsgAppRights{} }
func (m *CMsgAppRights) String() string { return proto.CompactTextString(m) }
func (*CMsgAppRights) ProtoMessage()    {}

func (m *CMsgAppRights) GetEditInfo() bool {
	if m != nil && m.EditInfo != nil {
		return *m.EditInfo
	}
	return false
}

func (m *CMsgAppRights) GetPublish() bool {
	if m != nil && m.Publish != nil {
		return *m.Publish
	}
	return false
}

func (m *CMsgAppRights) GetViewErrorData() bool {
	if m != nil && m.ViewErrorData != nil {
		return *m.ViewErrorData
	}
	return false
}

func (m *CMsgAppRights) GetDownload() bool {
	if m != nil && m.Download != nil {
		return *m.Download
	}
	return false
}

func (m *CMsgAppRights) GetUploadCdkeys() bool {
	if m != nil && m.UploadCdkeys != nil {
		return *m.UploadCdkeys
	}
	return false
}

func (m *CMsgAppRights) GetGenerateCdkeys() bool {
	if m != nil && m.GenerateCdkeys != nil {
		return *m.GenerateCdkeys
	}
	return false
}

func (m *CMsgAppRights) GetViewFinancials() bool {
	if m != nil && m.ViewFinancials != nil {
		return *m.ViewFinancials
	}
	return false
}

func (m *CMsgAppRights) GetManageCeg() bool {
	if m != nil && m.ManageCeg != nil {
		return *m.ManageCeg
	}
	return false
}

func (m *CMsgAppRights) GetManageSigning() bool {
	if m != nil && m.ManageSigning != nil {
		return *m.ManageSigning
	}
	return false
}

func (m *CMsgAppRights) GetManageCdkeys() bool {
	if m != nil && m.ManageCdkeys != nil {
		return *m.ManageCdkeys
	}
	return false
}

func (m *CMsgAppRights) GetEditMarketing() bool {
	if m != nil && m.EditMarketing != nil {
		return *m.EditMarketing
	}
	return false
}

func (m *CMsgAppRights) GetEconomySupport() bool {
	if m != nil && m.EconomySupport != nil {
		return *m.EconomySupport
	}
	return false
}

func (m *CMsgAppRights) GetEconomySupportSupervisor() bool {
	if m != nil && m.EconomySupportSupervisor != nil {
		return *m.EconomySupportSupervisor
	}
	return false
}

var E_MsgpoolSoftLimit = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*int32)(nil),
	Field:         50000,
	Name:          "msgpool_soft_limit",
	Tag:           "varint,50000,opt,name=msgpool_soft_limit,def=32",
}

var E_MsgpoolHardLimit = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*int32)(nil),
	Field:         50001,
	Name:          "msgpool_hard_limit",
	Tag:           "varint,50001,opt,name=msgpool_hard_limit,def=384",
}

func init() {
	proto.RegisterExtension(E_MsgpoolSoftLimit)
	proto.RegisterExtension(E_MsgpoolHardLimit)
}
