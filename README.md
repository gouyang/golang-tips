golang-tips and examples
========================

#Basics

##variable

- Short variable declarations could only be used inside a function.
- byte is alias for uint8
- var b bytes.Buffer // Buffer needs no initialization:
- var nil Type // Type must be a pointer, channel, func, interface, map, or slice type
- undefined variable

    if `go run` says undefined varible, try to include the dependency file to the commandline,
    like `go run main.go entry.go`.

##pointer

- Methods with pointer receivers

	There are two reasons to use a pointer receiver.

	First, to avoid copying the value on each method call (more efficient if the value type is a large struct).
	Second, so that the method can modify the value that its receiver points to.


##import

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

##non-local types

"It's impossible to define new methods on non-local type[s]", which is by design.
The best practice is to embed the non-local type into your own own
local type, and extend it.

	type MyExtension struct {
	    otherPackage.Type
	} 

	func (me *MyExtension) NewMethod() { ...  }

