package proptree

import "github.com/fatih/color"

// RenderTextOptions is a struct for rendering text.
type RenderTextOptions struct {
	// symbols
	HorizontalLink string
	VerticalLink   string
	RootLink       string
	ChildLink      string
	LastChildLink  string
	ChildrenLink   string

	// properties
	GlobalIcon         string
	GlobalTag          string
	GlobalDescriptions []string
	TagFormat          string

	// color
	GlobalNameColor        *color.Color
	GlobalIconColor        *color.Color
	GlobalTagColor         *color.Color
	GlobalDescriptionColor *color.Color

	// dimensions
	MarginTop            int
	MarginLeft           int
	ChildrenMarginTop    int
	ChildrenMarginBottom int
	HorizontalLinkLen    int
	NamePaddingLeftLen   int

	// behavior
	TurnOffProp bool

	// internal
	marginLeftString string
}

// RenderTextDefaultOptions is a dataset of rendering options by default.
func RenderTextDefaultOptions() *RenderTextOptions {
	return &RenderTextOptions{
		HorizontalLink: "─",
		VerticalLink:   "│",
		RootLink:       "┌",
		ChildLink:      "├",
		LastChildLink:  "└",
		ChildrenLink:   "┬",

		GlobalIcon:         "",
		GlobalTag:          "",
		GlobalDescriptions: []string{},
		TagFormat:          `: %s`,

		GlobalNameColor:        nil,
		GlobalIconColor:        nil,
		GlobalTagColor:         nil,
		GlobalDescriptionColor: nil,

		MarginTop:            1,
		MarginLeft:           0,
		ChildrenMarginTop:    0,
		ChildrenMarginBottom: 0,
		HorizontalLinkLen:    1,
		NamePaddingLeftLen:   1,

		TurnOffProp: false,
	}
}

// Relax is a optional data adjuster for relaxed view.
func (o *RenderTextOptions) Relax() *RenderTextOptions {
	o.MarginTop            = 1
	o.MarginLeft           = 1
	o.ChildrenMarginTop    = 1
	o.ChildrenMarginBottom = 1
	o.HorizontalLinkLen    = 3
	o.NamePaddingLeftLen   = 1

	return o
}

// SetGlobalNameColor is a setter to set GlobalNameColor
func (o *RenderTextOptions) SetGlobalNameColor(c ...color.Attribute) *RenderTextOptions {
	o.GlobalNameColor = color.New(c...)

	return o
}

// SetGlobalIconColor is a setter to set GlobalIconColor
func (o *RenderTextOptions) SetGlobalIconColor(c ...color.Attribute) *RenderTextOptions {
	o.GlobalIconColor = color.New(c...)

	return o
}

// SetGlobalTagColor is a setter to set GlobalTagColor
func (o *RenderTextOptions) SetGlobalTagColor(c ...color.Attribute) *RenderTextOptions {
	o.GlobalTagColor = color.New(c...)

	return o
}

// SetGlobalDescriptionColor is a setter to set GlobalDescriptionColor
func (o *RenderTextOptions) SetGlobalDescriptionColor(c ...color.Attribute) *RenderTextOptions {
	o.GlobalDescriptionColor = color.New(c...)

	return o
}
