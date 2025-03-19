package valueobject

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Uuid struct {
	value uuid.UUID
	isSet bool
}

func (u *Uuid) Value() uuid.UUID {
	return u.value
}

func (u *Uuid) IsSet() bool {
	return u.isSet
}

func (u *Uuid) MarshalJSON() ([]byte, error) {
	var value string
	if u.isSet {
		value = u.value.String()
	}
	return json.Marshal(value)
}

func (u *Uuid) UnmarshalJSON(b []byte) error {
	var strValue string
	u.isSet = false
	err := json.Unmarshal(b, &strValue)
	if err != nil {
		return err
	}
	if strValue == "" {
		return nil
	}
	parsedValue, err := uuid.Parse(strValue)
	if err != nil {
		return nil
	}
	u.value = parsedValue
	u.isSet = true
	return nil
}

func NewUuid() Uuid {
	return Uuid{value: uuid.New(), isSet: true}
}

func NewUuidFromUuid(uuid uuid.UUID) Uuid {
	return Uuid{value: uuid, isSet: true}
}

func NewUuidFromString(s string) (Uuid, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return Uuid{}, err
	}
	return Uuid{value: id, isSet: true}, nil
}

func NullUuid() Uuid {
	return Uuid{isSet: false}
}
