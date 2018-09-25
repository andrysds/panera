package entity

import (
	"strings"
	"time"

	"github.com/andrysds/clarity"
	"github.com/andrysds/panera/db"
)

const BirthdayKey = "panera:birthday"

type Birthday struct {
	Day    int
	Month  time.Month
	Name   string
	UserID int
}

func NewBirthday(data string) *Birthday {
	res := strings.Split(data, ":")
	obj := &Birthday{}
	if len(res) == 4 {
		obj.Day = clarity.Atoi(res[0])
		obj.Month = time.Month(clarity.Atoi(res[1]))
		obj.Name = res[2]
		obj.UserID = clarity.Atoi(res[3])
	}
	return obj
}

func Birthdays(day int, month time.Month) ([]*Birthday, error) {
	data, err := db.Redis.LRange(BirthdayKey, 0, -1).Result()
	result := []*Birthday{}
	getAll := day == 0 && month == 0
	for _, d := range data {
		birthday := NewBirthday(d)
		if getAll || (day == birthday.Day && month == birthday.Month) {
			result = append(result, birthday)
		}
	}
	return result, err
}
