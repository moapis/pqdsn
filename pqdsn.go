// Copyright (c) 2020, Mohlmann Solutions SRL. All rights reserved.
// Use of this source code is governed by a License that can be found in the LICENSE file.
// SPDX-License-Identifier: BSD-3-Clause

// Package pqdsn offers a type safe way of build Data Source Names for lib/pq.
package pqdsn

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

// SSLMode used for connection
type SSLMode string

const (
	// SSLDisable means: No SSL
	SSLDisable SSLMode = "disable"
	// SSLRequire means: Always SSL
	// (skip verification)
	SSLRequire SSLMode = "require"
	// SSLVerifyCA means: Always SSL
	// (verify that the certificate presented by the server was signed by a trusted CA)
	SSLVerifyCA SSLMode = "verify-ca"
	// SSLVerifyFull means: Always SSL
	// (verify that the certification presented by the server was signed by a trusted CA
	// and the server host name matches the one in the certificate)
	SSLVerifyFull SSLMode = "verify-full"
)

// Parameters as defined at https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters
type Parameters struct {
	DBname                  string  `json:"dbname,omitempty"`                                                    // The name of the database to connect to
	User                    string  `json:"user,omitempty"`                                                      // The user to sign in as
	Password                string  `json:"password,omitempty"`                                                  // The user's password
	Host                    string  `json:"host,omitempty"`                                                      // The host to connect to. Values that start with / are for unix domain sockets. (default is localhost)
	Port                    uint16  `json:"port,omitempty"`                                                      // The port to bind to. (default is 5432)
	SSLmode                 SSLMode `json:"sslmode,omitempty"`                                                   // Whether or not to use SSL (default is require, this is not the default for libpq)
	FallbackApplicationName string  `json:"fallback_application_name,omitempty" key:"fallback_application_name"` // An application_name to fall back to if one isn't provided.
	ConnectTimeout          uint    `json:"connect_timeout,omitempty" key:"connect_timeout"`                     // Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
	SSLcert                 string  `json:"sslcert,omitempty"`                                                   // Cert file location. The file must contain PEM encoded data.
	SSLkey                  string  `json:"sslkey,omitempty"`                                                    // Key file location. The file must contain PEM encoded data.
	SSLrootcert             string  `json:"sslrootcert,omitempty"`                                               // The location of the root certificate file. The file must contain PEM encoded data.

	buff *bytes.Buffer // Re-usable buffer
}

func addToBuffer(b *bytes.Buffer, k string, v interface{}, esc bool) {
	if b.Len() > 0 {
		b.WriteRune(' ')
	}

	if esc {
		if s, ok := v.(string); ok {
			s = strings.ReplaceAll(s, "'", `\'`) // Escape single quotes

			if strings.ContainsRune(s, ' ') {
				fmt.Fprintf(b, "%s='%s'", k, s)
			} else {
				fmt.Fprintf(b, "%s=%s", k, s)
			}

			return
		}
	}

	// strings.Builder never errors, so ignoring error
	fmt.Fprintf(b, "%s=%v", k, v)
}

func key(field reflect.StructField) string {
	if k := field.Tag.Get("key"); k != "" {
		return k
	}

	return strings.ToLower(field.Name)
}

func (p Parameters) buildString(esc bool) string {
	if p.buff == nil {
		p.buff = new(bytes.Buffer)
	}
	defer p.buff.Reset()

	v := reflect.ValueOf(p)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if v := v.Field(i); v.CanInterface() && !v.IsZero() {
			addToBuffer(p.buff, key(t.Field(i)), v.Interface(), esc)
		}
	}

	return p.buff.String()
}

// String returns a Data Source Name in the format of:
//
//     "dbname=pqgotest user=pqgotest sslmode=verify-full"
func (p *Parameters) String() string {
	return p.buildString(false)
}

// EscapedString single quote (') entries containing spaces.
// Existing single quotes in values are escaped by a back slash (\).
//
// Output is in the format of:
//
//    "dbname=pqgotest user='space man' password='it\'s valid'"
func (p *Parameters) EscapedString() string {
	return p.buildString(true)
}
