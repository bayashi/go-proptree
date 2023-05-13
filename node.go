package proptree

import (
	"github.com/fatih/color"
)

type prop struct {
	Name             string `json:"name,omitempty" yaml:"name,omitempty"`
	NameColor        *color.Color
	Icon             string `json:"icon,omitempty" yaml:"icon,omitempty"`
	IconColor        *color.Color
	Tag              string `json:"tag,omitempty" yaml:"tag,omitempty"`
	TagColor         *color.Color
	Descriptions     []string `json:"description,omitempty" yaml:"description,omitempty"`
	DescriptionColor *color.Color
}

type N struct {
	Prop      *prop `json:"prop,omitempty" yaml:"prop,omitempty"`
	Children  []*N  `json:"children,omitempty" yaml:"children,omitempty"`
	ancestors []*N
	isLast    bool
}

func Node(name string, c ...color.Attribute) *N {
	n := &N{
		isLast: true,
		Prop: &prop{
			Name: name,
		},
	}

	if len(c) > 0 {
		n.Prop.NameColor = color.New(c...)
	}

	return n
}

func (n *N) NameColor(c ...color.Attribute) *N {
	n.Prop.NameColor = color.New(c...)

	return n
}

func (n *N) RenderName() string {
	if n.Prop.NameColor != nil {
		return n.Prop.NameColor.Sprint(n.Prop.Name)
	}

	return n.Prop.Name
}

func (n *N) Icon(icon string, c ...color.Attribute) *N {
	n.Prop.Icon = icon

	if len(c) > 0 {
		n.Prop.IconColor = color.New(c...)
	}

	return n
}

func (n *N) IconColor(c ...color.Attribute) *N {
	n.Prop.IconColor = color.New(c...)

	return n
}

func (n *N) iconLen() int {
	return len(n.Prop.Icon)
}

func (n *N) Tag(tag string, c ...color.Attribute) *N {
	n.Prop.Tag = tag

	if len(c) > 0 {
		n.Prop.TagColor = color.New(c...)
	}

	return n
}

func (n *N) TagColor(c ...color.Attribute) *N {
	n.Prop.TagColor = color.New(c...)

	return n
}

func (n *N) Description(description string, c ...color.Attribute) *N {
	return n.Descriptions([]string{description}, c...)
}

func (n *N) Descriptions(descriptions []string, c ...color.Attribute) *N {
	n.Prop.Descriptions = append(n.Prop.Descriptions, descriptions...)

	if len(c) > 0 {
		n.Prop.DescriptionColor = color.New(c...)
	}

	return n
}

func (n *N) DescriptionColor(c ...color.Attribute) *N {
	n.Prop.DescriptionColor = color.New(c...)

	return n
}

func (n *N) hasDescription() bool {
	return len(n.Prop.Descriptions) > 0
}

func (n *N) hasChild() bool {
	return len(n.Children) > 0
}

func (n *N) isRoot() bool {
	return len(n.ancestors) == 0
}

func (n *N) depth() int {
	return len(n.ancestors)
}

func (n *N) Append(c *N) *N {
	if len(c.ancestors) != 0 {
		panic("Node already having a parent")
	}

	for _, child := range n.Children {
		child.isLast = false
	}

	c.ancestors = append(n.ancestors, c.ancestors...)
	c.ancestors = append([]*N{n}, c.ancestors...)
	c.isLast = true

	n.Children = append(n.Children, c)

	return n
}
