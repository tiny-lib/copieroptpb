package typeconverter

import "errors"

var (
	typeNotMatchError     = errors.New("type not match")
	timestampPbIsNotValid = errors.New("timestampPb Is Not Valid")
)
