package tests

import (
	"errors"

	"github.com/theTardigrade/validation/data"
)

var (
	ErrUnexpectedType = errors.New("unexpected type")
)

type Handler (func(m *data.Main) error)
