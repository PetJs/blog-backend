package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DateOnly struct {
	time.Time
}

func (DateOnly) GormDataType() string {
	return "date"
}

func (DateOnly) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return "date"
}

func (d DateOnly) Value() (driver.Value, error) {
	if !d.IsZero() {
		return d.Format("2006-01-02"), nil
	}
	return nil, nil
}

func (d *DateOnly) Scan(value interface{}) error {
	scanned, ok := value.(time.Time)
	if !ok {
		return errors.New(fmt.Sprint("Failed to scan DateOnly:", value))
	}
	*d = DateOnly{scanned}
	return nil
}

func (d DateOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format("2006-01-02"))
}

func (d *DateOnly) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*d = DateOnly{t}
	return nil
}

// ------------------- TIME ONLY --------------------

type TimeOnly struct {
	time.Time
}

func (TimeOnly) GormDataType() string {
	return "time"
}

func (TimeOnly) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return "time"
}

func (t TimeOnly) Value() (driver.Value, error) {
	if !t.IsZero() {
		return t.Format("15:04:05"), nil
	}
	return nil, nil
}

func (t *TimeOnly) Scan(value interface{}) error {
	scanned, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to scan TimeOnly:", value))
	}
	parsed, err := time.Parse("15:04:05", string(scanned))
	if err != nil {
		return err
	}
	*t = TimeOnly{parsed}
	return nil
}

func (t TimeOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Format("15:04:05"))
}

func (t *TimeOnly) UnmarshalJSON(bs []byte) error {
	var s string
	if err := json.Unmarshal(bs, &s); err != nil {
		return err
	}
	parsed, err := time.Parse("15:04:05", s)
	if err != nil {
		return err
	}
	*t = TimeOnly{parsed}
	return nil
}
