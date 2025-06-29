package model

import (
	"ZeonicaParser/parser"
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

var _ parser.ZeonicaParserVisitor = (*MyVisitor)(nil)

type MyVisitor struct {
	*parser.BaseZeonicaParserVisitor
}

func NewMyVisitor() *MyVisitor {
	return &MyVisitor{&parser.BaseZeonicaParserVisitor{}}
}

func (v *MyVisitor) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil {
		return nil
	}
	return tree.Accept(v)
}

func (v *MyVisitor) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	blocks := []PEBlock{}
	for _, block := range ctx.AllPeBlock() {
		if peBlock, ok := block.(*parser.PeBlockContext); ok {
			if result := v.VisitPeBlock(peBlock); result != nil {
				blocks = append(blocks, result.(PEBlock))
			}
		}
	}
	return Program{Blocks: blocks}
}

func (v *MyVisitor) VisitPeBlock(ctx *parser.PeBlockContext) interface{} {
	x, _ := strconv.Atoi(ctx.DECIMAL_LITERAL(0).GetText())
	y := 0
	if ctx.DECIMAL_LITERAL(1) != nil {
		y, _ = strconv.Atoi(ctx.DECIMAL_LITERAL(1).GetText())
	}

	loopType := "Loop"
	if ctx.LoopType() != nil {
		lt := ctx.LoopType().(*parser.LoopTypeContext)
		if lt.ONCE() != nil {
			loopType = "Once"
		}
	}

	var body interface{}
	peBody := ctx.PeBody()
	if peBody == nil {
		panic("PeBody is nil")
	}
	switch b := peBody.(type) {
	case *parser.FlatBodyContext:
		flatBody := FlatBody{}
		fs := b.FlatStyle()
		if fs == nil {
			body = flatBody
			break
		}
		for _, group := range fs.AllLabeledGroup() {
			flatBody.Groups = append(flatBody.Groups, v.Visit(group).(LabeledGroup))
		}
		body = flatBody
	case *parser.EntryBlockBodyContext:
		entryBody := EntryBlockBody{}
		for _, block := range b.AllEntryBlock() {
			entryBody.Blocks = append(entryBody.Blocks, v.Visit(block).(EntryBlock))
		}
		body = entryBody
	default:
		panic(fmt.Sprintf("Unknown PeBody type: %T", peBody))
	}

	return PEBlock{
		X: x, Y: y, LoopType: loopType, Body: body,
	}
}

var validOpCodes = map[string]bool{
	"ADD":           true,
	"ADDI":          true,
	"SUB":           true,
	"SUBI":          true,
	"MUL":           true,
	"DIV":           true,
	"MAC":           true,
	"INC":           true,
	"LLS":           true,
	"LRS":           true,
	"OR":            true,
	"AND":           true,
	"XOR":           true,
	"NOT":           true,
	"FADD":          true,
	"FSUB":          true,
	"FMUL":          true,
	"FDIV":          true,
	"FMAC":          true,
	"MOV":           true,
	"MUL_CONST_ADD": true,
	"NOP":           true,
	"CBR":           true,
	"CMOV":          true,
	"SGE":           true,
}

func (v *MyVisitor) VisitNormalInst(ctx *parser.NormalInstContext) interface{} {
	inst := Instruction{}

	if ctx.OpCode() == nil {
		panic(fmt.Sprintf("Error: Invalid instruction format: %s", ctx.GetText()))
	}

	opCode := ctx.OpCode().GetText()
	if !validOpCodes[opCode] {
		panic(fmt.Sprintf("Error: Invalid opcode '%s'", opCode))
	}
	inst.OpCode = opCode

	if ctx.OperandList(0) != nil {
		srcs := v.Visit(ctx.OperandList(0)).([]OperandGroup)
		inst.SrcOperands = srcs
		if len(ctx.AllOperandList()) == 2 {
			inst.DstOperands = v.Visit(ctx.OperandList(1)).([]OperandGroup)
		}
	} else if len(ctx.AllOperand()) == 2 {
		inst.SrcOperands = []OperandGroup{v.Visit(ctx.Operand(0)).(OperandGroup)}
		inst.DstOperands = []OperandGroup{v.Visit(ctx.Operand(1)).(OperandGroup)}
	}

	return inst
}

func (v *MyVisitor) VisitOperandList(ctx *parser.OperandListContext) interface{} {
	operands := []OperandGroup{}
	for _, o := range ctx.AllOperand() {
		operands = append(operands, v.Visit(o).(OperandGroup))
	}
	return operands
}

func (v *MyVisitor) VisitOperand(ctx *parser.OperandContext) interface{} {
	opr := Operand{}

	// 处理谓词标签
	if ctx.PredTag() != nil {
		opr.PredTag = ctx.PredTag().GetText()
	}

	if ctx.IMM() != nil {
		opr.OperandType = IMM
		switch {
		case ctx.DECIMAL_LITERAL() != nil:
			val, _ := strconv.Atoi(ctx.DECIMAL_LITERAL().GetText())
			opr.Value = val
		case ctx.OCT_LITERAL() != nil:
			val, _ := strconv.ParseInt(ctx.OCT_LITERAL().GetText(), 8, 64)
			opr.Value = val
		case ctx.HEX_LITERAL() != nil:
			val, _ := strconv.ParseInt(ctx.HEX_LITERAL().GetText()[2:], 16, 64)
			opr.Value = val
		case ctx.BINARY_LITERAL() != nil:
			val, _ := strconv.ParseInt(ctx.BINARY_LITERAL().GetText()[2:], 2, 64)
			opr.Value = val
		case ctx.FLOAT_LITERAL() != nil:
			fval, _ := strconv.ParseFloat(ctx.FLOAT_LITERAL().GetText(), 64)
			opr.Value = fval
		}
	} else if ctx.MEM() != nil {
		opr.OperandType = MEM
		if ctx.IdList() != nil {
			opr.IDs = v.VisitIdList(ctx.IdList().(*parser.IdListContext)).([]string)
		} else if ctx.HEX_LITERAL() != nil {
			val, _ := strconv.ParseInt(ctx.HEX_LITERAL().GetText()[2:], 16, 64)
			opr.Value = val
		}
	} else if ctx.LabelID() != nil {
		opr.OperandType = LABEL
		opr.Value = ctx.LabelID().GetText()
	} else if ctx.IdList() != nil {
		opr.OperandType = REG
		ids := v.VisitIdList(ctx.IdList().(*parser.IdListContext)).([]string)
		for _, id := range ids {
			if _, err := strconv.Atoi(id); err == nil {
				id = "$" + id
			}
			opr.IDs = append(opr.IDs, id)
		}
	}

	group := OperandGroup{
		PredTag:  opr.PredTag,
		Operands: []Operand{opr},
	}
	return group
}

func (v *MyVisitor) VisitIdList(ctx *parser.IdListContext) interface{} {
	ids := []string{}
	for _, id := range ctx.AllIDENTIFIER() {
		ids = append(ids, id.GetText())
	}
	return ids
}

func (v *MyVisitor) VisitLabeledGroup(ctx *parser.LabeledGroupContext) interface{} {
	group := LabeledGroup{}
	if ctx.Label() != nil {
		group.Label = ctx.Label().LabelID().GetText()
	}

	for _, inst := range ctx.AllNormalInst() {
		group.Instructions = append(group.Instructions, v.Visit(inst).(Instruction))
	}

	return group
}

func (v *MyVisitor) VisitEntryBlock(ctx *parser.EntryBlockContext) interface{} {
	block := EntryBlock{}
	if ctx.LoopType().ONCE() != nil {
		block.LoopType = "Once"
	} else {
		block.LoopType = "Loop"
	}

	for _, group := range ctx.InstGroupList().AllInstGroup() {
		block.Groups = append(block.Groups, v.Visit(group).(InstructionGroup))
	}

	return block
}

func (v *MyVisitor) VisitInstGroup(ctx *parser.InstGroupContext) interface{} {
	group := InstructionGroup{}
	for _, inst := range ctx.AllNormalInst() {
		group.Instructions = append(group.Instructions, v.Visit(inst).(Instruction))
	}
	return group
}
