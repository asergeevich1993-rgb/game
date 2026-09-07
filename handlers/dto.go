package handlers

import (
	"errors"
	"go-from-zero/concurrency3/miner"
	"strings"
	"time"
)

type MinerDTO struct {
	Class string `json:"class"`
}

func (md *MinerDTO) ValidateForCreate() error {
	md.Class = strings.TrimSpace(md.Class)

	if md.Class == "" {
		return errors.New("class is empty")
	}
	return nil
}

type ErrorDTO struct {
	Error string    `json:"error"`
	Time  time.Time `json:"time"`
}

type InfoDTO struct {
	Miners    int               `json:"miners"`
	Balance   int               `json:"balance"`
	Allminers []miner.MinerINFO `json:"miner_info"`
}
