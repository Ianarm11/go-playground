package finance

import (
	"fmt"
	"strconv"
)

const returnrate = .09

func Finance() {
	fmt.Println("I'm going to convince you that investing in the market with a small amount of money is not the best decision.")
	fmt.Println("********************")

	intialYearlyInvestment := 12000
	runningTotal := 0

	fmt.Println("Starting with " + strconv.Itoa(intialYearlyInvestment) + " in the merket.")
	fmt.Println("Assuming a 9% return each year and your yearly investment total is increased by $2000.")
	fmt.Println("----------------------------------")

	for i := 1; i < 41; i++ {
		fmt.Println("Year: " + strconv.Itoa(i))
		fmt.Println("Running profit: " + strconv.Itoa(runningTotal))
		fmt.Println("Total amount of money in the market: " + strconv.Itoa(intialYearlyInvestment))

		profit := ComputeProfitForTheYear(intialYearlyInvestment)

		fmt.Println("Profit made this year: " + strconv.Itoa(profit))

		runningTotal = runningTotal + profit
		intialYearlyInvestment = intialYearlyInvestment + profit + 2000

		fmt.Println("----------------------------------")
	}

}

func ComputeProfitForTheYear(yearlyInvested int) int {

	profit := int(float64(yearlyInvested) * returnrate)
	return profit
}
