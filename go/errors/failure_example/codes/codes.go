package codes

import "github.com/morikuni/failure"

// example codes from failure's README.md: https://github.com/morikuni/failure#usage
const (
	NotFound        failure.StringCode = "NotFound"
	InvalidArgument failure.StringCode = "InvalidArgument"
	Internal        failure.StringCode = "Internal"
)
