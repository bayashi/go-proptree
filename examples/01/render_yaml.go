package main

import (
	"fmt"
	"io"

	pt "github.com/bayashi/go-proptree"
	yaml "gopkg.in/yaml.v3"
)

func RenderAsYAML(w io.Writer, n *pt.N) error {
	err := yaml.NewEncoder(w).Encode(n)
	if err != nil {
		return fmt.Errorf("could not encode as YAML: %s", err)
	}

	return nil
}
