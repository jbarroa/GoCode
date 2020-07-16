//Jasmine Barroa
//June 16, 2020
//Bank Application 

package main

import (
	"fmt"
	"strconv"
//	"strings"
//	"math"
)

type account struct {
	accountNumber string
	accountOwner entity
	balance float64
	interestRate float64
	accountType string
}

type entity struct{
	entityId string
	address string
	entityType string
}

type wallet struct{
	walletId string
	walletOwner entity
	accountArray []account
}

//Account struct methods
func (a *account) withdraw(value float64){
	if a.balance >= value{
		a.balance = a.balance - value
		fmt.Println("New account balance is", a.balance)
	}else{
		fmt.Println("There are insufficient funds in your account")
	}
}

func (a *account) deposit(value float64){
	if value > 0{
		a.balance += value
		fmt.Println("New account balance is", a.balance)
	}
}

func (a *account) applyInterest(){
	a.balance = a.balance * a.interestRate
	fmt.Println("New account balance is", a.balance)
}

//a account is the source account, b account is the target account
func (a *account) wire(value float64, b *account){
	if a.balance >= value && value > 0{
		a.balance = a.balance - value
		b.balance = b.balance + value
		fmt.Println("New account balance is", a.balance)
		fmt.Println("New account balance is", b.balance)
	}else{
		fmt.Println("There are insufficient funds in your account")
	}
}

//Entity struct method
func (e *entity) changeAddress(newAddress string){
	e.address = newAddress
	fmt.Println("New address is:", e.address)
}

//Wallet struct method
func (w *wallet) displayAccounts(){
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

func (w *wallet) balance() float64{
	var total float64
	for _, v := range w.accountArray {
		total += v.balance
	}
	return total
}

//func (w *wallet) wire(value float64, b* wallet){

//}

func main(){
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

   var input string
   var walletOwn wallet

   Message1:
		fmt.Println("1. Jasmine Barroa")
		fmt.Println("2. Matthew Koblias")
		fmt.Println("3. Joel Karie")
		fmt.Println("4. Centene")
		fmt.Print("Enter the number that corresponds with your name: ")
		fmt.Scan(&input)

   if input == "1"{
	   walletOwn = jasmineBarroaWallet
   }else if input == "2"{
	   walletOwn = matthewKobilasWallet
   }else if input == "3"{
	   walletOwn = joelKarieWallet
   }else if input == "4"{
	   walletOwn = centeneWallet
   }else{
	   fmt.Println("Please enter a valid wallet number")
	   goto Message1
   }

   fmt.Println(walletOwn)
   for{
		fmt.Println("Menu:")
		fmt.Println("1. View Account")		
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Transfer")
		fmt.Println("5. Total Balance")
		fmt.Println("6. Change Address")
		fmt.Println("7. View Wallet")
		fmt.Println("8. Exit")
		fmt.Print("Enter a number from menu: ")
		fmt.Scan(&input)
		choice, _ := strconv.Atoi(input)

		if choice == 1{
			fmt.Println("Wallet Id is:", walletOwn.walletId) 
		}else if choice == 2{
			fmt.Println("Choice:", choice)
		}else if choice == 3{
			fmt.Println("Choice:", choice)
		}else if choice == 4{
			fmt.Println("Choice:", choice)
		}else if choice == 5{
			fmt.Println("Choice:", choice)
		}else if choice == 6{
			fmt.Println("Choice:", choice)
		}else if choice == 7{
			fmt.Println("Wallet Id is:", walletOwn.walletId)
		}else if choice == 8{
			fmt.Println("Thank you for using the bank today.")
			break
		}else{
			fmt.Println("Enter a valid menu item")
		}
   }
   
   //jasmineBarroa.changeAddress("Tampa")

}