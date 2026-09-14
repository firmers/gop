/*
 * Copyright (c) 2019, 2025, firmer.tech and/or its affiliates. All rights reserved.
 * Firmer Corporation PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 */

package kinds

import "github.com/firmers/mesh/client/golang/types"

const (
	InboxUnread int64 = 0
	InboxRead   int64 = 1
)

// InboxMessage is a message addressed to the authenticated user, not a delivery audit.
type InboxMessage struct {
	Id         string     `index:"0" json:"id" xml:"id" yaml:"id" comment:"站内信ID"`
	Title      string     `index:"1" json:"title" xml:"title" yaml:"title" comment:"消息标题，纯文本"`
	Body       string     `index:"2" json:"body" xml:"body" yaml:"body" comment:"消息正文，纯文本，不执行HTML或脚本"`
	Source     string     `index:"3" json:"source" xml:"source" yaml:"source" comment:"消息来源显示名称"`
	ReadStatus int64      `index:"4" json:"read_status" xml:"read_status" yaml:"read_status" comment:"阅读状态：0未读、1已读，与投递状态无关"`
	CreateAt   types.Time `index:"5" json:"create_at" xml:"create_at" yaml:"create_at" comment:"收件时间"`
}

// InboxCounts describes the whole active mailbox, independent of filters or pagination.
type InboxCounts struct {
	Total  int64 `index:"0" json:"total" xml:"total" yaml:"total" comment:"当前用户全部有效站内信数量，不受搜索和分页影响"`
	Unread int64 `index:"1" json:"unread" xml:"unread" yaml:"unread" comment:"当前用户未读站内信数量，不受搜索和分页影响"`
}
