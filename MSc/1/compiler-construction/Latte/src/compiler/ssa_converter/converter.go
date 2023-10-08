package ssa_converter

import (
	"fmt"
	"github.com/badochov/MRJP/Latte/src/IR/quadruple"
	"golang.org/x/exp/maps"
)

func ToSSA(program quadruple.Program) quadruple.Program {
	c := newConverter()

	for _, fn := range program.Fns {
		c.convertFn(fn)
	}
	for _, obj := range program.Objects {
		c.convertObj(obj)
	}
	return program
}

func newConverter() *converter {
	return &converter{
		replacementCounter: make(map[quadruple.Val]int),
	}
}

type phiData struct {
	var_ quadruple.Var
	phi  *quadruple.EPhi
}

type converter struct {
	currBlockData      blockData
	replacementCounter map[quadruple.Val]int

	cfg       *cfg
	blockData map[*quadruple.Block]blockData
	phis      map[*quadruple.Block][]phiData
	allVars   []quadruple.Var
	args      map[quadruple.Var]quadruple.Var
}

func (c *converter) VisitFn(fn *quadruple.Fn) any {
	return fn
}

func (c *converter) VisitMethod(method *quadruple.Method) any {
	return &quadruple.Method{
		Object: c.currBlockData.getReplacement(method.Object),
		Idx:    method.Idx,
	}
}

func (c *converter) convertFn(fn *quadruple.Fn) {
	c.blockData = make(map[*quadruple.Block]blockData)
	c.phis = make(map[*quadruple.Block][]phiData)
	c.cfg = buildCfg(fn)
	if c.cfg == nil {
		return
	}

	vars := make(map[quadruple.Var]struct{})
	for _, block := range fn.Blocks {
		varsUsed := gatherVarsUsed(block)
		for _, v := range varsUsed {
			vars[v] = struct{}{}
		}
	}
	c.allVars = maps.Keys(vars)

	c.args = make(map[quadruple.Var]quadruple.Var)
	for _, arg := range fn.Args {
		argVar := quadruple.Var{
			Type_: arg.Type,
			Id:    arg.Name,
		}
		c.args[argVar] = argVar
	}

	idx := 0
	for _, block := range fn.Blocks {
		if b := c.VisitBlock(block); b != nil {
			fn.Blocks[idx] = b.(*quadruple.Block)
			idx++
		}
	}
	fn.Blocks = fn.Blocks[:idx]

	fillPhis(c.cfg, c.blockData, c.phis)
}

func (c *converter) convertObj(obj *quadruple.Object) {
	for _, method := range obj.Methods {
		c.convertFn(method)
	}
}

func (c *converter) VisitEUnOp(e *quadruple.EUnOp) any {
	return &quadruple.EUnOp{
		Val: c.currBlockData.getReplacement(e.Val),
		Op:  e.Op,
	}
}

func (c *converter) VisitEVal(e *quadruple.EVal) any {
	return &quadruple.EVal{Val: c.currBlockData.getReplacement(e.Val)}
}

func (c *converter) VisitEBinOp(e *quadruple.EBinOp) any {
	return &quadruple.EBinOp{
		LVal: c.currBlockData.getReplacement(e.LVal),
		RVal: c.currBlockData.getReplacement(e.RVal),
		Op:   e.Op,
	}
}

func (c *converter) VisitECast(e *quadruple.ECast) any {
	return &quadruple.ECast{
		Val: c.currBlockData.getReplacement(e.Val),
		To:  e.To,
	}
}

func (c *converter) VisitECall(e *quadruple.ECall) any {
	args := make([]quadruple.Val, 0, len(e.Params))
	for _, arg := range e.Params {
		args = append(args, c.currBlockData.getReplacement(arg))
	}

	return &quadruple.ECall{
		Callable: e.Callable.Accept(c).(quadruple.Callable),
		Params:   args,
	}
}

func (c *converter) VisitEGetAttrPtr(e *quadruple.EGetAttrPtr) any {
	return &quadruple.EGetAttrPtr{
		Obj:     c.currBlockData.getReplacement(e.Obj),
		AttrIdx: e.AttrIdx,
	}
}

func (c *converter) VisitEGetArrElPtr(e *quadruple.EGetArrElPtr) any {
	return &quadruple.EGetArrElPtr{
		Arr: c.currBlockData.getReplacement(e.Arr),
		Idx: c.currBlockData.getReplacement(e.Idx),
	}
}

func (c *converter) VisitEGetArrLen(e *quadruple.EGetArrLen) any {
	return &quadruple.EGetArrLen{
		Arr: c.currBlockData.getReplacement(e.Arr),
	}
}

func (c *converter) VisitELoadPtr(e *quadruple.ELoadPtr) any {
	return &quadruple.ELoadPtr{
		Ptr: c.currBlockData.getReplacement(e.Ptr),
	}
}

func (c *converter) VisitENewArray(e *quadruple.ENewArray) any {
	return &quadruple.ENewArray{
		Size: c.currBlockData.getReplacement(e.Size),
		Type: e.Type,
	}
}

func (c *converter) VisitENewObj(e *quadruple.ENewObj) any {
	return &quadruple.ENewObj{
		Obj: e.Obj,
	}
}

func (c *converter) VisitEPhi(e *quadruple.EPhi) any {
	// PHIs are just dummies at this stage
	return e
}

type blockData struct {
	block           *quadruple.Block
	varReplacements map[quadruple.Var]quadruple.Var
}

func (c *converter) VisitBlock(b *quadruple.Block) any {
	c.currBlockData.block = b
	if c.cfg.root == b {
		c.currBlockData.varReplacements = c.args
	} else {
		if len(c.cfg.edgesIn[b]) == 0 {
			return nil
		}
		c.currBlockData.varReplacements = make(map[quadruple.Var]quadruple.Var)
	}

	c.blockData[b] = c.currBlockData

	if len(c.cfg.edgesIn[b]) > 0 {
		newInsts := make([]quadruple.Instruction, 0, len(c.allVars)+len(b.Instructions))
		phis := make([]phiData, 0, len(c.allVars))
		for _, v := range c.allVars {
			phi := &quadruple.EPhi{}
			newInsts = append(newInsts, &quadruple.Assignment{Var: v, Expr: phi})
			phis = append(phis, phiData{var_: v, phi: phi})
		}
		c.phis[b] = phis
		b.Instructions = append(newInsts, b.Instructions...)
	}

	for i, ex := range b.Instructions {
		b.Instructions[i] = ex.Accept(c).(quadruple.Instruction)
	}

	return b
}

func (c *converter) VisitUnreachable(u *quadruple.Unreachable) any {
	return &quadruple.Unreachable{}
}

func (c *converter) VisitGoto(g *quadruple.Goto) any {
	return &quadruple.Goto{
		Block: g.Block,
	}
}

func (c *converter) VisitPointerAssignment(pa *quadruple.PointerAssignment) any {
	return &quadruple.PointerAssignment{
		Pointer: c.currBlockData.getReplacement(pa.Pointer),
		Val:     c.currBlockData.getReplacement(pa.Val),
	}
}

func (c *converter) VisitAssignment(a *quadruple.Assignment) any {
	expr := a.Expr.Accept(c).(quadruple.Expr)
	var_ := c.replaceVarAssignment(a.Var)

	return &quadruple.Assignment{
		Var:  var_,
		Expr: expr,
	}
}

func (c *converter) VisitReturn(r *quadruple.Return) any {
	val := r.Val
	if val != nil {
		updated := c.currBlockData.getReplacement(*val)
		val = &updated
	}

	return &quadruple.Return{
		Val: val,
	}
}

func (c *converter) VisitVoidCall(fc *quadruple.VoidCall) any {
	args := make([]quadruple.Val, 0, len(fc.Params))
	for _, arg := range fc.Params {
		args = append(args, c.currBlockData.getReplacement(arg))
	}

	return &quadruple.VoidCall{
		Callable: fc.Callable.Accept(c).(quadruple.Callable),
		Params:   args,
	}
}

func (c *converter) VisitCondJump(cj *quadruple.CondJump) any {
	return &quadruple.CondJump{
		Cond:       c.currBlockData.getReplacement(cj.Cond),
		TrueBlock:  cj.TrueBlock,
		FalseBlock: cj.FalseBlock,
	}
}

func (c *converter) replaceVarAssignment(var_ quadruple.Var) quadruple.Var {
	counter := c.replacementCounter[var_]
	c.replacementCounter[var_] = counter + 1

	newVar := quadruple.Var{
		Type_: var_.Type_,
		Id:    fmt.Sprintf("%s_%d_%d", var_.Id, c.currBlockData.block.Id, counter),
	}
	c.currBlockData.varReplacements[var_] = newVar

	return newVar
}

func (b *blockData) getReplacement(val quadruple.Val) quadruple.Val {
	if var_, isVar := val.(quadruple.Var); isVar {
		return b.getVarReplacement(var_)
	}
	return val
}

func (b *blockData) getVarReplacement(var_ quadruple.Var) quadruple.Var {
	if replacement, ok := b.varReplacements[var_]; ok {
		return replacement
	}
	return var_
}

var _ quadruple.InstructionVisitor = &converter{}
var _ quadruple.ExprVisitor = &converter{}
var _ quadruple.BlockVisitor = &converter{}
var _ quadruple.CallableVisitor = &converter{}
