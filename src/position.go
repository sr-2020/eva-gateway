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

	ProxyOld(r, cfg.Position + "/api/v1/users", &positionUsers)
	ProxyOld(r, cfg.Auth + "/api/v1/users", &authUsers)

	var positionMap = make(map[int]PositionUser, 0)
	for _, v := range positionUsers {
		positionMap[v.Id] = v
	}

	temp := Positions{}
	for _, v := range authUsers {
		if val, ok := positionMap[v.Id]; ok {
			temp.join(v, val)
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

	ProxyOld(r, cfg.Position + "/api/v1/positions", &position)

	response, err := json.Marshal(position)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(response))
}
