package handlers

import (
	"context"
	"encoding/json"
	"go-from-zero/concurrency3/enterprise"
	"go-from-zero/concurrency3/equipment"
	"go-from-zero/concurrency3/miner"
	"net/http"
	"time"
)

type HttpHandlers struct {
	enterprice *enterprise.Enterprice
	cancel     context.CancelFunc
	ctx        context.Context
}

func NewHandler(e *enterprise.Enterprice, cancel context.CancelFunc, ctx context.Context) *HttpHandlers {
	return &HttpHandlers{
		enterprice: e,
		cancel:     cancel,
		ctx:        ctx,
	}
}

func (hh *HttpHandlers) HandleCreateMiner(w http.ResponseWriter, r *http.Request) {
	var md MinerDTO
	if err := json.NewDecoder(r.Body).Decode(&md); err != nil {
		writeErrors(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := md.ValidateForCreate(); err != nil {
		writeErrors(w, err.Error(), http.StatusBadRequest)
		return
	}
	miners := miner.NewMainer(md.Class)
	if err := hh.enterprice.Hire(miners, hh.ctx); err != nil {
		writeErrors(w, err.Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(miners)
}
func (hh *HttpHandlers) HandlerBuyEquipment(w http.ResponseWriter, r *http.Request) {
	equip := r.PathValue("name")
	newequip := equipment.NewEquip(equip)
	if newequip == nil {
		writeErrors(w, "equipment not found", http.StatusNotFound)
		return
	}
	if err := hh.enterprice.BuyEquip(newequip); err != nil {
		writeErrors(w, err.Error(), http.StatusConflict)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]string{"status": "buy"})

}
func (hh *HttpHandlers) HandleShowEquipment(w http.ResponseWriter, r *http.Request) {
	equip := hh.enterprice.ShowEquipment()
	if len(equip) == 0 {
		writeErrors(w, "not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(equip)
}
func (hh *HttpHandlers) HandleShowMiners(w http.ResponseWriter, r *http.Request) {
	miners := hh.enterprice.ShowMiners()
	if len(miners) == 0 {
		writeErrors(w, "nof found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(miners)
}
func (hh *HttpHandlers) HandleShowAllInfo(w http.ResponseWriter, r *http.Request) {
	balance := hh.enterprice.GetBalance()

	miners := hh.enterprice.ShowMinersInfo()
	if len(miners) == 0 {
		writeErrors(w, "nof found", http.StatusNotFound)
		return
	}

	countminers := hh.enterprice.TotalMiners()
	if countminers <= 0 {
		writeErrors(w, "not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(InfoDTO{
		Miners:    countminers,
		Balance:   int(balance),
		Allminers: miners,
	})

}
func (hh *HttpHandlers) HandleFinish(w http.ResponseWriter, r *http.Request) {
	hh.cancel()
	w.WriteHeader(200)
}

func writeErrors(w http.ResponseWriter, err string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorDTO{
		Error: err,
		Time:  time.Now(),
	})
}
