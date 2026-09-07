package miner

import (
	"context"
	"fmt"
	"time"
)

type MinerInterface interface {
	Run(ctx context.Context) <-chan Coal
	Info() MinerINFO
	MinerCost() Coal
}
type Coal int

type Miner struct {
	Class    string        `json:"class"`
	Energy   int           `json:"energy"`
	Income   int           `json:"income"`
	Interval time.Duration `json:"interval,omitempty"`
	Cost     Coal          `json:"cost"`
}
type MinerINFO struct {
	Class  string
	Energy int
	Income int
	Cost   Coal
}

func NewMainer(class string) *Miner {
	switch class {
	case "low":
		return &Miner{
			Class:    "low",
			Energy:   30,
			Income:   1,
			Interval: 3 * time.Second,
			Cost:     5,
		}
	case "middle":
		return &Miner{
			Class:    "middle",
			Energy:   45,
			Income:   3,
			Interval: 2 * time.Second,
			Cost:     50,
		}
	case "senior":
		return &Miner{
			Class:    "senior",
			Energy:   60,
			Income:   10,
			Interval: 1 * time.Second,
			Cost:     450,
		}
	}
	return nil
}

func (m *Miner) Info() MinerINFO {
	return MinerINFO{
		Class:  m.Class,
		Energy: m.Energy,
		Income: m.Income,
		Cost:   m.Cost,
	}
}

func (m *Miner) Run(ctx context.Context) <-chan Coal {
	out := make(chan Coal)

	go func() {
		defer close(out)
		for m.Energy > 0 {
			select {
			case <-ctx.Done():
				fmt.Println("Сбор окончен")
				return
			default:
				var sum Coal
				m.Energy--

				out <- Coal(m.Income)
				if m.Class == "senior" {
					m.Income = m.Income + 3
				}
				fmt.Printf("Добыто угля: %d, Miner: %s\n", m.Income, m.Class)
				time.Sleep(m.Interval)
				sum += Coal(m.Income)
				if m.Energy == 0 {
					fmt.Printf("Собрано угля %d шахтером %s\n", sum, m.Class)
				}
			}
		}

	}()
	return out
}
func (m *Miner) MinerCost() Coal {
	return m.Cost
}
