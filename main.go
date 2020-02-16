package main

import "fmt"

var (
	version  string
	revision string
)

func main() {
	fmt.Printf("goreleaser-work %s (%s)", version, revision)
}
