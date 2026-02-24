package creator

import "errors"

var (
	ErrInvalidName          = errors.New("invalid creator name")
	ErrInvalidIntroduction  = errors.New("invalid creator introduction")
	ErrInvalidLogoMimeType  = errors.New("invalid creator logo mime type")
	ErrInvalidLogoPath      = errors.New("invalid creator logo path")
	ErrInvalidLogoFileName  = errors.New("invalid creator logo file name")
	ErrCreatorRecordMissing = errors.New("creator is not found")
)
