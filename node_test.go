package proptree

import (
	"testing"

	a "github.com/bayashi/actually"
)

func TestNode(t *testing.T) {
	n := Node("foo")
	a.Got(n).Expect(&N{ Name:"foo", isLast:true }).Same(t)
	a.Got(n.Name).Expect("foo").Same(t)
	a.Got(n.hasDescription()).False(t)

	n.Icon("*")
	a.Got(n.Prop.Icon).Expect("*").Same(t)

	n.Tag("tag")
	a.Got(n.Prop.Tag).Expect("tag").Same(t)

	n.Description("desc")
	a.Got(n.hasDescription()).True(t)
	a.Got(n.Prop.Descriptions[0]).Expect("desc").Same(t)
}

func TestNodeAppend(t *testing.T) {
	n := Node("foo")
	a.Got(n.hasChild()).False(t)
	a.Got(n.isRoot()).True(t)
	a.Got(n.depth()).Expect(0).Same(t)
	a.Got(len(n.Children)).Expect(0).Same(t)

	m := Node("bar")
	a.Got(len(m.ancestors)).Expect(0).Same(t)
	n.Append(m)
	a.Got(len(n.Children)).Expect(1).Same(t)
	a.Got(len(m.ancestors)).Expect(1).Same(t)
}
