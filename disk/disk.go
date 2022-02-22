package disk

type Disk struct {
	Size       int // disk size
	Position_x int // disk position x
	Position_y int // disk position y
}

func (d *Disk) Move(_x int, _y int) {
	d.Position_x = _x
	d.Position_y = _y
}

/* 後々
func (o Object) GetSize() int{
	return o.size
}

func(o Object) GetPosition()
*/
