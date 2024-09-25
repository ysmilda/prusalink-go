package printer

import (
	"encoding/json"
	"time"
)

type Duration time.Duration

func (d Duration) String() string {
	return time.Duration(d).String()
}

func (d Duration) MarshallJSON() ([]byte, error) {
	return json.Marshal(time.Duration(d))
}

func (d *Duration) UnmarshallJSON(data []byte) error {
	var s int
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*d = Duration(time.Duration(s) * time.Second)
	return nil
}

type Time time.Time

func (t Time) String() string {
	return time.Time(t).String()
}

func (t Time) MarshallJSON() ([]byte, error) {
	return json.Marshal(time.Time(t))
}

func (t *Time) UnmarshallJSON(data []byte) error {
	var s int64
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	*t = Time(time.Unix(s, 0))
	return nil
}
