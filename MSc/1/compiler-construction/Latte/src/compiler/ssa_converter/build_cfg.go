package ssa_converter

import (
	"github.com/badochov/MRJP/Latte/src/IR/quadruple"
)

func buildCfg(fn *quadruple.Fn) *cfg {
	if len(fn.Blocks) == 0 {
		return nil
	}

	return newCfgBuilder().build(fn)
}

func newCfgBuilder() *cfgBuilder {
	return &cfgBuilder{
		cfg: newCfg(),
	}
}

type cfgBuilder struct {
	currBlock *quadruple.Block
	cfg       *cfg
}

func (c *cfgBuilder) build(fn *quadruple.Fn) *cfg {
	c.cfg.root = fn.Blocks[0]
	for _, b := range fn.Blocks {
		c.VisitBlock(b)
	}

	return c.cfg
}

func (c *cfgBuilder) VisitBlock(b *quadruple.Block) any {
	c.currBlock = b

	for _, inst := range b.Instructions {
		inst.Accept(c)
	}

	return nil
}

func (c *cfgBuilder) VisitUnreachable(u *quadruple.Unreachable) any {
	return nil
}

func (c *cfgBuilder) VisitGoto(g *quadruple.Goto) any {
	c.markGoesTo(g.Block)

	return nil
}

func (c *cfgBuilder) VisitPointerAssignment(pa *quadruple.PointerAssignment) any {
	return nil
}

func (c *cfgBuilder) VisitAssignment(a *quadruple.Assignment) any {
	return nil
}

func (c *cfgBuilder) VisitReturn(r *quadruple.Return) any {
	return nil
}

func (c *cfgBuilder) VisitVoidCall(fc *quadruple.VoidCall) any {
	return nil
}

func (c *cfgBuilder) VisitCondJump(cj *quadruple.CondJump) any {
	c.markGoesTo(cj.TrueBlock)
	c.markGoesTo(cj.FalseBlock)

	return nil
}

func (c *cfgBuilder) markGoesTo(b *quadruple.Block) {
	c.cfg.addEdge(c.currBlock, b)
}

type cfg struct {
	root     *quadruple.Block
	edgesIn  map[*quadruple.Block][]*quadruple.Block
	edgesOut map[*quadruple.Block][]*quadruple.Block
}

func newCfg() *cfg {
	return &cfg{
		edgesIn:  make(map[*quadruple.Block][]*quadruple.Block),
		edgesOut: make(map[*quadruple.Block][]*quadruple.Block),
	}
}

func (c *cfg) addEdge(from *quadruple.Block, to *quadruple.Block) {
	edgesIn := c.edgesIn[to]
	c.edgesIn[to] = append(edgesIn, from)

	edgesTo := c.edgesOut[from]
	c.edgesOut[from] = append(edgesTo, to)
}

var _ quadruple.InstructionVisitor = &cfgBuilder{}
var _ quadruple.BlockVisitor = &cfgBuilder{}
