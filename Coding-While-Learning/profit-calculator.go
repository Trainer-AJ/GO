/*
Variables inside Func SCOPED to that function only
*/

package main
import (
	"fmt"
	"math"
)

var a = 2.0
var b = 3.0
const Year = 2024 
/*
ERROR ==>
.\profit-calculator.go:12:1: syntax error: non-declaration statement outside function body 
*/
// globalVar := "I am a global variable"

var globalVar = "I am a global variable"

func main(){
	// calculateRevenue()
	returnTest(a,b)
	name, age := getInput()
	userName(name, age)
	calculateRevenue()
}


/* IF YOU RETURN BUT DON'T DEFINE RETURN TYPE BELOW ERROR WILL OCCUR: 
# command-line-arguments
.\profit-calculator.go:9:15: getInput() (no value) used as value
.\profit-calculator.go:20:9: too many return values
        have (string, string)
        want ()
*/

// in the other ( ) ==> you specify what value type you are returning
func getInput() (string, string) {
	var name, age string
	fmt.Println(globalVar)
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	return name, age
}
	/*
	fmt.Print(globalVar)
}
# command-line-arguments
.\profit-calculator.go:12:1: syntax error: non-declaration statement outside function body

The syntax error you are encountering is due to a statement (fmt.Print(globalVar)) placed after a return statement within the getInput function. In Go, once a return statement is executed, the function exits immediately, and any code following the return statement is unreachable.
*/
/*
# command-line-arguments
.\profit-calculator.go:10:17: cannot use age (variable of type string) as int value in argument to userName
	*/
func userName(name,age string){
	fmt.Println("Hello", name,"You are", age, "years old")
}

func calculateRevenue() {
	// ask revenue, expense , tax rates
	var revenue, expenses, taxRate float64
	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	// calculate profit 
	// var profit float64 ===> unneccessary
	profit := revenue - expenses


	// earning after tax
// Wrong LOGIC ==>
	// earningAfterTax := profit * (1 - taxRate)
	earningAfterTax := profit * (1 - taxRate/100)

	// calculate ratio earning after tax
	ratio := int((earningAfterTax / profit ) * 100)

	fmt.Println("Your profit is", profit, "and your earning after tax is", earningAfterTax, "with a net profit ratio of", ratio, "%", "in year", Year)
}

func returnTest(a,b float64) (powerAns float64){ 
	powerAns = math.Pow(a, b)
	fmt.Printf("%.1f", powerAns)
	return powerAns
}