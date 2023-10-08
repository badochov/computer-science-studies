package ast

type IStmt[Data any] interface {
	VisitStmt(visitor StmtVisitor[Data]) any
}

type StmtVisitor[Data any] interface {
	VisitStmtEmpty(s *StmtEmpty[Data]) any
	VisitStmtBlock(s *StmtBlock[Data]) any
	VisitStmtDecl(s *StmtDecl[Data]) any
	VisitStmtAss(s *StmtAss[Data]) any
	VisitStmtRet(s *StmtRet[Data]) any
	VisitStmtVRet(s *StmtVRet[Data]) any
	VisitStmtCond(s *StmtCond[Data]) any
	VisitStmtCondElse(s *StmtCondElse[Data]) any
	VisitStmtWhile(s *StmtWhile[Data]) any
	VisitStmtExp(s *StmtExp[Data]) any
	VisitStmtIndexAss(s *StmtIndexAss[Data]) any
	VisitStmtAttrAss(s *StmtAttrAss[Data]) any
}

type StmtEmpty[Data any] struct {
	Data Data
}

func (s *StmtEmpty[Data]) VisitStmt(visitor StmtVisitor[Data]) any {
	return visitor.VisitStmtEmpty(s)
}

type StmtBlock[Data any] struct {
	Data Data

	Block IBlock[Data]
}

func (s *StmtBlock[Data]) VisitStmt(visitor StmtVisitor[Data]) any {
	return visitor.VisitStmtBlock(s)
}

type StmtDecl[Data any] struct {
	Data Data

	Type  IType[Data]
	Items []IItem[Data]
}

func (s *StmtDecl[Data]) VisitStmt(visitor StmtVisitor[Data]) any {
	return visitor.VisitStmtDecl(s)
}

type StmtAss[Data any] struct {
	Data Data

	Id   string
	Expr IExpr[Data]
}

func (s *StmtAss[Data]) VisitStmt(visitor StmtVisitor[Data]) any {
	return visitor.VisitStmtAss(s)
}

type StmtRet[Data any] struct {
	Data Data

	Expr IExpr[Data]
}

func (s *StmtRet[Data]) VisitStmt(visitor StmtVisitor[Data]) any {
	return visitor.VisitStmtRet(s)
}

type StmtVRet[Data any] struct {
	Data Data

	Expr IExpr[Data]
}

func (s *StmtVRet[Data]) VisitStmt(visitor StmtVisitor[Data]) any {
	return visitor.VisitStmtVRet(s)
}

type StmtCond[Data any] struct {
	Data Data

	Cond IExpr[Data]
	Stmt IStmt[Data]
}

func (s *StmtCond[Data]) VisitStmt(visitor StmtVisitor[Data]) any {
	return visitor.VisitStmtCond(s)
}

type StmtCondElse[Data any] struct {
	Data Data

	Cond     IExpr[Data]
	Stmt     IStmt[Data]
	ElseStmt IStmt[Data]
}

func (s *StmtCondElse[Data]) VisitStmt(visitor StmtVisitor[Data]) any {
	return visitor.VisitStmtCondElse(s)
}

type StmtWhile[Data any] struct {
	Data Data

	Cond IExpr[Data]
	Stmt IStmt[Data]
}

func (s *StmtWhile[Data]) VisitStmt(visitor StmtVisitor[Data]) any {
	return visitor.VisitStmtWhile(s)
}

type StmtExp[Data any] struct {
	Data Data

	Expr IExpr[Data]
}

func (s *StmtExp[Data]) VisitStmt(visitor StmtVisitor[Data]) any {
	return visitor.VisitStmtExp(s)
}

type StmtIndexAss[Data any] struct {
	Data Data

	ArrExpr   IExpr[Data]
	IndexExpr IExpr[Data]
	Expr      IExpr[Data]
}

func (s *StmtIndexAss[Data]) VisitStmt(visitor StmtVisitor[Data]) any {
	return visitor.VisitStmtIndexAss(s)
}

type StmtAttrAss[Data any] struct {
	Data Data

	ObjExpr IExpr[Data]
	Attr    string
	Expr    IExpr[Data]
}

func (s *StmtAttrAss[Data]) VisitStmt(visitor StmtVisitor[Data]) any {
	return visitor.VisitStmtAttrAss(s)
}

var _ IStmt[any] = &StmtEmpty[any]{}
var _ IStmt[any] = &StmtBlock[any]{}
var _ IStmt[any] = &StmtDecl[any]{}
var _ IStmt[any] = &StmtAss[any]{}
var _ IStmt[any] = &StmtRet[any]{}
var _ IStmt[any] = &StmtVRet[any]{}
var _ IStmt[any] = &StmtCond[any]{}
var _ IStmt[any] = &StmtCondElse[any]{}
var _ IStmt[any] = &StmtWhile[any]{}
var _ IStmt[any] = &StmtExp[any]{}
var _ IStmt[any] = &StmtAttrAss[any]{}
var _ IStmt[any] = &StmtIndexAss[any]{}
