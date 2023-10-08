package listener

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/badochov/MRJP/instant/src/jvm/listener/code_template"
	parser2 "github.com/badochov/MRJP/instant/src/parser"
	"log"
	"math"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

type Listener struct {
	parser2.BaseInstantListener
	code         strings.Builder
	stackSize    int
	maxStackSize int
	locals       map[string]int
}

func (l *Listener) incStackSize() {
	l.stackSize++
	if l.stackSize > l.maxStackSize {
		l.maxStackSize = l.stackSize
	}
}

func (l *Listener) decStackSize() {
	l.stackSize--
}

// ExitEAdd is called when production add is exited.
func (l *Listener) ExitEAdd(ctx *parser2.EAddContext) {
	l.code.WriteString("iadd\n")
	l.decStackSize()
}

// ExitESub is called when production sub is exited.
func (l *Listener) ExitESub(ctx *parser2.ESubContext) {
	l.code.WriteString("isub\n")
	l.decStackSize()
}

// ExitEMul is called when production mul is exited.
func (l *Listener) ExitEMul(ctx *parser2.EMulContext) {
	op := ctx.MulOp().GetText()
	switch op {
	case "*":
		l.code.WriteString("imul\n")
		l.decStackSize()
	case "/":
		l.code.WriteString("idiv\n")
		l.decStackSize()
	default:
		log.Fatal("Unexpected mulOp:", op)
	}
}

// ExitEId is called when production id is exited.
func (l *Listener) ExitEId(ctx *parser2.EIdContext) {
	name := ctx.GetText()
	idx, ok := l.locals[name]
	if !ok {
		log.Fatalf("unknown variable '%s'", name)
	}

	var err error
	if idx <= 3 {
		_, err = fmt.Fprintf(&l.code, "iload_%d\n", idx)
	} else {
		_, err = fmt.Fprintf(&l.code, "iload %d\n", idx)
	}

	if err != nil {
		log.Fatal(err)
	}

	l.incStackSize()
}

// ExitEInt is called when production int is exited.
func (l *Listener) ExitEInt(ctx *parser2.EIntContext) {
	v, err := strconv.ParseInt(ctx.GetText(), 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	if v <= 5 {
		_, err = fmt.Fprintf(&l.code, "iconst_%d\n", v)
	} else if v < math.MaxInt8 {
		_, err = fmt.Fprintf(&l.code, "bipush %d\n", v)
	} else if v < math.MaxInt16 {
		_, err = fmt.Fprintf(&l.code, "sipush %d\n", v)
	} else {
		_, err = fmt.Fprintf(&l.code, "ldc %d\n", v)
	}

	if err != nil {
		log.Fatal(err)
	}

	l.incStackSize()
}

// ExitSAssign is called when production SAssign is exited.
func (l *Listener) ExitSAssign(ctx *parser2.SAssignContext) {
	ident := ctx.IDENT()
	name := ident.GetText()

	idx, ok := l.locals[name]
	if !ok {
		idx = len(l.locals)
		l.locals[name] = idx
	}

	var err error
	if idx <= 3 {
		_, err = fmt.Fprintf(&l.code, "istore_%d\n", idx)
	} else {
		_, err = fmt.Fprintf(&l.code, "istore %d\n", idx)
	}

	if err != nil {
		log.Fatal(err)
	}

	l.decStackSize()
}

// EnterSExpr is called when production SExpr is entered.
func (l *Listener) EnterSExpr(ctx *parser2.SExprContext) {
	l.incStackSize()
	l.code.WriteString("getstatic java/lang/System/out Ljava/io/PrintStream;\n")
}

// ExitSExpr is called when production SExpr is exited.
func (l *Listener) ExitSExpr(ctx *parser2.SExprContext) {
	l.code.WriteString("invokevirtual java/io/PrintStream/println(I)V\n")
	l.decStackSize()
	l.decStackSize()
}

// VisitErrorNode is called when an error node is visited.
func (l *Listener) VisitErrorNode(node antlr.ErrorNode) {
	log.Fatal(node.GetText())
}

func (l *Listener) getCode() string {
	return l.code.String()
}

func (l *Listener) Save(originalName string) error {
	savePath := getSavePath(originalName)
	f, err := os.Create(savePath)
	if err != nil {
		return err
	}
	name := getClassName(originalName)

	data := code_template.Data{
		ClassName:   name,
		StackLimit:  l.maxStackSize,
		Code:        l.getCode(),
		LocalsLimit: len(l.locals),
	}
	if err = code_template.Save(f, data); err != nil {
		return err
	}

	return l.compile(savePath)
}

const jasminPath = "./lib/jasmin.jar"

func (l *Listener) compile(path string) error {
	outPath := getCompilePath(path)
	cmd := exec.Command("java", "-jar", jasminPath, "-d", outPath, path)
	return cmd.Run()
}

func getSavePath(name string) string {
	withoutExtension := strings.TrimSuffix(name, ".ins")
	return withoutExtension + ".j"
}

func getCompilePath(name string) string {
	return path.Dir(name)
}

func getClassName(name string) string {
	withoutExtension := strings.TrimSuffix(name, ".ins")
	return path.Base(withoutExtension)
}

func New() *Listener {
	return &Listener{
		locals: make(map[string]int),
	}
}
