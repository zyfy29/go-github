// Copyright 2025 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package omitempty is a custom linter to be used by
// golangci-lint to check that omitempty is used if and only if
// the field is a pointer type or composite type (slice, map, etc.).
package omitempty

import (
	"go/ast"
	"reflect"
	"strings"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("omitempty", New)
}

// OmitemptyPlugin is a custom linter plugin for golangci-lint.
type OmitemptyPlugin struct{}

// New returns an analysis.Analyzer to use with golangci-lint.
func New(_ any) (register.LinterPlugin, error) {
	return &OmitemptyPlugin{}, nil
}

// BuildAnalyzers builds the analyzers for the OmitemptyPlugin.
func (o *OmitemptyPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		{
			Name: "omitempty",
			Doc:  "Reports incorrect usage of omitempty in JSON tags.",
			Run:  run,
		},
	}, nil
}

// GetLoadMode returns the load mode for the OmitemptyPlugin.
func (o *OmitemptyPlugin) GetLoadMode() string {
	return register.LoadModeSyntax
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if n == nil {
				return false
			}

			switch t := n.(type) {
			case *ast.TypeSpec:
				if structType, ok := t.Type.(*ast.StructType); ok {
					checkStructFields(structType, pass)
				}
			}

			return true
		})
	}
	return nil, nil
}

func checkStructFields(structType *ast.StructType, pass *analysis.Pass) {
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

		var fieldName string
		if len(field.Names) > 0 {
			fieldName = field.Names[0].Name
		} else {
			// Embedded field
			fieldName = getTypeName(field.Type)
		}

		switch typeCategory {
		case typeValue:
			if hasOmitempty {
				pass.Reportf(field.Pos(), "field %s: value type should not use omitempty", fieldName)
			}
		case typePointer:
			if !hasOmitempty {
				pass.Reportf(field.Pos(), "field %s: pointer type should use omitempty", fieldName)
			}
		case typeComposite:
			// No restrictions on composite types (slice, map, etc.)
		}
	}
}

type typeCategory int

const (
	typeValue     typeCategory = iota // Value types (int, string, bool, time.Time, etc.)
	typePointer                       // Pointer types (*string, *int, etc.)
	typeComposite                     // Composite types ([]string, map[string]string, etc.)
)

func getTypeCategory(expr ast.Expr) typeCategory {
	switch t := expr.(type) {
	case *ast.StarExpr:
		// Pointer type
		return typePointer
	case *ast.ArrayType:
		// Slice or array
		return typeComposite
	case *ast.MapType:
		// Map
		return typeComposite
	case *ast.InterfaceType:
		// Interface
		return typeComposite
	case *ast.ChanType:
		// Channel
		return typeComposite
	case *ast.Ident:
		// Check for basic types
		basicTypes := map[string]bool{
			"int": true, "int8": true, "int16": true, "int32": true, "int64": true,
			"uint": true, "uint8": true, "uint16": true, "uint32": true, "uint64": true,
			"float32": true, "float64": true,
			"string": true, "bool": true,
			"byte": true, "rune": true,
			"complex64": true, "complex128": true,
		}
		if basicTypes[t.Name] {
			return typeValue
		}
		// Named types (structs, etc.) - treat as value type
		return typeValue
	case *ast.SelectorExpr:
		// Qualified identifier (e.g., time.Time) - treat as value type
		return typeValue
	case *ast.StructType:
		// Inline struct - treat as value type
		return typeValue
	default:
		// Unknown type, assume value type
		return typeValue
	}
}

func getTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.StarExpr:
		return getTypeName(t.X)
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return getTypeName(t.X) + "." + t.Sel.Name
	default:
		return "unknown"
	}
}
