// Approach from https://github.com/theckman/goconstraint/

// Package goversion should only be used as a blank import. If imported, it
// will only compile if the Go runtime version is >= 1.18.
package goversion

// This will fail to compile if the Go runtime version isn't >= 1.18.
var _ = __EVCC_REQUIRES_GO_VERSION_1_18__
