package footballKit

import (
	"fmt"
	"testing"
)

func Test_004Time(t *testing.T) {
	//1.09
	//0.88
	//0.9
	a, b := EurOdds2AsiaOdds(1.09, 0.88, 0.9)
	fmt.Println(a, b)
}
