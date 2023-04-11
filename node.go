package proptree

type prop struct {
	Icon         string   `json:"icon,omitempty" yaml:"icon,omitempty"`
	Tag          string   `json:"tag,omitempty" yaml:"tag,omitempty"`
	Descriptions []string `json:"description,omitempty" yaml:"description,omitempty"`
}

type N struct {
	Name      string  `json:"name,omitempty" yaml:"name,omitempty"`
	Prop      *prop   `json:"prop,omitempty" yaml:"prop,omitempty"`
	Children  []*N    `json:"children,omitempty" yaml:"children,omitempty"`
	ancestors []*N
	isLast    bool
}

func Node(name string) *N {
	return &N{
		Name: name,
		isLast: true,
	}
}

func (n *N) Icon(icon string) *N {
	if n.Prop == nil {
		n.Prop = &prop{Icon: icon}
	} else {
		n.Prop.Icon = icon
	}

	return n
}

func (n *N) iconLen() int {
	if n.Prop == nil {
		return 0
	} else {
		return len(n.Prop.Icon)
	}
}

func (n *N) Tag(tag string) *N {
	if n.Prop == nil {
		n.Prop = &prop{Tag: tag}
	} else {
		n.Prop.Tag = tag
	}

	return n
}

func (n *N) Description(description string) *N {
	if n.Prop == nil {
		n.Prop = &prop{Descriptions: []string{description}}
	} else {
		n.Prop.Descriptions = append(n.Prop.Descriptions, description)
	}

	return n
}

func (n *N) hasDescription() bool {
	if n.Prop == nil {
		return false
	} else {
		return len(n.Prop.Descriptions) > 0
	}
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
