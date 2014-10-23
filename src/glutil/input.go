package glutil

type MouseState struct {
	LeftDown   bool
	MiddleDown bool
	RightDown  bool
	X          int
	Y          int
}

func CreateMouseState() MouseState {
	return MouseState{false, false, false, 0, 0}
}
