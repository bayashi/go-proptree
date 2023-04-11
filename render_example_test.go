package proptree_test

import (
	"os"

	pt "github.com/bayashi/go-proptree"
)

func ExampleRenderAsText() {
	n := tree()
	n.RenderAsText(os.Stdout)
	// Output:
	// ┌ Version History
	// │  This is version history of Fake Software.
	// │  Life is full of ups and downs.
	// └─┬ 1.0
	//   ├─┬ 1.1
	//   │ ├──! 1.1.1
	//   │ └──* 1.1.2
	//   ├─┬ 1.2
	//   │ └──* 1.2.1
	//   ├─┬ 1.3
	//   │ │  Implemented GUI from this version.
	//   │ ├── 1.3.1
	//   │ ├─┬ 1.3.2
	//   │ │ ├── 1.3.2.1
	//   │ │ └──* 1.3.2.2
	//   │ └── 1.3.3: Stable
	//   ├──! 1.4
	//   └──* 1.5: Newest
}

func tree() *pt.N {
	tree := pt.Node("Version History").
			Description("This is version history of Fake Software.").
			Description("Life is full of ups and downs.")

	node1 := pt.Node("1.0")
	tree.Append(node1)

	node11 := pt.Node("1.1")
	node1.Append(node11)
	node11.Append(pt.Node("1.1.1").Icon("!"))
	node11.Append(pt.Node("1.1.2").Icon("*"))

	node12 := pt.Node("1.2")
	node1.Append(node12)
	node12.Append(pt.Node("1.2.1").Icon("*"))

	node13 := pt.Node("1.3").Description("Implemented GUI from this version.")
	node1.Append(node13)
	node13.Append(pt.Node("1.3.1"))
	node132 := pt.Node("1.3.2")
	node13.Append(node132)
	node132.Append(pt.Node("1.3.2.1"))
	node132.Append(pt.Node("1.3.2.2").Icon("*"))
	node13.Append(pt.Node("1.3.3").Tag("Stable"))

	node1.Append(pt.Node("1.4").Icon("!"))
	node1.Append(pt.Node("1.5").Icon("*").Tag("Newest"))

	return tree
}
