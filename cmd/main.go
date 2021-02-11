package main

import (
	pack "gitlab.com/celikelozdinc/static_check/pkg/model"
)

func main() {

	printer := pack.CreateNewPrinter()
	printer.Print()
	printer.PrintWithType()
}
