package handler

import (
	"fmt"

	"github.com/andrysds/panera/entity"
)

// Command handles commands
func Command(command string) string {
	switch command {
	case "birthdays":
		return birthday()
	case "standup":
		return standup()
	case "standup_list":
		return standupList()
	case "standup_new_day":
		return standupNewDay()
	case "standup_skip":
		return standupSkip()
	}
	return ""
}

func birthday() string {
	var users []entity.User
	err := entity.Users.All("name", &users)
	if err != nil {
		return err.Error()
	}
	result := "Birthdays:\n"
	for _, u := range users {
		result += u.Birthday + " - " + u.Name + "\n"
	}
	return result
}

func standup() string {
	standup, err := entity.GetStandup()
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
		if s.State == entity.StandupStateDone {
			message += "`[x]` "
		} else if s.State == entity.StandupStateSkipped {
			message += "`[s]` "
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
	return "standupNewDay ok"
}

func standupSkip() string {
	standup, skipped, err := entity.SkipStandup()
	if err == nil {
		var standupUser, skippedUser *entity.User
		if standupUser, err = standup.User(); err == nil {
			if skippedUser, err = skipped.User(); err == nil {
				return fmt.Sprintf(
					"Karena %s tidak bisa, penggantinya %s (@%s)",
					skippedUser.Name, standupUser.Name, standupUser.Username,
				)
			}
		}
	}
	return err.Error()
}
