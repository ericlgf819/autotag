package main

import (
	"fmt"

	"github.com/ericlgf819/autotag/autotagbuilder"
)

func main() {
	// TODO... input/output path should be set
	err := autotagbuilder.BuildProcessor().Process("", "")

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
