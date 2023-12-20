package sqlf

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// NamedDB is a type that represents database with additional named parameter support.
// It is a wrapper around *sqlx.DB, which is an extension of the standard sql.DB type.
type NamedDB struct {
	*sqlx.DB
}

// NamedGet performs a named query and retrieves a single row data into the specified destination.
// It takes the destination pointer, the query string, and the argument as parameters.
// The destination must be a pointer to a struct, and the argument must be a struct or a map.
// If any error occurs during the preparation or retrieval, it is returned.
func (db NamedDB) NamedGet(dest any, query string, arg any) error {
	preparedQuery, err := db.PrepareNamed(query)
	if err != nil {
		return err
	}
	err = preparedQuery.Get(dest, arg)
	return err
}

// NamedSelect executes a named query and retrieves multiple rows of data into the specified destination.
// It takes the destination pointer, the query string, and the argument as parameters.
// The destination must be a pointer to a slice of structs, and the argument must be a struct or a map.
// If any error occurs during the preparation or retrieval, it is returned.
func (db NamedDB) NamedSelect(dest any, query string, arg any) error {
	preparedQuery, err := db.PrepareNamed(query)
	if err != nil {
		return err
	}
	err = preparedQuery.Select(dest, arg)
	return err
}

// NamedGetContext performs a named query and retrieves a single row data into the specified destination using a context.
// It takes the context, the destination pointer, the query string, and the argument as parameters.
// The destination must be a pointer to a struct, and the argument must be a struct or a map.
// If any error occurs during the preparation or retrieval, it is returned.
func (db NamedDB) NamedGetContext(ctx context.Context, dest any, query string, arg any) error {
	preparedQuery, err := db.PrepareNamed(query)
	if err != nil {
		return err
	}
	err = preparedQuery.GetContext(ctx, dest, arg)
	return err
}

// NamedSelectContext performs a named query and retrieves multiple rows of data into the specified destination using a context.
// It takes the context, the destination pointer, the query string, and the argument as parameters.
// The destination must be a pointer to a slice of structs, and the argument must be a struct or a map.
// If any error occurs during the preparation or retrieval, it is returned.
func (db NamedDB) NamedSelectContext(ctx context.Context, dest any, query string, arg any) error {
	preparedQuery, err := db.PrepareNamed(query)
	if err != nil {
		return err
	}
	err = preparedQuery.SelectContext(ctx, dest, arg)
	return err
}
