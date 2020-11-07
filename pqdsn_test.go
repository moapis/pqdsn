// Copyright (c) 2020, Mohlmann Solutions SRL. All rights reserved.
// Use of this source code is governed by a License that can be found in the LICENSE file.
// SPDX-License-Identifier: BSD-3-Clause

// Package pqdsn offers a type safe way of build Data Source Names for lib/pq.
package pqdsn

import (
	"reflect"
	"strings"
	"testing"
)

func Test_addToBuilder(t *testing.T) {
	var b strings.Builder

	addToBuilder(&b, "foo", "bar")

	want := "foo=bar"
	got := b.String()

	if got != want {
		t.Errorf("addToBuilder() = %s, want %s", got, want)
	}

	addToBuilder(&b, "spanac", 123)

	want = "foo=bar spanac=123"
	got = b.String()

	if got != want {
		t.Errorf("addToBuilder() = %s, want %s", got, want)
	}
}

func Test_key(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"DBname",
			"dbname",
		},
		{
			"FallbackApplicationName",
			"fallback_application_name",
		},
		{
			"ConnectTimeout",
			"connect_timeout",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pt := reflect.TypeOf(Parameters{})
			field, ok := pt.FieldByName(tt.name)
			if !ok {
				t.Fatalf("Test Field Name %q does not exist", tt.name)
			}

			if got := key(field); got != tt.want {
				t.Errorf("key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameters_String(t *testing.T) {
	type fields struct {
		DBname                  string
		User                    string
		Password                string
		Host                    string
		Port                    uint16
		SSLmode                 SSLMode
		FallbackApplicationName string
		ConnectTimeout          uint
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
				Password:                "secret",
				Host:                    "db.example.com",
				Port:                    1234,
				SSLmode:                 SSLVerifyFull,
				FallbackApplicationName: "'pqdsn test'",
			},
			"dbname=pqgotest user=pqgotest password=secret host=db.example.com port=1234 sslmode=verify-full fallback_application_name='pqdsn test'",
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
				t.Errorf("Parameters.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
