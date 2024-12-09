package day6

type Guard struct {
	x            int
	y            int
	direction    string
	traveledPath []visitedCell
}

func newGuard(x int, y int) *Guard {
	guard := new(Guard)
	guard.x = x
	guard.y = y
	guard.direction = "up"
	return guard
}

func (g *Guard) set(x int, y int, direction string) {
	g.x = x
	g.y = y
	g.direction = direction
}

func (g Guard) getNext() (int, int) {
	switch g.direction {
	case "up":
		return g.x, g.y - 1
	case "left":
		return g.x - 1, g.y
	case "right":
		return g.x + 1, g.y
	case "down":
		return g.x, g.y + 1
	}
	return g.x, g.y
}

func (g *Guard) moveNext() {
	nextX, nextY := g.getNext()
	g.x = nextX
	g.y = nextY
}

func (g *Guard) changeDirection() {
	g.direction = g.getNextDirection()
}

func (g *Guard) getNextDirection() string {
	switch g.direction {
	case "up":
		return "right"
	case "right":
		return "down"
	case "down":
		return "left"
	case "left":
		return "up"
	}
	return "up"
}

func (g *Guard) addToPath(x int, y int, direction string) {
	for _, value := range g.traveledPath {
		if value.x == x && value.y == y {
			if !value.checkDirections(direction) {
				value.directions = append(value.directions, direction)
				return
			}
		}
	}
	g.traveledPath = append(g.traveledPath, *newPosAndDirection(x, y, direction))
}

type visitedCell struct {
	x          int
	y          int
	directions []string
}

func newPosAndDirection(x int, y int, direction string) *visitedCell {
	result := new(visitedCell)
	result.x = x
	result.y = y
	result.directions = []string{direction}
	return result
}

func (c *visitedCell) checkDirections(direction string) bool {
	for _, value := range c.directions {
		if value == direction {
			return true
		}
	}
	return false
}
