package usecases

import "github.com/sr-2020/eva-gateway/app/entity"

func JoinUsers(positionUsers []entity.PositionUser, authUsers []entity.AuthUser) []entity.Positions {
	result := []entity.Positions{}

	var positionMap = make(map[int]entity.PositionUser, 0)
	for _, v := range positionUsers {
		positionMap[v.Id] = v
	}

	temp := entity.Positions{}
	for _, v := range authUsers {
		if val, ok := positionMap[v.Id]; ok {
			temp.Join(v, val)
		} else {
			temp.Join(v, entity.PositionUser{})
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
