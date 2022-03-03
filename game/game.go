package game

import (
	"fmt"
)

func GamePlay() {
	/* this is game of console*/
	fmt.Println("hello player")
	// Init game parameters and print field.
	var isLoop bool = true
	field := stickInit()
	PrintField(field)

	for isLoop {
		sel_stick := selectStick()             //入力1
		sel_direction := selectMoveDirection() //入力2
		isMove := moveDisk(sel_stick, sel_direction, field)
		if isMove {
			fmt.Println("you can move disk.")
		} else {
			fmt.Println("you cannot move this!")
		}
		PrintField(field)
		if CheckEnd(field) {
			fmt.Println("Complete!")
			isLoop = false
		}
	}
}

func stickInit() [][]int {
	fields := make([][]int, 3)
	for i := 0; i < 3; i++ {
		fields[i] = make([]int, 3)
		fields[i][0] = i + 1
	}
	return fields
}

func selectStick() int {
	var input int
	fmt.Println("Which disk do you want to move?")
	fmt.Scan(&input)
	return input
}

func selectMoveDirection() int {
	var input int
	fmt.Println("Where do you want to move?")
	fmt.Scan(&input)
	return input
}

func moveDisk(_s, _x int, _f [][]int) bool {
	var selDisk int
	var prePos [2]int

	// get disk
	for i := 0; i < 3; i++ {
		if _f[i][_s-1] != 0 {
			prePos[1] = i
			prePos[0] = _s - 1
			selDisk = _f[i][_s-1]
			_f[i][_s-1] = 0
			break
		}
	}
	if selDisk == 0 {
		return false
	}

	// move disk
	for i := 0; i < 3; i++ {
		if _f[2-i][_x-1] == 0 {
			if i == 0 { // 一番下なら無条件におk
				_f[2-i][_x-1] = selDisk
				return true
			} else { // 1、2番目に上の方
				// 自分より下がサイズが大きい時
				if selDisk < _f[2-i+1][_x-1] {
					_f[2-i][_x-1] = selDisk
					return true
				}
			}
		}
	}
	_f[prePos[1]][prePos[0]] = selDisk // 移動できなかった場合、元の位置にディスクを戻す
	return false
}

func PrintField(fields [][]int) {
	// [line][column]
	for i := 0; i < 3; i++ {
		fmt.Printf("%d,%d,%d\n", fields[i][0], fields[i][1], fields[i][2])
	}
}

func CheckEnd(fields [][]int) bool {
	return fields[0][2] == 1
}
