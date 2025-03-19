package valueobject

import (
	"encoding/json"
	"time"
)

type OptionalTime struct {
	value time.Time
	isSet bool
}

func (ot *OptionalTime) Value() time.Time {
	return ot.value
}

func (ot *OptionalTime) IsSet() bool {
	return ot.isSet
}

func NewOptionalTime(value time.Time) OptionalTime {
	return OptionalTime{value: value, isSet: true}
}

func NullOptionalTime() OptionalTime {
	return OptionalTime{isSet: false}
}

func (ot *OptionalTime) MarshalJSON() ([]byte, error) {
	var value string
	if ot.isSet {
		value = ot.value.Format(time.RFC3339)
	}
	return json.Marshal(value)
}

func (ot *OptionalTime) UnmarshalJSON(b []byte) error {
	var strValue string
	ot.isSet = false
	err := json.Unmarshal(b, &strValue)
	if err != nil {
		return err
	}
	if strValue == "" {
		return nil
	}
	parsedValue, err := time.Parse(time.RFC3339, strValue)
	if err != nil {
		return nil
	}
	ot.value = parsedValue
	ot.isSet = true
	return nil
}

type OptionalString struct {
	value string
	isSet bool
}

func (os *OptionalString) Value() string {
	return os.value
}

func (os *OptionalString) IsSet() bool {
	return os.isSet
}

func (os *OptionalString) MarshalJSON() ([]byte, error) {
	var value string
	if os.isSet {
		value = os.value
	}
	return json.Marshal(value)
}

func (os *OptionalString) UnmarshalJSON(b []byte) error {
	var strValue string
	os.isSet = false
	err := json.Unmarshal(b, &strValue)
	if err != nil {
		return err
	}
	if strValue == "" {
		return nil
	}
	os.value = strValue
	return nil
}

func NewOptionalString(value string) OptionalString {
	return OptionalString{value: value, isSet: true}
}

func NullOptionalString() OptionalString {
	return OptionalString{isSet: false}
}
