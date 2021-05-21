module main

go 1.15

require (
	collector v0.0.0
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/prometheus/client_golang v1.9.0
	github.com/prometheus/common v0.15.0
	resource v0.0.0
	util v0.0.0
)

replace collector => ./collector

replace util => ./util

replace resource => ./resource
