package numbers

const preamble = 25

func inSum(number int, numbers []int) (bool, []int) {
	n := []int{}

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[i] != numbers[j] {
				if numbers[i]+numbers[j] == number {
					n = append(n, numbers[i], numbers[j])
					return true, n
				}
			}
		}
	}

	if len(n) > 0 {
		return true, n
	}

	return false, nil
}

func FindWeakness(number int, numbers []int) int {
	min, max := 0, 0

	ncompares := 0
ii:
	for i := 0; i < len(numbers); i++ {
		sum := 0

		for j := i; j < len(numbers); j++ {
			sum += numbers[j]
			ncompares++
			if number == sum {
				min, max = 9999999999, -1
				for k := i; k <= j; k++ {
					if numbers[k] < min {
						min = numbers[k]
					}
					if numbers[k] > max {
						max = numbers[k]
					}
				}
				break ii
			}
		}
	}

	return min + max
}

func FindNumber(numbers []int) int {
	counter := preamble

	for {
		if counter < len(numbers) {
			c := counter / preamble
			ok, _ := inSum(numbers[counter], numbers[((c-1)*preamble)+(counter%preamble):(c*preamble)+(counter%preamble)])

			if !ok {
				return numbers[counter]
			}

			counter++
			continue
		}

		break
	}

	return -1
}
