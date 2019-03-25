package message

import "github.com/andrysds/panera/entity"

// Birthdays is message handler for /birthdays
func Birthdays() string {
	var users []entity.User
	err := entity.Users.All("birthday", &users)
	if err != nil {
		return err.Error()
	}
	result := "Birthdays:\n"
	for _, u := range users {
		result += u.Birthday + " (" + u.Name + ")\n"
	}
	return result
}
