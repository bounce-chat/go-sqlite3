// Copyright (C) 2026 BitBoxSwiss
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//go:build !libsqlcipher && (libsqlite3 || sqlite_plain)
// +build !libsqlcipher
// +build libsqlite3 sqlite_plain

package sqlite3

import (
	"database/sql"
	"strings"
	"testing"
)

func TestCipherOptionsRequireSQLCipher(t *testing.T) {
	tests := []string{
		"file::memory:?_key=passphrase",
		"file::memory:?_cipher_compatibility=4",
		"file::memory:?_cipher_migrate",
		"file::memory:?_cipher_page_size=4096",
		"file::memory:?_cipher_plaintext_header_size=32",
		"file::memory:?_cipher_use_hmac=on",
	}

	for _, dsn := range tests {
		t.Run(dsn, func(t *testing.T) {
			db, err := sql.Open("sqlite3", dsn)
			if err != nil {
				t.Fatalf("sql.Open: %v", err)
			}
			defer db.Close()

			err = db.Ping()
			if err == nil {
				t.Fatal("expected SQLCipher option to fail without SQLCipher support")
			}
			if !strings.Contains(err.Error(), "SQLCipher") {
				t.Fatalf("expected SQLCipher error, got: %v", err)
			}
		})
	}
}

// An empty _key is refused in every build, so that a DSN which asks for a key
// never opens a database without one.
func TestCipherEmptyKeyRefused(t *testing.T) {
	db, err := sql.Open("sqlite3", "file::memory:?_key=")
	if err != nil {
		t.Fatalf("sql.Open: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err == nil {
		t.Fatal("expected an empty key to be refused")
	}
	if !strings.Contains(err.Error(), "_key") {
		t.Fatalf("expected an error naming _key, got: %v", err)
	}
}
