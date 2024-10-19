package scim2

import (
	"errors"
	"github.com/spf13/cast"
	"regexp"
	"strings"
	"time"
)

func strArrToIface(values ...string) (ret []interface{}) {
	for _, s := range values {
		ret = append(ret, s)
	}

	return
}
func NewScimError(status int, detail, scimType string) ScimError {
	return ScimError{
		Status:   status,
		Detail:   detail,
		ScimType: scimType,
		Schemas:  []string{URN_ERROR},
	}
}

func IsInteger(val any) bool {
	_, err := cast.ToInt64E(val)
	return err == nil
}

func IsFloat(val any) bool {
	str := cast.ToString(val)
	if !strings.Contains(str, ".") {
		return false
	}
	_, err := cast.ToFloat64E(val)
	return err == nil
}

func IsDate(val any) bool {
	str := cast.ToString(val)
	dateRegex := `^\d{4}-\d{2}-\d{2}\s?[TZ0-9:.+\-]+$`
	matched, _ := regexp.MatchString(dateRegex, str)
	return matched
}

func ParseDate(dateString string) (*time.Time, error) {
	dateFormats := []string{
		"2006-01-02T15:04:05.000-07:00Z",
		"2006-01-02T15:04:05.000Z",
		"2006-01-02T15:04:05Z",
		"2006-01-02 15:04:05.000-07:00",
		"2006-01-02 15:04:05.000",
		"2006-01-02 15:04:05",
	}

	for _, format := range dateFormats {
		t, err := time.Parse(format, dateString)
		if err == nil {
			return &t, nil
		}
	}

	return nil, errors.New("invalid date format")
}
