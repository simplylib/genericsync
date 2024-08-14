[![Go Reference](https://pkg.go.dev/badge/github.com/simplylib/genericsync.svg)](https://pkg.go.dev/github.com/simplylib/genericsync)
[![Go Report Card](https://goreportcard.com/badge/github.com/simplylib/genericsync)](https://goreportcard.com/report/github.com/simplylib/genericsync)

# genericsync
A Go (golang) library to add generic wrappers to some standard library sync package types 

Currently the only implemented struct from [sync](https://pkg.go.dev/sync) is [sync.Map](https://pkg.go.dev/sync#Map)
and for [sync.Map](https://pkg.go.dev/sync#Map) we make sure to test our functions with the tests in those packages,
but rewritten to use generics. These are not pulled into the package, but ad-hoc run locally when a new function is added
