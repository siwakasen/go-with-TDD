// package multiplereturnfunctions
package multiplereturnfunctions

func getMinMax(n []int) (int, int) {
	var min, max int

	for i, v := range n {
		if i == 0 {
			min, max = v, v
		} else if v > max {
			max = v
		} else if v < min {
			min = v
		}
	}
	return min, max
}
