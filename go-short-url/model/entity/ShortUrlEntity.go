package entity

type ShortUrlEntity struct {
	BaseEntity
	ShortUrl  string
	NativeUrl string
	ValidTime uint64
}

func (ShortUrlEntity) TableName() string {
	return "t_short_url"
}
