package top150

func canCompleteCircuit(gas []int, cost []int) int {

	n := len(gas)

	for i := 0; i < n; {
		gasSum, costSum, cnt := 0, 0, 0
		for cnt < n {
			j := (i + cnt) % n
			gasSum += gas[j]
			costSum += cost[j]

			if gasSum < costSum {
				break
			}
			cnt++
		}

		if cnt == n {
			return i
		} else {
			i += cnt + 1
		}

	}
	return -1
}
