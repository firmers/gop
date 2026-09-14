/*
 * Copyright (c) 2019, 2026, firmer.tech and/or its affiliates. All rights reserved.
 * Firmer Corporation PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package kinds

import "github.com/firmers/mesh/client/golang/types"

const (
	PushBusStateDraft    = "draft"
	PushBusStateActive   = "active"
	PushBusStateDisabled = "disabled"
	PushBusStateInvalid  = "invalid"

	PushBusVersionDraft     = "draft"
	PushBusVersionValidated = "validated"
	PushBusVersionPublished = "published"
	PushBusVersionRetired   = "retired"
	PushBusVersionInvalid   = "invalid"

	PushBusCardEntry     = "entry"
	PushBusCardAudience  = "audience"
	PushBusCardTemplate  = "template"
	PushBusCardTransform = "transform"
	PushBusCardCondition = "condition"
	PushBusCardFrequency = "frequency"
	PushBusCardProvider  = "provider"
	PushBusCardFallback  = "fallback"
	PushBusCardRetry     = "retry"
	PushBusCardWebhook   = "webhook"
	PushBusCardSuccess   = "success"
	PushBusCardReject    = "reject"
	PushBusCardFailure   = "failure"

	PushBusIssueError   = "error"
	PushBusIssueWarning = "warning"

	PushBusComponentStateActive     = "active"
	PushBusComponentStateDeprecated = "deprecated"
	PushBusComponentStateDisabled   = "disabled"

	PushBusPortInput  = "input"
	PushBusPortOutput = "output"

	PushBusResourceAudience = "audience"
	PushBusResourceTemplate = "template"
	PushBusResourceQuota    = "quota"
	PushBusResourceProvider = "provider"
	PushBusResourceRoute    = "route"
	PushBusResourceWebhook  = "webhook"
)

// MessageRecord states are independent of provider-specific codes.
// 消息列表的文本状态；不用于MessageReceipt的int64回执状态码。sent不代表设备已送达。
const (
	MessageStatusPending   = "pending"   // 待发送，包含尚未到达计划发送时间。
	MessageStatusRunning   = "running"   // 正在处理，尚无最终投递结论。
	MessageStatusSent      = "sent"      // 服务商已受理，尚无设备送达证明。
	MessageStatusDelivered = "delivered" // 已确认送达。
	MessageStatusOpened    = "opened"    // 已确认打开。
	MessageStatusFailed    = "failed"    // 已确认失败。
	MessageStatusUncertain = "uncertain" // 结果未知，不推断成功或失败。
	MessageStatusPartial   = "partial"   // 多接收对象的处理结果不一致。
	MessageStatusCanceled  = "canceled"  // 已取消尚未发送的消息。
)

// MessageReceipt status codes are shared by business queries and callbacks.
// 业务回执状态码，数值属于对外契约，不随服务商状态码变化。
const (
	MessageReceiptStatusSent      int64 = 1  // 已受理，尚无设备送达证明。
	MessageReceiptStatusDelivered int64 = 2  // 已确认送达。
	MessageReceiptStatusOpened    int64 = 4  // 已确认打开。
	MessageReceiptStatusFailed    int64 = 8  // 已确认失败。
	MessageReceiptStatusUncertain int64 = 16 // 结果未知，不推断成功或失败。
)

// MessageReceipt is a business delivery event shared by queries and callbacks.
// MessageReceipt 是查询和业务回调共用的回执，不包含通道、Token 或服务商原始响应。
type MessageReceipt struct {
	BizId         string     `index:"1" json:"biz_id" xml:"biz_id" yaml:"biz_id" comment:"业务编号，对应Message.BizId"`
	JobId         string     `index:"2" json:"job_id" xml:"job_id" yaml:"job_id" comment:"推送任务ID，对应MsgBus.Push返回值"`
	Recipient     string     `index:"3" json:"recipient" xml:"recipient" yaml:"recipient" comment:"业务接收对象ID，不返回设备Token"`
	RecipientType string     `index:"4" json:"recipient_type" xml:"recipient_type" yaml:"recipient_type" comment:"接收对象类型：crowd人群、uid用户、did设备，与Message.Recipient标识体系一致"`
	Status        int64      `index:"5" json:"status" xml:"status" yaml:"status" comment:"业务回执状态码：1已受理、2已送达、4已打开、8失败、16结果未知；不使用服务商原始状态码"`
	Memo          string     `index:"7" json:"memo" xml:"memo" yaml:"memo" comment:"业务可读且已脱敏的回执备注，可说明失败或结果未知原因，不返回内部异常和原始响应"`
	EventAt       types.Time `index:"8" json:"event_at" xml:"event_at" yaml:"event_at" comment:"回执事件发生时间；重投同一事件时保持不变"`
}

// MessageRecord summarizes one business push request, not an internal delivery attempt.
// MessageRecord 表示一条业务推送请求，重试或多设备投递不会扩展成多条业务请求记录。
type MessageRecord struct {
	BizId     string            `index:"1" json:"biz_id" xml:"biz_id" yaml:"biz_id" comment:"业务编号，对应Message.BizId"`
	JobId     string            `index:"2" json:"job_id" xml:"job_id" yaml:"job_id" comment:"推送任务ID，对应MsgBus.Push返回值"`
	BusId     string            `index:"3" json:"bus_id" xml:"bus_id" yaml:"bus_id" comment:"执行推送的总线ID"`
	Recipient []*MessageReceipt `index:"4" json:"recipient" xml:"recipient" yaml:"recipient" comment:"业务接收对象的回执集合，包含对象标识、类型、回执状态、备注和事件时间，不返回Token和内部投递信息"`
	Status    string            `index:"5" json:"status" xml:"status" yaml:"status" comment:"整体状态：pending待发送、running处理中、sent已受理、delivered已送达、opened已打开、partial部分结果、failed失败、uncertain结果未知、canceled已取消"`
	ActiveAt  types.Time        `index:"6" json:"active_at" xml:"active_at" yaml:"active_at" comment:"计划发送时间，对应Message.ActiveAt，零值表示立即发送"`
	CreateAt  types.Time        `index:"7" json:"create_at" xml:"create_at" yaml:"create_at" comment:"平台受理业务推送请求的时间"`
}

// Message submits business parameters to a published message bus.
// 业务消息只包含业务标识、上下文、接收对象和计划发送信息。
type Message struct {
	BizId     string            `index:"1" json:"biz_id" xml:"biz_id" yaml:"biz_id" comment:"业务编号，也是幂等键，可用于推送回执查询"`
	BusId     string            `index:"2" json:"bus_id" xml:"bus_id" yaml:"bus_id" comment:"总线ID，执行推送的消息总线"`
	Ctx       map[string]string `index:"3" json:"ctx" xml:"ctx" yaml:"ctx" comment:"业务上下文参数，模板中的占位符或者DAG节点需要的参数"`
	Recipient map[string]string `index:"4" json:"recipient" xml:"recipient" yaml:"recipient" comment:"接收人列表 人群/用户/设备ID:类型（crowd/uid/did）"`
	ActiveAt  types.Time        `index:"5" json:"active_at" xml:"active_at" yaml:"active_at" comment:"真实推送时间，营销延迟推送"`
	CreateBy  string            `index:"6" json:"create_by" xml:"create_by" yaml:"create_by" comment:"提交人"`
}

// PushTraceQuery identifies a business request by business or job identity.
// 业务回执查询仅按BizId或JobId定位请求，不暴露内部投递标识。
type PushTraceQuery struct {
	BizId string `index:"1" json:"biz_id" xml:"biz_id" yaml:"biz_id" comment:"业务编号，也是幂等键，可用于推送回执查询"`
	JobId string `index:"2" json:"job_id" xml:"job_id" yaml:"job_id" comment:"总线推送任务ID"`
}

type PushReceipt struct {
	Id             string     `index:"1" json:"id" xml:"id" yaml:"id" comment:"回执ID"`
	NotificationId string     `index:"2" json:"notification_id" xml:"notification_id" yaml:"notification_id" comment:"通知ID"`
	TaskId         string     `index:"3" json:"task_id" xml:"task_id" yaml:"task_id" comment:"任务ID"`
	AppId          string     `index:"4" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	ChannelId      string     `index:"5" json:"channel_id" xml:"channel_id" yaml:"channel_id" comment:"通道ID"`
	Platform       string     `index:"6" json:"platform" xml:"platform" yaml:"platform" comment:"平台"`
	ProviderId     string     `index:"7" json:"provider_id" xml:"provider_id" yaml:"provider_id" comment:"服务商ID"`
	ProviderStatus string     `index:"8" json:"provider_status" xml:"provider_status" yaml:"provider_status" comment:"服务商状态"`
	ReceiptType    string     `index:"9" json:"receipt_type" xml:"receipt_type" yaml:"receipt_type" comment:"回执类型"`
	TokenHash      string     `index:"10" json:"token_hash" xml:"token_hash" yaml:"token_hash" comment:"Token摘要"`
	Response       string     `index:"11" json:"response" xml:"response" yaml:"response" comment:"响应JSON"`
	ErrorCode      string     `index:"12" json:"error_code" xml:"error_code" yaml:"error_code" comment:"错误码"`
	ErrorMessage   string     `index:"13" json:"error_message" xml:"error_message" yaml:"error_message" comment:"错误消息"`
	CreateAt       types.Time `index:"14" json:"create_at" xml:"create_at" yaml:"create_at" comment:"创建时间"`
}

type PushNotification struct {
	Id         string     `index:"1" json:"id" xml:"id" yaml:"id" comment:"通知ID"`
	TaskId     string     `index:"2" json:"task_id" xml:"task_id" yaml:"task_id" comment:"任务ID"`
	AppId      string     `index:"3" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	ChannelId  string     `index:"4" json:"channel_id" xml:"channel_id" yaml:"channel_id" comment:"通道ID"`
	Platform   string     `index:"5" json:"platform" xml:"platform" yaml:"platform" comment:"平台"`
	Target     string     `index:"6" json:"target" xml:"target" yaml:"target" comment:"目标"`
	TokenHash  string     `index:"7" json:"token_hash" xml:"token_hash" yaml:"token_hash" comment:"Token摘要"`
	UserId     string     `index:"8" json:"user_id" xml:"user_id" yaml:"user_id" comment:"用户ID"`
	DeviceId   string     `index:"9" json:"device_id" xml:"device_id" yaml:"device_id" comment:"设备ID"`
	Title      string     `index:"10" json:"title" xml:"title" yaml:"title" comment:"标题"`
	Body       string     `index:"11" json:"body" xml:"body" yaml:"body" comment:"内容"`
	Payload    string     `index:"12" json:"payload" xml:"payload" yaml:"payload" comment:"载荷JSON"`
	ProviderId string     `index:"13" json:"provider_id" xml:"provider_id" yaml:"provider_id" comment:"服务商ID"`
	State      string     `index:"14" json:"state" xml:"state" yaml:"state" comment:"状态文本"`
	Result     string     `index:"15" json:"result" xml:"result" yaml:"result" comment:"结果JSON"`
	TraceId    string     `index:"16" json:"trace_id" xml:"trace_id" yaml:"trace_id" comment:"TraceID"`
	MessageId  string     `index:"17" json:"message_id" xml:"message_id" yaml:"message_id" comment:"消息ID"`
	Status     int64      `index:"18" json:"status" xml:"status" yaml:"status" comment:"状态"`
	CreateAt   types.Time `index:"19" json:"create_at" xml:"create_at" yaml:"create_at" comment:"创建时间"`
	UpdateAt   types.Time `index:"20" json:"update_at" xml:"update_at" yaml:"update_at" comment:"更新时间"`
}

type PushBusIdentity struct {
	Id          string `index:"1" json:"id" xml:"id" yaml:"id" comment:"主键"`
	BusId       string `index:"2" json:"bus_id" xml:"bus_id" yaml:"bus_id" comment:"总线ID"`
	Environment string `index:"3" json:"environment" xml:"environment" yaml:"environment" comment:"环境"`
	Version     int64  `index:"4" json:"version" xml:"version" yaml:"version" comment:"版本"`
	RunId       string `index:"5" json:"run_id" xml:"run_id" yaml:"run_id" comment:"运行ID"`
	TraceId     string `index:"6" json:"trace_id" xml:"trace_id" yaml:"trace_id" comment:"TraceID"`
	Reason      string `index:"7" json:"reason" xml:"reason" yaml:"reason" comment:"操作原因"`
}

type PushBusQuery struct {
	Id          string     `index:"1" json:"id" xml:"id" yaml:"id" comment:"主键"`
	BusId       string     `index:"2" json:"bus_id" xml:"bus_id" yaml:"bus_id" comment:"总线ID"`
	AppId       string     `index:"3" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	Domain      string     `index:"4" json:"domain" xml:"domain" yaml:"domain" comment:"业务域"`
	Environment string     `index:"5" json:"environment" xml:"environment" yaml:"environment" comment:"环境"`
	Owner       string     `index:"6" json:"owner" xml:"owner" yaml:"owner" comment:"Owner"`
	State       string     `index:"7" json:"state" xml:"state" yaml:"state" comment:"状态"`
	RunId       string     `index:"8" json:"run_id" xml:"run_id" yaml:"run_id" comment:"运行ID"`
	TraceId     string     `index:"9" json:"trace_id" xml:"trace_id" yaml:"trace_id" comment:"TraceID"`
	FromAt      types.Time `index:"10" json:"from_at" xml:"from_at" yaml:"from_at" comment:"开始时间"`
	ToAt        types.Time `index:"11" json:"to_at" xml:"to_at" yaml:"to_at" comment:"结束时间"`
	Limit       int64      `index:"12" json:"limit" xml:"limit" yaml:"limit" comment:"限制条数"`
}

type PushBus struct {
	Id               string     `index:"1" json:"id" xml:"id" yaml:"id" comment:"主键"`
	BusId            string     `index:"2" json:"bus_id" xml:"bus_id" yaml:"bus_id" comment:"总线ID"`
	Name             string     `index:"3" json:"name" xml:"name" yaml:"name" comment:"总线名称"`
	AppId            string     `index:"4" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	AppName          string     `index:"5" json:"app_name" xml:"app_name" yaml:"app_name" comment:"应用名称"`
	Domain           string     `index:"6" json:"domain" xml:"domain" yaml:"domain" comment:"业务域"`
	Environment      string     `index:"7" json:"environment" xml:"environment" yaml:"environment" comment:"环境"`
	Owner            string     `index:"8" json:"owner" xml:"owner" yaml:"owner" comment:"Owner"`
	DraftVersion     int64      `index:"9" json:"draft_version" xml:"draft_version" yaml:"draft_version" comment:"草稿版本"`
	PublishedVersion int64      `index:"10" json:"published_version" xml:"published_version" yaml:"published_version" comment:"发布版本"`
	CardCount        int64      `index:"11" json:"card_count" xml:"card_count" yaml:"card_count" comment:"Card数量"`
	CallCount        int64      `index:"12" json:"call_count" xml:"call_count" yaml:"call_count" comment:"调用量"`
	SuccessRate      float64    `index:"13" json:"success_rate" xml:"success_rate" yaml:"success_rate" comment:"成功率"`
	Labels           string     `index:"14" json:"labels" xml:"labels" yaml:"labels" comment:"标签JSON"`
	Memo             string     `index:"15" json:"memo" xml:"memo" yaml:"memo" comment:"备注"`
	State            string     `index:"16" json:"state" xml:"state" yaml:"state" comment:"状态"`
	Status           int64      `index:"17" json:"status" xml:"status" yaml:"status" comment:"状态值"`
	CreateAt         types.Time `index:"18" json:"create_at" xml:"create_at" yaml:"create_at" comment:"创建时间"`
	UpdateAt         types.Time `index:"19" json:"update_at" xml:"update_at" yaml:"update_at" comment:"更新时间"`
	CreateBy         string     `index:"20" json:"create_by" xml:"create_by" yaml:"create_by" comment:"创建人"`
	UpdateBy         string     `index:"21" json:"update_by" xml:"update_by" yaml:"update_by" comment:"更新人"`
}

type PushBusPosition struct {
	X float64 `index:"1" json:"x" xml:"x" yaml:"x" comment:"横坐标"`
	Y float64 `index:"2" json:"y" xml:"y" yaml:"y" comment:"纵坐标"`
}

type PushBusPort struct {
	Id        string `index:"1" json:"id" xml:"id" yaml:"id" comment:"端口ID"`
	Label     string `index:"2" json:"label" xml:"label" yaml:"label" comment:"端口标签"`
	Direction string `index:"3" json:"direction" xml:"direction" yaml:"direction" comment:"端口方向"`
}

type PushBusComponentQuery struct {
	State string `index:"1" json:"state" xml:"state" yaml:"state" comment:"组件状态"`
}

type PushBusComponentResourceQuery struct {
	BusId         string `index:"1" json:"bus_id" xml:"bus_id" yaml:"bus_id" comment:"总线ID"`
	Environment   string `index:"2" json:"environment" xml:"environment" yaml:"environment" comment:"环境"`
	ComponentType string `index:"3" json:"component_type" xml:"component_type" yaml:"component_type" comment:"组件类型"`
	Keyword       string `index:"4" json:"keyword" xml:"keyword" yaml:"keyword" comment:"搜索关键字"`
	Limit         int64  `index:"5" json:"limit" xml:"limit" yaml:"limit" comment:"限制条数"`
}

type PushBusComponentGroup struct {
	Id          string `index:"1" json:"id" xml:"id" yaml:"id" comment:"组件组ID"`
	Name        string `index:"2" json:"name" xml:"name" yaml:"name" comment:"组件组名称"`
	Description string `index:"3" json:"description" xml:"description" yaml:"description" comment:"组件组说明"`
	Order       int64  `index:"4" json:"order" xml:"order" yaml:"order" comment:"排序"`
}

type PushBusComponentAppearance struct {
	Icon       string `index:"1" json:"icon" xml:"icon" yaml:"icon" comment:"图标令牌"`
	Color      string `index:"2" json:"color" xml:"color" yaml:"color" comment:"主色"`
	Background string `index:"3" json:"background" xml:"background" yaml:"background" comment:"背景色"`
}

type PushBusComponentOption struct {
	Value string `index:"1" json:"value" xml:"value" yaml:"value" comment:"选项值"`
	Label string `index:"2" json:"label" xml:"label" yaml:"label" comment:"选项名称"`
}

type PushBusComponentField struct {
	Key          string                    `index:"1" json:"key" xml:"key" yaml:"key" comment:"字段键"`
	Label        string                    `index:"2" json:"label" xml:"label" yaml:"label" comment:"字段名称"`
	Description  string                    `index:"3" json:"description" xml:"description" yaml:"description" comment:"字段说明"`
	Control      string                    `index:"4" json:"control" xml:"control" yaml:"control" comment:"控件类型"`
	Required     bool                      `index:"5" json:"required" xml:"required" yaml:"required" comment:"是否必填"`
	DefaultValue string                    `index:"6" json:"default_value" xml:"default_value" yaml:"default_value" comment:"默认值"`
	Placeholder  string                    `index:"7" json:"placeholder" xml:"placeholder" yaml:"placeholder" comment:"占位提示"`
	Options      []*PushBusComponentOption `index:"8" json:"options" xml:"options" yaml:"options" comment:"选项列表"`
}

type PushBusComponentResourceSpec struct {
	Kind     string `index:"1" json:"kind" xml:"kind" yaml:"kind" comment:"资源类型"`
	Required bool   `index:"2" json:"required" xml:"required" yaml:"required" comment:"是否必选"`
}

type PushBusComponentConstraints struct {
	Entry           bool   `index:"1" json:"entry" xml:"entry" yaml:"entry" comment:"是否入口"`
	Terminal        bool   `index:"2" json:"terminal" xml:"terminal" yaml:"terminal" comment:"是否终止"`
	Provider        bool   `index:"3" json:"provider" xml:"provider" yaml:"provider" comment:"是否投递节点"`
	Singleton       bool   `index:"4" json:"singleton" xml:"singleton" yaml:"singleton" comment:"是否单例"`
	RequireOutgoing bool   `index:"5" json:"require_outgoing" xml:"require_outgoing" yaml:"require_outgoing" comment:"是否要求出边"`
	ResultState     string `index:"6" json:"result_state" xml:"result_state" yaml:"result_state" comment:"终止结果状态"`
}

type PushBusComponentDefinition struct {
	Type          string                        `index:"1" json:"type" xml:"type" yaml:"type" comment:"组件类型"`
	Version       string                        `index:"2" json:"version" xml:"version" yaml:"version" comment:"组件版本"`
	Name          string                        `index:"3" json:"name" xml:"name" yaml:"name" comment:"组件名称"`
	Group         string                        `index:"4" json:"group" xml:"group" yaml:"group" comment:"组件组"`
	Description   string                        `index:"5" json:"description" xml:"description" yaml:"description" comment:"组件说明"`
	Handler       string                        `index:"6" json:"handler" xml:"handler" yaml:"handler" comment:"运行时处理器"`
	State         string                        `index:"7" json:"state" xml:"state" yaml:"state" comment:"组件状态"`
	Order         int64                         `index:"8" json:"order" xml:"order" yaml:"order" comment:"排序"`
	Appearance    *PushBusComponentAppearance   `index:"9" json:"appearance" xml:"appearance" yaml:"appearance" comment:"展示样式"`
	Inputs        []*PushBusPort                `index:"10" json:"inputs" xml:"inputs" yaml:"inputs" comment:"输入端口"`
	Outputs       []*PushBusPort                `index:"11" json:"outputs" xml:"outputs" yaml:"outputs" comment:"输出端口"`
	Resource      *PushBusComponentResourceSpec `index:"12" json:"resource" xml:"resource" yaml:"resource" comment:"绑定资源"`
	Fields        []*PushBusComponentField      `index:"13" json:"fields" xml:"fields" yaml:"fields" comment:"属性字段"`
	Constraints   *PushBusComponentConstraints  `index:"14" json:"constraints" xml:"constraints" yaml:"constraints" comment:"图约束"`
	DefaultConfig string                        `index:"15" json:"default_config" xml:"default_config" yaml:"default_config" comment:"默认配置JSON"`
	Metadata      string                        `index:"16" json:"metadata" xml:"metadata" yaml:"metadata" comment:"扩展元数据JSON"`
}

type PushBusComponentResource struct {
	Id          string `index:"1" json:"id" xml:"id" yaml:"id" comment:"资源ID"`
	Code        string `index:"2" json:"code" xml:"code" yaml:"code" comment:"资源编码"`
	Name        string `index:"3" json:"name" xml:"name" yaml:"name" comment:"资源名称"`
	Description string `index:"4" json:"description" xml:"description" yaml:"description" comment:"资源说明"`
	State       string `index:"5" json:"state" xml:"state" yaml:"state" comment:"资源状态"`
	Metadata    string `index:"6" json:"metadata" xml:"metadata" yaml:"metadata" comment:"资源元数据JSON"`
}

type PushBusCard struct {
	Id         string           `index:"1" json:"id" xml:"id" yaml:"id" comment:"Card ID"`
	Type       string           `index:"2" json:"type" xml:"type" yaml:"type" comment:"Card类型"`
	Name       string           `index:"3" json:"name" xml:"name" yaml:"name" comment:"Card名称"`
	ResourceId string           `index:"4" json:"resource_id" xml:"resource_id" yaml:"resource_id" comment:"资源ID"`
	Config     string           `index:"5" json:"config" xml:"config" yaml:"config" comment:"Card配置JSON"`
	Position   *PushBusPosition `index:"6" json:"position" xml:"position" yaml:"position" comment:"画布位置"`
	Inputs     []*PushBusPort   `index:"7" json:"inputs" xml:"inputs" yaml:"inputs" comment:"输入端口"`
	Outputs    []*PushBusPort   `index:"8" json:"outputs" xml:"outputs" yaml:"outputs" comment:"输出端口"`
	Metadata   string           `index:"9" json:"metadata" xml:"metadata" yaml:"metadata" comment:"展示元数据JSON"`
}

type PushBusEdge struct {
	Id         string `index:"1" json:"id" xml:"id" yaml:"id" comment:"边ID"`
	Source     string `index:"2" json:"source" xml:"source" yaml:"source" comment:"源Card"`
	SourcePort string `index:"3" json:"source_port" xml:"source_port" yaml:"source_port" comment:"源端口"`
	Target     string `index:"4" json:"target" xml:"target" yaml:"target" comment:"目标Card"`
	TargetPort string `index:"5" json:"target_port" xml:"target_port" yaml:"target_port" comment:"目标端口"`
	Label      string `index:"6" json:"label" xml:"label" yaml:"label" comment:"边标签"`
}

type PushBusGraph struct {
	Cards            []*PushBusCard `index:"1" json:"cards" xml:"cards" yaml:"cards" comment:"Card列表"`
	Edges            []*PushBusEdge `index:"2" json:"edges" xml:"edges" yaml:"edges" comment:"边列表"`
	Layout           string         `index:"3" json:"layout" xml:"layout" yaml:"layout" comment:"布局JSON"`
	Metadata         string         `index:"4" json:"metadata" xml:"metadata" yaml:"metadata" comment:"图元数据JSON"`
	RegistryVersion  string         `index:"5" json:"registry_version" xml:"registry_version" yaml:"registry_version" comment:"组件注册表版本"`
	RegistryChecksum string         `index:"6" json:"registry_checksum" xml:"registry_checksum" yaml:"registry_checksum" comment:"组件注册表校验和"`
}

type PushBusComponentRegistry struct {
	Version      string                        `index:"1" json:"version" xml:"version" yaml:"version" comment:"注册表版本"`
	Checksum     string                        `index:"2" json:"checksum" xml:"checksum" yaml:"checksum" comment:"注册表校验和"`
	Groups       []*PushBusComponentGroup      `index:"3" json:"groups" xml:"groups" yaml:"groups" comment:"组件组"`
	Components   []*PushBusComponentDefinition `index:"4" json:"components" xml:"components" yaml:"components" comment:"组件定义"`
	StarterGraph *PushBusGraph                 `index:"5" json:"starter_graph" xml:"starter_graph" yaml:"starter_graph" comment:"默认图"`
}

type PushBusValidationIssue struct {
	Code       string `index:"1" json:"code" xml:"code" yaml:"code" comment:"问题编码"`
	Severity   string `index:"2" json:"severity" xml:"severity" yaml:"severity" comment:"严重级别"`
	ObjectKind string `index:"3" json:"object_kind" xml:"object_kind" yaml:"object_kind" comment:"对象类型"`
	ObjectId   string `index:"4" json:"object_id" xml:"object_id" yaml:"object_id" comment:"对象ID"`
	Message    string `index:"5" json:"message" xml:"message" yaml:"message" comment:"问题描述"`
	Hint       string `index:"6" json:"hint" xml:"hint" yaml:"hint" comment:"修复提示"`
}

type PushBusValidation struct {
	Valid       bool                      `index:"1" json:"valid" xml:"valid" yaml:"valid" comment:"是否有效"`
	Checksum    string                    `index:"2" json:"checksum" xml:"checksum" yaml:"checksum" comment:"图校验和"`
	Issues      []*PushBusValidationIssue `index:"3" json:"issues" xml:"issues" yaml:"issues" comment:"问题列表"`
	ValidatedAt types.Time                `index:"4" json:"validated_at" xml:"validated_at" yaml:"validated_at" comment:"校验时间"`
}

type PushBusVersion struct {
	Id          string     `index:"1" json:"id" xml:"id" yaml:"id" comment:"版本主键"`
	BusId       string     `index:"2" json:"bus_id" xml:"bus_id" yaml:"bus_id" comment:"总线ID"`
	AppId       string     `index:"3" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	Environment string     `index:"4" json:"environment" xml:"environment" yaml:"environment" comment:"环境"`
	Version     int64      `index:"5" json:"version" xml:"version" yaml:"version" comment:"版本号"`
	Graph       string     `index:"6" json:"graph" xml:"graph" yaml:"graph" comment:"图JSON"`
	Checksum    string     `index:"7" json:"checksum" xml:"checksum" yaml:"checksum" comment:"图校验和"`
	Validation  string     `index:"8" json:"validation" xml:"validation" yaml:"validation" comment:"校验结果JSON"`
	State       string     `index:"9" json:"state" xml:"state" yaml:"state" comment:"状态"`
	Status      int64      `index:"10" json:"status" xml:"status" yaml:"status" comment:"状态值"`
	PublishedAt types.Time `index:"11" json:"published_at" xml:"published_at" yaml:"published_at" comment:"发布时间"`
	CreateAt    types.Time `index:"12" json:"create_at" xml:"create_at" yaml:"create_at" comment:"创建时间"`
	UpdateAt    types.Time `index:"13" json:"update_at" xml:"update_at" yaml:"update_at" comment:"更新时间"`
	CreateBy    string     `index:"14" json:"create_by" xml:"create_by" yaml:"create_by" comment:"创建人"`
	UpdateBy    string     `index:"15" json:"update_by" xml:"update_by" yaml:"update_by" comment:"更新人"`
}

type PushBusDraftSave struct {
	BusId       string        `index:"1" json:"bus_id" xml:"bus_id" yaml:"bus_id" comment:"总线ID"`
	Environment string        `index:"2" json:"environment" xml:"environment" yaml:"environment" comment:"环境"`
	Version     int64         `index:"3" json:"version" xml:"version" yaml:"version" comment:"草稿版本"`
	Graph       *PushBusGraph `index:"4" json:"graph" xml:"graph" yaml:"graph" comment:"图定义"`
	Reason      string        `index:"5" json:"reason" xml:"reason" yaml:"reason" comment:"保存原因"`
}

type PushBusDraftRequest struct {
	BusId       string `index:"1" json:"bus_id" xml:"bus_id" yaml:"bus_id" comment:"总线ID"`
	Environment string `index:"2" json:"environment" xml:"environment" yaml:"environment" comment:"环境"`
	Version     int64  `index:"3" json:"version" xml:"version" yaml:"version" comment:"草稿版本"`
}

type PushBusDryRunRequest struct {
	BusId       string `index:"1" json:"bus_id" xml:"bus_id" yaml:"bus_id" comment:"总线ID"`
	Environment string `index:"2" json:"environment" xml:"environment" yaml:"environment" comment:"环境"`
	Version     int64  `index:"3" json:"version" xml:"version" yaml:"version" comment:"草稿版本"`
	Payload     string `index:"4" json:"payload" xml:"payload" yaml:"payload" comment:"业务载荷JSON"`
	Recipient   string `index:"5" json:"recipient" xml:"recipient" yaml:"recipient" comment:"接收人JSON"`
	Context     string `index:"6" json:"context" xml:"context" yaml:"context" comment:"上下文JSON"`
}

type PushBusDryRunResult struct {
	Validation   *PushBusValidation `index:"1" json:"validation" xml:"validation" yaml:"validation" comment:"校验结果"`
	SelectedPath []string           `index:"2" json:"selected_path" xml:"selected_path" yaml:"selected_path" comment:"模拟路径"`
	Output       string             `index:"3" json:"output" xml:"output" yaml:"output" comment:"模拟输出JSON"`
}

type PushBusPublishRequest struct {
	BusId       string `index:"1" json:"bus_id" xml:"bus_id" yaml:"bus_id" comment:"总线ID"`
	Environment string `index:"2" json:"environment" xml:"environment" yaml:"environment" comment:"环境"`
	Version     int64  `index:"3" json:"version" xml:"version" yaml:"version" comment:"草稿版本"`
	Reason      string `index:"4" json:"reason" xml:"reason" yaml:"reason" comment:"发布原因"`
}

type PushBusRollbackRequest struct {
	BusId         string `index:"1" json:"bus_id" xml:"bus_id" yaml:"bus_id" comment:"总线ID"`
	Environment   string `index:"2" json:"environment" xml:"environment" yaml:"environment" comment:"环境"`
	TargetVersion int64  `index:"3" json:"target_version" xml:"target_version" yaml:"target_version" comment:"目标版本"`
	Reason        string `index:"4" json:"reason" xml:"reason" yaml:"reason" comment:"回滚原因"`
}

type PushBusVersionDiffRequest struct {
	BusId       string `index:"1" json:"bus_id" xml:"bus_id" yaml:"bus_id" comment:"总线ID"`
	Environment string `index:"2" json:"environment" xml:"environment" yaml:"environment" comment:"环境"`
	FromVersion int64  `index:"3" json:"from_version" xml:"from_version" yaml:"from_version" comment:"源版本"`
	ToVersion   int64  `index:"4" json:"to_version" xml:"to_version" yaml:"to_version" comment:"目标版本"`
}

type PushBusVersionDiff struct {
	FromVersion  int64    `index:"1" json:"from_version" xml:"from_version" yaml:"from_version" comment:"源版本"`
	ToVersion    int64    `index:"2" json:"to_version" xml:"to_version" yaml:"to_version" comment:"目标版本"`
	AddedCards   []string `index:"3" json:"added_cards" xml:"added_cards" yaml:"added_cards" comment:"新增Card"`
	RemovedCards []string `index:"4" json:"removed_cards" xml:"removed_cards" yaml:"removed_cards" comment:"删除Card"`
	ChangedCards []string `index:"5" json:"changed_cards" xml:"changed_cards" yaml:"changed_cards" comment:"变更Card"`
	Summary      string   `index:"6" json:"summary" xml:"summary" yaml:"summary" comment:"差异摘要"`
}

type PushBusRun struct {
	Id             string     `index:"1" json:"id" xml:"id" yaml:"id" comment:"主键"`
	RunId          string     `index:"2" json:"run_id" xml:"run_id" yaml:"run_id" comment:"运行ID"`
	BusId          string     `index:"3" json:"bus_id" xml:"bus_id" yaml:"bus_id" comment:"总线ID"`
	BusVersion     int64      `index:"4" json:"bus_version" xml:"bus_version" yaml:"bus_version" comment:"执行版本"`
	Environment    string     `index:"5" json:"environment" xml:"environment" yaml:"environment" comment:"环境"`
	TraceId        string     `index:"6" json:"trace_id" xml:"trace_id" yaml:"trace_id" comment:"TraceID"`
	TaskId         string     `index:"7" json:"task_id" xml:"task_id" yaml:"task_id" comment:"推送任务ID"`
	IdempotencyKey string     `index:"8" json:"idempotency_key" xml:"idempotency_key" yaml:"idempotency_key" comment:"幂等键"`
	Payload        string     `index:"9" json:"payload" xml:"payload" yaml:"payload" comment:"业务载荷JSON"`
	Recipient      string     `index:"10" json:"recipient" xml:"recipient" yaml:"recipient" comment:"接收人JSON"`
	Context        string     `index:"11" json:"context" xml:"context" yaml:"context" comment:"上下文JSON"`
	SelectedPath   string     `index:"12" json:"selected_path" xml:"selected_path" yaml:"selected_path" comment:"选择路径JSON"`
	State          string     `index:"13" json:"state" xml:"state" yaml:"state" comment:"运行状态"`
	Status         int64      `index:"14" json:"status" xml:"status" yaml:"status" comment:"状态值"`
	ErrorCardId    string     `index:"15" json:"error_card_id" xml:"error_card_id" yaml:"error_card_id" comment:"失败Card"`
	ErrorMessage   string     `index:"16" json:"error_message" xml:"error_message" yaml:"error_message" comment:"失败原因"`
	RetryEligible  bool       `index:"17" json:"retry_eligible" xml:"retry_eligible" yaml:"retry_eligible" comment:"是否可重试"`
	StartAt        types.Time `index:"18" json:"start_at" xml:"start_at" yaml:"start_at" comment:"开始时间"`
	EndAt          types.Time `index:"19" json:"end_at" xml:"end_at" yaml:"end_at" comment:"结束时间"`
	DurationMs     int64      `index:"20" json:"duration_ms" xml:"duration_ms" yaml:"duration_ms" comment:"耗时毫秒"`
	CreateAt       types.Time `index:"21" json:"create_at" xml:"create_at" yaml:"create_at" comment:"创建时间"`
}

type PushBusOverview struct {
	BusCount       int64             `index:"1" json:"bus_count" xml:"bus_count" yaml:"bus_count" comment:"总线数量"`
	ActiveCount    int64             `index:"2" json:"active_count" xml:"active_count" yaml:"active_count" comment:"运行总线"`
	CallCount      int64             `index:"3" json:"call_count" xml:"call_count" yaml:"call_count" comment:"今日调用"`
	SuccessCount   int64             `index:"4" json:"success_count" xml:"success_count" yaml:"success_count" comment:"成功调用"`
	ThrottleCount  int64             `index:"5" json:"throttle_count" xml:"throttle_count" yaml:"throttle_count" comment:"控频拒绝"`
	FailureCount   int64             `index:"6" json:"failure_count" xml:"failure_count" yaml:"failure_count" comment:"失败调用"`
	SuccessRate    float64           `index:"7" json:"success_rate" xml:"success_rate" yaml:"success_rate" comment:"成功率"`
	UnhealthyCount int64             `index:"8" json:"unhealthy_count" xml:"unhealthy_count" yaml:"unhealthy_count" comment:"异常总线"`
	RecentRuns     []*PushBusRun     `index:"9" json:"recent_runs" xml:"recent_runs" yaml:"recent_runs" comment:"最近运行"`
	RecentVersions []*PushBusVersion `index:"10" json:"recent_versions" xml:"recent_versions" yaml:"recent_versions" comment:"最近发布"`
}

type PushCardRun struct {
	Id               string     `index:"1" json:"id" xml:"id" yaml:"id" comment:"主键"`
	RunId            string     `index:"2" json:"run_id" xml:"run_id" yaml:"run_id" comment:"运行ID"`
	Sequence         int64      `index:"3" json:"sequence" xml:"sequence" yaml:"sequence" comment:"执行序号"`
	CardId           string     `index:"4" json:"card_id" xml:"card_id" yaml:"card_id" comment:"Card ID"`
	CardType         string     `index:"5" json:"card_type" xml:"card_type" yaml:"card_type" comment:"Card类型"`
	CardName         string     `index:"6" json:"card_name" xml:"card_name" yaml:"card_name" comment:"Card名称"`
	SelectedEdgeId   string     `index:"7" json:"selected_edge_id" xml:"selected_edge_id" yaml:"selected_edge_id" comment:"选择边"`
	Input            string     `index:"8" json:"input" xml:"input" yaml:"input" comment:"输入JSON"`
	Output           string     `index:"9" json:"output" xml:"output" yaml:"output" comment:"输出JSON"`
	ProviderResponse string     `index:"10" json:"provider_response" xml:"provider_response" yaml:"provider_response" comment:"脱敏服务商响应"`
	State            string     `index:"11" json:"state" xml:"state" yaml:"state" comment:"运行状态"`
	Status           int64      `index:"12" json:"status" xml:"status" yaml:"status" comment:"状态值"`
	ErrorMessage     string     `index:"13" json:"error_message" xml:"error_message" yaml:"error_message" comment:"失败原因"`
	DurationMs       int64      `index:"14" json:"duration_ms" xml:"duration_ms" yaml:"duration_ms" comment:"耗时毫秒"`
	StartAt          types.Time `index:"15" json:"start_at" xml:"start_at" yaml:"start_at" comment:"开始时间"`
	EndAt            types.Time `index:"16" json:"end_at" xml:"end_at" yaml:"end_at" comment:"结束时间"`
	CreateAt         types.Time `index:"17" json:"create_at" xml:"create_at" yaml:"create_at" comment:"创建时间"`
}

type PushBusTraceQuery struct {
	BusId       string `index:"1" json:"bus_id" xml:"bus_id" yaml:"bus_id" comment:"总线ID"`
	Environment string `index:"2" json:"environment" xml:"environment" yaml:"environment" comment:"环境"`
	RunId       string `index:"3" json:"run_id" xml:"run_id" yaml:"run_id" comment:"运行ID"`
	TraceId     string `index:"4" json:"trace_id" xml:"trace_id" yaml:"trace_id" comment:"TraceID"`
	TaskId      string `index:"5" json:"task_id" xml:"task_id" yaml:"task_id" comment:"任务ID"`
}

type PushBusTrace struct {
	Run      *PushBusRun     `index:"1" json:"run" xml:"run" yaml:"run" comment:"总线运行"`
	Version  *PushBusVersion `index:"2" json:"version" xml:"version" yaml:"version" comment:"执行版本"`
	CardRuns []*PushCardRun  `index:"3" json:"card_runs" xml:"card_runs" yaml:"card_runs" comment:"Card运行列表"`
}

type PushBusRetryRequest struct {
	RunId  string `index:"1" json:"run_id" xml:"run_id" yaml:"run_id" comment:"运行ID"`
	Reason string `index:"2" json:"reason" xml:"reason" yaml:"reason" comment:"重试原因"`
}
