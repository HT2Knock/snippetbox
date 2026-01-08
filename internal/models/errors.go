package models

import "errors"

var ErrNoRecord = errors.New("model: no matching user record")

var ErrInvalidCredentials = errors.New("model: invalid credentials")

var ErrDuplicateEmail = errors.New("model: duplicate email")
