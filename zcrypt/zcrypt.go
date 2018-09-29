package zcrypt

type Zcrypt interface {
	Exec([]byte) ([]byte, error)
}