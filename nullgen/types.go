package nullgen

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type UUIDType struct {
	Value uuid.UUID
	Null  uuid.NullUUID
}

type TimeType struct {
	Value time.Time
	Null  sql.NullTime
}

type BoolType struct {
	Value bool
	Null  sql.NullBool
}

type StringType struct {
	Value string
	Null  sql.NullString
}

type Int64Type struct {
	Value int64
	Null  sql.NullInt64
}
