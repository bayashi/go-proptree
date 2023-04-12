package proptree

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

		MarginTop:            1,
		MarginLeft:           0,
		ChildrenMarginTop:    0,
		ChildrenMarginBottom: 0,
		HorizontalLinkLen:    1,
		NamePaddingLeftLen:   1,

		TurnOffProp: false,
	}
}

func (o *RenderTextOptions) Relax() *RenderTextOptions {
	o.MarginTop            = 1
	o.MarginLeft           = 1
	o.ChildrenMarginTop    = 1
	o.ChildrenMarginBottom = 1
	o.HorizontalLinkLen    = 3
	o.NamePaddingLeftLen   = 1

	return o
}
