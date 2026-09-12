package role

import (
	"code.gitea.io/gitea/models"
	pg "github.com/lib/pq"
)

func AddUserRole(userId int64, roleType models.RoleType) error {
	role := GetRole(roleType)
	if role == nil {
		return models.ErrRoleNotExists{}
	}
	_, err := models.NewUserRole(models.UserRole{UserId: userId, RoleType: roleType})
	if err != nil {
		e := err.(*pg.Error)
		//23505 is postgrey error code for unique_violation
		//see https://www.postgresql.org/docs/current/errcodes-appendix.html
		//if unique_violation,user role record exists,just return
		if e.Code == "23505" {
			return nil
		}
	}
	return err
}

func DeleteUserRole(userId int64, roleType models.RoleType) error {
	role := GetRole(roleType)
	if role == nil {
		return models.ErrRoleNotExists{}
	}
	_, err := models.DeleteUserRole(roleType, userId)
	return err
}

func UserHasRole(userId int64, roleType models.RoleType) bool {
	role := GetRole(roleType)
	if role == nil {
		return false
	}
	_, err := models.GetUserRoleByUserAndRole(userId, roleType)
	if err != nil {
		return false
	}
	return true
}
