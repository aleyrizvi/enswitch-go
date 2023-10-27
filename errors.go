package enswitch

import "errors"

var (
	ErrRequestWithContext = errors.New("unable to create new http request with context provided")
	ErrHTTPRequest        = errors.New("unable to perform http request")
	ErrNotFound           = errors.New("this does not exist or does not belong to you")
	ErrBadURL             = errors.New("invalid base url")
	ErrDecodingRequest    = errors.New("unable to decode")
	ErrContextNil         = errors.New("context must not be nil")
)
