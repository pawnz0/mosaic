// Code generated from PolicyLexer.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 20, 136,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 3, 2, 3, 2, 6, 2, 56,
	10, 2, 13, 2, 14, 2, 57, 3, 2, 3, 2, 3, 2, 7, 2, 63, 10, 2, 12, 2, 14,
	2, 66, 11, 2, 3, 3, 6, 3, 69, 10, 3, 13, 3, 14, 3, 70, 3, 4, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 6, 18,
	108, 10, 18, 13, 18, 14, 18, 109, 3, 18, 3, 18, 3, 19, 6, 19, 115, 10,
	19, 13, 19, 14, 19, 116, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 5, 21,
	125, 10, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3,
	26, 3, 26, 2, 2, 27, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10,
	19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19,
	37, 20, 39, 2, 41, 2, 43, 2, 45, 2, 47, 2, 49, 2, 51, 2, 3, 2, 5, 4, 2,
	12, 12, 15, 15, 4, 2, 11, 11, 34, 34, 6, 2, 37, 40, 47, 48, 165, 165, 169,
	169, 2, 137, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2,
	17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2,
	2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2,
	2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 3, 55, 3, 2,
	2, 2, 5, 68, 3, 2, 2, 2, 7, 72, 3, 2, 2, 2, 9, 75, 3, 2, 2, 2, 11, 78,
	3, 2, 2, 2, 13, 80, 3, 2, 2, 2, 15, 83, 3, 2, 2, 2, 17, 85, 3, 2, 2, 2,
	19, 88, 3, 2, 2, 2, 21, 90, 3, 2, 2, 2, 23, 93, 3, 2, 2, 2, 25, 96, 3,
	2, 2, 2, 27, 98, 3, 2, 2, 2, 29, 100, 3, 2, 2, 2, 31, 102, 3, 2, 2, 2,
	33, 104, 3, 2, 2, 2, 35, 107, 3, 2, 2, 2, 37, 114, 3, 2, 2, 2, 39, 120,
	3, 2, 2, 2, 41, 124, 3, 2, 2, 2, 43, 126, 3, 2, 2, 2, 45, 128, 3, 2, 2,
	2, 47, 130, 3, 2, 2, 2, 49, 132, 3, 2, 2, 2, 51, 134, 3, 2, 2, 2, 53, 56,
	5, 41, 21, 2, 54, 56, 5, 39, 20, 2, 55, 53, 3, 2, 2, 2, 55, 54, 3, 2, 2,
	2, 56, 57, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 64,
	3, 2, 2, 2, 59, 63, 5, 41, 21, 2, 60, 63, 5, 47, 24, 2, 61, 63, 5, 39,
	20, 2, 62, 59, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 61, 3, 2, 2, 2, 63,
	66, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 4, 3, 2, 2,
	2, 66, 64, 3, 2, 2, 2, 67, 69, 5, 47, 24, 2, 68, 67, 3, 2, 2, 2, 69, 70,
	3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 6, 3, 2, 2, 2,
	72, 73, 7, 63, 2, 2, 73, 74, 7, 63, 2, 2, 74, 8, 3, 2, 2, 2, 75, 76, 7,
	35, 2, 2, 76, 77, 7, 63, 2, 2, 77, 10, 3, 2, 2, 2, 78, 79, 7, 62, 2, 2,
	79, 12, 3, 2, 2, 2, 80, 81, 7, 62, 2, 2, 81, 82, 7, 63, 2, 2, 82, 14, 3,
	2, 2, 2, 83, 84, 7, 64, 2, 2, 84, 16, 3, 2, 2, 2, 85, 86, 7, 64, 2, 2,
	86, 87, 7, 63, 2, 2, 87, 18, 3, 2, 2, 2, 88, 89, 7, 66, 2, 2, 89, 20, 3,
	2, 2, 2, 90, 91, 7, 49, 2, 2, 91, 92, 7, 94, 2, 2, 92, 22, 3, 2, 2, 2,
	93, 94, 7, 94, 2, 2, 94, 95, 7, 49, 2, 2, 95, 24, 3, 2, 2, 2, 96, 97, 7,
	46, 2, 2, 97, 26, 3, 2, 2, 2, 98, 99, 7, 93, 2, 2, 99, 28, 3, 2, 2, 2,
	100, 101, 7, 95, 2, 2, 101, 30, 3, 2, 2, 2, 102, 103, 7, 42, 2, 2, 103,
	32, 3, 2, 2, 2, 104, 105, 7, 43, 2, 2, 105, 34, 3, 2, 2, 2, 106, 108, 9,
	2, 2, 2, 107, 106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 107, 3, 2, 2,
	2, 109, 110, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 112, 8, 18, 2, 2, 112,
	36, 3, 2, 2, 2, 113, 115, 9, 3, 2, 2, 114, 113, 3, 2, 2, 2, 115, 116, 3,
	2, 2, 2, 116, 114, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 118, 3, 2, 2,
	2, 118, 119, 8, 19, 2, 2, 119, 38, 3, 2, 2, 2, 120, 121, 9, 4, 2, 2, 121,
	40, 3, 2, 2, 2, 122, 125, 5, 43, 22, 2, 123, 125, 5, 45, 23, 2, 124, 122,
	3, 2, 2, 2, 124, 123, 3, 2, 2, 2, 125, 42, 3, 2, 2, 2, 126, 127, 4, 67,
	92, 2, 127, 44, 3, 2, 2, 2, 128, 129, 4, 99, 124, 2, 129, 46, 3, 2, 2,
	2, 130, 131, 4, 50, 59, 2, 131, 48, 3, 2, 2, 2, 132, 133, 4, 51, 59, 2,
	133, 50, 3, 2, 2, 2, 134, 135, 7, 50, 2, 2, 135, 52, 3, 2, 2, 2, 11, 2,
	55, 57, 62, 64, 70, 109, 116, 124, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'@'", "'/\\'",
	"'\\/'", "','", "'['", "']'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "Name", "Number", "TK_EQ", "TK_NEQ", "TK_LT", "TK_LTEQ", "TK_GT", "TK_GTEQ",
	"TK_AT", "TK_AND", "TK_OR", "TK_COMMA", "TK_LBRACKET", "TK_RBRACKET", "TK_LPAREN",
	"TK_RPAREN", "TK_LFCR", "TK_WS",
}

var lexerRuleNames = []string{
	"Name", "Number", "TK_EQ", "TK_NEQ", "TK_LT", "TK_LTEQ", "TK_GT", "TK_GTEQ",
	"TK_AT", "TK_AND", "TK_OR", "TK_COMMA", "TK_LBRACKET", "TK_RBRACKET", "TK_LPAREN",
	"TK_RPAREN", "TK_LFCR", "TK_WS", "SpecialCharacter", "Letter", "LetterUppercase",
	"LetterLowercase", "Digit", "NonZeroDigit", "ZeroDigit",
}

type PolicyLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewPolicyLexer(input antlr.CharStream) *PolicyLexer {

	l := new(PolicyLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "PolicyLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PolicyLexer tokens.
const (
	PolicyLexerName        = 1
	PolicyLexerNumber      = 2
	PolicyLexerTK_EQ       = 3
	PolicyLexerTK_NEQ      = 4
	PolicyLexerTK_LT       = 5
	PolicyLexerTK_LTEQ     = 6
	PolicyLexerTK_GT       = 7
	PolicyLexerTK_GTEQ     = 8
	PolicyLexerTK_AT       = 9
	PolicyLexerTK_AND      = 10
	PolicyLexerTK_OR       = 11
	PolicyLexerTK_COMMA    = 12
	PolicyLexerTK_LBRACKET = 13
	PolicyLexerTK_RBRACKET = 14
	PolicyLexerTK_LPAREN   = 15
	PolicyLexerTK_RPAREN   = 16
	PolicyLexerTK_LFCR     = 17
	PolicyLexerTK_WS       = 18
)
