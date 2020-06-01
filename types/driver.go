package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Scan implements the Scanner interface.
func (t *RequestTemp) Scan(src interface{}) error {
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
func (t RequestTemp) Value() (driver.Value, error) {
	tmp, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(tmp), nil
}

// Scan implements the Scanner interface.
func (t *ResponseTemp) Scan(src interface{}) error {
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
func (t ResponseTemp) Value() (driver.Value, error) {
	tmp, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(tmp), nil
}
