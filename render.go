package proptree

import (
	"fmt"
	"io"
	"strings"
)

const (
	stringSpace   = " "
	stringNewLine = "\n"
)

func (n *N) RenderAsText(w io.Writer, opts ...*RenderTextOptions) error {
	opt := settleOpt(opts...)

	tree := ""
	n.growTree(&tree, opt)
	_, err := fmt.Fprint(w, tree)

	return err
}

func settleOpt(opts ...*RenderTextOptions) *RenderTextOptions {
	opt := RenderTextDefaultOptions()
	if opts != nil && opts[0] != nil {
		opt = opts[0]
	}

	opt.marginLeftString = spaces(opt.MarginLeft)

	return opt
}

func (n *N) growTree(treeLine *string, opt *RenderTextOptions) {
	if opt.TurnOffProp {
		n.Prop = nil
	}

	if n.isRoot() && opt.MarginTop > 0 {
		*treeLine += strings.Repeat(stringNewLine, opt.MarginTop)
	}

	*treeLine += n.buildNodeNameLine(opt) + n.buildDescriptions(opt) + n.buildNodeTopMargin(opt)

	for _, c := range n.Children {
		c.growTree(treeLine, opt)
	}

	*treeLine += n.buildNodeBottomMargin(opt)
}

func (n *N) buildNodeNameLine(opt *RenderTextOptions) string {
	var line string
	if n.isRoot() {
		line = opt.marginLeftString + opt.RootLink
	} else {
		line = n.buildAncestorBranchesLine(&n.ancestors, opt)
		if n.isLast {
			line += opt.LastChildLink
		} else {
			line += opt.ChildLink
		}
		if n.hasChild() {
			line += strings.Repeat(opt.HorizontalLink, opt.HorizontalLinkLen) + opt.ChildrenLink
		} else {
			line += strings.Repeat(opt.HorizontalLink, opt.HorizontalLinkLen + 1)
		}
	}

	if n.Prop != nil && !isBlank(n.Prop.Icon) {
		line += n.Prop.Icon
	} else if !isBlank(opt.GlobalIcon) {
		line += opt.GlobalIcon
	}

	line += spaces(opt.NamePaddingLeftLen) + n.Name

	if n.Prop != nil && !isBlank(n.Prop.Tag) {
		line += fmt.Sprintf(opt.TagFormat, n.Prop.Tag)
	} else if !isBlank(opt.GlobalTag) {
		line += fmt.Sprintf(opt.TagFormat, opt.GlobalTag)
	}

	return line + stringNewLine
}

func (n *N) buildDescriptions(opt *RenderTextOptions) string {
	var descriptions []string
	if n.hasDescription() {
		descriptions = n.Prop.Descriptions
	} else if len(opt.GlobalDescriptions) > 0 {
		descriptions = opt.GlobalDescriptions
	} else {
		return ""
	}

	line := ""
	for _, description := range descriptions {
		line += n.buildStringBeforeDescription(opt) + description + stringNewLine
	}

	return line
}

func (n *N) buildStringBeforeDescription(opt *RenderTextOptions) string {
	var line string
	if n.hasChild() {
		line = n.buildAncestorBranchesLine(&n.Children[0].ancestors, opt)
		if n.depth() == 1 {
			line += opt.VerticalLink
			if !n.isLast {
				line += stringSpace
			}
		}
	} else {
		line = n.buildAncestorBranchesLine(&n.ancestors, opt)
	}

	if n.isRoot() {
		line += opt.VerticalLink
	} else if n.isLast {
		line += stringSpace
	} else if n.depth() > 1 {
		line += opt.VerticalLink
	}

	if n.hasChild() {
		if n.depth() == 1 {
			line += spaces(n.iconLen() + opt.NamePaddingLeftLen)
		} else {
			line += spaces(n.iconLen() + opt.NamePaddingLeftLen + 1)
		}
	} else {
		line += spaces(opt.HorizontalLinkLen + n.iconLen() + opt.NamePaddingLeftLen + 2)
	}

	return line
}

func (n *N) buildNodeTopMargin(opt *RenderTextOptions) string {
	if n.hasChild() && opt.ChildrenMarginTop > 0 {
		line := n.buildAncestorBranchesLine(&n.Children[0].ancestors, opt)
		line += opt.VerticalLink
		return strings.Repeat(line + stringNewLine, opt.ChildrenMarginTop)
	}

	return ""
}

func (n *N) buildNodeBottomMargin(opt *RenderTextOptions ) string {
	if !n.hasChild() && n.isLast && opt.ChildrenMarginBottom > 0 {
		line := n.buildAncestorBranchesLine(&n.ancestors, opt)
		return strings.Repeat(line, opt.ChildrenMarginBottom)
	}

	return ""
}

func (n *N) buildAncestorBranchesLine(ancestors *[]*N, opt *RenderTextOptions) string {
	line := opt.marginLeftString
	for _, ancestor := range reverseNodes(*ancestors) {
		if ancestor.isRoot() {
			continue
		}
		if ancestor.isLast {
			line += spaces(opt.HorizontalLinkLen + 1)
		} else {
			line += opt.VerticalLink
			line += spaces(opt.HorizontalLinkLen)
		}
	}

	return line
}

func reverseNodes(Nodes []*N) []*N {
	reversed := []*N{}
	for i := len(Nodes) - 1; i >= 0; i-- {
		reversed = append(reversed, Nodes[i])
	}

	return reversed
}

func spaces(count int) string {
	return strings.Repeat(stringSpace, count)
}

func isBlank(str string) bool {
	return str == ""
}
