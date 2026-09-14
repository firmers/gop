/*
 * Copyright (c) 2019, 2025, firmer.tech and/or its affiliates. All rights reserved.
 * Firmer Corporation PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 */

package mpi

import (
	"context"

	"github.com/firmers/gop/client/kinds"
	"github.com/firmers/mesh/client/golang/types"
)

// Inbox spi
//
//	当前登录用户的站内信收件箱。所有方法从服务端Session获取用户，拒绝未登录调用；不接受调用方指定收件用户。
//
// @SPI(macro.MeshMPI)
type Inbox interface {
	// Index lists messages belonging to the authenticated user.
	//	分页查询本人站内信，按收件时间和ID倒序排列，不自动标记已读。
	//	@i idx index从0开始，limit默认20且最多100；factor.keyword按标题、正文、来源做文本包含查询；factor.unread_only为true时仅返回未读
	//	@o content 筛选后的消息分页；忽略order，不接受user_id等用户身份覆盖参数
	// @MPI("gop.inbox.index")
	Index(ctx context.Context, idx *types.Paging) (*types.Page[*kinds.InboxMessage], error)

	// Get reads a message without changing its read status.
	//	读取本人消息详情；他人消息与不存在消息返回相同的未找到错误，不改变阅读状态。
	//	@i id 消息ID
	//	@o content 消息详情
	// @MPI("gop.inbox.get")
	Get(ctx context.Context, id string) (*kinds.InboxMessage, error)

	// Counts returns totals for the authenticated user's entire mailbox.
	//	返回本人全部有效消息数和未读数，与搜索、标签及当前页无关。
	//	@o content 收件箱统计
	// @MPI("gop.inbox.counts")
	Counts(ctx context.Context) (*kinds.InboxCounts, error)

	// Read idempotently marks one owned message as read.
	//	事务内标记本人消息已读；重复操作成功，他人消息与不存在消息均返回未找到。
	//	@i id 消息ID
	// @MPI("gop.inbox.read")
	Read(ctx context.Context, id string) error

	// ReadAll marks the authenticated user's entire mailbox as read.
	//	事务内一键标记本人当前全部未读消息，不受搜索和分页限制；不修改其他用户消息，重复调用成功。
	// @MPI("gop.inbox.read.all")
	ReadAll(ctx context.Context) error
}
