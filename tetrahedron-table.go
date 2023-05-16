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
	// 0b00000010 case 2
	{13, 1, 0, 9},
	// 0b00000011 case 3
	{
		12, 13, 3, 8,
		13, 1, 3, 8,
		13, 1, 8, 9,
	},
	// 0b00000100 case 4
	{14, 2, 1, 10},
	// 0b00000101 case 5
	{
		12, 0, 3, 8,
		14, 2, 1, 10,
	},
	// 0b00000110 case 6
	{
		13, 14, 0, 9,
		14, 2, 0, 9,
		14, 2, 9, 10,
	},
	// 0b00000111 case 7
	{
		2, 15, 3, 19,
		2, 3, 8, 19,
		2, 8, 10, 19,
		10, 19, 8, 16,
		10, 19, 16, 18,
		10, 8, 9, 16,
		9, 10, 16, 18,
		9, 18, 16, 17,
	},
	// 0b00001000 case 8
	{
		2, 11, 3, 19,
		2, 19, 3, 16,
		2, 3, 12, 16,
		13, 2, 12, 16,
		16, 2, 19, 18,
		13, 2, 16, 18,
		13, 18, 16, 17,
		14, 2, 13, 18,
	},

	// ...

	// 0b11111100
	{},
	// 0b11111101
	{},
	// 0b11111110
	{},
	// 0b11111111 case 255: all cube corners have zero/negative values.
	{
		12, 16, 19, 18,
		12, 15, 19, 18,
		12, 16, 17, 18,
		12, 13, 17, 18,
		12, 15, 14, 18,
		12, 13, 14, 18,
	},
}
