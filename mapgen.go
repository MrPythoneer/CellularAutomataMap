package mapgen

import (
	"math/rand"
)

type Map struct {
	Width  int      // Map width
	Height int      // Map height
	Ratio  byte     // When filling in the map with random bytes, if byte > ratio then wall is set otherwise is not
	Seed   int64    // Seed for the random function
	Smooth int      // How many times run the smoothing function
	Array  [][]bool // True - wall is set, False - no wall
}

func imap(vs []byte, f func(byte) bool) []bool {
	vsm := make([]bool, len(vs))

	for i, v := range vs {
		vsm[i] = f(v)
	}

	return vsm
}

// Fills in the array with random values
func (m *Map) generateArray() {
	array := make([][]bool, m.Height)

	for y := 0; y < m.Height; y++ {
		bytes := make([]byte, m.Width)
		rand.Read(bytes)

		array[y] = imap(bytes, func(b byte) bool {
			return (b > m.Ratio)
		})
	}

	m.Array = array
}

// Smoothes the map
func (m *Map) smooth() {
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			if m.surroundingCount(x, y) > 4 {
				m.Array[y][x] = true
			} else if m.surroundingCount(x, y) < 4 {
				m.Array[y][x] = false
			}
		}
	}
}

// Counts number of walls around (x, y)
func (m *Map) surroundingCount(x, y int) int {
	count := 0

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}

			if x+dx < 0 || x+dx >= m.Width || y+dy < 0 || y+dy >= m.Height {
				count++
			}

			if m.Get(x+dx, y+dy) {
				count++
			}
		}
	}

	return count
}

// Generates the map into inner array
func (m *Map) Generate() {
	rand.Seed(m.Seed)
	m.generateArray()
	for i := 0; i < m.Smooth; i++ {
		m.smooth()
	}
}

// Returns map value at (x, y). If overboard, true returned
func (m *Map) Get(x, y int) bool {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		return true
	}

	return m.Array[y][x]
}

// Returns new Map instance
func NewMap(width, height int) Map {
	return Map{
		Width:  width,
		Height: height,
		Ratio:  127,
		Seed:   rand.Int63(),
		Smooth: 5,
	}
}
