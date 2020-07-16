package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)
type account struct {
    accNum      string
    entity      entity
    balance     float64
    interest    float64
    accountType string
}
type entity struct {
    entityId   string
    address    string
    entityType string
}
type wallet struct {
    walletId     string
    walletOwner  entity
    accountArray []account
}
func main() {
    //create entitys
    jasmineBarroa := entity{"001", "Middle of the Screen", "Individual"}
    matthewKobilas := entity{"002", "Bottom of the Screen", "Individual"}
    joelKarie := entity{"003", "Top of the Screen", "Individual"}
    centene := entity{"004", "Detroit and Tampa", "Business"}
    //create three accounts for each entity
    jasmineBarroaCheck := account{"C001", jasmineBarroa, 100, 1.01, "checking"}
    matthewKobilasCheck := account{"C002", matthewKobilas, 100, 1.01, "checking"}
    joelKarieCheck := account{"C003", joelKarie, 100, 1.01, "checking"}
    centeneCheck := account{"C004", centene, 1000, 1.005, "checking"}
    jasmineBarroaSav := account{"S001", jasmineBarroa, 500, 1.05, "savings"}
    matthewKobilasSav := account{"S002", matthewKobilas, 500, 1.05, "savings"}
    joelKarieSav := account{"S003", joelKarie, 500, 1.05, "savings"}
    centeneSav := account{"S004", centene, 5000, 1.02, "savings"}
    jasmineBarroaInv := account{"I001", jasmineBarroa, 1000, 1.02, "investment"}
    matthewKobilasInv := account{"I002", matthewKobilas, 1000, 1.02, "investment"}
    joelKarieInv := account{"I003", joelKarie, 1000, 1.02, "investment"}
    centeneInv := account{"I004", centene, 10000, 1.01, "investment"}
    //fmt.Println(jasmineBarroaSav, matthewKobilasSav, joelKarieSav, centeneSav)
    //fmt.Println(jasmineBarroaCheck, matthewKobilasCheck, joelKarieCheck, centeneCheck)
    //fmt.Println(jasmineBarroaInv, matthewKobilasInv, joelKarieInv, centeneInv)
    jasmineBarroaAccArr := []account{jasmineBarroaCheck, jasmineBarroaSav, jasmineBarroaInv}
    matthewKobilasAccArr := []account{matthewKobilasCheck, matthewKobilasSav, matthewKobilasInv}
    joelKarieAccArr := []account{joelKarieCheck, joelKarieSav, joelKarieInv}
    centeneAccArr := []account{centeneCheck, centeneSav, centeneInv}
    jasmineBarroaWallet := wallet{"W001", jasmineBarroa, jasmineBarroaAccArr}
    matthewKobilasWallet := wallet{"W002", matthewKobilas, matthewKobilasAccArr}
    joelKarieWallet := wallet{"W003", joelKarie, joelKarieAccArr}
    centeneWallet := wallet{"W004", centene, centeneAccArr}
    fmt.Println(jasmineBarroaWallet, matthewKobilasWallet, joelKarieWallet, centeneWallet)
    fmt.Println("Matthew's account before withdrawal:", matthewKobilasCheck.balance)
    matthewKobilasCheck.withdraw(25)
    fmt.Println("Matthew's account after withdrawal:", matthewKobilasCheck.balance)
    fmt.Println("Jasmine's checking balance before $100 deposit:", jasmineBarroaCheck.balance)
    jasmineBarroaCheck.deposit(100)
    fmt.Println("Jasmine's checking balance after deposit:", jasmineBarroaCheck.balance)
    fmt.Println("It's time to apply interest to Joel's investment account. The balance before:", joelKarieInv.balance)
    joelKarieInv.applyInterest()
    fmt.Println("Joel's new investment balance:", joelKarieInv.balance)
    fmt.Println("Centene is wiring $500 from their checking to Jasmine's checking. \nCentene's balance prior:", centeneCheck.balance, "Jasmine's balance prior:", jasmineBarroaCheck.balance)
    centeneCheck.wire(&jasmineBarroaCheck, 500)
    fmt.Println("New balances:\nCentene:", centeneCheck.balance, "Jasmine:", jasmineBarroaCheck.balance)
    fmt.Println("Let's move Matthew to Detroit. Address before:", matthewKobilas.address)
    matthewKobilas.changeAddress("Detroit")
    fmt.Println("Mathew's new address:", matthewKobilas.address)
    fmt.Println("Let's display Joel's accounts:")
    joelKarieWallet.displayAccounts()
    fmt.Println("Let's look at Centene's balance:")
    fmt.Println(centeneWallet.balance())
    joelKarieWallet.userDeposit()
    joelKarieWallet.userWithdrawal()
    fmt.Println(joelKarieCheck.balance)
}
/*
func (a *account) withdraw(amount float64) {
    if amount <= a.balance && amount > 0 {
        a.balance -= amount
    }
}
*/
func (a *account) withdraw(value float64) {
    if a.balance >= value {
        a.balance = a.balance - value
        fmt.Println("New account balance is", a.balance)
    } else {
        fmt.Println("There are insufficient funds in your account")
    }
}
func (a *account) deposit(amount float64) {
    if amount > 0 {
        a.balance += amount
        fmt.Println("New account balance is", a.balance)
    }
}
func (a *account) applyInterest() {
    a.balance *= a.interest
}
func (a *account) wire(b *account, amount float64) {
    if amount <= a.balance && amount > 0 {
        a.balance -= amount
        b.balance += amount
    }
}
func (e *entity) changeAddress(newAddress string) {
    e.address = newAddress
}
func (w wallet) displayAccounts() {
    for _, v := range w.accountArray {
        if v.accountType == "checking" {
            fmt.Println(v)
        }
    }
    for _, v := range w.accountArray {
        if v.accountType == "investment" {
            fmt.Println(v)
        }
    }
    for _, v := range w.accountArray {
        if v.accountType == "savings" {
            fmt.Println(v)
        }
    }
}
func (w wallet) balance() float64 {
    var total float64
    for _, v := range w.accountArray {
        total += v.balance
    }
    return total
}
//deposit money
func (userWallet wallet) userDeposit() {
    in := bufio.NewReader(os.Stdin) //reader
    fmt.Println("Hello, User", userWallet.walletOwner.entityId, "\nWhich account would you like to deposit to (Checking, Savings, or Investment)?")
    var input string
    var accountType string
    var depositAmount int
    for {
        input, _ = in.ReadString('\n')
        nowLower := strings.ToLower(input)
        nowLower = strings.TrimRight(input, "\r\n")
        if nowLower == "checking" || nowLower == "savings" || nowLower == "investment" {
            accountType = nowLower
            break
        } else {
            fmt.Println("Please enter checking, savings, or investment.")
        }
    }
    var thisAccount account
    if accountType == "checking" {
        thisAccount = userWallet.accountArray[0]
    }
    if accountType == "savings" {
        thisAccount = userWallet.accountArray[1]
    }
    if accountType == "investment" {
        thisAccount = userWallet.accountArray[2]
    }
    //fmt.Println(thisAccount)
    fmt.Println("How much would you like to deposit?")
    for {
        input, _ = in.ReadString('\n')
        //nowLower := strings.ToLower(input)
        input = strings.TrimRight(input, "\r\n")
        //convert string to int
        stringAsInt, e := strconv.Atoi(input)
        if e == nil {
            depositAmount = stringAsInt
            break
        } else {
            fmt.Println("Please entere a number")
        }
    }
    //fmt.Println(accountType, depositAmount)
    thisAccount.deposit(float64(depositAmount))
    //fmt.Println(thisAccount)
}
//withdraw money
func (userWallet wallet) userWithdrawal() {
    in := bufio.NewReader(os.Stdin) //reader
    fmt.Println("Hello, User", userWallet.walletOwner.entityId, "\nWhich account would you like to withdraw from (Checking, Savings, or Investment)?")
    var input string
    var accountType string
    var depositAmount int
    for {
        input, _ = in.ReadString('\n')
        nowLower := strings.ToLower(input)
        nowLower = strings.TrimRight(input, "\r\n")
        if nowLower == "checking" || nowLower == "savings" || nowLower == "investment" {
            accountType = nowLower
            break
        } else {
            fmt.Println("Please enter checking, savings, or investment.")
        }
    }
    var thisAccount account
    if accountType == "checking" {
        thisAccount = userWallet.accountArray[0]
    }
    if accountType == "savings" {
        thisAccount = userWallet.accountArray[1]
    }
    if accountType == "investment" {
        thisAccount = userWallet.accountArray[2]
    }
    //fmt.Println(thisAccount)
    fmt.Println("How much would you like to withdraw?")
    for {
        input, _ = in.ReadString('\n')
        //nowLower := strings.ToLower(input)
        input = strings.TrimRight(input, "\r\n")
        //convert string to int
        stringAsInt, e := strconv.Atoi(input)
        if e == nil {
            depositAmount = stringAsInt
            break
        } else {
            fmt.Println("Please entere a number")
        }
    }
    fmt.Println(depositAmount)
    //fmt.Println(accountType, depositAmount)
    thisAccount.withdraw(float64(depositAmount))
    //fmt.Println(thisAccount)
}
