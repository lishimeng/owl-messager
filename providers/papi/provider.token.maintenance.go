package papi

import (
	"context"
	"time"
)

type TokenMaintenance interface {
	RegisterAccessToken(id string, d time.Time)
}

type tokenMaintenance struct {
	ctx context.Context
	h   *PriorityQueue
}

func (tm *tokenMaintenance) Init() {
	tm.h = NewPq()
}
