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

func TestRoot1Child(t *testing.T) {
	expect := `
┌ root
└── Child
`

	tree := render(Node("root").Append(Node("Child")))

	actually.Got(tree).Expect(expect).Same(t)
}