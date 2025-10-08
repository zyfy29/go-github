// Copyright 2025 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "time"

type TestStruct struct {
	// Value types without omitempty - correct
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	Score     float64   `json:"score"`
	Count     int64     `json:"count"`
	CreatedAt time.Time `json:"created_at"`

	// Pointer/composite types with omitempty - correct
	Title       *string           `json:"title,omitempty"`
	Description *string           `json:"description,omitempty"`
	Age         *int              `json:"age,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	Data        []byte            `json:"data,omitempty"`

	// Fields without JSON tags - should be ignored
	Internal string

	// JSON tag with "-" - should be ignored
	Ignored string `json:"-"`
}

func main() {}
