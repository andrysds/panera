package standup

import (
	"fmt"
	"strings"
)

const Key = "panera:standup"

type Standup struct {
	Name     string
	Username string
	State    string
	Order    int
}

func NewStandup(data string, order int) *Standup {
	res := strings.Split(data, ":")
	obj := &Standup{}
	if len(res) == 3 {
		obj.Name = res[0]
		obj.Username = res[1]
		obj.State = res[2]
		obj.Order = order
	}
	return obj
}

func (self *Standup) Raw() string {
	return fmt.Sprintf("%s:%s:%s", self.Name, self.Username, self.State)
}
