// Code generated from ZeonicaParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ZeonicaParser

import "github.com/antlr4-go/antlr/v4"

// ZeonicaParserListener is a complete listener for a parse tree produced by ZeonicaParser.
type ZeonicaParserListener interface {
	antlr.ParseTreeListener

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterPeBlock is called when entering the peBlock production.
	EnterPeBlock(c *PeBlockContext)

	// EnterFlatBody is called when entering the FlatBody production.
	EnterFlatBody(c *FlatBodyContext)

	// EnterEntryBlockBody is called when entering the EntryBlockBody production.
	EnterEntryBlockBody(c *EntryBlockBodyContext)

	// EnterFlatStyle is called when entering the flatStyle production.
	EnterFlatStyle(c *FlatStyleContext)

	// EnterLabeledGroup is called when entering the labeledGroup production.
	EnterLabeledGroup(c *LabeledGroupContext)

	// EnterEntryBlock is called when entering the entryBlock production.
	EnterEntryBlock(c *EntryBlockContext)

	// EnterLoopType is called when entering the loopType production.
	EnterLoopType(c *LoopTypeContext)

	// EnterInstGroupList is called when entering the instGroupList production.
	EnterInstGroupList(c *InstGroupListContext)

	// EnterInstGroup is called when entering the instGroup production.
	EnterInstGroup(c *InstGroupContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterLabelID is called when entering the labelID production.
	EnterLabelID(c *LabelIDContext)

	// EnterNormalInst is called when entering the normalInst production.
	EnterNormalInst(c *NormalInstContext)

	// EnterOperandList is called when entering the operandList production.
	EnterOperandList(c *OperandListContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterIdList is called when entering the idList production.
	EnterIdList(c *IdListContext)

	// EnterPredTag is called when entering the predTag production.
	EnterPredTag(c *PredTagContext)

	// EnterOpCode is called when entering the opCode production.
	EnterOpCode(c *OpCodeContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitPeBlock is called when exiting the peBlock production.
	ExitPeBlock(c *PeBlockContext)

	// ExitFlatBody is called when exiting the FlatBody production.
	ExitFlatBody(c *FlatBodyContext)

	// ExitEntryBlockBody is called when exiting the EntryBlockBody production.
	ExitEntryBlockBody(c *EntryBlockBodyContext)

	// ExitFlatStyle is called when exiting the flatStyle production.
	ExitFlatStyle(c *FlatStyleContext)

	// ExitLabeledGroup is called when exiting the labeledGroup production.
	ExitLabeledGroup(c *LabeledGroupContext)

	// ExitEntryBlock is called when exiting the entryBlock production.
	ExitEntryBlock(c *EntryBlockContext)

	// ExitLoopType is called when exiting the loopType production.
	ExitLoopType(c *LoopTypeContext)

	// ExitInstGroupList is called when exiting the instGroupList production.
	ExitInstGroupList(c *InstGroupListContext)

	// ExitInstGroup is called when exiting the instGroup production.
	ExitInstGroup(c *InstGroupContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitLabelID is called when exiting the labelID production.
	ExitLabelID(c *LabelIDContext)

	// ExitNormalInst is called when exiting the normalInst production.
	ExitNormalInst(c *NormalInstContext)

	// ExitOperandList is called when exiting the operandList production.
	ExitOperandList(c *OperandListContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitIdList is called when exiting the idList production.
	ExitIdList(c *IdListContext)

	// ExitPredTag is called when exiting the predTag production.
	ExitPredTag(c *PredTagContext)

	// ExitOpCode is called when exiting the opCode production.
	ExitOpCode(c *OpCodeContext)
}
