// Code generated from ZeonicaParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ZeonicaParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ZeonicaParser.
type ZeonicaParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ZeonicaParser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

	// Visit a parse tree produced by ZeonicaParser#peBlock.
	VisitPeBlock(ctx *PeBlockContext) interface{}

	// Visit a parse tree produced by ZeonicaParser#FlatBody.
	VisitFlatBody(ctx *FlatBodyContext) interface{}

	// Visit a parse tree produced by ZeonicaParser#EntryBlockBody.
	VisitEntryBlockBody(ctx *EntryBlockBodyContext) interface{}

	// Visit a parse tree produced by ZeonicaParser#flatStyle.
	VisitFlatStyle(ctx *FlatStyleContext) interface{}

	// Visit a parse tree produced by ZeonicaParser#labeledGroup.
	VisitLabeledGroup(ctx *LabeledGroupContext) interface{}

	// Visit a parse tree produced by ZeonicaParser#entryBlock.
	VisitEntryBlock(ctx *EntryBlockContext) interface{}

	// Visit a parse tree produced by ZeonicaParser#loopType.
	VisitLoopType(ctx *LoopTypeContext) interface{}

	// Visit a parse tree produced by ZeonicaParser#instGroupList.
	VisitInstGroupList(ctx *InstGroupListContext) interface{}

	// Visit a parse tree produced by ZeonicaParser#instGroup.
	VisitInstGroup(ctx *InstGroupContext) interface{}

	// Visit a parse tree produced by ZeonicaParser#label.
	VisitLabel(ctx *LabelContext) interface{}

	// Visit a parse tree produced by ZeonicaParser#labelID.
	VisitLabelID(ctx *LabelIDContext) interface{}

	// Visit a parse tree produced by ZeonicaParser#normalInst.
	VisitNormalInst(ctx *NormalInstContext) interface{}

	// Visit a parse tree produced by ZeonicaParser#operandList.
	VisitOperandList(ctx *OperandListContext) interface{}

	// Visit a parse tree produced by ZeonicaParser#operand.
	VisitOperand(ctx *OperandContext) interface{}

	// Visit a parse tree produced by ZeonicaParser#idList.
	VisitIdList(ctx *IdListContext) interface{}

	// Visit a parse tree produced by ZeonicaParser#predTag.
	VisitPredTag(ctx *PredTagContext) interface{}

	// Visit a parse tree produced by ZeonicaParser#opCode.
	VisitOpCode(ctx *OpCodeContext) interface{}
}
