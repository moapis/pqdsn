[![Build Status](https://travis-ci.org/moapis/pqdsn.svg?branch=main)](https://travis-ci.org/moapis/pqdsn)
[![codecov](https://codecov.io/gh/moapis/pqdsn/branch/main/graph/badge.svg?token=F5LXD10VK9)](https://codecov.io/gh/moapis/pqdsn)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/moapis/pqdsn)](https://pkg.go.dev/github.com/moapis/pqdsn)
[![Go Report Card](https://goreportcard.com/badge/github.com/moapis/pqdsn)](https://goreportcard.com/report/github.com/moapis/pqdsn)

# PQDSN

Package pqdsn offers a type safe way of build Data Source Names for [lib/pq](https://github.com/lib/pq).
`Parameters` struct type holds all accepted *lib/pq* parameters.
The `String()` method returns a Data Source Name in the format of:

````
"user=pqgotest dbname=pqgotest sslmode=verify-full"
````

## Example

````
p := Parameters{
    DBname:   "pqgotest",
    User:     "pqgotest",
    Password: "secret",
    Host:     "db.example.com",
    Port:     1234,
    SSLmode:  SSLVerifyFull,
    // Use single quotes in values with space!
    FallbackApplicationName: "'pqdsn test'",
}

dsn := p.String()
fmt.Println(dsn)
````

Output:

````
dbname=pqgotest user=pqgotest password=secret host=db.example.com port=1234 sslmode=verify-full fallback_application_name='pqdsn test'
````

## Benchmarks

````
goos: linux
goarch: amd64
pkg: github.com/moapis/pqdsn
BenchmarkBuilder_addString-8                     2407327               461 ns/op             273 B/op          1 allocs/op
BenchmarkParameters_String-8                     2853232               423 ns/op             356 B/op          4 allocs/op
BenchmarkParameters_EscapedString_esc-8          1000000              1099 ns/op             644 B/op          5 allocs/op
PASS
ok      github.com/moapis/pqdsn 4.398s
````

## License

Copyright (c) 2020, Mohlmann Solutions SRL. All rights reserved.
Use of this source code is governed by a License that can be found in the LICENSE file.
