package service

import "github.com/fabian4/Fyi_sever/utils"

type PushService struct {
}

func (p *PushService) DoPush() {
	utils.Post()
}
