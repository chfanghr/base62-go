package base62_go

const encodeInverted = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NewInvertedEncoding() *Encoding {
	return NewEncoding(encodeInverted)
}

var InvertedEncoding = NewInvertedEncoding()
