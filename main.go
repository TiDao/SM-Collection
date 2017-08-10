package main

import (
	"fmt"
//	"SM/SM4"
//	"SM/SM3"
	"SM/SM2"
)

func main() {
/*
	fmt.Println("<SM4>")

	var key = []uint8{0x01,0x23,0x45,0x67,0x89,0xab,0xcd,0xef,0xfe,0xdc,0xba,0x98,0x76,0x54,0x32,0x10}
	var plain = []uint8{0x01,0x23,0x45,0x67,0x89,0xab,0xcd,0xef,0xfe,0xdc,0xba,0x98,0x76,0x54,0x32,0x10}

	var CypherText =  SM4.SM4_En( key, plain)
	var OriginText =  SM4.SM4_De( key, CypherText)

	fmt.Printf("key:: %x\n",key)
	fmt.Printf("plain:: %x\n",plain)
	fmt.Printf("CypherText:: %x\n",CypherText)
	fmt.Printf("OriginText:: %x\n",OriginText)

	fmt.Println("<SM4/>")


	fmt.Println("<SM3>")

	var Msg  = [] uint8 {0x61,0x62,0x63}

	var Hash =  SM3.SM3_To_256( Msg )

	fmt.Printf("Msg:: %x\n",Msg)
	fmt.Printf("Hash:: %x\n",Hash)

	fmt.Println("<SM3/>")

	fmt.Println("<SM2_EnDe>")

	var PriKey = []uint8{0x16,0x49,0xAB,0x77,0xA0,0x06,0x37,0xBD,0x5E,0x2E,0xFE,0x28,0x3F,0xBF,0x35,0x35,0x34,0xAA,0x7F,
				0x7C,0xB8,0x94,0x63,0xF2,0x08,0xDD,0xBC,0x29,0x20,0xBB,0x0D,0xA0}
	var Msg  = []uint8{0x65, 0x6E, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6F, 0x6E, 0x20, 0x73, 0x74, 0x61, 0x6E, 0x64, 0x61, 0x72, 0x64}

	CyperText, Err :=  SM2.SM2_En(PriKey, Msg)
	OriginText, Err :=  SM2.SM2_De(PriKey, CyperText)

	fmt.Printf("Msg:: %x\n",Msg)
	fmt.Printf("CyperText::%x\n",CyperText)
	fmt.Printf("OriginText::%x\n",OriginText)
	fmt.Printf("Err::%x\n",Err)

	fmt.Println("<SM2_EnDe/>")
*/
	fmt.Println("<SM2_Singature>")

	var PriKey = []uint8{0x39, 0x45, 0x20, 0x8f, 0x7b, 0x21, 0x44, 0xb1, 0x3f, 0x36, 0xe3, 0x8a, 0xc6, 0xd3, 0x9f, 0x95, 0x88, 0x93, 0x93, 0x69, 0x28, 0x60, 0xb5, 0x1a, 0x42, 0xfb, 0x81, 0xef, 0x4d, 0xf7, 0xc5, 0xb8}
	var IDA = []uint8{0x42, 0x4C, 0x49, 0x43, 0x45, 0x31, 0x32, 0x33, 0x40, 0x59, 0x41,0x48, 0x4F, 0x4F, 0x2E, 0x43, 0x4F,0x11} //ASCII code of userA's identification
	var Msg string = "message digest"

	R, S, PubKey, Err :=  SM2.SM2_Si(PriKey, IDA, []uint8(Msg))
	fmt.Printf("@Signature\n")
	fmt.Printf("PriKey:: %x\n",PriKey)
	fmt.Printf("IDA:: %x\n",IDA)
	fmt.Printf("Msg(String):: %v\n",Msg)
	fmt.Printf("R::%x\n",R)
	fmt.Printf("S::%x\n",S)
	fmt.Printf("Err::%x\n",Err)

	VerifyResult, Err :=  SM2.SM2_Ve(PubKey, IDA, []uint8(Msg), R, S)
	fmt.Printf("@Verify\n")
	fmt.Printf("VerifyResult::%v\n",VerifyResult)
	fmt.Printf("Err::%x\n",Err)

	fmt.Println("<SM2_Singature/>")
}
