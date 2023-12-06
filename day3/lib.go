package day3

import "unicode"

type Part struct {
	number string
	pos    [2]int
}

type Gear struct {
	x, y int
}

func isSpecialChar(char rune) bool {
	return char != '.' && !unicode.IsDigit(char)
}

func (p *Part) isValid(mtx [][]rune) bool {
	for i, _ := range p.number {
		x, y := p.pos[0], p.pos[1]+i

		onTopRow := x == 0
		onBottomRow := x == len(mtx)-1
		onFirstCol := y == 0
		onLastCol := y == len(mtx[x])-1
		onLeftmostDigit := i == 0
		onRightmostDigit := i == len(p.number)-1

		// check NORD
		if onLeftmostDigit && !onTopRow && isSpecialChar(mtx[x-1][y]) {
			return true
		}
		// check SOUTH
		if onLeftmostDigit && !onBottomRow && isSpecialChar(mtx[x+1][y]) {
			return true
		}
		// ckeck WEST
		if onLeftmostDigit && !onFirstCol && isSpecialChar(mtx[x][y-1]) {
			return true
		}
		// check EAST
		if onRightmostDigit && !onLastCol && isSpecialChar(mtx[x][y+1]) {
			return true
		}
		// check NW
		if onLeftmostDigit && !onFirstCol && !onTopRow && isSpecialChar(mtx[x-1][y-1]) {
			return true
		}
		// check NE
		if !onLastCol && !onTopRow && isSpecialChar(mtx[x-1][y+1]) {
			return true
		}
		// check SW
		if onLeftmostDigit && !onFirstCol && !onBottomRow && isSpecialChar(mtx[x+1][y-1]) {
			return true
		}
		// check SE
		if !onLastCol && !onBottomRow && isSpecialChar(mtx[x+1][y+1]) {
			return true
		}
	}

	return false
}
