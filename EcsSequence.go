package main

type Sequence struct {
	frames       []int
	currentFrame int
	fps          float64
	loop         bool
}

func NewSequence(frames []int, fps float64, loop bool) *Sequence {

	var seq Sequence
	seq.frames = frames
	seq.fps = fps
	seq.loop = loop

	return &seq
}

func (seq *Sequence) nextFrame() bool {
	if seq.currentFrame == len(seq.frames)-1 {
		if seq.loop {
			seq.currentFrame = 0
		} else {
			return true
		}
	} else {
		seq.currentFrame++
	}
	return false
}
