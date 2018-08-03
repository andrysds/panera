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
			}
		}

		// master
		if update.Message.Chat.ID == p.MasterId {
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "init_standup":
					p.HandleInitStandup(&update)
				case "init_standup_list":
					p.HandleInitStandup(&update)
				}
			} else {
				p.HandleMasterMessage(&update)
			}
		}
	}
}
