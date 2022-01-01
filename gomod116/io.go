package gomod116

import "io"

//go:embed io.go
var s string

var ReadAll = io.ReadAll
