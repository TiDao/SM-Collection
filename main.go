package main

import (
	"fmt"
	"SM/SM4"
)

func main() {
	fmt.Println("<SM4>")

	//SM.SM4_SelfCheck()

	var key = []uint8{0x01,0x23,0x45,0x67,0x89,0xab,0xcd,0xef,0xfe,0xdc,0xba,0x98,0x76,0x54,0x32,0x10}
	var plain = []uint8{0x01,0x23,0x45,0x67,0x89,0xab,0xcd,0xef,0xfe,0xdc,0xba,0x98,0x76,0x54,0x32,0x10}

	var CypherText =  SM4.SM4_En( key, plain)
	var OriginText =  SM4.SM4_De( key, CypherText)

	fmt.Printf("key:: %x\n",key)
	fmt.Printf("plain:: %x\n",plain)
	fmt.Printf("CypherText:: %x\n",CypherText)
	fmt.Printf("OriginText:: %x\n",OriginText)

	fmt.Println("<SM4/>")
}