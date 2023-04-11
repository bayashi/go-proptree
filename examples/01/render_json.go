package main

import (
	"encoding/json"
	"fmt"
	"io"

	pt "github.com/bayashi/go-proptree"
)

func RenderAsJSON(w io.Writer, n *pt.N) error {
	err := json.NewEncoder(w).Encode(n)
	if err != nil {
		return fmt.Errorf("could not encode as JSON: %s", err)
	}

	return nil
}
