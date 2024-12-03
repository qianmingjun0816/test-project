package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/qianmingjun0816/test-project/softwares"
)

func main() {
	fmt.Println(errors.New("xxx"))
	softwares.Test()
}
