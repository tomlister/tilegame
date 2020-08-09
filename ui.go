package main

//UIcreateSpeechBubble Creates a speech bubble
func (world *World) UIcreateSpeechBubble(textstr string, x, y, width int) {
	world.createText(textstr, x, y, width, true)
}
