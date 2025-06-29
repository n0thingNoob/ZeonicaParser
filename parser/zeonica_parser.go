// Code generated from ZeonicaParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ZeonicaParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ZeonicaParser struct {
	*antlr.BaseParser
}

var ZeonicaParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func zeonicaparserParserInit() {
	staticData := &ZeonicaParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'PE'", "'Loop'", "'Once'", "'Entry'", "'ADD'", "'ADDI'", "'SUB'",
		"'SUBI'", "'MUL'", "'DIV'", "'MAC'", "'INC'", "'LLS'", "'LRS'", "'OR'",
		"'AND'", "'XOR'", "'NOT'", "'FADD'", "'FSUB'", "'FMUL'", "'FDIV'", "'FMAC'",
		"'MOV'", "'MUL_CONST_ADD'", "'NOP'", "'CBR'", "'CMOV'", "'SGE'", "'LOAD'",
		"'STORE'", "'=>'", "'->'", "'IMM'", "'MEM'", "", "", "", "", "", "",
		"", "", "", "", "'null'", "'('", "')'", "'{'", "'}'", "'['", "']'",
		"';'", "','", "'.'", "':'", "'!'", "'^'",
	}
	staticData.SymbolicNames = []string{
		"", "PE", "LOOP", "ONCE", "ENTRY_BLOCK", "ADD", "ADDI", "SUB", "SUBI",
		"MUL", "DIV", "MAC", "INC", "LLS", "LRS", "OR", "AND", "XOR", "NOT",
		"FADD", "FSUB", "FMUL", "FDIV", "FMAC", "MOV", "MUL_CONST_ADD", "NOP",
		"CBR", "CMOV", "SGE", "LOAD", "STORE", "ENTRY_ARROW", "RIGHT_ARROW",
		"IMM", "MEM", "DECIMAL_LITERAL", "HEX_LITERAL", "OCT_LITERAL", "BINARY_LITERAL",
		"FLOAT_LITERAL", "HEX_FLOAT_LITERAL", "BOOL_LITERAL", "CHAR_LITERAL",
		"STRING_LITERAL", "TEXT_BLOCK", "NULL_LITERAL", "LPAREN", "RPAREN",
		"LBRACE", "RBRACE", "LBRACK", "RBRACK", "SEMI", "COMMA", "DOT", "COLON",
		"AND_PRED", "OR_PRED", "WS", "COMMENT", "LINE_COMMENT", "IDENTIFIER",
	}
	staticData.RuleNames = []string{
		"compilationUnit", "peBlock", "peBody", "flatStyle", "labeledGroup",
		"entryBlock", "loopType", "instGroupList", "instGroup", "label", "labelID",
		"normalInst", "operandList", "operand", "idList", "predTag", "opCode",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 62, 191, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 1, 0, 5, 0, 36, 8, 0, 10, 0, 12, 0, 39, 9, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 48, 8, 1, 1, 1, 1, 1, 1, 1, 3, 1, 53,
		8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 4, 2, 61, 8, 2, 11, 2, 12, 2,
		62, 3, 2, 65, 8, 2, 1, 3, 4, 3, 68, 8, 3, 11, 3, 12, 3, 69, 1, 4, 1, 4,
		1, 4, 5, 4, 75, 8, 4, 10, 4, 12, 4, 78, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 5, 7, 92, 8, 7, 10, 7, 12,
		7, 95, 9, 7, 1, 8, 1, 8, 5, 8, 99, 8, 8, 10, 8, 12, 8, 102, 9, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3,
		11, 126, 8, 11, 1, 12, 1, 12, 1, 12, 5, 12, 131, 8, 12, 10, 12, 12, 12,
		134, 9, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 174,
		8, 13, 1, 14, 1, 14, 1, 14, 5, 14, 179, 8, 14, 10, 14, 12, 14, 182, 9,
		14, 1, 15, 1, 15, 1, 15, 3, 15, 187, 8, 15, 1, 16, 1, 16, 1, 16, 0, 0,
		17, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 0, 2,
		1, 0, 2, 3, 1, 0, 5, 29, 197, 0, 37, 1, 0, 0, 0, 2, 42, 1, 0, 0, 0, 4,
		64, 1, 0, 0, 0, 6, 67, 1, 0, 0, 0, 8, 71, 1, 0, 0, 0, 10, 81, 1, 0, 0,
		0, 12, 88, 1, 0, 0, 0, 14, 93, 1, 0, 0, 0, 16, 96, 1, 0, 0, 0, 18, 105,
		1, 0, 0, 0, 20, 108, 1, 0, 0, 0, 22, 125, 1, 0, 0, 0, 24, 127, 1, 0, 0,
		0, 26, 173, 1, 0, 0, 0, 28, 175, 1, 0, 0, 0, 30, 186, 1, 0, 0, 0, 32, 188,
		1, 0, 0, 0, 34, 36, 3, 2, 1, 0, 35, 34, 1, 0, 0, 0, 36, 39, 1, 0, 0, 0,
		37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 40, 1, 0, 0, 0, 39, 37, 1,
		0, 0, 0, 40, 41, 5, 0, 0, 1, 41, 1, 1, 0, 0, 0, 42, 43, 5, 1, 0, 0, 43,
		44, 5, 47, 0, 0, 44, 47, 5, 36, 0, 0, 45, 46, 5, 54, 0, 0, 46, 48, 5, 36,
		0, 0, 47, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 52,
		5, 48, 0, 0, 50, 51, 5, 32, 0, 0, 51, 53, 3, 12, 6, 0, 52, 50, 1, 0, 0,
		0, 52, 53, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 55, 5, 49, 0, 0, 55, 56,
		3, 4, 2, 0, 56, 57, 5, 50, 0, 0, 57, 3, 1, 0, 0, 0, 58, 65, 3, 6, 3, 0,
		59, 61, 3, 10, 5, 0, 60, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 60, 1,
		0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 65, 1, 0, 0, 0, 64, 58, 1, 0, 0, 0, 64,
		60, 1, 0, 0, 0, 65, 5, 1, 0, 0, 0, 66, 68, 3, 8, 4, 0, 67, 66, 1, 0, 0,
		0, 68, 69, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 7, 1,
		0, 0, 0, 71, 72, 3, 18, 9, 0, 72, 76, 5, 49, 0, 0, 73, 75, 3, 22, 11, 0,
		74, 73, 1, 0, 0, 0, 75, 78, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 76, 77, 1,
		0, 0, 0, 77, 79, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 79, 80, 5, 50, 0, 0, 80,
		9, 1, 0, 0, 0, 81, 82, 5, 4, 0, 0, 82, 83, 5, 32, 0, 0, 83, 84, 3, 12,
		6, 0, 84, 85, 5, 49, 0, 0, 85, 86, 3, 14, 7, 0, 86, 87, 5, 50, 0, 0, 87,
		11, 1, 0, 0, 0, 88, 89, 7, 0, 0, 0, 89, 13, 1, 0, 0, 0, 90, 92, 3, 16,
		8, 0, 91, 90, 1, 0, 0, 0, 92, 95, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 94,
		1, 0, 0, 0, 94, 15, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 96, 100, 5, 49, 0,
		0, 97, 99, 3, 22, 11, 0, 98, 97, 1, 0, 0, 0, 99, 102, 1, 0, 0, 0, 100,
		98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 103, 1, 0, 0, 0, 102, 100, 1,
		0, 0, 0, 103, 104, 5, 50, 0, 0, 104, 17, 1, 0, 0, 0, 105, 106, 3, 20, 10,
		0, 106, 107, 5, 56, 0, 0, 107, 19, 1, 0, 0, 0, 108, 109, 5, 62, 0, 0, 109,
		21, 1, 0, 0, 0, 110, 111, 3, 32, 16, 0, 111, 112, 5, 54, 0, 0, 112, 113,
		3, 24, 12, 0, 113, 114, 5, 33, 0, 0, 114, 115, 3, 24, 12, 0, 115, 126,
		1, 0, 0, 0, 116, 117, 3, 32, 16, 0, 117, 118, 5, 54, 0, 0, 118, 119, 3,
		24, 12, 0, 119, 126, 1, 0, 0, 0, 120, 126, 3, 32, 16, 0, 121, 122, 3, 26,
		13, 0, 122, 123, 5, 33, 0, 0, 123, 124, 3, 26, 13, 0, 124, 126, 1, 0, 0,
		0, 125, 110, 1, 0, 0, 0, 125, 116, 1, 0, 0, 0, 125, 120, 1, 0, 0, 0, 125,
		121, 1, 0, 0, 0, 126, 23, 1, 0, 0, 0, 127, 132, 3, 26, 13, 0, 128, 129,
		5, 54, 0, 0, 129, 131, 3, 26, 13, 0, 130, 128, 1, 0, 0, 0, 131, 134, 1,
		0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 25, 1, 0, 0,
		0, 134, 132, 1, 0, 0, 0, 135, 136, 3, 30, 15, 0, 136, 137, 5, 51, 0, 0,
		137, 138, 3, 28, 14, 0, 138, 139, 5, 52, 0, 0, 139, 174, 1, 0, 0, 0, 140,
		141, 5, 34, 0, 0, 141, 142, 5, 51, 0, 0, 142, 143, 5, 36, 0, 0, 143, 174,
		5, 52, 0, 0, 144, 145, 5, 34, 0, 0, 145, 146, 5, 51, 0, 0, 146, 147, 5,
		38, 0, 0, 147, 174, 5, 52, 0, 0, 148, 149, 5, 34, 0, 0, 149, 150, 5, 51,
		0, 0, 150, 151, 5, 37, 0, 0, 151, 174, 5, 52, 0, 0, 152, 153, 5, 34, 0,
		0, 153, 154, 5, 51, 0, 0, 154, 155, 5, 39, 0, 0, 155, 174, 5, 52, 0, 0,
		156, 157, 5, 34, 0, 0, 157, 158, 5, 51, 0, 0, 158, 159, 5, 40, 0, 0, 159,
		174, 5, 52, 0, 0, 160, 161, 3, 30, 15, 0, 161, 162, 5, 35, 0, 0, 162, 163,
		5, 51, 0, 0, 163, 164, 3, 28, 14, 0, 164, 165, 5, 52, 0, 0, 165, 174, 1,
		0, 0, 0, 166, 167, 3, 30, 15, 0, 167, 168, 5, 35, 0, 0, 168, 169, 5, 51,
		0, 0, 169, 170, 5, 37, 0, 0, 170, 171, 5, 52, 0, 0, 171, 174, 1, 0, 0,
		0, 172, 174, 3, 20, 10, 0, 173, 135, 1, 0, 0, 0, 173, 140, 1, 0, 0, 0,
		173, 144, 1, 0, 0, 0, 173, 148, 1, 0, 0, 0, 173, 152, 1, 0, 0, 0, 173,
		156, 1, 0, 0, 0, 173, 160, 1, 0, 0, 0, 173, 166, 1, 0, 0, 0, 173, 172,
		1, 0, 0, 0, 174, 27, 1, 0, 0, 0, 175, 180, 5, 62, 0, 0, 176, 177, 5, 54,
		0, 0, 177, 179, 5, 62, 0, 0, 178, 176, 1, 0, 0, 0, 179, 182, 1, 0, 0, 0,
		180, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 29, 1, 0, 0, 0, 182, 180,
		1, 0, 0, 0, 183, 187, 5, 57, 0, 0, 184, 187, 5, 58, 0, 0, 185, 187, 1,
		0, 0, 0, 186, 183, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 186, 185, 1, 0, 0,
		0, 187, 31, 1, 0, 0, 0, 188, 189, 7, 1, 0, 0, 189, 33, 1, 0, 0, 0, 14,
		37, 47, 52, 62, 64, 69, 76, 93, 100, 125, 132, 173, 180, 186,
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

// ZeonicaParserInit initializes any static state used to implement ZeonicaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewZeonicaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ZeonicaParserInit() {
	staticData := &ZeonicaParserParserStaticData
	staticData.once.Do(zeonicaparserParserInit)
}

// NewZeonicaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewZeonicaParser(input antlr.TokenStream) *ZeonicaParser {
	ZeonicaParserInit()
	this := new(ZeonicaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ZeonicaParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ZeonicaParser.g4"

	return this
}

// ZeonicaParser tokens.
const (
	ZeonicaParserEOF               = antlr.TokenEOF
	ZeonicaParserPE                = 1
	ZeonicaParserLOOP              = 2
	ZeonicaParserONCE              = 3
	ZeonicaParserENTRY_BLOCK       = 4
	ZeonicaParserADD               = 5
	ZeonicaParserADDI              = 6
	ZeonicaParserSUB               = 7
	ZeonicaParserSUBI              = 8
	ZeonicaParserMUL               = 9
	ZeonicaParserDIV               = 10
	ZeonicaParserMAC               = 11
	ZeonicaParserINC               = 12
	ZeonicaParserLLS               = 13
	ZeonicaParserLRS               = 14
	ZeonicaParserOR                = 15
	ZeonicaParserAND               = 16
	ZeonicaParserXOR               = 17
	ZeonicaParserNOT               = 18
	ZeonicaParserFADD              = 19
	ZeonicaParserFSUB              = 20
	ZeonicaParserFMUL              = 21
	ZeonicaParserFDIV              = 22
	ZeonicaParserFMAC              = 23
	ZeonicaParserMOV               = 24
	ZeonicaParserMUL_CONST_ADD     = 25
	ZeonicaParserNOP               = 26
	ZeonicaParserCBR               = 27
	ZeonicaParserCMOV              = 28
	ZeonicaParserSGE               = 29
	ZeonicaParserLOAD              = 30
	ZeonicaParserSTORE             = 31
	ZeonicaParserENTRY_ARROW       = 32
	ZeonicaParserRIGHT_ARROW       = 33
	ZeonicaParserIMM               = 34
	ZeonicaParserMEM               = 35
	ZeonicaParserDECIMAL_LITERAL   = 36
	ZeonicaParserHEX_LITERAL       = 37
	ZeonicaParserOCT_LITERAL       = 38
	ZeonicaParserBINARY_LITERAL    = 39
	ZeonicaParserFLOAT_LITERAL     = 40
	ZeonicaParserHEX_FLOAT_LITERAL = 41
	ZeonicaParserBOOL_LITERAL      = 42
	ZeonicaParserCHAR_LITERAL      = 43
	ZeonicaParserSTRING_LITERAL    = 44
	ZeonicaParserTEXT_BLOCK        = 45
	ZeonicaParserNULL_LITERAL      = 46
	ZeonicaParserLPAREN            = 47
	ZeonicaParserRPAREN            = 48
	ZeonicaParserLBRACE            = 49
	ZeonicaParserRBRACE            = 50
	ZeonicaParserLBRACK            = 51
	ZeonicaParserRBRACK            = 52
	ZeonicaParserSEMI              = 53
	ZeonicaParserCOMMA             = 54
	ZeonicaParserDOT               = 55
	ZeonicaParserCOLON             = 56
	ZeonicaParserAND_PRED          = 57
	ZeonicaParserOR_PRED           = 58
	ZeonicaParserWS                = 59
	ZeonicaParserCOMMENT           = 60
	ZeonicaParserLINE_COMMENT      = 61
	ZeonicaParserIDENTIFIER        = 62
)

// ZeonicaParser rules.
const (
	ZeonicaParserRULE_compilationUnit = 0
	ZeonicaParserRULE_peBlock         = 1
	ZeonicaParserRULE_peBody          = 2
	ZeonicaParserRULE_flatStyle       = 3
	ZeonicaParserRULE_labeledGroup    = 4
	ZeonicaParserRULE_entryBlock      = 5
	ZeonicaParserRULE_loopType        = 6
	ZeonicaParserRULE_instGroupList   = 7
	ZeonicaParserRULE_instGroup       = 8
	ZeonicaParserRULE_label           = 9
	ZeonicaParserRULE_labelID         = 10
	ZeonicaParserRULE_normalInst      = 11
	ZeonicaParserRULE_operandList     = 12
	ZeonicaParserRULE_operand         = 13
	ZeonicaParserRULE_idList          = 14
	ZeonicaParserRULE_predTag         = 15
	ZeonicaParserRULE_opCode          = 16
)

// ICompilationUnitContext is an interface to support dynamic dispatch.
type ICompilationUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllPeBlock() []IPeBlockContext
	PeBlock(i int) IPeBlockContext

	// IsCompilationUnitContext differentiates from other interfaces.
	IsCompilationUnitContext()
}

type CompilationUnitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompilationUnitContext() *CompilationUnitContext {
	var p = new(CompilationUnitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_compilationUnit
	return p
}

func InitEmptyCompilationUnitContext(p *CompilationUnitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_compilationUnit
}

func (*CompilationUnitContext) IsCompilationUnitContext() {}

func NewCompilationUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompilationUnitContext {
	var p = new(CompilationUnitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZeonicaParserRULE_compilationUnit

	return p
}

func (s *CompilationUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *CompilationUnitContext) EOF() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserEOF, 0)
}

func (s *CompilationUnitContext) AllPeBlock() []IPeBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPeBlockContext); ok {
			len++
		}
	}

	tst := make([]IPeBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPeBlockContext); ok {
			tst[i] = t.(IPeBlockContext)
			i++
		}
	}

	return tst
}

func (s *CompilationUnitContext) PeBlock(i int) IPeBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPeBlockContext); ok {
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

	return t.(IPeBlockContext)
}

func (s *CompilationUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompilationUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompilationUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.EnterCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.ExitCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZeonicaParserVisitor:
		return t.VisitCompilationUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZeonicaParser) CompilationUnit() (localctx ICompilationUnitContext) {
	localctx = NewCompilationUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ZeonicaParserRULE_compilationUnit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ZeonicaParserPE {
		{
			p.SetState(34)
			p.PeBlock()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(40)
		p.Match(ZeonicaParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPeBlockContext is an interface to support dynamic dispatch.
type IPeBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllDECIMAL_LITERAL() []antlr.TerminalNode
	DECIMAL_LITERAL(i int) antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	PeBody() IPeBodyContext
	RBRACE() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	ENTRY_ARROW() antlr.TerminalNode
	LoopType() ILoopTypeContext

	// IsPeBlockContext differentiates from other interfaces.
	IsPeBlockContext()
}

type PeBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPeBlockContext() *PeBlockContext {
	var p = new(PeBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_peBlock
	return p
}

func InitEmptyPeBlockContext(p *PeBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_peBlock
}

func (*PeBlockContext) IsPeBlockContext() {}

func NewPeBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PeBlockContext {
	var p = new(PeBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZeonicaParserRULE_peBlock

	return p
}

func (s *PeBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *PeBlockContext) PE() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserPE, 0)
}

func (s *PeBlockContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserLPAREN, 0)
}

func (s *PeBlockContext) AllDECIMAL_LITERAL() []antlr.TerminalNode {
	return s.GetTokens(ZeonicaParserDECIMAL_LITERAL)
}

func (s *PeBlockContext) DECIMAL_LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(ZeonicaParserDECIMAL_LITERAL, i)
}

func (s *PeBlockContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserRPAREN, 0)
}

func (s *PeBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserLBRACE, 0)
}

func (s *PeBlockContext) PeBody() IPeBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPeBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPeBodyContext)
}

func (s *PeBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserRBRACE, 0)
}

func (s *PeBlockContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserCOMMA, 0)
}

func (s *PeBlockContext) ENTRY_ARROW() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserENTRY_ARROW, 0)
}

func (s *PeBlockContext) LoopType() ILoopTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopTypeContext)
}

func (s *PeBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PeBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PeBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.EnterPeBlock(s)
	}
}

func (s *PeBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.ExitPeBlock(s)
	}
}

func (s *PeBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZeonicaParserVisitor:
		return t.VisitPeBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZeonicaParser) PeBlock() (localctx IPeBlockContext) {
	localctx = NewPeBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ZeonicaParserRULE_peBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Match(ZeonicaParserPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)
		p.Match(ZeonicaParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Match(ZeonicaParserDECIMAL_LITERAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ZeonicaParserCOMMA {
		{
			p.SetState(45)
			p.Match(ZeonicaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)
			p.Match(ZeonicaParserDECIMAL_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(49)
		p.Match(ZeonicaParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ZeonicaParserENTRY_ARROW {
		{
			p.SetState(50)
			p.Match(ZeonicaParserENTRY_ARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.LoopType()
		}

	}
	{
		p.SetState(54)
		p.Match(ZeonicaParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)
		p.PeBody()
	}
	{
		p.SetState(56)
		p.Match(ZeonicaParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPeBodyContext is an interface to support dynamic dispatch.
type IPeBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPeBodyContext differentiates from other interfaces.
	IsPeBodyContext()
}

type PeBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPeBodyContext() *PeBodyContext {
	var p = new(PeBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_peBody
	return p
}

func InitEmptyPeBodyContext(p *PeBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_peBody
}

func (*PeBodyContext) IsPeBodyContext() {}

func NewPeBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PeBodyContext {
	var p = new(PeBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZeonicaParserRULE_peBody

	return p
}

func (s *PeBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *PeBodyContext) CopyAll(ctx *PeBodyContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PeBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PeBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FlatBodyContext struct {
	PeBodyContext
}

func NewFlatBodyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FlatBodyContext {
	var p = new(FlatBodyContext)

	InitEmptyPeBodyContext(&p.PeBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*PeBodyContext))

	return p
}

func (s *FlatBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlatBodyContext) FlatStyle() IFlatStyleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFlatStyleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFlatStyleContext)
}

func (s *FlatBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.EnterFlatBody(s)
	}
}

func (s *FlatBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.ExitFlatBody(s)
	}
}

func (s *FlatBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZeonicaParserVisitor:
		return t.VisitFlatBody(s)

	default:
		return t.VisitChildren(s)
	}
}

type EntryBlockBodyContext struct {
	PeBodyContext
}

func NewEntryBlockBodyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EntryBlockBodyContext {
	var p = new(EntryBlockBodyContext)

	InitEmptyPeBodyContext(&p.PeBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*PeBodyContext))

	return p
}

func (s *EntryBlockBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntryBlockBodyContext) AllEntryBlock() []IEntryBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEntryBlockContext); ok {
			len++
		}
	}

	tst := make([]IEntryBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEntryBlockContext); ok {
			tst[i] = t.(IEntryBlockContext)
			i++
		}
	}

	return tst
}

func (s *EntryBlockBodyContext) EntryBlock(i int) IEntryBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEntryBlockContext); ok {
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

	return t.(IEntryBlockContext)
}

func (s *EntryBlockBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.EnterEntryBlockBody(s)
	}
}

func (s *EntryBlockBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.ExitEntryBlockBody(s)
	}
}

func (s *EntryBlockBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZeonicaParserVisitor:
		return t.VisitEntryBlockBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZeonicaParser) PeBody() (localctx IPeBodyContext) {
	localctx = NewPeBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ZeonicaParserRULE_peBody)
	var _la int

	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ZeonicaParserIDENTIFIER:
		localctx = NewFlatBodyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.FlatStyle()
		}

	case ZeonicaParserENTRY_BLOCK:
		localctx = NewEntryBlockBodyContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ZeonicaParserENTRY_BLOCK {
			{
				p.SetState(59)
				p.EntryBlock()
			}

			p.SetState(62)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFlatStyleContext is an interface to support dynamic dispatch.
type IFlatStyleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLabeledGroup() []ILabeledGroupContext
	LabeledGroup(i int) ILabeledGroupContext

	// IsFlatStyleContext differentiates from other interfaces.
	IsFlatStyleContext()
}

type FlatStyleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFlatStyleContext() *FlatStyleContext {
	var p = new(FlatStyleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_flatStyle
	return p
}

func InitEmptyFlatStyleContext(p *FlatStyleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_flatStyle
}

func (*FlatStyleContext) IsFlatStyleContext() {}

func NewFlatStyleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FlatStyleContext {
	var p = new(FlatStyleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZeonicaParserRULE_flatStyle

	return p
}

func (s *FlatStyleContext) GetParser() antlr.Parser { return s.parser }

func (s *FlatStyleContext) AllLabeledGroup() []ILabeledGroupContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILabeledGroupContext); ok {
			len++
		}
	}

	tst := make([]ILabeledGroupContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILabeledGroupContext); ok {
			tst[i] = t.(ILabeledGroupContext)
			i++
		}
	}

	return tst
}

func (s *FlatStyleContext) LabeledGroup(i int) ILabeledGroupContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabeledGroupContext); ok {
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

	return t.(ILabeledGroupContext)
}

func (s *FlatStyleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlatStyleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FlatStyleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.EnterFlatStyle(s)
	}
}

func (s *FlatStyleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.ExitFlatStyle(s)
	}
}

func (s *FlatStyleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZeonicaParserVisitor:
		return t.VisitFlatStyle(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZeonicaParser) FlatStyle() (localctx IFlatStyleContext) {
	localctx = NewFlatStyleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ZeonicaParserRULE_flatStyle)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ZeonicaParserIDENTIFIER {
		{
			p.SetState(66)
			p.LabeledGroup()
		}

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILabeledGroupContext is an interface to support dynamic dispatch.
type ILabeledGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Label() ILabelContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllNormalInst() []INormalInstContext
	NormalInst(i int) INormalInstContext

	// IsLabeledGroupContext differentiates from other interfaces.
	IsLabeledGroupContext()
}

type LabeledGroupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabeledGroupContext() *LabeledGroupContext {
	var p = new(LabeledGroupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_labeledGroup
	return p
}

func InitEmptyLabeledGroupContext(p *LabeledGroupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_labeledGroup
}

func (*LabeledGroupContext) IsLabeledGroupContext() {}

func NewLabeledGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabeledGroupContext {
	var p = new(LabeledGroupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZeonicaParserRULE_labeledGroup

	return p
}

func (s *LabeledGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *LabeledGroupContext) Label() ILabelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *LabeledGroupContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserLBRACE, 0)
}

func (s *LabeledGroupContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserRBRACE, 0)
}

func (s *LabeledGroupContext) AllNormalInst() []INormalInstContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INormalInstContext); ok {
			len++
		}
	}

	tst := make([]INormalInstContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INormalInstContext); ok {
			tst[i] = t.(INormalInstContext)
			i++
		}
	}

	return tst
}

func (s *LabeledGroupContext) NormalInst(i int) INormalInstContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INormalInstContext); ok {
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

	return t.(INormalInstContext)
}

func (s *LabeledGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabeledGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabeledGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.EnterLabeledGroup(s)
	}
}

func (s *LabeledGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.ExitLabeledGroup(s)
	}
}

func (s *LabeledGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZeonicaParserVisitor:
		return t.VisitLabeledGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZeonicaParser) LabeledGroup() (localctx ILabeledGroupContext) {
	localctx = NewLabeledGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ZeonicaParserRULE_labeledGroup)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Label()
	}
	{
		p.SetState(72)
		p.Match(ZeonicaParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&5046283435081990112) != 0 {
		{
			p.SetState(73)
			p.NormalInst()
		}

		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(79)
		p.Match(ZeonicaParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEntryBlockContext is an interface to support dynamic dispatch.
type IEntryBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENTRY_BLOCK() antlr.TerminalNode
	ENTRY_ARROW() antlr.TerminalNode
	LoopType() ILoopTypeContext
	LBRACE() antlr.TerminalNode
	InstGroupList() IInstGroupListContext
	RBRACE() antlr.TerminalNode

	// IsEntryBlockContext differentiates from other interfaces.
	IsEntryBlockContext()
}

type EntryBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntryBlockContext() *EntryBlockContext {
	var p = new(EntryBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_entryBlock
	return p
}

func InitEmptyEntryBlockContext(p *EntryBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_entryBlock
}

func (*EntryBlockContext) IsEntryBlockContext() {}

func NewEntryBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntryBlockContext {
	var p = new(EntryBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZeonicaParserRULE_entryBlock

	return p
}

func (s *EntryBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *EntryBlockContext) ENTRY_BLOCK() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserENTRY_BLOCK, 0)
}

func (s *EntryBlockContext) ENTRY_ARROW() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserENTRY_ARROW, 0)
}

func (s *EntryBlockContext) LoopType() ILoopTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopTypeContext)
}

func (s *EntryBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserLBRACE, 0)
}

func (s *EntryBlockContext) InstGroupList() IInstGroupListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstGroupListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstGroupListContext)
}

func (s *EntryBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserRBRACE, 0)
}

func (s *EntryBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntryBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntryBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.EnterEntryBlock(s)
	}
}

func (s *EntryBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.ExitEntryBlock(s)
	}
}

func (s *EntryBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZeonicaParserVisitor:
		return t.VisitEntryBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZeonicaParser) EntryBlock() (localctx IEntryBlockContext) {
	localctx = NewEntryBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ZeonicaParserRULE_entryBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(ZeonicaParserENTRY_BLOCK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Match(ZeonicaParserENTRY_ARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.LoopType()
	}
	{
		p.SetState(84)
		p.Match(ZeonicaParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		p.InstGroupList()
	}
	{
		p.SetState(86)
		p.Match(ZeonicaParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoopTypeContext is an interface to support dynamic dispatch.
type ILoopTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LOOP() antlr.TerminalNode
	ONCE() antlr.TerminalNode

	// IsLoopTypeContext differentiates from other interfaces.
	IsLoopTypeContext()
}

type LoopTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopTypeContext() *LoopTypeContext {
	var p = new(LoopTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_loopType
	return p
}

func InitEmptyLoopTypeContext(p *LoopTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_loopType
}

func (*LoopTypeContext) IsLoopTypeContext() {}

func NewLoopTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopTypeContext {
	var p = new(LoopTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZeonicaParserRULE_loopType

	return p
}

func (s *LoopTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopTypeContext) LOOP() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserLOOP, 0)
}

func (s *LoopTypeContext) ONCE() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserONCE, 0)
}

func (s *LoopTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.EnterLoopType(s)
	}
}

func (s *LoopTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.ExitLoopType(s)
	}
}

func (s *LoopTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZeonicaParserVisitor:
		return t.VisitLoopType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZeonicaParser) LoopType() (localctx ILoopTypeContext) {
	localctx = NewLoopTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ZeonicaParserRULE_loopType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ZeonicaParserLOOP || _la == ZeonicaParserONCE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstGroupListContext is an interface to support dynamic dispatch.
type IInstGroupListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllInstGroup() []IInstGroupContext
	InstGroup(i int) IInstGroupContext

	// IsInstGroupListContext differentiates from other interfaces.
	IsInstGroupListContext()
}

type InstGroupListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstGroupListContext() *InstGroupListContext {
	var p = new(InstGroupListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_instGroupList
	return p
}

func InitEmptyInstGroupListContext(p *InstGroupListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_instGroupList
}

func (*InstGroupListContext) IsInstGroupListContext() {}

func NewInstGroupListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstGroupListContext {
	var p = new(InstGroupListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZeonicaParserRULE_instGroupList

	return p
}

func (s *InstGroupListContext) GetParser() antlr.Parser { return s.parser }

func (s *InstGroupListContext) AllInstGroup() []IInstGroupContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstGroupContext); ok {
			len++
		}
	}

	tst := make([]IInstGroupContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstGroupContext); ok {
			tst[i] = t.(IInstGroupContext)
			i++
		}
	}

	return tst
}

func (s *InstGroupListContext) InstGroup(i int) IInstGroupContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstGroupContext); ok {
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

	return t.(IInstGroupContext)
}

func (s *InstGroupListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstGroupListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstGroupListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.EnterInstGroupList(s)
	}
}

func (s *InstGroupListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.ExitInstGroupList(s)
	}
}

func (s *InstGroupListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZeonicaParserVisitor:
		return t.VisitInstGroupList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZeonicaParser) InstGroupList() (localctx IInstGroupListContext) {
	localctx = NewInstGroupListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ZeonicaParserRULE_instGroupList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ZeonicaParserLBRACE {
		{
			p.SetState(90)
			p.InstGroup()
		}

		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstGroupContext is an interface to support dynamic dispatch.
type IInstGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllNormalInst() []INormalInstContext
	NormalInst(i int) INormalInstContext

	// IsInstGroupContext differentiates from other interfaces.
	IsInstGroupContext()
}

type InstGroupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstGroupContext() *InstGroupContext {
	var p = new(InstGroupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_instGroup
	return p
}

func InitEmptyInstGroupContext(p *InstGroupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_instGroup
}

func (*InstGroupContext) IsInstGroupContext() {}

func NewInstGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstGroupContext {
	var p = new(InstGroupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZeonicaParserRULE_instGroup

	return p
}

func (s *InstGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *InstGroupContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserLBRACE, 0)
}

func (s *InstGroupContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserRBRACE, 0)
}

func (s *InstGroupContext) AllNormalInst() []INormalInstContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INormalInstContext); ok {
			len++
		}
	}

	tst := make([]INormalInstContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INormalInstContext); ok {
			tst[i] = t.(INormalInstContext)
			i++
		}
	}

	return tst
}

func (s *InstGroupContext) NormalInst(i int) INormalInstContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INormalInstContext); ok {
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

	return t.(INormalInstContext)
}

func (s *InstGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.EnterInstGroup(s)
	}
}

func (s *InstGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.ExitInstGroup(s)
	}
}

func (s *InstGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZeonicaParserVisitor:
		return t.VisitInstGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZeonicaParser) InstGroup() (localctx IInstGroupContext) {
	localctx = NewInstGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ZeonicaParserRULE_instGroup)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(ZeonicaParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&5046283435081990112) != 0 {
		{
			p.SetState(97)
			p.NormalInst()
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(103)
		p.Match(ZeonicaParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LabelID() ILabelIDContext
	COLON() antlr.TerminalNode

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_label
	return p
}

func InitEmptyLabelContext(p *LabelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_label
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZeonicaParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) LabelID() ILabelIDContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelIDContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelIDContext)
}

func (s *LabelContext) COLON() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserCOLON, 0)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.ExitLabel(s)
	}
}

func (s *LabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZeonicaParserVisitor:
		return t.VisitLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZeonicaParser) Label() (localctx ILabelContext) {
	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ZeonicaParserRULE_label)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.LabelID()
	}
	{
		p.SetState(106)
		p.Match(ZeonicaParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILabelIDContext is an interface to support dynamic dispatch.
type ILabelIDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsLabelIDContext differentiates from other interfaces.
	IsLabelIDContext()
}

type LabelIDContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelIDContext() *LabelIDContext {
	var p = new(LabelIDContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_labelID
	return p
}

func InitEmptyLabelIDContext(p *LabelIDContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_labelID
}

func (*LabelIDContext) IsLabelIDContext() {}

func NewLabelIDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelIDContext {
	var p = new(LabelIDContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZeonicaParserRULE_labelID

	return p
}

func (s *LabelIDContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelIDContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserIDENTIFIER, 0)
}

func (s *LabelIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelIDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelIDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.EnterLabelID(s)
	}
}

func (s *LabelIDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.ExitLabelID(s)
	}
}

func (s *LabelIDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZeonicaParserVisitor:
		return t.VisitLabelID(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZeonicaParser) LabelID() (localctx ILabelIDContext) {
	localctx = NewLabelIDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ZeonicaParserRULE_labelID)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(ZeonicaParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INormalInstContext is an interface to support dynamic dispatch.
type INormalInstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpCode() IOpCodeContext
	COMMA() antlr.TerminalNode
	AllOperandList() []IOperandListContext
	OperandList(i int) IOperandListContext
	RIGHT_ARROW() antlr.TerminalNode
	AllOperand() []IOperandContext
	Operand(i int) IOperandContext

	// IsNormalInstContext differentiates from other interfaces.
	IsNormalInstContext()
}

type NormalInstContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNormalInstContext() *NormalInstContext {
	var p = new(NormalInstContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_normalInst
	return p
}

func InitEmptyNormalInstContext(p *NormalInstContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_normalInst
}

func (*NormalInstContext) IsNormalInstContext() {}

func NewNormalInstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NormalInstContext {
	var p = new(NormalInstContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZeonicaParserRULE_normalInst

	return p
}

func (s *NormalInstContext) GetParser() antlr.Parser { return s.parser }

func (s *NormalInstContext) OpCode() IOpCodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpCodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpCodeContext)
}

func (s *NormalInstContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserCOMMA, 0)
}

func (s *NormalInstContext) AllOperandList() []IOperandListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperandListContext); ok {
			len++
		}
	}

	tst := make([]IOperandListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperandListContext); ok {
			tst[i] = t.(IOperandListContext)
			i++
		}
	}

	return tst
}

func (s *NormalInstContext) OperandList(i int) IOperandListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandListContext); ok {
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

	return t.(IOperandListContext)
}

func (s *NormalInstContext) RIGHT_ARROW() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserRIGHT_ARROW, 0)
}

func (s *NormalInstContext) AllOperand() []IOperandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperandContext); ok {
			len++
		}
	}

	tst := make([]IOperandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperandContext); ok {
			tst[i] = t.(IOperandContext)
			i++
		}
	}

	return tst
}

func (s *NormalInstContext) Operand(i int) IOperandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
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

	return t.(IOperandContext)
}

func (s *NormalInstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NormalInstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NormalInstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.EnterNormalInst(s)
	}
}

func (s *NormalInstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.ExitNormalInst(s)
	}
}

func (s *NormalInstContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZeonicaParserVisitor:
		return t.VisitNormalInst(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZeonicaParser) NormalInst() (localctx INormalInstContext) {
	localctx = NewNormalInstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ZeonicaParserRULE_normalInst)
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.OpCode()
		}
		{
			p.SetState(111)
			p.Match(ZeonicaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.OperandList()
		}
		{
			p.SetState(113)
			p.Match(ZeonicaParserRIGHT_ARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.OperandList()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.OpCode()
		}
		{
			p.SetState(117)
			p.Match(ZeonicaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)
			p.OperandList()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(120)
			p.OpCode()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(121)
			p.Operand()
		}
		{
			p.SetState(122)
			p.Match(ZeonicaParserRIGHT_ARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Operand()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperandListContext is an interface to support dynamic dispatch.
type IOperandListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOperand() []IOperandContext
	Operand(i int) IOperandContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsOperandListContext differentiates from other interfaces.
	IsOperandListContext()
}

type OperandListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandListContext() *OperandListContext {
	var p = new(OperandListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_operandList
	return p
}

func InitEmptyOperandListContext(p *OperandListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_operandList
}

func (*OperandListContext) IsOperandListContext() {}

func NewOperandListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandListContext {
	var p = new(OperandListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZeonicaParserRULE_operandList

	return p
}

func (s *OperandListContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandListContext) AllOperand() []IOperandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperandContext); ok {
			len++
		}
	}

	tst := make([]IOperandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperandContext); ok {
			tst[i] = t.(IOperandContext)
			i++
		}
	}

	return tst
}

func (s *OperandListContext) Operand(i int) IOperandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
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

	return t.(IOperandContext)
}

func (s *OperandListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ZeonicaParserCOMMA)
}

func (s *OperandListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ZeonicaParserCOMMA, i)
}

func (s *OperandListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.EnterOperandList(s)
	}
}

func (s *OperandListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.ExitOperandList(s)
	}
}

func (s *OperandListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZeonicaParserVisitor:
		return t.VisitOperandList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZeonicaParser) OperandList() (localctx IOperandListContext) {
	localctx = NewOperandListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ZeonicaParserRULE_operandList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Operand()
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ZeonicaParserCOMMA {
		{
			p.SetState(128)
			p.Match(ZeonicaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.Operand()
		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PredTag() IPredTagContext
	LBRACK() antlr.TerminalNode
	IdList() IIdListContext
	RBRACK() antlr.TerminalNode
	IMM() antlr.TerminalNode
	DECIMAL_LITERAL() antlr.TerminalNode
	OCT_LITERAL() antlr.TerminalNode
	HEX_LITERAL() antlr.TerminalNode
	BINARY_LITERAL() antlr.TerminalNode
	FLOAT_LITERAL() antlr.TerminalNode
	MEM() antlr.TerminalNode
	LabelID() ILabelIDContext

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_operand
	return p
}

func InitEmptyOperandContext(p *OperandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_operand
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZeonicaParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) PredTag() IPredTagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPredTagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPredTagContext)
}

func (s *OperandContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserLBRACK, 0)
}

func (s *OperandContext) IdList() IIdListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdListContext)
}

func (s *OperandContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserRBRACK, 0)
}

func (s *OperandContext) IMM() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserIMM, 0)
}

func (s *OperandContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserDECIMAL_LITERAL, 0)
}

func (s *OperandContext) OCT_LITERAL() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserOCT_LITERAL, 0)
}

func (s *OperandContext) HEX_LITERAL() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserHEX_LITERAL, 0)
}

func (s *OperandContext) BINARY_LITERAL() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserBINARY_LITERAL, 0)
}

func (s *OperandContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserFLOAT_LITERAL, 0)
}

func (s *OperandContext) MEM() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserMEM, 0)
}

func (s *OperandContext) LabelID() ILabelIDContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelIDContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelIDContext)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.EnterOperand(s)
	}
}

func (s *OperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.ExitOperand(s)
	}
}

func (s *OperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZeonicaParserVisitor:
		return t.VisitOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZeonicaParser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ZeonicaParserRULE_operand)
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
			p.PredTag()
		}
		{
			p.SetState(136)
			p.Match(ZeonicaParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.IdList()
		}
		{
			p.SetState(138)
			p.Match(ZeonicaParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.Match(ZeonicaParserIMM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Match(ZeonicaParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.Match(ZeonicaParserDECIMAL_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.Match(ZeonicaParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(144)
			p.Match(ZeonicaParserIMM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)
			p.Match(ZeonicaParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.Match(ZeonicaParserOCT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Match(ZeonicaParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(148)
			p.Match(ZeonicaParserIMM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.Match(ZeonicaParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Match(ZeonicaParserHEX_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Match(ZeonicaParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(152)
			p.Match(ZeonicaParserIMM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.Match(ZeonicaParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Match(ZeonicaParserBINARY_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)
			p.Match(ZeonicaParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(156)
			p.Match(ZeonicaParserIMM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
			p.Match(ZeonicaParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Match(ZeonicaParserFLOAT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
			p.Match(ZeonicaParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(160)
			p.PredTag()
		}
		{
			p.SetState(161)
			p.Match(ZeonicaParserMEM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.Match(ZeonicaParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.IdList()
		}
		{
			p.SetState(164)
			p.Match(ZeonicaParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(166)
			p.PredTag()
		}
		{
			p.SetState(167)
			p.Match(ZeonicaParserMEM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.Match(ZeonicaParserLBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Match(ZeonicaParserHEX_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.Match(ZeonicaParserRBRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(172)
			p.LabelID()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdListContext is an interface to support dynamic dispatch.
type IIdListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsIdListContext differentiates from other interfaces.
	IsIdListContext()
}

type IdListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdListContext() *IdListContext {
	var p = new(IdListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_idList
	return p
}

func InitEmptyIdListContext(p *IdListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_idList
}

func (*IdListContext) IsIdListContext() {}

func NewIdListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdListContext {
	var p = new(IdListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZeonicaParserRULE_idList

	return p
}

func (s *IdListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdListContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ZeonicaParserIDENTIFIER)
}

func (s *IdListContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ZeonicaParserIDENTIFIER, i)
}

func (s *IdListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ZeonicaParserCOMMA)
}

func (s *IdListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ZeonicaParserCOMMA, i)
}

func (s *IdListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.EnterIdList(s)
	}
}

func (s *IdListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.ExitIdList(s)
	}
}

func (s *IdListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZeonicaParserVisitor:
		return t.VisitIdList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZeonicaParser) IdList() (localctx IIdListContext) {
	localctx = NewIdListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ZeonicaParserRULE_idList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(ZeonicaParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ZeonicaParserCOMMA {
		{
			p.SetState(176)
			p.Match(ZeonicaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.Match(ZeonicaParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPredTagContext is an interface to support dynamic dispatch.
type IPredTagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AND_PRED() antlr.TerminalNode
	OR_PRED() antlr.TerminalNode

	// IsPredTagContext differentiates from other interfaces.
	IsPredTagContext()
}

type PredTagContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredTagContext() *PredTagContext {
	var p = new(PredTagContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_predTag
	return p
}

func InitEmptyPredTagContext(p *PredTagContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_predTag
}

func (*PredTagContext) IsPredTagContext() {}

func NewPredTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredTagContext {
	var p = new(PredTagContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZeonicaParserRULE_predTag

	return p
}

func (s *PredTagContext) GetParser() antlr.Parser { return s.parser }

func (s *PredTagContext) AND_PRED() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserAND_PRED, 0)
}

func (s *PredTagContext) OR_PRED() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserOR_PRED, 0)
}

func (s *PredTagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredTagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredTagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.EnterPredTag(s)
	}
}

func (s *PredTagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.ExitPredTag(s)
	}
}

func (s *PredTagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZeonicaParserVisitor:
		return t.VisitPredTag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZeonicaParser) PredTag() (localctx IPredTagContext) {
	localctx = NewPredTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ZeonicaParserRULE_predTag)
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ZeonicaParserAND_PRED:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(183)
			p.Match(ZeonicaParserAND_PRED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ZeonicaParserOR_PRED:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(184)
			p.Match(ZeonicaParserOR_PRED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ZeonicaParserMEM, ZeonicaParserLBRACK:
		p.EnterOuterAlt(localctx, 3)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOpCodeContext is an interface to support dynamic dispatch.
type IOpCodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ADD() antlr.TerminalNode
	ADDI() antlr.TerminalNode
	SUB() antlr.TerminalNode
	SUBI() antlr.TerminalNode
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MAC() antlr.TerminalNode
	INC() antlr.TerminalNode
	LLS() antlr.TerminalNode
	LRS() antlr.TerminalNode
	OR() antlr.TerminalNode
	AND() antlr.TerminalNode
	XOR() antlr.TerminalNode
	NOT() antlr.TerminalNode
	FADD() antlr.TerminalNode
	FSUB() antlr.TerminalNode
	FMUL() antlr.TerminalNode
	FDIV() antlr.TerminalNode
	FMAC() antlr.TerminalNode
	MOV() antlr.TerminalNode
	MUL_CONST_ADD() antlr.TerminalNode
	NOP() antlr.TerminalNode
	CBR() antlr.TerminalNode
	CMOV() antlr.TerminalNode
	SGE() antlr.TerminalNode

	// IsOpCodeContext differentiates from other interfaces.
	IsOpCodeContext()
}

type OpCodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpCodeContext() *OpCodeContext {
	var p = new(OpCodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_opCode
	return p
}

func InitEmptyOpCodeContext(p *OpCodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZeonicaParserRULE_opCode
}

func (*OpCodeContext) IsOpCodeContext() {}

func NewOpCodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpCodeContext {
	var p = new(OpCodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZeonicaParserRULE_opCode

	return p
}

func (s *OpCodeContext) GetParser() antlr.Parser { return s.parser }

func (s *OpCodeContext) ADD() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserADD, 0)
}

func (s *OpCodeContext) ADDI() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserADDI, 0)
}

func (s *OpCodeContext) SUB() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserSUB, 0)
}

func (s *OpCodeContext) SUBI() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserSUBI, 0)
}

func (s *OpCodeContext) MUL() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserMUL, 0)
}

func (s *OpCodeContext) DIV() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserDIV, 0)
}

func (s *OpCodeContext) MAC() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserMAC, 0)
}

func (s *OpCodeContext) INC() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserINC, 0)
}

func (s *OpCodeContext) LLS() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserLLS, 0)
}

func (s *OpCodeContext) LRS() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserLRS, 0)
}

func (s *OpCodeContext) OR() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserOR, 0)
}

func (s *OpCodeContext) AND() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserAND, 0)
}

func (s *OpCodeContext) XOR() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserXOR, 0)
}

func (s *OpCodeContext) NOT() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserNOT, 0)
}

func (s *OpCodeContext) FADD() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserFADD, 0)
}

func (s *OpCodeContext) FSUB() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserFSUB, 0)
}

func (s *OpCodeContext) FMUL() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserFMUL, 0)
}

func (s *OpCodeContext) FDIV() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserFDIV, 0)
}

func (s *OpCodeContext) FMAC() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserFMAC, 0)
}

func (s *OpCodeContext) MOV() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserMOV, 0)
}

func (s *OpCodeContext) MUL_CONST_ADD() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserMUL_CONST_ADD, 0)
}

func (s *OpCodeContext) NOP() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserNOP, 0)
}

func (s *OpCodeContext) CBR() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserCBR, 0)
}

func (s *OpCodeContext) CMOV() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserCMOV, 0)
}

func (s *OpCodeContext) SGE() antlr.TerminalNode {
	return s.GetToken(ZeonicaParserSGE, 0)
}

func (s *OpCodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpCodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpCodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.EnterOpCode(s)
	}
}

func (s *OpCodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZeonicaParserListener); ok {
		listenerT.ExitOpCode(s)
	}
}

func (s *OpCodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZeonicaParserVisitor:
		return t.VisitOpCode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZeonicaParser) OpCode() (localctx IOpCodeContext) {
	localctx = NewOpCodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ZeonicaParserRULE_opCode)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1073741792) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
