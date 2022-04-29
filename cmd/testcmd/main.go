package main

import (
	"fmt"

	root "github.com/currysoba/go-modtest"
	"github.com/currysoba/go-modtest/internal"
	"github.com/currysoba/go-modtest/pub"
	"github.com/currysoba/go-modtest/sub"
)

func main() {
	fmt.Println(root.RootFunc())
	fmt.Println(pub.PublicFunc())
	fmt.Println(pub.PublicFunc2())
	fmt.Println(sub.SubFunc())
	fmt.Println(internal.InternalFunc())
}
