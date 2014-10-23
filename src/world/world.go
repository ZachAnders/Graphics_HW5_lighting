package world

import (
	"glutil"
)

type World struct {
	Entities      []Entity
	Actors        []Actor
	gravity       glutil.Point3D
	actor_counter int
}

func NewWorld() World {
	return World{nil, nil, glutil.Point3D{0, -.5, 0}, 0}
}

func (self *World) TranslateActor(actor Actor, dx, dy, dz float64) {
	var highestPoint *float64 = new(float64)
	actor_center := actor.GetCenter()
	next := actor_center.Add(glutil.Point3D{dx, dy, dz})

	if actor.CanClip() {
		// Skip collision detection if can clip
		actor_center.X += dx
		actor_center.Y += dy
		actor_center.Z += dz
	}

	for _, entity := range self.Entities {
		success, value := entity.HeightAt(next)
		if success {
			if highestPoint == nil || value > *highestPoint {
				*highestPoint = value
			}
		}
	}

	for _, ea_actor := range self.Actors {
		if ea_actor.GetID() != actor.GetID() {
			//fmt.Printf("Obj1: %s Obj2 %s\n", actor.GetBoundingBox().Center.ToString(),
			//				ea_actor.GetBoundingBox().Center.ToString())
			success, value := ea_actor.HeightAt(next)
			if success && actor.CanObstruct() {
				if highestPoint == nil || value > *highestPoint {
					if *highestPoint < next.Y {
						*highestPoint = value
					}
				}
			}
		}
	}

	if highestPoint != nil {
		//fmt.Printf("Highest: %f\n", *highestPoint)
		player_bot := actor_center.Y - actor.GetSize().Y
		diff := *highestPoint - player_bot
		if diff < 3 {
			actor_center.X += dx
			actor_center.Z += dz
			if diff >= 0 {
				actor_center.Y = *highestPoint + actor.GetSize().Y
			} else {
				actor_center.Y += dy
			}

		}
	} else {
		actor_center.X += dx
		actor_center.Y += dy
		actor_center.Z += dz
	}
}

func (self *World) Tick() {
	for _, actor := range self.Actors {
		actor.Tick()
		if actor.CanFall() {
			actor.Translate(self.gravity.Unpack())
		}
	}
}

func (self *World) AddEntity(thing Entity) {
	self.Entities = append(self.Entities, thing)
}

func (self *World) AddActor(thing Actor) {
	thing.SetID(self.actor_counter)
	self.actor_counter++
	self.Actors = append(self.Actors, thing)
}
