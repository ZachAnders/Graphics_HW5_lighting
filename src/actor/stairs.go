package actor

import (
	"github.com/go-gl/glh"
	"glutil"
	"os"
	"world"
)

type Stairs struct {
	stairs   []Block
	Position glutil.Point3D
}

func NewStairs(myworld *world.World, position glutil.Point3D) Stairs {
	stairTex := glh.NewTexture(512, 512)
	file, _ := os.Open("stone.png")
	stairTex.FromPngReader(file, 0)
	file.Close()

	size := glutil.Point3D{2, 2, 10}
	trans := glutil.Point3D{2, 0, 0}
	size_inc := glutil.Point3D{0, 2, 0}
	tmp := Stairs{nil, position}

	pt1 := glutil.Point3D{-10, 0, 5 + 2}
	left := NewBlock(myworld, position.Minus(pt1), glutil.Point3D{20, 20, 4}, glutil.Color4D{.5, .3, .1, 1}, false)
	left.Texture = stairTex
	left.TexScale = 6
	myworld.AddActor(&left)
	tmp.stairs = append(tmp.stairs, left)

	pt2 := glutil.Point3D{-10, 0, -5 - 2}
	right := NewBlock(myworld, position.Minus(pt2), glutil.Point3D{20, 20, 4}, glutil.Color4D{.5, .3, .1, 1}, false)
	right.Texture = stairTex
	right.TexScale = 6
	myworld.AddActor(&right)
	tmp.stairs = append(tmp.stairs, right)

	// Backing
	back := NewBlock(myworld, glutil.Point3D{22.5, 0, 0}, glutil.Point3D{10, 20.5, 20}, glutil.Color4D{.5, .3, .1, 1}, false)
	back.Texture = stairTex
	back.TexScale = 6
	myworld.AddActor(&back)
	tmp.stairs = append(tmp.stairs, back)

	for i := 0; i < 10; i++ {
		blk := NewBlock(myworld, position, size, glutil.Color4D{.5, .3, .1, 1}, false)
		blk.Texture = stairTex
		blk.TexScale = 6
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
