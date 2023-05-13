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

// RenderAsText renders a node tree as text, and output.
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
		n.Prop = &prop{
			Name: n.Prop.Name,
		}
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
			line += strings.Repeat(opt.HorizontalLink, opt.HorizontalLinkLen+1)
		}
	}

	return line + n.buildIcon(opt) + spaces(opt.NamePaddingLeftLen) + n.buildName(opt) + n.buildTag(opt) + stringNewLine
}

func (n *N) buildIcon(opt *RenderTextOptions) string {
	if !isBlank(n.Prop.Icon) {
		if n.Prop.IconColor != nil {
			return n.Prop.IconColor.Sprint(n.Prop.Icon)
		} else if opt.GlobalIconColor != nil {
			return opt.GlobalIconColor.Sprint(n.Prop.Icon)
		} else {
			return n.Prop.Icon
		}
	} else if !isBlank(opt.GlobalIcon) {
		if n.Prop.IconColor != nil {
			return n.Prop.IconColor.Sprint(opt.GlobalIcon)
		} else if opt.GlobalIconColor != nil {
			return opt.GlobalIconColor.Sprint(opt.GlobalIcon)
		} else {
			return opt.GlobalIcon
		}
	}

	return ""
}

func (n *N) buildName(opt *RenderTextOptions) string {
	if !isBlank(n.Prop.Name) {
		if n.Prop.NameColor != nil {
			return n.RenderName()
		} else if opt.GlobalNameColor != nil {
			return opt.GlobalNameColor.Sprint(n.Prop.Name)
		} else {
			return n.Prop.Name
		}
	}

	return ""
}

func (n *N) buildTag(opt *RenderTextOptions) string {
	if !isBlank(n.Prop.Tag) {
		tagString := fmt.Sprintf(opt.TagFormat, n.Prop.Tag)
		if n.Prop.TagColor != nil {
			return n.Prop.TagColor.Sprint(tagString)
		} else if opt.GlobalTagColor != nil {
			return opt.GlobalTagColor.Sprint(tagString)
		} else {
			return tagString
		}
	} else if !isBlank(opt.GlobalTag) {
		tagString := fmt.Sprintf(opt.TagFormat, opt.GlobalTag)
		if n.Prop.TagColor != nil {
			return n.Prop.TagColor.Sprint(tagString)
		} else if opt.GlobalTagColor != nil {
			return opt.GlobalTagColor.Sprint(tagString)
		} else {
			return tagString
		}
	}

	return ""
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
		if n.Prop.DescriptionColor != nil {
			description = n.Prop.DescriptionColor.Sprint(description)
		} else if opt.GlobalDescriptionColor != nil {
			description = opt.GlobalDescriptionColor.Sprint(description)
		}
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
	} else if n.depth() > 1 || !n.hasChild() {
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
		return strings.Repeat(line+stringNewLine, opt.ChildrenMarginTop)
	}

	return ""
}

func (n *N) buildNodeBottomMargin(opt *RenderTextOptions) string {
	if !n.hasChild() && n.isLast && opt.ChildrenMarginBottom > 0 {
		line := n.buildAncestorBranchesLine(&n.ancestors, opt)
		return strings.Repeat(line, opt.ChildrenMarginBottom) + stringNewLine
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
