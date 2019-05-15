package main

import (
	"fmt"

	b "github.com/codysoyland/gomod-b"
	c "github.com/codysoyland/gomod-c"
)

func main() {
	fmt.Println("Hello from A")
	b.Hello()
	c.Hello()
}
