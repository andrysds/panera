package app

func (p *Panera) Run() {
	for update := range p.Updates {
		if update.Message == nil {
			continue
		}
		p.LogMessage(update.Message)

		switch {
		case update.Message.NewChatMembers != nil:
			p.HandleGroupInvitation(&update)
		case update.Message.IsCommand():
			switch update.Message.Command() {
			case "standup":
				p.HandleStandup(&update)
			}
		case update.Message.Chat.ID == p.MasterId:
			p.HandleMasterMessage(&update)
		}
	}
}
