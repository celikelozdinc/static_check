package main

import (
	"fmt"
	"sort"

	pack "github.com/celikelozdinc/static_check/pkg/model"
)

func main() {

	printer := pack.CreateNewPrinter()
	printer.Print()
	printer.PrintWithType()

	if !!printer.Stub.My_param { // => negating a boolean twice has no effect; is this a typo? (SA4013)
		fmt.Println("We negated a boolean twice and we got T")
	} else {
		fmt.Println("We negated a boolean twice and we got F")
	}

	//var custom [5]int
	sort.Slice([]int{5}, nil)
	//sort.Slice(custom, nil) // => sort.Slice must only be called on slices, was called on [5]int (SA1028)

	x1, y := printer.Get() // => this value of x1 is never used (SA4006)
	fmt.Println(y)
	x1, _ = printer.Get()
	fmt.Println(x1)

}
