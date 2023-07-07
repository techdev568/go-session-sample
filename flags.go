package main

import (
	"flag"
	"fmt"
)

// flag will parse the commandline arguments
// init() -> execute before main
// init() -> initilization of data we will do in init()

var env *string
var port *int

func init() {
	fmt.Println("init")
	// varname = flag.String(varName, defaultvalue, msg)
	env = flag.String("env", "development", "a string")
	port = flag.Int("port", 3000, "an int")
}

func Flags() {
	flag.Parse()
	fmt.Println("Hello")
	fmt.Println("env", *env)
	fmt.Println("port", *port)
}
