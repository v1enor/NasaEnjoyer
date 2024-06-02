package domain

import (
	"errors"
)

var (
	ErrAPODNotFound         = errors.New("apod not found")
	ErrCreateRequest        = errors.New("error creating request")
	ErrFetchData            = errors.New("error fetching data")
	ErrDecodeResponse       = errors.New("error decoding response body")
	ErrRequestCreation      = errors.New("error creating request")
	ErrImageDownload        = errors.New("error downloading image")
	ErrUnexpectedStatusCode = errors.New("unexpected status code")
	ErrNotAnImage           = errors.New("content is not an image")
	ErrDirectoryCreation    = errors.New("could not create directory")
)
