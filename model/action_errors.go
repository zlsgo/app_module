package model

import (
	"github.com/sohaha/zlsgo/zerror"
)

var (
	errRelationMismatch = zerror.WrapTag(zerror.InvalidInput)
	errDataValidation   = zerror.WrapTag(zerror.InvalidInput)
	errDecryptionFailed = zerror.WrapTag(zerror.InvalidInput)
)
