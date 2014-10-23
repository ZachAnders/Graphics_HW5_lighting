package glutil

import (
	"fmt"
	"github.com/go-gl/gl"
	"github.com/rhencke/glut"
	"math"
	"reflect"
)

// Displays the given format string and arguments at the current position
func Print(format string, vargs ...interface{}) {
	var tmp string
	if len(vargs) != 0 {
		tmp = fmt.Sprintf(format, vargs)
		fval := reflect.ValueOf(fmt.Sprintf) //.(*reflect.FuncValue)
		print_args := []reflect.Value{reflect.ValueOf(format)}
		// TODO: Make sure vargs[0] is actually a []interface{}
		for _, val := range vargs[0].([]interface{}) {
			print_args = append(print_args, reflect.ValueOf(val))
		}
		tmp = fval.Call(print_args)[0].String()
	} else {
		tmp = fmt.Sprintf("%s", format)
	}
	for _, eachChar := range tmp {
		glut.BITMAP_HELVETICA_12.Character(eachChar)
	}
}

// Displays the given format string and arguments at the location
// (x, y), which is from the lower right corner
func Print2d(x int, y int, format string, vargs ...interface{}) {
	// This function based off of the code at
	// http://programmingexamples.net/wiki/OpenGL/Text
	// This is due to the fact that golang is missing the WindowPos* extension
	_, _, width, height := gl.GetInteger4(gl.VIEWPORT)

	gl.MatrixMode(gl.PROJECTION)
	gl.PushMatrix()
	gl.LoadIdentity()
	gl.Ortho(0, float64(width), 0, float64(height), -1.0, 1.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.PushMatrix()
	gl.LoadIdentity()

	gl.RasterPos2i(x, y)
	Print(format, vargs)

	gl.MatrixMode(gl.PROJECTION)
	gl.PopMatrix()
	gl.MatrixMode(gl.MODELVIEW)
	gl.PopMatrix()
}

func Sin(th int32) float64 {
	return math.Sin((float64(th) / 180) * math.Pi)
}

func Cos(th int32) float64 {
	return math.Cos((float64(th) / 180) * math.Pi)
}
