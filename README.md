# go-cmp

Package **cmp** provider generic functions for **less-than**, **equal**, and **greater-than** operations, for the Go programming language.

This is to provide a more _friendly_ (and less error-prone) developer-experience to the `Cmp()` method.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-cmp

[![GoDoc](https://godoc.org/github.com/reiver/go-cmp?status.svg)](https://godoc.org/github.com/reiver/go-cmp)

## Examples

```golang
import "github.com/reiver/go-cmp"

// ...

var x *big.Int = big.NewInt(-5)
var y *big.Int = big.NewInt(10)

// ...

if cmp.LessThan(x,y) {
	// ...
}

// ...

if cmp.Equal(x,y) {
	// ...
}

// ...

if cmp.GreaterThan(x,y) {
	// ...
}
```

## Import

To import package **cmp** use `import` code like the following:
```
import "github.com/reiver/go-cmp"
```

## Installation

To install package **cmp** do the following:
```
GOPROXY=direct go get github.com/reiver/go-cmp
```

## Author

Package **cmp** was written by [Charles Iliya Krempeaux](http://reiver.link)
