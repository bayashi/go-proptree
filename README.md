# go-proptree

<a href="https://github.com/bayashi/go-proptree/blob/main/LICENSE"><img src="https://img.shields.io/badge/LICENSE-MIT-GREEN.png"></a>
<a href="https://github.com/bayashi/go-proptree/actions"><img src="https://github.com/bayashi/go-proptree/workflows/main/badge.svg?_t=1681289447"/></a>
<a href="https://goreportcard.com/report/github.com/bayashi/go-proptree" title="go-proptree report card" target="_blank"><img src="https://goreportcard.com/badge/github.com/bayashi/go-proptree" alt="go-proptree report card"></a>
<a href="https://pkg.go.dev/github.com/bayashi/go-proptree"><img src="https://pkg.go.dev/badge/github.com/bayashi/go-proptree.svg" alt="Go Reference"></a>

`go-proptree` provides a text tree view of nesting nodes with properties.

Building tree data structure easily, and you can add Icon, Tag and Description properties on each node.

## Overview

This is example code of `examples/01/main.go` in this repository.

```go
package main

import (
	"os"
	pt "github.com/bayashi/go-proptree"
)

func main() {
	tree := pt.Node("Root").Icon("*").Tag("tag")
	child := pt.Node("Child A").
		Description("This is a description about Child.").
		Description("You can set multiple lines.")
	tree.Append(child).
		Append(pt.Node("Child B").Description("This is a description about Child B."))
	child.Append(pt.Node("Grandchild").Icon("@"))
	tree.RenderAsText(os.Stdout)
}
```

It renders:

    $ go run examples/01/main.go
    
    ┌* Root: tag
    ├─┬ Child A
    │ │  This is a description about Child.
    │ │  You can set multiple lines.
    │ └──@ Grandchild
    └── Child B
         This is a description about Child B.
    

Another tree from `render_example_test.go`.

    ┌ Version History
    │  This is version history of Fake Software.
    │  Life is full of ups and downs.
    └─┬ 1.0
      ├─┬ 1.1
      │ ├──! 1.1.1
      │ └──* 1.1.2
      ├─┬ 1.2
      │ └──* 1.2.1
      ├─┬ 1.3
      │ │  Implemented GUI from this version.
      │ ├── 1.3.1
      │ ├─┬ 1.3.2
      │ │ ├── 1.3.2.1
      │ │ └──* 1.3.2.2
      │ └── 1.3.3: Stable
      ├──! 1.4
      └──* 1.5: Newest

You can add color to each node like below. See `examples/02/main.go`

![colorized tree](https://user-images.githubusercontent.com/42190/238113810-df59e086-b580-43a8-8c2f-3eda1641cb3a.png)


## Installation

    go get github.com/bayashi/go-proptree

## License

MIT License

## Author

Dai Okabayashi: https://github.com/bayashi

## Special Thanks

Most of the logic was copied from https://github.com/plouc/textree
