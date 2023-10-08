package tree_simplifier

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	ast2 "github.com/badochov/MRJP/Latte/src/IR/ast"
	"github.com/badochov/MRJP/Latte/src/frontend/types"
	"github.com/badochov/MRJP/Latte/src/status"
	"github.com/badochov/MRJP/Latte/src/topdefs/signatures"
	"golang.org/x/exp/maps"
)

func Simplify[Data antlr.ParserRuleContext](tree ast2.IProgram[Data], topDefs signatures.TopDefs) ast2.IProgram[Data] {
	tree = simplifyConst(tree, topDefs)
	return tree
}

func simplifyConst[Data antlr.ParserRuleContext](tree ast2.IProgram[Data], topDefs signatures.TopDefs) ast2.IProgram[Data] {
	return tree.VisitProgram(newConstSimplifier[Data](topDefs)).(ast2.IProgram[Data])
}

type constSimplifier[Data antlr.ParserRuleContext] struct {
	blockCtx     *blockCtx
	currType     types.Type
	loopCounter  int
	ifCounter    int
	classes      map[string]signatures.Class
	currClassVar *var_
}

func (c *constSimplifier[Data]) VisitClassAttribute(attribute *ast2.Attribute[Data]) any {
	return attribute
}

func (c *constSimplifier[Data]) VisitClassMethod(m *ast2.Method[Data]) any {
	ret := &ast2.Method[Data]{
		Data: m.Data,
		Type: m.Type,
		Id:   m.Id,
	}

	c.copyBlock()
	defer c.exitBlock()
	for _, arg := range m.Args {
		ret.Args = append(ret.Args, arg.VisitArg(c).(ast2.IArg[Data]))
	}

	c.blockCtx.DeclareVar("self", *c.currClassVar)

	ret.Block = m.Block.VisitBlock(c).(ast2.IBlock[Data])
	return ret
}

func (c *constSimplifier[Data]) VisitBlock(b *ast2.Block[Data]) any {
	block := &ast2.Block[Data]{
		Data: b.Data,
	}

	c.enterBlock()
	defer c.exitBlock()

	for _, stmt := range b.Stmts {
		block.Stmts = append(block.Stmts, stmt.VisitStmt(c).(ast2.IStmt[Data]))
	}

	return block
}

func (c *constSimplifier[Data]) VisitStmtEmpty(s *ast2.StmtEmpty[Data]) any {
	return s
}

func (c *constSimplifier[Data]) VisitStmtBlock(s *ast2.StmtBlock[Data]) any {
	return &ast2.StmtBlock[Data]{
		Data:  s.Data,
		Block: s.Block.VisitBlock(c).(ast2.IBlock[Data]),
	}
}

func (c *constSimplifier[Data]) VisitStmtDecl(s *ast2.StmtDecl[Data]) any {
	ret := &ast2.StmtDecl[Data]{
		Data: s.Data,
		Type: s.Type,
	}
	c.currType = ast2.GetType(s.Type)
	for _, i := range s.Items {
		ret.Items = append(ret.Items, i.VisitItem(c).(ast2.IItem[Data]))
	}

	return ret
}

func (c *constSimplifier[Data]) VisitStmtAss(s *ast2.StmtAss[Data]) any {
	expr := s.Expr.VisitExpr(c).(ast2.IExpr[Data])
	if c.inIf() || c.inLoop() {
		c.blockCtx.EraseVarVal(s.Id)
	}
	c.blockCtx.UpdateVarVal(s.Id, getVal(expr))

	return &ast2.StmtAss[Data]{
		Data: s.Data,
		Id:   s.Id,
		Expr: expr,
	}
}

func (c *constSimplifier[Data]) VisitStmtRet(s *ast2.StmtRet[Data]) any {
	return &ast2.StmtRet[Data]{
		Data: s.Data,
		Expr: s.Expr.VisitExpr(c).(ast2.IExpr[Data]),
	}
}

func (c *constSimplifier[Data]) VisitStmtVRet(s *ast2.StmtVRet[Data]) any {
	return &ast2.StmtVRet[Data]{
		Data: s.Data,
	}
}

func (c *constSimplifier[Data]) VisitStmtCond(s *ast2.StmtCond[Data]) any {
	cond := s.Cond.VisitExpr(c).(ast2.IExpr[Data])

	c.enterIf()
	defer c.exitIf()

	if b, ok := cond.(*ast2.ExprBool[Data]); ok {
		if b.Bool {
			return s.Stmt.VisitStmt(c).(ast2.IStmt[Data])
		}
		return &ast2.StmtEmpty[Data]{
			Data: s.Data,
		}
	}
	return &ast2.StmtCond[Data]{
		Data: s.Data,
		Cond: cond,
		Stmt: s.Stmt.VisitStmt(c).(ast2.IStmt[Data]),
	}
}

func (c *constSimplifier[Data]) VisitStmtCondElse(s *ast2.StmtCondElse[Data]) any {
	cond := s.Cond.VisitExpr(c).(ast2.IExpr[Data])

	c.enterIf()
	defer c.exitIf()

	if b, ok := cond.(*ast2.ExprBool[Data]); ok {
		if b.Bool {
			return s.Stmt.VisitStmt(c).(ast2.IStmt[Data])
		}
		return s.ElseStmt.VisitStmt(c).(ast2.IStmt[Data])
	}
	return &ast2.StmtCondElse[Data]{
		Data:     s.Data,
		Cond:     cond,
		Stmt:     s.Stmt.VisitStmt(c).(ast2.IStmt[Data]),
		ElseStmt: s.ElseStmt.VisitStmt(c).(ast2.IStmt[Data]),
	}
}

func (c *constSimplifier[Data]) enterLoop() {
	c.loopCounter++
}

func (c *constSimplifier[Data]) exitLoop() {
	c.loopCounter--
}

func (c *constSimplifier[Data]) inLoop() bool {
	return c.loopCounter != 0
}

func (c *constSimplifier[Data]) VisitStmtWhile(s *ast2.StmtWhile[Data]) any {
	c.enterLoop()
	defer c.exitLoop()

	cond := s.Cond.VisitExpr(c).(ast2.IExpr[Data])
	if b, ok := cond.(*ast2.ExprBool[Data]); ok {
		if !b.Bool {
			return &ast2.StmtEmpty[Data]{
				Data: s.Data,
			}
		}
	}

	return &ast2.StmtWhile[Data]{
		Data: s.Data,
		Cond: cond,
		Stmt: s.Stmt.VisitStmt(c).(ast2.IStmt[Data]),
	}
}

func (c *constSimplifier[Data]) VisitStmtExp(s *ast2.StmtExp[Data]) any {
	return &ast2.StmtExp[Data]{
		Data: s.Data,
		Expr: s.Expr.VisitExpr(c).(ast2.IExpr[Data]),
	}
}

func (c *constSimplifier[Data]) VisitStmtIndexAss(s *ast2.StmtIndexAss[Data]) any {
	return &ast2.StmtIndexAss[Data]{
		Data:      s.Data,
		ArrExpr:   s.ArrExpr.VisitExpr(c).(ast2.IExpr[Data]),
		IndexExpr: s.IndexExpr.VisitExpr(c).(ast2.IExpr[Data]),
		Expr:      s.Expr.VisitExpr(c).(ast2.IExpr[Data]),
	}
}

func (c *constSimplifier[Data]) VisitStmtAttrAss(s *ast2.StmtAttrAss[Data]) any {
	return &ast2.StmtAttrAss[Data]{
		Data:    s.Data,
		ObjExpr: s.ObjExpr.VisitExpr(c).(ast2.IExpr[Data]),
		Attr:    s.Attr,
		Expr:    s.Expr.VisitExpr(c).(ast2.IExpr[Data]),
	}
}

func (c *constSimplifier[Data]) VisitExprParen(e *ast2.ExprParen[Data]) any {
	return &ast2.ExprParen[Data]{
		Data: e.Data,
		Expr: e.Expr,
	}
}

func (c *constSimplifier[Data]) VisitExprUnOp(e *ast2.ExprUnOp[Data]) any {
	ex := e.Expr.VisitExpr(c).(ast2.IExpr[Data])

	switch e.UnOp {
	case ast2.UnaryMinus:
		if expr, ok := ex.(*ast2.ExprInt[Data]); ok {
			return &ast2.ExprInt[Data]{
				Data: expr.Data,
				Int:  -expr.Int,
			}
		}
	case ast2.LogicNot:
		if expr, ok := ex.(*ast2.ExprBool[Data]); ok {
			return &ast2.ExprBool[Data]{
				Data: expr.Data,
				Bool: !expr.Bool,
			}
		}
	}

	return &ast2.ExprUnOp[Data]{
		Data: e.Data,
		UnOp: e.UnOp,
		Expr: ex,
	}
}

func (c *constSimplifier[Data]) VisitExprMulOp(e *ast2.ExprMulOp[Data]) any {
	le := e.LExpr.VisitExpr(c).(ast2.IExpr[Data])
	re := e.RExpr.VisitExpr(c).(ast2.IExpr[Data])

	switch e.MulOp {
	case ast2.Times:
		if r, ok := re.(*ast2.ExprInt[Data]); ok {
			switch r.Int {
			case 1:
				return re
			case 0:
				return &ast2.ExprInt[Data]{
					Data: e.Data,
					Int:  0,
				}
			default:
				if l, ok := le.(*ast2.ExprInt[Data]); ok {
					return &ast2.ExprInt[Data]{
						Data: e.Data,
						Int:  r.Int * l.Int,
					}
				}
			}
		}
	case ast2.Divide:
		if r, ok := re.(*ast2.ExprInt[Data]); ok {
			switch r.Int {
			case 1:
				return le
			case 0:
				status.HandleError(fmt.Errorf("division by 0"), e.Data)
			default:
				if l, ok := le.(*ast2.ExprInt[Data]); ok {
					return &ast2.ExprInt[Data]{
						Data: e.Data,
						Int:  l.Int / r.Int,
					}
				}
			}
		}
	case ast2.Modulo:
		if r, ok := re.(*ast2.ExprInt[Data]); ok {
			switch r.Int {
			case 1:
				return &ast2.ExprInt[Data]{
					Data: e.Data,
					Int:  0,
				}
			case 0:
				status.HandleError(fmt.Errorf("modulo by 0"), e.Data)
			default:
				if l, ok := le.(*ast2.ExprInt[Data]); ok {
					return &ast2.ExprInt[Data]{
						Data: e.Data,
						Int:  l.Int % r.Int,
					}
				}
			}
		}
	}

	return &ast2.ExprMulOp[Data]{
		Data:  e.Data,
		MulOp: e.MulOp,
		LExpr: le,
		RExpr: re,
	}
}

func (c *constSimplifier[Data]) VisitExprAddOp(e *ast2.ExprAddOp[Data]) any {
	le := e.LExpr.VisitExpr(c).(ast2.IExpr[Data])
	re := e.RExpr.VisitExpr(c).(ast2.IExpr[Data])

	switch e.AddOp {
	case ast2.Minus:
		if r, ok := re.(*ast2.ExprInt[Data]); ok {
			switch r.Int {
			case 0:
				return le
			default:
				if l, ok := le.(*ast2.ExprInt[Data]); ok {
					return &ast2.ExprInt[Data]{
						Data: e.Data,
						Int:  l.Int - r.Int,
					}
				}
			}
		} else if l, ok := le.(*ast2.ExprInt[Data]); ok && l.Int == 0 {
			return &ast2.ExprUnOp[Data]{
				Data: e.Data,
				Expr: re,
				UnOp: ast2.UnaryMinus,
			}
		}
	case ast2.Plus:
		// Int
		if r, ok := re.(*ast2.ExprInt[Data]); ok {
			switch r.Int {
			case 0:
				return le
			default:
				if l, ok := le.(*ast2.ExprInt[Data]); ok {
					return &ast2.ExprInt[Data]{
						Data: e.Data,
						Int:  l.Int + r.Int,
					}
				}
			}
		} else if l, ok := le.(*ast2.ExprInt[Data]); ok && l.Int == 0 {
			return re
		}

		// String
		if r, ok := re.(*ast2.ExprString[Data]); ok {
			switch r.String {
			case "":
				return le
			default:
				if l, ok := le.(*ast2.ExprString[Data]); ok {
					return &ast2.ExprString[Data]{
						Data:   e.Data,
						String: l.String + r.String,
					}
				}
			}
		} else if l, ok := le.(*ast2.ExprString[Data]); ok && l.String == "" {
			return re
		}
	}

	return &ast2.ExprAddOp[Data]{
		Data:  e.Data,
		AddOp: e.AddOp,
		LExpr: le,
		RExpr: re,
	}
}

func (c *constSimplifier[Data]) VisitExprRelOp(e *ast2.ExprRelOp[Data]) any {
	le := e.LExpr.VisitExpr(c).(ast2.IExpr[Data])
	re := e.RExpr.VisitExpr(c).(ast2.IExpr[Data])

	if l, ok := le.(*ast2.ExprInt[Data]); ok {
		if r, ok := re.(*ast2.ExprInt[Data]); ok {
			switch e.RelOp {
			case ast2.Ne:
				return &ast2.ExprBool[Data]{
					Data: e.Data,
					Bool: l.Int != r.Int,
				}
			case ast2.Eq:
				return &ast2.ExprBool[Data]{
					Data: e.Data,
					Bool: l.Int == r.Int,
				}
			case ast2.Le:
				return &ast2.ExprBool[Data]{
					Data: e.Data,
					Bool: l.Int <= r.Int,
				}
			case ast2.Lt:
				return &ast2.ExprBool[Data]{
					Data: e.Data,
					Bool: l.Int < r.Int,
				}
			case ast2.Gt:
				return &ast2.ExprBool[Data]{
					Data: e.Data,
					Bool: l.Int > r.Int,
				}
			case ast2.Ge:
				return &ast2.ExprBool[Data]{
					Data: e.Data,
					Bool: l.Int >= r.Int,
				}
			}
		}
	}

	if l, ok := le.(*ast2.ExprBool[Data]); ok {
		if r, ok := re.(*ast2.ExprBool[Data]); ok {
			switch e.RelOp {
			case ast2.Ne:
				return &ast2.ExprBool[Data]{
					Data: e.Data,
					Bool: l.Bool != r.Bool,
				}
			case ast2.Eq:
				return &ast2.ExprBool[Data]{
					Data: e.Data,
					Bool: l.Bool == r.Bool,
				}
			}
		}
	}

	return &ast2.ExprRelOp[Data]{
		Data:  e.Data,
		LExpr: le,
		RExpr: re,
		RelOp: e.RelOp,
	}
}

func (c *constSimplifier[Data]) VisitExprAnd(e *ast2.ExprAnd[Data]) any {
	le := e.LExpr.VisitExpr(c).(ast2.IExpr[Data])
	re := e.RExpr.VisitExpr(c).(ast2.IExpr[Data])

	if l, ok := le.(*ast2.ExprBool[Data]); ok {
		if l.Bool {
			return re
		}
		return &ast2.ExprBool[Data]{
			Data: e.Data,
			Bool: false,
		}
	}
	if r, ok := re.(*ast2.ExprBool[Data]); ok {
		if r.Bool {
			return le
		}
		return &ast2.ExprBool[Data]{
			Data: e.Data,
			Bool: false,
		}
	}

	return &ast2.ExprAnd[Data]{
		Data:  e.Data,
		LExpr: le,
		RExpr: re,
	}
}

func (c *constSimplifier[Data]) VisitExprOr(e *ast2.ExprOr[Data]) any {
	le := e.LExpr.VisitExpr(c).(ast2.IExpr[Data])
	re := e.RExpr.VisitExpr(c).(ast2.IExpr[Data])

	if l, ok := le.(*ast2.ExprBool[Data]); ok {
		if l.Bool {
			return &ast2.ExprBool[Data]{
				Data: e.Data,
				Bool: true,
			}
		}
		return re
	}
	if r, ok := re.(*ast2.ExprBool[Data]); ok {
		if r.Bool {
			return &ast2.ExprBool[Data]{
				Data: e.Data,
				Bool: true,
			}
		}
		return le
	}
	return &ast2.ExprOr[Data]{
		Data:  e.Data,
		LExpr: le,
		RExpr: re,
	}
}

func (c *constSimplifier[Data]) VisitExprId(e *ast2.ExprId[Data]) any {
	if c.inLoop() {
		return &ast2.ExprId[Data]{
			Data: e.Data,
			Id:   e.Id,
		}
	}
	if v := c.blockCtx.GetVar(e.Id); v.val != nil {
		switch v.t {
		case types.Int:
			return &ast2.ExprInt[Data]{
				Data: e.Data,
				Int:  v.val.(int64),
			}
		case types.Bool:
			return &ast2.ExprBool[Data]{
				Data: e.Data,
				Bool: v.val.(bool),
			}
		case types.String:
			return &ast2.ExprString[Data]{
				Data:   e.Data,
				String: v.val.(string),
			}
		}
	}

	return e
}

func (c *constSimplifier[Data]) VisitExprInt(e *ast2.ExprInt[Data]) any {
	return e
}

func (c *constSimplifier[Data]) VisitExprBool(e *ast2.ExprBool[Data]) any {
	return e
}

func (c *constSimplifier[Data]) VisitExprString(e *ast2.ExprString[Data]) any {
	return e
}

func (c *constSimplifier[Data]) VisitExprFunCall(e *ast2.ExprFunCall[Data]) any {
	ret := &ast2.ExprFunCall[Data]{
		Data: e.Data,
		Id:   e.Id,
	}

	for _, param := range e.Params {
		ret.Params = append(ret.Params, param.VisitExpr(c).(ast2.IExpr[Data]))
	}

	return ret
}

func (c *constSimplifier[Data]) VisitExprIndex(e *ast2.ExprIndex[Data]) any {
	return &ast2.ExprIndex[Data]{
		Data:      e.Data,
		ArrExpr:   e.ArrExpr.VisitExpr(c).(ast2.IExpr[Data]),
		IndexExpr: e.IndexExpr.VisitExpr(c).(ast2.IExpr[Data]),
	}
}

func (c *constSimplifier[Data]) VisitExprNull(e *ast2.ExprNull[Data]) any {
	return e
}

func (c *constSimplifier[Data]) VisitExprNewClass(e *ast2.ExprNewClass[Data]) any {
	return e
}

func (c *constSimplifier[Data]) VisitExprMethodCall(e *ast2.ExprMethodCall[Data]) any {
	ret := &ast2.ExprMethodCall[Data]{
		Data:    e.Data,
		Id:      e.Id,
		ObjExpr: e.ObjExpr.VisitExpr(c).(ast2.IExpr[Data]),
	}

	for _, param := range e.Params {
		ret.Params = append(ret.Params, param.VisitExpr(c).(ast2.IExpr[Data]))
	}

	return ret
}

func (c *constSimplifier[Data]) VisitExprAttr(e *ast2.ExprAttr[Data]) any {
	return &ast2.ExprAttr[Data]{
		Data:    e.Data,
		ObjExpr: e.ObjExpr.VisitExpr(c).(ast2.IExpr[Data]),
		Attr:    e.Attr,
	}
}

func (c *constSimplifier[Data]) VisitExprNewArray(e *ast2.ExprNewArray[Data]) any {
	return &ast2.ExprNewArray[Data]{
		Data:     e.Data,
		Type:     e.Type,
		SizeExpr: e.SizeExpr.VisitExpr(c).(ast2.IExpr[Data]),
	}
}

func (c *constSimplifier[Data]) VisitItemAss(i *ast2.ItemAss[Data]) any {
	ex := i.Expr.VisitExpr(c).(ast2.IExpr[Data])

	c.blockCtx.DeclareVar(i.Id, var_{
		t:   c.currType,
		val: getVal(ex),
	})

	return &ast2.ItemAss[Data]{
		Data: i.Data,
		Id:   i.Id,
		Expr: ex,
	}
}

func getVal[Data antlr.ParserRuleContext](ex ast2.IExpr[Data]) any {
	switch expr := ex.(type) {
	case *ast2.ExprInt[Data]:
		return expr.Int
	case *ast2.ExprBool[Data]:
		return expr.Bool
	case *ast2.ExprString[Data]:
		return expr.String
	default:
		return nil
	}
}

func (c *constSimplifier[Data]) VisitArg(a *ast2.Arg[Data]) any {
	c.blockCtx.DeclareVar(a.Id, var_{
		t: ast2.GetType(a.Type),
	})

	return a
}

func (c *constSimplifier[Data]) VisitFunc(f *ast2.Func[Data]) any {
	ret := &ast2.Func[Data]{
		Data: f.Data,
		Type: f.Type,
		Id:   f.Id,
	}

	c.enterBlock()
	defer c.exitBlock()
	for _, arg := range f.Args {
		ret.Args = append(ret.Args, arg.VisitArg(c).(ast2.IArg[Data]))
	}

	ret.Block = f.Block.VisitBlock(c).(ast2.IBlock[Data])
	return ret
}

func (c *constSimplifier[Data]) VisitClass(c2 *ast2.Class[Data]) any {
	class := &ast2.Class[Data]{
		Data:    c2.Data,
		Id:      c2.Id,
		Extends: c2.Extends,
	}

	c.enterClass(c2.Id)
	defer c.exitClass(c2.Id)

	for _, el := range c2.ClassElements {
		class.ClassElements = append(class.ClassElements, el.VisitClassElement(c).(ast2.IClassElement[Data]))
	}

	return class
}

func (c *constSimplifier[Data]) VisitProgram(p *ast2.Program[Data]) any {
	ret := &ast2.Program[Data]{
		Data: p.Data,
	}

	for _, topDef := range p.TopDefs {
		ret.TopDefs = append(ret.TopDefs, topDef.VisitTopDef(c).(ast2.ITopDef[Data]))
	}

	return ret
}

func (c *constSimplifier[Data]) enterBlock() {
	c.blockCtx = &blockCtx{
		prev: c.blockCtx,
		vars: make(map[string]var_),
		fns:  make(map[string]signatures.Function),
	}
}

func (c *constSimplifier[Data]) copyBlock() {
	c.blockCtx = &blockCtx{
		prev: c.blockCtx,
		vars: maps.Clone(c.blockCtx.vars),
		fns:  maps.Clone(c.blockCtx.fns),
	}
}

func (c *constSimplifier[Data]) exitBlock() {
	c.blockCtx = c.blockCtx.prev
}

type var_ struct {
	t   types.Type
	val any
}

type blockCtx struct {
	prev *blockCtx
	vars map[string]var_
	fns  map[string]signatures.Function
}

func (b *blockCtx) GetVar(varName string) var_ {
	t, ok := b.vars[varName]
	if !ok {
		return b.prev.GetVar(varName)
	}
	return t
}

func (b *blockCtx) UpdateVarVal(varName string, val any) {
	t, ok := b.vars[varName]
	if !ok {
		b.prev.UpdateVarVal(varName, val)
		return
	}
	t.val = val
	b.vars[varName] = t
}

func (b *blockCtx) EraseVarVal(varName string) {
	if t, ok := b.vars[varName]; ok {
		t.val = nil
		b.vars[varName] = t
	}
	if b.prev != nil {
		b.prev.EraseVarVal(varName)
	}
}

func (b *blockCtx) DeclareVar(varName string, v var_) {
	b.vars[varName] = v
}

func (b *blockCtx) getVars() map[string]var_ {
	vars := make(map[string]var_)
	for c := b; c != nil; c = c.prev {
		maps.Copy(vars, c.vars)
	}
	return vars
}

func convertAttributes(attrs map[string]types.Type) map[string]var_ {
	ret := make(map[string]var_, len(attrs))
	for name, t := range attrs {
		ret[name] = var_{
			t: t,
		}
	}
	return ret
}

func (c *constSimplifier[Data]) enterClass(name string) {
	class := c.classes[name]
	parents := class.GetParents()
	for i := len(parents) - 1; i >= 0; i-- {
		c.blockCtx = &blockCtx{
			prev: c.blockCtx,
			vars: convertAttributes(parents[i].Attributes),
			fns:  parents[i].Methods,
		}
	}
	c.blockCtx = &blockCtx{
		prev: c.blockCtx,
		vars: convertAttributes(class.Attributes),
		fns:  class.Methods,
	}

	c.currClassVar = &var_{
		t: types.Type{
			Name:   name,
			Struct: true,
		},
	}
}

func (c *constSimplifier[Data]) exitClass(name string) {
	for range c.classes[name].GetParents() {
		c.blockCtx = c.blockCtx.prev
	}
	c.blockCtx = c.blockCtx.prev

	c.currClassVar = nil
}

func (c *constSimplifier[Data]) enterIf() {
	c.blockCtx = &blockCtx{
		prev: c.blockCtx,
		vars: c.blockCtx.getVars(),
		fns:  nil,
	}
	c.ifCounter++
}

func (c *constSimplifier[Data]) exitIf() {
	c.blockCtx = c.blockCtx.prev
	c.ifCounter--
}

func (c *constSimplifier[Data]) inIf() bool {
	return c.ifCounter != 0
}

func newConstSimplifier[Data antlr.ParserRuleContext](defs signatures.TopDefs) *constSimplifier[Data] {
	return &constSimplifier[Data]{
		blockCtx: &blockCtx{
			prev: nil,
			vars: make(map[string]var_),
			fns:  defs.Fns,
		},
		currType:    types.Type{},
		loopCounter: 0,
		classes:     defs.Classes,
	}
}
