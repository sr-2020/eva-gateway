package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type PositionLocation struct {
	Id    int    `json:"id"`
	Label string `json:"label"`
}

type PositionUser struct {
	Id         int               `json:"id"`
	LocationId *int              `json:"location_id"`
	CreatedAt  string            `json:"created_at"`
	UpdatedAt  string            `json:"updated_at"`
	Location   *PositionLocation `json:"location"`
}

type Positions struct {
	Id                int               `json:"id"`
	Admin             bool              `json:"admin"`
	Status            string            `json:"status"`
	Name              string            `json:"name"`
	CreatedAt         string            `json:"created_at"`
	UpdatedAt         string            `json:"updated_at"`
	LocationUpdatedAt string            `json:"location_updated_at"`
	LocationId        *int              `json:"location_id"`
	Location          *PositionLocation `json:"location"`
}

func (pl *Positions) join(u AuthUser, p PositionUser) {
	pl.Id = u.Id
	pl.Admin = u.Admin
	pl.Status = u.Status
	pl.Name = u.Name
	pl.CreatedAt = u.CreatedAt
	pl.UpdatedAt = u.UpdatedAt

	pl.LocationUpdatedAt = p.UpdatedAt
	pl.LocationId = p.LocationId
	pl.Location = p.Location
}

func GetUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	resp := []Positions{}
	var positionUsers []PositionUser
	var authUsers []AuthUser

	Proxy(r, cfg.Position + "/api/v1/users", &positionUsers, nil)
	Proxy(r, cfg.Auth + "/api/v1/users", &authUsers, nil)

	temp := Positions{}
	for i, v := range authUsers {
		if len(positionUsers) > i {
			temp.join(v, positionUsers[i])
		} else {
			temp.join(v, PositionUser{})
		}
		resp = append(resp, temp)
	}

	response, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(response))
}

func PostPositions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Auth(r)

	var position PositionUser

	Proxy(r, cfg.Position + "/api/v1/positions", &position, nil)

	response, err := json.Marshal(position)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(response))
}

func PositionService(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	path := "/api/v1" + ps.ByName("path")

	Auth(r)

	var resp interface{}
	res, err := ProxyLite(r, cfg.Position + path, &resp)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(res.StatusCode)
	response, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(response))
}
