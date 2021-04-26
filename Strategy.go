package main

import (
	"fmt"
)

//往后预判几手
var deepMind = 2

//测试智力
func Strategy() {


	for i := 0; i < BOARDSIZE; i++ {
		for j := 0; j < BOARDSIZE; j++ {
			rooms = append(rooms, Coordinate{i,j})
		}
	}
	fmt.Println("Hello World!")
	fail := 0
	for fail < 50 {
		maxValue := 0
		bestChoice := 0
		var sugests = sugestFur(nowColor,HAND, rooms, status)
		for index, sugest := range sugests {
			value := tryPut(sugest.x, sugest.y, nowColor, HAND, status, rooms, dajie)
			if value > maxValue {
				maxValue = value
				bestChoice = index
			}
		}
		x := sugests[bestChoice].x
		y := sugests[bestChoice].y
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
func tryPut(x, y, myColor, pre_Hand int, statusV [9][9]int, roomsV []Coordinate, dajieV Coordinate) int {
	myPreHand := pre_Hand

	var myState [9][9]int
	for index, son := range statusV {
		for i, value := range son {
			myState[index][i] = value
		}
	}

	var myRooms []Coordinate
	myRooms = roomsV

	myDajie := Coordinate{dajieV.x, dajieV.y}

	var linkV []Coordinate
	var eatenV []Coordinate
	if putChessV(x, y, myColor, myState, myRooms, myDajie, linkV, eatenV) {
		//雙方各往下預判几手
		if myPreHand == HAND + deepMind*2 - 1{
			return calculate(myColor, myRooms, myState)
		}
		myPreHand++
		for i := 0; i < len(myRooms); i++ {
			if myRooms[i].x == x && myRooms[i].y == y {
				myRooms = append(myRooms[:i], myRooms[i+1:]...)
			}
		}
		//updateRoomsV(x, y, myRooms)

		var mySugests = sugestFur(myColor,myPreHand ,myRooms, myState)
		for _, sugest := range mySugests {
			tryPut(sugest.x, sugest.y, - myColor, myPreHand, myState, myRooms, myDajie)
		}
	}

	return 0
}

//落子
func putChessV(x, y, color int, statusV [9][9]int, roomsV []Coordinate, dajieV Coordinate, linkV, eatenV []Coordinate) bool {
	if statusV[x][y] != 0 {
		return false
	}
	//模拟落子
	statusV[x][y] = color
	var eatLeft = false
	var eatRight = false
	var eatTop = false
	var eatBotton = false
	if x+1 < BOARDSIZE {
		eatRight = killEnemyV(x+1, y, color, statusV, dajieV, roomsV, linkV, eatenV)
	}
	if x-1 >= 0 {
		eatLeft = killEnemyV(x-1, y, color, statusV, dajieV, roomsV, linkV, eatenV)
	}
	if y+1 < BOARDSIZE {
		eatTop = killEnemyV(x, y+1, color, statusV, dajieV, roomsV, linkV, eatenV)
	}
	if y-1 >= 0 {
		eatBotton = killEnemyV(x, y-1, color, statusV, dajieV, roomsV, linkV, eatenV)
	}

	if eatLeft || eatRight || eatTop || eatBotton {
		return true
	}

	//判断自己有没有气
	linkV = linkV[0:0]
	if findWayV(x, y, color,statusV,eatenV) {
		dajieV = Coordinate{-1, -1}
		return true
	}
	statusV[x][y] = 0
	return false
}

//更新 空点信息
func updateRoomsV(x, y int, roomsV []Coordinate) {
	for i := 0; i < len(roomsV); i++ {
		if roomsV[i].x == x && roomsV[i].y == y {
			roomsV = append(roomsV[:i], roomsV[i+1:]...)
		}
	}
}

//吃掉对方
func killEnemyV(x, y, color int, statusV [9][9]int, dajieV Coordinate, roomsV, linkV, eatenV []Coordinate) bool {
	//先判断能否吃对方
	linkV = linkV[0:0]
	eatenV = eatenV[0:0]
	if statusV[x][y] == -color && !findWay(x, y, -color) {
		if len(eatenV) == 1 && dajieV == eatenV[0] {
			return false
		}
		if len(eatenV) == 1 && dajieV != eatenV[0] {
			dajieV = Coordinate{x, y}
			statusV[x][y] = 0
		}
		if len(eatenV) > 1 {
			dajieV = Coordinate{-1, -1}
		}
		//吃棋子，将对方所占棋盘清空
		for i := 0; i < len(eatenV); i++ {
			statusV[eatenV[i].x][eatenV[i].y] = 0
			roomsV = append(roomsV, eatenV[i])
		}
		return true
	}
	return false
}


//寻气
func findWayV(x, y, color int,statusV [9][9]int, eatenV []Coordinate) bool {
	var northAlive = false
	var southAlive = false
	var westAlive = false
	var eastAlive = false
	link = append(link, Coordinate{x, y})

	if x+1 < BOARDSIZE && statusV[x+1][y] == 0 {

		return true
	}
	if x-1 >= 0 && statusV[x-1][y] == 0 {

		return true
	}
	if y+1 < BOARDSIZE && statusV[x][y+1] == 0 {

		return true
	}
	if y-1 >= 0 && statusV[x][y-1] == 0 {

		return true
	}
	if !linkContains(link, Coordinate{x + 1, y}) && (x+1) < BOARDSIZE && statusV[x+1][y] == color {
		westAlive = findWayV(x+1, y, color,statusV,eatenV)
	}
	if !linkContains(link, Coordinate{x - 1, y}) && x-1 >= 0 && statusV[x-1][y] == color {
		eastAlive = findWayV(x-1, y, color,statusV,eatenV)
	}
	if !linkContains(link, Coordinate{x, y + 1}) && y+1 < BOARDSIZE && statusV[x][y+1] == color {
		northAlive = findWayV(x, y+1, color,statusV,eatenV)
	}
	if !linkContains(link, Coordinate{x, y - 1}) && y-1 >= 0 && statusV[x][y-1] == color {
		southAlive = findWayV(x, y-1, color,statusV,eatenV)
	}
	if !(westAlive || eastAlive || northAlive || southAlive) {
		eatenV = append(eatenV,Coordinate{x, y})
	}

	return westAlive || eastAlive || northAlive || southAlive
}

