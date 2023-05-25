package boot

import "github.com/echoframe/pkg/es"

// InitEs
func InitEs(url string) {
	es.NewEs(url)
}
