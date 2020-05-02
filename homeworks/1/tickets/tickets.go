package tickets

// Calculate calculate number of lucky tickets
// n is length of half of length of ticket.
func Calculate(n int) int64 {
	if n <= 0 {
		return 0
	}

	combinations := calculateCombinationsTable(n)

	// go-mnd linter doesn't like magic numbers, so use var
	maxDigit := 9

	// further code deals with n only
	maxSum := maxDigit * n

	// square combinations because original task is calculate combinations of lucky tickets
	// sum of left half of ticket must be equal sum of right half of ticket
	// each half is independent
	//		for example for n = 2 and sum=3:
	//			0+3 = 0+3
	//			0+3 = 1+2
	//			0+3 = 2+1
	//			0+3 = 3+0
	//			1+2 = 0+3
	//			...
	//			3+0 = 3+0
	//		total is 16 combinations for whole ticket of length 2*n (square combination for half)
	combinationsForTicket := make([]int64, maxSum+1)

	for sum := 0; sum <= maxSum; sum++ {
		combinationsForHalf := combinations[n-1][sum]
		combinationsForTicket[sum] = combinationsForHalf * combinationsForHalf
	}

	// summing all combinations for each possible sums
	totalCombinations := int64(0)

	for sum := 0; sum <= maxSum; sum++ {
		totalCombinations += combinationsForTicket[sum]
	}

	return totalCombinations
}

// calculateCombinationsTable calculate dynamic programming calculation table
// keep possible combinations of k terms for concrete sum
// sum goes from 0 to 9*k (for example for k = 3 max sum of 3 terms is 27 = 9 + 9 +9).
func calculateCombinationsTable(n int) [][]int64 {
	// go-mnd linter doesn't like magic numbers, so use var
	maxDigit := 9

	// dynamic programming calculation table (aka cache), to prevent exponent growth
	var combinations [][]int64

	// calculate combinations and fill the table
	for k := 1; k <= n; k++ {
		// max sum possible with k digits
		maxSum := maxDigit * k

		// allocate vector of need length
		vector := make([]int64, maxSum+1)

		// k starts from 1, bun in combinations index starts from 0
		combinations = append(combinations, vector)

		// k=1 it is easy case
		if k == 1 {
			for sum := 0; sum <= maxSum; sum++ {
				combinations[k-1][sum] = 1
			}

			continue
		}

		for sum := 0; sum <= maxSum; sum++ {
			maxSumForPrevK := maxSum - maxDigit

			minPossibleDigit := sum - maxSumForPrevK
			if minPossibleDigit <= 0 {
				minPossibleDigit = 0
			}

			maxPossibleDigit := maxDigit
			if sum <= maxDigit {
				maxPossibleDigit = sum
			}

			combinationsForCurrentK := int64(0)

			for d := minPossibleDigit; d <= maxPossibleDigit; d++ {
				restSum := sum - d
				combinationsForPrevK := combinations[k-2][restSum]
				combinationsForCurrentK += combinationsForPrevK
			}

			// save calculated combination for current sum and current k
			combinations[k-1][sum] = combinationsForCurrentK
		}
	}

	return combinations
}
