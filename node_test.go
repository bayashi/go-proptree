package proptree

import (
	"testing"

	a "github.com/bayashi/actually"
	"github.com/fatih/color"
)

func TestNode(t *testing.T) {
	n := Node("foo")
	a.Got(n).Expect(&N{Prop: &prop{Name: "foo"}, isLast: true}).Same(t)
	a.Got(n.Prop.Name).Expect("foo").Same(t)
	a.Got(n.hasDescription()).False(t)

	n.Icon("*")
	a.Got(n.Prop.Icon).Expect("*").Same(t)

	n.Tag("tag")
	a.Got(n.Prop.Tag).Expect("tag").Same(t)

	n.Description("desc")
	a.Got(n.hasDescription()).True(t)
	a.Got(n.Prop.Descriptions[0]).Expect("desc").Same(t)

	n.Descriptions([]string{"desc2", "desc3"})
	a.Got(n.Prop.Descriptions[1]).Expect("desc2").Same(t)
	a.Got(n.Prop.Descriptions[2]).Expect("desc3").Same(t)
}

func TestNodeColor(t *testing.T) {
	red := color.FgHiRed
	blue := color.FgBlue

	n := Node("foo", red)
	a.Got(n).Expect(&N{Prop: &prop{Name: "foo", NameColor: color.New(red)}, isLast: true}).Same(t)
	n.NameColor(blue)
	a.Got(n.Prop.NameColor).Expect(color.New(blue)).Same(t)

	n.Icon("*", red)
	a.Got(n.Prop.IconColor).Expect(color.New(red)).Same(t)
	n.IconColor(blue)
	a.Got(n.Prop.IconColor).Expect(color.New(blue)).Same(t)

	n.Tag("tag", red)
	a.Got(n.Prop.TagColor).Expect(color.New(red)).Same(t)
	n.TagColor(blue)
	a.Got(n.Prop.TagColor).Expect(color.New(blue)).Same(t)

	n.Description("description", red)
	a.Got(n.Prop.DescriptionColor).Expect(color.New(red)).Same(t)
	n.DescriptionColor(blue)
	a.Got(n.Prop.DescriptionColor).Expect(color.New(blue)).Same(t)
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
