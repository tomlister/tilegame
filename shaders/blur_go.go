// Code generated by file2byteslice. DO NOT EDIT.
// (gofmt is fine after generating)

package shaders

var blur_go = []byte("// build +ignore\npackage shaders\n\nvar Size vec3\n\nfunc Fragment(position vec4, texCoord vec2, color vec4) vec4 {\n\tquality := 6\n\tdirections := 12\n\ttau := 6.28318530718\n\tradius := Size.z / Size.xy\n\tcol := image0TextureAt(texCoord)\n\tfor d := 0.0; d < tau; d += tau / float(directions) {\n\t\tfor i := 1.0 / float(quality); i <= 1.0; i += 1.0 / float(quality) {\n\t\t\tcol += image0TextureAt(texCoord + vec2(cos(d), sin(d))*i)\n\t\t}\n\t}\n\tcol /= float(quality)*float(directions) + 1.0\n\treturn col * image0TextureAt(texCoord)\n}\n")