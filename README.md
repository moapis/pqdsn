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
