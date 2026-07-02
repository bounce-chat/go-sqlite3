// Copyright (C) 2026 BitBoxSwiss.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

//go:build !libsqlite3 && !libsqlcipher && !sqlite_plain && !darwin
// +build !libsqlite3,!libsqlcipher,!sqlite_plain,!darwin

package sqlite3

/*
#cgo CFLAGS: -DSQLCIPHER_CRYPTO_LIBTOMCRYPT
#cgo CFLAGS: -I${SRCDIR}/internal/libtomcrypt
*/
import "C"

import _ "github.com/mattn/go-sqlite3/internal/libtomcrypt"
