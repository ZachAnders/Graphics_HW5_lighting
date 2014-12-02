package world

import (
	"github.com/go-gl/gl"
)

type DisplayListManager struct {
	Mappings map[int]int
	inList   bool
}

func NewDisplayListManager() DisplayListManager {
	return DisplayListManager{make(map[int]int), false}
}

func (self *DisplayListManager) CachedRender(key int, obj Renderable) {
	if val, ok := self.Mappings[key]; ok {
		gl.CallList(uint(val))
		return
	}
	val := gl.GenLists(1)
	self.Mappings[key] = int(val)
	gl.NewList(val, gl.COMPILE)
	obj.Render()
	gl.EndList()
}
