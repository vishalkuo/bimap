# bimap
[![GoDoc](https://godoc.org/github.com/vishalkuo/bimap?status.svg)](https://godoc.org/github.com/vishalkuo/bimap)

A generic bidirectional map written in Go

## Installation
```
go get github.com/vishalkuo/bimap
```

## Usage
```go
import "github.com/vishalkuo/bimap"

biMap := bimap.NewBiMap[string, string]()
biMap.Insert("key", "value")
val, ok := biMap.GetInverse("value") // val should be "key", ok should be true
biMap.Delete("key")
biMap.Size() // == 0

biMap2 := bimap.NewBiMap[string, int]()
biMap2.Insert("key", 1) // Notice different types
val, ok = biMap2.Get("key") // Returns 1
val, ok = biMap2.GetInverse(1) // Returns "key"
```
