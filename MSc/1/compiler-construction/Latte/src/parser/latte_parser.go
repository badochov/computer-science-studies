// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Latte

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type LatteParser struct {
	*antlr.BaseParser
}

var latteParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func latteParserInit() {
	staticData := &latteParserStaticData
	staticData.literalNames = []string{
		"", "'('", "')'", "'class'", "'extends'", "'{'", "'}'", "','", "'&'",
		"';'", "'['", "']'", "'='", "'.'", "'++'", "'--'", "'return'", "'if'",
		"'else'", "'while'", "'for'", "':'", "'[]'", "'int'", "'string'", "'boolean'",
		"'void'", "'new'", "'&&'", "'||'", "'true'", "'false'", "'null'", "'!'",
		"'-'", "'+'", "'*'", "'/'", "'%'", "'<'", "'<='", "'>'", "'>='", "'=='",
		"'!='",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "COMMENT", "MULTICOMMENT",
		"INT", "ID", "WS", "STR",
	}
	staticData.ruleNames = []string{
		"program", "topDef", "arg", "ref", "block", "stmt", "type_", "type_single",
		"classElement", "item", "expr", "unOp", "addOp", "mulOp", "relOp",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 50, 306, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 5, 0,
		32, 8, 0, 10, 0, 12, 0, 35, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3,
		1, 43, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 52, 8, 1,
		1, 1, 1, 1, 5, 1, 56, 8, 1, 10, 1, 12, 1, 59, 9, 1, 1, 1, 3, 1, 62, 8,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 72, 8, 2, 10,
		2, 12, 2, 75, 9, 2, 1, 3, 3, 3, 78, 8, 3, 1, 4, 1, 4, 5, 4, 82, 8, 4, 10,
		4, 12, 4, 85, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 5, 5, 110, 8, 5, 10, 5, 12, 5, 113, 9, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 3, 5, 166, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 172, 8, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 3, 7, 179, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 185, 8,
		8, 10, 8, 12, 8, 188, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 196,
		8, 8, 1, 8, 1, 8, 1, 8, 3, 8, 201, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9,
		207, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 5, 10, 230, 8, 10, 10, 10, 12, 10, 233, 9, 10, 3, 10, 235, 8,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 243, 8, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 250, 8, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 282, 8, 10, 10, 10, 12, 10,
		285, 9, 10, 3, 10, 287, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 293,
		8, 10, 10, 10, 12, 10, 296, 9, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 14, 0, 1, 20, 15, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 0, 4, 1, 0, 33, 34, 1, 0, 34, 35, 1, 0, 36, 38, 1,
		0, 39, 44, 345, 0, 33, 1, 0, 0, 0, 2, 61, 1, 0, 0, 0, 4, 63, 1, 0, 0, 0,
		6, 77, 1, 0, 0, 0, 8, 79, 1, 0, 0, 0, 10, 165, 1, 0, 0, 0, 12, 171, 1,
		0, 0, 0, 14, 178, 1, 0, 0, 0, 16, 200, 1, 0, 0, 0, 18, 206, 1, 0, 0, 0,
		20, 249, 1, 0, 0, 0, 22, 297, 1, 0, 0, 0, 24, 299, 1, 0, 0, 0, 26, 301,
		1, 0, 0, 0, 28, 303, 1, 0, 0, 0, 30, 32, 3, 2, 1, 0, 31, 30, 1, 0, 0, 0,
		32, 35, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 36, 1,
		0, 0, 0, 35, 33, 1, 0, 0, 0, 36, 37, 5, 0, 0, 1, 37, 1, 1, 0, 0, 0, 38,
		39, 3, 12, 6, 0, 39, 40, 5, 48, 0, 0, 40, 42, 5, 1, 0, 0, 41, 43, 3, 4,
		2, 0, 42, 41, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 45,
		5, 2, 0, 0, 45, 46, 3, 8, 4, 0, 46, 62, 1, 0, 0, 0, 47, 48, 5, 3, 0, 0,
		48, 51, 5, 48, 0, 0, 49, 50, 5, 4, 0, 0, 50, 52, 5, 48, 0, 0, 51, 49, 1,
		0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 57, 5, 5, 0, 0, 54,
		56, 3, 16, 8, 0, 55, 54, 1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57, 55, 1, 0,
		0, 0, 57, 58, 1, 0, 0, 0, 58, 60, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 60, 62,
		5, 6, 0, 0, 61, 38, 1, 0, 0, 0, 61, 47, 1, 0, 0, 0, 62, 3, 1, 0, 0, 0,
		63, 64, 3, 12, 6, 0, 64, 65, 3, 6, 3, 0, 65, 73, 5, 48, 0, 0, 66, 67, 5,
		7, 0, 0, 67, 68, 3, 12, 6, 0, 68, 69, 3, 6, 3, 0, 69, 70, 5, 48, 0, 0,
		70, 72, 1, 0, 0, 0, 71, 66, 1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73, 71, 1,
		0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 5, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 76,
		78, 5, 8, 0, 0, 77, 76, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 7, 1, 0, 0,
		0, 79, 83, 5, 5, 0, 0, 80, 82, 3, 10, 5, 0, 81, 80, 1, 0, 0, 0, 82, 85,
		1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 86, 1, 0, 0, 0,
		85, 83, 1, 0, 0, 0, 86, 87, 5, 6, 0, 0, 87, 9, 1, 0, 0, 0, 88, 166, 5,
		9, 0, 0, 89, 90, 3, 20, 10, 0, 90, 91, 5, 10, 0, 0, 91, 92, 3, 20, 10,
		0, 92, 93, 5, 11, 0, 0, 93, 94, 5, 12, 0, 0, 94, 95, 3, 20, 10, 0, 95,
		96, 5, 9, 0, 0, 96, 166, 1, 0, 0, 0, 97, 98, 3, 20, 10, 0, 98, 99, 5, 13,
		0, 0, 99, 100, 5, 48, 0, 0, 100, 101, 5, 12, 0, 0, 101, 102, 3, 20, 10,
		0, 102, 103, 5, 9, 0, 0, 103, 166, 1, 0, 0, 0, 104, 166, 3, 8, 4, 0, 105,
		106, 3, 12, 6, 0, 106, 111, 3, 18, 9, 0, 107, 108, 5, 7, 0, 0, 108, 110,
		3, 18, 9, 0, 109, 107, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 109, 1, 0,
		0, 0, 111, 112, 1, 0, 0, 0, 112, 114, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0,
		114, 115, 5, 9, 0, 0, 115, 166, 1, 0, 0, 0, 116, 117, 5, 48, 0, 0, 117,
		118, 5, 12, 0, 0, 118, 119, 3, 20, 10, 0, 119, 120, 5, 9, 0, 0, 120, 166,
		1, 0, 0, 0, 121, 122, 5, 48, 0, 0, 122, 123, 5, 14, 0, 0, 123, 166, 5,
		9, 0, 0, 124, 125, 5, 48, 0, 0, 125, 126, 5, 15, 0, 0, 126, 166, 5, 9,
		0, 0, 127, 128, 5, 16, 0, 0, 128, 129, 3, 20, 10, 0, 129, 130, 5, 9, 0,
		0, 130, 166, 1, 0, 0, 0, 131, 132, 5, 16, 0, 0, 132, 166, 5, 9, 0, 0, 133,
		134, 5, 17, 0, 0, 134, 135, 5, 1, 0, 0, 135, 136, 3, 20, 10, 0, 136, 137,
		5, 2, 0, 0, 137, 138, 3, 10, 5, 0, 138, 166, 1, 0, 0, 0, 139, 140, 5, 17,
		0, 0, 140, 141, 5, 1, 0, 0, 141, 142, 3, 20, 10, 0, 142, 143, 5, 2, 0,
		0, 143, 144, 3, 10, 5, 0, 144, 145, 5, 18, 0, 0, 145, 146, 3, 10, 5, 0,
		146, 166, 1, 0, 0, 0, 147, 148, 5, 19, 0, 0, 148, 149, 5, 1, 0, 0, 149,
		150, 3, 20, 10, 0, 150, 151, 5, 2, 0, 0, 151, 152, 3, 10, 5, 0, 152, 166,
		1, 0, 0, 0, 153, 154, 5, 20, 0, 0, 154, 155, 5, 1, 0, 0, 155, 156, 3, 12,
		6, 0, 156, 157, 5, 48, 0, 0, 157, 158, 5, 21, 0, 0, 158, 159, 3, 20, 10,
		0, 159, 160, 5, 2, 0, 0, 160, 161, 3, 10, 5, 0, 161, 166, 1, 0, 0, 0, 162,
		163, 3, 20, 10, 0, 163, 164, 5, 9, 0, 0, 164, 166, 1, 0, 0, 0, 165, 88,
		1, 0, 0, 0, 165, 89, 1, 0, 0, 0, 165, 97, 1, 0, 0, 0, 165, 104, 1, 0, 0,
		0, 165, 105, 1, 0, 0, 0, 165, 116, 1, 0, 0, 0, 165, 121, 1, 0, 0, 0, 165,
		124, 1, 0, 0, 0, 165, 127, 1, 0, 0, 0, 165, 131, 1, 0, 0, 0, 165, 133,
		1, 0, 0, 0, 165, 139, 1, 0, 0, 0, 165, 147, 1, 0, 0, 0, 165, 153, 1, 0,
		0, 0, 165, 162, 1, 0, 0, 0, 166, 11, 1, 0, 0, 0, 167, 168, 3, 14, 7, 0,
		168, 169, 5, 22, 0, 0, 169, 172, 1, 0, 0, 0, 170, 172, 3, 14, 7, 0, 171,
		167, 1, 0, 0, 0, 171, 170, 1, 0, 0, 0, 172, 13, 1, 0, 0, 0, 173, 179, 5,
		23, 0, 0, 174, 179, 5, 24, 0, 0, 175, 179, 5, 25, 0, 0, 176, 179, 5, 26,
		0, 0, 177, 179, 5, 48, 0, 0, 178, 173, 1, 0, 0, 0, 178, 174, 1, 0, 0, 0,
		178, 175, 1, 0, 0, 0, 178, 176, 1, 0, 0, 0, 178, 177, 1, 0, 0, 0, 179,
		15, 1, 0, 0, 0, 180, 181, 3, 12, 6, 0, 181, 186, 5, 48, 0, 0, 182, 183,
		5, 7, 0, 0, 183, 185, 5, 48, 0, 0, 184, 182, 1, 0, 0, 0, 185, 188, 1, 0,
		0, 0, 186, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 189, 1, 0, 0, 0,
		188, 186, 1, 0, 0, 0, 189, 190, 5, 9, 0, 0, 190, 201, 1, 0, 0, 0, 191,
		192, 3, 12, 6, 0, 192, 193, 5, 48, 0, 0, 193, 195, 5, 1, 0, 0, 194, 196,
		3, 4, 2, 0, 195, 194, 1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 197, 1, 0,
		0, 0, 197, 198, 5, 2, 0, 0, 198, 199, 3, 8, 4, 0, 199, 201, 1, 0, 0, 0,
		200, 180, 1, 0, 0, 0, 200, 191, 1, 0, 0, 0, 201, 17, 1, 0, 0, 0, 202, 207,
		5, 48, 0, 0, 203, 204, 5, 48, 0, 0, 204, 205, 5, 12, 0, 0, 205, 207, 3,
		20, 10, 0, 206, 202, 1, 0, 0, 0, 206, 203, 1, 0, 0, 0, 207, 19, 1, 0, 0,
		0, 208, 209, 6, 10, -1, 0, 209, 210, 5, 27, 0, 0, 210, 211, 3, 14, 7, 0,
		211, 212, 5, 10, 0, 0, 212, 213, 3, 20, 10, 0, 213, 214, 5, 11, 0, 0, 214,
		250, 1, 0, 0, 0, 215, 216, 5, 27, 0, 0, 216, 250, 3, 14, 7, 0, 217, 218,
		3, 22, 11, 0, 218, 219, 3, 20, 10, 14, 219, 250, 1, 0, 0, 0, 220, 250,
		5, 48, 0, 0, 221, 250, 5, 47, 0, 0, 222, 250, 5, 30, 0, 0, 223, 250, 5,
		31, 0, 0, 224, 225, 5, 48, 0, 0, 225, 234, 5, 1, 0, 0, 226, 231, 3, 20,
		10, 0, 227, 228, 5, 7, 0, 0, 228, 230, 3, 20, 10, 0, 229, 227, 1, 0, 0,
		0, 230, 233, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232,
		235, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 234, 226, 1, 0, 0, 0, 234, 235,
		1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 250, 5, 2, 0, 0, 237, 250, 5, 50,
		0, 0, 238, 239, 5, 1, 0, 0, 239, 240, 3, 12, 6, 0, 240, 241, 5, 2, 0, 0,
		241, 243, 1, 0, 0, 0, 242, 238, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 243,
		244, 1, 0, 0, 0, 244, 250, 5, 32, 0, 0, 245, 246, 5, 1, 0, 0, 246, 247,
		3, 20, 10, 0, 247, 248, 5, 2, 0, 0, 248, 250, 1, 0, 0, 0, 249, 208, 1,
		0, 0, 0, 249, 215, 1, 0, 0, 0, 249, 217, 1, 0, 0, 0, 249, 220, 1, 0, 0,
		0, 249, 221, 1, 0, 0, 0, 249, 222, 1, 0, 0, 0, 249, 223, 1, 0, 0, 0, 249,
		224, 1, 0, 0, 0, 249, 237, 1, 0, 0, 0, 249, 242, 1, 0, 0, 0, 249, 245,
		1, 0, 0, 0, 250, 294, 1, 0, 0, 0, 251, 252, 10, 13, 0, 0, 252, 253, 3,
		26, 13, 0, 253, 254, 3, 20, 10, 14, 254, 293, 1, 0, 0, 0, 255, 256, 10,
		12, 0, 0, 256, 257, 3, 24, 12, 0, 257, 258, 3, 20, 10, 13, 258, 293, 1,
		0, 0, 0, 259, 260, 10, 11, 0, 0, 260, 261, 3, 28, 14, 0, 261, 262, 3, 20,
		10, 12, 262, 293, 1, 0, 0, 0, 263, 264, 10, 10, 0, 0, 264, 265, 5, 28,
		0, 0, 265, 293, 3, 20, 10, 10, 266, 267, 10, 9, 0, 0, 267, 268, 5, 29,
		0, 0, 268, 293, 3, 20, 10, 9, 269, 270, 10, 19, 0, 0, 270, 271, 5, 10,
		0, 0, 271, 272, 3, 20, 10, 0, 272, 273, 5, 11, 0, 0, 273, 293, 1, 0, 0,
		0, 274, 275, 10, 18, 0, 0, 275, 276, 5, 13, 0, 0, 276, 277, 5, 48, 0, 0,
		277, 286, 5, 1, 0, 0, 278, 283, 3, 20, 10, 0, 279, 280, 5, 7, 0, 0, 280,
		282, 3, 20, 10, 0, 281, 279, 1, 0, 0, 0, 282, 285, 1, 0, 0, 0, 283, 281,
		1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284, 287, 1, 0, 0, 0, 285, 283, 1, 0,
		0, 0, 286, 278, 1, 0, 0, 0, 286, 287, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0,
		288, 293, 5, 2, 0, 0, 289, 290, 10, 17, 0, 0, 290, 291, 5, 13, 0, 0, 291,
		293, 5, 48, 0, 0, 292, 251, 1, 0, 0, 0, 292, 255, 1, 0, 0, 0, 292, 259,
		1, 0, 0, 0, 292, 263, 1, 0, 0, 0, 292, 266, 1, 0, 0, 0, 292, 269, 1, 0,
		0, 0, 292, 274, 1, 0, 0, 0, 292, 289, 1, 0, 0, 0, 293, 296, 1, 0, 0, 0,
		294, 292, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 21, 1, 0, 0, 0, 296, 294,
		1, 0, 0, 0, 297, 298, 7, 0, 0, 0, 298, 23, 1, 0, 0, 0, 299, 300, 7, 1,
		0, 0, 300, 25, 1, 0, 0, 0, 301, 302, 7, 2, 0, 0, 302, 27, 1, 0, 0, 0, 303,
		304, 7, 3, 0, 0, 304, 29, 1, 0, 0, 0, 24, 33, 42, 51, 57, 61, 73, 77, 83,
		111, 165, 171, 178, 186, 195, 200, 206, 231, 234, 242, 249, 283, 286, 292,
		294,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// LatteParserInit initializes any static state used to implement LatteParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLatteParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LatteParserInit() {
	staticData := &latteParserStaticData
	staticData.once.Do(latteParserInit)
}

// NewLatteParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLatteParser(input antlr.TokenStream) *LatteParser {
	LatteParserInit()
	this := new(LatteParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &latteParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// LatteParser tokens.
const (
	LatteParserEOF          = antlr.TokenEOF
	LatteParserT__0         = 1
	LatteParserT__1         = 2
	LatteParserT__2         = 3
	LatteParserT__3         = 4
	LatteParserT__4         = 5
	LatteParserT__5         = 6
	LatteParserT__6         = 7
	LatteParserT__7         = 8
	LatteParserT__8         = 9
	LatteParserT__9         = 10
	LatteParserT__10        = 11
	LatteParserT__11        = 12
	LatteParserT__12        = 13
	LatteParserT__13        = 14
	LatteParserT__14        = 15
	LatteParserT__15        = 16
	LatteParserT__16        = 17
	LatteParserT__17        = 18
	LatteParserT__18        = 19
	LatteParserT__19        = 20
	LatteParserT__20        = 21
	LatteParserT__21        = 22
	LatteParserT__22        = 23
	LatteParserT__23        = 24
	LatteParserT__24        = 25
	LatteParserT__25        = 26
	LatteParserT__26        = 27
	LatteParserT__27        = 28
	LatteParserT__28        = 29
	LatteParserT__29        = 30
	LatteParserT__30        = 31
	LatteParserT__31        = 32
	LatteParserT__32        = 33
	LatteParserT__33        = 34
	LatteParserT__34        = 35
	LatteParserT__35        = 36
	LatteParserT__36        = 37
	LatteParserT__37        = 38
	LatteParserT__38        = 39
	LatteParserT__39        = 40
	LatteParserT__40        = 41
	LatteParserT__41        = 42
	LatteParserT__42        = 43
	LatteParserT__43        = 44
	LatteParserCOMMENT      = 45
	LatteParserMULTICOMMENT = 46
	LatteParserINT          = 47
	LatteParserID           = 48
	LatteParserWS           = 49
	LatteParserSTR          = 50
)

// LatteParser rules.
const (
	LatteParserRULE_program      = 0
	LatteParserRULE_topDef       = 1
	LatteParserRULE_arg          = 2
	LatteParserRULE_ref          = 3
	LatteParserRULE_block        = 4
	LatteParserRULE_stmt         = 5
	LatteParserRULE_type_        = 6
	LatteParserRULE_type_single  = 7
	LatteParserRULE_classElement = 8
	LatteParserRULE_item         = 9
	LatteParserRULE_expr         = 10
	LatteParserRULE_unOp         = 11
	LatteParserRULE_addOp        = 12
	LatteParserRULE_mulOp        = 13
	LatteParserRULE_relOp        = 14
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(LatteParserEOF, 0)
}

func (s *ProgramContext) AllTopDef() []ITopDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITopDefContext); ok {
			len++
		}
	}

	tst := make([]ITopDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITopDefContext); ok {
			tst[i] = t.(ITopDefContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) TopDef(i int) ITopDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopDefContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LatteParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&281475102539784) != 0 {
		{
			p.SetState(30)
			p.TopDef()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(36)
		p.Match(LatteParserEOF)
	}

	return localctx
}

// ITopDefContext is an interface to support dynamic dispatch.
type ITopDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopDefContext differentiates from other interfaces.
	IsTopDefContext()
}

type TopDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopDefContext() *TopDefContext {
	var p = new(TopDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_topDef
	return p
}

func (*TopDefContext) IsTopDefContext() {}

func NewTopDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopDefContext {
	var p = new(TopDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_topDef

	return p
}

func (s *TopDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TopDefContext) CopyFrom(ctx *TopDefContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TopDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DefClassContext struct {
	*TopDefContext
}

func NewDefClassContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefClassContext {
	var p = new(DefClassContext)

	p.TopDefContext = NewEmptyTopDefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TopDefContext))

	return p
}

func (s *DefClassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefClassContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(LatteParserID)
}

func (s *DefClassContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(LatteParserID, i)
}

func (s *DefClassContext) AllClassElement() []IClassElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IClassElementContext); ok {
			len++
		}
	}

	tst := make([]IClassElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IClassElementContext); ok {
			tst[i] = t.(IClassElementContext)
			i++
		}
	}

	return tst
}

func (s *DefClassContext) ClassElement(i int) IClassElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassElementContext)
}

func (s *DefClassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterDefClass(s)
	}
}

func (s *DefClassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitDefClass(s)
	}
}

func (s *DefClassContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitDefClass(s)

	default:
		return t.VisitChildren(s)
	}
}

type DefFuncContext struct {
	*TopDefContext
}

func NewDefFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefFuncContext {
	var p = new(DefFuncContext)

	p.TopDefContext = NewEmptyTopDefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TopDefContext))

	return p
}

func (s *DefFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefFuncContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *DefFuncContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *DefFuncContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DefFuncContext) Arg() IArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgContext)
}

func (s *DefFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterDefFunc(s)
	}
}

func (s *DefFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitDefFunc(s)
	}
}

func (s *DefFuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitDefFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) TopDef() (localctx ITopDefContext) {
	this := p
	_ = this

	localctx = NewTopDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LatteParserRULE_topDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(61)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LatteParserT__22, LatteParserT__23, LatteParserT__24, LatteParserT__25, LatteParserID:
		localctx = NewDefFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.Type_()
		}
		{
			p.SetState(39)
			p.Match(LatteParserID)
		}
		{
			p.SetState(40)
			p.Match(LatteParserT__0)
		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&281475102539776) != 0 {
			{
				p.SetState(41)
				p.Arg()
			}

		}
		{
			p.SetState(44)
			p.Match(LatteParserT__1)
		}
		{
			p.SetState(45)
			p.Block()
		}

	case LatteParserT__2:
		localctx = NewDefClassContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)
			p.Match(LatteParserT__2)
		}
		{
			p.SetState(48)
			p.Match(LatteParserID)
		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LatteParserT__3 {
			{
				p.SetState(49)
				p.Match(LatteParserT__3)
			}
			{
				p.SetState(50)
				p.Match(LatteParserID)
			}

		}
		{
			p.SetState(53)
			p.Match(LatteParserT__4)
		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&281475102539776) != 0 {
			{
				p.SetState(54)
				p.ClassElement()
			}

			p.SetState(59)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(60)
			p.Match(LatteParserT__5)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArgContext is an interface to support dynamic dispatch.
type IArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgContext differentiates from other interfaces.
	IsArgContext()
}

type ArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgContext() *ArgContext {
	var p = new(ArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_arg
	return p
}

func (*ArgContext) IsArgContext() {}

func NewArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgContext {
	var p = new(ArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_arg

	return p
}

func (s *ArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgContext) AllType_() []IType_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType_Context); ok {
			len++
		}
	}

	tst := make([]IType_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType_Context); ok {
			tst[i] = t.(IType_Context)
			i++
		}
	}

	return tst
}

func (s *ArgContext) Type_(i int) IType_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ArgContext) AllRef() []IRefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRefContext); ok {
			len++
		}
	}

	tst := make([]IRefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRefContext); ok {
			tst[i] = t.(IRefContext)
			i++
		}
	}

	return tst
}

func (s *ArgContext) Ref(i int) IRefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRefContext)
}

func (s *ArgContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(LatteParserID)
}

func (s *ArgContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(LatteParserID, i)
}

func (s *ArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterArg(s)
	}
}

func (s *ArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitArg(s)
	}
}

func (s *ArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Arg() (localctx IArgContext) {
	this := p
	_ = this

	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LatteParserRULE_arg)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Type_()
	}
	{
		p.SetState(64)
		p.Ref()
	}
	{
		p.SetState(65)
		p.Match(LatteParserID)
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LatteParserT__6 {
		{
			p.SetState(66)
			p.Match(LatteParserT__6)
		}
		{
			p.SetState(67)
			p.Type_()
		}
		{
			p.SetState(68)
			p.Ref()
		}
		{
			p.SetState(69)
			p.Match(LatteParserID)
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRefContext is an interface to support dynamic dispatch.
type IRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRefContext differentiates from other interfaces.
	IsRefContext()
}

type RefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRefContext() *RefContext {
	var p = new(RefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_ref
	return p
}

func (*RefContext) IsRefContext() {}

func NewRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RefContext {
	var p = new(RefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_ref

	return p
}

func (s *RefContext) GetParser() antlr.Parser { return s.parser }
func (s *RefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterRef(s)
	}
}

func (s *RefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitRef(s)
	}
}

func (s *RefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitRef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Ref() (localctx IRefContext) {
	this := p
	_ = this

	localctx = NewRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LatteParserRULE_ref)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LatteParserT__7 {
		{
			p.SetState(76)
			p.Match(LatteParserT__7)
		}

	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LatteParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(LatteParserT__4)
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1548145919722018) != 0 {
		{
			p.SetState(80)
			p.Stmt()
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(86)
		p.Match(LatteParserT__5)
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) CopyFrom(ctx *StmtContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssContext struct {
	*StmtContext
}

func NewAssContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssContext {
	var p = new(AssContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *AssContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *AssContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterAss(s)
	}
}

func (s *AssContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitAss(s)
	}
}

func (s *AssContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitAss(s)

	default:
		return t.VisitChildren(s)
	}
}

type RetContext struct {
	*StmtContext
}

func NewRetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RetContext {
	var p = new(RetContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *RetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterRet(s)
	}
}

func (s *RetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitRet(s)
	}
}

func (s *RetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitRet(s)

	default:
		return t.VisitChildren(s)
	}
}

type IndexAssContext struct {
	*StmtContext
}

func NewIndexAssContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexAssContext {
	var p = new(IndexAssContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *IndexAssContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexAssContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *IndexAssContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IndexAssContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterIndexAss(s)
	}
}

func (s *IndexAssContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitIndexAss(s)
	}
}

func (s *IndexAssContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitIndexAss(s)

	default:
		return t.VisitChildren(s)
	}
}

type CondContext struct {
	*StmtContext
}

func NewCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CondContext {
	var p = new(CondContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *CondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CondContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *CondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterCond(s)
	}
}

func (s *CondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitCond(s)
	}
}

func (s *CondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitCond(s)

	default:
		return t.VisitChildren(s)
	}
}

type CondElseContext struct {
	*StmtContext
}

func NewCondElseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CondElseContext {
	var p = new(CondElseContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *CondElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondElseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CondElseContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *CondElseContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *CondElseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterCondElse(s)
	}
}

func (s *CondElseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitCondElse(s)
	}
}

func (s *CondElseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitCondElse(s)

	default:
		return t.VisitChildren(s)
	}
}

type VRetContext struct {
	*StmtContext
}

func NewVRetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VRetContext {
	var p = new(VRetContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *VRetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VRetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterVRet(s)
	}
}

func (s *VRetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitVRet(s)
	}
}

func (s *VRetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitVRet(s)

	default:
		return t.VisitChildren(s)
	}
}

type BlockStmtContext struct {
	*StmtContext
}

func NewBlockStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockStmtContext {
	var p = new(BlockStmtContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *BlockStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *BlockStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterBlockStmt(s)
	}
}

func (s *BlockStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitBlockStmt(s)
	}
}

func (s *BlockStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitBlockStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclContext struct {
	*StmtContext
}

func NewDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclContext {
	var p = new(DeclContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *DeclContext) AllItem() []IItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IItemContext); ok {
			len++
		}
	}

	tst := make([]IItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IItemContext); ok {
			tst[i] = t.(IItemContext)
			i++
		}
	}

	return tst
}

func (s *DeclContext) Item(i int) IItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IItemContext)
}

func (s *DeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterDecl(s)
	}
}

func (s *DeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitDecl(s)
	}
}

func (s *DeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type WhileContext struct {
	*StmtContext
}

func NewWhileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileContext {
	var p = new(WhileContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *WhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *WhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterWhile(s)
	}
}

func (s *WhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitWhile(s)
	}
}

func (s *WhileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitWhile(s)

	default:
		return t.VisitChildren(s)
	}
}

type SExpContext struct {
	*StmtContext
}

func NewSExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SExpContext {
	var p = new(SExpContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *SExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SExpContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterSExp(s)
	}
}

func (s *SExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitSExp(s)
	}
}

func (s *SExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitSExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type AttrAssContext struct {
	*StmtContext
}

func NewAttrAssContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AttrAssContext {
	var p = new(AttrAssContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *AttrAssContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrAssContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AttrAssContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AttrAssContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *AttrAssContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterAttrAss(s)
	}
}

func (s *AttrAssContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitAttrAss(s)
	}
}

func (s *AttrAssContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitAttrAss(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForEachContext struct {
	*StmtContext
}

func NewForEachContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForEachContext {
	var p = new(ForEachContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *ForEachContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForEachContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ForEachContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *ForEachContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForEachContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ForEachContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterForEach(s)
	}
}

func (s *ForEachContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitForEach(s)
	}
}

func (s *ForEachContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitForEach(s)

	default:
		return t.VisitChildren(s)
	}
}

type DecrContext struct {
	*StmtContext
}

func NewDecrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecrContext {
	var p = new(DecrContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *DecrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecrContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *DecrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterDecr(s)
	}
}

func (s *DecrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitDecr(s)
	}
}

func (s *DecrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitDecr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EmptyContext struct {
	*StmtContext
}

func NewEmptyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptyContext {
	var p = new(EmptyContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *EmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEmpty(s)
	}
}

func (s *EmptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEmpty(s)
	}
}

func (s *EmptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEmpty(s)

	default:
		return t.VisitChildren(s)
	}
}

type IncrContext struct {
	*StmtContext
}

func NewIncrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IncrContext {
	var p = new(IncrContext)

	p.StmtContext = NewEmptyStmtContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StmtContext))

	return p
}

func (s *IncrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncrContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *IncrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterIncr(s)
	}
}

func (s *IncrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitIncr(s)
	}
}

func (s *IncrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitIncr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Stmt() (localctx IStmtContext) {
	this := p
	_ = this

	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LatteParserRULE_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEmptyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Match(LatteParserT__8)
		}

	case 2:
		localctx = NewIndexAssContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.expr(0)
		}
		{
			p.SetState(90)
			p.Match(LatteParserT__9)
		}
		{
			p.SetState(91)
			p.expr(0)
		}
		{
			p.SetState(92)
			p.Match(LatteParserT__10)
		}
		{
			p.SetState(93)
			p.Match(LatteParserT__11)
		}
		{
			p.SetState(94)
			p.expr(0)
		}
		{
			p.SetState(95)
			p.Match(LatteParserT__8)
		}

	case 3:
		localctx = NewAttrAssContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(97)
			p.expr(0)
		}
		{
			p.SetState(98)
			p.Match(LatteParserT__12)
		}
		{
			p.SetState(99)
			p.Match(LatteParserID)
		}
		{
			p.SetState(100)
			p.Match(LatteParserT__11)
		}
		{
			p.SetState(101)
			p.expr(0)
		}
		{
			p.SetState(102)
			p.Match(LatteParserT__8)
		}

	case 4:
		localctx = NewBlockStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(104)
			p.Block()
		}

	case 5:
		localctx = NewDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(105)
			p.Type_()
		}
		{
			p.SetState(106)
			p.Item()
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LatteParserT__6 {
			{
				p.SetState(107)
				p.Match(LatteParserT__6)
			}
			{
				p.SetState(108)
				p.Item()
			}

			p.SetState(113)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(114)
			p.Match(LatteParserT__8)
		}

	case 6:
		localctx = NewAssContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(116)
			p.Match(LatteParserID)
		}
		{
			p.SetState(117)
			p.Match(LatteParserT__11)
		}
		{
			p.SetState(118)
			p.expr(0)
		}
		{
			p.SetState(119)
			p.Match(LatteParserT__8)
		}

	case 7:
		localctx = NewIncrContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(121)
			p.Match(LatteParserID)
		}
		{
			p.SetState(122)
			p.Match(LatteParserT__13)
		}
		{
			p.SetState(123)
			p.Match(LatteParserT__8)
		}

	case 8:
		localctx = NewDecrContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(124)
			p.Match(LatteParserID)
		}
		{
			p.SetState(125)
			p.Match(LatteParserT__14)
		}
		{
			p.SetState(126)
			p.Match(LatteParserT__8)
		}

	case 9:
		localctx = NewRetContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(127)
			p.Match(LatteParserT__15)
		}
		{
			p.SetState(128)
			p.expr(0)
		}
		{
			p.SetState(129)
			p.Match(LatteParserT__8)
		}

	case 10:
		localctx = NewVRetContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(131)
			p.Match(LatteParserT__15)
		}
		{
			p.SetState(132)
			p.Match(LatteParserT__8)
		}

	case 11:
		localctx = NewCondContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(133)
			p.Match(LatteParserT__16)
		}
		{
			p.SetState(134)
			p.Match(LatteParserT__0)
		}
		{
			p.SetState(135)
			p.expr(0)
		}
		{
			p.SetState(136)
			p.Match(LatteParserT__1)
		}
		{
			p.SetState(137)
			p.Stmt()
		}

	case 12:
		localctx = NewCondElseContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(139)
			p.Match(LatteParserT__16)
		}
		{
			p.SetState(140)
			p.Match(LatteParserT__0)
		}
		{
			p.SetState(141)
			p.expr(0)
		}
		{
			p.SetState(142)
			p.Match(LatteParserT__1)
		}
		{
			p.SetState(143)
			p.Stmt()
		}
		{
			p.SetState(144)
			p.Match(LatteParserT__17)
		}
		{
			p.SetState(145)
			p.Stmt()
		}

	case 13:
		localctx = NewWhileContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(147)
			p.Match(LatteParserT__18)
		}
		{
			p.SetState(148)
			p.Match(LatteParserT__0)
		}
		{
			p.SetState(149)
			p.expr(0)
		}
		{
			p.SetState(150)
			p.Match(LatteParserT__1)
		}
		{
			p.SetState(151)
			p.Stmt()
		}

	case 14:
		localctx = NewForEachContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(153)
			p.Match(LatteParserT__19)
		}
		{
			p.SetState(154)
			p.Match(LatteParserT__0)
		}
		{
			p.SetState(155)
			p.Type_()
		}
		{
			p.SetState(156)
			p.Match(LatteParserID)
		}
		{
			p.SetState(157)
			p.Match(LatteParserT__20)
		}
		{
			p.SetState(158)
			p.expr(0)
		}
		{
			p.SetState(159)
			p.Match(LatteParserT__1)
		}
		{
			p.SetState(160)
			p.Stmt()
		}

	case 15:
		localctx = NewSExpContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(162)
			p.expr(0)
		}
		{
			p.SetState(163)
			p.Match(LatteParserT__8)
		}

	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) CopyFrom(ctx *Type_Context) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeArrayContext struct {
	*Type_Context
}

func NewTypeArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeArrayContext {
	var p = new(TypeArrayContext)

	p.Type_Context = NewEmptyType_Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type_Context))

	return p
}

func (s *TypeArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeArrayContext) Type_single() IType_singleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_singleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_singleContext)
}

func (s *TypeArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterTypeArray(s)
	}
}

func (s *TypeArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitTypeArray(s)
	}
}

func (s *TypeArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitTypeArray(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeSingleContext struct {
	*Type_Context
}

func NewTypeSingleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeSingleContext {
	var p = new(TypeSingleContext)

	p.Type_Context = NewEmptyType_Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Type_Context))

	return p
}

func (s *TypeSingleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSingleContext) Type_single() IType_singleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_singleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_singleContext)
}

func (s *TypeSingleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterTypeSingle(s)
	}
}

func (s *TypeSingleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitTypeSingle(s)
	}
}

func (s *TypeSingleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitTypeSingle(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Type_() (localctx IType_Context) {
	this := p
	_ = this

	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LatteParserRULE_type_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeArrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)
			p.Type_single()
		}
		{
			p.SetState(168)
			p.Match(LatteParserT__21)
		}

	case 2:
		localctx = NewTypeSingleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(170)
			p.Type_single()
		}

	}

	return localctx
}

// IType_singleContext is an interface to support dynamic dispatch.
type IType_singleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_singleContext differentiates from other interfaces.
	IsType_singleContext()
}

type Type_singleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_singleContext() *Type_singleContext {
	var p = new(Type_singleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_type_single
	return p
}

func (*Type_singleContext) IsType_singleContext() {}

func NewType_singleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_singleContext {
	var p = new(Type_singleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_type_single

	return p
}

func (s *Type_singleContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_singleContext) CopyFrom(ctx *Type_singleContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Type_singleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_singleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StrContext struct {
	*Type_singleContext
}

func NewStrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StrContext {
	var p = new(StrContext)

	p.Type_singleContext = NewEmptyType_singleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Type_singleContext))

	return p
}

func (s *StrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterStr(s)
	}
}

func (s *StrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitStr(s)
	}
}

func (s *StrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitStr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolContext struct {
	*Type_singleContext
}

func NewBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolContext {
	var p = new(BoolContext)

	p.Type_singleContext = NewEmptyType_singleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Type_singleContext))

	return p
}

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterBool(s)
	}
}

func (s *BoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitBool(s)
	}
}

func (s *BoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitBool(s)

	default:
		return t.VisitChildren(s)
	}
}

type ClassNameContext struct {
	*Type_singleContext
}

func NewClassNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ClassNameContext {
	var p = new(ClassNameContext)

	p.Type_singleContext = NewEmptyType_singleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Type_singleContext))

	return p
}

func (s *ClassNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassNameContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *ClassNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterClassName(s)
	}
}

func (s *ClassNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitClassName(s)
	}
}

func (s *ClassNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitClassName(s)

	default:
		return t.VisitChildren(s)
	}
}

type VoidContext struct {
	*Type_singleContext
}

func NewVoidContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VoidContext {
	var p = new(VoidContext)

	p.Type_singleContext = NewEmptyType_singleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Type_singleContext))

	return p
}

func (s *VoidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VoidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterVoid(s)
	}
}

func (s *VoidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitVoid(s)
	}
}

func (s *VoidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitVoid(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntContext struct {
	*Type_singleContext
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	p.Type_singleContext = NewEmptyType_singleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Type_singleContext))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterInt(s)
	}
}

func (s *IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitInt(s)
	}
}

func (s *IntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitInt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Type_single() (localctx IType_singleContext) {
	this := p
	_ = this

	localctx = NewType_singleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LatteParserRULE_type_single)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(178)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LatteParserT__22:
		localctx = NewIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)
			p.Match(LatteParserT__22)
		}

	case LatteParserT__23:
		localctx = NewStrContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
			p.Match(LatteParserT__23)
		}

	case LatteParserT__24:
		localctx = NewBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(175)
			p.Match(LatteParserT__24)
		}

	case LatteParserT__25:
		localctx = NewVoidContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(176)
			p.Match(LatteParserT__25)
		}

	case LatteParserID:
		localctx = NewClassNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(177)
			p.Match(LatteParserID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IClassElementContext is an interface to support dynamic dispatch.
type IClassElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassElementContext differentiates from other interfaces.
	IsClassElementContext()
}

type ClassElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassElementContext() *ClassElementContext {
	var p = new(ClassElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_classElement
	return p
}

func (*ClassElementContext) IsClassElementContext() {}

func NewClassElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassElementContext {
	var p = new(ClassElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_classElement

	return p
}

func (s *ClassElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassElementContext) CopyFrom(ctx *ClassElementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ClassElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ClassAttributeContext struct {
	*ClassElementContext
}

func NewClassAttributeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ClassAttributeContext {
	var p = new(ClassAttributeContext)

	p.ClassElementContext = NewEmptyClassElementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ClassElementContext))

	return p
}

func (s *ClassAttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassAttributeContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ClassAttributeContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(LatteParserID)
}

func (s *ClassAttributeContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(LatteParserID, i)
}

func (s *ClassAttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterClassAttribute(s)
	}
}

func (s *ClassAttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitClassAttribute(s)
	}
}

func (s *ClassAttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitClassAttribute(s)

	default:
		return t.VisitChildren(s)
	}
}

type ClassMethodContext struct {
	*ClassElementContext
}

func NewClassMethodContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ClassMethodContext {
	var p = new(ClassMethodContext)

	p.ClassElementContext = NewEmptyClassElementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ClassElementContext))

	return p
}

func (s *ClassMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassMethodContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ClassMethodContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *ClassMethodContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ClassMethodContext) Arg() IArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgContext)
}

func (s *ClassMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterClassMethod(s)
	}
}

func (s *ClassMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitClassMethod(s)
	}
}

func (s *ClassMethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitClassMethod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) ClassElement() (localctx IClassElementContext) {
	this := p
	_ = this

	localctx = NewClassElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LatteParserRULE_classElement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewClassAttributeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Type_()
		}
		{
			p.SetState(181)
			p.Match(LatteParserID)
		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LatteParserT__6 {
			{
				p.SetState(182)
				p.Match(LatteParserT__6)
			}
			{
				p.SetState(183)
				p.Match(LatteParserID)
			}

			p.SetState(188)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(189)
			p.Match(LatteParserT__8)
		}

	case 2:
		localctx = NewClassMethodContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(191)
			p.Type_()
		}
		{
			p.SetState(192)
			p.Match(LatteParserID)
		}
		{
			p.SetState(193)
			p.Match(LatteParserT__0)
		}
		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&281475102539776) != 0 {
			{
				p.SetState(194)
				p.Arg()
			}

		}
		{
			p.SetState(197)
			p.Match(LatteParserT__1)
		}
		{
			p.SetState(198)
			p.Block()
		}

	}

	return localctx
}

// IItemContext is an interface to support dynamic dispatch.
type IItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsItemContext differentiates from other interfaces.
	IsItemContext()
}

type ItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItemContext() *ItemContext {
	var p = new(ItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_item
	return p
}

func (*ItemContext) IsItemContext() {}

func NewItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItemContext {
	var p = new(ItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_item

	return p
}

func (s *ItemContext) GetParser() antlr.Parser { return s.parser }

func (s *ItemContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *ItemContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterItem(s)
	}
}

func (s *ItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitItem(s)
	}
}

func (s *ItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Item() (localctx IItemContext) {
	this := p
	_ = this

	localctx = NewItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LatteParserRULE_item)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(202)
			p.Match(LatteParserID)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(203)
			p.Match(LatteParserID)
		}
		{
			p.SetState(204)
			p.Match(LatteParserT__11)
		}
		{
			p.SetState(205)
			p.expr(0)
		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EIdContext struct {
	*ExprContext
}

func NewEIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EIdContext {
	var p = new(EIdContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EIdContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *EIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEId(s)
	}
}

func (s *EIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEId(s)
	}
}

func (s *EIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEId(s)

	default:
		return t.VisitChildren(s)
	}
}

type EFunCallContext struct {
	*ExprContext
}

func NewEFunCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EFunCallContext {
	var p = new(EFunCallContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EFunCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EFunCallContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *EFunCallContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EFunCallContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EFunCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEFunCall(s)
	}
}

func (s *EFunCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEFunCall(s)
	}
}

func (s *EFunCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEFunCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type ENewArrayContext struct {
	*ExprContext
}

func NewENewArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ENewArrayContext {
	var p = new(ENewArrayContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ENewArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ENewArrayContext) Type_single() IType_singleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_singleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_singleContext)
}

func (s *ENewArrayContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ENewArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterENewArray(s)
	}
}

func (s *ENewArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitENewArray(s)
	}
}

func (s *ENewArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitENewArray(s)

	default:
		return t.VisitChildren(s)
	}
}

type ERelOpContext struct {
	*ExprContext
}

func NewERelOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ERelOpContext {
	var p = new(ERelOpContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ERelOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ERelOpContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ERelOpContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ERelOpContext) RelOp() IRelOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelOpContext)
}

func (s *ERelOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterERelOp(s)
	}
}

func (s *ERelOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitERelOp(s)
	}
}

func (s *ERelOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitERelOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ETrueContext struct {
	*ExprContext
}

func NewETrueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ETrueContext {
	var p = new(ETrueContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ETrueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ETrueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterETrue(s)
	}
}

func (s *ETrueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitETrue(s)
	}
}

func (s *ETrueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitETrue(s)

	default:
		return t.VisitChildren(s)
	}
}

type EOrContext struct {
	*ExprContext
}

func NewEOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EOrContext {
	var p = new(EOrContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EOrContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EOrContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEOr(s)
	}
}

func (s *EOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEOr(s)
	}
}

func (s *EOrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EIntContext struct {
	*ExprContext
}

func NewEIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EIntContext {
	var p = new(EIntContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EIntContext) INT() antlr.TerminalNode {
	return s.GetToken(LatteParserINT, 0)
}

func (s *EIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEInt(s)
	}
}

func (s *EIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEInt(s)
	}
}

func (s *EIntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEInt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ENewClassContext struct {
	*ExprContext
}

func NewENewClassContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ENewClassContext {
	var p = new(ENewClassContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ENewClassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ENewClassContext) Type_single() IType_singleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_singleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_singleContext)
}

func (s *ENewClassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterENewClass(s)
	}
}

func (s *ENewClassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitENewClass(s)
	}
}

func (s *ENewClassContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitENewClass(s)

	default:
		return t.VisitChildren(s)
	}
}

type EUnOpContext struct {
	*ExprContext
}

func NewEUnOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EUnOpContext {
	var p = new(EUnOpContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EUnOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EUnOpContext) UnOp() IUnOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnOpContext)
}

func (s *EUnOpContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EUnOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEUnOp(s)
	}
}

func (s *EUnOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEUnOp(s)
	}
}

func (s *EUnOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEUnOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type EStrContext struct {
	*ExprContext
}

func NewEStrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EStrContext {
	var p = new(EStrContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EStrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EStrContext) STR() antlr.TerminalNode {
	return s.GetToken(LatteParserSTR, 0)
}

func (s *EStrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEStr(s)
	}
}

func (s *EStrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEStr(s)
	}
}

func (s *EStrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEStr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EMulOpContext struct {
	*ExprContext
}

func NewEMulOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EMulOpContext {
	var p = new(EMulOpContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EMulOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EMulOpContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EMulOpContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EMulOpContext) MulOp() IMulOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMulOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMulOpContext)
}

func (s *EMulOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEMulOp(s)
	}
}

func (s *EMulOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEMulOp(s)
	}
}

func (s *EMulOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEMulOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type EAndContext struct {
	*ExprContext
}

func NewEAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EAndContext {
	var p = new(EAndContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EAndContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EAndContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEAnd(s)
	}
}

func (s *EAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEAnd(s)
	}
}

func (s *EAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type EMethodCallContext struct {
	*ExprContext
}

func NewEMethodCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EMethodCallContext {
	var p = new(EMethodCallContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EMethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EMethodCallContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EMethodCallContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EMethodCallContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *EMethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEMethodCall(s)
	}
}

func (s *EMethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEMethodCall(s)
	}
}

func (s *EMethodCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEMethodCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type EParenContext struct {
	*ExprContext
}

func NewEParenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EParenContext {
	var p = new(EParenContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EParenContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EParenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEParen(s)
	}
}

func (s *EParenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEParen(s)
	}
}

func (s *EParenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEParen(s)

	default:
		return t.VisitChildren(s)
	}
}

type EFalseContext struct {
	*ExprContext
}

func NewEFalseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EFalseContext {
	var p = new(EFalseContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EFalseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EFalseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEFalse(s)
	}
}

func (s *EFalseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEFalse(s)
	}
}

func (s *EFalseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEFalse(s)

	default:
		return t.VisitChildren(s)
	}
}

type EIndexContext struct {
	*ExprContext
}

func NewEIndexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EIndexContext {
	var p = new(EIndexContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EIndexContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EIndexContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEIndex(s)
	}
}

func (s *EIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEIndex(s)
	}
}

func (s *EIndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

type EAddOpContext struct {
	*ExprContext
}

func NewEAddOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EAddOpContext {
	var p = new(EAddOpContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EAddOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EAddOpContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EAddOpContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EAddOpContext) AddOp() IAddOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddOpContext)
}

func (s *EAddOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEAddOp(s)
	}
}

func (s *EAddOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEAddOp(s)
	}
}

func (s *EAddOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEAddOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ENullContext struct {
	*ExprContext
}

func NewENullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ENullContext {
	var p = new(ENullContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ENullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ENullContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ENullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterENull(s)
	}
}

func (s *ENullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitENull(s)
	}
}

func (s *ENullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitENull(s)

	default:
		return t.VisitChildren(s)
	}
}

type EAttrContext struct {
	*ExprContext
}

func NewEAttrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EAttrContext {
	var p = new(EAttrContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EAttrContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EAttrContext) ID() antlr.TerminalNode {
	return s.GetToken(LatteParserID, 0)
}

func (s *EAttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterEAttr(s)
	}
}

func (s *EAttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitEAttr(s)
	}
}

func (s *EAttrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitEAttr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *LatteParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, LatteParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewENewArrayContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(209)
			p.Match(LatteParserT__26)
		}
		{
			p.SetState(210)
			p.Type_single()
		}
		{
			p.SetState(211)
			p.Match(LatteParserT__9)
		}
		{
			p.SetState(212)
			p.expr(0)
		}
		{
			p.SetState(213)
			p.Match(LatteParserT__10)
		}

	case 2:
		localctx = NewENewClassContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(215)
			p.Match(LatteParserT__26)
		}
		{
			p.SetState(216)
			p.Type_single()
		}

	case 3:
		localctx = NewEUnOpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(217)
			p.UnOp()
		}
		{
			p.SetState(218)
			p.expr(14)
		}

	case 4:
		localctx = NewEIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(220)
			p.Match(LatteParserID)
		}

	case 5:
		localctx = NewEIntContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(221)
			p.Match(LatteParserINT)
		}

	case 6:
		localctx = NewETrueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(222)
			p.Match(LatteParserT__29)
		}

	case 7:
		localctx = NewEFalseContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(223)
			p.Match(LatteParserT__30)
		}

	case 8:
		localctx = NewEFunCallContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(224)
			p.Match(LatteParserID)
		}
		{
			p.SetState(225)
			p.Match(LatteParserT__0)
		}
		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1548145792122882) != 0 {
			{
				p.SetState(226)
				p.expr(0)
			}
			p.SetState(231)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == LatteParserT__6 {
				{
					p.SetState(227)
					p.Match(LatteParserT__6)
				}
				{
					p.SetState(228)
					p.expr(0)
				}

				p.SetState(233)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(236)
			p.Match(LatteParserT__1)
		}

	case 9:
		localctx = NewEStrContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(237)
			p.Match(LatteParserSTR)
		}

	case 10:
		localctx = NewENullContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LatteParserT__0 {
			{
				p.SetState(238)
				p.Match(LatteParserT__0)
			}
			{
				p.SetState(239)
				p.Type_()
			}
			{
				p.SetState(240)
				p.Match(LatteParserT__1)
			}

		}
		{
			p.SetState(244)
			p.Match(LatteParserT__31)
		}

	case 11:
		localctx = NewEParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(245)
			p.Match(LatteParserT__0)
		}
		{
			p.SetState(246)
			p.expr(0)
		}
		{
			p.SetState(247)
			p.Match(LatteParserT__1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(292)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEMulOpContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LatteParserRULE_expr)
				p.SetState(251)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(252)
					p.MulOp()
				}
				{
					p.SetState(253)
					p.expr(14)
				}

			case 2:
				localctx = NewEAddOpContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LatteParserRULE_expr)
				p.SetState(255)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(256)
					p.AddOp()
				}
				{
					p.SetState(257)
					p.expr(13)
				}

			case 3:
				localctx = NewERelOpContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LatteParserRULE_expr)
				p.SetState(259)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(260)
					p.RelOp()
				}
				{
					p.SetState(261)
					p.expr(12)
				}

			case 4:
				localctx = NewEAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LatteParserRULE_expr)
				p.SetState(263)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(264)
					p.Match(LatteParserT__27)
				}
				{
					p.SetState(265)
					p.expr(10)
				}

			case 5:
				localctx = NewEOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LatteParserRULE_expr)
				p.SetState(266)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(267)
					p.Match(LatteParserT__28)
				}
				{
					p.SetState(268)
					p.expr(9)
				}

			case 6:
				localctx = NewEIndexContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LatteParserRULE_expr)
				p.SetState(269)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(270)
					p.Match(LatteParserT__9)
				}
				{
					p.SetState(271)
					p.expr(0)
				}
				{
					p.SetState(272)
					p.Match(LatteParserT__10)
				}

			case 7:
				localctx = NewEMethodCallContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LatteParserRULE_expr)
				p.SetState(274)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(275)
					p.Match(LatteParserT__12)
				}
				{
					p.SetState(276)
					p.Match(LatteParserID)
				}
				{
					p.SetState(277)
					p.Match(LatteParserT__0)
				}
				p.SetState(286)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1548145792122882) != 0 {
					{
						p.SetState(278)
						p.expr(0)
					}
					p.SetState(283)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == LatteParserT__6 {
						{
							p.SetState(279)
							p.Match(LatteParserT__6)
						}
						{
							p.SetState(280)
							p.expr(0)
						}

						p.SetState(285)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(288)
					p.Match(LatteParserT__1)
				}

			case 8:
				localctx = NewEAttrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LatteParserRULE_expr)
				p.SetState(289)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(290)
					p.Match(LatteParserT__12)
				}
				{
					p.SetState(291)
					p.Match(LatteParserID)
				}

			}

		}
		p.SetState(296)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// IUnOpContext is an interface to support dynamic dispatch.
type IUnOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnOpContext differentiates from other interfaces.
	IsUnOpContext()
}

type UnOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnOpContext() *UnOpContext {
	var p = new(UnOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_unOp
	return p
}

func (*UnOpContext) IsUnOpContext() {}

func NewUnOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnOpContext {
	var p = new(UnOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_unOp

	return p
}

func (s *UnOpContext) GetParser() antlr.Parser { return s.parser }
func (s *UnOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterUnOp(s)
	}
}

func (s *UnOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitUnOp(s)
	}
}

func (s *UnOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitUnOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) UnOp() (localctx IUnOpContext) {
	this := p
	_ = this

	localctx = NewUnOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LatteParserRULE_unOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LatteParserT__32 || _la == LatteParserT__33) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAddOpContext is an interface to support dynamic dispatch.
type IAddOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddOpContext differentiates from other interfaces.
	IsAddOpContext()
}

type AddOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddOpContext() *AddOpContext {
	var p = new(AddOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_addOp
	return p
}

func (*AddOpContext) IsAddOpContext() {}

func NewAddOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddOpContext {
	var p = new(AddOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_addOp

	return p
}

func (s *AddOpContext) GetParser() antlr.Parser { return s.parser }
func (s *AddOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterAddOp(s)
	}
}

func (s *AddOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitAddOp(s)
	}
}

func (s *AddOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitAddOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) AddOp() (localctx IAddOpContext) {
	this := p
	_ = this

	localctx = NewAddOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LatteParserRULE_addOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LatteParserT__33 || _la == LatteParserT__34) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMulOpContext is an interface to support dynamic dispatch.
type IMulOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMulOpContext differentiates from other interfaces.
	IsMulOpContext()
}

type MulOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulOpContext() *MulOpContext {
	var p = new(MulOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_mulOp
	return p
}

func (*MulOpContext) IsMulOpContext() {}

func NewMulOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulOpContext {
	var p = new(MulOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_mulOp

	return p
}

func (s *MulOpContext) GetParser() antlr.Parser { return s.parser }
func (s *MulOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MulOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterMulOp(s)
	}
}

func (s *MulOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitMulOp(s)
	}
}

func (s *MulOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitMulOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) MulOp() (localctx IMulOpContext) {
	this := p
	_ = this

	localctx = NewMulOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LatteParserRULE_mulOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&481036337152) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRelOpContext is an interface to support dynamic dispatch.
type IRelOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelOpContext differentiates from other interfaces.
	IsRelOpContext()
}

type RelOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelOpContext() *RelOpContext {
	var p = new(RelOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatteParserRULE_relOp
	return p
}

func (*RelOpContext) IsRelOpContext() {}

func NewRelOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelOpContext {
	var p = new(RelOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatteParserRULE_relOp

	return p
}

func (s *RelOpContext) GetParser() antlr.Parser { return s.parser }
func (s *RelOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.EnterRelOp(s)
	}
}

func (s *RelOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatteListener); ok {
		listenerT.ExitRelOp(s)
	}
}

func (s *RelOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LatteVisitor:
		return t.VisitRelOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LatteParser) RelOp() (localctx IRelOpContext) {
	this := p
	_ = this

	localctx = NewRelOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LatteParserRULE_relOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34634616274944) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *LatteParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LatteParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 17)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
