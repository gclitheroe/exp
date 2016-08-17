// +build devmode

package token

// allow running with no TLS for development when compiled with
// go build -tags devmode ...

func init() {
	secure = false
}
