# qp
[![Go Report Card](https://goreportcard.com/badge/github.com/alexandergrom/null)](https://goreportcard.com/report/github.com/alexandergrom/null) [![GoDoc](https://godoc.org/github.com/alexandergrom/null?status.svg)](https://godoc.org/github.com/alexandergrom/null)

Package null is a simple package for dealing with nullable SQL and JSON values.

```go
type User struct {
	ID          int64        `json:"id"`
	CountryID   null.Int64   `json:"country_id"`
	LocaleID    int64        `json:"locale_id"`
	Email       string       `json:"email"`
	Phone       null.String  `json:"phone"`
	Rating      null.Float64 `json:"rating"`
	IsTrial     bool         `json:"is_trial"`
	IsActivated null.Bool    `json:"is_activated"`
	Timezone    string       `json:"timezone"`
	CreatedAt   time.Time    `json:"created_at"`
	DeletedAt   null.Time    `json:"deleted_at"`
}

```
