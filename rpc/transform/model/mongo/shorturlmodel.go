package mongo

type ShortUrl struct {
	Shorten string `o:"find,get,set,del" c:"短链"`
	Url     string `o:"find,get,set,del" c:"长链"`
}
