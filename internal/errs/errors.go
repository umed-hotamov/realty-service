package errs

import "errors"

var (
	ErrDuplicate          = errors.New("record already exists")
	ErrNotExist           = errors.New("record does not exist")
	ErrUpdateFailed       = errors.New("record update failed")
	ErrDeleteFailed       = errors.New("record delete failed")
	ErrPersistenceFailed  = errors.New("persistence internal error")
)

var (
  ErrInvalidTokenClaims = errors.New("token or claims are invalid")
  ErrInvalidFingerprint = errors.New("invalid client fingerprint")
  ErrInvalidToken       = errors.New("invalid jwt token")
	ErrEmail              = errors.New("email error")
	ErrPassword           = errors.New("password error")
)
