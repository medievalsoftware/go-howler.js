package howler

func SetStereo(pan float64) {
	howler.Call("stereo", pan)
}

func Pos() (x, y, z float64) {
	arr := howler.Call("pos")
	x = arr.Index(0).Float()
	y = arr.Index(1).Float()
	z = arr.Index(2).Float()
	return
}

func SetPos(x, y, z float32) {
	howler.Call("pos", x, y, z)
}

// Orientation gets the direction the listener is pointing in the 3D cartesian
// space. A front and up vector must be provided. The front is the direction the
// face of the listener is pointing, and up is the direction the top of the
// listener is pointing. Thus, these values are expected to be at right angles
// from each other.
func Orientation() (x, y, z, upX, upY, upZ float64) {
	arr := howler.Call("orientation")
	x = arr.Index(0).Float()
	y = arr.Index(1).Float()
	z = arr.Index(2).Float()
	upX = arr.Index(3).Float()
	upY = arr.Index(4).Float()
	upZ = arr.Index(5).Float()
	return
}

func SetOrientation(x, y, z, upX, upY, upZ float64) {
	howler.Call("orientation", x, y, z, upX, upY, upZ)
}
