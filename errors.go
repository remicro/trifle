package trifle

import "errors"

func UnexpectedError() (err error) {
	return errors.New("unexpected error: " + StringN(16))
}
