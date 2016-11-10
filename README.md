# bimap
A bidirectional map written in Go

## Installation
```
go get github.com/vishalkuo/bimap
```

## Usage
```go
import "github.com/vishalkuo/bimap"

biMap := biMap.NewBiMap()
biMap.Insert("key", "value")
val, ok := biMap.InverseGet("value") // val should be "key", ok should be true
biMap.Delete("key")
biMap.Size() // == 0
```
