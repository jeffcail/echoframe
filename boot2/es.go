package boot2

import "github.com/echoframe/pkg/es"

// InitEs
func InitEs(url string) {
	es.NewEs(url)
}
