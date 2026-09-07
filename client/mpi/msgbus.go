/*
 * Copyright (c) 2019, 2025, firmer.tech and/or its affiliates. All rights reserved.
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

// MsgBus spi
// MsgBus exposes message delivery and receipt queries to business callers.
//
//	面向业务调用方的消息总线，只提供发推、回执查询及已发送消息分页检索。
//
// @SPI(macro.MeshMPI)
type MsgBus interface {
	// Push sends a message through a published bus workflow.
	//	按照 bus_id 执行已发布流程，使用幂等键防止重复提交。
	//	@i input  BizId业务幂等键、BusId总线、Ctx业务参数、Recipient接收对象、ActiveAt计划时间及CreateBy提交人
	//	@o content  JobId推送任务唯一ID，可用于Stats查询，不代表设备已送达
	// @MPI("gop.msgbus.push")
	Push(ctx context.Context, input *kinds.Message) (string, error)

	// Stats queries business delivery receipts by business or job identity.
	//	按业务编号或推送任务ID查询业务回执，与MsgCallback使用相同状态语义，不暴露投递实现细节。
	//	@i query  BizId业务编号或JobId推送任务ID，至少填写一个；同时填写时必须指向同一业务请求
	//	@o content  业务回执集合，包含接收对象、明确状态、备注及事件时间，无回执时返回空集合
	// @MPI("gop.msgbus.stats")
	Stats(ctx context.Context, query *kinds.PushTraceQuery) ([]*kinds.MessageReceipt, error)

	// Index retrieves a page of business push requests visible to the caller.
	//	分页检索当前业务调用方可见的推送请求，按业务请求计数，不返回内部通知、通道或执行审计。
	//	@i idx  分页及业务筛选条件，index为从零开始的页码，limit为每页条数
	//	@o content  业务消息分页，包含BizId、JobId、BusId、接收对象回执集合、整体状态和计划及受理时间
	// @MPI("gop.msgbus.index")
	Index(ctx context.Context, idx *types.Paging) (*types.Page[*kinds.MessageRecord], error)
}

// MsgCallback spi
// MsgCallback is implemented by business systems to receive message receipts.
//
//	由业务系统实现，推送平台调用，用于通知业务回执；不是极光或设备向推送平台提交回执的入口。
//
// @SPI(macro.MeshMPI)
type MsgCallback interface {
	// Receipt notifies the business system of a message delivery event.
	//	推送平台向业务系统通知回执，参数与Stats结果一致；业务方应幂等处理重复事件，受理不能视为送达。
	//	@i receipt BizId、JobId、业务接收对象及类型、int64回执状态码（1已受理、2已送达、4已打开、8失败、16结果未知）、Memo备注及事件时间
	// @MPI("${mesh.name}.msgbus.receipt")
	Receipt(ctx context.Context, receipt *kinds.MessageReceipt) error
}
