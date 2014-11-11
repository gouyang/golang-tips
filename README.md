golang-tips
===========

#Basics

##variable

- byte // alias for uint8
- rune // alias for int32

###shorthand

// constant

    const HELLO string = "HELLO, WORLD"

// int

    i := 10

// string

    s := "hello world"

// array

    a := [3]array{1, 2, 3}
    a := [...]array{1, 2, 3}

// slice

    s := []slice{1, 2, 3}

// map

    m := make(map[string]int)

###if

    if a := b + c; a > 10 {
	    return a
    }


#import

- import format "fmt" 

	Creates an alias of fmt.
	Preceed all fmt package content with format. instead of fmt.. 

- import . "fmt"

	Allows content of the package to be accessed directly,
	without the need for it to be preceeded with fmt.
	For readiability, don't do this!

- import _ "fmt"

	Suppresses compiler warnings related to fmt if it is not being used,
	and executes initialization functions if there are any.
	The remainder of fmt is inaccessible.
