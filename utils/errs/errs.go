package errs

import "errors"

var (
	ErrInvalidToken       = errors.New("invalid token")
	ErrNoReportsFound     = errors.New("no reports found")
	ErrSomethingWentWrong = errors.New("something went wrong")
	ErrUserNotFound       = errors.New("user not found")
	ErrSongNotFound       = errors.New("song not found")
	ErrAlbumNotFound      = errors.New("album not found")
	ErrArtistNotFound     = errors.New("artist not found")
	ErrPermissionDenied   = errors.New("permission denied")
	ErrValidationFailed   = errors.New("validation failed")
	ErrUsernameExists     = errors.New("username already exists")
	ErrRecordNotFound     = errors.New("record not found")
	ErrUserBlocked        = errors.New("user blocked")
)
