package repo

import (
	dto "final-project/dto"
	model "final-project/model"
)

func CreateRole(req *dto.CreateRole) (role *model.Role, err error) {
	role = &model.Role{
		// Id:   util.NewID(),
		Name: req.Name,
	}

	if err := DB.Table("role").Create(role).Error; err != nil {
		return nil, err
	}

	return role, nil
}
