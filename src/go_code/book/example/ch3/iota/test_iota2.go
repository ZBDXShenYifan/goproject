package main

import "fmt"

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopBack
	FlagPointToPoint
	FlagMulticast
)

func main() {
	fmt.Println(FlagUp, FlagBroadcast, FlagLoopBack)
}

func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

func TurnDown(v *Flags) {
	*v &^=FlagUp
}

func SetBoardcast(v *Flags) {
	*v |= FlagBroadcast
}