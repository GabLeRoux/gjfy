package main

const (
	cssFileName  = "custom.css"
	logoFileName = "logo.png"
)

var (
	favicon = []byte{0x0, 0x0, 0x1, 0x0, 0x1, 0x0, 0x10, 0x10, 0x10, 0x0, 0x1,
		0x0, 0x4, 0x0, 0x28, 0x1, 0x0, 0x0, 0x16, 0x0, 0x0, 0x0, 0x28, 0x0, 0x0,
		0x0, 0x10, 0x0, 0x0, 0x0, 0x20, 0x0, 0x0, 0x0, 0x1, 0x0, 0x4, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x10, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0x0, 0xff, 0xff, 0x0, 0x0, 0xff, 0xff, 0x0, 0x0, 0xff,
		0xff, 0x0, 0x0, 0xfe, 0x7f, 0x0, 0x0, 0xfe, 0x7f, 0x0, 0x0, 0xff, 0xff,
		0x0, 0x0, 0xfe, 0x7f, 0x0, 0x0, 0xfe, 0x7f, 0x0, 0x0, 0xff, 0x3f, 0x0, 0x0,
		0xff, 0x9f, 0x0, 0x0, 0xfd, 0xdf, 0x0, 0x0, 0xfc, 0x9f, 0x0, 0x0, 0xfe,
		0x3f, 0x0, 0x0, 0xff, 0xff, 0x0, 0x0, 0xff, 0xff, 0x0, 0x0, 0xff, 0xff,
		0x0, 0x0}
	gjfyLogoSmall = []byte{0x89, 0x50, 0x4e, 0x47, 0xd, 0xa, 0x1a,
		0xa, 0x0, 0x0, 0x0, 0xd, 0x49, 0x48, 0x44, 0x52, 0x0, 0x0, 0x0, 0x3f,
		0x0, 0x0, 0x0, 0x2a, 0x8, 0x6, 0x0, 0x0, 0x0, 0x4, 0x44, 0x83, 0xf7,
		0x0, 0x0, 0x0, 0x4, 0x73, 0x42, 0x49, 0x54, 0x8, 0x8, 0x8, 0x8, 0x7c,
		0x8, 0x64, 0x88, 0x0, 0x0, 0x0, 0x9, 0x70, 0x48, 0x59, 0x73, 0x0, 0x0,
		0xb, 0x12, 0x0, 0x0, 0xb, 0x12, 0x1, 0xd2, 0xdd, 0x7e, 0xfc, 0x0, 0x0,
		0x0, 0x1c, 0x74, 0x45, 0x58, 0x74, 0x53, 0x6f, 0x66, 0x74, 0x77, 0x61,
		0x72, 0x65, 0x0, 0x41, 0x64, 0x6f, 0x62, 0x65, 0x20, 0x46, 0x69, 0x72,
		0x65, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x20, 0x43, 0x53, 0x36, 0xe8, 0xbc,
		0xb2, 0x8c, 0x0, 0x0, 0x0, 0x18, 0x74, 0x45, 0x58, 0x74, 0x43, 0x72,
		0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x54, 0x69, 0x6d, 0x65, 0x0,
		0x32, 0x37, 0x2e, 0x30, 0x31, 0x2e, 0x32, 0x30, 0x31, 0x37, 0x94, 0xa1,
		0x3d, 0x17, 0x0, 0x0, 0x4, 0xa8, 0x49, 0x44, 0x41, 0x54, 0x68, 0x81,
		0xcd, 0x9a, 0xc1, 0x75, 0xe3, 0x46, 0xc, 0x86, 0xbf, 0xd5, 0xdb, 0xbb,
		0xdc, 0x81, 0x95, 0xa, 0xac, 0xe, 0xa2, 0xad, 0x60, 0x95, 0xa, 0x96,
		0x39, 0xe1, 0xb8, 0xde, 0xa, 0xa2, 0xad, 0x20, 0xda, 0x23, 0x4e, 0x4b,
		0x55, 0x10, 0xb9, 0x82, 0xd0, 0x15, 0x44, 0xae, 0x20, 0x52, 0x7, 0x56,
		0x7, 0x39, 0xc, 0x28, 0xd3, 0x23, 0x70, 0x38, 0xa4, 0x68, 0xc5, 0xff,
		0x7b, 0x7e, 0x16, 0x87, 0x98, 0xe1, 0x60, 0x8, 0x60, 0x7e, 0x60, 0xf8,
		0x81, 0x2b, 0x41, 0x55, 0x6f, 0x80, 0x2d, 0xf0, 0x2b, 0x70, 0x0, 0xa,
		0x11, 0xa9, 0x12, 0xf2, 0x5, 0xb0, 0x4, 0x16, 0xc0, 0x14, 0x38, 0x2,
		0xa5, 0x88, 0xdc, 0x8f, 0x35, 0xa7, 0x8f, 0x63, 0xd, 0x94, 0x81, 0x82,
		0xa0, 0x38, 0xc0, 0x2d, 0x50, 0x2, 0xb3, 0x58, 0x28, 0x5a, 0xa4, 0x26,
		0xa6, 0xc0, 0x57, 0x55, 0xad, 0x44, 0x64, 0x3b, 0xc6, 0x84, 0x26, 0x63,
		0xc, 0x92, 0x89, 0x9b, 0xe8, 0xfa, 0xb6, 0x45, 0xae, 0xe4, 0x5c, 0xf1,
		0xd4, 0x38, 0x83, 0x71, 0x4d, 0xe5, 0x3b, 0xa1, 0xaa, 0xb, 0xe0, 0x73,
		0x42, 0xe4, 0x8, 0x54, 0x63, 0x3d, 0xef, 0x9a, 0x66, 0x9f, 0x83, 0xa5,
		0xd3, 0x76, 0x24, 0xb8, 0xc1, 0x9e, 0xe0, 0xf3, 0xfb, 0xb1, 0x1e, 0xf6,
		0xde, 0x94, 0x9f, 0x3b, 0x6d, 0xb, 0x11, 0xd9, 0xbd, 0xc5, 0xc3, 0xae,
		0x69, 0xf6, 0xfb, 0xe8, 0xfa, 0xd1, 0x91, 0x89, 0xfd, 0xf9, 0xf1, 0xad,
		0x14, 0x87, 0x2b, 0x2a, 0x2f, 0x22, 0x25, 0xf0, 0x9d, 0xa0, 0xf4, 0x86,
		0x10, 0xfd, 0x63, 0xdc, 0x5d, 0x6b, 0x3e, 0x90, 0x61, 0xf6, 0xaa, 0x3a,
		0x27, 0xf8, 0xe2, 0x9c, 0x97, 0x37, 0xb3, 0x27, 0x4, 0x9e, 0xad, 0x88,
		0x3c, 0x47, 0xf2, 0x33, 0x60, 0xe6, 0xed, 0xe1, 0x22, 0xb2, 0xba, 0x64,
		0xb2, 0xce, 0xdc, 0x6e, 0xe2, 0xe7, 0xf7, 0xe9, 0xf3, 0x21, 0x21, 0x54,
		0x0, 0x2b, 0xda, 0xb7, 0x24, 0x8, 0xc1, 0x68, 0xd, 0xac, 0x45, 0xe4,
		0x59, 0x55, 0xef, 0x81, 0x3f, 0xed, 0xde, 0x13, 0xc1, 0x5f, 0x4f, 0x93,
		0xb3, 0x85, 0x59, 0x0, 0xbb, 0xa6, 0x39, 0x5b, 0x7b, 0x49, 0x7a, 0x8b,
		0xfb, 0x51, 0x13, 0x1c, 0x93, 0xdf, 0x12, 0x2c, 0xe5, 0x48, 0x46, 0x5c,
		0x30, 0xfe, 0x50, 0x35, 0xfa, 0xdc, 0x9f, 0x99, 0xbd, 0xaa, 0xce, 0x54,
		0xb5, 0x2, 0x7e, 0x92, 0x56, 0x1c, 0x2, 0xf1, 0xf8, 0x3, 0xa8, 0xcc,
		0x42, 0x56, 0x8d, 0x7b, 0x77, 0x34, 0x2, 0x98, 0xdd, 0xdf, 0xd9, 0xb8,
		0xff, 0xa8, 0xea, 0xba, 0x21, 0xbb, 0x26, 0xad, 0x38, 0x4, 0x82, 0x53,
		0xef, 0x6, 0x2b, 0x5e, 0x5c, 0x64, 0x4a, 0x58, 0xb8, 0x2e, 0x14, 0x51,
		0x9f, 0xf5, 0x2b, 0xe5, 0x1b, 0x13, 0xec, 0x9a, 0x48, 0x8c, 0x3b, 0xc2,
		0xaa, 0x4e, 0x13, 0x32, 0xcb, 0xe8, 0xfe, 0xd7, 0xc6, 0xef, 0x5c, 0xe2,
		0x52, 0x2f, 0x66, 0x2c, 0x7f, 0x67, 0xd6, 0x90, 0x42, 0x4c, 0x8b, 0xa7,
		0x27, 0xe5, 0x4d, 0xf1, 0x8a, 0xb4, 0x2, 0x8f, 0x4, 0x73, 0xf6, 0x90,
		0xea, 0xd7, 0x85, 0x2a, 0x53, 0xae, 0xb4, 0xff, 0x1e, 0xbd, 0x6d, 0xe5,
		0xfc, 0xa6, 0x5b, 0x6c, 0xc5, 0xf, 0xcd, 0x80, 0x57, 0xe2, 0x2b, 0xb0,
		0x21, 0xf8, 0xf4, 0x2b, 0x9f, 0xb2, 0x98, 0x70, 0xcf, 0x8, 0x11, 0x5a,
		0x44, 0x56, 0xaa, 0xfa, 0xcc, 0x4b, 0xbc, 0xa8, 0x71, 0xe0, 0x45, 0xe1,
		0xaa, 0x41, 0x70, 0xb6, 0x4, 0xf7, 0x69, 0x62, 0x49, 0xfb, 0x2, 0x14,
		0x4e, 0xdb, 0x76, 0x2, 0xa0, 0xaa, 0x2b, 0xce, 0x95, 0x38, 0x2, 0xbf,
		0x89, 0x48, 0xe1, 0x5, 0x13, 0x11, 0x29, 0x45, 0x64, 0x4e, 0x58, 0x9c,
		0x8b, 0x21, 0x22, 0x6b, 0xa7, 0x79, 0x2f, 0x22, 0x2b, 0xfb, 0xab, 0x1a,
		0xb2, 0xcf, 0xc0, 0x43, 0x24, 0x7b, 0x6b, 0x6f, 0xd8, 0x43, 0xcc, 0x1c,
		0x8f, 0x22, 0x52, 0xd6, 0x66, 0xef, 0xad, 0xd8, 0x32, 0x27, 0x7b, 0x12,
		0x91, 0xc2, 0x99, 0xc8, 0x35, 0x50, 0x3a, 0x6d, 0x45, 0xdc, 0xd0, 0x62,
		0xf2, 0x5b, 0x80, 0x89, 0x99, 0x6f, 0x6c, 0xee, 0x9b, 0x54, 0xae, 0xed,
		0x60, 0xb4, 0x1c, 0x3b, 0x17, 0xf6, 0x62, 0x8e, 0x51, 0xb3, 0x97, 0x1b,
		0x14, 0x4e, 0xdb, 0x1a, 0x2, 0xc3, 0x5b, 0x38, 0x37, 0x57, 0x3d, 0x27,
		0xb2, 0x67, 0x24, 0xf3, 0xef, 0x89, 0xd8, 0x32, 0x3d, 0xd3, 0x8f, 0x17,
		0xe4, 0x50, 0xbb, 0xf1, 0x84, 0xf3, 0x82, 0xc2, 0x61, 0x60, 0xe6, 0x54,
		0xd, 0xe8, 0x73, 0x29, 0xbc, 0x38, 0x71, 0xb2, 0xc2, 0x16, 0x93, 0x3f,
		0xf5, 0xf1, 0xb8, 0xfd, 0x7e, 0xe0, 0x44, 0x86, 0xf6, 0x1b, 0xc, 0x7b,
		0x83, 0x87, 0xa8, 0xb9, 0xf9, 0xa6, 0xb, 0xa7, 0xdb, 0xc9, 0x5a, 0xde,
		0x55, 0x31, 0x63, 0x20, 0xca, 0xe8, 0x7a, 0xda, 0x60, 0x82, 0xb1, 0xc9,
		0x3f, 0x34, 0xad, 0xda, 0x53, 0xbe, 0x2f, 0xbb, 0xab, 0xb1, 0x18, 0xd8,
		0xef, 0x52, 0x94, 0x4e, 0xdb, 0x32, 0x15, 0xe5, 0x6b, 0x4c, 0x70, 0x7c,
		0xd5, 0x76, 0x80, 0xbe, 0x18, 0xd2, 0xe7, 0x62, 0xd8, 0x9b, 0x8c, 0x59,
		0xa7, 0x47, 0x78, 0x8e, 0x96, 0x56, 0x9f, 0x30, 0xc1, 0x5f, 0xb9, 0x95,
		0x65, 0x41, 0x59, 0xb0, 0xc5, 0xea, 0x4a, 0x82, 0xde, 0x12, 0x71, 0xe0,
		0x9b, 0x2, 0x5f, 0xa2, 0xb6, 0x33, 0xce, 0x32, 0xb1, 0x95, 0x8b, 0xab,
		0x2a, 0xb7, 0xce, 0x80, 0x2e, 0xcc, 0xbc, 0xb2, 0x64, 0xdf, 0x10, 0x39,
		0xa5, 0xec, 0x32, 0x6e, 0x48, 0x31, 0xbc, 0x2f, 0xaa, 0xba, 0x4b, 0x65,
		0x4b, 0x16, 0x58, 0x2a, 0x2e, 0x4b, 0x6a, 0x2e, 0x46, 0xb, 0xdd, 0x6d,
		0xe2, 0xe0, 0x91, 0xb6, 0x8f, 0xd6, 0x79, 0xa7, 0xaa, 0xdf, 0x9, 0xb9,
		0x79, 0x13, 0x77, 0xc0, 0xbf, 0xaa, 0xfa, 0x40, 0x48, 0x75, 0x2b, 0x42,
		0x3a, 0x39, 0x27, 0xf8, 0xf8, 0xff, 0x69, 0xea, 0x31, 0x4a, 0xda, 0xcb,
		0xde, 0xa5, 0xd7, 0x78, 0x8a, 0xf6, 0x56, 0x62, 0x6a, 0x63, 0x69, 0x9f,
		0x9, 0xb, 0xf3, 0x37, 0xf0, 0x97, 0xfd, 0x7e, 0x4f, 0x8a, 0xb7, 0xd1,
		0xdd, 0x1a, 0xa5, 0xd7, 0xf8, 0x6a, 0xab, 0xb3, 0x24, 0xe5, 0xc7, 0xc0,
		0xe7, 0xb7, 0xe5, 0xf9, 0xd7, 0x84, 0xe7, 0xfb, 0x8f, 0x6d, 0x8c, 0xf5,
		0x6c, 0x9f, 0xb7, 0x3a, 0xd9, 0x27, 0xce, 0x99, 0x53, 0xa, 0x1b, 0x7a,
		0xe6, 0x3, 0x2d, 0x88, 0x3, 0x6f, 0xaf, 0xe2, 0x24, 0x3e, 0xc5, 0x2e,
		0xdb, 0x84, 0x5d, 0x86, 0x27, 0x22, 0x95, 0x88, 0xcc, 0x80, 0xdf, 0x49,
		0x7, 0x92, 0x7, 0xe0, 0x93, 0x59, 0x4c, 0x5b, 0x2e, 0xdd, 0x7, 0xcd,
		0x5d, 0xe3, 0x48, 0xff, 0x5, 0x3d, 0xcb, 0xdb, 0x49, 0xec, 0x4, 0xc9,
		0xd2, 0xb5, 0x91, 0x82, 0x12, 0x4e, 0x5b, 0x5a, 0xbd, 0xf7, 0x3f, 0x3b,
		0x5, 0x8e, 0x59, 0x8f, 0x49, 0xb6, 0x3d, 0x6f, 0xab, 0xaa, 0xbf, 0xd8,
		0x58, 0xbb, 0x3e, 0x65, 0x69, 0xdb, 0x95, 0xe2, 0x80, 0x77, 0x56, 0x5a,
		0x6f, 0x22, 0xfb, 0xb8, 0x2a, 0xe3, 0xe4, 0x64, 0xe1, 0xb4, 0xf5, 0x3e,
		0x6d, 0x31, 0xff, 0xdc, 0xf7, 0xed, 0x47, 0x47, 0x12, 0xe3, 0x61, 0x94,
		0xc4, 0xa6, 0x85, 0xe1, 0x3d, 0xf5, 0x3d, 0x50, 0xb8, 0x10, 0x45, 0x74,
		0x7d, 0xe8, 0xaa, 0x44, 0x5d, 0xac, 0x7c, 0x82, 0xe1, 0x95, 0xd1, 0xf5,
		0x68, 0xe7, 0xea, 0xce, 0x1c, 0x96, 0x74, 0x24, 0x31, 0x1e, 0x2e, 0x52,
		0x3e, 0xc1, 0xf0, 0x8e, 0x9c, 0x2b, 0xef, 0x95, 0x98, 0xc6, 0x82, 0x37,
		0x76, 0x27, 0xe5, 0xce, 0xf6, 0xf9, 0x46, 0xc0, 0x9b, 0xd9, 0xdf, 0x92,
		0xf6, 0xb2, 0xf5, 0xca, 0x8e, 0xaf, 0xe6, 0x26, 0x7b, 0xcf, 0xf9, 0x9b,
		0xf1, 0x4e, 0x69, 0x7b, 0xc3, 0x12, 0xb0, 0x38, 0x89, 0x79, 0xca, 0xa9,
		0x46, 0x65, 0x29, 0x6f, 0x3e, 0x1d, 0xd7, 0xc9, 0xdb, 0xb0, 0x11, 0x91,
		0xb5, 0xaa, 0xee, 0x48, 0xd7, 0xf4, 0x47, 0xf9, 0xae, 0x86, 0x44, 0x81,
		0xb2, 0xb, 0xb9, 0x66, 0x9f, 0x5b, 0x9d, 0xdd, 0x88, 0x48, 0x61, 0x9f,
		0x97, 0xa4, 0x14, 0xf7, 0xdc, 0x62, 0x28, 0x3c, 0x93, 0xcf, 0x5a, 0xd8,
		0x5c, 0xe5, 0xbb, 0xa2, 0xf6, 0x11, 0xf8, 0x66, 0x64, 0xa7, 0xb, 0xf5,
		0xa9, 0xea, 0x58, 0x3b, 0x41, 0x1c, 0x48, 0x37, 0xb9, 0x63, 0xe7, 0x2a,
		0x5f, 0xb6, 0xb4, 0x3f, 0x2, 0xdf, 0x8, 0xe7, 0xf1, 0x27, 0x53, 0xb3,
		0xf4, 0xd1, 0xf3, 0xe9, 0xd, 0x30, 0x1f, 0xf9, 0x6b, 0x8b, 0x15, 0x2f,
		0x9, 0x4d, 0x7d, 0x64, 0x9e, 0x85, 0xd6, 0xf3, 0xf9, 0x18, 0xf5, 0x47,
		0x7, 0x76, 0xe9, 0x31, 0x3c, 0xaf, 0xcf, 0x89, 0x15, 0xf6, 0x3c, 0x4,
		0xe9, 0x5, 0xb, 0x7a, 0x73, 0x7a, 0xb2, 0xc2, 0xff, 0x0, 0x1c, 0xc8,
		0xaf, 0xaa, 0xbd, 0x1, 0xf1, 0x94, 0x0, 0x0, 0x0, 0x0, 0x49, 0x45,
		0x4e, 0x44, 0xae, 0x42, 0x60, 0x82}
)
