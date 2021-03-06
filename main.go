// BETA VERSION //

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	const ROW, COLUMN = 15, 15
	var count uint = 1
	var turn string
	var piece string
	var finished bool = false

	grid := map[int]map[int]string{} // 2d map grid for the main game

	for row := 0; row <= ROW; row++ {
		grid[row] = map[int]string{}
		for column := 0; column <= COLUMN; column++ {
			if row == 0 {
				if column < 10 {
					grid[row][column] = strconv.Itoa(column) + " "
				} else {
					grid[row][column] = strconv.Itoa(column)
				}
			} else if column == 0 {
				if row < 10 {
					grid[row][column] = strconv.Itoa(row) + " "
				} else {
					grid[row][column] = strconv.Itoa(row)
				}
			} else {
				grid[row][column] = " "
			}
		}
	}

	for !finished {

		for row := 0; row <= ROW; row++ { // printing the grid in slice form for better screen view
			var slice []string
			for column := 0; column <= COLUMN; column++ {
				slice = append(slice, grid[row][column])
			}
			fmt.Println(slice)
		}

		if numType(count) {
			turn = "Black"
			piece = "B "
		} else {
			turn = "White"
			piece = "W "
		}

		fmt.Printf("It's %v's turn\n", turn)
		a, b := input(ROW, COLUMN)

		grid[a][b] = piece

		for row := 0; row <= ROW; row++ { // win conditions
			for column := 0; column <= COLUMN; column++ {
				for b+4 <= COLUMN {
					if grid[a][b] == grid[a][b+1] && grid[a][b] == grid[a][b+2] && grid[a][b] == grid[a][b+3] && grid[a][b] == grid[a][b+4] && grid[a][b] != "  " {
						fmt.Printf("%v wins!", turn)
						finished = true
					}
				}
				for a+4 <= ROW {
					if grid[a][b] == grid[a+1][b] && grid[a][b] == grid[a+2][b] && grid[a][b] == grid[a+3][b] && grid[a][b] == grid[a+4][b] && grid[a][b] != "  " {
						fmt.Printf("%v wins!", turn)
						finished = true
					}
				}
				for b+4 <= COLUMN && a+4 <= ROW {
					if grid[a][b] == grid[a+1][b+1] && grid[a][b] == grid[a+2][b+2] && grid[a][b] == grid[a+3][b+3] && grid[a][b] == grid[a+4][b+4] && grid[a][b] != "  " {
						fmt.Printf("%v wins!", turn)
						finished = true
					}
				}
			}
		}
		count++
	}
	end()
}

func numType(a uint) bool { // false if even, true if odd
	if a%2 == 0 {
		return false
	}
	return true
}

func input(i, j int) (int, int) { // row and column input control
	x := bufio.NewScanner(os.Stdin)
	y := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter row number, press ENTER, then enter column number: \n")
	x.Scan()
	y.Scan()
	inputx, errx := strconv.ParseInt(x.Text(), 10, 64)
	inputy, erry := strconv.ParseInt(y.Text(), 10, 64)

	intInputx := int(inputx)
	intInputy := int(inputy)

	for errx != nil || erry != nil || intInputx <= 0 || intInputx > i || intInputy <= 0 || intInputy > j {
		x = bufio.NewScanner(os.Stdin)
		y = bufio.NewScanner(os.Stdin)
		fmt.Printf("Invalid coordinates for row or column. Re-enter inputs row and column: \n")
		x.Scan()
		y.Scan()
		inputx, errx = strconv.ParseInt(x.Text(), 10, 64)
		inputy, erry = strconv.ParseInt(y.Text(), 10, 64)

		intInputx = int(inputx)
		intInputy = int(inputy)
	}

	return intInputx, intInputy
}

func end() { // game end conditions
	end := bufio.NewScanner(os.Stdin)
	fmt.Printf("Press ENTER to end")
	end.Scan()
}
