package model

import (
	"fmt"
	"strings"
)

type Program struct {
	Blocks []PEBlock
}

type PEBlock struct {
	X, Y     int
	LoopType string
	Body     interface{} // 可以是 FlatBody 或 EntryBlockBody
}

type FlatBody struct {
	Groups []LabeledGroup
}

type EntryBlockBody struct {
	Blocks []EntryBlock
}

type LabeledGroup struct {
	Label        string
	Instructions []Instruction
}

type EntryBlock struct {
	LoopType string
	Groups   []InstructionGroup
}

type InstructionGroup struct {
	Instructions []Instruction
}

type Instruction struct {
	OpCode      string
	SrcOperands []OperandGroup
	DstOperands []OperandGroup
}

type OperandGroup struct {
	PredTag  string
	Operands []Operand
}

type OperandType int

const (
	REG OperandType = iota
	IMM
	MEM
	LABEL
)

type Operand struct {
	OperandType OperandType
	PredTag     string
	IDs         []string
	Value       interface{}
}

func (o Operand) String() string {
	switch o.OperandType {
	case REG:
		if len(o.IDs) > 0 {
			return strings.Join(o.IDs, ", ")
		}
	case IMM:
		return fmt.Sprintf("%v", o.Value)
	case MEM:
		if len(o.IDs) > 0 {
			return fmt.Sprintf("MEM[%s]", strings.Join(o.IDs, ", "))
		}
		return fmt.Sprintf("MEM[0x%x]", o.Value)
	case LABEL:
		return fmt.Sprintf("%v", o.Value)
	}
	return ""
}

func (og OperandGroup) String() string {
	var result string
	if og.PredTag != "" {
		result += og.PredTag
	}
	if len(og.Operands) > 0 {
		var strs []string
		for _, o := range og.Operands {
			strs = append(strs, o.String())
		}
		result += strings.Join(strs, ", ")
	}
	return result
}
