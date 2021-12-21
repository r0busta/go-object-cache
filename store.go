package goobjectstore

type Reader interface {
	Read(v interface{}) error
}

type Writer interface {
	Write(v interface{}) error
}
