package main

import (
	pack "github.com/celikelozdinc/static_check/pkg/model"
)

func main() {

	printer := pack.CreateNewPrinter()
	printer.Print()
	printer.PrintWithType()
}
