package dbhelpers

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONStringSlice []string

func (j *JSONStringSlice) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, j)
}

func (j JSONStringSlice) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}
