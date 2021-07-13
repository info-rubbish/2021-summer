package database

func CheckPermission(id string) error {
	user := &User{}
	if err := DB.First(user, "id=?", id).Error; err != nil {
		return err
	}
	if user.Permission < 2 {
		return ErrPermissionDenied
	}
	return nil
}
