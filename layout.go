package gnomoria

import (
	"math"
	"math/rand"
)

const (
	GRID_SIZE = 10
	NUM_TILES = GRID_SIZE * GRID_SIZE
)

const (
	DUPLICATE_WORKSHOP_PENALTY = 20
	STOCKPILE_PENALTY          = 5
	MISSING_WORKSHOP_PENALTY   = 100
	HAULING_DISTANCE_PENALTY   = 1
	GOODS_PENALTY              = 2
	MISSING_STOCKPILE_PENALTY  = 100
)

type Tiles [NUM_TILES]Tile

type Layout struct {
	tiles     Tiles
	fitness   float64
	copied    bool
	locations map[string][]int
}

func (layout *Layout) Evaluate() {
	findLocations(layout)
	layout.fitness = 0
	layout.fitness -= penalizeDuplicateWorkshops(layout)
	layout.fitness -= penalizeMissingWorkshops(layout)
	layout.fitness -= penalizeStockpiles(layout)
	layout.fitness -= penalizeGoods(layout)
	layout.fitness -= penalizeHaulingDistance(layout)
}

func (layout *Layout) Fitness() float64 {
	return layout.fitness
}

func (layout *Layout) Mutate() {
	layout.RandomTile().Mutate()
}

func (layout *Layout) RandomTile() *Tile {
	i, j := rand.Intn(GRID_SIZE), rand.Intn(GRID_SIZE)
	return &layout.tiles[i*GRID_SIZE+j]
}

func (a *Layout) Combine(b *Layout) (*Layout, *Layout) {
	c := rand.Intn(GRID_SIZE)
	var d int
	if rand.Float64() < 0.5 {
		d = 1
	} else {
		d = GRID_SIZE
	}
	for i := 0; i < c*GRID_SIZE; i += d {
		a.tiles[i], b.tiles[i] = b.tiles[i], a.tiles[i]
	}
	return a, b
}

func (layout *Layout) Copy() *Layout {
	if layout.copied {
		theCopy := new(Layout)
		*theCopy = *layout
		return theCopy
	}
	layout.copied = true
	return layout
}

func (layout *Layout) Randomize() {
	for i := 0; i < NUM_TILES; i++ {
		(&layout.tiles[i]).Randomize()
	}
}

func penalizeDuplicateWorkshops(layout *Layout) (penalty float64) {
	workshops := map[string]bool{}
	for i := 0; i < NUM_TILES; i++ {
		tile := &layout.tiles[i]
		if layout.tiles[i].kind == WORKSHOP {
			if workshops[tile.workshop] {
				penalty += DUPLICATE_WORKSHOP_PENALTY
			} else {
				workshops[tile.workshop] = true
			}
		}
	}
	return
}

func penalizeMissingWorkshops(layout *Layout) float64 {
	workshops := map[string]bool{}
	for i := 0; i < NUM_TILES; i++ {
		tile := &layout.tiles[i]
		workshops[tile.workshop] = true
	}
	return float64((NumWorkshops() - len(workshops)) * MISSING_WORKSHOP_PENALTY)
}

func penalizeStockpiles(layout *Layout) float64 {
	count := 0
	for i := 0; i < NUM_TILES; i++ {
		if layout.tiles[i].kind == WORKSHOP {
			count++
		}
	}
	return float64(count * STOCKPILE_PENALTY)
}

func penalizeGoods(layout *Layout) float64 {
	total_count := 0
	for i := 0; i < NUM_TILES; i++ {
		count := 0
		for _, b := range layout.tiles[i].goods {
			if b {
				count++
			}
		}
		if count > 0 {
			total_count += (count - 1)
		}
	}
	return float64(total_count * GOODS_PENALTY)
}

func penalizeHaulingDistance(layout *Layout) (penalty float64) {
	for i := 0; i < NUM_TILES; i++ {
		if layout.tiles[i].kind == WORKSHOP {
			y, x := i/GRID_SIZE, i%GRID_SIZE
			for _, good := range WorkshopGoods(layout.tiles[i].workshop) {
				penalty += distanceToHaul(x, y, layout.locations[good])
			}
		}
	}
	return
}

func findLocations(layout *Layout) {
	layout.locations = make(map[string][]int)
	for i := 0; i < NUM_TILES; i++ {
		for good, b := range layout.tiles[i].goods {
			if b {
				layout.locations[good] = append(layout.locations[good], i)
			}
		}
	}
}

func distanceToHaul(x1, y1 int, locations []int) float64 {
	if locations == nil {
		return MISSING_STOCKPILE_PENALTY
	}
	dist := float64(0)
	for _, i := range locations {
		y2, x2 := i/GRID_SIZE, i%GRID_SIZE
		dist += math.Abs(float64(x2-x1)) + math.Abs(float64(y2-y1))
	}
	return (dist / float64(len(locations))) * HAULING_DISTANCE_PENALTY
}
