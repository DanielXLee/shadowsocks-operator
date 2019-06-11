package controller

import (
	"github.com/DanielXLee/shadowsocks-operator/pkg/controller/sss"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, sss.Add)
}
