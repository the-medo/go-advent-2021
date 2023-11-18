package day16

import (
	"fmt"
	"github.com/the-medo/go-advent-2021/utils"
)

type Packet struct {
	Version    int
	TypeId     int
	Value      int
	Subpackets []Packet
}

func Solve(input string) {
	binaryString, err := utils.HexStringToBinary(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	processBinaryPacket(binaryString)
}

func processBinaryPacket(binaryPacket string) *Packet {
	fmt.Println("Binary packet:", binaryPacket)

	binaryPacketVersion := binaryPacket[0:3]
	fmt.Println("Computing packet version from:", binaryPacketVersion)
	packetVersion := utils.BinaryToInt(binaryPacketVersion)
	fmt.Println("Packet version:", packetVersion)

	binaryPacketTypeId := binaryPacket[3:6]
	fmt.Println("Computing packet type id:", binaryPacketTypeId)
	packetTypeId := utils.BinaryToInt(binaryPacketTypeId)
	fmt.Println("Packet type ID:", packetTypeId)

	//Literal value!
	if packetTypeId == 4 {
		loop := true
		startingPoint := 6
		value := ""
		for loop {
			infoPart := binaryPacket[startingPoint]
			valuePart := binaryPacket[startingPoint+1 : startingPoint+5]
			startingPoint += 5
			value = value + valuePart
			if infoPart == '0' {
				loop = false
			}
		}

		fmt.Println("Literal binary value:", value)
		fmt.Println("Literal value:", utils.BinaryToInt(value))
	}

	return &Packet{
		Version: packetVersion,
		TypeId:  packetTypeId,
	}
}
