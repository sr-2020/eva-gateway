package entity

type PositionLocation struct {
	Id    int    `json:"id"`
	Label string `json:"label"`
	Options map[string]interface{}
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

func (pl *Positions) Join(u AuthUser, p PositionUser) {
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
