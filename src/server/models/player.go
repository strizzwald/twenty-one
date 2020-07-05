package models

import "github.com/google/uuid"

type Player struct {
	Id   uuid.UUID `json:"id"`
	Hand []Card
}

func (p *Player) SetId(id uuid.UUID) {
	p.Id = id
}
