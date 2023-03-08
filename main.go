package main

import (
	"fmt"

	say "github.com/bmahmood/go-private-module-repo"
)

func main() {
	result := say.Hello()
	fmt.Println(result)
}
