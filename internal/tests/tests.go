package tests

import (
	"errors"

	"github.com/theTardigrade/validation/internal/data"
)

var (
	ErrUnexpectedType = errors.New("unexpected type")
)

type Handler (func(m *data.Main) error)
