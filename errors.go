package enswitch

import "errors"

var (
	ErrRequestWithContext = errors.New("unable to create new http request with context provided")
	ErrHTTPRequest        = errors.New("unable to perform http request")
	ErrBadURL             = errors.New("invalid base url")
	ErrDecodingRequest    = errors.New("unable to decode")
)
