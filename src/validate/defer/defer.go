package _defer

import (
	"errors"
	"fmt"
)

func Defer() (err error) {
	defer func() {
		fmt.Println(err)
		err = errors.New("a")
	}()
	defer func(e error) {
		fmt.Println(e)
	}(err)
	return errors.New("c")
}
