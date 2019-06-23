package base62_go

const encodeStd = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func NewStdEncoding() *Encoding {
	return NewEncoding(encodeStd)
}

var StdEncoding = NewStdEncoding()
