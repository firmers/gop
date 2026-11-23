/*
 * Copyright (c) 2019, 2025, firmer.tech and/or its affiliates.
 * SPDX-License-Identifier: Apache-2.0
 */

package kinds

import (
	"encoding/json"
	"time"

	"github.com/firmers/mesh/client/golang/types"
)

// MarshalJSON represents the absence of an unbinding event as null, not a date.
func (device PushDevice) MarshalJSON() ([]byte, error) {
	type fields PushDevice
	var unregistered *types.Time
	if !time.Time(device.UnregisterAt).IsZero() {
		unregistered = &device.UnregisterAt
	}
	return json.Marshal(struct {
		*fields
		UnregisterAt *types.Time `json:"unregister_at"`
	}{fields: (*fields)(&device), UnregisterAt: unregistered})
}

func (device *PushDevice) UnmarshalJSON(data []byte) error {
	type fields PushDevice
	var value struct {
		fields
		UnregisterAt *types.Time `json:"unregister_at"`
	}
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*device = PushDevice(value.fields)
	if value.UnregisterAt != nil {
		device.UnregisterAt = *value.UnregisterAt
	}
	return nil
}
