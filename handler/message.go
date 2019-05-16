package handler

import (
	"fmt"

	"github.com/andrysds/panera/entity"
)

// Message is router function for message
func Message(command string) string {
	switch command {
	case "/birthdays":
		return birthday()
	case "/standup":
		return standup()
	case "/standup_list":
		return standupList()
	}
	return ""
}

func birthday() string {
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

func standup() string {
	standup := entity.GetTodayStandup()
	return fmt.Sprintf(
		"Yang dapat giliran untuk memimpin stand up hari ini adalah %s (%s)",
		standup.User().Name, standup.User().Username,
	)
}

func standupList() string {
	standups := entity.GetStandupList()
	message := "Stand up lead periode ini:"
	for _, s := range standups {
		message += "\n"
		if s.State == "1" {
			message += "`[x]` "
		} else {
			message += "`[ ]` "
		}
		message += s.User().Name
	}
	return message
}
