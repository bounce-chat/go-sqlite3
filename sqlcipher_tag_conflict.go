// Copyright (C) 2026 BitBoxSwiss
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//go:build libsqlite3 && libsqlcipher
// +build libsqlite3,libsqlcipher

package sqlite3

// With both tags the key pragma could reach the system's plain SQLite, which
// ignores it and leaves the database unencrypted.

/*
#error "the libsqlite3 and libsqlcipher build tags are mutually exclusive"
*/
import "C"
