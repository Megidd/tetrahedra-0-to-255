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
	// 0b00110011 case 51
	{
		12, 1, 3, 5,
		3, 12, 5, 7,
		12, 13, 1, 5,
		12, 13, 5, 17,
		12, 17, 5, 7,
		12, 17, 7, 16,
	},
	// 0b00110100 case 52
	{
		1, 14, 2, 10,
		7, 8, 5, 16,
		8, 9, 5, 16,
		16, 9, 5, 17,
	},
	// 0b00110101 case 53
	{
		1, 14, 2, 10,
		3, 12, 0, 5,
		3, 12, 5, 7,
		0, 9, 5, 17,
		0, 17, 5, 7,
		0, 17, 7, 12,
		12, 17, 7, 16,
	},
	// 0b00110110 case 54
	{
		13, 14, 2, 10,
		0, 13, 2, 10,
		0, 10, 2, 8,
		0, 13, 10, 8,
		13, 10, 8, 5,
		13, 5, 8, 17,
		8, 17, 5, 7,
		8, 17, 7, 16,
	},
	// 0b00110111 case 55
	{
		13, 14, 2, 10,
		13, 2, 3, 10,
		13, 10, 3, 5,
		12, 13, 3, 5,
		12, 13, 5, 17,
		3, 12, 5, 7,
		12, 17, 5, 7,
		12, 17, 7, 16,
	},
	// 0b00111000 case 56
	{
		15, 3, 2, 11,
		8, 5, 7, 16,
		8, 9, 5, 16,
		9, 5, 16, 17,
	},
	// 0b00111001 case 57
	{
		12, 0, 2, 9,
		15, 12, 2, 9,
		2, 15, 9, 11,
		15, 12, 9, 11,
		12, 9, 11, 7,
		12, 9, 7, 16,
		9, 7, 16, 5,
		9, 5, 16, 17,
	},
	// 0b00111010 case 58
	{
		15, 3, 2, 11,
		0, 13, 1, 8,
		13, 1, 8, 7,
		13, 1, 7, 5,
		13, 7, 8, 16,
		13, 7, 16, 5,
		13, 5, 16, 17,
	},
	// 0b00111011 case 59
	{
		12, 13, 1, 17,
		12, 1, 17, 5,
		12, 1, 5, 7,
		12, 17, 5, 7,
		12, 17, 7, 16,
		2, 12, 1, 7,
		15, 12, 2, 7,
		15, 7, 2, 11,
	},
	// 0b00111100 case 60
	{
		8, 5, 7, 16,
		8, 9, 5, 16,
		9, 5, 16, 17,
		14, 3, 1, 10,
		14, 15, 3, 10,
		15, 3, 10, 11,
	},
	// 0b00111101 case 61
	{
		0, 9, 5, 17,
		0, 17, 5, 7,
		0, 17, 7, 16,
		0, 16, 7, 12,
		12, 0, 7, 11,
		14, 0, 1, 10,
		12, 0, 15, 11,
		14, 15, 0, 11,
		14, 11, 0, 10,
	},
	// 0b00111110 case 62
	{
		15, 3, 0, 11,
		14, 15, 0, 11,
		14, 11, 0, 10,
		14, 0, 13, 10,
		0, 13, 10, 5,
		0, 13, 5, 17,
		0, 17, 5, 7,
		0, 17, 7, 8,
		8, 17, 7, 16,
	},
	// 0b00111111 case 63
	{
		12, 14, 15, 11,
		12, 14, 11, 10,
		12, 13, 14, 10,
		11, 12, 10, 7,
		12, 13, 10, 7,
		12, 13, 7, 16,
		13, 10, 7, 5,
		13, 5, 7, 17,
		13, 7, 16, 17,
	},
	// 0b01000000 case 64
	{10, 6, 5, 18},
	// 0b01000001 case 65
	{
		12, 0, 3, 8,
		10, 6, 5, 18,
	},
	// 0b01000010 case 66
	{
		0, 13, 1, 9,
		10, 6, 5, 18,
	},
	// 0b01000011 case 67
	{
		10, 6, 5, 18,
		3, 12, 1, 8,
		12, 13, 1, 8,
		13, 1, 8, 9,
	},
	// 0b01000100 case 68
	{
		1, 14, 2, 5,
		14, 2, 5, 6,
		14, 6, 5, 18,
	},
	// 0b01000101 case 69
	{
		12, 0, 3, 8,
		14, 2, 1, 5,
		14, 2, 5, 6,
		14, 6, 5, 18,
	},
	// 0b01000110 case 70
	{
		0, 13, 2, 6,
		0, 13, 6, 9,
		13, 14, 2, 6,
		13, 14, 6, 9,
		9, 14, 6, 5,
		14, 6, 5, 18,
	},
	// 0b01000111 case 71
	{
		3, 12, 2, 8,
		12, 13, 2, 8,
		13, 2, 8, 9,
		13, 14, 2, 9,
		9, 2, 8, 5,
		9, 14, 2, 5,
		14, 2, 5, 6,
		14, 6, 5, 18,
	},
	// 0b01001000 case 72
	{
		10, 6, 5, 18,
		15, 3, 2, 11,
	},
	// 0b01001001 case 73
	{
		10, 6, 5, 18,
		2, 15, 0, 11,
		15, 12, 0, 11,
		12, 0, 11, 8,
	},
	// 0b01001010 case 74
	{
		0, 13, 1, 9,
		2, 15, 3, 11,
		10, 6, 5, 18,
	},
	// 0b01001011 case 75
	{
		10, 6, 5, 18,
		2, 15, 1, 11,
		15, 13, 1, 11,
		13, 1, 11, 9,
		15, 12, 13, 11,
		11, 12, 13, 9,
		11, 12, 9, 8,
	},
	// 0b01001100 case 76
	{
		14, 15, 3, 11,
		1, 14, 3, 11,
		1, 11, 3, 5,
		1, 14, 11, 5,
		5, 14, 11, 6,
		5, 14, 6, 18,
	},
	// 0b01001101 case 77
	{
		12, 0, 15, 8,
		15, 8, 0, 11,
		14, 0, 1, 5,
		14, 15, 0, 11,
		14, 0, 5, 6,
		14, 11, 0, 6,
		14, 6, 5, 18,
	},
	// 0b01001110 case 78
	{
		0, 13, 1, 9,
		1, 3, 0, 9,
		1, 3, 9, 11,
		1, 15, 3, 11,
		1, 14, 15, 11,
		1, 14, 11, 9,
		9, 14, 11, 5,
		14, 11, 5, 6,
		14, 6, 5, 18,
	},
	// 0b01001111 case 79
	{
		15, 12, 13, 8,
		13, 15, 8, 11,
		13, 11, 8, 9,
		13, 14, 15, 9,
		14, 15, 9, 11,
		9, 14, 11, 5,
		14, 11, 5, 6,
		14, 6, 5, 18,
	},
	// 0b01010000 case 80
	{
		8, 4, 7, 16,
		10, 6, 5, 18,
	},
	// 0b01010001 case 81
	{
		10, 6, 5, 18,
		12, 0, 3, 7,
		12, 0, 7, 4,
		12, 4, 7, 16,
	},
	// 0b01010010 case 82
	{
		13, 1, 0, 9,
		8, 4, 7, 16,
		10, 6, 5, 18,
	},
	// 0b01010011 case 83
	{
		10, 6, 5, 18,
		12, 13, 1, 9,
		3, 12, 1, 9,
		3, 9, 1, 7,
		3, 12, 9, 7,
		12, 9, 7, 4,
		12, 4, 7, 16,
	},
	// 0b01010100 case 84
	{
		8, 4, 7, 16,
		14, 2, 1, 6,
		1, 14, 6, 5,
		5, 14, 6, 18,
	},
	// 0b01010101 case 85
	{
		3, 12, 0, 7,
		12, 0, 7, 4,
		12, 4, 7, 16,
		1, 14, 2, 5,
		14, 2, 5, 6,
		5, 14, 6, 18,
	},
	// 0b01010110 case 86
	{
		8, 4, 7, 16,
		0, 13, 2, 6,
		0, 13, 6, 9,
		13, 14, 2, 6,
		13, 14, 6, 9,
		14, 6, 9, 5,
		14, 6, 5, 18,
	},
	// 0b01010111 case 87
	{
		13, 2, 3, 9,
		13, 3, 12, 9,
		13, 14, 2, 9,
		14, 2, 9, 6,
		3, 12, 9, 7,
		12, 9, 7, 4,
		14, 6, 9, 5,
		12, 4, 7, 16,
		14, 6, 5, 18,
	},
	// 0b01011000 case 88
	{
		8, 4, 7, 16,
		10, 6, 5, 18,
		15, 3, 2, 11,
	},
	// 0b01011001 case 89
	{
		10, 6, 5, 18,
		2, 15, 0, 4,
		2, 15, 4, 11,
		15, 12, 0, 4,
		15, 12, 4, 11,
		11, 12, 4, 7,
		7, 12, 4, 16,
	},
	// 0b01011010 case 90
	{
		13, 1, 0, 9,
		2, 15, 3, 11,
		8, 4, 7, 16,
		10, 6, 5, 18,
	},
	// 0b01011011 case 91
	{
		10, 6, 5, 18,
		15, 1, 2, 11,
		15, 13, 1, 11,
		13, 1, 11, 9,
		15, 12, 13, 11,
		12, 13, 11, 9,
		12, 9, 11, 7,
		12, 9, 7, 4,
		12, 4, 7, 16,
	},
	// 0b01011100 case 92
	{
		8, 4, 7, 16,
		1, 15, 3, 5,
		15, 3, 5, 6,
		15, 3, 6, 11,
		1, 14, 15, 6,
		1, 14, 6, 5,
		14, 6, 5, 18,
	},
	// 0b01011101 case 93
	{
		15, 0, 1, 11,
		15, 12, 0, 11,
		1, 14, 15, 11,
		11, 12, 0, 4,
		11, 1, 14, 5,
		11, 12, 4, 7,
		11, 5, 14, 6,
		14, 6, 5, 18,
		12, 4, 7, 16,
	},
	// 0b01011110 case 94
	{
		8, 4, 7, 16,
		13, 3, 0, 9,
		13, 15, 3, 9,
		15, 3, 9, 11,
		13, 14, 15, 9,
		14, 15, 9, 11,
		14, 11, 9, 5,
		14, 11, 5, 6,
		14, 6, 5, 18,
	},
	// 0b01011111 case 95
	{
		15, 12, 13, 11,
		11, 12, 13, 9,
		11, 12, 9, 7,
		12, 9, 7, 4,
		12, 4, 7, 16,
		13, 14, 15, 9,
		14, 15, 9, 11,
		9, 14, 11, 5,
		5, 14, 11, 6,
		14, 6, 5, 18,
	},
	// 0b01100000 case 96
	{
		9, 10, 4, 17,
		10, 6, 4, 17,
		10, 6, 17, 18,
	},
	// 0b01100001 case 97
	{
		12, 0, 3, 8,
		4, 9, 6, 17,
		9, 10, 6, 17,
		17, 10, 6, 18,
	},
	// 0b01100010 case 98
	{
		0, 13, 1, 6,
		1, 10, 6, 18,
		13, 1, 6, 18,
		0, 13, 6, 4,
		13, 18, 6, 4,
		13, 18, 4, 17,
	},
	// 0b01100011 case 99
	{
		3, 12, 1, 8,
		12, 13, 1, 8,
		8, 13, 1, 17,
		1, 8, 17, 4,
		1, 6, 8, 4,
		1, 4, 17, 6,
		1, 6, 17, 10,
		17, 10, 6, 18,
	},
	// 0b01100100 case 100
	{
		1, 14, 2, 4,
		14, 2, 4, 6,
		1, 4, 9, 6,
		1, 14, 6, 9,
		14, 6, 9, 18,
		9, 18, 6, 4,
		4, 9, 18, 17,
	},
	// 0b01100101 case 101
	{
		3, 12, 0, 8,
		1, 14, 2, 9,
		9, 14, 2, 4,
		14, 2, 4, 6,
		9, 14, 4, 17,
		14, 17, 4, 6,
		14, 6, 17, 18,
	},
	// 0b01100110 case 102
	{
		0, 13, 2, 4,
		13, 14, 2, 4,
		14, 2, 4, 6,
		14, 18, 6, 4,
		14, 18, 4, 13,
		13, 18, 4, 17,
	},
	// 0b01100111 case 103
	{
		12, 2, 3, 8,
		12, 13, 2, 8,
		13, 14, 2, 8,
		14, 2, 8, 4,
		13, 14, 8, 4,
		14, 2, 4, 6,
		13, 14, 4, 17,
		14, 4, 17, 6,
		14, 6, 17, 18,
	},
	// 0b01101000 case 104
	{
		2, 15, 3, 11,
		4, 9, 6, 17,
		9, 10, 6, 17,
		10, 6, 17, 18,
	},
	// 0b01101001 case 105
	{
		2, 15, 0, 11,
		15, 12, 0, 11,
		12, 0, 11, 8,
		4, 9, 6, 17,
		9, 10, 6, 17,
		10, 6, 17, 18,
	},
	// 0b01101010 case 106
	{
		2, 15, 3, 11,
		0, 13, 1, 6,
		0, 13, 6, 4,
		1, 10, 6, 4,
		13, 1, 4, 17,
		10, 6, 4, 18,
		17, 1, 4, 10,
		17, 10, 4, 18,
	},
	// 0b01101011 case 107
	{
		15, 1, 2, 11,
		15, 12, 1, 11,
		12, 1, 11, 8,
		12, 13, 1, 8,
		8, 13, 1, 4,
		13, 1, 4, 17,
		1, 6, 4, 17,
		1, 6, 17, 10,
		10, 6, 17, 18,
	},
	// 0b01101100 case 108
	{
		1, 14, 3, 9,
		14, 15, 3, 9,
		15, 3, 9, 11,
		14, 15, 9, 11,
		14, 11, 9, 6,
		14, 6, 9, 18,
		9, 18, 6, 4,
		4, 9, 18, 17,
	},
	// 0b01101101 case 109
	{
		12, 0, 1, 8,
		15, 12, 1, 8,
		15, 8, 1, 11,
		14, 15, 1, 11,
		1, 14, 11, 6,
		1, 14, 6, 18,
		1, 18, 6, 4,
		1, 18, 4, 17,
		1, 17, 4, 9,
	},
	// 0b01101110 case 110
	{
		14, 15, 3, 11,
		0, 14, 3, 11,
		0, 13, 14, 11,
		0, 13, 11, 6,
		13, 14, 11, 6,
		13, 14, 6, 18,
		0, 13, 6, 4,
		13, 18, 6, 4,
		13, 18, 4, 17,
	},
	// 0b01101111 case 111
	{
		15, 12, 13, 11,
		13, 14, 15, 11,
		12, 13, 11, 8,
		8, 13, 14, 4,
		14, 8, 4, 11,
		14, 11, 4, 6,
		13, 14, 4, 17,
		14, 4, 17, 6,
		14, 6, 17, 18,
	},
	// 0b01110000 case 112
	{
		8, 9, 10, 7,
		9, 10, 7, 6,
		8, 9, 7, 16,
		9, 10, 6, 18,
		9, 7, 16, 17,
		9, 6, 7, 17,
		9, 18, 6, 17,
	},
	// 0b01110001 case 113
	{
		12, 0, 3, 7,
		12, 0, 7, 16,
		0, 9, 10, 7,
		0, 9, 7, 16,
		9, 10, 16, 17,
		10, 7, 16, 17,
		10, 7, 17, 6,
		10, 6, 17, 18,
	},
	// 0b01110010 case 114
	{
		0, 13, 1, 10,
		0, 13, 10, 8,
		13, 10, 8, 6,
		13, 6, 8, 7,
		13, 10, 6, 18,
		13, 7, 8, 16,
		13, 6, 7, 17,
		13, 18, 6, 17,
		13, 7, 16, 17,
	},
	// 0b01110011 case 115
	{
		3, 12, 13, 7,
		12, 13, 7, 16,
		13, 1, 3, 7,
		13, 1, 16, 17,
		1, 7, 16, 6,
		1, 6, 16, 17,
		1, 6, 17, 10,
		10, 6, 17, 18,
	},
	// 0b01110100 case 116
	{
		1, 14, 2, 6,
		1, 14, 6, 18,
		1, 18, 6, 7,
		1, 8, 9, 16,
		1, 7, 8, 16,
		1, 18, 7, 16,
		1, 18, 16, 9,
		9, 18, 16, 17,
	},
	// 0b01110101 case 117
	{
		3, 12, 0, 9,
		3, 12, 9, 7,
		12, 9, 7, 16,
		9, 7, 16, 17,
		9, 6, 7, 17,
		1, 14, 2, 9,
		14, 2, 9, 6,
		14, 6, 9, 17,
		14, 6, 17, 18,
	},
	// 0b01110110 case 118
	{
		13, 14, 2, 6,
		0, 13, 2, 6,
		13, 14, 6, 18,
		0, 13, 18, 17,
		0, 17, 18, 6,
		0, 17, 6, 7,
		0, 17, 7, 16,
		0, 7, 8, 16,
	},
	// 0b01110111 case 119
	{
		12, 2, 3, 7,
		12, 14, 2, 7,
		14, 2, 7, 6,
		12, 13, 14, 7,
		12, 13, 7, 16,
		14, 6, 7, 18,
		14, 18, 7, 16,
		13, 14, 16, 18,
		13, 18, 16, 17,
	},
	// 0b01111000 case 120
	{
		2, 15, 3, 11,
		8, 9, 10, 7,
		9, 10, 7, 6,
		9, 10, 6, 18,
		8, 9, 7, 16,
		9, 18, 6, 17,
		9, 6, 7, 17,
		9, 7, 16, 17,
	},
	// 0b01111001 case 121
	{
		15, 0, 2, 7,
		15, 7, 2, 11,
		15, 12, 0, 7,
		12, 0, 7, 16,
		16, 0, 7, 9,
		16, 9, 7, 17,
		9, 10, 7, 17,
		10, 7, 17, 18,
		10, 7, 18, 6,
	},
	// 0b01111010 case 122
	{
		15, 3, 2, 11,
		0, 13, 1, 10,
		0, 13, 10, 8,
		13, 10, 8, 6,
		8, 13, 6, 7,
		13, 6, 7, 17,
		13, 10, 6, 17,
		13, 7, 8, 17,
		17, 10, 6, 18,
		17, 7, 8, 16,
	},
	// 0b01111011 case 123
	{
		15, 1, 2, 11,
		15, 12, 1, 11,
		12, 13, 1, 11,
		13, 1, 11, 7,
		12, 13, 11, 7,
		12, 13, 7, 16,
		13, 7, 16, 17,
		13, 1, 7, 17,
		1, 7, 17, 6,
		1, 6, 17, 10,
		10, 6, 17, 18,
	},
	// 0b01111100 case 124
	{
		3, 11, 6, 15,
		1, 15, 3, 6,
		1, 14, 15, 6,
		1, 14, 6, 18,
		1, 18, 6, 9,
		9, 18, 6, 17,
		8, 9, 6, 17,
		8, 17, 6, 7,
		8, 17, 7, 16,
	},
	// 0b01111101 case 125
	{
		12, 0, 1, 9,
		12, 1, 14, 9,
		12, 14, 15, 9,
		12, 9, 15, 17,
		14, 15, 9, 17,
		15, 12, 17, 16,
		15, 17, 14, 18,
		15, 16, 17, 7,
		15, 18, 17, 6,
		15, 7, 17, 11,
		15, 17, 6, 11,
		11, 6, 7, 17,
	},
	// 0b01111110 case 126
	{
		15, 3, 0, 11,
		14, 15, 0, 11,
		0, 14, 11, 6,
		0, 13, 14, 6,
		13, 14, 6, 18,
		0, 13, 18, 17,
		0, 18, 6, 17,
		0, 17, 6, 7,
		0, 17, 7, 8,
		8, 17, 7, 16,
	},
	// 0b01111111 case 127
	{
		6, 7, 11, 15,
		12, 15, 13, 7,
		13, 14, 15, 7,
		14, 15, 7, 6,
		13, 14, 7, 6,
		13, 14, 6, 18,
		12, 13, 7, 16,
		13, 18, 6, 17,
		13, 6, 7, 17,
		13, 7, 16, 17,
	},
	// 0b10000000 case 128
	{
		11, 7, 6, 19,
	},
	// 0b10000001 case 129
	{
		3, 12, 0, 8,
		11, 7, 6, 19,
	},
	// 0b10000010 case 130
	{
		0, 13, 1, 9,
		11, 7, 6, 19,
	},
	// 0b10000011 case 131
	{
		11, 7, 6, 19,
		3, 12, 1, 8,
		12, 13, 1, 8,
		13, 1, 8, 9,
	},
	// 0b10000100 case 132
	{
		1, 14, 2, 10,
		11, 7, 6, 19,
	},
	// 0b10000101 case 133
	{
		3, 12, 0, 8,
		1, 14, 2, 10,
		11, 7, 6, 19,
	},
	// 0b10000110 case 134
	{
		0, 13, 2, 9,
		13, 14, 2, 9,
		14, 2, 9, 10,
		11, 7, 6, 19,
	},
	// 0b10000111 case 135
	{
		12, 2, 3, 8,
		12, 14, 2, 8,
		14, 2, 8, 10,
		12, 13, 14, 8,
		13, 14, 8, 10,
		13, 10, 8, 9,
		11, 7, 6, 19,
	},
	// 0b10001000 case 136
	{
		2, 15, 3, 6,
		15, 3, 6, 7,
		15, 7, 6, 19,
	},
	// 0b10001001 case 137
	{
		2, 15, 0, 6,
		15, 12, 0, 6,
		12, 0, 6, 7,
		12, 0, 7, 8,
		15, 12, 6, 7,
		15, 7, 6, 19,
	},
	// 0b10001010 case 138
	{
		0, 13, 1, 9,
		2, 15, 3, 6,
		15, 3, 6, 7,
		15, 7, 6, 19,
	},
	// 0b10001011 case 139
	{
		12, 13, 1, 9,
		12, 9, 1, 8,
		12, 1, 2, 8,
		2, 8, 1, 7,
		2, 7, 1, 6,
		2, 15, 12, 8,
		2, 15, 8, 7,
		2, 15, 7, 6,
		15, 7, 6, 19,
	},
	// 0b10001100 case 140
	{
		15, 3, 1, 7,
		14, 15, 1, 7,
		1, 14, 7, 6,
		1, 14, 6, 10,
		14, 15, 7, 6,
		15, 7, 6, 19,
	},
	// 0b10001101 case 141
	{
		14, 0, 1, 10,
		14, 12, 0, 10,
		12, 0, 10, 8,
		14, 15, 12, 10,
		15, 12, 10, 8,
		10, 15, 8, 6,
		6, 15, 8, 7,
		6, 15, 7, 19,
	},
	// 0b10001110 case 142
	{
		0, 13, 14, 9,
		14, 0, 9, 10,
		14, 15, 0, 10,
		15, 3, 0, 10,
		10, 3, 0, 7,
		10, 15, 3, 7,
		10, 15, 7, 6,
		6, 15, 7, 19,
	},
	// 0b10001111 case 143
	{
		14, 15, 12, 8,
		14, 15, 8, 10,
		10, 15, 8, 6,
		6, 15, 8, 7,
		6, 15, 7, 19,
		12, 13, 14, 10,
		12, 13, 10, 8,
		8, 13, 10, 9,
	},
	// 0b10010000 case 144
	{
		11, 8, 6, 19,
		8, 6, 19, 4,
		8, 4, 19, 16,
	},
	// 0b10010001 case 145
	{
		3, 12, 0, 4,
		3, 4, 0, 6,
		12, 4, 16, 6,
		3, 12, 6, 16,
		3, 16, 6, 19,
		3, 19, 6, 11,
	},
	// 0b10010010 case 146
	{
		0, 13, 1, 9,
		11, 8, 6, 19,
		8, 6, 19, 4,
		8, 4, 19, 16,
	},
	// 0b10010011 case 147
	{
		3, 13, 1, 9,
		3, 12, 13, 9,
		3, 12, 9, 6,
		12, 9, 6, 4,
		3, 12, 4, 16,
		3, 16, 4, 6,
		3, 16, 6, 11,
		11, 16, 6, 19,
	},
	// 0b10010100 case 148
	{
		1, 14, 2, 10,
		11, 8, 6, 19,
		8, 4, 6, 19,
		8, 4, 19, 16,
	},
	// 0b10010101 case 149
	{
		1, 14, 2, 10,
		3, 12, 0, 11,
		12, 0, 11, 6,
		12, 6, 11, 19,
		0, 6, 19, 4,
		12, 0, 19, 4,
		12, 4, 19, 16,
	},
	// 0b10010110 case 150
	{
		13, 14, 0, 9,
		14, 2, 0, 9,
		14, 2, 9, 10,
		11, 8, 6, 19,
		8, 6, 19, 4,
		8, 4, 19, 16,
	},
	// 0b10010111 case 151
	{
		3, 14, 2, 10,
		3, 13, 14, 10,
		3, 13, 10, 9,
		3, 12, 13, 9,
		3, 12, 9, 4,
		3, 12, 4, 16,
		3, 16, 4, 6,
		3, 16, 6, 11,
		11, 16, 6, 19,
	},
	// 0b10011000 case 152
	{
		2, 15, 3, 8,
		2, 15, 8, 4,
		15, 8, 4, 16,
		2, 15, 4, 6,
		15, 16, 4, 6,
		15, 16, 6, 19,
	},
	// 0b10011001 case 153
	{
		12, 0, 2, 4,
		15, 12, 2, 4,
		2, 15, 4, 6,
		12, 4, 15, 16,
		15, 16, 4, 6,
		15, 16, 6, 19,
	},
	// 0b10011010 case 154
	{
		0, 13, 1, 9,
		2, 15, 3, 8,
		2, 15, 8, 4,
		15, 8, 4, 16,
		2, 15, 4, 6,
		15, 16, 4, 6,
		15, 16, 6, 19,
	},
	// 0b10011011 case 155
	{
		2, 13, 1, 9,
		2, 13, 9, 4,
		12, 13, 2, 4,
		15, 12, 2, 4,
		2, 15, 4, 6,
		15, 12, 4, 16,
		15, 16, 4, 6,
		15, 16, 6, 19,
	},
	// 0b10011100 case 156
	{
		14, 3, 1, 8,
		14, 8, 1, 10,
		14, 15, 3, 8,
		14, 15, 8, 10,
		10, 15, 8, 6,
		15, 8, 6, 19,
		6, 19, 8, 4,
		19, 8, 4, 16,
	},
	// 0b10011101 case 157
	{
		14, 0, 1, 10,
		14, 15, 0, 10,
		15, 12, 0, 10,
		12, 0, 10, 6,
		15, 12, 10, 6,
		12, 0, 6, 4,
		15, 12, 6, 19,
		12, 4, 6, 19,
		12, 4, 19, 16,
	},
	// 0b10011110 case 158
	{
		13, 3, 0, 9,
		13, 14, 3, 9,
		14, 3, 9, 10,
		14, 15, 3, 10,
		15, 3, 10, 6,
		15, 3, 6, 19,
		3, 6, 19, 4,
		3, 4, 19, 8,
		8, 4, 19, 16,
	},
	// 0b10011111 case 159
	{
		12, 13, 14, 10,
		12, 13, 10, 9,
		12, 14, 15, 10,
		15, 9, 10, 6,
		15, 12, 9, 6,
		12, 9, 6, 4,
		15, 12, 6, 19,
		12, 6, 19, 4,
		12, 4, 19, 16,
	},
	// 0b10100000 case 160
	{
		9, 5, 4, 17,
		11, 7, 6, 19,
	},
	// 0b10100001 case 161
	{
		3, 12, 0, 8,
		9, 5, 4, 17,
		11, 7, 6, 19,
	},
	// 0b10100010 case 162
	{
		0, 13, 1, 4,
		13, 1, 4, 5,
		4, 13, 5, 17,
		11, 7, 6, 19,
	},
	// 0b10100011 case 163
	{
		3, 13, 1, 5,
		3, 12, 13, 5,
		12, 13, 5, 17,
		3, 17, 5, 4,
		3, 12, 17, 4,
		3, 12, 4, 8,
		11, 7, 6, 19,
	},
	// 0b10100100 case 164
	{
		1, 14, 2, 10,
		11, 7, 6, 19,
		9, 5, 4, 17,
	},
	// 0b10100101 case 165
	{
		3, 12, 0, 8,
		1, 14, 2, 10,
		9, 5, 4, 17,
		11, 7, 6, 19,
	},
	// 0b10100110 case 166
	{
		14, 2, 0, 10,
		13, 14, 0, 10,
		0, 13, 10, 5,
		0, 13, 5, 4,
		13, 5, 4, 17,
		11, 7, 6, 19,
	},
	// 0b10100111 case 167
	{
		11, 7, 6, 19,
		3, 14, 2, 10,
		3, 12, 14, 10,
		3, 12, 10, 8,
		12, 13, 14, 10,
		12, 13, 10, 8,
		8, 13, 10, 5,
		8, 13, 5, 4,
		13, 5, 4, 17,
	},
	// 0b10101000 case 168
	{
		9, 5, 4, 17,
		2, 15, 3, 6,
		15, 3, 6, 7,
		6, 15, 7, 19,
	},
	// 0b10101001 case 169
	{
		9, 5, 4, 17,
		2, 12, 0, 6,
		12, 0, 6, 7,
		12, 0, 7, 8,
		2, 15, 12, 6,
		15, 12, 6, 7,
		6, 15, 7, 19,
	},
	// 0b10101010 case 170
	{
		0, 13, 1, 4,
		13, 1, 4, 5,
		13, 5, 4, 17,
		2, 15, 3, 6,
		15, 3, 6, 7,
		15, 7, 6, 19,
	},
	// 0b10101011 case 171
	{
		2, 12, 1, 8,
		12, 13, 1, 8,
		2, 15, 12, 8,
		8, 13, 1, 5,
		8, 13, 5, 4,
		4, 13, 5, 17,
		2, 15, 8, 6,
		15, 8, 6, 7,
		6, 15, 7, 19,
	},
	// 0b10101100 case 172
	{
		9, 5, 4, 17,
		1, 15, 3, 7,
		1, 14, 15, 7,
		1, 14, 7, 6,
		1, 14, 6, 10,
		14, 15, 7, 6,
		15, 7, 6, 19,
	},
	// 0b10101101 case 173
	{
		9, 5, 4, 17,
		1, 12, 0, 8,
		14, 12, 1, 8,
		14, 15, 12, 8,
		14, 8, 1, 10,
		14, 15, 8, 10,
		10, 15, 8, 6,
		6, 15, 8, 7,
		6, 15, 7, 19,
	},
	// 0b10101110 case 174
	{
		13, 14, 0, 10,
		14, 3, 0, 10,
		14, 15, 3, 10,
		0, 13, 10, 5,
		0, 5, 10, 4,
		0, 13, 5, 4,
		13, 5, 4, 17,
		15, 3, 10, 7,
		15, 7, 10, 6,
		15, 7, 6, 19,
	},
	// 0b10101111 case 175
	{
		12, 13, 14, 10,
		12, 14, 15, 10,
		12, 13, 10, 8,
		10, 15, 12, 8,
		10, 15, 8, 6,
		15, 8, 6, 7,
		15, 7, 6, 19,
		13, 10, 8, 4,
		13, 10, 4, 5,
		13, 5, 4, 17,
	},
	// 0b10110000 case 176
	{
		11, 8, 9, 5,
		11, 8, 5, 6,
		8, 9, 5, 17,
		8, 6, 11, 19,
		8, 17, 5, 16,
		8, 5, 6, 16,
		8, 6, 19, 16,
	},
	// 0b10110001 case 177
	{
		12, 0, 3, 11,
		11, 12, 0, 9,
		11, 12, 9, 6,
		12, 9, 6, 5,
		12, 6, 11, 19,
		12, 9, 5, 17,
		12, 17, 5, 16,
		12, 5, 6, 16,
		12, 6, 19, 16,
	},
	// 0b10110010 case 178
	{
		0, 13, 1, 5,
		0, 13, 5, 17,
		0, 17, 5, 6,
		0, 17, 6, 11,
		11, 17, 6, 19,
		0, 17, 11, 19,
		0, 19, 11, 8,
		0, 17, 19, 8,
		8, 17, 19, 16,
	},
	// 0b10110011 case 179
	{
		3, 13, 1, 5,
		3, 12, 13, 5,
		12, 13, 5, 17,
		3, 12, 17, 16,
		3, 17, 5, 16,
		3, 16, 5, 6,
		3, 16, 6, 11,
		16, 6, 11, 19,
	},
	// 0b10110100 case 180
	{
		11, 8, 9, 5,
		11, 8, 5, 6,
		8, 9, 5, 17,
		8, 17, 5, 16,
		8, 5, 6, 16,
		8, 6, 11, 19,
		8, 6, 19, 16,
		1, 14, 2, 10,
	},
	// 0b10110101 case 181
	{
		1, 14, 2, 10,
		3, 12, 0, 9,
		3, 12, 9, 11,
		11, 12, 9, 5,
		11, 12, 5, 6,
		12, 9, 5, 17,
		11, 12, 6, 19,
		12, 17, 5, 16,
		12, 5, 6, 16,
		12, 6, 19, 16,
	},
	// 0b10110110 case 182
	{
		13, 14, 2, 10,
		0, 13, 2, 10,
		0, 10, 2, 5,
		0, 13, 10, 5,
		0, 13, 5, 17,
		0, 17, 5, 8,
		8, 17, 5, 16,
		11, 8, 5, 16,
		11, 16, 5, 6,
		11, 16, 6, 19,
	},
	// 0b10110111 case 183
	{
		3, 14, 2, 10,
		3, 13, 14, 10,
		3, 12, 13, 10,
		3, 12, 10, 5,
		12, 13, 10, 5,
		12, 13, 5, 17,
		3, 12, 17, 16,
		3, 17, 5, 16,
		3, 16, 5, 6,
		3, 16, 6, 11,
		11, 16, 6, 19,
	},
	// 0b10111000 case 184
	{
		8, 9, 5, 17,
		8, 17, 5, 16,
		2, 15, 3, 8,
		2, 15, 8, 5,
		2, 15, 5, 6,
		15, 8, 5, 6,
		6, 8, 5, 16,
		6, 15, 8, 19,
		6, 19, 8, 16,
	},
	// 0b10111001 case 185
	{
		2, 15, 0, 6,
		15, 12, 0, 6,
		15, 12, 6, 19,
		12, 0, 19, 16,
		0, 6, 19, 16,
		0, 6, 16, 5,
		0, 5, 16, 9,
		16, 9, 5, 17,
	},
	// 0b10111010 case 186
	{
		0, 13, 1, 8,
		13, 1, 8, 5,
		8, 13, 5, 17,
		8, 17, 5, 16,
		8, 16, 5, 6,
		2, 15, 3, 8,
		2, 15, 8, 6,
		15, 8, 6, 16,
		15, 16, 6, 19,
	},
	// 0b10111011 case 187
	{
		2, 13, 1, 5,
		2, 12, 13, 5,
		12, 13, 5, 17,
		12, 17, 5, 16,
		12, 2, 15, 5,
		2, 15, 5, 6,
		15, 12, 5, 16,
		15, 16, 5, 6,
		15, 16, 6, 19,
	},
	// 0b10111100 case 188
	{
		14, 3, 1, 6,
		1, 14, 6, 10,
		14, 15, 3, 6,
		15, 3, 6, 19,
		3, 6, 19, 8,
		8, 6, 19, 16,
		8, 6, 16, 9,
		9, 6, 16, 5,
		9, 5, 16, 17,
	},
	// 0b10111101 case 189
	{
		14, 0, 1, 10,
		14, 12, 0, 10,
		14, 15, 12, 10,
		12, 0, 10, 6,
		10, 15, 12, 6,
		6, 15, 12, 19,
		12, 0, 19, 16,
		0, 6, 19, 16,
		0, 6, 16, 5,
		0, 5, 16, 9,
		16, 9, 5, 17,
	},
	// 0b10111110 case 190
	{
		0, 15, 3, 8,
		0, 14, 15, 8,
		0, 13, 14, 8,
		8, 14, 15, 16,
		14, 15, 16, 19,
		13, 14, 8, 16,
		13, 14, 16, 17,
		14, 16, 17, 5,
		14, 16, 5, 10,
		10, 16, 5, 6,
		14, 19, 16, 10,
		10, 19, 16, 6,
	},
	// 0b10111111 case 191
	{
		0, 15, 12, 8,
		0, 14, 15, 8,
		0, 13, 14, 8,
		8, 14, 15, 16,
		14, 15, 16, 19,
		13, 14, 8, 16,
		13, 14, 16, 17,
		14, 16, 17, 5,
		14, 16, 5, 10,
		10, 16, 5, 6,
		14, 19, 16, 10,
		10, 19, 16, 6,
	},
	// 0b11000000 case 192
	{
		10, 11, 5, 18,
		18, 11, 5, 7,
		11, 7, 18, 19,
	},
	// 0b11000001 case 193
	{
		3, 12, 0, 8,
		5, 10, 11, 18,
		5, 18, 11, 7,
		11, 7, 18, 19,
	},
	// 0b11000010 case 194
	{
		0, 13, 1, 9,
		10, 11, 5, 18,
		11, 7, 5, 18,
		11, 7, 18, 19,
	},
	// 0b11000011 case 195
	{
		12, 1, 3, 8,
		12, 13, 1, 8,
		13, 1, 8, 9,
		10, 11, 5, 18,
		11, 7, 5, 18,
		11, 7, 18, 19,
	},
	// 0b11000100 case 196
	{
		1, 14, 2, 7,
		1, 14, 7, 5,
		14, 2, 5, 18,
		2, 7, 5, 18,
		2, 7, 18, 11,
		11, 7, 18, 19,
	},
	// 0b11000101 case 197
	{
		1, 14, 2, 7,
		1, 14, 7, 5,
		14, 2, 5, 18,
		2, 7, 5, 18,
		2, 7, 18, 11,
		11, 7, 18, 19,
		3, 12, 0, 8,
	},
	// 0b11000110 case 198
	{
		0, 13, 2, 9,
		13, 14, 2, 9,
		14, 2, 9, 5,
		14, 2, 5, 18,
		2, 9, 5, 7,
		2, 5, 18, 7,
		2, 7, 18, 11,
		11, 7, 18, 19,
	},
	// 0b11000111 case 199
	{
		12, 2, 3, 8,
		12, 13, 2, 8,
		8, 13, 2, 9,
		13, 14, 2, 9,
		9, 14, 2, 5,
		14, 2, 5, 18,
		2, 5, 18, 7,
		7, 18, 2, 11,
		18, 11, 7, 19,
	},
	// 0b11001000 case 200
	{
		2, 15, 3, 10,
		15, 3, 10, 5,
		5, 10, 15, 18,
		5, 15, 3, 7,
		5, 18, 15, 7,
		18, 15, 7, 19,
	},
	// 0b11001001 case 201
	{
		12, 0, 2, 8,
		15, 12, 2, 8,
		2, 15, 8, 7,
		2, 15, 7, 19,
		2, 7, 8, 5,
		2, 19, 7, 5,
		2, 19, 5, 10,
		10, 19, 5, 18,
	},
	// 0b11001010 case 202
	{
		0, 13, 1, 9,
		2, 15, 3, 5,
		15, 3, 5, 7,
		2, 15, 7, 19,
		2, 7, 5, 19,
		2, 19, 5, 10,
		10, 19, 5, 18,
	},
	// 0b11001011 case 203
	{
		2, 13, 1, 9,
		2, 12, 13, 9,
		2, 12, 9, 8,
		2, 15, 12, 8,
		2, 15, 8, 7,
		2, 15, 7, 19,
		2, 19, 7, 5,
		2, 19, 5, 10,
		10, 19, 5, 18,
	},
	// 0b11001100 case 204
	{
		1, 14, 3, 5,
		14, 15, 3, 5,
		15, 3, 5, 7,
		14, 15, 5, 18,
		15, 7, 5, 18,
		15, 7, 18, 19,
	},
	// 0b11001101 case 205
	{
		1, 12, 0, 8,
		1, 15, 12, 8,
		1, 15, 8, 7,
		1, 14, 15, 7,
		14, 15, 7, 19,
		14, 7, 1, 5,
		14, 7, 5, 19,
		14, 19, 5, 18,
	},
	// 0b11001110 case 206
	{
		13, 3, 0, 9,
		13, 14, 3, 9,
		14, 3, 9, 5,
		14, 15, 3, 5,
		14, 15, 5, 18,
		15, 3, 5, 7,
		15, 7, 5, 18,
		15, 7, 18, 19,
	},
	// 0b11001111 case 207
	{
		15, 12, 13, 8,
		13, 15, 8, 9,
		14, 15, 13, 9,
		15, 8, 9, 7,
		14, 15, 9, 7,
		14, 7, 9, 5,
		15, 7, 14, 19,
		14, 19, 7, 5,
		14, 19, 5, 18,
	},
	// 0b11010000 case 208
	{
		10, 11, 8, 4,
		10, 11, 4, 5,
		11, 8, 4, 16,
		11, 5, 10, 18,
		11, 16, 4, 19,
		11, 4, 5, 19,
		11, 5, 18, 19,
	},
	// 0b11010001 case 209
	{
		3, 12, 0, 5,
		12, 0, 5, 4,
		12, 4, 5, 16,
		3, 12, 5, 16,
		3, 16, 5, 11,
		11, 16, 5, 19,
		11, 19, 5, 10,
		10, 19, 5, 18,
	},
	// 0b11010010 case 210
	{
		0, 13, 1, 9,
		10, 11, 8, 4,
		10, 11, 4, 5,
		10, 11, 5, 18,
		11, 8, 4, 16,
		11, 5, 18, 19,
		11, 4, 5, 19,
		11, 16, 4, 19,
	},
	// 0b11010011 case 211
	{
		3, 13, 1, 4,
		13, 1, 4, 9,
		3, 12, 13, 4,
		3, 12, 4, 16,
		3, 16, 4, 11,
		11, 16, 4, 19,
		10, 11, 4, 19,
		10, 19, 4, 5,
		10, 19, 5, 18,
	},
	// 0b11010100 case 212
	{
		1, 14, 2, 5,
		14, 2, 5, 18,
		2, 5, 18, 4,
		2, 4, 18, 8,
		2, 8, 18, 11,
		18, 8, 4, 16,
		18, 8, 16, 11,
		18, 11, 16, 19,
	},
	// 0b11010101 case 213
	{
		12, 0, 3, 11,
		12, 0, 11, 4,
		12, 4, 11, 16,
		4, 11, 16, 19,
		11, 19, 4, 5,
		11, 19, 5, 18,
		1, 14, 2, 11,
		1, 14, 11, 5,
		14, 11, 5, 18,
	},
	// 0b11010110 case 214
	{
		13, 2, 0, 5,
		0, 13, 5, 9,
		13, 14, 2, 5,
		5, 14, 2, 18,
		5, 18, 2, 11,
		5, 18, 11, 19,
		5, 11, 8, 19,
		8, 5, 19, 4,
		8, 4, 19, 16,
	},
	// 0b11010111 case 215
	{
		3, 2, 11, 19,
		3, 14, 2, 19,
		3, 12, 14, 19,
		12, 13, 14, 19,
		12, 13, 19, 16,
		13, 14, 16, 9,
		9, 14, 16, 4,
		14, 19, 16, 4,
		14, 4, 9, 5,
		14, 19, 4, 5,
		14, 19, 5, 18,
	},
	// 0b11011000 case 216
	{
		2, 15, 3, 10,
		15, 3, 10, 8,
		10, 15, 8, 5,
		15, 8, 5, 4,
		15, 5, 10, 18,
		15, 8, 4, 16,
		15, 4, 5, 19,
		15, 16, 4, 19,
		15, 5, 18, 19,
	},
	// 0b11011001 case 217
	{
		2, 12, 0, 4,
		2, 15, 12, 4,
		15, 12, 4, 16,
		2, 15, 16, 19,
		2, 19, 16, 4,
		2, 19, 4, 5,
		2, 19, 5, 10,
		10, 19, 5, 18,
	},
	// 0b11011010 case 218
	{
		0, 13, 1, 9,
		2, 15, 3, 10,
		15, 3, 10, 8,
		10, 15, 8, 5,
		5, 15, 8, 4,
		15, 5, 10, 18,
		15, 4, 5, 19,
		15, 8, 4, 16,
		15, 5, 18, 19,
		15, 16, 4, 19,
	},
	// 0b11011011 case 219
	{
		13, 1, 2, 9,
		12, 13, 2, 9,
		12, 2, 15, 9,
		15, 9, 2, 4,
		15, 12, 9, 4,
		15, 12, 4, 16,
		15, 16, 2, 19,
		2, 16, 4, 19,
		2, 19, 4, 5,
		2, 19, 5, 10,
		10, 19, 5, 18,
	},
	// 0b11011100 case 220
	{
		1, 14, 3, 5,
		14, 15, 3, 5,
		14, 15, 5, 18,
		15, 3, 5, 8,
		15, 8, 5, 4,
		15, 8, 4, 16,
		15, 16, 4, 19,
		15, 4, 5, 19,
		15, 5, 18, 19,
	},
	// 0b11011101 case 221
	{
		14, 0, 1, 5,
		14, 15, 0, 5,
		14, 15, 5, 18,
		15, 12, 0, 5,
		12, 0, 5, 4,
		15, 12, 18, 19,
		12, 4, 5, 16,
		12, 5, 18, 16,
		12, 18, 19, 16,
	},
	// 0b11011110 case 222
	{
		13, 3, 0, 9,
		13, 14, 3, 9,
		14, 3, 9, 5,
		14, 3, 5, 18,
		14, 15, 3, 18,
		15, 3, 18, 19,
		3, 18, 19, 5,
		3, 5, 19, 4,
		3, 4, 19, 8,
		8, 4, 19, 16,
	},
	// 0b11011111 case 223
	{
		14, 12, 13, 9,
		14, 15, 12, 9,
		14, 15, 9, 5,
		14, 15, 5, 18,
		15, 9, 5, 4,
		15, 12, 9, 4,
		15, 12, 4, 16,
		15, 16, 4, 19,
		15, 4, 5, 19,
		15, 5, 18, 19,
	},
	// 0b11100000 case 224
	{
		9, 10, 11, 7,
		9, 10, 7, 4,
		10, 11, 7, 19,
		10, 4, 9, 17,
		10, 19, 7, 18,
		10, 7, 4, 18,
		10, 4, 17, 18,
	},
	// 0b11100001 case 225
	{
		3, 12, 0, 8,
		9, 10, 11, 7,
		9, 10, 7, 4,
		10, 11, 7, 19,
		10, 4, 9, 17,
		10, 19, 7, 18,
		10, 7, 4, 18,
		10, 4, 17, 18,
	},
	// 0b11100010 case 226
	{
		0, 13, 1, 4,
		13, 1, 4, 17,
		1, 4, 17, 10,
		10, 4, 17, 18,
		1, 4, 10, 11,
		10, 11, 4, 18,
		18, 11, 4, 7,
		18, 11, 7, 19,
	},
	// 0b11100011 case 227
	{
		12, 1, 3, 4,
		12, 4, 3, 8,
		12, 13, 1, 4,
		13, 1, 4, 17,
		1, 4, 17, 10,
		10, 4, 17, 18,
		10, 11, 4, 18,
		11, 4, 18, 7,
		11, 7, 18, 19,
	},
	// 0b11100100 case 228
	{
		1, 14, 2, 9,
		14, 2, 9, 11,
		14, 11, 9, 4,
		14, 11, 4, 7,
		14, 11, 7, 19,
		14, 4, 9, 17,
		14, 19, 7, 18,
		14, 7, 4, 18,
		14, 4, 17, 18,
	},
	// 0b11100101 case 229
	{
		1, 14, 2, 9,
		14, 2, 9, 11,
		14, 11, 9, 4,
		14, 11, 4, 7,
		14, 11, 7, 19,
		14, 4, 9, 17,
		14, 19, 7, 18,
		14, 7, 4, 18,
		14, 4, 17, 18,
		3, 12, 0, 8,
	},
	// 0b11100110 case 230
	{
		0, 13, 2, 4,
		13, 14, 2, 4,
		13, 14, 4, 17,
		14, 2, 17, 18,
		2, 4, 17, 18,
		2, 4, 18, 7,
		2, 7, 18, 11,
		11, 7, 18, 19,
	},
	// 0b11100111 case 231
	{
		12, 2, 3, 8,
		12, 13, 2, 8,
		13, 2, 8, 4,
		13, 14, 2, 4,
		13, 14, 4, 17,
		14, 2, 17, 18,
		2, 4, 17, 18,
		2, 4, 18, 7,
		2, 7, 18, 11,
		11, 7, 18, 19,
	},
	// 0b11101000 case 232
	{
		2, 15, 3, 7,
		2, 15, 7, 19,
		2, 19, 7, 4,
		2, 19, 4, 9,
		9, 19, 4, 17,
		2, 19, 9, 10,
		9, 10, 19, 17,
		10, 19, 17, 18,
	},
	// 0b11101001 case 233
	{
		2, 12, 0, 7,
		12, 0, 7, 8,
		2, 15, 12, 7,
		2, 15, 7, 19,
		2, 19, 7, 10,
		10, 19, 7, 18,
		9, 10, 7, 18,
		9, 18, 7, 4,
		9, 18, 4, 17,
	},
	// 0b11101010 case 234
	{
		0, 13, 1, 10,
		0, 13, 10, 17,
		0, 17, 10, 4,
		10, 4, 17, 18,
		10, 7, 4, 18,
		2, 15, 3, 10,
		15, 3, 10, 7,
		10, 15, 7, 18,
		15, 7, 18, 19,
	},
	// 0b11101011 case 235
	{
		13, 1, 2, 10,
		13, 2, 15, 10,
		13, 15, 12, 10,
		12, 13, 10, 18,
		10, 15, 12, 18,
		15, 12, 18, 19,
		12, 13, 18, 17,
		12, 17, 18, 8,
		12, 18, 19, 8,
		8, 17, 18, 4,
		8, 4, 18, 7,
		19, 8, 18, 7,
	},
	// 0b11101100 case 236
	{
		1, 15, 3, 7,
		1, 14, 15, 7,
		14, 15, 7, 19,
		1, 14, 19, 18,
		1, 18, 19, 7,
		1, 18, 7, 4,
		1, 18, 4, 9,
		9, 18, 4, 17,
	},
	// 0b11101101 case 237
	{
		1, 12, 0, 8,
		1, 15, 12, 8,
		1, 14, 15, 8,
		1, 14, 8, 7,
		14, 15, 8, 7,
		14, 15, 7, 19,
		1, 14, 19, 18,
		1, 18, 19, 7,
		1, 18, 7, 4,
		4, 1, 18, 9,
		9, 18, 4, 17,
	},
	// 0b11101110 case 238
	{
		13, 3, 0, 4,
		13, 15, 3, 4,
		15, 3, 4, 7,
		13, 14, 15, 7,
		13, 14, 7, 4,
		14, 15, 7, 19,
		13, 14, 4, 17,
		14, 19, 7, 18,
		14, 7, 4, 18,
		14, 4, 17, 18,
	},
	// 0b11101111 case 239
	{
		14, 15, 12, 18,
		12, 13, 14, 18,
		15, 12, 18, 19,
		12, 13, 18, 17,
		12, 17, 18, 8,
		12, 18, 19, 8,
		8, 17, 18, 4,
		18, 8, 4, 7,
		19, 8, 18, 7,
	},
	// 0b11110000 case 240
	{
		9, 10, 11, 18,
		8, 9, 11, 18,
		8, 9, 18, 17,
		18, 11, 8, 19,
		8, 17, 18, 16,
		8, 18, 19, 16,
	},
}
