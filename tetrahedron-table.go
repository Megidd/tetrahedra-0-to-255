package main

// Specify the edges & corners used to create the tetrahedra.
// Keep the index from 0 to 11 for edges,
// but shift the index for corners.
// Corners were originally indexed from 0 to 7.
// So, corners would be indexed from 0+12 to 7+12 i.e. from 12 to 19.
// Inspired by:
// https://github.com/deadsy/sdfx/blob/1a71e404e4b2aa00c59f53cffc219a9e83e62d85/render/march3.go#L360
var mcTetrahedronTable = [256][]int{
	// 0b00000000 case 0: no cube corner has zero/negative value.
	{},
	// 0b00000001 case 1: first cube corner has zero/negative value.
	{12, 0, 3, 8},
	// 0b00000010
	{},
	// 0b00000011
	{},
	// 0b00000100
	{},

	// ...

	// 0b11111100
	{},
	// 0b11111101
	{},
	// 0b11111110
	{},
	// 0b11111111 case 255: all cube corners have zero/negative values.
	{},
}
