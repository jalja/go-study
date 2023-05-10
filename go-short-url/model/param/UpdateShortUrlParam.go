package param

type UpdateShortUrlParam struct {
	Id        uint64
	ShortUrl  string
	NativeUrl string
	ValidTime uint64
}
