package game2048

import (
	"fmt"
	"math/rand"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

//Game2048 the main struct for the library
type Game2048 struct {
	Matrix [][]int
	Size   int
	Moves  int
	Score  int
}

/*
NewGame2048 is the constructor of the Game2048,
returns a pointer to Game2048 instance
*/
func NewGame2048(size int) *Game2048 {
	m := make([][]int, size)
	for i := range m {
		m[i] = make([]int, size)
	}
	for i := range m {
		for j := range m[i] {
			m[i][j] = 0
		}
	}
	g := Game2048{
		m,
		size,
		0,
		0,
	}
	return &g
}

//AddRandom adds a 2 or 4 on a space in the matrix randomly
func (g *Game2048) AddRandom() {
	var positions []int
	size := 0
	pos := 0

	for i := range g.Matrix {
		for _, v := range g.Matrix[i] {
			if v == 0 {
				positions = append(positions, pos)
				size++
			}
			pos++
		}
	}
	pos = positions[rnd.Intn(size)]
	n := (rnd.Intn(2) + 1) * 2
	g.Matrix[pos/g.Size][pos%g.Size] = n
}

//Print prints well formated data of the game
func (g *Game2048) Print() {
	fmt.Printf("Moves: %d\nScore: %d\n", g.Moves, g.Score)
	for i := range g.Matrix {
		for _, v := range g.Matrix[i] {
			fmt.Printf("%d\t", v)
		}
		fmt.Printf("\n")
	}
}

//MoveUP moves up all possible cells
func (g *Game2048) MoveUP() {
	fmt.Printf("MOVE UP\n")
	v := make([]int, g.Size)
	for i := 0; i < g.Size; i++ {
		for j := 0; j < g.Size; j++ {
			v[j] = g.Matrix[j][i]
		}
		g.Score += group(v)
		for j := 0; j < g.Size; j++ {
			g.Matrix[j][i] = v[j]
		}
	}
	g.Moves++
}

//MoveDOWN moves down all possible cells
func (g *Game2048) MoveDOWN() {
	fmt.Printf("MOVE DOWN\n")
	v := make([]int, g.Size)
	for i := 0; i < g.Size; i++ {
		for j := 0; j < g.Size; j++ {
			v[j] = g.Matrix[(g.Size-1)-j][i]
		}
		g.Score += group(v)
		for j := 0; j < g.Size; j++ {
			g.Matrix[(g.Size-1)-j][i] = v[j]
		}
	}
	g.Moves++
}

//MoveLEFT moves left all possible cells
func (g *Game2048) MoveLEFT() {
	fmt.Printf("MOVE RIGHT\n")
	v := make([]int, g.Size)
	for i := 0; i < g.Size; i++ {
		for j := 0; j < g.Size; j++ {
			v[j] = g.Matrix[i][j]
		}
		g.Score += group(v)
		for j := 0; j < g.Size; j++ {
			g.Matrix[i][j] = v[j]
		}
	}
	g.Moves++
}

//MoveRIGHT moves right all possible cells
func (g *Game2048) MoveRIGHT() {
	fmt.Printf("MOVE LEFT\n")
	v := make([]int, g.Size)
	for i := 0; i < g.Size; i++ {
		for j := 0; j < g.Size; j++ {
			v[j] = g.Matrix[i][(g.Size-1)-j]
		}
		g.Score += group(v)
		for j := 0; j < g.Size; j++ {
			g.Matrix[i][(g.Size-1)-j] = v[j]
		}
	}
	g.Moves++
}

//ValidateUP checks if can move up
func (g *Game2048) ValidateUP() bool {
	for i := 0; i < g.Size; i++ {
		zeroFound := false
		for j := 0; j < g.Size; j++ {
			if g.Matrix[j][i] == 0 {
				zeroFound = true
			}
			if g.Matrix[j][i] != 0 {
				if zeroFound {
					return true
				}
				if j < (g.Size-1) && g.Matrix[j][i] == g.Matrix[j+1][i] {
					return true
				}
			}
		}
	}
	return false
}

//ValidateDOWN checks if can move down
func (g *Game2048) ValidateDOWN() bool {
	for i := 0; i < g.Size; i++ {
		zeroFound := false
		for j := 0; j < g.Size; j++ {
			if g.Matrix[(g.Size-1)-j][i] == 0 {
				zeroFound = true
			}
			if g.Matrix[(g.Size-1)-j][i] != 0 {
				if zeroFound {
					return true
				}
				if j > 0 && g.Matrix[j][i] == g.Matrix[j-1][i] {
					return true
				}
			}
		}
	}
	return false
}

//ValidateLEFT checks if can move left
func (g *Game2048) ValidateLEFT() bool {
	for i := 0; i < g.Size; i++ {
		zeroFound := false
		for j := 0; j < g.Size; j++ {
			if g.Matrix[i][j] == 0 {
				zeroFound = true
			}
			if g.Matrix[i][j] != 0 {
				if zeroFound {
					return true
				}
				if j < (g.Size-1) && g.Matrix[i][j] == g.Matrix[i][j+1] {
					return true
				}
			}
		}
	}
	return false
}

//ValidateRIGHT checks if can move right
func (g *Game2048) ValidateRIGHT() bool {
	for i := 0; i < g.Size; i++ {
		zeroFound := false
		for j := 0; j < g.Size; j++ {
			if g.Matrix[i][(g.Size-1)-j] == 0 {
				zeroFound = true
			}
			if g.Matrix[i][(g.Size-1)-j] != 0 {
				if zeroFound {
					return true
				}
				if j > 0 && g.Matrix[i][j] == g.Matrix[i][j-1] {
					return true
				}
			}
		}
	}
	return false
}

func group(v []int) int {
	score := 0
	vAux := []int{0, 0, 0, 0}

	pos := 0
	for i := range v {
		if v[i] != 0 {
			vAux[pos] = v[i]
			pos++
		}
		v[i] = 0
	}
	pos = 0
	for i := 0; i < len(vAux); i++ {
		if i < 3 && vAux[i] == vAux[i+1] {
			v[pos] = vAux[i] * 2
			score += v[pos]
			i++
		} else {
			v[pos] = vAux[i]
		}
		pos++
	}
	return score
}
