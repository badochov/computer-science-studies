package listener

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	parser2 "github.com/badochov/MRJP/instant/src/parser"
	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"

	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var intType = types.I32

type Listener struct {
	parser2.BaseInstantListener
	module      *ir.Module
	block       *ir.Block
	allocations map[string]*ir.InstAlloca
	stack       []value.Value
	print       *ir.Func
}

func (l *Listener) pop() value.Value {
	el := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]
	return el
}

func (l *Listener) push(el value.Value) {
	l.stack = append(l.stack, el)
}

// ExitEAdd is called when production add is exited.
func (l *Listener) ExitEAdd(ctx *parser2.EAddContext) {
	b := l.pop()
	a := l.pop()

	add := l.block.NewAdd(a, b)
	l.push(add)
}

// ExitESub is called when production sub is exited.
func (l *Listener) ExitESub(ctx *parser2.ESubContext) {
	b := l.pop()
	a := l.pop()

	sub := l.block.NewSub(a, b)
	l.push(sub)
}

// ExitEMul is called when production mul is exited.
func (l *Listener) ExitEMul(ctx *parser2.EMulContext) {
	b := l.pop()
	a := l.pop()

	op := ctx.MulOp().GetText()
	switch op {
	case "*":
		mul := l.block.NewMul(a, b)
		l.push(mul)
	case "/":
		div := l.block.NewSDiv(a, b)
		l.push(div)
	default:
		log.Fatal("Unexpected mulOp:", op)
	}

}

// ExitEId is called when production id is exited.
func (l *Listener) ExitEId(ctx *parser2.EIdContext) {
	name := ctx.GetText()
	v, ok := l.allocations[name]
	if !ok {
		log.Fatalf("unknown variable `%s`", name)
	}

	s := l.block.NewLoad(intType, v)
	l.push(s)
}

// ExitEInt is called when production int is exited.
func (l *Listener) ExitEInt(ctx *parser2.EIntContext) {
	v, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		panic(err)
	}

	l.push(constant.NewInt(intType, v))
}

// ExitSAssign is called when production SAssign is exited.
func (l *Listener) ExitSAssign(ctx *parser2.SAssignContext) {
	val := l.pop()
	identNode := ctx.IDENT()
	ident := identNode.GetText()

	al, ok := l.allocations[ident]
	if !ok {
		al = l.block.NewAlloca(intType)
		l.allocations[ident] = al
	}

	l.block.NewStore(val, al)
}

// ExitSExpr is called when production SExpr is exited.
func (l *Listener) ExitSExpr(ctx *parser2.SExprContext) {
	val := l.pop()
	l.block.NewCall(l.print, val)
}

// ExitProgram is called when production program is exited.
func (l *Listener) ExitProgram(ctx *parser2.ProgramContext) {
	l.block.NewRet(constant.NewInt(types.I32, 0))
}

// VisitErrorNode is called when an error node is visited.
func (l *Listener) VisitErrorNode(node antlr.ErrorNode) {
	log.Fatal(node.GetText())
}

func (l *Listener) getCode() string {
	return l.module.String()
}

func (l *Listener) Save(originalName string) error {
	path := getSavePath(originalName)
	if err := os.WriteFile(path, []byte(l.getCode()), 0644); err != nil {
		return err
	}

	return l.compile(path)
}

func (l *Listener) compile(path string) error {
	cmd := exec.Command("llvm-as", path)
	return cmd.Run()
}

func getSavePath(name string) string {
	withoutExtension := strings.TrimSuffix(name, ".ins")
	return withoutExtension + ".ll"
}

func findPrintInt(fns []*ir.Func) *ir.Func {
	const name = "printInt"

	for _, fn := range fns {
		if fn.Name() == name {
			return fn
		}
	}
	log.Fatalf("cannot find function `%s`", name)
	return nil
}

func New() *Listener {
	mod, err := asm.ParseString("runtime.ll", runtime)
	if err != nil {
		log.Fatal(err)
	}

	main := mod.NewFunc("main", types.I32)
	entry := main.NewBlock("")

	return &Listener{
		BaseInstantListener: parser2.BaseInstantListener{},
		module:              mod,
		block:               entry,
		print:               findPrintInt(mod.Funcs),
		allocations:         make(map[string]*ir.InstAlloca),
		stack:               nil,
	}
}
