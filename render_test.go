package proptree

import (
	"bytes"
	"testing"

	"github.com/bayashi/actually"
)

func render(n *N) string {
	buf := &bytes.Buffer{}
	n.RenderAsText(buf)

	return buf.String()
}

func TestRoot(t *testing.T) {
	expect := `
┌ root
`
	tree := render(Node("root"))
	actually.Got(tree).Expect(expect).Same(t)
}

func TestRootChild(t *testing.T) {
	expect := `
┌ root
└── Child
`
	tree := render(Node("root").Append(Node("Child")))
	actually.Got(tree).Expect(expect).Same(t)
}

func TestRootChildGChild(t *testing.T) {
	expect := `
┌ root
└─┬ Child
  └── G-Child
`
	tree := render(Node("root").Append(Node("Child").Append(Node("G-Child"))))
	actually.Got(tree).Expect(expect).Same(t)
}

func TestRootChildWithProp(t *testing.T) {
	expect := `
┌* root: tag
│   root description
└──+ Child: tag
      Child description
`
	root  := Node("root").Icon("*").Tag("tag").Description("root description")
	child := Node("Child").Icon("+").Tag("tag").Description("Child description")
	tree := render(root.Append(child))
	actually.Got(tree).Expect(expect).Same(t)
}

func TestRootChildGChildWithProp(t *testing.T) {
	expect := `
┌* root: tag
│   root description
└─┬+ Child: tag
  │   Child description
  └──$ G-Child: tag
        G-Child description
`
	root   := Node("root").Icon("*").Tag("tag").Description("root description")
	child  := Node("Child").Icon("+").Tag("tag").Description("Child description")
	gchild := Node("G-Child").Icon("$").Tag("tag").Description("G-Child description")
	tree := root.Append(child.Append(gchild))
	actually.Got(render(tree)).Expect(expect).Same(t)
}

func TestRoot2ChildrenGChildWithProp(t *testing.T) {
	expect := `
┌* root: tag
│   root description
├─┬+ Child: tag
│ │   Child description
│ └──$ G-Child: tag
│       G-Child description
└──+ Child2: tag
      Child2 description
`
	root   := Node("root").Icon("*").Tag("tag").Description("root description")
	child  := Node("Child").Icon("+").Tag("tag").Description("Child description")
	gchild := Node("G-Child").Icon("$").Tag("tag").Description("G-Child description")
	child2 := Node("Child2").Icon("+").Tag("tag").Description("Child2 description")
	tree := root.Append(child.Append(gchild)).Append(child2)
	actually.Got(render(tree)).Expect(expect).Same(t)
}

func TestRoot3ChildrenGChildWithProp(t *testing.T) {
	expect := `
┌* root: tag
│   root description
├─┬+ Child: tag
│ │   Child description
│ └──$ G-Child: tag
│       G-Child description
├──+ Child2: tag
│     Child2 description
└──+ Child3: tag
      Child3 description
`
	root   := Node("root").Icon("*").Tag("tag").Description("root description")
	child  := Node("Child").Icon("+").Tag("tag").Description("Child description")
	gchild := Node("G-Child").Icon("$").Tag("tag").Description("G-Child description")
	child2 := Node("Child2").Icon("+").Tag("tag").Description("Child2 description")
	child3 := Node("Child3").Icon("+").Tag("tag").Description("Child3 description")
	tree := root.Append(child.Append(gchild)).Append(child2).Append(child3)
	actually.Got(render(tree)).Expect(expect).Same(t)
}
