package store

import "time"

type TidyUrl struct {
	Id         string
	LongUrl    string
	EncodedStr string
	CreatedAt  time.Time
}
