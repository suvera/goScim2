package scim2

import (
	"encoding/json"
	"strings"
	"time"
)

type JsonDate time.Time

func (j *JsonDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02T15:04:05Z", s)
	if err != nil {
		return err
	}
	*j = JsonDate(t)
	return nil
}

func (j JsonDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

func toJsonDate(secs int64) JsonDate {
	return JsonDate(time.Unix(secs, 0))
}
