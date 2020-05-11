// +build pbc

package service

import (
	"github.com/unicredit/mosaic/abe"
)

// get curve on pbc library and stores it
func GetPbcCurve(curveStr string) string {
	curveJson := abe.GetPbcCurve(curveStr)
	StoreCurve(curveStr, curveJson)
	return curveJson
}
