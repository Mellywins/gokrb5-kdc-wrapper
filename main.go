package main

import (
	"fmt"
	"github.com/Mellywins/gokrb5-kdc-wrapper/kadmin"
)

func main() {
	atts := kadmin.CreateAddPrincipalAttributes().
		SetDupKey(0).
		SetForwardable(1).
		SetNeedChange(1).
		SetOkToAuthAsDelegate()
	addP := kadmin.AddPrincipal(*atts).
		WithPrincipal("Mellywins").
		WithExpDate("10 hours").
		WithPwExpDate("5 days").
		WithKvno(0).
		WithPassword("#!a2==!QsfK").
		WithClearPolicy().
		ParseCommand()
	kadmExecutor, err := kadmin.NewExecutorSpecBuilder().Local(true).MakeVerbose(true).Build()
	if err != nil {
		panic(err)
	}
	fmt.Println(kadmExecutor.Execute(addP))
}
