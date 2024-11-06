package rest

import (
	"crypto/sha256"
	"github.com/beauxarts/fedorov/data"
	"github.com/boggydigital/kevlar"
	"github.com/boggydigital/middleware"
)

var (
	rdx kevlar.ReadableRedux
)

func SetUsername(role, u string) {
	middleware.SetUsername(role, sha256.Sum256([]byte(u)))
}

func SetPassword(role, p string) {
	middleware.SetPassword(role, sha256.Sum256([]byte(p)))
}

func Init() error {

	var err error

	if rdx, err = data.NewReduxReader(data.ReduxProperties()...); err != nil {
		return err
	}

	return err
}
