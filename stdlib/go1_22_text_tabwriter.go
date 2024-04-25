// Code generated by 'yaegi extract text/tabwriter'. DO NOT EDIT.

//go:build go1.22
// +build go1.22

package stdlib

import (
	"go/constant"
	"go/token"
	"reflect"
	"text/tabwriter"
)

func init() {
	Symbols["text/tabwriter/tabwriter"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AlignRight":          reflect.ValueOf(tabwriter.AlignRight),
		"Debug":               reflect.ValueOf(tabwriter.Debug),
		"DiscardEmptyColumns": reflect.ValueOf(tabwriter.DiscardEmptyColumns),
		"Escape":              reflect.ValueOf(constant.MakeFromLiteral("255", token.INT, 0)),
		"FilterHTML":          reflect.ValueOf(tabwriter.FilterHTML),
		"NewWriter":           reflect.ValueOf(tabwriter.NewWriter),
		"StripEscape":         reflect.ValueOf(tabwriter.StripEscape),
		"TabIndent":           reflect.ValueOf(tabwriter.TabIndent),

		// type definitions
		"Writer": reflect.ValueOf((*tabwriter.Writer)(nil)),
	}
}