package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"net/http"
)

type Middleware func(http.ResponseWriter, *http.Request, httprouter.Params) (*http.Response, error)
type NewMiddleware func(http.ResponseWriter, *http.Request, httprouter.Params)

type Middlewares struct {
	Billing ServiceMiddleware
	Auth ServiceMiddleware
	Position ServiceMiddleware
	Push ServiceMiddleware
}

type ServiceMiddleware struct {
	Global []alice.Constructor
	Route map[string][]alice.Constructor
}

var MiddlewareMap = Middlewares{
	Auth: ServiceMiddleware{
		Global: nil,
		Route: map[string][]alice.Constructor{
			"/profile": {
				NewAuthMiddleware,
			},
			"/login": {
				NewLoginMiddleware,
			},
		},
	},
	Billing: ServiceMiddleware{
		Global: []alice.Constructor{
			NewAuthMiddleware,
		},
		Route: map[string][]alice.Constructor{
			"/account_info": {
				NewAccountInfoMiddleware,
			},
			"/transfer": {
				NewTransferMiddleware,
			},
		},
	},
	Position: ServiceMiddleware{
		Global: []alice.Constructor{
			NewAuthMiddleware,
		},
		Route: nil,
	},
	Push: ServiceMiddleware{
		Global: []alice.Constructor{
			NewAuthMiddleware,
		},
		Route: nil,
	},
}
