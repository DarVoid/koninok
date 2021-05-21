package main

import (
	"fmt"
	"time"
)

func main()  {
	hello:="Hello peeps"
	hope:="I've been just researching what folder structure I should use for GO projects"
	andThis:="and this is what I came up with :D (with some help from uncle Google)"
	came:="Hope you have a GOod one ;)"
	fmt.Printf("%v\n", hello)
	time.Sleep( time.Second)
	fmt.Printf("%v\n", hope)
	time.Sleep(2 * time.Second)
	fmt.Printf("%v\n", andThis)
	time.Sleep(1 * time.Second)
	fmt.Printf("%v\n", came)


}