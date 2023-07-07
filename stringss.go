package main

import (
	"fmt"
	st "strings"
)

// aliaising -> renaming the package
var p = fmt.Println

func Stringss() {
	fmt.Println("hello")

	p("Contains", st.Contains("testing", "et"))
	p("Count", st.Count("testing", "t"))
	p("Touppercase", st.ToUpper("tEsTing"))
	p("Touppercase", st.ToLower("dsfsdfs"))
	p("Replace", st.Replace("foo", "o","0",-1))
	p("Split", st.Split("test-dev-prod", "-"))

}
