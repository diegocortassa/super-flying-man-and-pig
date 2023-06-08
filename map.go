package main

import "image"

var gameMap = []int{
	0, 1, 0, 1, 1, 0, 1, 1,
	1, 0, 1, 0, 1, 0, 1, 0,
	0, 1, 0, 1, 1, 1, 0, 1,
	0, 0, 0, 1, 0, 0, 0, 0,
	0, 1, 0, 1, 0, 1, 0, 0,
	0, 0, 0, 1, 1, 0, 1, 0,
	1, 0, 1, 0, 0, 0, 1, 1,
	0, 0, 1, 0, 1, 0, 1, 0,
	0, 1, 1, 0, 1, 1, 1, 0,
	0, 1, 0, 1, 0, 1, 1, 1,
	0, 0, 1, 1, 0, 1, 0, 15,
	2, 3, 4, 5, 6, 7, 10, 11,
	18, 17, 83, 62, 63, 8, 9, 14,
	17, 16, 62, 60, 61, 16, 12, 13,
	18, 60, 16, 17, 16, 83, 16, 17,
	83, 18, 17, 16, 18, 64, 83, 59,
	63, 16, 65, 18, 63, 18, 17, 17,
	17, 18, 17, 17, 56, 57, 18, 17,
	18, 54, 18, 17, 55, 58, 18, 17,
	17, 19, 17, 18, 18, 17, 18, 17,
	18, 17, 18, 18, 18, 21, 18, 17,
	18, 17, 18, 17, 18, 18, 17, 18,
	17, 22, 20, 17, 18, 18, 76, 18,
	18, 18, 17, 59, 18, 23, 24, 18,
	17, 55, 18, 18, 17, 26, 25, 18,
	59, 18, 18, 18, 83, 17, 18, 18,
	17, 27, 18, 17, 20, 54, 57, 17,
	83, 17, 17, 18, 17, 17, 18, 17,
	18, 56, 18, 18, 59, 18, 22, 18,
	54, 18, 18, 18, 18, 18, 18, 18,
	28, 29, 30, 31, 32, 33, 34, 35,
	41, 36, 39, 41, 39, 36, 38, 37,
	37, 38, 37, 38, 36, 37, 38, 40,
	39, 38, 51, 52, 41, 38, 41, 37,
	40, 36, 50, 53, 40, 41, 38, 40,
	37, 40, 37, 41, 38, 40, 36, 38,
	37, 36, 37, 36, 37, 36, 37, 36,
	36, 37, 36, 37, 36, 37, 36, 37,
	37, 36, 39, 39, 36, 39, 37, 36,
	36, 37, 36, 37, 36, 37, 36, 37,
	37, 36, 37, 36, 37, 36, 37, 36,
	36, 41, 51, 52, 36, 41, 36, 41,
	41, 36, 50, 53, 41, 36, 41, 40,
	40, 41, 36, 38, 40, 38, 40, 41,
	41, 36, 38, 36, 38, 36, 41, 40,
	40, 38, 38, 41, 40, 38, 40, 41,
	40, 37, 40, 36, 40, 36, 41, 40,
	39, 40, 39, 40, 39, 40, 39, 40,
	40, 51, 52, 39, 40, 39, 36, 39,
	39, 50, 53, 40, 39, 37, 36, 40,
	40, 37, 40, 39, 40, 39, 40, 39,
	39, 40, 39, 40, 37, 40, 39, 37,
	42, 43, 44, 45, 46, 47, 48, 49,
	18, 57, 18, 18, 18, 18, 17, 54,
	17, 18, 17, 56, 18, 17, 18, 18,
	17, 18, 18, 18, 17, 18, 17, 17,
	18, 54, 18, 17, 18, 18, 17, 18,
	18, 17, 18, 18, 18, 17, 57, 18,
	17, 17, 66, 72, 17, 17, 17, 17,
	69, 68, 75, 71, 73, 17, 76, 18,
	70, 71, 75, 71, 75, 74, 18, 17,
	75, 71, 75, 81, 75, 71, 73, 18,
	75, 67, 75, 71, 71, 75, 71, 74,
	75, 75, 86, 75, 75, 75, 71, 75,
	71, 71, 71, 71, 75, 75, 71, 82,
	71, 71, 71, 75, 75, 75, 71, 75,
	71, 75, 75, 80, 77, 75, 75, 75,
	75, 75, 75, 79, 78, 75, 75, 75,
	75, 82, 71, 71, 71, 75, 75, 75,
	71, 75, 75, 75, 75, 81, 75, 71,
	75, 75, 75, 71, 75, 71, 75, 75,
	75, 75, 75, 75, 75, 71, 75, 75,
	75, 75, 75, 75, 75, 75, 75, 75,
	71, 71, 75, 71, 71, 86, 71, 75,
	75, 75, 82, 71, 71, 71, 75, 75,
	71, 75, 71, 75, 75, 71, 71, 71,
	71, 75, 75, 75, 75, 71, 75, 75,
	71, 81, 71, 71, 71, 71, 82, 71,
	71, 71, 71, 71, 71, 71, 75, 71,
	71, 75, 75, 75, 75, 75, 71, 71,
	75, 71, 75, 71, 71, 82, 75, 71,
	75, 75, 75, 75, 71, 71, 75, 75,
	75, 75, 67, 75, 75, 75, 75, 75,
	75, 75, 75, 75, 75, 71, 75, 75,
	86, 75, 75, 75, 75, 75, 81, 75,
	85, 84, 85, 84, 85, 84, 85, 84,
	75, 75, 75, 71, 75, 75, 71, 81,
	71, 75, 71, 75, 71, 71, 75, 75,
	82, 75, 75, 75, 75, 71, 71, 75,
	75, 71, 71, 75, 71, 86, 75, 71,
	71, 75, 71, 75, 71, 71, 71, 75,
	75, 71, 71, 75, 81, 75, 71, 75,
	71, 80, 77, 71, 75, 71, 67, 75,
	71, 79, 78, 71, 71, 71, 75, 71,
	75, 71, 75, 71, 75, 82, 71, 75,
	82, 86, 75, 71, 75, 71, 75, 71,
	75, 71, 71, 71, 75, 75, 71, 71,
	84, 85, 84, 85, 84, 85, 84, 85,
	75, 75, 75, 75, 75, 75, 75, 75,
	75, 82, 71, 86, 75, 80, 77, 71,
	75, 67, 71, 75, 81, 79, 78, 71,
	75, 71, 71, 75, 71, 75, 71, 75,
	87, 88, 75, 81, 75, 75, 75, 75,
	17, 17, 88, 75, 71, 75, 75, 75,
	17, 17, 17, 88, 71, 75, 82, 75,
	17, 17, 17, 17, 88, 71, 71, 75,
	54, 17, 17, 17, 17, 88, 75, 71,
	17, 18, 17, 18, 17, 17, 88, 71,
	17, 18, 17, 17, 18, 17, 17, 88,
	17, 17, 18, 59, 18, 18, 17, 17,
	17, 17, 17, 18, 17, 17, 18, 17,
	18, 17, 17, 17, 18, 17, 18, 17,
	55, 17, 18, 17, 17, 18, 17, 17,
	17, 18, 17, 17, 17, 18, 17, 17,
	18, 17, 17, 18, 17, 17, 17, 17,
	17, 17, 18, 17, 17, 18, 54, 17,
	18, 18, 55, 17, 17, 17, 18, 17,
	17, 17, 18, 17, 17, 56, 57, 17,
	17, 18, 63, 65, 17, 55, 58, 17,
	54, 60, 18, 18, 16, 18, 18, 17,
	18, 17, 16, 83, 17, 18, 62, 17,
	18, 17, 17, 18, 18, 69, 18, 17,
	18, 18, 18, 18, 17, 17, 17, 17,
	18, 59, 18, 17, 18, 17, 17, 18,
	17, 18, 18, 22, 21, 17, 58, 17,
	17, 19, 17, 18, 17, 17, 22, 17,
	28, 29, 30, 31, 32, 33, 34, 35,
	36, 36, 36, 36, 38, 36, 38, 37,
	36, 37, 38, 38, 36, 37, 36, 36,
	36, 38, 39, 38, 36, 38, 40, 36,
	37, 36, 37, 36, 37, 36, 37, 37,
	36, 51, 52, 36, 40, 39, 36, 36,
	36, 50, 53, 37, 39, 36, 37, 36,
	36, 36, 38, 36, 37, 38, 36, 37,
	37, 37, 38, 40, 36, 38, 39, 36,
	36, 38, 39, 38, 36, 38, 36, 38,
	42, 43, 44, 45, 46, 47, 48, 49,
	18, 17, 17, 17, 17, 18, 17, 17,
	17, 18, 62, 18, 18, 17, 17, 83,
	17, 16, 83, 17, 16, 18, 63, 17,
	18, 17, 16, 61, 18, 64, 18, 17,
	63, 18, 18, 17, 83, 17, 16, 18,
	17, 54, 17, 62, 17, 17, 17, 18,
	18, 17, 17, 18, 17, 17, 17, 17,
	17, 17, 18, 17, 17, 18, 17, 54,
	17, 18, 17, 55, 18, 17, 17, 18,
	17, 59, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17,
	18, 17, 18, 17, 18, 17, 56, 17,
	17, 18, 17, 17, 19, 17, 18, 17,
	17, 57, 17, 18, 17, 18, 17, 17,
	17, 18, 17, 18, 17, 18, 17, 17,
	18, 17, 18, 17, 17, 17, 18, 17,
	17, 17, 17, 17, 17, 17, 59, 17,
	17, 23, 24, 17, 18, 17, 18, 17,
	18, 26, 25, 17, 18, 18, 18, 17,
	17, 18, 17, 20, 18, 17, 18, 17,
	17, 18, 17, 17, 18, 18, 17, 18,
	18, 17, 18, 17, 18, 17, 18, 17,
	17, 19, 22, 17, 17, 18, 17, 18,
	17, 17, 17, 17, 17, 17, 17, 17,
	59, 18, 17, 18, 17, 18, 18, 20,
	17, 17, 17, 18, 18, 17, 18, 18,
	18, 21, 18, 18, 18, 17, 18, 54,
	18, 17, 18, 17, 18, 17, 18, 17,
	17, 18, 18, 17, 18, 17, 22, 17,
	17, 17, 18, 17, 17, 18, 17, 17,
	17, 54, 22, 17, 18, 17, 18, 17,
	18, 17, 17, 17, 18, 18, 17, 18,
	89, 90, 91, 92, 91, 92, 91, 92,
	93, 95, 93, 94, 95, 94, 93, 109,
	94, 108, 110, 93, 93, 93, 94, 93,
	93, 106, 112, 93, 93, 94, 95, 94,
	94, 95, 94, 95, 94, 93, 94, 93,
	93, 94, 93, 94, 93, 93, 93, 94,
	93, 93, 95, 93, 94, 93, 95, 93,
	93, 93, 93, 95, 93, 95, 93, 94,
	94, 93, 95, 93, 95, 109, 95, 93,
	93, 95, 93, 94, 93, 95, 93, 94,
	94, 93, 94, 95, 94, 93, 95, 93,
	94, 109, 94, 93, 95, 93, 94, 93,
	93, 95, 93, 95, 93, 95, 93, 94,
	94, 93, 95, 93, 94, 93, 95, 93,
	93, 94, 93, 95, 93, 94, 93, 94,
	94, 93, 94, 93, 95, 93, 94, 93,
	94, 95, 94, 108, 97, 110, 94, 95,
	95, 94, 93, 107, 83, 111, 95, 94,
	94, 95, 93, 106, 89, 112, 93, 95,
	95, 93, 95, 93, 95, 94, 95, 109,
	94, 95, 94, 95, 94, 95, 93, 95,
	95, 94, 95, 93, 95, 94, 95, 94,
	93, 109, 93, 94, 95, 93, 94, 93,
	96, 97, 96, 97, 98, 94, 94, 95,
	56, 18, 17, 18, 100, 99, 95, 94,
	17, 18, 54, 18, 59, 100, 99, 94,
	18, 21, 17, 18, 17, 22, 100, 99,
	18, 18, 18, 55, 18, 18, 54, 18,
	17, 59, 17, 18, 18, 18, 17, 18,
	18, 17, 18, 59, 18, 18, 18, 18,
	54, 20, 60, 17, 17, 59, 18, 56,
	17, 18, 62, 61, 65, 18, 17, 17,
	17, 17, 18, 64, 17, 17, 21, 17,
	17, 56, 17, 17, 22, 17, 17, 18,
	19, 18, 17, 18, 17, 18, 57, 17,
	18, 70, 73, 17, 18, 17, 70, 73,
	68, 71, 71, 73, 72, 68, 71, 71,
	71, 75, 71, 71, 71, 75, 71, 75,
	75, 71, 75, 75, 71, 75, 71, 75,
	101, 101, 101, 101, 101, 101, 101, 101,
	103, 102, 102, 102, 103, 103, 102, 103,
	103, 103, 103, 102, 102, 103, 103, 103,
	102, 103, 103, 102, 103, 102, 103, 103,
	104, 105, 105, 104, 105, 104, 105, 104,
	0, 0, 1, 0, 1, 1, 1, 0,
	0, 1, 0, 1, 0, 1, 0, 1,
	0, 0, 1, 1, 0, 1, 0, 1,
	0, 1, 0, 1, 0, 1, 0, 0,
	1, 1, 1, 1, 1, 1, 0, 1,
	0, 0, 1, 1, 0, 1, 0, 1,
	1, 1, 0, 0, 1, 1, 0, 0,
	1, 1, 1, 1, 0, 1, 1, 0,
	0, 0, 1, 0, 1, 0, 1, 0,
	1, 0, 0, 0, 1, 0, 0, 1,
	0, 0, 1, 0, 1, 0, 1, 0,
	0, 1, 0, 0, 1, 0, 1, 0,
	1, 0, 0, 1, 0, 0, 0, 1,
	0, 1, 0, 0, 1, 0, 1, 0,
	0, 1, 0, 0, 0, 0, 1, 0,
	1, 0, 1, 0, 1, 0, 0, 0,
	0, 1, 1, 1, 0, 0, 0, 1,
	0, 1, 1, 0, 1, 0, 0, 0,
	0, 1, 0, 1, 0, 0, 1, 0,
	0, 1, 0, 0, 1, 1, 0, 0,
	0, 1, 0, 1, 1, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 1, 1,
	1, 0, 1, 0, 0, 0, 0, 1,
	0, 0, 1, 1, 1, 1, 0, 0,
	0, 1, 0, 1, 1, 0, 1, 1,
	1, 0, 1, 1, 1, 0, 1, 0,
	1, 0, 0, 1, 0, 0, 1, 0,
	0, 1, 1, 1, 1, 1, 1, 0,
}

func getTile(tileIndex int) image.Image {
	return tilesImage.SubImage(image.Rect(tileSize*(tileIndex), 0, tileSize*(tileIndex+1), tileSize))
}
