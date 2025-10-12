// Copyright 2025 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package omitempty is a custom linter to be used by
// golangci-lint to check that omitempty is used if and only if
// the field is a pointer type or composite type (slice, map, etc.).
package omitempty

import (
	"encoding/json"
	"go/ast"
	"reflect"
	"strings"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("omitempty", New)
}

// Config holds the configuration for the omitempty linter.
type Config struct {
	// Unnecessary enables checking that value types should not use omitempty.
	// Reports "omitempty is unnecessary" for value types with omitempty.
	// Default: true
	Unnecessary bool `json:"unnecessary"`
	// Missing enables checking that pointer types should use omitempty.
	// Reports "omitempty is missing" for pointer types without omitempty.
	// Default: true
	Missing bool `json:"missing"`
}

// OmitemptyPlugin is a custom linter plugin for golangci-lint.
type OmitemptyPlugin struct {
	config Config
}

// New returns an analysis.Analyzer to use with golangci-lint.
func New(conf any) (register.LinterPlugin, error) {
	// Default configuration: both checks enabled
	config := Config{
		Unnecessary: true,
		Missing:     true,
	}

	// Parse configuration if provided
	if conf != nil {
		// The configuration comes as a map[string]any from golangci-lint
		if confMap, ok := conf.(map[string]any); ok {
			// Convert to JSON and back to properly parse into our Config struct
			if jsonBytes, err := json.Marshal(confMap); err == nil {
				_ = json.Unmarshal(jsonBytes, &config)
			}
		}
	}

	return &OmitemptyPlugin{config: config}, nil
}

// BuildAnalyzers builds the analyzers for the OmitemptyPlugin.
func (o *OmitemptyPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		{
			Name: "omitempty",
			Doc:  "Reports incorrect usage of omitempty in JSON tags.",
			Run: func(pass *analysis.Pass) (any, error) {
				return run(pass, o.config)
			},
		},
	}, nil
}

// GetLoadMode returns the load mode for the OmitemptyPlugin.
func (o *OmitemptyPlugin) GetLoadMode() string {
	return register.LoadModeSyntax
}

func run(pass *analysis.Pass, config Config) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if n == nil {
				return false
			}

			switch t := n.(type) {
			case *ast.TypeSpec:
				if structType, ok := t.Type.(*ast.StructType); ok {
					checkStructFields(structType, pass, config)
				}
			}

			return true
		})
	}
	return nil, nil
}

func checkStructFields(structType *ast.StructType, pass *analysis.Pass, config Config) {
	for _, field := range structType.Fields.List {
		if field.Tag == nil {
			continue
		}

		tag := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
		jsonTag := tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		hasOmitempty := strings.Contains(jsonTag, "omitempty")
		typeCategory := getTypeCategory(field.Type)

		switch typeCategory {
		case typeValue:
			if config.Unnecessary && hasOmitempty {
				pass.Reportf(field.Pos(), "omitempty is unnecessary for value types")
			}
		case typePointer:
			if config.Missing && !hasOmitempty {
				pass.Reportf(field.Pos(), "omitempty is missing for pointer types")
			}
		case typeOther:
			// No restrictions on other types (slice, map, interface, chan, etc.)
		}
	}
}

const (
	typeValue   = iota // Value types that we care about checking (basic types and named structs)
	typePointer        // Pointer types (*string, *int, etc.)
	typeOther          // Other types we don't check (slice, map, interface, chan, any, etc.)
)

func getTypeCategory(expr ast.Expr) int {
	switch t := expr.(type) {
	case *ast.StarExpr:
		// Pointer type: *T
		return typePointer
	case *ast.Ident:
		// All identifier types (int, string, bool, CustomType, any, error, etc.)
		return typeValue
	case *ast.SelectorExpr:
		// Qualified types like time.Time, pkg.Type
		// Only exclude json.RawMessage
		if x, ok := t.X.(*ast.Ident); ok && x.Name == "json" && t.Sel.Name == "RawMessage" {
			return typeOther
		}
		return typeValue
	case *ast.StructType:
		// Inline struct
		return typeValue
	default:
		return typeOther
	}
}
