package jwt

import (
	"strings"

	"github.com/sohaha/zlsgo/zerror"
)

func ParseError(err error) error {
	if err == nil {
		return nil
	}

	errMsg := err.Error()
	if !strings.Contains(errMsg, "expired") {
		return zerror.InvalidInput.Wrap(err, "token is illegal")
	}

	return zerror.Unauthorized.Text(errMsg)
}
