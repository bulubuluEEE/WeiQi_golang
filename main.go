package main

import "fmt"

func main() {
	Strategy()
}


//绘制棋盘
func drawBoard(x, y, color int) {

	for index, son := range status {
		for i, value := range son {
			if index == x && i == y {
				if color == BLACK {
					fmt.Print("◕ ")
				}
				if color == WHITE {
					fmt.Print("◔ ")
				}

			} else {
				if value == EMPTY {
					fmt.Print("✛ ")
				}
				if value == BLACK {
					fmt.Print("⚫ ")
				}
				if value == WHITE {
					fmt.Print("⚪ ")
				}
			}
		}
		fmt.Println()
	}
}

func linkContains(link []Coordinate, c Coordinate) bool {
	for i := 0; i < len(link); i++ {
		if link[i].x == c.x && link[i].y == c.y {
			return true
		}
	}
	return false
}