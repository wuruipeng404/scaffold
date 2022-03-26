package dt

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type SliceStr []string

func (s *SliceStr) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal json value: %v", value)
	}

	var result []string
	err := json.Unmarshal(bytes, &result)

	*s = result
	return err
}

func (s SliceStr) Value() (driver.Value, error) {
	if len(s) == 0 {
		return nil, nil
	}

	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	return string(b), nil
}

func (SliceStr) GormDataType() string {
	return "text"
}

type SliceInt []int

func (s *SliceInt) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal json value: %v", value)
	}

	var result []int
	err := json.Unmarshal(bytes, &result)

	*s = result
	return err
}

func (s SliceInt) Value() (driver.Value, error) {
	if len(s) == 0 {
		return nil, nil
	}

	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	return string(b), nil
}

func (SliceInt) GormDataType() string {
	return "text"
}

type SliceAny []any

func (s *SliceAny) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal json value: %v", value)
	}

	var result []any
	err := json.Unmarshal(bytes, &result)

	*s = result
	return err
}

func (s SliceAny) Value() (driver.Value, error) {
	if len(s) == 0 {
		return nil, nil
	}

	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	return string(b), nil
}

func (SliceAny) GormDataType() string {
	return "text"
}
