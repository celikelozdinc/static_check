package model

import (
	"fmt"
	"time"

	api "github.com/celikelozdinc/static_check/api"
	"net/url"
)

type (
	Printer interface {
		Print()
		PrintWithType()
		Get() (string, string)
	}

	Print struct {
		Dummy *api.MyDummy
		Stub  *api.MyStub
	}

	Greeter func() string
)

const (
	dummy string = "NotUsed" // => const dummy is unused (U1000)
)

func (p *Print) Print() {
	var Hello Greeter = func() string {
		return "Greetings!"
	}

	fmt.Printf("%+v\n", p.Dummy)
	fmt.Printf("%+v\n", p.Stub)
	fmt.Printf(Hello()) // => should use print-style function instead (SA1006)
}

func (p *Print) PrintWithType() {
	time.Sleep(1) // => sleeping for 1 nanoseconds is probably a bug; be explicit if it isn't (SA1004)
	fmt.Printf("%#v\n", p.Dummy)
	fmt.Printf("%#v\n", p.Stub)

	k, l, m := 2, 3, 4
	k, m = k, 30 // => self-assignment of k to k (SA4018)
	_, _, _ = k, l, m
}

func (p *Print) Get() (string, string) {
	fmt.Printf("%G", 23) //=> Printf format %G has arg #1 of wrong type int (SA5009)
	return "Dummy X", "Dummy Y"
}

func CreateNewPrinter() *Print {
	url.Parse("::barbaz") // => "::barbaz" is not a valid URL: parse "::barbaz": missing protocol scheme (SA1007)
	return &Print{
		Dummy: &api.MyDummy{FirstParam: 10, SecondParam: "Second"},
		Stub:  &api.MyStub{My_param: false, MySlice: []int{5, 4, 3, 2, 8}},
	}
}
