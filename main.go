package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)
	for i, j := 0, len(sd)-1; i < j; i, j = i+1, j-1 {
		sd[i], sd[j] = sd[j], sd[i]
	}
	fmt.Println("🍔🍔 " + string(sd) + " 🍔🍔")

}
