// Copyright (c) 2020, Mohlmann Solutions SRL. All rights reserved.
// Use of this source code is governed by a License that can be found in the LICENSE file.
// SPDX-License-Identifier: BSD-3-Clause

package pqdsn

import (
	"testing"
)

func TestParameters_String(t *testing.T) {
	type fields struct {
		DBname                  string
		User                    string
		Password                string
		Host                    string
		Port                    uint16
		SSLmode                 SSLMode
		FallbackApplicationName string
		ConnectTimeout          int
		SSLcert                 string
		SSLkey                  string
		SSLrootcert             string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"No fields",
			fields{},
			"",
		},
		{
			"Some fields",
			fields{
				DBname:                  "pqgotest",
				User:                    "pqgotest",
				Password:                `'it\'s secret'`,
				Host:                    "db.example.com",
				Port:                    1234,
				SSLmode:                 SSLVerifyFull,
				FallbackApplicationName: "'pqdsn test'",
			},
			`dbname=pqgotest user=pqgotest password='it\'s secret' host=db.example.com port=1234 sslmode=verify-full fallback_application_name='pqdsn test'`,
		},
		{
			"All fields",
			fields{
				"nameofdb",
				"itisme",
				"secret",
				"localtoast",
				789,
				SSLVerifyFull,
				"This-App",
				22,
				"/path/to/cert",
				"/path/to/key",
				"/path/to/rootcert",
			},
			"dbname=nameofdb user=itisme password=secret host=localtoast port=789 sslmode=verify-full fallback_application_name=This-App connect_timeout=22 sslcert=/path/to/cert sslkey=/path/to/key sslrootcert=/path/to/rootcert",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parameters{
				DBname:                  tt.fields.DBname,
				User:                    tt.fields.User,
				Password:                tt.fields.Password,
				Host:                    tt.fields.Host,
				Port:                    tt.fields.Port,
				SSLmode:                 tt.fields.SSLmode,
				FallbackApplicationName: tt.fields.FallbackApplicationName,
				ConnectTimeout:          tt.fields.ConnectTimeout,
				SSLcert:                 tt.fields.SSLcert,
				SSLkey:                  tt.fields.SSLkey,
				SSLrootcert:             tt.fields.SSLrootcert,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("Parameters.String() = \n%v\nwant\n%v", got, tt.want)
			}
		})
	}
}

func TestParameters_EscapedString(t *testing.T) {
	type fields struct {
		DBname                  string
		User                    string
		Password                string
		Host                    string
		Port                    uint16
		SSLmode                 SSLMode
		FallbackApplicationName string
		ConnectTimeout          int
		SSLcert                 string
		SSLkey                  string
		SSLrootcert             string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"No fields",
			fields{},
			"",
		},
		{
			"Some fields",
			fields{
				DBname:                  "pqgotest",
				User:                    "pqgotest",
				Password:                "it's secret",
				Host:                    "db.example.com",
				Port:                    1234,
				SSLmode:                 SSLVerifyFull,
				FallbackApplicationName: "pqdsn test",
			},
			`dbname=pqgotest user=pqgotest password='it\'s secret' host=db.example.com port=1234 sslmode=verify-full fallback_application_name='pqdsn test'`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parameters{
				DBname:                  tt.fields.DBname,
				User:                    tt.fields.User,
				Password:                tt.fields.Password,
				Host:                    tt.fields.Host,
				Port:                    tt.fields.Port,
				SSLmode:                 tt.fields.SSLmode,
				FallbackApplicationName: tt.fields.FallbackApplicationName,
				ConnectTimeout:          tt.fields.ConnectTimeout,
				SSLcert:                 tt.fields.SSLcert,
				SSLkey:                  tt.fields.SSLkey,
				SSLrootcert:             tt.fields.SSLrootcert,
			}
			if got := p.EscapedString(); got != tt.want {
				t.Errorf("Parameters.EscapedString() =\n%v\nwant\n%v", got, tt.want)
			}
		})
	}
}

func BenchmarkParameters_String(t *testing.B) {
	p := Parameters{
		DBname:                  "pqgotest",
		User:                    "pqgotest",
		Password:                `'it\'s secret'`,
		Host:                    "db.example.com",
		Port:                    1234,
		SSLmode:                 SSLVerifyFull,
		FallbackApplicationName: "'pqdsn test'",
	}

	for i := 0; i < t.N; i++ {
		_ = p.String()
	}
}

func BenchmarkParameters_EscapedString_esc(t *testing.B) {
	p := Parameters{
		DBname:                  "pqgotest",
		User:                    "pqgotest",
		Password:                "it's secret",
		Host:                    "db.example.com",
		Port:                    1234,
		SSLmode:                 SSLVerifyFull,
		FallbackApplicationName: "pqdsn test",
	}

	for i := 0; i < t.N; i++ {
		_ = p.EscapedString()
	}
}

func TestMerge(t *testing.T) {
	a := Parameters{
		DBname:                  "pqgotest",
		User:                    "pqgotest",
		Password:                "it's secret",
		Host:                    "db.example.com",
		Port:                    1234,
		SSLmode:                 SSLVerifyFull,
		FallbackApplicationName: "pqdsn test",
	}

	got := Merge(a, Parameters{})
	if got != a {
		t.Errorf("Merge() = \n%v\nwant\n%v", got, a)
	}

	b := Parameters{
		DBname:                  "name",
		User:                    "user",
		Password:                "spanac",
		Host:                    "example.com",
		Port:                    21,
		SSLmode:                 SSLRequire,
		FallbackApplicationName: "who hoo",
		ConnectTimeout:          33,
		SSLcert:                 "/path/to/cert",
		SSLkey:                  "/path/to/key",
		SSLrootcert:             "/path/to/rootcert",
	}

	got = Merge(a, b)
	if got != b {
		t.Errorf("Merge() = \n%v\nwant\n%v", got, b)
	}
}
