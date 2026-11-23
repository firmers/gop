/*
 * Copyright (c) 2019, 2026, firmer.tech and/or its affiliates. All rights reserved.
 * Firmer Corporation PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package kinds

import "github.com/firmers/mesh/client/golang/types"

const (
	PushPlatformIOS     = "ios"
	PushPlatformAndroid = "android"
	PushPlatformHuawei  = "huawei"

	PushStateDraft     = "draft"
	PushStateActive    = "active"
	PushStateDisabled  = "disabled"
	PushStatePending   = "pending"
	PushStateRunning   = "running"
	PushStateSuccess   = "success"
	PushStateFailed    = "failed"
	PushStatePartial   = "partial"
	PushStateRetrying  = "retrying"
	PushStatePaused    = "paused"
	PushStateCanceled  = "canceled"
	PushStateDead      = "dead"
	PushStatePublished = "published"
)

type PushIdentity struct {
	Id        string `index:"1" json:"id" xml:"id" yaml:"id" comment:"主键"`
	AppId     string `index:"2" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	ChannelId string `index:"3" json:"channel_id" xml:"channel_id" yaml:"channel_id" comment:"通道ID"`
	Kind      string `index:"4" json:"kind" xml:"kind" yaml:"kind" comment:"对象类型"`
	Reason    string `index:"5" json:"reason" xml:"reason" yaml:"reason" comment:"操作原因"`
}

type PushQuery struct {
	Id             string     `index:"1" json:"id" xml:"id" yaml:"id" comment:"主键"`
	AppId          string     `index:"2" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	ChannelId      string     `index:"3" json:"channel_id" xml:"channel_id" yaml:"channel_id" comment:"通道ID"`
	Platform       string     `index:"4" json:"platform" xml:"platform" yaml:"platform" comment:"平台"`
	Environment    string     `index:"5" json:"environment" xml:"environment" yaml:"environment" comment:"环境"`
	UserId         string     `index:"6" json:"user_id" xml:"user_id" yaml:"user_id" comment:"用户ID"`
	DeviceId       string     `index:"7" json:"device_id" xml:"device_id" yaml:"device_id" comment:"设备ID"`
	TaskId         string     `index:"8" json:"task_id" xml:"task_id" yaml:"task_id" comment:"任务ID"`
	NotificationId string     `index:"9" json:"notification_id" xml:"notification_id" yaml:"notification_id" comment:"通知ID"`
	MessageId      string     `index:"10" json:"message_id" xml:"message_id" yaml:"message_id" comment:"消息ID"`
	TraceId        string     `index:"11" json:"trace_id" xml:"trace_id" yaml:"trace_id" comment:"TraceID"`
	Status         string     `index:"12" json:"status" xml:"status" yaml:"status" comment:"状态"`
	FromAt         types.Time `index:"13" json:"from_at" xml:"from_at" yaml:"from_at" comment:"开始时间"`
	ToAt           types.Time `index:"14" json:"to_at" xml:"to_at" yaml:"to_at" comment:"结束时间"`
	Limit          int64      `index:"15" json:"limit" xml:"limit" yaml:"limit" comment:"限制条数"`
}

type PushApp struct {
	Id             string     `index:"1" json:"id" xml:"id" yaml:"id" comment:"应用ID"`
	Code           string     `index:"2" json:"code" xml:"code" yaml:"code" comment:"应用编码"`
	Name           string     `index:"3" json:"name" xml:"name" yaml:"name" comment:"应用名称"`
	DisplayName    string     `index:"4" json:"display_name" xml:"display_name" yaml:"display_name" comment:"显示名称"`
	Environment    string     `index:"5" json:"environment" xml:"environment" yaml:"environment" comment:"环境"`
	LdcPath        string     `index:"6" json:"ldc_path" xml:"ldc_path" yaml:"ldc_path" comment:"LDC路径"`
	Owner          string     `index:"7" json:"owner" xml:"owner" yaml:"owner" comment:"Owner"`
	AuthMode       string     `index:"8" json:"auth_mode" xml:"auth_mode" yaml:"auth_mode" comment:"认证方式"`
	CallbackConfig string     `index:"9" json:"callback_config" xml:"callback_config" yaml:"callback_config" comment:"回调配置JSON"`
	Labels         string     `index:"10" json:"labels" xml:"labels" yaml:"labels" comment:"标签JSON"`
	Memo           string     `index:"11" json:"memo" xml:"memo" yaml:"memo" comment:"备注"`
	State          string     `index:"12" json:"state" xml:"state" yaml:"state" comment:"状态文本"`
	Status         int64      `index:"13" json:"status" xml:"status" yaml:"status" comment:"状态"`
	CreateAt       types.Time `index:"14" json:"create_at" xml:"create_at" yaml:"create_at" comment:"创建时间"`
	UpdateAt       types.Time `index:"15" json:"update_at" xml:"update_at" yaml:"update_at" comment:"更新时间"`
	CreateBy       string     `index:"16" json:"create_by" xml:"create_by" yaml:"create_by" comment:"创建人"`
	UpdateBy       string     `index:"17" json:"update_by" xml:"update_by" yaml:"update_by" comment:"更新人"`
}

type PushChannel struct {
	Id            string              `index:"1" json:"id" xml:"id" yaml:"id" comment:"通道ID"`
	AppId         string              `index:"2" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	AppName       string              `index:"3" json:"app_name" xml:"app_name" yaml:"app_name" comment:"应用名称"`
	Code          string              `index:"4" json:"code" xml:"code" yaml:"code" comment:"通道编码"`
	Name          string              `index:"5" json:"name" xml:"name" yaml:"name" comment:"通道名称"`
	Platform      string              `index:"6" json:"platform" xml:"platform" yaml:"platform" comment:"平台"`
	Provider      string              `index:"7" json:"provider" xml:"provider" yaml:"provider" comment:"服务商"`
	Environment   string              `index:"8" json:"environment" xml:"environment" yaml:"environment" comment:"环境"`
	Region        string              `index:"9" json:"region" xml:"region" yaml:"region" comment:"地域"`
	Endpoint      string              `index:"10" json:"endpoint" xml:"endpoint" yaml:"endpoint" comment:"端点"`
	Config        string              `index:"11" json:"config" xml:"config" yaml:"config" comment:"通道配置JSON"`
	Capabilities  string              `index:"12" json:"capabilities" xml:"capabilities" yaml:"capabilities" comment:"能力JSON"`
	QuotaPolicyId string              `index:"13" json:"quota_policy_id" xml:"quota_policy_id" yaml:"quota_policy_id" comment:"配额策略ID"`
	RoutePolicyId string              `index:"14" json:"route_policy_id" xml:"route_policy_id" yaml:"route_policy_id" comment:"路由策略ID"`
	Memo          string              `index:"15" json:"memo" xml:"memo" yaml:"memo" comment:"备注"`
	State         string              `index:"16" json:"state" xml:"state" yaml:"state" comment:"状态文本"`
	Status        int64               `index:"17" json:"status" xml:"status" yaml:"status" comment:"状态"`
	LastHealthAt  types.Time          `index:"18" json:"last_health_at" xml:"last_health_at" yaml:"last_health_at" comment:"最近健康检查时间"`
	Credentials   *PushProviderConfig `index:"23" json:"credentials" xml:"credentials" yaml:"credentials" comment:"通用供应商配置，空值兼容旧配置"`
	CreateAt      types.Time          `index:"19" json:"create_at" xml:"create_at" yaml:"create_at" comment:"创建时间"`
	UpdateAt      types.Time          `index:"20" json:"update_at" xml:"update_at" yaml:"update_at" comment:"更新时间"`
	CreateBy      string              `index:"21" json:"create_by" xml:"create_by" yaml:"create_by" comment:"创建人"`
	UpdateBy      string              `index:"22" json:"update_by" xml:"update_by" yaml:"update_by" comment:"更新人"`
}

type PushChannelTest struct {
	ChannelId string `index:"1" json:"channel_id" xml:"channel_id" yaml:"channel_id" comment:"通道ID"`
	AppId     string `index:"2" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	Platform  string `index:"3" json:"platform" xml:"platform" yaml:"platform" comment:"平台"`
	Token     string `index:"4" json:"token" xml:"token" yaml:"token" comment:"设备Token"`
	Topic     string `index:"5" json:"topic" xml:"topic" yaml:"topic" comment:"Topic"`
	Condition string `index:"6" json:"condition" xml:"condition" yaml:"condition" comment:"条件"`
	Title     string `index:"7" json:"title" xml:"title" yaml:"title" comment:"标题"`
	Body      string `index:"8" json:"body" xml:"body" yaml:"body" comment:"内容"`
	Data      string `index:"9" json:"data" xml:"data" yaml:"data" comment:"数据JSON"`
	Reason    string `index:"10" json:"reason" xml:"reason" yaml:"reason" comment:"原因"`
}

type PushChannelTestResult struct {
	ChannelId     string     `index:"1" json:"channel_id" xml:"channel_id" yaml:"channel_id" comment:"通道ID"`
	Platform      string     `index:"2" json:"platform" xml:"platform" yaml:"platform" comment:"平台"`
	Status        string     `index:"3" json:"status" xml:"status" yaml:"status" comment:"状态"`
	Engine        string     `index:"4" json:"engine" xml:"engine" yaml:"engine" comment:"引擎"`
	EngineVersion string     `index:"5" json:"engine_version" xml:"engine_version" yaml:"engine_version" comment:"引擎版本"`
	Request       string     `index:"6" json:"request" xml:"request" yaml:"request" comment:"请求JSON"`
	Response      string     `index:"7" json:"response" xml:"response" yaml:"response" comment:"响应JSON"`
	ErrorCode     string     `index:"8" json:"error_code" xml:"error_code" yaml:"error_code" comment:"错误码"`
	ErrorMessage  string     `index:"9" json:"error_message" xml:"error_message" yaml:"error_message" comment:"错误消息"`
	TestedAt      types.Time `index:"10" json:"tested_at" xml:"tested_at" yaml:"tested_at" comment:"测试时间"`
}

type PushTemplate struct {
	Id            string     `index:"1" json:"id" xml:"id" yaml:"id" comment:"模板ID"`
	AppId         string     `index:"2" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	Code          string     `index:"3" json:"code" xml:"code" yaml:"code" comment:"模板编码"`
	Name          string     `index:"4" json:"name" xml:"name" yaml:"name" comment:"模板名称"`
	Version       string     `index:"5" json:"version" xml:"version" yaml:"version" comment:"版本"`
	Locale        string     `index:"6" json:"locale" xml:"locale" yaml:"locale" comment:"语言"`
	Title         string     `index:"7" json:"title" xml:"title" yaml:"title" comment:"标题模板"`
	Body          string     `index:"8" json:"body" xml:"body" yaml:"body" comment:"内容模板"`
	DataSchema    string     `index:"9" json:"data_schema" xml:"data_schema" yaml:"data_schema" comment:"数据结构JSON"`
	DefaultData   string     `index:"10" json:"default_data" xml:"default_data" yaml:"default_data" comment:"默认数据JSON"`
	PublishStatus string     `index:"11" json:"publish_status" xml:"publish_status" yaml:"publish_status" comment:"发布状态"`
	Memo          string     `index:"12" json:"memo" xml:"memo" yaml:"memo" comment:"备注"`
	Status        int64      `index:"13" json:"status" xml:"status" yaml:"status" comment:"状态"`
	CreateAt      types.Time `index:"14" json:"create_at" xml:"create_at" yaml:"create_at" comment:"创建时间"`
	UpdateAt      types.Time `index:"15" json:"update_at" xml:"update_at" yaml:"update_at" comment:"更新时间"`
	CreateBy      string     `index:"16" json:"create_by" xml:"create_by" yaml:"create_by" comment:"创建人"`
	UpdateBy      string     `index:"17" json:"update_by" xml:"update_by" yaml:"update_by" comment:"更新人"`
	Provider      string     `index:"18" json:"provider" xml:"provider" yaml:"provider" comment:"渠道类型"`
}

type PushPublishRequest struct {
	Id            string `index:"1" json:"id" xml:"id" yaml:"id" comment:"对象ID"`
	AppId         string `index:"2" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	PublishStatus string `index:"3" json:"publish_status" xml:"publish_status" yaml:"publish_status" comment:"发布状态"`
	Reason        string `index:"4" json:"reason" xml:"reason" yaml:"reason" comment:"原因"`
}

type PushAudience struct {
	Id            string     `index:"1" json:"id" xml:"id" yaml:"id" comment:"人群ID"`
	AppId         string     `index:"2" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	Code          string     `index:"3" json:"code" xml:"code" yaml:"code" comment:"人群编码"`
	Name          string     `index:"4" json:"name" xml:"name" yaml:"name" comment:"人群名称"`
	Selector      string     `index:"5" json:"selector" xml:"selector" yaml:"selector" comment:"选择器JSON"`
	EstimateCount int64      `index:"6" json:"estimate_count" xml:"estimate_count" yaml:"estimate_count" comment:"预估数量"`
	Memo          string     `index:"7" json:"memo" xml:"memo" yaml:"memo" comment:"备注"`
	Status        int64      `index:"8" json:"status" xml:"status" yaml:"status" comment:"状态"`
	CreateAt      types.Time `index:"9" json:"create_at" xml:"create_at" yaml:"create_at" comment:"创建时间"`
	UpdateAt      types.Time `index:"10" json:"update_at" xml:"update_at" yaml:"update_at" comment:"更新时间"`
	CreateBy      string     `index:"11" json:"create_by" xml:"create_by" yaml:"create_by" comment:"创建人"`
	UpdateBy      string     `index:"12" json:"update_by" xml:"update_by" yaml:"update_by" comment:"更新人"`
}

type PushRoutePolicy struct {
	Id           string     `index:"1" json:"id" xml:"id" yaml:"id" comment:"路由策略ID"`
	AppId        string     `index:"2" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	Code         string     `index:"3" json:"code" xml:"code" yaml:"code" comment:"策略编码"`
	Name         string     `index:"4" json:"name" xml:"name" yaml:"name" comment:"策略名称"`
	Priority     int64      `index:"5" json:"priority" xml:"priority" yaml:"priority" comment:"优先级"`
	MatchRule    string     `index:"6" json:"match_rule" xml:"match_rule" yaml:"match_rule" comment:"匹配规则JSON"`
	ChannelRule  string     `index:"7" json:"channel_rule" xml:"channel_rule" yaml:"channel_rule" comment:"通道规则JSON"`
	FallbackRule string     `index:"8" json:"fallback_rule" xml:"fallback_rule" yaml:"fallback_rule" comment:"降级规则JSON"`
	Memo         string     `index:"9" json:"memo" xml:"memo" yaml:"memo" comment:"备注"`
	Status       int64      `index:"10" json:"status" xml:"status" yaml:"status" comment:"状态"`
	CreateAt     types.Time `index:"11" json:"create_at" xml:"create_at" yaml:"create_at" comment:"创建时间"`
	UpdateAt     types.Time `index:"12" json:"update_at" xml:"update_at" yaml:"update_at" comment:"更新时间"`
	CreateBy     string     `index:"13" json:"create_by" xml:"create_by" yaml:"create_by" comment:"创建人"`
	UpdateBy     string     `index:"14" json:"update_by" xml:"update_by" yaml:"update_by" comment:"更新人"`
}

type PushQuotaPolicy struct {
	Id             string     `index:"1" json:"id" xml:"id" yaml:"id" comment:"配额策略ID"`
	AppId          string     `index:"2" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	Code           string     `index:"3" json:"code" xml:"code" yaml:"code" comment:"策略编码"`
	Name           string     `index:"4" json:"name" xml:"name" yaml:"name" comment:"策略名称"`
	QuotaScope     string     `index:"5" json:"quota_scope" xml:"quota_scope" yaml:"quota_scope" comment:"配额范围"`
	MaxQps         int64      `index:"6" json:"max_qps" xml:"max_qps" yaml:"max_qps" comment:"最大QPS"`
	MaxConcurrency int64      `index:"7" json:"max_concurrency" xml:"max_concurrency" yaml:"max_concurrency" comment:"最大并发"`
	DailyLimit     int64      `index:"8" json:"daily_limit" xml:"daily_limit" yaml:"daily_limit" comment:"每日上限"`
	RuleConfig     string     `index:"9" json:"rule_config" xml:"rule_config" yaml:"rule_config" comment:"规则配置JSON"`
	Memo           string     `index:"10" json:"memo" xml:"memo" yaml:"memo" comment:"备注"`
	Status         int64      `index:"11" json:"status" xml:"status" yaml:"status" comment:"状态"`
	CreateAt       types.Time `index:"12" json:"create_at" xml:"create_at" yaml:"create_at" comment:"创建时间"`
	UpdateAt       types.Time `index:"13" json:"update_at" xml:"update_at" yaml:"update_at" comment:"更新时间"`
	CreateBy       string     `index:"14" json:"create_by" xml:"create_by" yaml:"create_by" comment:"创建人"`
	UpdateBy       string     `index:"15" json:"update_by" xml:"update_by" yaml:"update_by" comment:"更新人"`
}

type PushWebhook struct {
	Id             string     `index:"1" json:"id" xml:"id" yaml:"id" comment:"Webhook ID"`
	AppId          string     `index:"2" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	Code           string     `index:"3" json:"code" xml:"code" yaml:"code" comment:"Webhook编码"`
	Name           string     `index:"4" json:"name" xml:"name" yaml:"name" comment:"Webhook名称"`
	HookType       string     `index:"5" json:"hook_type" xml:"hook_type" yaml:"hook_type" comment:"回调类型"`
	Url            string     `index:"6" json:"url" xml:"url" yaml:"url" comment:"URL"`
	Secret         string     `index:"7" json:"secret" xml:"secret" yaml:"secret" comment:"密钥"`
	Headers        string     `index:"8" json:"headers" xml:"headers" yaml:"headers" comment:"请求头JSON"`
	TimeoutSeconds int64      `index:"9" json:"timeout_seconds" xml:"timeout_seconds" yaml:"timeout_seconds" comment:"超时秒数"`
	Memo           string     `index:"10" json:"memo" xml:"memo" yaml:"memo" comment:"备注"`
	Status         int64      `index:"11" json:"status" xml:"status" yaml:"status" comment:"状态"`
	CreateAt       types.Time `index:"12" json:"create_at" xml:"create_at" yaml:"create_at" comment:"创建时间"`
	UpdateAt       types.Time `index:"13" json:"update_at" xml:"update_at" yaml:"update_at" comment:"更新时间"`
	CreateBy       string     `index:"14" json:"create_by" xml:"create_by" yaml:"create_by" comment:"创建人"`
	UpdateBy       string     `index:"15" json:"update_by" xml:"update_by" yaml:"update_by" comment:"更新人"`
}

type PushAudit struct {
	Id          string     `index:"1" json:"id" xml:"id" yaml:"id" comment:"审计ID"`
	AuditType   string     `index:"2" json:"audit_type" xml:"audit_type" yaml:"audit_type" comment:"审计类型"`
	Operation   string     `index:"3" json:"operation" xml:"operation" yaml:"operation" comment:"操作"`
	Operator    string     `index:"4" json:"operator" xml:"operator" yaml:"operator" comment:"操作人"`
	AppId       string     `index:"5" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	ObjectKind  string     `index:"6" json:"object_kind" xml:"object_kind" yaml:"object_kind" comment:"对象类型"`
	ObjectId    string     `index:"7" json:"object_id" xml:"object_id" yaml:"object_id" comment:"对象ID"`
	ObjectName  string     `index:"8" json:"object_name" xml:"object_name" yaml:"object_name" comment:"对象名称"`
	Environment string     `index:"9" json:"environment" xml:"environment" yaml:"environment" comment:"环境"`
	Result      string     `index:"10" json:"result" xml:"result" yaml:"result" comment:"结果"`
	Request     string     `index:"11" json:"request" xml:"request" yaml:"request" comment:"请求JSON"`
	Response    string     `index:"12" json:"response" xml:"response" yaml:"response" comment:"响应JSON"`
	Diff        string     `index:"13" json:"diff" xml:"diff" yaml:"diff" comment:"变更Diff"`
	Evidence    string     `index:"14" json:"evidence" xml:"evidence" yaml:"evidence" comment:"证据JSON"`
	CreateAt    types.Time `index:"15" json:"create_at" xml:"create_at" yaml:"create_at" comment:"创建时间"`
}

type PushDevice struct {
	Id           string     `index:"1" json:"id" xml:"id" yaml:"id" comment:"设备ID"`
	AppId        string     `index:"2" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	AppName      string     `index:"3" json:"app_name" xml:"app_name" yaml:"app_name" comment:"应用名称"`
	Platform     string     `index:"4" json:"platform" xml:"platform" yaml:"platform" comment:"平台"`
	DeviceId     string     `index:"5" json:"device_id" xml:"device_id" yaml:"device_id" comment:"设备标识"`
	UserId       string     `index:"6" json:"user_id" xml:"user_id" yaml:"user_id" comment:"用户ID"`
	Token        string     `index:"7" json:"token" xml:"token" yaml:"token" comment:"设备Token"`
	TokenHash    string     `index:"8" json:"token_hash" xml:"token_hash" yaml:"token_hash" comment:"Token摘要"`
	Tags         string     `index:"9" json:"tags" xml:"tags" yaml:"tags" comment:"标签JSON"`
	Attributes   string     `index:"10" json:"attributes" xml:"attributes" yaml:"attributes" comment:"属性JSON"`
	Locale       string     `index:"11" json:"locale" xml:"locale" yaml:"locale" comment:"语言"`
	Timezone     string     `index:"12" json:"timezone" xml:"timezone" yaml:"timezone" comment:"时区"`
	AppVersion   string     `index:"13" json:"app_version" xml:"app_version" yaml:"app_version" comment:"应用版本"`
	SdkVersion   string     `index:"14" json:"sdk_version" xml:"sdk_version" yaml:"sdk_version" comment:"SDK版本"`
	State        string     `index:"15" json:"state" xml:"state" yaml:"state" comment:"状态文本"`
	LastActiveAt types.Time `index:"16" json:"last_active_at" xml:"last_active_at" yaml:"last_active_at" comment:"最近活跃时间"`
	RegisterAt   types.Time `index:"17" json:"register_at" xml:"register_at" yaml:"register_at" comment:"注册时间"`
	UnregisterAt types.Time `index:"18" json:"unregister_at" xml:"unregister_at" yaml:"unregister_at" comment:"注销时间"`
	Memo         string     `index:"19" json:"memo" xml:"memo" yaml:"memo" comment:"备注"`
	Status       int64      `index:"20" json:"status" xml:"status" yaml:"status" comment:"状态"`
	CreateAt     types.Time `index:"21" json:"create_at" xml:"create_at" yaml:"create_at" comment:"创建时间"`
	UpdateAt     types.Time `index:"22" json:"update_at" xml:"update_at" yaml:"update_at" comment:"更新时间"`
	CreateBy     string     `index:"23" json:"create_by" xml:"create_by" yaml:"create_by" comment:"创建人"`
	UpdateBy     string     `index:"24" json:"update_by" xml:"update_by" yaml:"update_by" comment:"更新人"`
	ChannelId    string     `index:"25" json:"channel_id,omitempty" xml:"channel_id" yaml:"channel_id" comment:"设备注册通道ID"`
}

type PushDeviceRegistration struct {
	AppId      string `index:"1" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	Platform   string `index:"2" json:"platform" xml:"platform" yaml:"platform" comment:"平台"`
	DeviceId   string `index:"3" json:"device_id" xml:"device_id" yaml:"device_id" comment:"设备标识"`
	UserId     string `index:"4" json:"user_id" xml:"user_id" yaml:"user_id" comment:"用户ID"`
	Token      string `index:"5" json:"token" xml:"token" yaml:"token" comment:"设备Token"`
	Tags       string `index:"6" json:"tags" xml:"tags" yaml:"tags" comment:"标签JSON"`
	Attributes string `index:"7" json:"attributes" xml:"attributes" yaml:"attributes" comment:"属性JSON"`
	Locale     string `index:"8" json:"locale" xml:"locale" yaml:"locale" comment:"语言"`
	Timezone   string `index:"9" json:"timezone" xml:"timezone" yaml:"timezone" comment:"时区"`
	AppVersion string `index:"10" json:"app_version" xml:"app_version" yaml:"app_version" comment:"应用版本"`
	SdkVersion string `index:"11" json:"sdk_version" xml:"sdk_version" yaml:"sdk_version" comment:"SDK版本"`
	ChannelId  string `index:"12" json:"channel_id,omitempty" xml:"channel_id" yaml:"channel_id" comment:"设备注册通道ID"`
}

type PushSendRequest struct {
	AppId          string     `index:"1" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	ChannelId      string     `index:"2" json:"channel_id" xml:"channel_id" yaml:"channel_id" comment:"通道ID"`
	Platform       string     `index:"3" json:"platform" xml:"platform" yaml:"platform" comment:"平台"`
	TargetType     string     `index:"4" json:"target_type" xml:"target_type" yaml:"target_type" comment:"目标类型"`
	Token          string     `index:"5" json:"token" xml:"token" yaml:"token" comment:"单个Token"`
	Tokens         []string   `index:"6" json:"tokens" xml:"tokens" yaml:"tokens" comment:"Token列表"`
	DeviceIds      []string   `index:"7" json:"device_ids" xml:"device_ids" yaml:"device_ids" comment:"设备ID列表"`
	UserIds        []string   `index:"8" json:"user_ids" xml:"user_ids" yaml:"user_ids" comment:"用户ID列表"`
	AudienceId     string     `index:"9" json:"audience_id" xml:"audience_id" yaml:"audience_id" comment:"人群ID"`
	Topic          string     `index:"10" json:"topic" xml:"topic" yaml:"topic" comment:"Topic"`
	Condition      string     `index:"11" json:"condition" xml:"condition" yaml:"condition" comment:"条件"`
	TemplateId     string     `index:"12" json:"template_id" xml:"template_id" yaml:"template_id" comment:"模板ID"`
	TemplateCode   string     `index:"13" json:"template_code" xml:"template_code" yaml:"template_code" comment:"模板编码"`
	TemplateData   string     `index:"14" json:"template_data" xml:"template_data" yaml:"template_data" comment:"模板数据JSON"`
	Title          string     `index:"15" json:"title" xml:"title" yaml:"title" comment:"标题"`
	Body           string     `index:"16" json:"body" xml:"body" yaml:"body" comment:"内容"`
	Image          string     `index:"17" json:"image" xml:"image" yaml:"image" comment:"图片"`
	Sound          string     `index:"18" json:"sound" xml:"sound" yaml:"sound" comment:"声音"`
	Badge          int64      `index:"19" json:"badge" xml:"badge" yaml:"badge" comment:"角标"`
	Data           string     `index:"20" json:"data" xml:"data" yaml:"data" comment:"数据JSON"`
	Priority       string     `index:"21" json:"priority" xml:"priority" yaml:"priority" comment:"优先级"`
	TtlSeconds     int64      `index:"22" json:"ttl_seconds" xml:"ttl_seconds" yaml:"ttl_seconds" comment:"TTL秒数"`
	CollapseKey    string     `index:"23" json:"collapse_key" xml:"collapse_key" yaml:"collapse_key" comment:"折叠Key"`
	MessageId      string     `index:"24" json:"message_id" xml:"message_id" yaml:"message_id" comment:"消息ID"`
	TraceId        string     `index:"25" json:"trace_id" xml:"trace_id" yaml:"trace_id" comment:"TraceID"`
	IdempotencyKey string     `index:"26" json:"idempotency_key" xml:"idempotency_key" yaml:"idempotency_key" comment:"幂等键"`
	ScheduleAt     types.Time `index:"27" json:"schedule_at" xml:"schedule_at" yaml:"schedule_at" comment:"计划时间"`
	Reason         string     `index:"28" json:"reason" xml:"reason" yaml:"reason" comment:"原因"`
	Environment    string     `index:"29" json:"environment" xml:"environment" yaml:"environment" comment:"环境"`
	Tags           string     `index:"30" json:"tags" xml:"tags" yaml:"tags" comment:"标签JSON"`
	Attributes     string     `index:"31" json:"attributes" xml:"attributes" yaml:"attributes" comment:"属性JSON"`
}

type PushBatchSendRequest struct {
	AppId          string     `index:"1" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	ChannelId      string     `index:"2" json:"channel_id" xml:"channel_id" yaml:"channel_id" comment:"通道ID"`
	Platform       string     `index:"3" json:"platform" xml:"platform" yaml:"platform" comment:"平台"`
	TargetType     string     `index:"4" json:"target_type" xml:"target_type" yaml:"target_type" comment:"目标类型"`
	AudienceId     string     `index:"5" json:"audience_id" xml:"audience_id" yaml:"audience_id" comment:"人群ID"`
	Tokens         []string   `index:"6" json:"tokens" xml:"tokens" yaml:"tokens" comment:"Token列表"`
	DeviceIds      []string   `index:"7" json:"device_ids" xml:"device_ids" yaml:"device_ids" comment:"设备ID列表"`
	UserIds        []string   `index:"8" json:"user_ids" xml:"user_ids" yaml:"user_ids" comment:"用户ID列表"`
	TemplateId     string     `index:"9" json:"template_id" xml:"template_id" yaml:"template_id" comment:"模板ID"`
	TemplateCode   string     `index:"10" json:"template_code" xml:"template_code" yaml:"template_code" comment:"模板编码"`
	TemplateData   string     `index:"11" json:"template_data" xml:"template_data" yaml:"template_data" comment:"模板数据JSON"`
	Title          string     `index:"12" json:"title" xml:"title" yaml:"title" comment:"标题"`
	Body           string     `index:"13" json:"body" xml:"body" yaml:"body" comment:"内容"`
	Data           string     `index:"14" json:"data" xml:"data" yaml:"data" comment:"数据JSON"`
	MessageId      string     `index:"15" json:"message_id" xml:"message_id" yaml:"message_id" comment:"消息ID"`
	TraceId        string     `index:"16" json:"trace_id" xml:"trace_id" yaml:"trace_id" comment:"TraceID"`
	IdempotencyKey string     `index:"17" json:"idempotency_key" xml:"idempotency_key" yaml:"idempotency_key" comment:"幂等键"`
	ScheduleAt     types.Time `index:"18" json:"schedule_at" xml:"schedule_at" yaml:"schedule_at" comment:"计划时间"`
	Reason         string     `index:"19" json:"reason" xml:"reason" yaml:"reason" comment:"原因"`
	Environment    string     `index:"20" json:"environment" xml:"environment" yaml:"environment" comment:"环境"`
	Tags           string     `index:"21" json:"tags" xml:"tags" yaml:"tags" comment:"标签JSON"`
	Attributes     string     `index:"22" json:"attributes" xml:"attributes" yaml:"attributes" comment:"属性JSON"`
}

type PushTaskSubmit struct {
	AppId          string     `index:"1" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	ChannelId      string     `index:"2" json:"channel_id" xml:"channel_id" yaml:"channel_id" comment:"通道ID"`
	Platform       string     `index:"3" json:"platform" xml:"platform" yaml:"platform" comment:"平台"`
	TaskType       string     `index:"4" json:"task_type" xml:"task_type" yaml:"task_type" comment:"任务类型"`
	Request        string     `index:"5" json:"request" xml:"request" yaml:"request" comment:"请求JSON"`
	IdempotencyKey string     `index:"6" json:"idempotency_key" xml:"idempotency_key" yaml:"idempotency_key" comment:"幂等键"`
	ScheduleAt     types.Time `index:"7" json:"schedule_at" xml:"schedule_at" yaml:"schedule_at" comment:"计划时间"`
	Reason         string     `index:"8" json:"reason" xml:"reason" yaml:"reason" comment:"原因"`
}

type PushDispatchWindow struct {
	AppId         string     `index:"1" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	ChannelId     string     `index:"2" json:"channel_id" xml:"channel_id" yaml:"channel_id" comment:"通道ID"`
	Platform      string     `index:"3" json:"platform" xml:"platform" yaml:"platform" comment:"平台"`
	TaskType      string     `index:"4" json:"task_type" xml:"task_type" yaml:"task_type" comment:"任务类型"`
	State         string     `index:"5" json:"state" xml:"state" yaml:"state" comment:"任务状态"`
	Limit         int64      `index:"6" json:"limit" xml:"limit" yaml:"limit" comment:"调度数量"`
	ExecuteBefore types.Time `index:"7" json:"execute_before" xml:"execute_before" yaml:"execute_before" comment:"最晚计划时间"`
	DryRun        bool       `index:"8" json:"dry_run" xml:"dry_run" yaml:"dry_run" comment:"是否试运行"`
	Reason        string     `index:"9" json:"reason" xml:"reason" yaml:"reason" comment:"原因"`
}

type PushTaskControl struct {
	TaskId     string     `index:"1" json:"task_id" xml:"task_id" yaml:"task_id" comment:"任务ID"`
	State      string     `index:"2" json:"state" xml:"state" yaml:"state" comment:"目标状态"`
	Reason     string     `index:"3" json:"reason" xml:"reason" yaml:"reason" comment:"原因"`
	ScheduleAt types.Time `index:"4" json:"schedule_at" xml:"schedule_at" yaml:"schedule_at" comment:"计划时间"`
}

type PushTask struct {
	Id              string     `index:"1" json:"id" xml:"id" yaml:"id" comment:"任务ID"`
	AppId           string     `index:"2" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	AppName         string     `index:"3" json:"app_name" xml:"app_name" yaml:"app_name" comment:"应用名称"`
	ChannelId       string     `index:"4" json:"channel_id" xml:"channel_id" yaml:"channel_id" comment:"通道ID"`
	ChannelName     string     `index:"5" json:"channel_name" xml:"channel_name" yaml:"channel_name" comment:"通道名称"`
	Platform        string     `index:"6" json:"platform" xml:"platform" yaml:"platform" comment:"平台"`
	TaskType        string     `index:"7" json:"task_type" xml:"task_type" yaml:"task_type" comment:"任务类型"`
	MessageId       string     `index:"8" json:"message_id" xml:"message_id" yaml:"message_id" comment:"消息ID"`
	TraceId         string     `index:"9" json:"trace_id" xml:"trace_id" yaml:"trace_id" comment:"TraceID"`
	IdempotencyKey  string     `index:"10" json:"idempotency_key" xml:"idempotency_key" yaml:"idempotency_key" comment:"幂等键"`
	TargetCount     int64      `index:"11" json:"target_count" xml:"target_count" yaml:"target_count" comment:"目标数量"`
	SuccessCount    int64      `index:"12" json:"success_count" xml:"success_count" yaml:"success_count" comment:"成功数量"`
	FailureCount    int64      `index:"13" json:"failure_count" xml:"failure_count" yaml:"failure_count" comment:"失败数量"`
	RetryCount      int64      `index:"14" json:"retry_count" xml:"retry_count" yaml:"retry_count" comment:"重试数量"`
	DeadLetterCount int64      `index:"15" json:"dead_letter_count" xml:"dead_letter_count" yaml:"dead_letter_count" comment:"死信数量"`
	Request         string     `index:"16" json:"request" xml:"request" yaml:"request" comment:"请求JSON"`
	Result          string     `index:"17" json:"result" xml:"result" yaml:"result" comment:"结果JSON"`
	State           string     `index:"18" json:"state" xml:"state" yaml:"state" comment:"状态文本"`
	Status          int64      `index:"19" json:"status" xml:"status" yaml:"status" comment:"状态"`
	ScheduleAt      types.Time `index:"20" json:"schedule_at" xml:"schedule_at" yaml:"schedule_at" comment:"计划时间"`
	ExecuteAt       types.Time `index:"21" json:"execute_at" xml:"execute_at" yaml:"execute_at" comment:"执行时间"`
	FinishAt        types.Time `index:"22" json:"finish_at" xml:"finish_at" yaml:"finish_at" comment:"完成时间"`
	CreateAt        types.Time `index:"23" json:"create_at" xml:"create_at" yaml:"create_at" comment:"创建时间"`
	UpdateAt        types.Time `index:"24" json:"update_at" xml:"update_at" yaml:"update_at" comment:"更新时间"`
	CreateBy        string     `index:"25" json:"create_by" xml:"create_by" yaml:"create_by" comment:"创建人"`
	UpdateBy        string     `index:"26" json:"update_by" xml:"update_by" yaml:"update_by" comment:"更新人"`
}

type PushDeliveryAttempt struct {
	Id             string     `index:"1" json:"id" xml:"id" yaml:"id" comment:"投递尝试ID"`
	NotificationId string     `index:"2" json:"notification_id" xml:"notification_id" yaml:"notification_id" comment:"通知ID"`
	TaskId         string     `index:"3" json:"task_id" xml:"task_id" yaml:"task_id" comment:"任务ID"`
	AppId          string     `index:"4" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	ChannelId      string     `index:"5" json:"channel_id" xml:"channel_id" yaml:"channel_id" comment:"通道ID"`
	Platform       string     `index:"6" json:"platform" xml:"platform" yaml:"platform" comment:"平台"`
	AttemptNo      int64      `index:"7" json:"attempt_no" xml:"attempt_no" yaml:"attempt_no" comment:"尝试次数"`
	ProviderId     string     `index:"8" json:"provider_id" xml:"provider_id" yaml:"provider_id" comment:"服务商ID"`
	Request        string     `index:"9" json:"request" xml:"request" yaml:"request" comment:"请求JSON"`
	Response       string     `index:"10" json:"response" xml:"response" yaml:"response" comment:"响应JSON"`
	ErrorCode      string     `index:"11" json:"error_code" xml:"error_code" yaml:"error_code" comment:"错误码"`
	ErrorMessage   string     `index:"12" json:"error_message" xml:"error_message" yaml:"error_message" comment:"错误消息"`
	State          string     `index:"13" json:"state" xml:"state" yaml:"state" comment:"状态文本"`
	Status         int64      `index:"14" json:"status" xml:"status" yaml:"status" comment:"状态"`
	CreateAt       types.Time `index:"15" json:"create_at" xml:"create_at" yaml:"create_at" comment:"创建时间"`
	UpdateAt       types.Time `index:"16" json:"update_at" xml:"update_at" yaml:"update_at" comment:"更新时间"`
}

type PushReceiptCallback struct {
	AppId          string     `index:"1" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	ChannelId      string     `index:"2" json:"channel_id" xml:"channel_id" yaml:"channel_id" comment:"通道ID"`
	Platform       string     `index:"3" json:"platform" xml:"platform" yaml:"platform" comment:"平台"`
	ProviderId     string     `index:"4" json:"provider_id" xml:"provider_id" yaml:"provider_id" comment:"服务商ID"`
	NotificationId string     `index:"5" json:"notification_id" xml:"notification_id" yaml:"notification_id" comment:"通知ID"`
	TaskId         string     `index:"6" json:"task_id" xml:"task_id" yaml:"task_id" comment:"任务ID"`
	MessageId      string     `index:"7" json:"message_id" xml:"message_id" yaml:"message_id" comment:"消息ID"`
	TraceId        string     `index:"8" json:"trace_id" xml:"trace_id" yaml:"trace_id" comment:"TraceID"`
	Status         string     `index:"9" json:"status" xml:"status" yaml:"status" comment:"回执状态：sent服务商受理、delivered设备送达、opened打开、failed失败；不使用success代替送达"`
	Response       string     `index:"10" json:"response" xml:"response" yaml:"response" comment:"响应JSON"`
	Signature      string     `index:"11" json:"signature" xml:"signature" yaml:"signature" comment:"签名"`
	CreateAt       types.Time `index:"12" json:"create_at" xml:"create_at" yaml:"create_at" comment:"创建时间"`
}

type PushRetryRequest struct {
	AppId          string `index:"1" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	TaskId         string `index:"2" json:"task_id" xml:"task_id" yaml:"task_id" comment:"任务ID"`
	NotificationId string `index:"3" json:"notification_id" xml:"notification_id" yaml:"notification_id" comment:"通知ID"`
	DeadLetterId   string `index:"4" json:"dead_letter_id" xml:"dead_letter_id" yaml:"dead_letter_id" comment:"死信ID"`
	Reason         string `index:"5" json:"reason" xml:"reason" yaml:"reason" comment:"原因"`
}

type PushDeadLetter struct {
	Id             string     `index:"1" json:"id" xml:"id" yaml:"id" comment:"死信ID"`
	NotificationId string     `index:"2" json:"notification_id" xml:"notification_id" yaml:"notification_id" comment:"通知ID"`
	TaskId         string     `index:"3" json:"task_id" xml:"task_id" yaml:"task_id" comment:"任务ID"`
	AppId          string     `index:"4" json:"app_id" xml:"app_id" yaml:"app_id" comment:"应用ID"`
	ChannelId      string     `index:"5" json:"channel_id" xml:"channel_id" yaml:"channel_id" comment:"通道ID"`
	Platform       string     `index:"6" json:"platform" xml:"platform" yaml:"platform" comment:"平台"`
	TokenHash      string     `index:"7" json:"token_hash" xml:"token_hash" yaml:"token_hash" comment:"Token摘要"`
	Request        string     `index:"8" json:"request" xml:"request" yaml:"request" comment:"请求JSON"`
	LastResponse   string     `index:"9" json:"last_response" xml:"last_response" yaml:"last_response" comment:"最后响应JSON"`
	ErrorCode      string     `index:"10" json:"error_code" xml:"error_code" yaml:"error_code" comment:"错误码"`
	ErrorMessage   string     `index:"11" json:"error_message" xml:"error_message" yaml:"error_message" comment:"错误消息"`
	RetryCount     int64      `index:"12" json:"retry_count" xml:"retry_count" yaml:"retry_count" comment:"重试次数"`
	State          string     `index:"13" json:"state" xml:"state" yaml:"state" comment:"状态文本"`
	Status         int64      `index:"14" json:"status" xml:"status" yaml:"status" comment:"状态"`
	CreateAt       types.Time `index:"15" json:"create_at" xml:"create_at" yaml:"create_at" comment:"创建时间"`
	UpdateAt       types.Time `index:"16" json:"update_at" xml:"update_at" yaml:"update_at" comment:"更新时间"`
}

type PushSendResult struct {
	Task          *PushTask           `index:"1" json:"task" xml:"task" yaml:"task" comment:"任务"`
	Notifications []*PushNotification `index:"2" json:"notifications" xml:"notifications" yaml:"notifications" comment:"通知列表"`
	Receipts      []*PushReceipt      `index:"3" json:"receipts" xml:"receipts" yaml:"receipts" comment:"回执列表"`
	Status        string              `index:"4" json:"status" xml:"status" yaml:"status" comment:"状态"`
	SuccessCount  int64               `index:"5" json:"success_count" xml:"success_count" yaml:"success_count" comment:"成功数量"`
	FailureCount  int64               `index:"6" json:"failure_count" xml:"failure_count" yaml:"failure_count" comment:"失败数量"`
}

type PushTaskResult struct {
	Task            *PushTask      `index:"1" json:"task" xml:"task" yaml:"task" comment:"任务"`
	SuccessCount    int64          `index:"2" json:"success_count" xml:"success_count" yaml:"success_count" comment:"成功数量"`
	FailureCount    int64          `index:"3" json:"failure_count" xml:"failure_count" yaml:"failure_count" comment:"失败数量"`
	RetryCount      int64          `index:"4" json:"retry_count" xml:"retry_count" yaml:"retry_count" comment:"重试数量"`
	DeadLetterCount int64          `index:"5" json:"dead_letter_count" xml:"dead_letter_count" yaml:"dead_letter_count" comment:"死信数量"`
	Receipts        []*PushReceipt `index:"6" json:"receipts" xml:"receipts" yaml:"receipts" comment:"回执列表"`
}

type PushDispatchReport struct {
	ScannedCount    int64             `index:"1" json:"scanned_count" xml:"scanned_count" yaml:"scanned_count" comment:"扫描数量"`
	ExecutedCount   int64             `index:"2" json:"executed_count" xml:"executed_count" yaml:"executed_count" comment:"执行数量"`
	SuccessCount    int64             `index:"3" json:"success_count" xml:"success_count" yaml:"success_count" comment:"成功数量"`
	FailureCount    int64             `index:"4" json:"failure_count" xml:"failure_count" yaml:"failure_count" comment:"失败数量"`
	RetryCount      int64             `index:"5" json:"retry_count" xml:"retry_count" yaml:"retry_count" comment:"重试数量"`
	DeadLetterCount int64             `index:"6" json:"dead_letter_count" xml:"dead_letter_count" yaml:"dead_letter_count" comment:"死信数量"`
	SkippedCount    int64             `index:"7" json:"skipped_count" xml:"skipped_count" yaml:"skipped_count" comment:"跳过数量"`
	Tasks           []*PushTask       `index:"8" json:"tasks" xml:"tasks" yaml:"tasks" comment:"任务列表"`
	Results         []*PushTaskResult `index:"9" json:"results" xml:"results" yaml:"results" comment:"执行结果"`
	Errors          []string          `index:"10" json:"errors" xml:"errors" yaml:"errors" comment:"错误列表"`
}

type PushTrace struct {
	Task          *PushTask              `index:"1" json:"task" xml:"task" yaml:"task" comment:"任务"`
	Notifications []*PushNotification    `index:"2" json:"notifications" xml:"notifications" yaml:"notifications" comment:"通知列表"`
	Attempts      []*PushDeliveryAttempt `index:"3" json:"attempts" xml:"attempts" yaml:"attempts" comment:"投递尝试"`
	Receipts      []*PushReceipt         `index:"4" json:"receipts" xml:"receipts" yaml:"receipts" comment:"回执列表"`
	DeadLetters   []*PushDeadLetter      `index:"5" json:"dead_letters" xml:"dead_letters" yaml:"dead_letters" comment:"死信列表"`
	Audits        []*PushAudit           `index:"6" json:"audits" xml:"audits" yaml:"audits" comment:"审计列表"`
	Precise       bool                   `index:"7" json:"precise" xml:"precise" yaml:"precise" comment:"是否精确"`
}

type PushOverview struct {
	AppCount        int64        `index:"1" json:"app_count" xml:"app_count" yaml:"app_count" comment:"应用数量"`
	ChannelCount    int64        `index:"2" json:"channel_count" xml:"channel_count" yaml:"channel_count" comment:"通道数量"`
	DeviceCount     int64        `index:"3" json:"device_count" xml:"device_count" yaml:"device_count" comment:"设备数量"`
	TemplateCount   int64        `index:"4" json:"template_count" xml:"template_count" yaml:"template_count" comment:"模板数量"`
	TaskCount       int64        `index:"5" json:"task_count" xml:"task_count" yaml:"task_count" comment:"任务数量"`
	SuccessCount    int64        `index:"6" json:"success_count" xml:"success_count" yaml:"success_count" comment:"成功数量"`
	FailureCount    int64        `index:"7" json:"failure_count" xml:"failure_count" yaml:"failure_count" comment:"失败数量"`
	DeadLetterCount int64        `index:"8" json:"dead_letter_count" xml:"dead_letter_count" yaml:"dead_letter_count" comment:"死信数量"`
	RecentTasks     []*PushTask  `index:"9" json:"recent_tasks" xml:"recent_tasks" yaml:"recent_tasks" comment:"最近任务"`
	RecentAudits    []*PushAudit `index:"10" json:"recent_audits" xml:"recent_audits" yaml:"recent_audits" comment:"最近审计"`
	RiskSummaries   string       `index:"11" json:"risk_summaries" xml:"risk_summaries" yaml:"risk_summaries" comment:"风险摘要JSON"`
}
