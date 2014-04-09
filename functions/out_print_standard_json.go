package functions

import (
	"github.com/simonz05/graphite-ng/chains"
	"fmt"
)

func init() {
	Functions["printStandardJson"] = []string{"OutPrintStandardJson", "metric", "int"}
}

func OutPrintStandardJson(dep_el chains.ChainEl, until int32) {
	for {
		d := <-dep_el.Link
		fmt.Printf("[%f, %d]", d.Value, d.Ts)
		if d.Ts >= until {
			break
		} else {
			fmt.Printf(", ")
		}
	}

}
