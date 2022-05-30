package commModel

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

type Time time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

// UnmarshalJSON ...
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

// MarshalJSON ...
func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

// Implement  beego orm Fielder

// Value return the datetime value
func (e Time) Value() time.Time {
	return time.Time(e)
}

// Set set the time.Time to datetime
func (e *Time) Set(d time.Time) {
	*e = Time(d)
}

// String return the time's String
func (e *Time) String() string {
	return e.Value().String()
}

// FieldType return the enum TypeDateTimeField
func (e *Time) FieldType() int {
	return 64
}

// SetRaw convert the string or time.Time to DateTimeField
func (e *Time) SetRaw(value interface{}) error {
	switch d := value.(type) {
	case time.Time:
		e.Set(d)
	case string:
		v, err := time.ParseInLocation(timeFormart, d, time.Local)
		if err == nil {
			e.Set(v)
		}
		return err
	default:
		logs.Error(fmt.Errorf("<DateTimeField.SetRaw> unknown value `%s`", value))
		return nil
	}
	return nil
}

// RawValue return the datetime value
func (e *Time) RawValue() interface{} {
	return e.Value()
}
