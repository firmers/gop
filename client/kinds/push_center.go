/*
 * Copyright (c) 2019, 2025, firmer.tech and/or its affiliates. All rights reserved.
 * Firmer Corporation PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 */

package kinds

// PushChannelState changes provider availability without replacing configuration.
// 通道启停请求，不替换服务商配置或凭据。
type PushChannelState struct {
	Id      string `index:"1" json:"id" xml:"id" yaml:"id" comment:"Channel ID 通道ID"`
	AppId   string `index:"2" json:"app_id" xml:"app_id" yaml:"app_id" comment:"App ID 应用ID"`
	Enabled bool   `index:"3" json:"enabled" xml:"enabled" yaml:"enabled" comment:"Enabled 是否启用"`
	Reason  string `index:"4" json:"reason" xml:"reason" yaml:"reason" comment:"Operation reason 操作原因"`
}
