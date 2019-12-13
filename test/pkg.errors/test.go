package main

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

func a() error {
	return b()
}

func b() error {
	return c()
}

func c() error {
	// return errors.Errorf("hello")
	// return errors.Wrap(io.EOF, "hi")
	return errors.WithStack(io.EOF)
}

func main() {
	fmt.Println(a().Error())
	fmt.Printf("%+v\n", a())
}
