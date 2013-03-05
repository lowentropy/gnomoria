package gnomoria

import (
	"math/rand"
)

const (
	EMPTY = iota
	WORKSHOP
	STOCKPILE
	NUM_KIND
)

type Tile struct {
	kind     int
	workshop string
	goods    map[string]bool
}

func (tile *Tile) Mutate() {
	if rand.Float64() < 0.2 {
		tile.Randomize()
		return
	}
	switch tile.kind {
	case WORKSHOP:
		tile.MutateWorkshop()
	case STOCKPILE:
		tile.MutateStockpile()
	}
}

func (tile *Tile) Randomize() {
	tile.kind = rand.Intn(NUM_KIND)
	tile.RandomizeSettings()
}

func (tile *Tile) RandomizeSettings() {
	switch tile.kind {
	case WORKSHOP:
		tile.MutateWorkshop()
	case STOCKPILE:
		tile.goods = make(map[string]bool)
		tile.MutateStockpile()
	}
}

func (tile *Tile) MutateWorkshop() {
	tile.workshop = RandomWorkshop()
}

func (tile *Tile) MutateStockpile() {
	n := rand.Intn(3) + 1
	for i := 0; i < n; i++ {
		good := RandomGood()
		tile.goods[good] = !tile.goods[good]
	}
}
