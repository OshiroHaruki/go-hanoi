package gamegui

import (
	"fmt"
	"go-hanoi/game"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GamePlay_view() {
	field := stickInit()

	myApp := app.New()
	w := myApp.NewWindow("tower of hanoi")
	w.Resize(fyne.NewSize(600, 500))

	backGround := canvas.NewImageFromFile("./game/Stick.png")
	backGround.Resize(fyne.NewSize(600, 400))
	backGround.Move(fyne.NewPos(0, 0))

	disk1 := canvas.NewImageFromFile("./game/Red.png")
	disk1.Resize(fyne.NewSize(80, 50)) //yは50区切りでしたから320,280,230
	disk1.Move(fyne.NewPos(85, 230))   //xは85,260,435

	disk2 := canvas.NewImageFromFile("./game/Blue.png")
	disk2.Resize(fyne.NewSize(100, 50)) //yは同じ
	disk2.Move(fyne.NewPos(72, 280))    //xは,72,247,425

	disk3 := canvas.NewImageFromFile("./game/Green.png")
	disk3.Resize(fyne.NewSize(120, 50)) //yは同じ
	disk3.Move(fyne.NewPos(68, 320))    //xは68,240,415

	disks := []*canvas.Image{disk1, disk2, disk3}

	button1 := widget.NewButton("left", func() {
		log.Println("1")
		gameOneStep(1, field, disks)
	})
	button1.Resize(fyne.NewSize(100, 50))
	button1.Move(fyne.NewPos(0, 400))

	button2 := widget.NewButton("center", func() {
		log.Println("2")
		gameOneStep(2, field, disks)
	})
	button2.Resize(fyne.NewSize(100, 50))
	button2.Move(fyne.NewPos(100, 400))

	button3 := widget.NewButton("right", func() {
		log.Println("3")
		gameOneStep(3, field, disks)
	})
	button3.Resize(fyne.NewSize(100, 50))
	button3.Move(fyne.NewPos(200, 400))

	w.SetContent(
		container.NewWithoutLayout(
			backGround,
			button1,
			button2,
			button3,
			disk1,
			disk2,
			disk3,
		),
	)

	w.ShowAndRun()

}

var D1Pos = []float32{85, 260, 435} //cannot define const on array(golang)
var D2Pos = []float32{72, 247, 425}
var D3Pos = []float32{68, 240, 415}
var XPos = [][]float32{D1Pos, D2Pos, D3Pos}
var YPos = []float32{30, 230, 280, 320}
var step int = 0
var sel_stick int = 0
var sel_disk int = 0
var sel_move int = 0
var isEnd bool = false
var animPrePos [2]int
var animPrePosX float32
var putPosY int
var prePos [2]int

func gameOneStep(_s int, _f [][]int, disks []*canvas.Image) {
	if !isEnd {
		if step == 0 {
			// 選べるかどうかの判定が必要...
			fmt.Println(canGetDisk(_f, _s))
			if canGetDisk(_f, _s) {
				fmt.Println("えらべるかな")
				sel_stick = _s
				sd := seleDisk(sel_stick, _f)
				// animation select disk
				animSelectDisk(_s, sd, _f, disks)
				step = 1
				fmt.Println("えらべたよ")
			}
		} else {
			fmt.Println("おけるかな")
			sel_move = _s
			fmt.Println(sel_stick, sel_move)
			isMove := moveDisk(sel_stick, sel_move, _f)
			fmt.Println(isMove)
			if isMove {
				// move animation
				animPutDisk(_s, sel_disk, putPosY, _f, disks, 1)
			} else {
				// put animation
				animPutDisk(animPrePos[0], sel_disk, 0, _f, disks, 0)
			}
			game.PrintField(_f)
			if game.CheckEnd(_f) {
				fmt.Println("Complete!")
				isEnd = true
			}
			step = 0
			fmt.Println("は？")
		}
	}
}

func canGetDisk(_f [][]int, _s int) bool {
	var selDisk int
	// get disk
	for i := 0; i < 3; i++ {
		if _f[i][_s-1] != 0 {
			animPrePos[1] = i
			animPrePos[0] = _s - 1
			selDisk = _f[i][_s-1]
			//_f[i][_s-1] = 0
			break
		}
	}
	if selDisk == 0 {
		return false
	}
	return true
}

func animSelectDisk(_s, sd int, _f [][]int, disks []*canvas.Image) {
	// _sは位置、_sdは選択したディスク,
	// get disk
	/*
		for i := 0; i < 3; i++ {
			if _f[i][_s-1] != 0 {
				get := _f[i][_s-1] - 1
				//fmt.Println(get)
				x := disks[get].Position().X
				animPrePosX = x
				disks[get].Move(fyne.NewPos(x, YPos[0]))
				//return true おけるか確認してから関数呼ぶからいらない。
				break
			}
		}
	*/
	//return false
	x := disks[sd-1].Position().X
	disks[sd-1].Move(fyne.NewPos(x, YPos[0]))
}

func animPutDisk(_s int, disk int, _yy int, _f [][]int, disks []*canvas.Image, _y int) {
	// put disk
	if _y == 0 {
		fmt.Println("check")
		x := XPos[disk-1][_s-1]
		disks[disk-1].Move(fyne.NewPos(x, YPos[_y]))
	}
	/*
		for i := 0; i < 3; i++ {
			if _f[2-i][_s-1] == 0 {
				//get := _f[i][_s-1]
				x := XPos[disk-1][_s-1]
				disks[disk-1].Move(fyne.NewPos(x, YPos[3-i]))
				fmt.Println(x, YPos[3-i])
				break
				// おけるか確認してから関数呼ぶからいらないreturn true
			}
		}
	*/
	x := XPos[disk-1][_s-1]
	y := YPos[_yy+1]
	disks[disk-1].Move(fyne.NewPos(x, y))
	fmt.Println(y)
	//return false
}
func stickInit() [][]int {
	fields := make([][]int, 3)
	for i := 0; i < 3; i++ {
		fields[i] = make([]int, 3)
		fields[i][0] = i + 1
	}
	return fields
}

func seleDisk(_s int, _f [][]int) int {
	var prePos [2]int
	var selDisk int
	// get disk
	for i := 0; i < 3; i++ {
		if _f[i][_s-1] != 0 {
			prePos[1] = i
			prePos[0] = _s - 1
			selDisk = _f[i][_s-1]
			sel_disk = selDisk
			_f[i][_s-1] = 0
			break
		}
	}
	if selDisk == 0 {
		return 0
	}
	return selDisk
}

func moveDisk(_s, _x int, _f [][]int) bool {
	//var selDisk int

	// get disk
	/*
		for i := 0; i < 3; i++ {
			if _f[i][_s-1] != 0 {
				prePos[1] = i
				prePos[0] = _s - 1
				selDisk = _f[i][_s-1]
				sel_disk = selDisk
				_f[i][_s-1] = 0
				break
			}
		}
		if selDisk == 0 {
			return false
		}
	*/

	// move disk
	for i := 0; i < 3; i++ {
		if _f[2-i][_x-1] == 0 {
			if i == 0 { // 一番下なら無条件におk
				_f[2-i][_x-1] = sel_disk
				putPosY = 2 - i
				return true
			} else { // 1、2番目に上の方
				// 自分より下がサイズが大きい時
				if sel_disk < _f[2-i+1][_x-1] {
					_f[2-i][_x-1] = sel_disk
					putPosY = 2 - i
					return true
				}
			}
		}
	}
	_f[prePos[1]][prePos[0]] = sel_disk // 移動できなかった場合、元の位置にディスクを戻す
	return false
}
