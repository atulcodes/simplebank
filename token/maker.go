package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken creates a token for specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)
	// VerifyToken checks whether the token is valid or not
	VerifyToken(token string) (*Payload, error)
}