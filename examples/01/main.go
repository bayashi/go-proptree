package main

import (
	"os"

	pt "github.com/bayashi/proptree"
)

// This is for README example

func main() {
	tree := pt.Node("Root").Icon("*").Tag("tag")

	child := pt.Node("Child A").
		Description("This is a description about Child.").
		Description("You can set multiple lines.")

	tree.Append(child).
		Append(pt.Node("Child B").Description("This is a description about Child B."))

	child.Append(pt.Node("Grandchild").Icon("@"))

	tree.RenderAsText(os.Stdout)
	//
    // ┌* Root: tag
    // ├─┬ Child A
    // │ │  This is a description about Child.
    // │ │  You can set multiple lines.
    // │ └──@ Grandchild
    // └── Child B
    //      This is a description about Child B.
	//

	// Examples for another type of rendering
	//RenderAsJSON(os.Stdout, tree)
	//RenderAsYAML(os.Stdout, tree)
}
