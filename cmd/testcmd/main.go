package main

import (
	"fmt"

	root "github.com/currysoba/go-modtest"
	"github.com/currysoba/go-modtest/pub"
)

func main() {
	fmt.Println(root.RootFunc())
	fmt.Println(pub.PublicFunc())
}
