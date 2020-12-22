package controller

import (
	"github.com/snghnaveen/hello-operator/pkg/controller/traveller"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, traveller.Add)
}
