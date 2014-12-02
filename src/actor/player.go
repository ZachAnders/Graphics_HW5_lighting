package actor

import (
	//	"fmt"
	"github.com/ianremmler/ode"
	"glutil"
	"math"
	"world"
)

type Player struct {
	*BasicActor
	camera      *glutil.DisplayViewport
	height      float64
	gravity     glutil.Point3D
	ray         *ode.Ray
	world       *world.World
	target      ode.Geom
	target_dist float64
	Grab        bool
}

func (self *Player) Tick() {
	//self.ray.SetEnabled(true)
	//cts := self.ray.Collide(self.world.Space., 1, 0)
	//mt.Println(len(cts))
	self.UpdateState()
}

func NewPlayer(myworld *world.World, position glutil.Point3D, height float64) *Player {
	box := myworld.Space.NewBox(ode.V3(height, height, height))

	mass := ode.NewMass()
	mass.SetBox(1, ode.V3(height, height, height))
	mass.Adjust(1)
	body := myworld.Ode_world.NewBody()
	//body.SetLinearDamping(.45)
	body.SetMass(mass)
	body.SetPosition(position.ToODE())
	body.SetMass(mass)
	box.SetBody(body)

	box.SetPosition(position.ToODE())
	camera := glutil.CreateDefaultViewport()
	gravity := glutil.Point3D{0, -1, 0}
	actor := NewBasicActor(myworld, &box)
	box.SetData(&actor)

	ray := myworld.Space.NewRay(25)

	// Make sure the ray and player do not collide
	ray.SetCategoryBits(1)
	ray.SetCollideBits(1)
	box.SetCategoryBits(2)
	box.SetCollideBits(2)

	player := Player{&actor, &camera, height, gravity, &ray, myworld, nil, 0, false}
	player.UpdateRay()
	ray.SetData(&player)

	return &player
}

func (self *Player) UpdateState() {
	self.updateCameraFromPos()
	self.UpdateRay()
	self.UpdateTarget()
}

func (self *Player) UpdateTarget() {
	if self.target != nil {
		//fmt.Printf("Geom: %s Dist: %f\n", self.target, self.target_dist)
		Ex, Ey, Ez := self.RayCoords()
		pos := glutil.NewODEPoint3D(self.Model.Position())
		x, y, z := pos.Unpack()
		self.target.SetPosition(ode.V3(
			self.target_dist*Ex+x,
			self.target_dist*Ey+y,
			self.target_dist*Ez+z))
	}
}

func (self *Player) RayCoords() (float64, float64, float64) {
	theta, phi := self.camera.GetViewAngle()
	var Ex float64 = -1 * glutil.Sin(theta) * glutil.Cos(phi)
	var Ey float64 = 1 * glutil.Sin(phi)
	var Ez float64 = 1 * glutil.Cos(theta) * glutil.Cos(phi)
	return Ex, Ey, Ez
}

func (self *Player) UpdateRay() {
	pos := self.Model.Position()
	self.ray.SetPosDir(pos, ode.V3(self.RayCoords()))
}

func (self *Player) Translate(dx float64, dy float64, dz float64) {
	rel_dx := (-glutil.Sin(self.camera.Theta) * dx) +
		(glutil.Cos(self.camera.Theta) * dz)
	rel_dz := (glutil.Cos(self.camera.Theta) * dx) +
		(glutil.Sin(self.camera.Theta) * dz)

	pos := self.Model.Position()
	//self.Model.SetPosition(ode.V3(rel_dx+pos[0], pos[1], rel_dz+pos[2]))
	self.Model.SetPosition(ode.V3(pos[0], pos[1]+.025, pos[2]))
	self.Model.Body().SetForce(ode.V3(20*rel_dx, 20*dy, 20*rel_dz))
	//self.Model.Body().SetForce(ode.V3(20*rel0*rel_dx, 20*dy, 20*rel_dz))
	//self.Model.Body().SetForce(ode.V3(0, 20*dy, 0))
	//self.World.TranslateActor(self, rel_dx, dy, rel_dz)
}

func (self *Player) Rotate(dtheta, dphi int32) {
	self.camera.Rotate(dtheta, dphi)
	self.UpdateState()
}
func (self *Player) GetViewAngle() (int32, int32) {
	return self.camera.GetViewAngle()
}
func (self *Player) ImmediateLook() {
	self.UpdateState()
}
func (self *Player) ImmediateJump() {
	//	self.Position.Center.X = self.camera.Position.X
	//	self.Position.Center.Y = self.camera.Position.Y
	//	self.Position.Center.Z = self.camera.Position.Z
}
func (self *Player) PositionSelf() {
	self.camera.PositionSelf()
}
func (self *Player) ToString() string {
	return self.camera.ToString()
}

func (self *Player) updateCameraFromPos() {
	self.camera.Position = glutil.NewODEPoint3D(self.Model.Position())
}

func (self *Player) GetProjection() *glutil.DisplayProjection {
	return &self.camera.Projection
}

func (self *Player) CanObstruct() bool {
	return true
}

func (self *Player) Pickup(geom ode.Geom) {
	self.target = geom
	pos := glutil.NewODEPoint3D(geom.Position())
	pos = pos.Minus(glutil.NewODEPoint3D(self.Model.Position()))
	x, y, z := pos.Unpack()

	self.target_dist = math.Sqrt((x * x) + (y * y) + (z * z))
}

func (self *Player) GrabThrow() {
	if self.target == nil {
		self.Grab = true
	} else {
		x, y, z := self.RayCoords()
		self.target.Body().SetForce(ode.V3(1000*x, 1000*y+200, 1000*z))
		self.target = nil
		self.Grab = false
	}
}

func (self *Player) Interact(my_geom, other_geom ode.Geom) bool {
	if _, ok := my_geom.(ode.Ray); ok {
		if self.Grab && self.target == nil {
			self.Grab = false
			//fmt.Println(other_geom.Data())
			if other_geom.Data() != nil && self.target != other_geom {
				// Do not pick up anonymous geom, do not attempt to re-pickup
				// already held geom
				if _, ok := other_geom.Data().(*Block); ok {
					// Only pick up blocks
					self.Pickup(other_geom)
				}
			}
		}
		// Don't let ODE make a physical joint between the Ray and other Geom
		return false
	}
	return true
}
