package shaders

//TitleShader provides access to the title shader
func TitleShader() []byte {
	return title_go
}

//LightShader provides access to the light shader
func LightShader() []byte {
	return light_go
}

//BlurShader provides access to the blur shader
func BlurShader() []byte {
	return blur_go
}
