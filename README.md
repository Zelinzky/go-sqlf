# go-sqlf
go-sqlf is a collection of helpers that work hand in hand with sqlx to make your life easier.

sqlf stands for SQL File, this library takes inspiration on https://github.com/sqlc-dev/sqlc way of parsing sql files to generate go files, but only cares about the parsing.
SQLf only has one method called load, it loads several queries from a string which should have been previously embedded into your go binary.



