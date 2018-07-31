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
		default:
			p.HandleDefault(&update)
		}
	}
}
