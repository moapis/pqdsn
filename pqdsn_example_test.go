package pqdsn

import "fmt"

func Example() {
	p := Parameters{
		DBname:                  "pqgotest",
		User:                    "pqgotest",
		Password:                "it's secret",
		Host:                    "db.example.com",
		Port:                    1234,
		SSLmode:                 SSLVerifyFull,
		FallbackApplicationName: "pqdsn test",
	}

	dsn := p.EscapedString()
	fmt.Println(dsn)

	// Output: dbname=pqgotest user=pqgotest password='it\'s secret' host=db.example.com port=1234 sslmode=verify-full fallback_application_name='pqdsn test'
}
