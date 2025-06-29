// Generated from ZeonicaParser.g4 by ANTLR 4.13.2
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link ZeonicaParser}.
 */
public interface ZeonicaParserListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link ZeonicaParser#compilationUnit}.
	 * @param ctx the parse tree
	 */
	void enterCompilationUnit(ZeonicaParser.CompilationUnitContext ctx);
	/**
	 * Exit a parse tree produced by {@link ZeonicaParser#compilationUnit}.
	 * @param ctx the parse tree
	 */
	void exitCompilationUnit(ZeonicaParser.CompilationUnitContext ctx);
	/**
	 * Enter a parse tree produced by {@link ZeonicaParser#peBlock}.
	 * @param ctx the parse tree
	 */
	void enterPeBlock(ZeonicaParser.PeBlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link ZeonicaParser#peBlock}.
	 * @param ctx the parse tree
	 */
	void exitPeBlock(ZeonicaParser.PeBlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link ZeonicaParser#loopType}.
	 * @param ctx the parse tree
	 */
	void enterLoopType(ZeonicaParser.LoopTypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link ZeonicaParser#loopType}.
	 * @param ctx the parse tree
	 */
	void exitLoopType(ZeonicaParser.LoopTypeContext ctx);
	/**
	 * Enter a parse tree produced by {@link ZeonicaParser#instList}.
	 * @param ctx the parse tree
	 */
	void enterInstList(ZeonicaParser.InstListContext ctx);
	/**
	 * Exit a parse tree produced by {@link ZeonicaParser#instList}.
	 * @param ctx the parse tree
	 */
	void exitInstList(ZeonicaParser.InstListContext ctx);
	/**
	 * Enter a parse tree produced by {@link ZeonicaParser#inst}.
	 * @param ctx the parse tree
	 */
	void enterInst(ZeonicaParser.InstContext ctx);
	/**
	 * Exit a parse tree produced by {@link ZeonicaParser#inst}.
	 * @param ctx the parse tree
	 */
	void exitInst(ZeonicaParser.InstContext ctx);
	/**
	 * Enter a parse tree produced by {@link ZeonicaParser#label}.
	 * @param ctx the parse tree
	 */
	void enterLabel(ZeonicaParser.LabelContext ctx);
	/**
	 * Exit a parse tree produced by {@link ZeonicaParser#label}.
	 * @param ctx the parse tree
	 */
	void exitLabel(ZeonicaParser.LabelContext ctx);
	/**
	 * Enter a parse tree produced by {@link ZeonicaParser#labelID}.
	 * @param ctx the parse tree
	 */
	void enterLabelID(ZeonicaParser.LabelIDContext ctx);
	/**
	 * Exit a parse tree produced by {@link ZeonicaParser#labelID}.
	 * @param ctx the parse tree
	 */
	void exitLabelID(ZeonicaParser.LabelIDContext ctx);
	/**
	 * Enter a parse tree produced by {@link ZeonicaParser#normalInst}.
	 * @param ctx the parse tree
	 */
	void enterNormalInst(ZeonicaParser.NormalInstContext ctx);
	/**
	 * Exit a parse tree produced by {@link ZeonicaParser#normalInst}.
	 * @param ctx the parse tree
	 */
	void exitNormalInst(ZeonicaParser.NormalInstContext ctx);
	/**
	 * Enter a parse tree produced by {@link ZeonicaParser#operandList}.
	 * @param ctx the parse tree
	 */
	void enterOperandList(ZeonicaParser.OperandListContext ctx);
	/**
	 * Exit a parse tree produced by {@link ZeonicaParser#operandList}.
	 * @param ctx the parse tree
	 */
	void exitOperandList(ZeonicaParser.OperandListContext ctx);
	/**
	 * Enter a parse tree produced by {@link ZeonicaParser#operand}.
	 * @param ctx the parse tree
	 */
	void enterOperand(ZeonicaParser.OperandContext ctx);
	/**
	 * Exit a parse tree produced by {@link ZeonicaParser#operand}.
	 * @param ctx the parse tree
	 */
	void exitOperand(ZeonicaParser.OperandContext ctx);
	/**
	 * Enter a parse tree produced by {@link ZeonicaParser#idList}.
	 * @param ctx the parse tree
	 */
	void enterIdList(ZeonicaParser.IdListContext ctx);
	/**
	 * Exit a parse tree produced by {@link ZeonicaParser#idList}.
	 * @param ctx the parse tree
	 */
	void exitIdList(ZeonicaParser.IdListContext ctx);
	/**
	 * Enter a parse tree produced by {@link ZeonicaParser#predTag}.
	 * @param ctx the parse tree
	 */
	void enterPredTag(ZeonicaParser.PredTagContext ctx);
	/**
	 * Exit a parse tree produced by {@link ZeonicaParser#predTag}.
	 * @param ctx the parse tree
	 */
	void exitPredTag(ZeonicaParser.PredTagContext ctx);
	/**
	 * Enter a parse tree produced by {@link ZeonicaParser#opCode}.
	 * @param ctx the parse tree
	 */
	void enterOpCode(ZeonicaParser.OpCodeContext ctx);
	/**
	 * Exit a parse tree produced by {@link ZeonicaParser#opCode}.
	 * @param ctx the parse tree
	 */
	void exitOpCode(ZeonicaParser.OpCodeContext ctx);
}