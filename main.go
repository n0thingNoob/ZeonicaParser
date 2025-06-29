package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"ZeonicaParser/model"
	"ZeonicaParser/parser"
)

func main() {
	source, err := os.ReadFile("check.asm")
	if err != nil {
		log.Fatalf("cannot read check.asm: %v", err)
	}

	input := antlr.NewInputStream(string(source))
	lexer := parser.NewZeonicaLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewZeonicaParser(tokens)

	lexer.RemoveErrorListeners()
	p.RemoveErrorListeners()
	diag := antlr.NewDiagnosticErrorListener(true)
	lexer.AddErrorListener(diag)
	p.AddErrorListener(diag)

	fmt.Println("Starting parsing...")
	tree := p.CompilationUnit()
	fmt.Println("Parsing completed")

	fmt.Println("Starting visitor...")
	visitor := model.NewMyVisitor()
	res := visitor.Visit(tree)
	fmt.Println("Visitor completed")

	prog, ok := res.(model.Program)
	if !ok {
		fmt.Fprintln(os.Stderr, "parse failed or visitor returned unexpected type")
		os.Exit(1)
	}

	fmt.Printf("Parsed %d PE block(s)\n", len(prog.Blocks))
	for _, pe := range prog.Blocks {
		switch body := pe.Body.(type) {
		case model.FlatBody:
			fmt.Printf("PE(%d,%d) Mode: %s, %d group(s)\n", pe.X, pe.Y, pe.LoopType, len(body.Groups))
			for i, g := range body.Groups {
				if g.Label != "" {
					fmt.Printf("  Group %d (%s):\n", i, g.Label)
				} else {
					fmt.Printf("  Group %d:\n", i)
				}
				for _, inst := range g.Instructions {
					fmt.Printf("    Op: %-8s Src: %s  Dst: %s\n",
						inst.OpCode, formatOperands(inst.SrcOperands), formatOperands(inst.DstOperands))
				}
			}
		case model.EntryBlockBody:
			fmt.Printf("PE(%d,%d) Mode: %s, %d entry block(s)\n", pe.X, pe.Y, pe.LoopType, len(body.Blocks))
			for i, b := range body.Blocks {
				fmt.Printf("    EntryBlock %d (Mode: %s):\n", i, b.LoopType)
				for j, g := range b.Groups {
					fmt.Printf("        Group %d:\n", j)
					for _, inst := range g.Instructions {
						fmt.Printf("            Op: %-8s Src: %s  Dst: %s\n",
							inst.OpCode, formatOperands(inst.SrcOperands), formatOperands(inst.DstOperands))
					}
				}
			}
		}
	}
}

func formatOperands(operands []model.OperandGroup) string {
	var result []string
	for _, og := range operands {
		for _, o := range og.Operands {
			opStr := o.String()
			if og.PredTag != "" {
				result = append(result, og.PredTag+"["+opStr+"]")
			} else {
				result = append(result, "["+opStr+"]")
			}
		}
	}
	return strings.Join(result, " ")
}
