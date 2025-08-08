package store

type Store interface {
	CreateTidyUrl(p any) (string, error)
	GetTidyUrlByLongUrl(longUrl string) (string, error)
	GetTidyUrlByShortUrl(shortUrl string) (string, error)
}
