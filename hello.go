package hello

import (
	"embed"
	"fmt"
)

//go:embed xml/*
var Content embed.FS

func Aa() {
	fmt.Println("Hello, world.")
}
