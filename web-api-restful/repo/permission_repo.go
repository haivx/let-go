package repo

import (
	"final-project/dto"
	model "final-project/model"
)

func CreatePermission(req dto.CreatePermission) (per *model.Permission, err error) {
	per = &model.Permission{
		Name: req.Name,
		Code: req.Code,
	}

	if err := DB.Table("permission").Create(per).Error; err != nil {
		return nil, err
	}

	return per, nil
}

func CreateRolePermission(req dto.CreateRolePermission) (rolePermission *model.RolePermission, err error) {
	rolePermission = &model.RolePermission{
		RoleId:       req.RoleId,
		PermissionId: req.PermissionId,
	}

	if err := DB.Table("role_permission").Create(rolePermission).Error; err != nil {
		return nil, err
	}

	return rolePermission, nil
}
