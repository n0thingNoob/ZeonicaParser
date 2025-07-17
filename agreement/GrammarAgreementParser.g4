parser grammar GrammarAgreementParser;


options {
    tokenVocab = GrammarAgreementLexer;
}

compilationUnit 
    : pe* EOF
    ;
    
pe
    : PE_LIT LPAREN DECIMAL_LITERAL (COMMA DECIMAL_LITERAL)? RPAREN RIGHT_ARROW LBRACE entryBlock* RBRACE
    ;

entryBlock
    : entryHead LBRACE instList RBRACE
    ;

entryHead
    : ENTRY ENTRY_ARROW exitOption
    | ENTRY operandList ENTRY_ARROW exitOption
    ;

exitOption
    : ONCE
    | // default is LOOP, means if enter this entry, engine will find a entry to run after it ends.
    ;

instList
    : inst*
    ;

inst
    : LBRACE normalInst* RBRACE
    | normalInst
    | label
    ;

label
    : labelID COLON
    ;

labelID
    : IDENTIFIER
    ;


normalInst
    : opCode COMMA operandList (RIGHT_ARROW operandList)?  // normal instruction.
    | operand RIGHT_ARROW operand // syntax sugar for mov. // no dst operand.
    ;

// operandList is like  [],[],[],[]
operandList
    : operand (COMMA operand)*
    ;

operand
    : predTag LBRACK idList RBRACK
    | IMM LBRACK DECIMAL_LITERAL RBRACK
    | IMM LBRACK OCT_LITERAL RBRACK
    | IMM LBRACK HEX_LITERAL RBRACK
    | IMM LBRACK BINARY_LITERAL RBRACK
    | IMM LBRACK FLOAT_LITERAL RBRACK
    | predTag MEM LBRACK idList RBRACK
    | predTag MEM LBRACK HEX_LITERAL RBRACK // MEM
    | labelID
    ;

idList
    : IDENTIFIER (COMMA IDENTIFIER)*
    ;

predTag
    : AND_PRED
    | OR_PRED
    | 
    ;

// it will be finalized after the opcodes are finalized.
opCode
    : ADD
    | ADDI
    | SUB
    | SUBI
    | MUL
    | DIV
    | MAC
    | INC
    | LLS
    | LRS
    | OR
    | AND
    | XOR
    | NOT
    | FADD
    | FSUB
    | FMUL
    | FDIV
    | FMAC
    | MOV
    | MUL_CONST_ADD
    ;