package utils

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Wrapper for adding JSON array to MySQL or SQLite database
// Make sure to add `gorm:"serializer:json"` as tag on the field using this
type SerializableArray[T any] []T

func (j *SerializableArray[T]) Scan(value interface{}) error {
	switch val := value.(type) {
	case []uint8: // SQLite
		err := json.Unmarshal(val, j)
		if err != nil {
			return err
		}
	case string: // MySQL
		err := json.Unmarshal([]byte(val), j)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("not correct type for scanning value from SQL: %v", value)
	}

	return nil
}

func (j *SerializableArray[T]) Value() (driver.Value, error) {
	if j == nil || len(*j) == 0 {
		return nil, nil
	}

	data, err := json.Marshal(*j)
	if err != nil {
		return nil, err
	}

	return data, nil
}
