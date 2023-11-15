package hello

import (
	"embed"
	"fmt"
)

//go:embed xml/*
var Content embed.FS

func main() {
	fmt.Println("Hello, world.")
}
