package app

func (p *Panera) Run() {
	for update := range p.Updates {
		if update.Message == nil {
			continue
		}
		p.LogMessage(update.Message)

		// public
		switch {
		case update.Message.NewChatMembers != nil:
			p.HandleGroupInvitation(&update)
		case update.Message.IsCommand():
			switch update.Message.Command() {
			case "standup":
				p.HandleStandup(&update)
			case "standup_list":
				p.HandleStandupList(&update)
			case "standup_skip":
				p.HandleStandupSkip(&update)
			default:
				if update.Message.Chat.ID == p.MasterId {
					p.HandleMasterCommand(update.Message.Command())
				}
			}
		case update.Message.Chat.ID == p.MasterId:
			p.HandleMasterMessage(&update)
		}
	}
}
