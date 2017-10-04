package main

import (
	"fmt"

	"github.com/charles-haynes/mergo"
)

func merge(dst, src interface{}, overwrite bool) error {
	if overwrite {
		return mergo.MergeWithOverwrite(dst, src)
	}
	fmt.Printf("merging\n%#v\n\nwith\n%#v\n\n", dst, src)
	err := mergo.Merge(dst, src)
	fmt.Printf("result\n%#v\n\n", dst)
	return err
}
