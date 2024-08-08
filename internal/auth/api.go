package auth

import "errors"

// ErrNotFound - Error to return when something is not found
var ErrNotFound = errors.New("not found")

// ErrUnknown - Error to return when an unknown server error occurs
var ErrUnknown = errors.New("an unknown server error occurred, please try again")

var ErrNoAuthToken = errors.New("no access token available. please login with 'deepgram login'")
