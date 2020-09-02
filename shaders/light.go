// build +ignore
package shaders

/*
var LightColor [40]vec3
var LightIntensity [40]float
var LightPoint [40]vec2
var LightAmount float

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	finalcol := vec4(0.0, 0.0, 0.0, 0.0)
	for i := 0; i < 40; i++ {
		if float(i) == LightAmount {
			break
		}
		intensity := LightIntensity[i]
		distance := length(LightPoint[i] - position.xy)*256
		attenuation := intensity / (4*3.1415926538*pow(distance, 2))
		col := vec4(attenuation, attenuation, attenuation, pow(attenuation, 3)) * vec4(LightColor[i], 1)
		finalcol = finalcol + col
	}
	return image0TextureAt(texCoord) * finalcol
}
*/
