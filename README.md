golang-tips and examples
========================


##variable

- Short variable declarations could only be used inside a function.
- byte is alias for uint8
- rune is alias for int32 // an integer value identifying a Unicode code point
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

##tips

- use signed types for everything except bit manipulation
- [Why is rune in golang an alias for int32 and not uint32?]

        Go uses a lot of signed values, not just for runes but array indices,
        Read/Write byte counts, etc. That's because uints, in any language,
        behave confusingly unless you guard every piece of arithmetic against overflow
        (for example if var a, b uint = 1, 2, a-b > 0 and a-b > 1000000: play.golang.org/p/lsdiZJiN7V).
        ints behave more like numbers in everyday life,
        which is a compelling reason to use them,
        and there is no equally compelling reason not to use them.

[Why is rune in golang an alias for int32 and not uint32?]: http://stackoverflow.com/questions/24714665/why-is-rune-in-golang-an-alias-for-int32-and-not-uint32
