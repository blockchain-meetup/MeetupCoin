package main

import (
  "fmt"
  "time"
  "github.com/hashgraph/hedera-sdk-go"
)

func main() {
  client := hedera.Dial("testnet.hedera.com:50126")
  defer client.Close()

  myAccount := hedera.NewAccountID(0, 0, 1005)

  balance, err := client.GetAccountBalance(myAccount).Answer()
  if err != nil {
    panic(err)
  }
  fmt.Printf("Your balance: %v \n", balance)
}
