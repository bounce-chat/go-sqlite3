#!/bin/sh
#
# Checks what the SQLCipher support relies on and what a rebase onto a newer
# mattn/go-sqlite3 can undo without a merge conflict:
#
#   1. The generated amalgamations start with the guards that keep plain SQLite
#      and SQLCipher out of each other's builds. An upstream SQLite upgrade
#      regenerates sqlite3-binding.c, sqlite3-binding.h and sqlite3ext.h, and
#      taking those files as they come drops the guards.
#   2. Every file that includes sqlite3-binding.h does so through the switch
#      that selects sqlcipher-binding.h under USE_SQLCIPHER. A file added
#      upstream arrives without the switch.
#   3. Every build constraint that mentions libsqlite3 also accounts for
#      libsqlcipher, the other tag that links a system library.
#
# Run it after every rebase. The SQLCipher workflow runs it as well.

set -eu

cd "$(dirname "$0")/../.."

failed=0

fail() {
  echo "check_guards: $*" >&2
  failed=1
}

expect_first_line() {
  file="$1"
  expected="$2"
  actual=$(head -n 1 "$file")
  if [ "$actual" != "$expected" ]; then
    fail "$file must start with: $expected"
  fi
}

plain_guard='#if !defined(USE_LIBSQLITE3) && !defined(USE_LIBSQLCIPHER) && !defined(USE_SQLCIPHER)'
expect_first_line sqlite3-binding.c "$plain_guard"
expect_first_line sqlite3-binding.h "$plain_guard"
expect_first_line sqlite3ext.h '#if !defined(USE_LIBSQLITE3) && !defined(USE_LIBSQLCIPHER)'
expect_first_line sqlcipher-binding.c '#ifdef USE_SQLCIPHER'
expect_first_line sqlcipher-binding.h '#ifdef USE_SQLCIPHER'

for file in *.go *.c *.h; do
  case "$file" in
    sqlite3-binding.[ch] | sqlcipher-binding.[ch]) continue ;;
  esac

  if grep -q '#include "sqlite3-binding.h"' "$file" &&
    ! grep -q '#include "sqlcipher-binding.h"' "$file"; then
    fail "$file includes sqlite3-binding.h without the USE_SQLCIPHER switch"
  fi

  # sqlite3_libsqlite3.go is the file the libsqlite3 tag itself selects.
  if [ "$file" != sqlite3_libsqlite3.go ] &&
    grep -q '^//go:build.*libsqlite3' "$file" &&
    ! grep -q '^//go:build.*libsqlcipher' "$file"; then
    fail "$file has a libsqlite3 build constraint that does not mention libsqlcipher"
  fi
done

if [ "$failed" -ne 0 ]; then
  exit 1
fi

echo "check_guards: ok"
