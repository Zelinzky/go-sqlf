package sqlf

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type NamedDB struct {
	*sqlx.DB
}

func (db NamedDB) NamedGet(dest any, query string, arg any) error {
	preparedQuery, err := db.PrepareNamed(query)
	if err != nil {
		return err
	}
	err = preparedQuery.Get(dest, arg)
	return err
}

func (db NamedDB) NamedSelect(dest any, query string, arg any) error {
	preparedQuery, err := db.PrepareNamed(query)
	if err != nil {
		return err
	}
	err = preparedQuery.Select(dest, arg)
	return err
}

func (db NamedDB) NamedGetContext(ctx context.Context, dest any, query string, arg any) error {
	preparedQuery, err := db.PrepareNamed(query)
	if err != nil {
		return err
	}
	err = preparedQuery.GetContext(ctx, dest, arg)
	return err
}

func (db NamedDB) NamedSelectContext(ctx context.Context, dest any, query string, arg any) error {
	preparedQuery, err := db.PrepareNamed(query)
	if err != nil {
		return err
	}
	err = preparedQuery.SelectContext(ctx, dest, arg)
	return err
}
