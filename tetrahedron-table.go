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
	// 0b00010010 case 18
	{
		8, 4, 7, 16,
		13, 1, 0, 9,
	},
	// 0b00010011 case 19
	{
		3, 12, 1, 7,
		12, 13, 1, 7,
		12, 13, 7, 16,
		13, 1, 7, 9,
		13, 7, 16, 9,
		16, 9, 7, 4,
	},
	// 0b00010100 case 20
	{
		8, 4, 7, 16,
		1, 14, 2, 10,
	},
	// 0b00010101 case 21
	{
		1, 14, 2, 10,
		3, 12, 0, 7,
		12, 0, 7, 4,
		12, 4, 7, 16,
	},
	// 0b00010110 case 22
	{
		8, 4, 7, 16,
		0, 13, 2, 9,
		13, 14, 2, 9,
		14, 2, 9, 10,
	},
	// 0b00010111 case 23
	{
		3, 12, 2, 7,
		12, 13, 2, 7,
		13, 2, 7, 4,
		13, 2, 4, 9,
		12, 13, 7, 16,
		16, 13, 7, 4,
		13, 14, 2, 9,
		14, 2, 9, 10,
	},
	// 0b00011000 case 24
	{
		8, 4, 7, 16,
		2, 15, 3, 11,
	},
	// 0b00011001 case 25
	{
		12, 0, 2, 4,
		12, 2, 15, 4,
		2, 15, 4, 11,
		12, 4, 15, 11,
		12, 4, 11, 7,
		12, 4, 7, 16,
	},
	// 0b00011010 case 26
	{
		0, 13, 1, 9,
		2, 15, 3, 11,
		8, 4, 7, 16,
	},
	// 0b00011011 case 27
	{
		12, 2, 15, 11,
		12, 1, 2, 11,
		12, 13, 1, 11,
		13, 1, 11, 9,
		12, 13, 11, 9,
		12, 9, 11, 4,
		12, 4, 11, 16,
		11, 16, 4, 7,
	},
	// 0b00011100 case 28
	{
		8, 4, 7, 16,
		1, 14, 3, 10,
		14, 15, 3, 10,
		15, 3, 10, 11,
	},
	// 0b00011101 case 29
	{
		12, 0, 1, 4,
		3, 12, 1, 4,
		15, 3, 1, 4,
		15, 4, 1, 7,
		15, 7, 1, 11,
		15, 12, 4, 7,
		12, 4, 7, 16,
		1, 14, 15, 11,
		1, 14, 11, 10,
	},
	// 0b00011110 case 30
	{
		8, 4, 7, 16,
		13, 14, 0, 9,
		14, 3, 0, 9,
		14, 15, 3, 9,
		15, 3, 9, 11,
		14, 15, 9, 11,
		14, 11, 9, 10,
	},
	// 0b00011111 case 31
	{
		13, 14, 15, 11,
		13, 14, 11, 9,
		14, 11, 9, 10,
		12, 13, 15, 9,
		12, 9, 15, 11,
		12, 9, 11, 7,
		12, 9, 7, 16,
		16, 9, 7, 4,
	},
	// 0b00100000 case 32
	{9, 5, 4, 17},
	// 0b00100001 case 33
	{
		12, 0, 3, 8,
		9, 5, 4, 17,
	},
	// 0b00100010 case 34
	{
		0, 13, 1, 5,
		0, 13, 5, 4,
		4, 13, 5, 17,
	},
	// 0b00100011 case 35
	{
		13, 1, 3, 5,
		13, 3, 12, 5,
		12, 5, 3, 4,
		12, 4, 3, 8,
		12, 13, 5, 4,
		13, 5, 4, 17,
	},
	// 0b00100100 case 36
	{
		1, 14, 2, 10,
		9, 5, 4, 17,
	},
	// 0b00100101 case 37
	{
		3, 12, 0, 8,
		1, 14, 2, 10,
		4, 9, 5, 17,
	},
	// 0b00100110 case 38
	{
		0, 14, 2, 4,
		14, 2, 4, 5,
		14, 2, 5, 10,
		0, 13, 14, 4,
		13, 14, 4, 5,
		13, 5, 4, 17,
	},
	// 0b00100111 case 39
	{
		13, 14, 2, 10,
		12, 13, 2, 10,
		12, 2, 3, 10,
		12, 10, 3, 8,
		12, 13, 10, 8,
		13, 10, 8, 5,
		13, 5, 8, 4,
		13, 5, 4, 17,
	},
	// 0b00101000 case 40
	{
		9, 5, 4, 17,
		2, 15, 3, 11,
	},
	// 0b00101001 case 41
	{
		9, 5, 4, 17,
		2, 15, 0, 11,
		15, 12, 0, 11,
		12, 0, 11, 8,
	},
	// 0b00101010 case 42
	{
		2, 15, 3, 11,
		0, 13, 1, 4,
		4, 13, 1, 5,
		4, 13, 5, 17,
	},
	// 0b00101011 case 43
	{
		2, 15, 12, 11,
		11, 12, 2, 8,
		12, 1, 2, 8,
		12, 13, 1, 8,
		8, 1, 2, 5,
		13, 1, 8, 5,
		8, 13, 5, 4,
		13, 5, 4, 17,
	},
	// 0b00101100 case 44
	{
		9, 5, 4, 17,
		1, 14, 3, 10,
		14, 15, 3, 10,
		15, 3, 10, 11,
	},
	// 0b00101101 case 45
	{
		9, 5, 4, 17,
		1, 14, 0, 10,
		14, 12, 0, 10,
		12, 0, 10, 8,
		14, 15, 12, 10,
		15, 12, 10, 8,
		15, 8, 10, 11,
	},
	// 0b00101110 case 46
	{
		14, 15, 3, 11,
		14, 3, 0, 11,
		0, 14, 11, 5,
		14, 11, 5, 10,
		0, 13, 14, 5,
		0, 13, 5, 4,
		13, 5, 4, 17,
	},
	// 0b00101111 case 47
	{
		14, 15, 12, 8,
		14, 15, 8, 10,
		15, 8, 10, 11,
		12, 13, 14, 10,
		12, 13, 10, 8,
		13, 10, 8, 5,
		13, 5, 8, 4,
		13, 5, 4, 17,
	},
	// 0b00110000 case 48
	{
		8, 5, 7, 16,
		8, 9, 5, 16,
		9, 5, 16, 17,
	},
	// 0b00110001 case 49
	{
		3, 12, 0, 5,
		3, 12, 5, 7,
		12, 0, 7, 16,
		0, 5, 7, 16,
		0, 5, 16, 9,
		9, 5, 16, 17,
	},
	// 0b00110010 case 50
	{
		0, 13, 1, 7,
		13, 1, 7, 5,
		0, 13, 5, 17,
		0, 5, 7, 17,
		0, 17, 7, 8,
		8, 17, 7, 16,
	},
}
