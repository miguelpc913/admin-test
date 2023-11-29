package dbhelpers

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONIntSLice []int

func (j *JSONIntSLice) Scan(value interface{}) error {
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

func (j JSONIntSLice) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}
