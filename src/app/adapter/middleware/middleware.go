package middleware

import (
	"github.com/justinas/alice"
)

type Middleware = alice.Constructor

type ServiceMiddleware struct {
	Global []Middleware
	Route  map[string][]Middleware
}
