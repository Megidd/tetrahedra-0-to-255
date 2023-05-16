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
		12, 13, 3, 8,
		13, 2, 3, 8,
		13, 14, 2, 8,
		14, 2, 8, 10,
		13, 14, 8, 10,
		13, 10, 8, 9,
	},
	// 0b00001000 case 8
	{
		15, 3, 2, 11,
	},
	// 0b00001001 case 9
	{
		12, 0, 2, 11,
		12, 2, 15, 11,
		12, 0, 11, 8,
	},
	// 0b00001010 case 10
	{
		0, 13, 1, 9,
		2, 15, 3, 11,
	},
	// 0b00001011 case 11
	{
		15, 12, 2, 11,
		12, 1, 2, 11,
		12, 13, 1, 11,
		13, 1, 11, 9,
		12, 13, 11, 9,
		11, 12, 9, 8,
	},
	// 0b00001100 case 12
	{
		1, 14, 3, 10,
		14, 15, 3, 10,
		15, 3, 10, 11,
	},
	// 0b00001101 case 13
	{
		12, 0, 1, 8,
		12, 1, 14, 8,
		1, 14, 8, 10,
		14, 15, 12, 8,
		14, 15, 8, 10,
		10, 15, 8, 11,
	},
	// 0b00001110 case 14
	{
		0, 15, 3, 11,
		0, 14, 15, 11,
		0, 14, 11, 9,
		0, 13, 14, 9,
		14, 11, 9, 10,
	},
	// 0b00001111 case 15
	{
		12, 13, 14, 10,
		12, 13, 10, 9,
		12, 9, 10, 8,
		12, 14, 15, 8,
		14, 15, 8, 10,
		15, 8, 10, 11,
	},
	// 0b00010000 case 16
	{8, 4, 7, 16},
	// 0b00010001 case 17
	{
		12, 0, 3, 7,
		12, 0, 7, 4,
		12, 4, 7, 16,
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
