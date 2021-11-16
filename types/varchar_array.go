package types

import (
	"database/sql/driver"

	"github.com/jackc/pgtype"
)

type VarcharArray []string

func (va VarcharArray) Value() (driver.Value, error) {
	if va == nil {
		va = make(VarcharArray, 0)
	}

	t := pgtype.VarcharArray{}
	if err := t.Set([]string(va)); err != nil {
		return nil, err
	}

	return t.Value()
}

func (va *VarcharArray) Scan(src interface{}) error {
	t := pgtype.VarcharArray{}
	if err := t.Scan(src); err != nil {
		return err
	}

	return t.AssignTo(va)
}
