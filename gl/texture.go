// +build !windows
// MinGW fucks up the function definition for glPixelMap

package gl

// #ifdef __APPLE__
// # include <OpenGL/gl.h>
// #else
// # include <GL/gl.h>
// #endif
import "C"

//void glPixelMapfv (GLenum map, int mapsize, const float *values)
func PixelMapfv(map_ GLenum, mapsize int, values *float32) {
	C.glPixelMapfv(C.GLenum(map_), C.GLsizei(mapsize), (*C.GLfloat)(values))
}

//void glPixelMapuiv (GLenum map, int mapsize, const uint *values)
func PixelMapuiv(map_ GLenum, mapsize int, values *uint32) {
	C.glPixelMapuiv(C.GLenum(map_), C.GLsizei(mapsize), (*C.GLuint)(values))
}

//void glPixelMapusv (GLenum map, int mapsize, const uint16 *values)
func PixelMapusv(map_ GLenum, mapsize int, values *uint16) {
	C.glPixelMapusv(C.GLenum(map_), C.GLsizei(mapsize), (*C.GLushort)(values))
}
