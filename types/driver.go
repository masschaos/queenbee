package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"
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
		return "", err
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
		return "", err
	}
	return string(tmp), nil
}

type stringSlice []string

func (ss *stringSlice) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	tmp, ok := src.([]byte)
	if !ok {
		return errors.New("read slice string from db failed")
	}
	if len(tmp) == 0 {
		return nil
	}

	*ss = strings.Split(string(tmp), ";")
	return nil
}

func (ds stringSlice) Value() (driver.Value, error) {
	return strings.Join(ds, ";"), nil
}

type mapStrStr map[string]string

func (mss *mapStrStr) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	tmp, ok := src.([]byte)
	if !ok {
		return errors.New("read map value from db failed")
	}
	if len(tmp) == 0 {
		return nil
	}

	return json.Unmarshal(tmp, mss)
}

func (mss mapStrStr) Value() (driver.Value, error) {
	tmp, err := json.Marshal(mss)
	if err != nil {
		return "", err
	}
	return string(tmp), nil
}

type mapStrSliceStr map[string][]string

func (msss *mapStrSliceStr) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	tmp, ok := src.([]byte)
	if !ok {
		return errors.New("read map value from db failed")
	}
	if len(tmp) == 0 {
		return nil
	}

	return json.Unmarshal(tmp, msss)
}

func (msss mapStrSliceStr) Value() (driver.Value, error) {
	tmp, err := json.Marshal(msss)
	if err != nil {
		return "", err
	}
	return string(tmp), nil
}
