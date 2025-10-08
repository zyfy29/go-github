// Copyright 2025 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "time"

type TestStruct struct {
	// Value types with omitempty - should report
	ID        int       `json:"id,omitempty"`         // want `field ID: value type should not use omitempty`
	Name      string    `json:"name,omitempty"`       // want `field Name: value type should not use omitempty`
	Active    bool      `json:"active,omitempty"`     // want `field Active: value type should not use omitempty`
	Score     float64   `json:"score,omitempty"`      // want `field Score: value type should not use omitempty`
	Count     int64     `json:"count,omitempty"`      // want `field Count: value type should not use omitempty`
	CreatedAt time.Time `json:"created_at,omitempty"` // want `field CreatedAt: value type should not use omitempty`

	// Pointer types without omitempty - should report
	Title       *string `json:"title"`       // want `field Title: pointer type should use omitempty`
	Description *string `json:"description"` // want `field Description: pointer type should use omitempty`
	Age         *int    `json:"age"`         // want `field Age: pointer type should use omitempty`

	// Composite types (slice, map) - no restrictions, these should NOT report
	Tags     []string          `json:"tags"`
	Metadata map[string]string `json:"metadata"`
	Data     []byte            `json:"data"`
}

func main() {}
