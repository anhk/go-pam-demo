package main

import "io"

type AuthResult int

const (
	// AuthError is a failure.
	AuthError AuthResult = iota
	// AuthSuccess is a success.
	AuthSuccess
)

func pamAuthenticate(w io.Writer, uid int, username string, argv []string) AuthResult {
	log("PamAuthenticate [%v,%v,%v]", username, uid, argv)
	return AuthSuccess
}

func main() {}
