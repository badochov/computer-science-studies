package return_checker

import (
	"fmt"
	ast2 "github.com/badochov/MRJP/Latte/src/IR/ast"
	"github.com/badochov/MRJP/Latte/src/status"
)

func Check[Data any](program ast2.IProgram[Data]) {
	program.VisitProgram(returnChecker[Data]{})
}

type returnChecker[Data any] struct{}

func (r returnChecker[Data]) VisitStmtIndexAss(s *ast2.StmtIndexAss[Data]) any {
	return false
}

func (r returnChecker[Data]) VisitStmtAttrAss(s *ast2.StmtAttrAss[Data]) any {
	return false
}

func (r returnChecker[Data]) VisitBlock(b *ast2.Block[Data]) any {
	for _, stmt := range b.Stmts {
		if stmt.VisitStmt(r).(bool) {
			return true
		}
	}
	return false
}

func (r returnChecker[Data]) VisitStmtEmpty(s *ast2.StmtEmpty[Data]) any {
	return false
}

func (r returnChecker[Data]) VisitStmtBlock(s *ast2.StmtBlock[Data]) any {
	return s.Block.VisitBlock(r)
}

func (r returnChecker[Data]) VisitStmtDecl(s *ast2.StmtDecl[Data]) any {
	return false
}

func (r returnChecker[Data]) VisitStmtAss(s *ast2.StmtAss[Data]) any {
	return false
}

func (r returnChecker[Data]) VisitStmtRet(s *ast2.StmtRet[Data]) any {
	return true
}

func (r returnChecker[Data]) VisitStmtVRet(s *ast2.StmtVRet[Data]) any {
	return true
}

func (r returnChecker[Data]) VisitStmtCond(s *ast2.StmtCond[Data]) any {
	return false
}

func (r returnChecker[Data]) VisitStmtCondElse(s *ast2.StmtCondElse[Data]) any {
	return s.Stmt.VisitStmt(r).(bool) && s.ElseStmt.VisitStmt(r).(bool)
}

func (r returnChecker[Data]) VisitStmtWhile(s *ast2.StmtWhile[Data]) any {
	if b, ok := s.Cond.(*ast2.ExprBool[Data]); ok && b.Bool {
		return s.Stmt.VisitStmt(r)
	}
	return false
}

func (r returnChecker[Data]) VisitStmtExp(s *ast2.StmtExp[Data]) any {
	return false
}

func (r returnChecker[Data]) VisitFunc(f *ast2.Func[Data]) any {
	if _, isVoid := f.Type.(*ast2.TypeVoid[Data]); !isVoid {
		if !f.Block.VisitBlock(r).(bool) {
			status.Fatal(fmt.Errorf("function '%s', doesn't always return", f.Id))
		}
	}
	return nil
}

func (r returnChecker[Data]) VisitClass(c *ast2.Class[Data]) any {
	for _, el := range c.ClassElements {
		el.VisitClassElement(r)
	}
	return nil
}

func (r returnChecker[Data]) VisitClassAttribute(attribute *ast2.Attribute[Data]) any {
	return nil
}

func (r returnChecker[Data]) VisitClassMethod(m *ast2.Method[Data]) any {
	if _, isVoid := m.Type.(*ast2.TypeVoid[Data]); !isVoid {
		if !m.Block.VisitBlock(r).(bool) {
			status.Fatal(fmt.Errorf("method '%s', doesn't always return", m.Id))
		}
	}
	return nil
}

func (r returnChecker[Data]) VisitProgram(p *ast2.Program[Data]) any {
	for _, topDef := range p.TopDefs {
		topDef.VisitTopDef(r)
	}
	return nil
}
