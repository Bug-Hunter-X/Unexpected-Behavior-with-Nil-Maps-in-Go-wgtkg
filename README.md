# Unexpected Behavior with Nil Maps in Go

This repository demonstrates a potential pitfall when working with maps in Go: accessing keys in a nil map.

In many languages, attempting to access a key in a nil map would result in a runtime error (e.g., a `NullPointerException`).  However, in Go, accessing a key in a nil map simply returns the zero value for the map's value type. This can lead to subtle bugs if not handled carefully.

The `bug.go` file contains the code that exhibits this behavior. The `bugSolution.go` file provides a solution to prevent this kind of error by explicitly checking for `nil`. 