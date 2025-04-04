package user

func PresenterUsers(model []UserModel) []User {
	users := make([]User, len(model))

	for _, user := range model {
		users = append(users, User{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		})
	}

	return users
}
