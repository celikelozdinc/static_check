package model

import (
	"fmt"

	api "github.com/celikelozdinc/static_check/api"
)

type (
	Printer interface {
		PrintWithType()
	}

	Print struct {
		dummy *api.MyDummy
		stub  *api.MyStub
	}
)

func (p *Print) Print() {
	fmt.Printf("%+v\n", p.dummy)
	fmt.Printf("%+v\n", p.stub)
}

func (p *Print) PrintWithType() {
	fmt.Printf("%#v\n", p.dummy)
	fmt.Printf("%#v\n", p.stub)
}

func CreateNewPrinter() *Print {
	return &Print{
		dummy: &api.MyDummy{FirstParam: 10, SecondParam: "Second"},
		stub:  &api.MyStub{My_param: false},
	}
}
