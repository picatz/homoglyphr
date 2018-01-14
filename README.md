# homoglyphr
> Confusing domain name character generator.

A simple [homoglyph](https://en.wikipedia.org/wiki/Homoglyph) generation library for gophers focused on registrable domain name characters.

## Installation

```
go get github.com/picatz/homoglyphr
```

## Example Usage

```go
package main

import (
  "fmt"
  "github.com/picatz/homoglyphr"
)

func main() {
  for similar := range homoglyphr.StreamAllRelatedCharacters("g") {
    fmt.Println(similar)
  }
}
```
