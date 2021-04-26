package main

//计算双方的周长差（边界数之差）
func calculate(color int, roomsV []Coordinate, statusV [9][9]int) int {
	//周长总和
	myGirth := 0
	enemyGirth := 0
	for _, room := range roomsV {
		if (room.x+1 < BOARDSIZE && statusV[room.x+1][room.y] == color) ||
			(room.x-1 >= 0 && statusV[room.x-1][room.y] == color) ||
			(room.y+1 < BOARDSIZE && statusV[room.x][room.y+1] == color) ||
			(room.y-1 >= 0 && statusV[room.x][room.y-1] == color) {

			myGirth++
		}
		if (room.x+1 < BOARDSIZE && statusV[room.x+1][room.y] == -color) ||
			(room.x-1 >= 0 && statusV[room.x-1][room.y] == -color) ||
			(room.y+1 < BOARDSIZE && statusV[room.x][room.y+1] == -color) ||
			(room.y-1 >= 0 && statusV[room.x][room.y-1] == -color) {

			enemyGirth++
		}
	}
	return myGirth - enemyGirth
}

//敌人的所有边界
func sugestFur(color, handV int, roomsV []Coordinate, statusV [9][9]int) []Coordinate {
	var results []Coordinate

	for _, room := range roomsV {
		x := room.x
		y := room.y
		if  handV == 1 || isNear(x, y, color, statusV) {
			if handV > 4 || (x > 1 && x < 7 && y > 1 && y < 7) {
				results = append(results, room)
			}
		}
	}
	return results
}

//判断是否临近
func isNear(x, y, color int, statusV [9][9]int) bool {
	//距离1
	distance1 := (x+1 < BOARDSIZE && statusV[x+1][y] == -color) ||
		(x-1 >= 0 && statusV[x-1][y] == -color) ||
		(y+1 < BOARDSIZE && statusV[x][y+1] == -color) ||
		(y-1 >= 0 && statusV[x][y-1] == -color)
	//距离2
	distance2 := (x+2 < BOARDSIZE && statusV[x+2][y] == -color) ||
		(x-2 >= 0 && statusV[x-2][y] == -color) ||
		(y+2 < BOARDSIZE && statusV[x][y+2] == -color) ||
		(y-2 >= 0 && statusV[x][y-2] == -color)
	//距离3
	//distance3 := (x+3 < BOARDSIZE && statusV[x+3][y] == -color) ||
	//	(x-3 >= 0 && statusV[x-3][y] == -color) ||
	//	(y+3 < BOARDSIZE && statusV[x][y+3] == -color) ||
	//	(y-3 >= 0 && statusV[x][y-3] == -color)

	//斜着距离1
	near1 := (x+1 < BOARDSIZE && y+1 < BOARDSIZE && statusV[x+1][y+1] == -color) ||
		(x-1 >= 0 && y+1 < BOARDSIZE && statusV[x-1][y+1] == -color) ||
		(x+1 < BOARDSIZE && y-1 >= 0 && statusV[x+1][y-1] == -color) ||
		(x-1 >= 0 && y-1 >= 0 && statusV[x-1][y-1] == -color)

	return distance1 || distance2 || near1
}
