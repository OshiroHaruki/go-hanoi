package game

import (
	"fmt"
	"go-hanoi/disk"
)

var d1 disk.Disk
var d2 disk.Disk
var d3 disk.Disk

var isLoop bool = true

func GamePlay() {
	fmt.Println("hello gamer")
	diskInit()
	fields := stickInit()
	printField(fields)
	for isLoop {
		sel_stick := selectStick()
		sel_direction := selectMoveDirection()
		isMove := moveDisk(sel_stick, sel_direction, fields)
		if isMove {
			fmt.Println("you can move disk.")
		} else {
			fmt.Println("you cannot move this!")
		}
		printField(fields)
		if checkEnd(fields) {
			fmt.Println("Complete!")
			isLoop = false
		}
	}
}

func stickInit() [][]int {
	fields := make([][]int, 3) //makeで作ってるからスライス＝他の関数で値を弄れる
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
	for i := 0; i < 3; i++ {
		if _f[i][_s-1] != 0 {
			prePos[1] = i
			prePos[0] = _s - 1
			selDisk = _f[i][_s-1]
			_f[i][_s-1] = 0
			break
		}
	}
	//if selDisk == 0{
	// 移動できるディスクが選択した棒にない。
	//}
	for i := 0; i < 3; i++ {
		//fmt.Println(_f[2-i][_x-1])
		//fmt.Println(selDisk)
		if _f[2-i][_x-1] == 0 {
			if i == 0 { // 一番下なら無条件におk
				_f[2-i][_x-1] = selDisk
				return true
			} else { // 1、2番目に上の方
				// 自分より下がサイズが大きい時
				if selDisk < _f[2-i+1][_x-1] {
					_f[2-i][_x-1] = selDisk
					return true
				} else {
					_f[prePos[1]][prePos[0]] = selDisk // 元の位置にディスクを戻す
					return false
				}
			}
		}
	}
	return false
}

func printField(fields [][]int) {
	// [行][列]
	for i := 0; i < 3; i++ {
		fmt.Printf("%d,%d,%d\n", fields[i][0], fields[i][1], fields[i][2])
	}
}

func checkEnd(fields [][]int) bool {
	if fields[0][2] == 1 {
		return true
		/*以下イランかも
		if fields[][]
		*/
	}
	return false
}

func diskInit() {
	d1 = disk.Disk{1, 1, 1}
	d2 = disk.Disk{2, 1, 2}
	d3 = disk.Disk{3, 1, 3}
	/*
				x
		y		(1,1)	(2,1)	(3,1)
				(1,2)	(2,2)	(3,2)
				(3,3)	(2,3)	(3,3)
	*/
}

func printDisks() {
	fmt.Printf("disk 1 is in (%d,%d)\n", d1.Position_x, d1.Position_y)
	fmt.Printf("disk 2 is in (%d,%d)\n", d2.Position_x, d2.Position_y)
	fmt.Printf("disk 3 is in (%d,%d)\n", d3.Position_x, d3.Position_y)
}

func selectDisk() int {
	var input int
	fmt.Println("Which disk do you want to move?")
	fmt.Scan(&input)
	return input
}

/*
func selectMoveDirection() int {
	var input int
	fmt.Println("Where do you want to move?")
	fmt.Scan(&input)
	return input
}
*/
/*
func diskMove(_d, _x int) {
	switch _d {
	case 1:
		d1.Move(_x, 1)
	case 2:
		d2.Move(_x, 1)
	case 3:
		d3.Move(_x, 1)
	}
}*/
/*
func moveCheck(_x int) bool {
	if _x < 0 || 3 < _x {
		return false
	}
	return true
}
*/
