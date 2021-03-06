// this file was auto-generated by internal/cmd/gentypes/main.go: DO NOT EDIT

package jwa

import (
	"github.com/pkg/errors"
)

//  EllipticCurveAlgorithm represents the algorithms used for EC keys
type EllipticCurveAlgorithm string

// Supported values for EllipticCurveAlgorithm
const (
	P256 EllipticCurveAlgorithm = "P-256"
	P384 EllipticCurveAlgorithm = "P-384"
	P521 EllipticCurveAlgorithm = "P-521"
)

// Accept is used when conversion from values given by
// outside sources (such as JSON payloads) is required
func (v *EllipticCurveAlgorithm) Accept(value interface{}) error {
	var tmp EllipticCurveAlgorithm
	switch x := value.(type) {
	case string:
		tmp = EllipticCurveAlgorithm(x)
	case EllipticCurveAlgorithm:
		tmp = x
	default:
		return errors.Errorf(`invalid type for jwa.EllipticCurveAlgorithm: %T`, value)
	}
	switch tmp {
	case P256, P384, P521:
	default:
		return errors.Errorf(`invalid jwa.EllipticCurveAlgorithm value`)
	}

	*v = tmp
	return nil
}

// String returns the string representation of a EllipticCurveAlgorithm
func (v EllipticCurveAlgorithm) String() string {
	return string(v)
}
