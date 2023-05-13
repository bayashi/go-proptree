package main

import (
	"os"

	pt "github.com/bayashi/go-proptree"
	c "github.com/fatih/color"
)

// Colored tree example

func main() {
	tree := pt.Node("Root", c.FgHiGreen, c.Bold).
		Icon("*", c.FgHiCyan).
		Tag("tag")

	child := pt.Node("Child A").
		Description("This is a description about Child.").
		Description("You can set multiple lines.")

	tree.Append(child).
		Append(pt.Node("Child B").Description("This is a description about Child B.", c.FgHiWhite))

	child.Append(pt.Node("Grandchild").Icon("@"))

	opt := pt.RenderTextDefaultOptions()
	opt.SetGlobalNameColor(c.FgCyan)
	opt.SetGlobalIconColor(c.FgHiYellow)
	opt.SetGlobalTagColor(c.FgHiRed)
	opt.SetGlobalDescriptionColor(c.FgWhite)
	tree.RenderAsText(os.Stdout, opt)
}
