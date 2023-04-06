package boot

import "github.com/echo-scaffolding/pkg/es"

// InitEs
func InitEs(url string) {
	es.NewEs(url)
}
