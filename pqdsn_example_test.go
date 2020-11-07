package pqdsn

import "fmt"

func Example() {
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

	// Output: dbname=pqgotest user=pqgotest password=secret host=db.example.com port=1234 sslmode=verify-full fallback_application_name='pqdsn test'
}
