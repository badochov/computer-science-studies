package converter

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/badochov/MRJP/Latte/src/IR/ast"
	"github.com/badochov/MRJP/Latte/src/IR/quadruple"
	"github.com/badochov/MRJP/Latte/src/frontend/types"
	"github.com/badochov/MRJP/Latte/src/topdefs/signatures"
)

func ConvertAst[Data antlr.ParserRuleContext](program ast.IProgram[Data], topDefs signatures.TopDefs) quadruple.Program {
	programStub := createProgramStub(topDefs)
	c := newConverter[Data](programStub)
	program.VisitProgram(c)
	return c.program
}

func createProgramStub(defs signatures.TopDefs) quadruple.Program {
	objs := convertObjStubs(defs.Classes)
	fns := convertFnStubs(defs.Fns, objs)
	return quadruple.Program{
		Objects: objs,
		Fns:     fns,
		Strings: make(map[string]struct{}),
	}
}

func convertObjStubs(classes map[string]signatures.Class) map[string]*quadruple.Object {
	allObjs := make(map[string]*quadruple.Object, len(classes))
	for name := range classes {
		allObjs[name] = &quadruple.Object{}
	}

	ret := make(map[string]*quadruple.Object, len(classes))
	for _, cl := range classes {
		convertClass(cl, allObjs, ret)
	}

	return ret
}

func convertClass(cl signatures.Class, allObjs map[string]*quadruple.Object, converted map[string]*quadruple.Object) *quadruple.Object {
	if obj, ok := converted[cl.Name]; ok {
		return obj
	}

	this := allObjs[cl.Name]
	if cl.Extends != nil {
		*this = convertClass(*cl.Extends, allObjs, converted).Copy()
		this.Extends = converted[cl.Extends.Name]
	} else {
		*this = quadruple.Object{
			Methods:        make([]*quadruple.Fn, 0, len(cl.Methods)),
			Attributes:     make([]quadruple.Type, 0, len(cl.Attributes)+1),
			AttributeToIdx: make(map[string]int, len(cl.Attributes)+1),
			MethodsMap:     make(map[string]int, len(cl.Methods)),
		}
		const vTablePlaceholderName = ".vtable"
		this.Attributes = append(this.Attributes, quadruple.Type{})
		this.AttributeToIdx[vTablePlaceholderName] = 0
	}
	this.Name = cl.Name

	for name, attrType := range cl.Attributes {
		this.AttributeToIdx[name] = len(this.Attributes)
		this.Attributes = append(this.Attributes, toQuadrupleType(attrType, allObjs))
	}

	for name, method := range cl.Methods {
		qM := &quadruple.Fn{
			Name: fmt.Sprintf(".%s__%s", cl.Name, name),
			Ret:  toQuadrupleType(method.RetType, allObjs),
			Args: make([]quadruple.Arg, 0, len(method.ArgTypes)+1),
		}
		qM.Args = append(qM.Args, quadruple.Arg{
			Name: "self",
			Type: quadruple.Type{
				Type:   quadruple.ObjectType,
				Object: this,
			},
			Ref: true,
		})
		for _, arg := range method.ArgTypes {
			qM.Args = append(qM.Args, quadruple.Arg{
				Name: arg.Name,
				Type: toQuadrupleType(arg.Type, allObjs),
				Ref:  arg.Ref,
			})
		}

		if prev, ok := this.MethodsMap[name]; ok {
			this.Methods[prev] = qM
		} else {
			this.MethodsMap[name] = len(this.Methods)
			this.Methods = append(this.Methods, qM)
		}
	}

	converted[cl.Name] = this

	return this
}

func convertFnStubs(fns map[string]signatures.Function, objs map[string]*quadruple.Object) map[string]*quadruple.Fn {
	ret := make(map[string]*quadruple.Fn, len(fns))
	for name, fn := range fns {
		qF := &quadruple.Fn{
			Name: name,
			Ret:  toQuadrupleType(fn.RetType, objs),
		}
		for _, arg := range fn.ArgTypes {
			qF.Args = append(qF.Args, quadruple.Arg{
				Name: arg.Name,
				Type: toQuadrupleType(arg.Type, objs),
				Ref:  arg.Ref,
			})
		}
		ret[name] = qF
	}

	return ret
}

type converter[Data antlr.ParserRuleContext] struct {
	blockCtx      *blockCtx
	currType      types.Type
	self          *quadruple.Var
	tempCounter   int
	blockCounter  int
	currFn        *quadruple.Fn
	currBlockInfo *currBlockInfo
	program       quadruple.Program
}

type currBlockInfo struct {
	block      *quadruple.Block
	terminates bool
	jumpTo     *quadruple.Block
}

func newConverter[Data antlr.ParserRuleContext](program quadruple.Program) *converter[Data] {
	return &converter[Data]{
		program: program,
		blockCtx: &blockCtx{
			prev: nil,
		},
	}
}

func (c *converter[Data]) newVar(t quadruple.Type) quadruple.Var {
	c.tempCounter++
	return quadruple.Var{
		Type_: t,
		Id:    fmt.Sprintf("%d", c.tempCounter),
	}
}

func (c *converter[Data]) resetLabels() {
	c.tempCounter = 0
	c.blockCounter = 0
}

func (c *converter[Data]) VisitClassAttribute(attribute *ast.Attribute[Data]) any {
	return nil
}

func (c *converter[Data]) VisitClassMethod(m *ast.Method[Data]) any {
	c.resetLabels()

	c.enterBlock()
	defer c.exitBlock()

	c.currFn = c.self.Type().Object.GetMethod(m.Id)

	for _, arg := range c.currFn.Arguments() {
		if arg.Name == "self" {
			c.blockCtx.DeclareVar(arg.Name, var_{val: *c.self})
		} else {
			c.blockCtx.DeclareVar(arg.Name, var_{
				val: quadruple.Var{
					Type_: arg.Type,
					Id:    fmt.Sprintf("%s", arg.Name),
				},
			})
		}
	}

	c.setBlock(c.newBlock(), nil)
	m.Block.VisitBlock(c)
	c.endBlock()

	return nil

}

func (c *converter[Data]) VisitBlock(b *ast.Block[Data]) any {
	c.enterBlock()
	defer c.exitBlock()

	for _, stmt := range b.Stmts {
		stmt.VisitStmt(c)
	}

	return nil
}

func (c *converter[Data]) VisitStmtEmpty(s *ast.StmtEmpty[Data]) any {
	return nil
}

func (c *converter[Data]) VisitStmtBlock(s *ast.StmtBlock[Data]) any {
	return s.Block.VisitBlock(c)
}

func (c *converter[Data]) VisitStmtDecl(s *ast.StmtDecl[Data]) any {
	c.currType = ast.GetType(s.Type)
	for _, i := range s.Items {
		i.VisitItem(c)
	}

	return nil
}

func (c *converter[Data]) VisitStmtAss(s *ast.StmtAss[Data]) any {
	v := c.blockCtx.GetVar(s.Id)

	ex := s.Expr.VisitExpr(c).(quadruple.Val)
	if v.self != nil {
		return c.attrAss(*v.self, s.Id, ex)
	}
	correctType := c.castIfNeeded(ex, v.val.Type())
	c.appendInstructions(&quadruple.Assignment{Var: v.val, Expr: &quadruple.EVal{Val: correctType}})

	return nil
}

func (c *converter[Data]) VisitStmtRet(s *ast.StmtRet[Data]) any {
	v := s.Expr.VisitExpr(c).(quadruple.Val)
	correctType := c.castIfNeeded(v, c.currFn.RetType())
	c.appendInstructions(&quadruple.Return{Val: &correctType})
	c.markBlockTerminates()

	return nil
}

func (c *converter[Data]) VisitStmtVRet(s *ast.StmtVRet[Data]) any {
	c.appendInstructions(&quadruple.Return{Val: nil})
	c.markBlockTerminates()

	return nil
}

func (c *converter[Data]) VisitStmtCond(s *ast.StmtCond[Data]) any {
	blockStmt := c.newBlock()
	blockIfEnd := c.newBlock()

	cond := s.Cond.VisitExpr(c).(quadruple.Val)
	c.appendInstructions(&quadruple.CondJump{
		Cond:       cond,
		TrueBlock:  blockStmt,
		FalseBlock: blockIfEnd,
	})
	jt := c.currJumpTo()
	c.markBlockTerminates()
	c.endBlock()

	c.setBlock(blockStmt, blockIfEnd)
	s.Stmt.VisitStmt(c)
	c.endBlock()

	c.setBlock(blockIfEnd, jt)

	return nil
}

func (c *converter[Data]) VisitStmtCondElse(s *ast.StmtCondElse[Data]) any {
	blockStmtTrue := c.newBlock()
	blockStmtFalse := c.newBlock()
	blockIfEnd := c.newBlock()

	cond := s.Cond.VisitExpr(c).(quadruple.Val)
	c.appendInstructions(&quadruple.CondJump{
		Cond:       cond,
		TrueBlock:  blockStmtTrue,
		FalseBlock: blockStmtFalse,
	})
	jt := c.currJumpTo()
	c.markBlockTerminates()
	c.endBlock()

	c.setBlock(blockStmtTrue, blockIfEnd)
	s.Stmt.VisitStmt(c)
	c.endBlock()

	c.setBlock(blockStmtFalse, blockIfEnd)
	s.ElseStmt.VisitStmt(c)
	c.endBlock()

	c.setBlock(blockIfEnd, jt)

	return nil
}

func (c *converter[Data]) VisitStmtWhile(s *ast.StmtWhile[Data]) any {
	blockCond := c.newBlock()
	blockStmt := c.newBlock()
	blockWhileEnd := c.newBlock()

	c.appendInstructions(&quadruple.Goto{
		Block: blockCond,
	})
	jt := c.currJumpTo()
	c.markBlockTerminates()
	c.endBlock()

	c.setBlock(blockCond, jt)
	cond := s.Cond.VisitExpr(c).(quadruple.Val)
	c.appendInstructions(&quadruple.CondJump{
		Cond:       cond,
		TrueBlock:  blockStmt,
		FalseBlock: blockWhileEnd,
	})
	c.markBlockTerminates()
	c.endBlock()

	c.setBlock(blockStmt, blockCond)
	s.Stmt.VisitStmt(c)
	c.endBlock()

	c.setBlock(blockWhileEnd, jt)

	return nil
}

func (c *converter[Data]) VisitStmtExp(s *ast.StmtExp[Data]) any {
	return s.Expr.VisitExpr(c)
}

func (c *converter[Data]) VisitStmtIndexAss(s *ast.StmtIndexAss[Data]) any {
	op := &quadruple.EGetArrElPtr{
		Arr: s.ArrExpr.VisitExpr(c).(quadruple.Val),
		Idx: s.IndexExpr.VisitExpr(c).(quadruple.Val),
	}
	t := op.Arr.Type()
	t.Array = false
	ptrT := t
	ptrT.Pointer = true
	ptrV := c.newVar(ptrT)
	ptr := &quadruple.Assignment{
		Var:  ptrV,
		Expr: op,
	}
	exprV := c.castIfNeeded(s.Expr.VisitExpr(c).(quadruple.Val), t)
	ptrAss := &quadruple.PointerAssignment{
		Pointer: ptrV,
		Val:     exprV,
	}
	c.appendInstructions(ptr, ptrAss)

	return nil
}

func (c *converter[Data]) VisitStmtAttrAss(s *ast.StmtAttrAss[Data]) any {
	obj := s.ObjExpr.VisitExpr(c).(quadruple.Val)
	ex := s.Expr.VisitExpr(c).(quadruple.Val)

	return c.attrAss(obj, s.Attr, ex)
}

func (c *converter[Data]) VisitExprParen(e *ast.ExprParen[Data]) any {
	return e.Expr.VisitExpr(c)
}

func (c *converter[Data]) VisitExprUnOp(e *ast.ExprUnOp[Data]) any {
	var op quadruple.Expr
	var typ quadruple.Type
	switch e.UnOp {
	case ast.LogicNot:
		unOp := &quadruple.EUnOp{
			Val: e.Expr.VisitExpr(c).(quadruple.Val),
			Op:  quadruple.LogicNot,
		}
		typ = unOp.Val.Type()
		op = unOp
	case ast.UnaryMinus:
		binOp := &quadruple.EBinOp{
			LVal: quadruple.IntVal{Int: 0},
			RVal: e.Expr.VisitExpr(c).(quadruple.Val),
			Op:   quadruple.Minus,
		}
		typ = binOp.LVal.Type()
		op = binOp
	}
	v := c.newVar(typ)
	c.appendInstructions(&quadruple.Assignment{
		Var:  v,
		Expr: op,
	})

	return v
}

func (c *converter[Data]) VisitExprMulOp(e *ast.ExprMulOp[Data]) any {
	var binOp quadruple.BinOp
	switch e.MulOp {
	case ast.Times:
		binOp = quadruple.Times
	case ast.Divide:
		binOp = quadruple.Divide
	case ast.Modulo:
		binOp = quadruple.Modulo
	}

	op := &quadruple.EBinOp{
		LVal: e.LExpr.VisitExpr(c).(quadruple.Val),
		RVal: e.RExpr.VisitExpr(c).(quadruple.Val),
		Op:   binOp,
	}
	v := c.newVar(op.LVal.Type())
	c.appendInstructions(&quadruple.Assignment{
		Var:  v,
		Expr: op,
	})

	return v
}

func (c *converter[Data]) VisitExprAddOp(e *ast.ExprAddOp[Data]) any {
	lVal := e.LExpr.VisitExpr(c).(quadruple.Val)
	rVal := e.RExpr.VisitExpr(c).(quadruple.Val)

	var binOp quadruple.BinOp
	switch e.AddOp {
	case ast.Plus:
		binOp = quadruple.Plus
		if lVal.Type().Type == quadruple.StringType {
			op := &quadruple.ECall{
				Callable: c.program.Fns[".str_add"],
				Params:   []quadruple.Val{lVal, rVal},
			}
			v := c.newVar(lVal.Type())
			c.appendInstructions(&quadruple.Assignment{
				Var:  v,
				Expr: op,
			})
			return v
		}
	case ast.Minus:
		binOp = quadruple.Minus
	}

	op := &quadruple.EBinOp{
		LVal: lVal,
		RVal: rVal,
		Op:   binOp,
	}
	v := c.newVar(op.LVal.Type())
	c.appendInstructions(&quadruple.Assignment{
		Var:  v,
		Expr: op,
	})

	return v
}

func (c *converter[Data]) VisitExprRelOp(e *ast.ExprRelOp[Data]) any {
	lVal := e.LExpr.VisitExpr(c).(quadruple.Val)
	rVal := e.RExpr.VisitExpr(c).(quadruple.Val)

	var binOp quadruple.BinOp
	switch e.RelOp {
	case ast.Ne:
		binOp = quadruple.Ne
		lVal, rVal = c.toSameType(lVal, rVal)
	case ast.Eq:
		binOp = quadruple.Eq
		lVal, rVal = c.toSameType(lVal, rVal)
	case ast.Lt:
		binOp = quadruple.Gt
		lVal, rVal = rVal, lVal
	case ast.Le:
		binOp = quadruple.Ge
		lVal, rVal = rVal, lVal
	case ast.Gt:
		binOp = quadruple.Gt
	case ast.Ge:
		binOp = quadruple.Ge
	}

	op := &quadruple.EBinOp{
		LVal: lVal,
		RVal: rVal,
		Op:   binOp,
	}
	v := c.newVar(op.LVal.Type())
	c.appendInstructions(&quadruple.Assignment{
		Var:  v,
		Expr: op,
	})

	return v
}

func (c *converter[Data]) toSameType(lval quadruple.Val, rval quadruple.Val) (quadruple.Val, quadruple.Val) {
	lType := lval.Type()
	lobj := lType.Object
	rType := rval.Type()
	robj := rType.Object

	if rType.Type == quadruple.UntypedNullType {
		return c.castIfNeeded(rval, lType), lval
	}
	if lType.Type == quadruple.UntypedNullType {
		return rval, c.castIfNeeded(lval, rType)
	}

	if lobj == robj {
		return lval, rval
	}

	for currObj := robj.Extends; currObj != nil; currObj = currObj.Extends {
		if currObj == lobj {
			return lval, c.castIfNeeded(rval, lType)
		}
	}

	return c.castIfNeeded(lval, rType), rval
}

func (c *converter[Data]) VisitExprAnd(e *ast.ExprAnd[Data]) any {
	trueBlock := c.newBlock()
	falseBlock := c.newBlock()
	mergeBlock := c.newBlock()

	condVal := e.LExpr.VisitExpr(c).(quadruple.Val)
	cond := c.toVar(condVal)
	c.appendInstructions(&quadruple.CondJump{
		Cond:       cond,
		TrueBlock:  trueBlock,
		FalseBlock: falseBlock,
	})
	jt := c.currJumpTo()
	c.markBlockTerminates()
	c.endBlock()

	c.setBlock(trueBlock, nil)
	ass := &quadruple.Assignment{
		Var:  cond,
		Expr: &quadruple.EVal{Val: e.RExpr.VisitExpr(c).(quadruple.Val)},
	}
	gt := &quadruple.Goto{Block: mergeBlock}
	c.appendInstructions(ass, gt)
	c.markBlockTerminates()
	c.endBlock()

	c.setBlock(falseBlock, nil)
	assF := &quadruple.Assignment{
		Var:  cond,
		Expr: &quadruple.EVal{Val: quadruple.BoolVal{Bool: false}},
	}
	gtF := &quadruple.Goto{Block: mergeBlock}
	c.appendInstructions(assF, gtF)
	c.markBlockTerminates()
	c.endBlock()

	c.setBlock(mergeBlock, jt)

	return cond
}

func (c *converter[Data]) VisitExprOr(e *ast.ExprOr[Data]) any {
	trueBlock := c.newBlock()
	falseBlock := c.newBlock()
	mergeBlock := c.newBlock()

	condVal := e.LExpr.VisitExpr(c).(quadruple.Val)
	cond := c.toVar(condVal)
	c.appendInstructions(&quadruple.CondJump{
		Cond:       cond,
		TrueBlock:  trueBlock,
		FalseBlock: falseBlock,
	})
	jt := c.currJumpTo()
	c.markBlockTerminates()
	c.endBlock()

	c.setBlock(trueBlock, nil)
	assT := &quadruple.Assignment{
		Var:  cond,
		Expr: &quadruple.EVal{Val: quadruple.BoolVal{Bool: true}},
	}
	gtT := &quadruple.Goto{Block: mergeBlock}
	c.appendInstructions(assT, gtT)
	c.markBlockTerminates()
	c.endBlock()

	c.setBlock(falseBlock, nil)
	assF := &quadruple.Assignment{
		Var:  cond,
		Expr: &quadruple.EVal{Val: e.RExpr.VisitExpr(c).(quadruple.Val)},
	}
	gtF := &quadruple.Goto{Block: mergeBlock}
	c.appendInstructions(assF, gtF)
	c.markBlockTerminates()
	c.endBlock()

	c.setBlock(mergeBlock, jt)

	return cond
}

func (c *converter[Data]) VisitExprId(e *ast.ExprId[Data]) any {
	v := c.blockCtx.GetVar(e.Id)
	if v.self != nil {
		return c.attr(*v.self, e.Id)
	}

	val := c.newVar(v.val.Type())
	c.appendInstructions(&quadruple.Assignment{
		Var: val,
		Expr: &quadruple.EVal{
			Val: v.val,
		},
	})
	return val
}

func (c *converter[Data]) VisitExprInt(e *ast.ExprInt[Data]) any {
	return quadruple.IntVal{Int: e.Int}
}

func (c *converter[Data]) VisitExprBool(e *ast.ExprBool[Data]) any {
	return quadruple.BoolVal{Bool: e.Bool}
}

func (c *converter[Data]) VisitExprString(e *ast.ExprString[Data]) any {
	if _, ok := c.program.Strings[e.String]; !ok {
		c.program.Strings[e.String] = struct{}{}
	}

	return quadruple.StrVal{String: e.String}
}

func (c *converter[Data]) VisitExprFunCall(e *ast.ExprFunCall[Data]) any {
	params := make([]quadruple.Val, 0, len(e.Params))
	for _, p := range e.Params {
		params = append(params, p.VisitExpr(c).(quadruple.Val))
	}

	if f, ok := c.program.Fns[e.Id]; ok {
		return c.fnCall(f, params)
	}

	return c.methodCall(c.self, e.Id, params)
}

func (c *converter[Data]) methodCall(self quadruple.Val, name string, params []quadruple.Val) any {
	p := make([]quadruple.Val, 0, len(params)+1)
	p = append(p, self)
	p = append(p, params...)

	method := &quadruple.Method{
		Object: self,
		Idx:    self.Type().Object.MethodsMap[name],
	}

	return c.fnCall(method, p)
}

func (c *converter[Data]) fnCall(fn quadruple.Callable, params []quadruple.Val) any {
	if fn.RetType().Type == quadruple.VoidType {
		c.appendInstructions(&quadruple.VoidCall{
			Callable: fn,
			Params:   params,
		})
		return nil
	}

	op := &quadruple.ECall{
		Callable: fn,
		Params:   params,
	}
	v := c.newVar(fn.RetType())
	c.appendInstructions(&quadruple.Assignment{
		Var:  v,
		Expr: op,
	})
	return v
}

func (c *converter[Data]) VisitExprIndex(e *ast.ExprIndex[Data]) any {
	op := &quadruple.EGetArrElPtr{
		Arr: e.ArrExpr.VisitExpr(c).(quadruple.Val),
		Idx: e.IndexExpr.VisitExpr(c).(quadruple.Val),
	}
	t := op.Arr.Type()
	t.Pointer = true
	t.Array = false
	ptrV := c.newVar(t)
	ptrAss := &quadruple.Assignment{
		Var:  ptrV,
		Expr: op,
	}
	elV := c.newVar(quadruple.Type{Type: ptrAss.Var.Type().Type})
	ptrDeref := &quadruple.Assignment{
		Var:  elV,
		Expr: &quadruple.ELoadPtr{Ptr: ptrV},
	}
	c.appendInstructions(ptrAss, ptrDeref)
	return elV
}

func (c *converter[Data]) VisitExprNull(e *ast.ExprNull[Data]) any {
	return c.getNull(ast.GetType(e.Type))
}

func (c *converter[Data]) VisitExprNewClass(e *ast.ExprNewClass[Data]) any {
	obj := c.program.Objects[e.Type.Id]
	op := &quadruple.ENewObj{
		Obj: obj,
	}
	v := c.newVar(quadruple.Type{
		Type:   quadruple.ObjectType,
		Object: obj,
	})
	c.appendInstructions(&quadruple.Assignment{
		Var:  v,
		Expr: op,
	})

	return v
}

func (c *converter[Data]) VisitExprNewArray(e *ast.ExprNewArray[Data]) any {
	op := &quadruple.ENewArray{
		Type: c.toQuadrupleType(ast.GetTypeSingle(e.Type)),
		Size: e.SizeExpr.VisitExpr(c).(quadruple.Val),
	}
	arrT := op.Type
	arrT.Array = true
	v := c.newVar(arrT)
	c.appendInstructions(&quadruple.Assignment{
		Var:  v,
		Expr: op,
	})
	return v
}

func (c *converter[Data]) VisitExprMethodCall(e *ast.ExprMethodCall[Data]) any {
	obj := e.ObjExpr.VisitExpr(c).(quadruple.Val)

	params := make([]quadruple.Val, 0, len(e.Params))
	for _, p := range e.Params {
		params = append(params, p.VisitExpr(c).(quadruple.Val))
	}

	return c.methodCall(obj, e.Id, params)
}

func (c *converter[Data]) VisitExprAttr(e *ast.ExprAttr[Data]) any {
	return c.attr(e.ObjExpr.VisitExpr(c).(quadruple.Val), e.Attr)
}

func (c *converter[Data]) attr(obj quadruple.Val, attr string) quadruple.Val {
	objType := obj.Type()

	if objType.Array {
		op := &quadruple.EGetArrLen{
			Arr: obj,
		}

		vLen := c.newVar(quadruple.Type{Type: quadruple.IntType})
		c.appendInstructions(&quadruple.Assignment{
			Var:  vLen,
			Expr: op,
		})

		return vLen
	}

	attrIdx := objType.Object.AttributeToIdx[attr]
	op := &quadruple.EGetAttrPtr{
		Obj:     obj,
		AttrIdx: attrIdx,
	}
	elT := objType.Object.Attributes[attrIdx]
	ptrT := elT
	ptrT.Pointer = true
	ptrV := c.newVar(ptrT)
	ptrAss := &quadruple.Assignment{
		Var:  ptrV,
		Expr: op,
	}
	elV := c.newVar(elT)
	ptrDeref := &quadruple.Assignment{
		Var:  elV,
		Expr: &quadruple.ELoadPtr{Ptr: ptrV},
	}
	c.appendInstructions(ptrAss, ptrDeref)
	return elV
}

func (c *converter[Data]) VisitItemAss(i *ast.ItemAss[Data]) any {
	v := i.Expr.VisitExpr(c).(quadruple.Val)
	asVar := c.toVar(v)
	correctTypeVar := c.castIfNeeded(asVar, toQuadrupleType(c.currType, c.program.Objects)).(quadruple.Var)

	c.blockCtx.DeclareVar(i.Id, var_{
		val: correctTypeVar,
	})

	return asVar
}

func (c *converter[Data]) toVar(v quadruple.Val) quadruple.Var {
	asVar, isVar := v.(quadruple.Var)
	if !isVar {
		asVar = c.newVar(v.Type())
		c.appendInstructions(&quadruple.Assignment{
			Var:  asVar,
			Expr: &quadruple.EVal{Val: v},
		})
	}

	return asVar
}

func (c *converter[Data]) toQuadrupleType(t types.Type) quadruple.Type {
	return toQuadrupleType(t, c.program.Objects)
}

func (c *converter[Data]) VisitFunc(f *ast.Func[Data]) any {
	c.resetLabels()

	c.enterBlock()
	defer c.exitBlock()

	c.currFn = c.program.Fns[f.Id]

	for _, arg := range c.currFn.Arguments() {
		c.blockCtx.DeclareVar(arg.Name, var_{
			val: quadruple.Var{
				Type_: arg.Type,
				Id:    fmt.Sprintf("%s", arg.Name),
			},
		})
	}

	c.setBlock(c.newBlock(), nil)
	f.Block.VisitBlock(c)
	c.endBlock()

	return nil
}

func (c *converter[Data]) VisitClass(c2 *ast.Class[Data]) any {
	c.enterClass(c2.Id)
	defer c.exitClass()

	for _, el := range c2.ClassElements {
		el.VisitClassElement(c)
	}

	return nil
}

func (c *converter[Data]) VisitProgram(p *ast.Program[Data]) any {
	for _, topDef := range p.TopDefs {
		topDef.VisitTopDef(c)
	}
	return nil
}

func (c *converter[Data]) enterBlock() {
	c.blockCtx = &blockCtx{
		prev: c.blockCtx,
		vars: make(map[string]var_),
	}
}

func (c *converter[Data]) exitBlock() {
	c.blockCtx = c.blockCtx.prev
}

type var_ struct {
	val  quadruple.Var
	self *quadruple.Var
}

type blockCtx struct {
	prev *blockCtx
	vars map[string]var_
}

func (b *blockCtx) GetVar(varName string) var_ {
	t, ok := b.vars[varName]
	if !ok {
		return b.prev.GetVar(varName)
	}
	return t
}

func (b *blockCtx) DeclareVar(varName string, v var_) {
	b.vars[varName] = v
}

func (c *converter[Data]) convertAttributes(self *quadruple.Var, attrs []quadruple.Type, attrIdxMap map[string]int) map[string]var_ {
	ret := make(map[string]var_, len(attrIdxMap))
	for name, idx := range attrIdxMap {
		ret[name] = var_{
			val:  c.newVar(attrs[idx]),
			self: self,
		}
	}

	return ret
}

func (c *converter[Data]) enterClass(name string) {
	obj := c.program.Objects[name]

	self := quadruple.Var{
		Type_: quadruple.Type{
			Type:   quadruple.ObjectType,
			Object: obj,
		},
		Id: "self",
	}

	c.self = &self

	c.blockCtx = &blockCtx{
		prev: c.blockCtx,
		vars: c.convertAttributes(&self, obj.Attributes, obj.AttributeToIdx),
	}
}

func (c *converter[Data]) exitClass() {
	c.blockCtx = c.blockCtx.prev

	c.self = nil
}

func (c *converter[Data]) appendInstructions(instructions ...quadruple.Instruction) {
	c.currBlockInfo.block.Instructions = append(c.currBlockInfo.block.Instructions, instructions...)
}

func (c *converter[Data]) newBlock() *quadruple.Block {
	c.blockCounter++
	b := &quadruple.Block{
		Id:           c.blockCounter,
		Instructions: nil,
	}
	c.currFn.Blocks = append(c.currFn.Blocks, b)
	return b
}

func (c *converter[Data]) setBlock(b, jumpTo *quadruple.Block) {
	c.currBlockInfo = &currBlockInfo{block: b, jumpTo: jumpTo}
}

func (c *converter[Data]) currJumpTo() *quadruple.Block {
	return c.currBlockInfo.jumpTo
}

func (c *converter[Data]) markBlockTerminates() {
	c.currBlockInfo.terminates = true
}

func (c *converter[Data]) endBlock() {
	if !c.currBlockInfo.terminates {
		if c.currBlockInfo.jumpTo == nil {
			if c.currFn.RetType().Type == quadruple.VoidType {
				c.appendInstructions(&quadruple.Return{})
			} else {
				c.appendInstructions(&quadruple.Unreachable{})
			}
		} else {
			c.appendInstructions(&quadruple.Goto{Block: c.currBlockInfo.jumpTo})
		}
	}
	c.currBlockInfo = nil
}

func (c *converter[Data]) getNull(currType types.Type) quadruple.Val {
	return quadruple.NullVal{Type_: c.toQuadrupleType(currType)}
}

func (c *converter[Data]) attrAss(obj quadruple.Val, attr string, ex quadruple.Val) any {
	objType := obj.Type()

	attrIdx := objType.Object.AttributeToIdx[attr]
	op := &quadruple.EGetAttrPtr{
		Obj:     obj,
		AttrIdx: attrIdx,
	}
	t := objType.Object.Attributes[attrIdx]
	t.Pointer = true
	ptrV := c.newVar(t)
	ptr := &quadruple.Assignment{
		Var:  ptrV,
		Expr: op,
	}
	exprV := c.castIfNeeded(ex, objType.Object.Attributes[attrIdx])
	ptrAss := &quadruple.PointerAssignment{
		Pointer: ptrV,
		Val:     exprV,
	}

	c.appendInstructions(ptr, ptrAss)

	return nil
}

func (c *converter[Data]) castIfNeeded(from quadruple.Val, to quadruple.Type) quadruple.Val {
	if from.Type().Object != to.Object || from.Type().Type == quadruple.UntypedNullType {
		v := c.newVar(to)
		ass := &quadruple.Assignment{
			Var: v,
			Expr: &quadruple.ECast{
				Val: from,
				To:  to,
			},
		}
		c.appendInstructions(ass)
		return v
	}
	return from
}

func toQuadrupleType(t types.Type, objects map[string]*quadruple.Object) quadruple.Type {
	if t == types.Null {
		return quadruple.Type{Type: quadruple.UntypedNullType}
	}
	if t.Struct {
		return quadruple.Type{
			Type:   quadruple.ObjectType,
			Object: objects[t.Name],
		}
	}
	if t.Array != nil {
		ty := toQuadrupleType(*t.Array, objects)
		ty.Array = true
		return ty
	}
	switch t {
	case types.Int:
		return quadruple.Type{Type: quadruple.IntType}
	case types.Bool:
		return quadruple.Type{Type: quadruple.BoolType}
	case types.Void:
		return quadruple.Type{Type: quadruple.VoidType}
	case types.String:
		return quadruple.Type{Type: quadruple.StringType}
	default:
		panic("IMPOSSIBLE")
	}
}
