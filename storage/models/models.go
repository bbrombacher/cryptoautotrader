package models

import (
	"time"

	jsoninter "github.com/json-iterator/go"
)

type UserEntry struct {
	ID        string `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`

	CursorID  int        `db:"cursor_id,omitempty" json:"cursor_id"`
	CreatedAt *time.Time `db:"created_at,omitempty" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at,omitempty" json:"updated_at"`
}

func (e *UserEntry) RetrieveTagValues(tag string) (map[string]interface{}, error) {
	tagMap := map[string]interface{}{}
	var rjson = jsoninter.Config{TagKey: tag}.Froze()
	data, err := rjson.Marshal(e)
	if err != nil {
		return nil, err
	}
	if err := rjson.Unmarshal(data, &tagMap); err != nil {
		return nil, err
	}
	return tagMap, nil
}
