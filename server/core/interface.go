package core

type IApiCore interface {
	Get(key []byte) (value []byte, err error)
	Set(key, value []byte) error
}

type ApiCore struct {
	IApiCore
}
