package constant

import "errors"

const ErrNoDataFound = "data tidak ditemukan"
const ErrInvalidORExpiredJWT = "Invalid or expired JWT"
const ErrMissingOrMalformedJWT = "Missing or malformed JWT"
const ErrUserUnauthorized = "Anda tidak mempunyai akses untuk aplikasi ini"

var ErrWrongPassword = errors.New("password yang anda masukkan salah")
var ErrUserNotFound = errors.New("user yang anda masukkan tidak ada di dalam database")

// Error Validator golang

const ErrFieldStartsWith = "startswith"
