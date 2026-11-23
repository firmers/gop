/*
 * Copyright (c) 2019, 2025, firmer.tech and/or its affiliates. All rights reserved.
 * Firmer Corporation PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 */

package kinds

type PushNativeCredential struct {
	Platform   string `index:"1" json:"platform" xml:"platform" yaml:"platform" comment:"原生平台：apple、android、harmony"`
	Ak         string `index:"2" json:"ak" xml:"ak" yaml:"ak" comment:"认证标识"`
	Sk         string `index:"3" json:"sk" xml:"sk" yaml:"sk" comment:"密钥，仅写入，留空保留"`
	Options    string `index:"4" json:"options" xml:"options" yaml:"options" comment:"扩展JSON"`
	Configured bool   `index:"5" json:"configured" xml:"configured" yaml:"configured" comment:"已配置密钥，只读"`
}

// PushProviderConfig describes authentication; an empty SK preserves the stored secret on edit.
type PushProviderConfig struct {
	Icon           string                  `index:"1" json:"icon" xml:"icon" yaml:"icon" comment:"供应商图标"`
	Ak             string                  `index:"2" json:"ak" xml:"ak" yaml:"ak" comment:"认证标识"`
	Sk             string                  `index:"3" json:"sk" xml:"sk" yaml:"sk" comment:"密钥，仅写入，编辑留空保留"`
	AuthType       string                  `index:"4" json:"auth_type" xml:"auth_type" yaml:"auth_type" comment:"认证方式：ak_sk、p8、pem、p12、service_account"`
	Options        string                  `index:"5" json:"options" xml:"options" yaml:"options" comment:"认证附加参数JSON"`
	Configured     bool                    `index:"6" json:"configured" xml:"configured" yaml:"configured" comment:"已配置密钥，只读"`
	Capabilities   []string                `index:"7" json:"capabilities" xml:"capabilities" yaml:"capabilities" comment:"推送能力多选数组"`
	NativeSettings []*PushNativeCredential `index:"8" json:"native_settings" xml:"native_settings" yaml:"native_settings" comment:"原生平台配置"`
}
