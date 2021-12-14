package swagger

 import "testing"

func TestInterestMonthly(t *testing.T){
    var testLoanBody CalculateloanBody
    testLoanBody.LoanAmount = 350000 
	testLoanBody.LoanType = "interest" 
	testLoanBody.PaymentFrequency = "monthly" 
	testLoanBody.InterestRate = 0.03
	testLoanBody.LoanTerm = 1 

    got := calculate(testLoanBody)

    var want LoanRepayments
    want.MonthlyRepayments = 875
  
    if got.MonthlyRepayments != want.MonthlyRepayments {
        t.Errorf("got %q, wanted %q", got, want)
    } 

    want.TotalInterestPayable = 10500
    if got.TotalInterestPayable != want.TotalInterestPayable {
        t.Errorf("got %q, wanted %q", got, want)
    } 
    
    var loanRepaymentsAmountOwing LoanRepaymentsAmountOwing
    loanRepaymentsAmountOwing.Year = 0

	loanRepaymentsAmountOwing.Interest =10500

	loanRepaymentsAmountOwing.Principal =350000

	loanRepaymentsAmountOwing.Total =360500
    want.AmountOwing = append(want.AmountOwing, loanRepaymentsAmountOwing)
    
    if len(got.AmountOwing) != len(want.AmountOwing) {
        t.Errorf("got %q, wanted %q", got, want)
    } 

}
func TestInterestFort(t *testing.T){
    var testLoanBody CalculateloanBody
    testLoanBody.LoanAmount = 350000 
	testLoanBody.LoanType = "interest" 
	testLoanBody.PaymentFrequency = "fortnightly" 
	testLoanBody.InterestRate = 0.03
	testLoanBody.LoanTerm = 1 

    got := calculate(testLoanBody)

    var want LoanRepayments
    want.MonthlyRepayments = 404
  
    if got.MonthlyRepayments != want.MonthlyRepayments {
        t.Errorf("got %q, wanted %q", got, want)
    } 

    want.TotalInterestPayable = 10500
    if got.TotalInterestPayable != want.TotalInterestPayable {
        t.Errorf("got %q, wanted %q", got, want)
    } 
    
    var loanRepaymentsAmountOwing LoanRepaymentsAmountOwing
    loanRepaymentsAmountOwing.Year = 0

	loanRepaymentsAmountOwing.Interest =10500

	loanRepaymentsAmountOwing.Principal =350000

	loanRepaymentsAmountOwing.Total =360500
    want.AmountOwing = append(want.AmountOwing, loanRepaymentsAmountOwing)
    
    if len(got.AmountOwing) != len(want.AmountOwing) {
        t.Errorf("got %q, wanted %q", got, want)
    } 

}
func TestInterestweekly(t *testing.T){
    var testLoanBody CalculateloanBody
    testLoanBody.LoanAmount = 350000 
	testLoanBody.LoanType = "interest" 
	testLoanBody.PaymentFrequency = "weekly" 
	testLoanBody.InterestRate = 0.03
	testLoanBody.LoanTerm = 1 

    got := calculate(testLoanBody)

    var want LoanRepayments
    want.MonthlyRepayments = 202
  
    if got.MonthlyRepayments != want.MonthlyRepayments {
        t.Errorf("got %v, wanted %v", got, want)
    } 

    want.TotalInterestPayable = 10500
    if got.TotalInterestPayable != want.TotalInterestPayable {
        t.Errorf("got %v, wanted %v", got, want)
    } 
    
    var loanRepaymentsAmountOwing LoanRepaymentsAmountOwing
    loanRepaymentsAmountOwing.Year = 0

	loanRepaymentsAmountOwing.Interest =10500

	loanRepaymentsAmountOwing.Principal =350000

	loanRepaymentsAmountOwing.Total =360500
    want.AmountOwing = append(want.AmountOwing, loanRepaymentsAmountOwing)
    
    if len(got.AmountOwing) != len(want.AmountOwing) {
        t.Errorf("got %q, wanted %q", got, want)
    } 

}