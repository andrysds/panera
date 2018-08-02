package migrate

import (
	"github.com/andrysds/panera/db"
	"github.com/andrysds/panera/entity"
)

func InitStandup() string {
	initData := []string{
		"Herry:herrydev:1",
		"Olvi:olvilora:1",
		"Isti:tianaulia:1",
		"Setia:setiasimaremare:1",
		"Rifa:rifaMukhlisa:1",
		"Yohanes:yohanes77:1",
		"Regina:regina\\_avena:0",
		"Ben:benlemueltanasale:0",
		"Farida:faridaamila:0",
		"Adimas:addimas:0",
		"Andrys:andrysds:0",
		"Ai:ayshzkh:0",
		"Luthfi:luthfift:0",
	}

	_, err := db.Redis.Del(entity.StandupKey).Result()
	if err != nil {
		return err.Error()
	}

	for _, d := range initData {
		_, err = db.Redis.LPush(entity.StandupKey, d).Result()
		if err != nil {
			return err.Error()
		}
	}

	return "ok"
}
