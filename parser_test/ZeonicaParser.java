// Generated from ZeonicaParser.g4 by ANTLR 4.13.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class ZeonicaParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		PE=1, LOOP=2, ONCE=3, ADD=4, ADDI=5, SUB=6, SUBI=7, MUL=8, DIV=9, MAC=10, 
		INC=11, LLS=12, LRS=13, OR=14, AND=15, XOR=16, NOT=17, FADD=18, FSUB=19, 
		FMUL=20, FDIV=21, FMAC=22, BEQ=23, MOV=24, MUL_CONST_ADD=25, NOP=26, CBR=27, 
		CMOV=28, SGE=29, LOAD=30, STORE=31, ENTRY=32, ENTRY_ARROW=33, RIGHT_ARROW=34, 
		IMM=35, MEM=36, DECIMAL_LITERAL=37, HEX_LITERAL=38, OCT_LITERAL=39, BINARY_LITERAL=40, 
		FLOAT_LITERAL=41, HEX_FLOAT_LITERAL=42, BOOL_LITERAL=43, CHAR_LITERAL=44, 
		STRING_LITERAL=45, TEXT_BLOCK=46, NULL_LITERAL=47, LPAREN=48, RPAREN=49, 
		LBRACE=50, RBRACE=51, LBRACK=52, RBRACK=53, SEMI=54, COMMA=55, DOT=56, 
		COLON=57, AND_PRED=58, OR_PRED=59, WS=60, COMMENT=61, LINE_COMMENT=62, 
		IDENTIFIER=63;
	public static final int
		RULE_compilationUnit = 0, RULE_peBlock = 1, RULE_loopType = 2, RULE_instList = 3, 
		RULE_inst = 4, RULE_label = 5, RULE_labelID = 6, RULE_normalInst = 7, 
		RULE_operandList = 8, RULE_operand = 9, RULE_idList = 10, RULE_predTag = 11, 
		RULE_opCode = 12;
	private static String[] makeRuleNames() {
		return new String[] {
			"compilationUnit", "peBlock", "loopType", "instList", "inst", "label", 
			"labelID", "normalInst", "operandList", "operand", "idList", "predTag", 
			"opCode"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'PE'", "'Loop'", "'Once'", "'ADD'", "'ADDI'", "'SUB'", "'SUBI'", 
			"'MUL'", "'DIV'", "'MAC'", "'INC'", "'LLS'", "'LRS'", "'OR'", "'AND'", 
			"'XOR'", "'NOT'", "'FADD'", "'FSUB'", "'FMUL'", "'FDIV'", "'FMAC'", "'BEQ'", 
			"'MOV'", "'MUL_CONST_ADD'", "'NOP'", "'CBR'", "'CMOV'", "'SGE'", "'LOAD'", 
			"'STORE'", "'Entry'", "'=>'", "'->'", "'IMM'", "'MEM'", null, null, null, 
			null, null, null, null, null, null, null, "'null'", "'('", "')'", "'{'", 
			"'}'", "'['", "']'", "';'", "','", "'.'", "':'", "'!'", "'^'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PE", "LOOP", "ONCE", "ADD", "ADDI", "SUB", "SUBI", "MUL", "DIV", 
			"MAC", "INC", "LLS", "LRS", "OR", "AND", "XOR", "NOT", "FADD", "FSUB", 
			"FMUL", "FDIV", "FMAC", "BEQ", "MOV", "MUL_CONST_ADD", "NOP", "CBR", 
			"CMOV", "SGE", "LOAD", "STORE", "ENTRY", "ENTRY_ARROW", "RIGHT_ARROW", 
			"IMM", "MEM", "DECIMAL_LITERAL", "HEX_LITERAL", "OCT_LITERAL", "BINARY_LITERAL", 
			"FLOAT_LITERAL", "HEX_FLOAT_LITERAL", "BOOL_LITERAL", "CHAR_LITERAL", 
			"STRING_LITERAL", "TEXT_BLOCK", "NULL_LITERAL", "LPAREN", "RPAREN", "LBRACE", 
			"RBRACE", "LBRACK", "RBRACK", "SEMI", "COMMA", "DOT", "COLON", "AND_PRED", 
			"OR_PRED", "WS", "COMMENT", "LINE_COMMENT", "IDENTIFIER"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "ZeonicaParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public ZeonicaParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CompilationUnitContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(ZeonicaParser.EOF, 0); }
		public List<PeBlockContext> peBlock() {
			return getRuleContexts(PeBlockContext.class);
		}
		public PeBlockContext peBlock(int i) {
			return getRuleContext(PeBlockContext.class,i);
		}
		public CompilationUnitContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_compilationUnit; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).enterCompilationUnit(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).exitCompilationUnit(this);
		}
	}

	public final CompilationUnitContext compilationUnit() throws RecognitionException {
		CompilationUnitContext _localctx = new CompilationUnitContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_compilationUnit);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(29);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==PE) {
				{
				{
				setState(26);
				peBlock();
				}
				}
				setState(31);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(32);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PeBlockContext extends ParserRuleContext {
		public TerminalNode PE() { return getToken(ZeonicaParser.PE, 0); }
		public TerminalNode LPAREN() { return getToken(ZeonicaParser.LPAREN, 0); }
		public List<TerminalNode> DECIMAL_LITERAL() { return getTokens(ZeonicaParser.DECIMAL_LITERAL); }
		public TerminalNode DECIMAL_LITERAL(int i) {
			return getToken(ZeonicaParser.DECIMAL_LITERAL, i);
		}
		public TerminalNode RPAREN() { return getToken(ZeonicaParser.RPAREN, 0); }
		public TerminalNode ENTRY_ARROW() { return getToken(ZeonicaParser.ENTRY_ARROW, 0); }
		public LoopTypeContext loopType() {
			return getRuleContext(LoopTypeContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(ZeonicaParser.LBRACE, 0); }
		public InstListContext instList() {
			return getRuleContext(InstListContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(ZeonicaParser.RBRACE, 0); }
		public TerminalNode COMMA() { return getToken(ZeonicaParser.COMMA, 0); }
		public PeBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_peBlock; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).enterPeBlock(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).exitPeBlock(this);
		}
	}

	public final PeBlockContext peBlock() throws RecognitionException {
		PeBlockContext _localctx = new PeBlockContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_peBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(34);
			match(PE);
			setState(35);
			match(LPAREN);
			setState(36);
			match(DECIMAL_LITERAL);
			setState(39);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(37);
				match(COMMA);
				setState(38);
				match(DECIMAL_LITERAL);
				}
			}

			setState(41);
			match(RPAREN);
			setState(42);
			match(ENTRY_ARROW);
			setState(43);
			loopType();
			setState(44);
			match(LBRACE);
			setState(45);
			instList();
			setState(46);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LoopTypeContext extends ParserRuleContext {
		public TerminalNode LOOP() { return getToken(ZeonicaParser.LOOP, 0); }
		public TerminalNode ONCE() { return getToken(ZeonicaParser.ONCE, 0); }
		public LoopTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_loopType; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).enterLoopType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).exitLoopType(this);
		}
	}

	public final LoopTypeContext loopType() throws RecognitionException {
		LoopTypeContext _localctx = new LoopTypeContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_loopType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(48);
			_la = _input.LA(1);
			if ( !(_la==LOOP || _la==ONCE) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InstListContext extends ParserRuleContext {
		public List<InstContext> inst() {
			return getRuleContexts(InstContext.class);
		}
		public InstContext inst(int i) {
			return getRuleContext(InstContext.class,i);
		}
		public InstListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).enterInstList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).exitInstList(this);
		}
	}

	public final InstListContext instList() throws RecognitionException {
		InstListContext _localctx = new InstListContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_instList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(53);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==LBRACE || _la==IDENTIFIER) {
				{
				{
				setState(50);
				inst();
				}
				}
				setState(55);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InstContext extends ParserRuleContext {
		public TerminalNode LBRACE() { return getToken(ZeonicaParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(ZeonicaParser.RBRACE, 0); }
		public LabelContext label() {
			return getRuleContext(LabelContext.class,0);
		}
		public List<NormalInstContext> normalInst() {
			return getRuleContexts(NormalInstContext.class);
		}
		public NormalInstContext normalInst(int i) {
			return getRuleContext(NormalInstContext.class,i);
		}
		public InstContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_inst; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).enterInst(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).exitInst(this);
		}
	}

	public final InstContext inst() throws RecognitionException {
		InstContext _localctx = new InstContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_inst);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(57);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(56);
				label();
				}
			}

			setState(59);
			match(LBRACE);
			setState(63);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & -8354177204627701776L) != 0)) {
				{
				{
				setState(60);
				normalInst();
				}
				}
				setState(65);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(66);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LabelContext extends ParserRuleContext {
		public LabelIDContext labelID() {
			return getRuleContext(LabelIDContext.class,0);
		}
		public TerminalNode COLON() { return getToken(ZeonicaParser.COLON, 0); }
		public LabelContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_label; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).enterLabel(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).exitLabel(this);
		}
	}

	public final LabelContext label() throws RecognitionException {
		LabelContext _localctx = new LabelContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_label);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(68);
			labelID();
			setState(69);
			match(COLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LabelIDContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(ZeonicaParser.IDENTIFIER, 0); }
		public LabelIDContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_labelID; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).enterLabelID(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).exitLabelID(this);
		}
	}

	public final LabelIDContext labelID() throws RecognitionException {
		LabelIDContext _localctx = new LabelIDContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_labelID);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(71);
			match(IDENTIFIER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class NormalInstContext extends ParserRuleContext {
		public OpCodeContext opCode() {
			return getRuleContext(OpCodeContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(ZeonicaParser.COMMA, 0); }
		public List<OperandListContext> operandList() {
			return getRuleContexts(OperandListContext.class);
		}
		public OperandListContext operandList(int i) {
			return getRuleContext(OperandListContext.class,i);
		}
		public TerminalNode RIGHT_ARROW() { return getToken(ZeonicaParser.RIGHT_ARROW, 0); }
		public List<OperandContext> operand() {
			return getRuleContexts(OperandContext.class);
		}
		public OperandContext operand(int i) {
			return getRuleContext(OperandContext.class,i);
		}
		public NormalInstContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_normalInst; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).enterNormalInst(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).exitNormalInst(this);
		}
	}

	public final NormalInstContext normalInst() throws RecognitionException {
		NormalInstContext _localctx = new NormalInstContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_normalInst);
		try {
			setState(88);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(73);
				opCode();
				setState(74);
				match(COMMA);
				setState(75);
				operandList();
				setState(76);
				match(RIGHT_ARROW);
				setState(77);
				operandList();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(79);
				opCode();
				setState(80);
				match(COMMA);
				setState(81);
				operandList();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(83);
				opCode();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(84);
				operand();
				setState(85);
				match(RIGHT_ARROW);
				setState(86);
				operand();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OperandListContext extends ParserRuleContext {
		public List<OperandContext> operand() {
			return getRuleContexts(OperandContext.class);
		}
		public OperandContext operand(int i) {
			return getRuleContext(OperandContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(ZeonicaParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(ZeonicaParser.COMMA, i);
		}
		public OperandListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operandList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).enterOperandList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).exitOperandList(this);
		}
	}

	public final OperandListContext operandList() throws RecognitionException {
		OperandListContext _localctx = new OperandListContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_operandList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(90);
			operand();
			setState(95);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(91);
				match(COMMA);
				setState(92);
				operand();
				}
				}
				setState(97);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OperandContext extends ParserRuleContext {
		public PredTagContext predTag() {
			return getRuleContext(PredTagContext.class,0);
		}
		public TerminalNode LBRACK() { return getToken(ZeonicaParser.LBRACK, 0); }
		public IdListContext idList() {
			return getRuleContext(IdListContext.class,0);
		}
		public TerminalNode RBRACK() { return getToken(ZeonicaParser.RBRACK, 0); }
		public TerminalNode IMM() { return getToken(ZeonicaParser.IMM, 0); }
		public TerminalNode DECIMAL_LITERAL() { return getToken(ZeonicaParser.DECIMAL_LITERAL, 0); }
		public TerminalNode OCT_LITERAL() { return getToken(ZeonicaParser.OCT_LITERAL, 0); }
		public TerminalNode HEX_LITERAL() { return getToken(ZeonicaParser.HEX_LITERAL, 0); }
		public TerminalNode BINARY_LITERAL() { return getToken(ZeonicaParser.BINARY_LITERAL, 0); }
		public TerminalNode FLOAT_LITERAL() { return getToken(ZeonicaParser.FLOAT_LITERAL, 0); }
		public TerminalNode MEM() { return getToken(ZeonicaParser.MEM, 0); }
		public LabelIDContext labelID() {
			return getRuleContext(LabelIDContext.class,0);
		}
		public OperandContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operand; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).enterOperand(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).exitOperand(this);
		}
	}

	public final OperandContext operand() throws RecognitionException {
		OperandContext _localctx = new OperandContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_operand);
		try {
			setState(136);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(98);
				predTag();
				setState(99);
				match(LBRACK);
				setState(100);
				idList();
				setState(101);
				match(RBRACK);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(103);
				match(IMM);
				setState(104);
				match(LBRACK);
				setState(105);
				match(DECIMAL_LITERAL);
				setState(106);
				match(RBRACK);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(107);
				match(IMM);
				setState(108);
				match(LBRACK);
				setState(109);
				match(OCT_LITERAL);
				setState(110);
				match(RBRACK);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(111);
				match(IMM);
				setState(112);
				match(LBRACK);
				setState(113);
				match(HEX_LITERAL);
				setState(114);
				match(RBRACK);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(115);
				match(IMM);
				setState(116);
				match(LBRACK);
				setState(117);
				match(BINARY_LITERAL);
				setState(118);
				match(RBRACK);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(119);
				match(IMM);
				setState(120);
				match(LBRACK);
				setState(121);
				match(FLOAT_LITERAL);
				setState(122);
				match(RBRACK);
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(123);
				predTag();
				setState(124);
				match(MEM);
				setState(125);
				match(LBRACK);
				setState(126);
				idList();
				setState(127);
				match(RBRACK);
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(129);
				predTag();
				setState(130);
				match(MEM);
				setState(131);
				match(LBRACK);
				setState(132);
				match(HEX_LITERAL);
				setState(133);
				match(RBRACK);
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(135);
				labelID();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IdListContext extends ParserRuleContext {
		public List<TerminalNode> IDENTIFIER() { return getTokens(ZeonicaParser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(ZeonicaParser.IDENTIFIER, i);
		}
		public List<TerminalNode> COMMA() { return getTokens(ZeonicaParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(ZeonicaParser.COMMA, i);
		}
		public IdListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_idList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).enterIdList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).exitIdList(this);
		}
	}

	public final IdListContext idList() throws RecognitionException {
		IdListContext _localctx = new IdListContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_idList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(138);
			match(IDENTIFIER);
			setState(143);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(139);
				match(COMMA);
				setState(140);
				match(IDENTIFIER);
				}
				}
				setState(145);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PredTagContext extends ParserRuleContext {
		public TerminalNode AND_PRED() { return getToken(ZeonicaParser.AND_PRED, 0); }
		public TerminalNode OR_PRED() { return getToken(ZeonicaParser.OR_PRED, 0); }
		public PredTagContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_predTag; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).enterPredTag(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).exitPredTag(this);
		}
	}

	public final PredTagContext predTag() throws RecognitionException {
		PredTagContext _localctx = new PredTagContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_predTag);
		try {
			setState(149);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case AND_PRED:
				enterOuterAlt(_localctx, 1);
				{
				setState(146);
				match(AND_PRED);
				}
				break;
			case OR_PRED:
				enterOuterAlt(_localctx, 2);
				{
				setState(147);
				match(OR_PRED);
				}
				break;
			case MEM:
			case LBRACK:
				enterOuterAlt(_localctx, 3);
				{
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OpCodeContext extends ParserRuleContext {
		public TerminalNode ADD() { return getToken(ZeonicaParser.ADD, 0); }
		public TerminalNode ADDI() { return getToken(ZeonicaParser.ADDI, 0); }
		public TerminalNode SUB() { return getToken(ZeonicaParser.SUB, 0); }
		public TerminalNode SUBI() { return getToken(ZeonicaParser.SUBI, 0); }
		public TerminalNode MUL() { return getToken(ZeonicaParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(ZeonicaParser.DIV, 0); }
		public TerminalNode MAC() { return getToken(ZeonicaParser.MAC, 0); }
		public TerminalNode INC() { return getToken(ZeonicaParser.INC, 0); }
		public TerminalNode LLS() { return getToken(ZeonicaParser.LLS, 0); }
		public TerminalNode LRS() { return getToken(ZeonicaParser.LRS, 0); }
		public TerminalNode OR() { return getToken(ZeonicaParser.OR, 0); }
		public TerminalNode AND() { return getToken(ZeonicaParser.AND, 0); }
		public TerminalNode XOR() { return getToken(ZeonicaParser.XOR, 0); }
		public TerminalNode NOT() { return getToken(ZeonicaParser.NOT, 0); }
		public TerminalNode FADD() { return getToken(ZeonicaParser.FADD, 0); }
		public TerminalNode FSUB() { return getToken(ZeonicaParser.FSUB, 0); }
		public TerminalNode FMUL() { return getToken(ZeonicaParser.FMUL, 0); }
		public TerminalNode FDIV() { return getToken(ZeonicaParser.FDIV, 0); }
		public TerminalNode FMAC() { return getToken(ZeonicaParser.FMAC, 0); }
		public TerminalNode MOV() { return getToken(ZeonicaParser.MOV, 0); }
		public TerminalNode MUL_CONST_ADD() { return getToken(ZeonicaParser.MUL_CONST_ADD, 0); }
		public TerminalNode NOP() { return getToken(ZeonicaParser.NOP, 0); }
		public TerminalNode CBR() { return getToken(ZeonicaParser.CBR, 0); }
		public TerminalNode CMOV() { return getToken(ZeonicaParser.CMOV, 0); }
		public TerminalNode SGE() { return getToken(ZeonicaParser.SGE, 0); }
		public OpCodeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_opCode; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).enterOpCode(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ZeonicaParserListener ) ((ZeonicaParserListener)listener).exitOpCode(this);
		}
	}

	public final OpCodeContext opCode() throws RecognitionException {
		OpCodeContext _localctx = new OpCodeContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_opCode);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(151);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 1065353200L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\u0004\u0001?\u009a\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0001\u0000\u0005\u0000\u001c\b\u0000\n\u0000\f\u0000\u001f"+
		"\t\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0003\u0001(\b\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0002\u0001"+
		"\u0002\u0001\u0003\u0005\u00034\b\u0003\n\u0003\f\u00037\t\u0003\u0001"+
		"\u0004\u0003\u0004:\b\u0004\u0001\u0004\u0001\u0004\u0005\u0004>\b\u0004"+
		"\n\u0004\f\u0004A\t\u0004\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0003\u0007Y\b\u0007\u0001\b\u0001\b\u0001\b\u0005\b^\b\b\n\b\f\ba\t"+
		"\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0003\t\u0089\b\t\u0001\n\u0001\n\u0001\n\u0005\n\u008e"+
		"\b\n\n\n\f\n\u0091\t\n\u0001\u000b\u0001\u000b\u0001\u000b\u0003\u000b"+
		"\u0096\b\u000b\u0001\f\u0001\f\u0001\f\u0000\u0000\r\u0000\u0002\u0004"+
		"\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u0000\u0002\u0001\u0000"+
		"\u0002\u0003\u0002\u0000\u0004\u0016\u0018\u001d\u00a0\u0000\u001d\u0001"+
		"\u0000\u0000\u0000\u0002\"\u0001\u0000\u0000\u0000\u00040\u0001\u0000"+
		"\u0000\u0000\u00065\u0001\u0000\u0000\u0000\b9\u0001\u0000\u0000\u0000"+
		"\nD\u0001\u0000\u0000\u0000\fG\u0001\u0000\u0000\u0000\u000eX\u0001\u0000"+
		"\u0000\u0000\u0010Z\u0001\u0000\u0000\u0000\u0012\u0088\u0001\u0000\u0000"+
		"\u0000\u0014\u008a\u0001\u0000\u0000\u0000\u0016\u0095\u0001\u0000\u0000"+
		"\u0000\u0018\u0097\u0001\u0000\u0000\u0000\u001a\u001c\u0003\u0002\u0001"+
		"\u0000\u001b\u001a\u0001\u0000\u0000\u0000\u001c\u001f\u0001\u0000\u0000"+
		"\u0000\u001d\u001b\u0001\u0000\u0000\u0000\u001d\u001e\u0001\u0000\u0000"+
		"\u0000\u001e \u0001\u0000\u0000\u0000\u001f\u001d\u0001\u0000\u0000\u0000"+
		" !\u0005\u0000\u0000\u0001!\u0001\u0001\u0000\u0000\u0000\"#\u0005\u0001"+
		"\u0000\u0000#$\u00050\u0000\u0000$\'\u0005%\u0000\u0000%&\u00057\u0000"+
		"\u0000&(\u0005%\u0000\u0000\'%\u0001\u0000\u0000\u0000\'(\u0001\u0000"+
		"\u0000\u0000()\u0001\u0000\u0000\u0000)*\u00051\u0000\u0000*+\u0005!\u0000"+
		"\u0000+,\u0003\u0004\u0002\u0000,-\u00052\u0000\u0000-.\u0003\u0006\u0003"+
		"\u0000./\u00053\u0000\u0000/\u0003\u0001\u0000\u0000\u000001\u0007\u0000"+
		"\u0000\u00001\u0005\u0001\u0000\u0000\u000024\u0003\b\u0004\u000032\u0001"+
		"\u0000\u0000\u000047\u0001\u0000\u0000\u000053\u0001\u0000\u0000\u0000"+
		"56\u0001\u0000\u0000\u00006\u0007\u0001\u0000\u0000\u000075\u0001\u0000"+
		"\u0000\u00008:\u0003\n\u0005\u000098\u0001\u0000\u0000\u00009:\u0001\u0000"+
		"\u0000\u0000:;\u0001\u0000\u0000\u0000;?\u00052\u0000\u0000<>\u0003\u000e"+
		"\u0007\u0000=<\u0001\u0000\u0000\u0000>A\u0001\u0000\u0000\u0000?=\u0001"+
		"\u0000\u0000\u0000?@\u0001\u0000\u0000\u0000@B\u0001\u0000\u0000\u0000"+
		"A?\u0001\u0000\u0000\u0000BC\u00053\u0000\u0000C\t\u0001\u0000\u0000\u0000"+
		"DE\u0003\f\u0006\u0000EF\u00059\u0000\u0000F\u000b\u0001\u0000\u0000\u0000"+
		"GH\u0005?\u0000\u0000H\r\u0001\u0000\u0000\u0000IJ\u0003\u0018\f\u0000"+
		"JK\u00057\u0000\u0000KL\u0003\u0010\b\u0000LM\u0005\"\u0000\u0000MN\u0003"+
		"\u0010\b\u0000NY\u0001\u0000\u0000\u0000OP\u0003\u0018\f\u0000PQ\u0005"+
		"7\u0000\u0000QR\u0003\u0010\b\u0000RY\u0001\u0000\u0000\u0000SY\u0003"+
		"\u0018\f\u0000TU\u0003\u0012\t\u0000UV\u0005\"\u0000\u0000VW\u0003\u0012"+
		"\t\u0000WY\u0001\u0000\u0000\u0000XI\u0001\u0000\u0000\u0000XO\u0001\u0000"+
		"\u0000\u0000XS\u0001\u0000\u0000\u0000XT\u0001\u0000\u0000\u0000Y\u000f"+
		"\u0001\u0000\u0000\u0000Z_\u0003\u0012\t\u0000[\\\u00057\u0000\u0000\\"+
		"^\u0003\u0012\t\u0000][\u0001\u0000\u0000\u0000^a\u0001\u0000\u0000\u0000"+
		"_]\u0001\u0000\u0000\u0000_`\u0001\u0000\u0000\u0000`\u0011\u0001\u0000"+
		"\u0000\u0000a_\u0001\u0000\u0000\u0000bc\u0003\u0016\u000b\u0000cd\u0005"+
		"4\u0000\u0000de\u0003\u0014\n\u0000ef\u00055\u0000\u0000f\u0089\u0001"+
		"\u0000\u0000\u0000gh\u0005#\u0000\u0000hi\u00054\u0000\u0000ij\u0005%"+
		"\u0000\u0000j\u0089\u00055\u0000\u0000kl\u0005#\u0000\u0000lm\u00054\u0000"+
		"\u0000mn\u0005\'\u0000\u0000n\u0089\u00055\u0000\u0000op\u0005#\u0000"+
		"\u0000pq\u00054\u0000\u0000qr\u0005&\u0000\u0000r\u0089\u00055\u0000\u0000"+
		"st\u0005#\u0000\u0000tu\u00054\u0000\u0000uv\u0005(\u0000\u0000v\u0089"+
		"\u00055\u0000\u0000wx\u0005#\u0000\u0000xy\u00054\u0000\u0000yz\u0005"+
		")\u0000\u0000z\u0089\u00055\u0000\u0000{|\u0003\u0016\u000b\u0000|}\u0005"+
		"$\u0000\u0000}~\u00054\u0000\u0000~\u007f\u0003\u0014\n\u0000\u007f\u0080"+
		"\u00055\u0000\u0000\u0080\u0089\u0001\u0000\u0000\u0000\u0081\u0082\u0003"+
		"\u0016\u000b\u0000\u0082\u0083\u0005$\u0000\u0000\u0083\u0084\u00054\u0000"+
		"\u0000\u0084\u0085\u0005&\u0000\u0000\u0085\u0086\u00055\u0000\u0000\u0086"+
		"\u0089\u0001\u0000\u0000\u0000\u0087\u0089\u0003\f\u0006\u0000\u0088b"+
		"\u0001\u0000\u0000\u0000\u0088g\u0001\u0000\u0000\u0000\u0088k\u0001\u0000"+
		"\u0000\u0000\u0088o\u0001\u0000\u0000\u0000\u0088s\u0001\u0000\u0000\u0000"+
		"\u0088w\u0001\u0000\u0000\u0000\u0088{\u0001\u0000\u0000\u0000\u0088\u0081"+
		"\u0001\u0000\u0000\u0000\u0088\u0087\u0001\u0000\u0000\u0000\u0089\u0013"+
		"\u0001\u0000\u0000\u0000\u008a\u008f\u0005?\u0000\u0000\u008b\u008c\u0005"+
		"7\u0000\u0000\u008c\u008e\u0005?\u0000\u0000\u008d\u008b\u0001\u0000\u0000"+
		"\u0000\u008e\u0091\u0001\u0000\u0000\u0000\u008f\u008d\u0001\u0000\u0000"+
		"\u0000\u008f\u0090\u0001\u0000\u0000\u0000\u0090\u0015\u0001\u0000\u0000"+
		"\u0000\u0091\u008f\u0001\u0000\u0000\u0000\u0092\u0096\u0005:\u0000\u0000"+
		"\u0093\u0096\u0005;\u0000\u0000\u0094\u0096\u0001\u0000\u0000\u0000\u0095"+
		"\u0092\u0001\u0000\u0000\u0000\u0095\u0093\u0001\u0000\u0000\u0000\u0095"+
		"\u0094\u0001\u0000\u0000\u0000\u0096\u0017\u0001\u0000\u0000\u0000\u0097"+
		"\u0098\u0007\u0001\u0000\u0000\u0098\u0019\u0001\u0000\u0000\u0000\n\u001d"+
		"\'59?X_\u0088\u008f\u0095";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}