package ssa_converter

import (
	"github.com/badochov/MRJP/Latte/src/IR/quadruple"
)

func fillPhis(cfg *cfg, bD map[*quadruple.Block]blockData, phis map[*quadruple.Block][]phiData) {
	for block, phisData := range phis {
		data := make([]blockData, 0, len(cfg.edgesIn[block]))
		for _, bIn := range cfg.edgesIn[block] {
			data = append(data, bD[bIn])
		}
		fillPhisInBlock(phisData, data)
	}
}

func fillPhisInBlock(phisData []phiData, blockData []blockData) {
	for _, phi := range phisData {
		fillPhiInBlock(phi, blockData)
	}
}

func fillPhiInBlock(phi phiData, data []blockData) {
	phi.phi.Vals = make([]quadruple.PhiValData, 0, len(data))
	for _, bD := range data {
		if replacement, ok := bD.varReplacements[phi.var_]; ok {
			phi.phi.Vals = append(phi.phi.Vals, quadruple.PhiValData{Block: bD.block, Val: replacement})
		} else {
			phi.phi.Vals = nil
			return
		}
	}
}
