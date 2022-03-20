package main

import (
	kadmin "github.com/Mellywins/gokrb5-kdc-wrapper/kadmin"
)

func main() {
	// status, err := internal.EnsureServiceIsRunning()
	// fmt.Printf("Is service running: %v, status is : %s", err, status)

	atts := kadmin.CreateAddPrincipalAttributes().SetDupKey(0).SetForwardable(1).SetNeedChange(1).SetOkToAuthAsDelegate()
	kadmin.AddPrincipal(*atts).
		WithPrincipal("Mellywins").
		WithExpDate("10 hours").
		WithKvno(0).
		WithPassword("#!a2==!QsfK").
		WithClearPolicy().
		ParseCommand()
}
