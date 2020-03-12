package usecases

import "github.com/sr-2020/eva-gateway/app/entity"

func JoinUsers(positionUsers []entity.PositionUser, authUsers []entity.AuthUser) []entity.Positions {
	result := []entity.Positions{}

	var authMap = make(map[int]entity.AuthUser, 0)
	for _, v := range authUsers {
		authMap[v.Id] = v
	}

	temp := entity.Positions{}
	for _, v := range positionUsers {
		if val, ok := authMap[v.Id]; ok {
			temp.Join(val, v)
		} else {
			temp.Join(val, entity.PositionUser{})
		}

		result = append(result, temp)
	}

	return result
}

func Profile(authUser entity.AuthUser, positionUser entity.PositionUser) entity.ProfileUser {

	result := entity.ProfileUser{}
	result.Join(authUser, positionUser)

	return result
}
