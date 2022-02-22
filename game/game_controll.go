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
	//printDisks()
	//for isLoop {
	//sel_disk := selectDisk()
	//sel_direction := selectMoveDirection()
	//if moveCheck(sel_direction) {
	//	diskMove(sel_disk, sel_direction)
	//}
	//printDisks()
	//for i := 0; i < 3; i++ {
	//	fmt.Printf("%d,%d,%d", fields[i][0], fields[i][1], fields[i][2])
	//}
	//}
	for i := 0; i < 3; i++ {
		fmt.Printf("%d,%d,%d\n", fields[i][0], fields[i][1], fields[i][2])
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
func diskMove(_d, _x int) {
	switch _d {
	case 1:
		d1.Move(_x, 1)
	case 2:
		d2.Move(_x, 1)
	case 3:
		d3.Move(_x, 1)
	}
}

func moveCheck(_x int) bool {
	if _x < 0 || 3 < _x {
		return false
	}
	return true
}
