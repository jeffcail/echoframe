package vm

import (
	"github.com/jeffcail/echoframe/g"
	"github.com/jeffcail/gtools"
	"go.uber.org/zap"
)

func newLogger() *zap.Logger {
	pr, err := findProjectRoot()
	if err != nil {
		panic(err)
	}
	m := g.GM.Get("logger").(map[string]interface{})
	p, ok := m["path"].(string)
	if !ok {
		panic(ok)
	}
	p = gtools.CompactStr(pr, p)

	mx, ok := m["max"].(int)
	if !ok {
		panic(ok)
	}

	li, ok := m["live"].(int)
	if !ok {
		panic(ok)
	}
	co, ok := m["compress"].(bool)
	if !ok {
		panic(ok)
	}
	lo, ok := m["localtime"].(bool)
	if !ok {
		panic(ok)
	}
	return gtools.NewLogger(p, mx, li, lo, co)
}
