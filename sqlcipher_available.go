// Copyright (C) 2026 BitBoxSwiss
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//go:build libsqlcipher || (!libsqlite3 && !sqlite_plain)
// +build libsqlcipher !libsqlite3,!sqlite_plain

package sqlite3

const sqlcipherAvailable = true
