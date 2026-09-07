package enterprise

import (
	"context"
	"errors"
	"fmt"
	"go-from-zero/concurrency3/equipment"
	"go-from-zero/concurrency3/miner"
	"sync"
	"time"
)

type Enterprice struct {
	miner       []miner.MinerInterface
	balance     miner.Coal
	totalCoal   miner.Coal
	totalMiners int
	equipment   []equipment.Equipment
	mtx         sync.RWMutex
}

func (e *Enterprice) ShowEquipment() []equipment.Equipment {
	return e.equipment
}
func (e *Enterprice) ShowMinersInfo() []miner.MinerINFO {
	result := make([]miner.MinerINFO, 0, len(e.miner))
	for _, m := range e.miner {
		result = append(result, m.Info())
	}
	return result
}

func (e *Enterprice) BuyEquip(eq *equipment.Equipment) error {
	if e.balance < eq.Price {
		return errors.New("Low balance")
	}
	e.balance -= eq.Price
	eq.Bought = true
	e.equipment = append(e.equipment, *eq)
	fmt.Printf("Купили экипу %s, за %d Угля\n", eq.Name, eq.Price)
	return nil
}

func NewEnterprice() *Enterprice {
	return &Enterprice{
		miner: make([]miner.MinerInterface, 0),
	}
}
func (e *Enterprice) AddBalance(c miner.Coal) {
	defer e.mtx.Unlock()
	e.mtx.Lock()
	e.balance += miner.Coal(c)
}
func (e *Enterprice) GetBalance() miner.Coal {
	defer e.mtx.RUnlock()
	e.mtx.RLock()
	return e.balance
}

func (e *Enterprice) Hire(mi miner.MinerInterface, ctx context.Context) error {
	if e.balance >= mi.MinerCost() {
		e.balance -= mi.MinerCost()
		e.miner = append(e.miner, mi)
		e.totalMiners++
		go func() {
			ch := mi.Run(ctx)
			for c := range ch {
				e.AddBalance(c)
				e.AddTotalCoal(c)
			}

		}()
		return nil
	}
	return errors.New("Недостаточно угля для найма")

}
func (e *Enterprice) AddTotalCoal(c miner.Coal) {
	defer e.mtx.Unlock()
	e.mtx.Lock()
	e.totalCoal += c
}
func (e *Enterprice) TotalMiners() int {
	return e.totalMiners
}
func (e *Enterprice) ShowMiners() []miner.MinerInterface {
	return e.miner
}

func (e *Enterprice) passiveIncome(ctx context.Context) {

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Завершение пассивного дохода")
			return
		case <-ticker.C:

			e.AddBalance(1)
			fmt.Println("Пассивный доход в размере :", 1)

		}
	}
}
func (e *Enterprice) Start(ctx context.Context) {

	go e.passiveIncome(ctx)

}
