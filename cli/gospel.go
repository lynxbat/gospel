package main

import (
	"fmt"

	"github.com/lynxbat/gospel/truths"
)

func main() {
	fmt.Println("Gospel v0.0.1")

	t := truths.New()
	t.List()
}
