package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Scan implements the Scanner interface.
func (t *RequestTpl) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	tmp, ok := src.([]byte)
	if !ok {
		return errors.New("read json from DB failed")
	}
	if len(tmp) == 0 {
		return nil
	}
	return json.Unmarshal(tmp, t)
}

// Value implements the driver Valuer interface.
func (t RequestTpl) Value() (driver.Value, error) {
	tmp, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return string(tmp), nil
}

// Scan implements the Scanner interface.
func (t *ResponseTpl) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	tmp, ok := src.([]byte)
	if !ok {
		return errors.New("read json from DB failed")
	}
	if len(tmp) == 0 {
		return nil
	}
	return json.Unmarshal(tmp, t)
}

// Value implements the driver Valuer interface.
func (t ResponseTpl) Value() (driver.Value, error) {
	tmp, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return string(tmp), nil
}
