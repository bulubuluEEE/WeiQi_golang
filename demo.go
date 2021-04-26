package main

import (
	"fmt"
	"math/rand"
)

func demo(){

	for i := 0; i < BOARDSIZE; i++ {
		for j := 0; j < BOARDSIZE; j++ {
			rooms = append(rooms, Coordinate{i,j})
		}
	}
	fmt.Println("Hello World!")
	fail := 0
	for fail < 50 {
		x := rand.Intn(BOARDSIZE)
		y := rand.Intn(BOARDSIZE)
		if putChess(x, y, nowColor) {
			updateRooms(x,y)
			var name = "黑"
			if nowColor == WHITE {
				name = "白"
			}
			fmt.Printf("======第%d手,%s棋落子：坐标x%d,y%d=======\n", HAND, name, x, y)
			drawBoard(x, y, nowColor)
			fail = 0
			HAND ++
			nowColor = -nowColor
		} else {

			fail++
		}

	}

}

var link []Coordinate
var eaten []Coordinate
//落子
func putChess(x, y, color int) bool {
	if status[x][y] != 0 {
		return false
	}
	//模拟落子
	status[x][y] = color
	var eatLeft = false
	var eatRight = false
	var eatTop = false
	var eatBotton = false
	if x+1 < BOARDSIZE {
		eatRight = killEnemy(x+1, y, color)
	}
	if x-1 >= 0 {
		eatLeft = killEnemy(x-1, y, color)
	}
	if y+1 <BOARDSIZE {
		eatTop = killEnemy(x, y+1, color)
	}
	if y-1 >= 0 {
		eatBotton = killEnemy(x, y-1, color)
	}

	if eatLeft || eatRight || eatTop || eatBotton {
		return true
	}

	//判断自己有没有气
	link = link[0:0]
	if findWay(x, y, color) {
		dajie = Coordinate{-1, -1}
		return true
	}
	status[x][y] = 0
	return false
}
//更新 空点信息
func updateRooms(x,y int)  {
	for i := 0; i < len(rooms); i++ {
		if rooms[i].x == x && rooms[i].y == y {
			rooms = append(rooms[:i], rooms[i+1:]...)
		}
	}
}
//吃掉对方
func killEnemy(x, y, color int) bool{
	//先判断能否吃对方
	link = link[0:0]
	eaten = eaten[0:0]
	if status[x][y] == -color && !findWay(x, y, -color) {
		var name = "黑"
		if -color == WHITE {
			name = "白"
		}
		if len(eaten) == 1 && dajie == eaten[0] {
			fmt.Printf("打劫，坐标：x%d,y%d\n", x, y)
			return false
		}
		if len(eaten) == 1 && dajie != eaten[0] {
			dajie = Coordinate{x, y}
			status[x][y] = 0
		}
		if len(eaten) > 1 {
			dajie = Coordinate{-1, -1}
		}
		fmt.Printf("%s方被吃，坐标：", name)
		//吃棋子，将对方所占棋盘清空
		for i := 0; i < len(eaten); i++ {
			status[eaten[i].x][eaten[i].y] = 0
			rooms = append(rooms, eaten[i])
			fmt.Printf("(%d,%d)——", eaten[i].x, eaten[i].y)
		}
		fmt.Println()
		return true
	}
	return false
}

//寻气
func findWay(x, y, color int) bool {
	var northAlive = false
	var southAlive = false
	var westAlive = false
	var eastAlive = false
	link = append(link, Coordinate{x, y})

	if x+1 < BOARDSIZE && status[x+1][y] == 0 {

		return true
	}
	if x-1 >= 0 && status[x-1][y] == 0 {

		return true
	}
	if y+1 < BOARDSIZE && status[x][y+1] == 0 {

		return true
	}
	if y-1 >= 0 && status[x][y-1] == 0 {

		return true
	}
	if !linkContains(link, Coordinate{x + 1, y}) && (x+1) < BOARDSIZE && status[x+1][y] == color {
		westAlive = findWay(x+1, y, color)
	}
	if !linkContains(link, Coordinate{x - 1, y}) && x-1 >= 0 && status[x-1][y] == color {
		eastAlive = findWay(x-1, y, color)
	}
	if !linkContains(link, Coordinate{x, y + 1}) && y+1 < BOARDSIZE && status[x][y+1] == color {
		northAlive = findWay(x, y+1, color)
	}
	if !linkContains(link, Coordinate{x, y - 1}) && y-1 >= 0 && status[x][y-1] == color {
		southAlive = findWay(x, y-1, color)
	}
	if !(westAlive || eastAlive || northAlive || southAlive) {
		eaten = append(eaten,Coordinate{x, y})
	}

	return westAlive || eastAlive || northAlive || southAlive
}

func removeLink(link,free []Coordinate)  {

}