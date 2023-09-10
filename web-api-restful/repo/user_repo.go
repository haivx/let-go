package repo

import (
	"final-project/dto"
	model "final-project/model"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func CreateUser(req *model.CreateUser) (user *model.User, err error) {
	user = &model.User{
		Username:    req.UserName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := DB.Table("user").Create(user).Error; err != nil {
		return nil, err
	}

	userRole := &model.UserRole{
		UserId: user.Id,
		RoleId: 1, // USER
	}

	if err := DB.Table("user_role").Create(userRole).Error; err != nil {
		return nil, err
	}
	_ = DB.Preload("Roles").First(&user)

	return user, nil
}

func GetUser(email string) (user *model.User, err error) {
	err = DB.Preload("Roles").First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUserRole(dataUserRole dto.UpdateUserRole) (userRole *model.UserRole, err error) {
	userRole = &model.UserRole{
		UserId: dataUserRole.UserId,
		RoleId: dataUserRole.RoleId,
	}

	if err := DB.Table("user_role").Create(userRole).Error; err != nil {
		return nil, err
	}

	return userRole, nil
}

func GetUserList() ([]*model.User, error) {
	users := []*model.User{}
	err := DB.Preload("Roles").Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserPermission(email interface{}) (permissions *[]model.Permission, err error) {
	collectBuilder := DB.Table("permission")
	collectBuilder.Joins("INNER JOIN \"role_permission\" on \"role_permission\".\"permission_id\" = \"permission\".\"id\"")
	collectBuilder.Joins("INNER JOIN \"role\" on \"role\".\"id\" = \"role_permission\".\"role_id\"")
	collectBuilder.Joins("INNER JOIN \"user_role\" on \"user_role\".\"role_id\" = \"role\".\"id\"")
	collectBuilder.Joins("INNER JOIN \"user\" on \"user\".\"id\" = \"user_role\".\"user_id\"")
	collectBuilder.Where("\"user\".\"email\" = ?", email)

	if err = collectBuilder.Find(&permissions).Error; err != nil {
		return nil, err
	}

	return permissions, nil
}

func UpdateUser(dataUser dto.UpdateUser) (newUser *model.User, err interface{}) {
	newUser, err = GetUser(dataUser.Email)
	if newUser == nil {
		return nil, "user's not exist"
	}
	if err != nil {
		return nil, err
	}
	newUser = &model.User{
		Email:       dataUser.Email,
		PhoneNumber: dataUser.PhoneNumber,
		Username:    dataUser.Username,
	}
	collectBuilder := DB.Table("user")
	collectBuilder.Where("\"user\".\"email\" = ?", dataUser.Email)
	collectBuilder.Updates(&newUser)

	return newUser, nil
}

func DeleteUser(id string) (ok interface{}, err error) {
	transaction := DB.Begin()

	collectBuilderUserRole := DB.Table("user_role")
	if err = collectBuilderUserRole.Where("user_id = ?", id).Delete(&model.UserRole{}).Error; err != nil {
		return nil, err
	}

	collectBuilderUser := DB.Table("user")
	if err = collectBuilderUser.Delete(&model.User{}, id).Error; err != nil {
		return nil, err
	}

	transaction.Commit()
	err = transaction.Commit().Error
	if err != nil {
		return nil, err
	}
	return "ok", nil
}

func DumpUser() (err error) {
	transaction := DB.Begin()
	for i := 3; i < 10; i++ {
		user := &model.User{
			Username:    gofakeit.Username(),
			Email:       gofakeit.Email(),
			PhoneNumber: gofakeit.Phone(),
			Password:    gofakeit.Password(true, true, true, true, false, 60),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		if err := DB.Table("user").Create(user).Error; err != nil {
			return err
		}
	}

	for i := 3; i < 10; i++ {
		role := &model.Role{
			Name: gofakeit.Snack(),
		}
		if err := DB.Table("role").Create(role).Error; err != nil {
			return err
		}
	}

	for i := 3; i < 10; i++ {
		userRole := &model.UserRole{
			UserId: i,
			RoleId: 2,
		}
		if err := DB.Table("user_role").Create(userRole).Error; err != nil {
			return err
		}
	}

	transaction.Commit()
	return
}
