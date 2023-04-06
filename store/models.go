package store

type Payload struct {
	LongURL string `json:"long_url"`
}

type TidyUrl struct {
	Id       string `db:"id"`
	LongURL  string `db:"long_url"`
	ShortURL string `db:"short_url"`
}
