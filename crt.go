//go:build ignore

// Reference: a public domain CRT effect
// https://github.com/libretro/glsl-shaders/blob/master/crt/shaders/crt-lottes.glsl

package main

func Warp(pos vec2) vec2 {
	const (
		warpX = 0.031
		warpY = 0.041
	)
	pos = pos*2 - 1
	pos *= vec2(1+(pos.y*pos.y)*warpX, 1+(pos.x*pos.x)*warpY)
	return pos/2 + 0.5
}

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {

	// generate scan line effect
	ambient := 1.0
	if int(mod(position.y, 2)) == int(0) {
		ambient = 0.25
	}

	// Adjust the texture position to [0, 1].
	pos := texCoord
	origin, size := imageSrcRegionOnTexture()
	pos -= origin
	pos /= size

	pos = Warp(pos)

	return imageSrc0At(pos*size+origin) * (ambient)
	// return imageSrc0UnsafeAt(pos*size+origin) * (ambient)
}
