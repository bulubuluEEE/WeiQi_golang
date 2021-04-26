package main


const BOARDSIZE = 9
const BLACK = 1
const WHITE = -1
const EMPTY = 0

var nowColor = BLACK

//网点状态
var status [BOARDSIZE][BOARDSIZE]int
//空点 未落子的点
var rooms []Coordinate

//棋子
var pieces map[int]int

type Coordinate struct {
	x int
	y int
}

var dajie = Coordinate{-1, -1}

var HAND = 1
