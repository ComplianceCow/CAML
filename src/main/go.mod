module main

replace commonlibs => ../commonlibs

go 1.17

require commonlibs v0.0.0-00010101000000-000000000000

require (
	github.com/go-yaml/yaml v2.1.0+incompatible // indirect
	github.com/golang/snappy v0.0.3 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/yugabyte/gocql v0.0.0-20211022122817-e03782bd564c // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)
