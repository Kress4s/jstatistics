package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgtype"
)

type BigintArray []int64

// BigintArray数据库特殊数据的驱动

func (a BigintArray) Value() (driver.Value, error) {
	if a == nil {
		a = make(BigintArray, 0)
	}
	arr := pgtype.Int8Array{}
	if err := arr.Set([]int64(a)); err != nil {
		return nil, err
	}
	return arr.Value()
}

func (a *BigintArray) Scan(src interface{}) error {
	arr := pgtype.Int8Array{}
	if err := arr.Scan(src); err != nil {
		return err
	}
	return arr.AssignTo(a)
}

func (a BigintArray) MarshalJSON() ([]byte, error) {
	values := make([]string, len(a))
	for i, value := range []int64(a) {
		values[i] = fmt.Sprintf(`"%v"`, value)
	}
	return []byte(fmt.Sprintf("[%v]", strings.Join(values, ","))), nil
}

func (a *BigintArray) UnmarshalJSON(b []byte) error {
	// Try array of strings first.
	var values []string
	err := json.Unmarshal(b, &values)
	if err != nil {
		// Fall back to array of integers:
		var values []int64
		if err := json.Unmarshal(b, &values); err != nil {
			return err
		}
		*a = values
		return nil
	}
	*a = make([]int64, len(values))
	for i, value := range values {
		value, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		(*a)[i] = value
	}
	return nil
}
