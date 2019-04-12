# F1

Simple data structure to computes F1 scores.

https://en.wikipedia.org/wiki/F1_score


[![Build Status](https://travis-ci.org/dmarkham/f1.svg?branch=master)](https://travis-ci.org/dmarkham/f1)
[![GoDoc](https://godoc.org/github.com/dmarkham/f1?status.svg)](https://godoc.org/github.com/dmarkham/f1)

## Install

`go get github.com/dmarkham/f1`

## Example

```go
package main

import (
	"fmt"

	"github.com/dmarkham/f1"
)

func main() {
	f := f1.New()

	f.TruePositive++        // Correct Positive
	fmt.Println(f.Score(1)) // perfect score of 1

	f.TruePositive++        // Correct Negative
	fmt.Println(f.Score(1)) // Still a perfect score

	f.Type1Error++          // False Positive
	fmt.Println(f.Score(1)) // Now the score goes down

	f.Type2Error++          // False Negative
	fmt.Println(f.Score(1)) // Now the score goes down even more
}
```