package models

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type JSON []byte

// MarshalID tells gqlgen how to parse the ID
func MarshalID(id uint64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(fmt.Sprintf("%d", id)))
	})
}

// UnmarshalID tells gqlgen how to unp-arse the ID
func UnmarshalID(v interface{}) (uint64, error) {
	id, ok := v.(string)
	if !ok {
		return 0, fmt.Errorf("ids must be strings")
	}
	i, e := strconv.Atoi(id)
	return uint64(i), e
}

// MarshalID tells gqlgen how to parse the ID
func MarshalCUID(id string) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, id)
	})
}

// UnmarshalID tells gqlgen how to unp-arse the ID
func UnmarshalCUID(v interface{}) (string, error) {
	id, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("ids must be strings")
	}

	return id, nil
}

func MarshalTimestamp(t time.Time) graphql.Marshaler {
	timestamp := t.Unix() * 1000

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(timestamp, 10))
	})
}

func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(int); ok {
		return time.Unix(int64(tmpStr), 0), nil
	}
	return time.Time{}, errors.New("time should be a unix timestamp")
}

func (j JSON) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}

func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		errors.New("invalid scan source")
	}
	*j = append((*j)[0:0], s...)
	return nil
}

func (m JSON) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

func (m *JSON) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("null point exception")
	}
	*m = append((*m)[0:0], data...)
	return nil
}

func (j JSON) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}

func (j JSON) Equals(j1 JSON) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}
