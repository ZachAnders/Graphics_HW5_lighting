package util

import (
	"world"
)

type RenderQueue struct {
	Renderables []world.Renderable
	lookupTable map[string]world.Renderable
}

func NewEmptyRenderQueue() RenderQueue {
	queue := make([]world.Renderable, 0)
	lookupTable := make(map[string]world.Renderable)
	return RenderQueue{queue, lookupTable}
}

func NewRenderQueue(entities ...world.Renderable) RenderQueue {
	queue := make([]world.Renderable, 0)
	lookupTable := make(map[string]world.Renderable)
	rq := RenderQueue{queue, lookupTable}
	for _, entity := range entities {
		rq.Add(entity)
	}
	return rq
}

func (self *RenderQueue) Add(entity world.Renderable) {
	self.Renderables = append(self.Renderables, entity)
}

func (self *RenderQueue) AddNamed(entity world.Renderable, name string) {
	self.Renderables = append(self.Renderables, entity)
	self.lookupTable[name] = entity
}

func (self *RenderQueue) Get(name string) world.Renderable {
	return self.lookupTable[name]
}

func (self *RenderQueue) RenderAll() {
	for _, entity := range self.Renderables {
		entity.Render()
	}
}
