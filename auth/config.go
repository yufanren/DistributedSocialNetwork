package auth

import "time"

const (
	// port of auth service
	port = ":8010"
	// token validity duration
	ttl = 1000000 * time.Second
	// hash-based message authentication code
	hmacSecret = "not-secure"
)
