package actor

import (
	"glutil"
	"world"
)

type Stairs struct {
	stairs   []Block
	Position glutil.Point3D
}

func NewStairs(myworld *world.World, position glutil.Point3D) Stairs {
	size := glutil.Point3D{2, 2, 10}
	trans := glutil.Point3D{2, 0, 0}
	size_inc := glutil.Point3D{0, 2, 0}
	tmp := Stairs{nil, position}

	pt1 := glutil.Point3D{-10, 0, 5 + 2}
	left := NewBlock(myworld, position.Minus(pt1), glutil.Point3D{20, 20, 4}, glutil.Color4D{.5, .3, .1, 1})
	myworld.AddActor(&left)
	tmp.stairs = append(tmp.stairs, left)

	pt2 := glutil.Point3D{-10, 0, -5 - 2}
	right := NewBlock(myworld, position.Minus(pt2), glutil.Point3D{20, 20, 4}, glutil.Color4D{.5, .3, .1, 1})
	myworld.AddActor(&right)
	tmp.stairs = append(tmp.stairs, right)

	for i := 0; i < 10; i++ {
		blk := NewBlock(myworld, position, size, glutil.Color4D{.5, .3, .1, 1})
		//blk.Offset = i
		size = size.Add(size_inc)
		position = position.Add(trans)
		tmp.stairs = append(tmp.stairs, blk)
		myworld.AddActor(&blk)
	}
	return tmp
}

func (self *Stairs) Render() {
	for _, elem := range self.stairs {
		elem.Render()
	}
}
