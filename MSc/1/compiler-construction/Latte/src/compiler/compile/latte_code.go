package compile

import (
	"github.com/llir/llvm/ir"
	"os"
	"os/exec"
	"strings"
)

type LatteCode struct {
	module *ir.Module
}

func (l *LatteCode) getCode() string {
	return l.module.String()
}

func (l *LatteCode) Save(originalName string) error {
	llPath := getLLPath(originalName)
	if err := os.WriteFile(llPath, []byte(l.getCode()), 0644); err != nil {
		return err
	}

	if err := compile(llPath); err != nil {
		return err
	}
	return link(getBCPath(originalName))
}

func compile(path string) error {
	cmd := exec.Command("llvm-as", path)
	return cmd.Run()
}

func link(path string) error {
	const runtimePath string = "./lib/runtime.bc"
	cmd := exec.Command("llvm-link", "-o", path, runtimePath, path)
	return cmd.Run()
}

func getLLPath(name string) string {
	withoutExtension := strings.TrimSuffix(name, ".lat")
	return withoutExtension + ".ll"
}

func getBCPath(name string) string {
	withoutExtension := strings.TrimSuffix(name, ".lat")
	return withoutExtension + ".bc"
}
