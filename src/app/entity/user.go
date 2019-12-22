package entity

type AuthUser struct {
	Id        int         `json:"id"`
	Amount    int         `json:"amount"`
	Followers []int       `json:"followers"`
	Following []int       `json:"following"`
	Admin     bool        `json:"admin"`
	Status    string      `json:"status"`
	Role      string      `json:"role"`
	Items     interface{} `json:"items"`
	Name      string      `json:"name"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
}

type ProfileUser struct {
	Id                int               `json:"id"`
	Amount            int               `json:"amount"`
	Followers         []int             `json:"followers"`
	Following         []int             `json:"following"`
	Admin             bool              `json:"admin"`
	Status            string            `json:"status"`
	Role              string            `json:"role"`
	Items             interface{}       `json:"items"`
	Name              string            `json:"name"`
	CreatedAt         string            `json:"created_at"`
	UpdatedAt         string            `json:"updated_at"`
	LocationUpdatedAt string            `json:"location_updated_at"`
	LocationId        *int              `json:"location_id"`
	Location          *PositionLocation `json:"location"`
}

type AuthLogin struct {
	Email         string `json:"email"`
	Password      string `json:"password"`
	FirebaseToken string `json:"firebase_token"`
}

type AuthUserToken struct {
	Id     int    `json:"id"`
	ApiKey string `json:"api_key"`
}

type PushToken struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

func (pu *ProfileUser) Join(u AuthUser, p PositionUser) {
	pu.Id = u.Id
	pu.Admin = u.Admin
	pu.Status = u.Status
	pu.Name = u.Name
	pu.CreatedAt = u.CreatedAt
	pu.UpdatedAt = u.UpdatedAt

	pu.Amount = u.Amount
	pu.Followers = u.Followers
	pu.Following = u.Following
	pu.Role = u.Role
	pu.Items = u.Items

	pu.LocationUpdatedAt = p.UpdatedAt
	pu.LocationId = p.LocationId
	pu.Location = p.Location
}
