package cache

import (
	"github.com/orcaman/concurrent-map/v2"
	"server/internal/structures"
)

var Map cmap.ConcurrentMap[string, structures.Train]

func init() {
	Map = cmap.New[structures.Train]()
}