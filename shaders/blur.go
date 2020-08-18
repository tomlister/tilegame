// build +ignore
package shaders

/*
var Size vec3

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	quality := 6
	directions := 12
	tau := 6.28318530718
	radius := Size.z / Size.xy
	col := image0TextureAt(texCoord)
	for d := 0.0; d < tau; d += tau / float(directions) {
		for i := 1.0 / float(quality); i <= 1.0; i += 1.0 / float(quality) {
			col += image0TextureAt(texCoord + vec2(cos(d), sin(d))*i)
		}
	}
	col /= float(quality)*float(directions) + 1.0
	return col * image0TextureAt(texCoord)
}*/
