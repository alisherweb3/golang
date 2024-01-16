package main
import "fmt"

type BankAccount struct {
  holder string
  balance int
}

func (acc *BankAccount) withdraw(amount int) {
	if acc.balance < amount {
    fmt.Println("Insufficient Funds")
	} else {
    acc.balance -= amount
	}
}

func main() {
  acc := BankAccount{"James Smith", 100000}
  
  var amount int
  fmt.Scanln(&amount)
  
  acc.withdraw(amount)
  fmt.Println(acc)
}
