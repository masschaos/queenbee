package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// Network is worker ISP location
type Network struct {
	Region string `json:"region" gorm:"type:varchar(50)"` // country or region
	State  string `json:"state" gorm:"type:varchar(50)"`  // state or province
	City   string `json:"city" gorm:"type:varchar(50)"`   // city
	ISP    string `json:"isp" gorm:"type:varchar(50)"`    // ISP
}

// IPv4Network is a db table: ip_v4_networks
type IPv4Network struct {
	IP string `json:"ip" gorm:"type:varchar(15);primary_key"`
	Network
	CreatedAt time.Time `json:"createdAt"` // for gorm
}

// TableName is for gorm table name
func (n IPv4Network) TableName() string {
	return "ip_v4_networks"
}

// Scan implements the Scanner interface.
func (n *IPv4Network) Scan(src interface{}) error {
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
	return json.Unmarshal(tmp, n)
}

// Value implements the driver Valuer interface.
func (n IPv4Network) Value() (driver.Value, error) {
	tmp, err := json.Marshal(n)
	if err != nil {
		return nil, err
	}
	return string(tmp), nil
}
