// Code generated from ZeonicaParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ZeonicaParser

import "github.com/antlr4-go/antlr/v4"

// BaseZeonicaParserListener is a complete listener for a parse tree produced by ZeonicaParser.
type BaseZeonicaParserListener struct{}

var _ ZeonicaParserListener = &BaseZeonicaParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseZeonicaParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseZeonicaParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseZeonicaParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseZeonicaParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseZeonicaParserListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseZeonicaParserListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterPeBlock is called when production peBlock is entered.
func (s *BaseZeonicaParserListener) EnterPeBlock(ctx *PeBlockContext) {}

// ExitPeBlock is called when production peBlock is exited.
func (s *BaseZeonicaParserListener) ExitPeBlock(ctx *PeBlockContext) {}

// EnterFlatBody is called when production FlatBody is entered.
func (s *BaseZeonicaParserListener) EnterFlatBody(ctx *FlatBodyContext) {}

// ExitFlatBody is called when production FlatBody is exited.
func (s *BaseZeonicaParserListener) ExitFlatBody(ctx *FlatBodyContext) {}

// EnterEntryBlockBody is called when production EntryBlockBody is entered.
func (s *BaseZeonicaParserListener) EnterEntryBlockBody(ctx *EntryBlockBodyContext) {}

// ExitEntryBlockBody is called when production EntryBlockBody is exited.
func (s *BaseZeonicaParserListener) ExitEntryBlockBody(ctx *EntryBlockBodyContext) {}

// EnterFlatStyle is called when production flatStyle is entered.
func (s *BaseZeonicaParserListener) EnterFlatStyle(ctx *FlatStyleContext) {}

// ExitFlatStyle is called when production flatStyle is exited.
func (s *BaseZeonicaParserListener) ExitFlatStyle(ctx *FlatStyleContext) {}

// EnterLabeledGroup is called when production labeledGroup is entered.
func (s *BaseZeonicaParserListener) EnterLabeledGroup(ctx *LabeledGroupContext) {}

// ExitLabeledGroup is called when production labeledGroup is exited.
func (s *BaseZeonicaParserListener) ExitLabeledGroup(ctx *LabeledGroupContext) {}

// EnterEntryBlock is called when production entryBlock is entered.
func (s *BaseZeonicaParserListener) EnterEntryBlock(ctx *EntryBlockContext) {}

// ExitEntryBlock is called when production entryBlock is exited.
func (s *BaseZeonicaParserListener) ExitEntryBlock(ctx *EntryBlockContext) {}

// EnterLoopType is called when production loopType is entered.
func (s *BaseZeonicaParserListener) EnterLoopType(ctx *LoopTypeContext) {}

// ExitLoopType is called when production loopType is exited.
func (s *BaseZeonicaParserListener) ExitLoopType(ctx *LoopTypeContext) {}

// EnterInstGroupList is called when production instGroupList is entered.
func (s *BaseZeonicaParserListener) EnterInstGroupList(ctx *InstGroupListContext) {}

// ExitInstGroupList is called when production instGroupList is exited.
func (s *BaseZeonicaParserListener) ExitInstGroupList(ctx *InstGroupListContext) {}

// EnterInstGroup is called when production instGroup is entered.
func (s *BaseZeonicaParserListener) EnterInstGroup(ctx *InstGroupContext) {}

// ExitInstGroup is called when production instGroup is exited.
func (s *BaseZeonicaParserListener) ExitInstGroup(ctx *InstGroupContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseZeonicaParserListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseZeonicaParserListener) ExitLabel(ctx *LabelContext) {}

// EnterLabelID is called when production labelID is entered.
func (s *BaseZeonicaParserListener) EnterLabelID(ctx *LabelIDContext) {}

// ExitLabelID is called when production labelID is exited.
func (s *BaseZeonicaParserListener) ExitLabelID(ctx *LabelIDContext) {}

// EnterNormalInst is called when production normalInst is entered.
func (s *BaseZeonicaParserListener) EnterNormalInst(ctx *NormalInstContext) {}

// ExitNormalInst is called when production normalInst is exited.
func (s *BaseZeonicaParserListener) ExitNormalInst(ctx *NormalInstContext) {}

// EnterOperandList is called when production operandList is entered.
func (s *BaseZeonicaParserListener) EnterOperandList(ctx *OperandListContext) {}

// ExitOperandList is called when production operandList is exited.
func (s *BaseZeonicaParserListener) ExitOperandList(ctx *OperandListContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseZeonicaParserListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseZeonicaParserListener) ExitOperand(ctx *OperandContext) {}

// EnterIdList is called when production idList is entered.
func (s *BaseZeonicaParserListener) EnterIdList(ctx *IdListContext) {}

// ExitIdList is called when production idList is exited.
func (s *BaseZeonicaParserListener) ExitIdList(ctx *IdListContext) {}

// EnterPredTag is called when production predTag is entered.
func (s *BaseZeonicaParserListener) EnterPredTag(ctx *PredTagContext) {}

// ExitPredTag is called when production predTag is exited.
func (s *BaseZeonicaParserListener) ExitPredTag(ctx *PredTagContext) {}

// EnterOpCode is called when production opCode is entered.
func (s *BaseZeonicaParserListener) EnterOpCode(ctx *OpCodeContext) {}

// ExitOpCode is called when production opCode is exited.
func (s *BaseZeonicaParserListener) ExitOpCode(ctx *OpCodeContext) {}
