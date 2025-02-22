package schemas

import (
	_ "embed"
)

//go:embed user.sql
var UserSchema string
