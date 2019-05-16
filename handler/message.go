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
	case "/standup_new_day":
		return standupNewDay()
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
	standup, err := entity.GetTodayStandup()
	if err == nil {
		var user *entity.User
		user, err = standup.User()
		if err == nil {
			return fmt.Sprintf(
				"Yang dapat giliran untuk memimpin stand up hari ini adalah %s (%s)",
				user.Name, user.Username,
			)
		}
	}
	return err.Error()
}

func standupList() string {
	standups, err := entity.GetStandupList()
	if err != nil {
		return err.Error()
	}
	message := "Stand up lead periode ini:"
	for _, s := range standups {
		message += "\n"
		if s.State == "done" {
			message += "`[x]` "
		} else {
			message += "`[ ]` "
		}
		user, err := s.User()
		if err != nil {
			return err.Error()
		}
		message += user.Name
	}
	return message
}

func standupNewDay() string {
	if err := entity.NewDayStandup(); err != nil {
		return err.Error()
	}
	return "ok"
}
