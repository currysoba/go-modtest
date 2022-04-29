package pub

import (
	"github.com/currysoba/go-modtest/internal"
	"github.com/currysoba/go-modtest/sub"
)

func PublicFunc() string {
	return "PublicFunc"
}

func PublicFunc2() string {
	return "PublicFunc2 " + sub.SubFunc() + " " + internal.InternalFunc()
}
