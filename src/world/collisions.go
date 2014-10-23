package world

import (
	//	"fmt"
	"glutil"
	"math"
)

type BoundingBox struct {
	Center glutil.Point3D
	Size   glutil.Point3D
}

func (self *BoundingBox) GetCenter() *glutil.Point3D {
	return &self.Center
}

func (self *BoundingBox) GetSize() *glutil.Point3D {
	return &self.Size
}
func NewBoundingBox(center, size glutil.Point3D) BoundingBox {
	size.X /= 2
	size.Y /= 2
	size.Z /= 2
	//	fmt.Printf("Alloc'd Bbox Center %s, size %s\n", center.ToString(), size.ToString())
	return BoundingBox{center, size}
}

func (self *BoundingBox) Collides(other_entity Entity) bool {
	switch other := other_entity.(type) {
	case *BoundingBox:
		if math.Abs(self.Center.X-other.Center.X) > self.Size.X+other.Size.X {
			return false
		}
		if math.Abs(self.Center.Y-other.Center.Y) > self.Size.Y+other.Size.Y {
			return false
		}
		if math.Abs(self.Center.Z-other.Center.Z) > self.Size.Z+other.Size.Z {
			return false
		}
		return true
	}
	panic("Unknown type to collide with!")
}

func (self *BoundingBox) HeightAt(position glutil.Point3D) (bool, float64) {
	if math.Abs(self.Center.X-position.X) > self.Size.X {
		return false, 0
	}
	if math.Abs(self.Center.Z-position.Z) > self.Size.Z {
		return false, 0
	}
	return true, self.Center.Y + self.Size.Y
}
