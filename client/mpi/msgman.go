/*
 * Copyright (c) 2019, 2026, firmer.tech and/or its affiliates. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mpi

import (
	"context"

	"github.com/firmers/gop/client/kinds"
	"github.com/firmers/mesh/client/golang/types"
)

// MsgMan spi
//
// Manages push resources, delivery operations, and push bus lifecycles.
//
//	推送资源、投递操作及推送总线生命周期管理服务接口.
//
// @SPI(macro.MeshMPI)
type MsgMan interface {
	// Overview returns management overview.
	//	查询推送管理总览.
	//	@i query 查询条件
	//	@o overview 管理总览
	// @MPI("gop.push.admin.overview")
	Overview(ctx context.Context, query *kinds.PushQuery) (*kinds.PushOverview, error)

	// IndexApp lists push apps.
	//	分页查询推送应用.
	//	@i idx 分页参数
	//	@o page 应用分页
	// @MPI("gop.push.admin.app.index")
	IndexApp(ctx context.Context, idx *types.Paging) (*types.Page[*kinds.PushApp], error)

	// GetApp gets a push app.
	//	查询推送应用详情.
	//	@i identity 应用标识
	//	@o app 应用详情
	// @MPI("gop.push.admin.app.get")
	GetApp(ctx context.Context, identity *kinds.PushIdentity) (*kinds.PushApp, error)

	// SaveApp saves a push app.
	//	保存推送应用.
	//	@i input 应用信息
	// @MPI("gop.push.admin.app.save")
	SaveApp(ctx context.Context, input *kinds.PushApp) error

	// RemoveApp removes a push app.
	//	删除推送应用.
	//	@i identity 应用标识
	// @MPI("gop.push.admin.app.remove")
	RemoveApp(ctx context.Context, identity *kinds.PushIdentity) error

	// IndexChannel lists push channels.
	//	分页查询推送通道.
	//	@i idx 分页参数
	//	@o page 通道分页
	// @MPI("gop.push.admin.channel.index")
	IndexChannel(ctx context.Context, idx *types.Paging) (*types.Page[*kinds.PushChannel], error)

	// GetChannel gets a push channel.
	//	查询推送通道详情.
	//	@i identity 通道标识
	//	@o channel 通道详情
	// @MPI("gop.push.admin.channel.get")
	GetChannel(ctx context.Context, identity *kinds.PushIdentity) (*kinds.PushChannel, error)

	// SaveChannel saves a push channel.
	//	保存推送通道.
	//	@i input 通道信息
	// @MPI("gop.push.admin.channel.save")
	SaveChannel(ctx context.Context, input *kinds.PushChannel) error

	// TestChannel sends a channel test notification.
	//	发送通道测试通知.
	//	@i request 测试请求
	//	@o result 测试结果
	// @MPI("gop.push.admin.channel.test")
	TestChannel(ctx context.Context, request *kinds.PushChannelTest) (*kinds.PushChannelTestResult, error)

	// RemoveChannel removes a push channel.
	//	删除推送通道.
	//	@i identity 通道标识
	// @MPI("gop.push.admin.channel.remove")
	RemoveChannel(ctx context.Context, identity *kinds.PushIdentity) error

	// IndexDevice lists managed devices.
	//	分页查询设备目录.
	//	@i idx 分页参数
	//	@o page 设备分页
	// @MPI("gop.push.admin.device.index")
	IndexDevice(ctx context.Context, idx *types.Paging) (*types.Page[*kinds.PushDevice], error)

	// GetDevice gets a device.
	//	查询设备详情.
	//	@i identity 设备标识
	//	@o device 设备详情
	// @MPI("gop.push.admin.device.get")
	GetDevice(ctx context.Context, identity *kinds.PushIdentity) (*kinds.PushDevice, error)

	// SaveDevice saves a device.
	//	保存设备信息.
	//	@i input 设备信息
	// @MPI("gop.push.admin.device.save")
	SaveDevice(ctx context.Context, input *kinds.PushDevice) error

	// DisableDevice disables a device.
	//	禁用设备.
	//	@i identity 设备标识
	// @MPI("gop.push.admin.device.disable")
	DisableDevice(ctx context.Context, identity *kinds.PushIdentity) error

	// IndexTemplate lists templates.
	//	分页查询推送模板.
	//	@i idx 分页参数
	//	@o page 模板分页
	// @MPI("gop.push.admin.template.index")
	IndexTemplate(ctx context.Context, idx *types.Paging) (*types.Page[*kinds.PushTemplate], error)

	// GetTemplate gets a template.
	//	查询推送模板详情.
	//	@i identity 模板标识
	//	@o template 模板详情
	// @MPI("gop.push.admin.template.get")
	GetTemplate(ctx context.Context, identity *kinds.PushIdentity) (*kinds.PushTemplate, error)

	// SaveTemplate saves a template.
	//	保存推送模板.
	//	@i input 模板信息
	// @MPI("gop.push.admin.template.save")
	SaveTemplate(ctx context.Context, input *kinds.PushTemplate) error

	// PublishTemplate publishes a template.
	//	发布推送模板.
	//	@i request 发布请求
	// @MPI("gop.push.admin.template.publish")
	PublishTemplate(ctx context.Context, request *kinds.PushPublishRequest) error

	// RemoveTemplate removes a template.
	//	删除推送模板.
	//	@i identity 模板标识
	// @MPI("gop.push.admin.template.remove")
	RemoveTemplate(ctx context.Context, identity *kinds.PushIdentity) error

	// IndexAudience lists audiences.
	//	分页查询推送人群.
	//	@i idx 分页参数
	//	@o page 人群分页
	// @MPI("gop.push.admin.audience.index")
	IndexAudience(ctx context.Context, idx *types.Paging) (*types.Page[*kinds.PushAudience], error)

	// GetAudience gets an audience.
	//	查询推送人群详情.
	//	@i identity 人群标识
	//	@o audience 人群详情
	// @MPI("gop.push.admin.audience.get")
	GetAudience(ctx context.Context, identity *kinds.PushIdentity) (*kinds.PushAudience, error)

	// SaveAudience saves an audience.
	//	保存推送人群.
	//	@i input 人群信息
	// @MPI("gop.push.admin.audience.save")
	SaveAudience(ctx context.Context, input *kinds.PushAudience) error

	// RemoveAudience removes an audience.
	//	删除推送人群.
	//	@i identity 人群标识
	// @MPI("gop.push.admin.audience.remove")
	RemoveAudience(ctx context.Context, identity *kinds.PushIdentity) error

	// IndexRoutePolicy lists route policies.
	//	分页查询路由策略.
	//	@i idx 分页参数
	//	@o page 路由策略分页
	// @MPI("gop.push.admin.route.index")
	IndexRoutePolicy(ctx context.Context, idx *types.Paging) (*types.Page[*kinds.PushRoutePolicy], error)

	// GetRoutePolicy gets a route policy.
	//	查询路由策略详情.
	//	@i identity 路由策略标识
	//	@o policy 路由策略详情
	// @MPI("gop.push.admin.route.get")
	GetRoutePolicy(ctx context.Context, identity *kinds.PushIdentity) (*kinds.PushRoutePolicy, error)

	// SaveRoutePolicy saves a route policy.
	//	保存路由策略.
	//	@i input 路由策略
	// @MPI("gop.push.admin.route.save")
	SaveRoutePolicy(ctx context.Context, input *kinds.PushRoutePolicy) error

	// RemoveRoutePolicy removes a route policy.
	//	删除路由策略.
	//	@i identity 路由策略标识
	// @MPI("gop.push.admin.route.remove")
	RemoveRoutePolicy(ctx context.Context, identity *kinds.PushIdentity) error

	// IndexQuotaPolicy lists quota policies.
	//	分页查询配额策略.
	//	@i idx 分页参数
	//	@o page 配额策略分页
	// @MPI("gop.push.admin.quota.index")
	IndexQuotaPolicy(ctx context.Context, idx *types.Paging) (*types.Page[*kinds.PushQuotaPolicy], error)

	// GetQuotaPolicy gets a quota policy.
	//	查询配额策略详情.
	//	@i identity 配额策略标识
	//	@o policy 配额策略详情
	// @MPI("gop.push.admin.quota.get")
	GetQuotaPolicy(ctx context.Context, identity *kinds.PushIdentity) (*kinds.PushQuotaPolicy, error)

	// SaveQuotaPolicy saves a quota policy.
	//	保存配额策略.
	//	@i input 配额策略
	// @MPI("gop.push.admin.quota.save")
	SaveQuotaPolicy(ctx context.Context, input *kinds.PushQuotaPolicy) error

	// RemoveQuotaPolicy removes a quota policy.
	//	删除配额策略.
	//	@i identity 配额策略标识
	// @MPI("gop.push.admin.quota.remove")
	RemoveQuotaPolicy(ctx context.Context, identity *kinds.PushIdentity) error

	// IndexWebhook lists webhooks.
	//	分页查询Webhook.
	//	@i idx 分页参数
	//	@o page Webhook分页
	// @MPI("gop.push.admin.webhook.index")
	IndexWebhook(ctx context.Context, idx *types.Paging) (*types.Page[*kinds.PushWebhook], error)

	// GetWebhook gets a webhook.
	//	查询Webhook详情.
	//	@i identity Webhook标识
	//	@o webhook Webhook详情
	// @MPI("gop.push.admin.webhook.get")
	GetWebhook(ctx context.Context, identity *kinds.PushIdentity) (*kinds.PushWebhook, error)

	// SaveWebhook saves a webhook.
	//	保存Webhook.
	//	@i input Webhook信息
	// @MPI("gop.push.admin.webhook.save")
	SaveWebhook(ctx context.Context, input *kinds.PushWebhook) error

	// RemoveWebhook removes a webhook.
	//	删除Webhook.
	//	@i identity Webhook标识
	// @MPI("gop.push.admin.webhook.remove")
	RemoveWebhook(ctx context.Context, identity *kinds.PushIdentity) error

	// Audit lists push audits.
	//	分页查询推送审计.
	//	@i idx 分页参数
	//	@o page 审计分页
	// @MPI("gop.push.admin.audit")
	Audit(ctx context.Context, idx *types.Paging) (*types.Page[*kinds.PushAudit], error)

	// Components lists available push bus components.
	//	查询推送总线可用组件.
	//	@i query 组件查询条件
	//	@o registry 组件注册目录
	// @MPI("gop.push.bus.admin.components")
	Components(ctx context.Context, query *kinds.PushBusComponentQuery) (*kinds.PushBusComponentRegistry, error)

	// ComponentResources lists resources available to a bus component.
	//	分页查询总线组件可用资源.
	//	@i query 组件资源查询条件
	//	@o page 组件资源分页
	// @MPI("gop.push.bus.admin.component.resources")
	ComponentResources(ctx context.Context, query *kinds.PushBusComponentResourceQuery) (*types.Page[*kinds.PushBusComponentResource], error)

	// OverviewBus returns the push bus management overview.
	//	查询推送总线管理总览.
	//	@i query 总线查询条件
	//	@o overview 总线管理总览
	// @MPI("gop.push.bus.admin.overview")
	OverviewBus(ctx context.Context, query *kinds.PushBusQuery) (*kinds.PushBusOverview, error)

	// IndexBus lists push buses.
	//	分页查询推送总线.
	//	@i idx 分页参数
	//	@o page 总线分页
	// @MPI("gop.push.bus.admin.index")
	IndexBus(ctx context.Context, idx *types.Paging) (*types.Page[*kinds.PushBus], error)

	// GetBus gets a push bus.
	//	查询推送总线详情.
	//	@i identity 总线标识
	//	@o bus 总线详情
	// @MPI("gop.push.bus.admin.get")
	GetBus(ctx context.Context, identity *kinds.PushBusIdentity) (*kinds.PushBus, error)

	// SaveBus saves a push bus.
	//	保存推送总线.
	//	@i input 总线信息
	// @MPI("gop.push.bus.admin.save")
	SaveBus(ctx context.Context, input *kinds.PushBus) error

	// RemoveBus removes a push bus.
	//	删除推送总线.
	//	@i identity 总线标识
	// @MPI("gop.push.bus.admin.remove")
	RemoveBus(ctx context.Context, identity *kinds.PushBusIdentity) error

	// GetDraft gets a push bus draft.
	//	查询推送总线草稿.
	//	@i identity 总线标识
	//	@o draft 总线草稿
	// @MPI("gop.push.bus.admin.draft.get")
	GetDraft(ctx context.Context, identity *kinds.PushBusIdentity) (*kinds.PushBusVersion, error)

	// SaveDraft saves a push bus draft.
	//	保存推送总线草稿.
	//	@i input 草稿保存请求
	//	@o draft 总线草稿
	// @MPI("gop.push.bus.admin.draft.save")
	SaveDraft(ctx context.Context, input *kinds.PushBusDraftSave) (*kinds.PushBusVersion, error)

	// ValidateDraft validates a push bus draft.
	//	校验推送总线草稿.
	//	@i input 草稿校验请求
	//	@o validation 草稿校验结果
	// @MPI("gop.push.bus.admin.draft.validate")
	ValidateDraft(ctx context.Context, input *kinds.PushBusDraftRequest) (*kinds.PushBusValidation, error)

	// TestDraft dry-runs a push bus draft.
	//	试运行推送总线草稿.
	//	@i input 草稿试运行请求
	//	@o result 草稿试运行结果
	// @MPI("gop.push.bus.admin.draft.test")
	TestDraft(ctx context.Context, input *kinds.PushBusDryRunRequest) (*kinds.PushBusDryRunResult, error)

	// Publish publishes a push bus version.
	//	发布推送总线版本.
	//	@i input 版本发布请求
	//	@o version 已发布版本
	// @MPI("gop.push.bus.admin.publish")
	Publish(ctx context.Context, input *kinds.PushBusPublishRequest) (*kinds.PushBusVersion, error)

	// Rollback rolls back a push bus version.
	//	回滚推送总线版本.
	//	@i input 版本回滚请求
	//	@o version 回滚后的版本
	// @MPI("gop.push.bus.admin.rollback")
	Rollback(ctx context.Context, input *kinds.PushBusRollbackRequest) (*kinds.PushBusVersion, error)

	// Versions lists push bus versions.
	//	分页查询推送总线版本.
	//	@i idx 分页参数
	//	@o page 总线版本分页
	// @MPI("gop.push.bus.admin.versions")
	Versions(ctx context.Context, idx *types.Paging) (*types.Page[*kinds.PushBusVersion], error)

	// GetVersion gets a push bus version.
	//	查询推送总线版本详情.
	//	@i identity 总线版本标识
	//	@o version 总线版本详情
	// @MPI("gop.push.bus.admin.version.get")
	GetVersion(ctx context.Context, identity *kinds.PushBusIdentity) (*kinds.PushBusVersion, error)

	// DiffVersion compares push bus versions.
	//	比较推送总线版本差异.
	//	@i input 版本比较请求
	//	@o diff 版本差异
	// @MPI("gop.push.bus.admin.version.diff")
	DiffVersion(ctx context.Context, input *kinds.PushBusVersionDiffRequest) (*kinds.PushBusVersionDiff, error)

	// Runs lists push bus runs.
	//	分页查询推送总线运行记录.
	//	@i idx 分页参数
	//	@o page 总线运行记录分页
	// @MPI("gop.push.bus.admin.runs")
	Runs(ctx context.Context, idx *types.Paging) (*types.Page[*kinds.PushBusRun], error)

	// GetRun gets a push bus run.
	//	查询推送总线运行详情.
	//	@i identity 总线运行标识
	//	@o run 总线运行详情
	// @MPI("gop.push.bus.admin.run.get")
	GetRun(ctx context.Context, identity *kinds.PushBusIdentity) (*kinds.PushBusRun, error)

	// AuditBus lists push bus audits.
	//	分页查询推送总线审计.
	//	@i idx 分页参数
	//	@o page 总线审计分页
	// @MPI("gop.push.bus.admin.audit")
	AuditBus(ctx context.Context, idx *types.Paging) (*types.Page[*kinds.PushAudit], error)

	// RegisterDevice registers a push device.
	//	注册推送设备.
	//	@i input 设备注册信息
	//	@o device 设备信息
	// @MPI("gop.push.runtime.device.register")
	RegisterDevice(ctx context.Context, input *kinds.PushDeviceRegistration) (*kinds.PushDevice, error)

	// UnregisterDevice unregisters a push device.
	//	注销推送设备.
	//	@i identity 设备标识
	// @MPI("gop.push.runtime.device.unregister")
	UnregisterDevice(ctx context.Context, identity *kinds.PushIdentity) error

	// Send sends a push notification.
	//	发送单条推送.
	//	@i input 发送请求
	//	@o result 发送结果
	// @MPI("gop.push.runtime.send")
	Send(ctx context.Context, input *kinds.PushSendRequest) (*kinds.PushSendResult, error)

	// BatchSend submits a batch push task.
	//	提交批量推送任务.
	//	@i input 批量发送请求
	//	@o task 推送任务
	// @MPI("gop.push.runtime.send.batch")
	BatchSend(ctx context.Context, input *kinds.PushBatchSendRequest) (*kinds.PushTask, error)

	// SubmitTask submits a push task.
	//	提交推送任务.
	//	@i input 任务提交请求
	//	@o task 推送任务
	// @MPI("gop.push.runtime.task.submit")
	SubmitTask(ctx context.Context, input *kinds.PushTaskSubmit) (*kinds.PushTask, error)

	// PauseTask pauses a push task.
	//	暂停推送任务.
	//	@i control 任务控制参数
	// @MPI("gop.push.runtime.task.pause")
	PauseTask(ctx context.Context, control *kinds.PushTaskControl) error

	// ResumeTask resumes a push task.
	//	恢复推送任务.
	//	@i control 任务控制参数
	// @MPI("gop.push.runtime.task.resume")
	ResumeTask(ctx context.Context, control *kinds.PushTaskControl) error

	// CancelTask cancels a push task.
	//	取消推送任务.
	//	@i control 任务控制参数
	// @MPI("gop.push.runtime.task.cancel")
	CancelTask(ctx context.Context, control *kinds.PushTaskControl) error

	// ExecuteTask executes a push task.
	//	执行推送任务.
	//	@i identity 任务标识
	//	@o result 任务执行结果
	// @MPI("gop.push.runtime.task.execute")
	ExecuteTask(ctx context.Context, identity *kinds.PushIdentity) (*kinds.PushTaskResult, error)

	// Dispatch executes scheduled push tasks.
	//	调度执行到期推送任务.
	//	@i window 调度窗口
	//	@o report 调度报告
	// @MPI("gop.push.runtime.dispatch")
	Dispatch(ctx context.Context, window *kinds.PushDispatchWindow) (*kinds.PushDispatchReport, error)

	// Trace returns push delivery trace.
	//	查询推送链路.
	//	@i query 链路查询
	//	@o trace 推送链路
	// @MPI("gop.push.runtime.trace")
	Trace(ctx context.Context, query *kinds.PushQuery) (*kinds.PushTrace, error)

	// Tasks lists runtime push tasks.
	//	分页查询运行态推送任务.
	//	@i idx 分页参数
	//	@o page 任务分页
	// @MPI("gop.push.runtime.tasks")
	Tasks(ctx context.Context, idx *types.Paging) (*types.Page[*kinds.PushTask], error)

	// Notifications lists runtime push notifications.
	//	分页查询运行态推送通知.
	//	@i idx 分页参数
	//	@o page 通知分页
	// @MPI("gop.push.runtime.notifications")
	Notifications(ctx context.Context, idx *types.Paging) (*types.Page[*kinds.PushNotification], error)

	// Attempts lists runtime push delivery attempts.
	//	分页查询运行态投递尝试.
	//	@i idx 分页参数
	//	@o page 投递尝试分页
	// @MPI("gop.push.runtime.attempts")
	Attempts(ctx context.Context, idx *types.Paging) (*types.Page[*kinds.PushDeliveryAttempt], error)

	// Receipts lists push receipts.
	//	分页查询推送回执.
	//	@i idx 分页参数
	//	@o page 回执分页
	// @MPI("gop.push.runtime.receipts")
	Receipts(ctx context.Context, idx *types.Paging) (*types.Page[*kinds.PushReceipt], error)

	// Retry retries a push task or notification.
	//	重试推送任务或通知.
	//	@i request 重试请求
	//	@o task 推送任务
	// @MPI("gop.push.runtime.retry")
	Retry(ctx context.Context, request *kinds.PushRetryRequest) (*kinds.PushTask, error)

	// DeadLetters lists push dead letters.
	//	分页查询推送死信.
	//	@i idx 分页参数
	//	@o page 死信分页
	// @MPI("gop.push.runtime.dead_letters")
	DeadLetters(ctx context.Context, idx *types.Paging) (*types.Page[*kinds.PushDeadLetter], error)

	// TraceBus returns a push bus execution trace.
	//	查询推送总线执行链路.
	//	@i query 总线链路查询条件
	//	@o trace 总线执行链路
	// @MPI("gop.push.bus.runtime.trace")
	TraceBus(ctx context.Context, query *kinds.PushBusTraceQuery) (*kinds.PushBusTrace, error)

	// RetryBus retries a push bus run.
	//	重试推送总线运行.
	//	@i input 总线重试请求
	//	@o run 总线运行记录
	// @MPI("gop.push.bus.runtime.retry")
	RetryBus(ctx context.Context, input *kinds.PushBusRetryRequest) (*kinds.PushBusRun, error)
}
