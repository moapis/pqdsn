// Copyright (c) 2020, Mohlmann Solutions SRL. All rights reserved.
// Use of this source code is governed by a License that can be found in the LICENSE file.
// SPDX-License-Identifier: BSD-3-Clause

// Package pqdsn offers a type safe way of build Data Source Names for lib/pq.
package pqdsn

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

const (
	dbname                  = "dbname"
	user                    = "user"
	password                = "password"
	host                    = "host"
	port                    = "port"
	sslmode                 = "sslmode"
	fallbackApplicationName = "fallback_application_name"
	connectTimeout          = "connect_timeout"
	sslcert                 = "sslcert"
	sslkey                  = "sslkey"
	sslrootcert             = "sslrootcert"
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
	ConnectTimeout          int     `json:"connect_timeout,omitempty" key:"connect_timeout"`                     // Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
	SSLcert                 string  `json:"sslcert,omitempty"`                                                   // Cert file location. The file must contain PEM encoded data.
	SSLkey                  string  `json:"sslkey,omitempty"`                                                    // Key file location. The file must contain PEM encoded data.
	SSLrootcert             string  `json:"sslrootcert,omitempty"`                                               // The location of the root certificate file. The file must contain PEM encoded data.
}

func (p *Parameters) buildString(escape bool) string {
	b := &builder{escape: escape}

	if p.DBname != "" {
		b.addString(dbname, p.DBname)
	}
	if p.User != "" {
		b.addString(user, p.User)
	}
	if p.Password != "" {
		b.addString(password, p.Password)
	}
	if p.Host != "" {
		b.addString(host, p.Host)
	}
	if p.Port != 0 {
		b.addInt(port, int(p.Port))
	}
	if p.SSLmode != "" {
		b.addString(sslmode, string(p.SSLmode))
	}
	if p.FallbackApplicationName != "" {
		b.addString(fallbackApplicationName, p.FallbackApplicationName)
	}
	if p.ConnectTimeout != 0 {
		b.addInt(connectTimeout, p.ConnectTimeout)
	}
	if p.SSLcert != "" {
		b.addString(sslcert, p.SSLcert)
	}
	if p.SSLkey != "" {
		b.addString(sslkey, p.SSLkey)
	}
	if p.SSLrootcert != "" {
		b.addString(sslrootcert, p.SSLrootcert)
	}

	return b.String()
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

// Merge b into a and return a copy.
// Non-emtpy values from b overwrite a.
func Merge(a, b Parameters) Parameters {
	if b.DBname != "" {
		a.DBname = b.DBname
	}
	if b.User != "" {
		a.User = b.User
	}
	if b.Password != "" {
		a.Password = b.Password
	}
	if b.Host != "" {
		a.Host = b.Host
	}
	if b.Port != 0 {
		a.Port = b.Port
	}
	if b.SSLmode != "" {
		a.SSLmode = b.SSLmode
	}
	if b.FallbackApplicationName != "" {
		a.FallbackApplicationName = b.FallbackApplicationName
	}
	if b.ConnectTimeout != 0 {
		a.ConnectTimeout = b.ConnectTimeout
	}
	if b.SSLcert != "" {
		a.SSLcert = b.SSLcert
	}
	if b.SSLkey != "" {
		a.SSLkey = b.SSLkey
	}
	if b.SSLrootcert != "" {
		a.SSLrootcert = b.SSLrootcert
	}

	return a
}
