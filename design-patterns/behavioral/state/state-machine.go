package state

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type State2 int

const (
	OffHook State2 = iota
	Connecting
	Connected
	OnHold
	OnHook
)

func (s State2) String() string {
	switch s {
	case OffHook:
		return "OffHook"
	case Connecting:
		return "Connecting"
	case Connected:
		return "Connected"
	case OnHold:
		return "OnHold"
	case OnHook:
		return "OnHook"
	}
	return "Unknown"
}

type Trigger int

const (
	CallDialed Trigger = iota
	HungUp
	CallConnected
	PlacedOnHold
	TakenOffHold
	LeftMessage
)

func (t Trigger) String() string {
	switch t {
	case CallDialed:
		return "CallDialed"
	case HungUp:
		return "HungUp"
	case CallConnected:
		return "CallConnected"
	case PlacedOnHold:
		return "PlacedOnHold"
	case TakenOffHold:
		return "TakenOffHold"
	case LeftMessage:
		return "LeftMessage"
	}
	return "Unknown"
}

type TriggerResult struct {
	Trigger Trigger
	State2  State2
}

var rules = map[State2][]TriggerResult{
	OffHook: {
		{CallDialed, Connecting},
	},
	Connecting: {
		{HungUp, OffHook},
		{CallConnected, Connected},
	},
	Connected: {
		{LeftMessage, OnHook},
		{HungUp, OnHook},
		{PlacedOnHold, OnHold},
	},
	OnHold: {
		{TakenOffHold, Connected},
		{HungUp, OnHook},
	},
}

func main2() {
	state2, exitState2 := OffHook, OnHook
	for ok := true; ok; ok = state2 != exitState2 {
		fmt.Println("The phone is currently", state2)
		fmt.Println("Select a trigger:")

		for i := 0; i < len(rules[state2]); i++ {
			tr := rules[state2][i]
			fmt.Println(strconv.Itoa(i), ".", tr.Trigger)
		}

		input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		i, _ := strconv.Atoi(string(input))

		tr := rules[state2][i]
		state2 = tr.State2
	}
	fmt.Println("We are done using the phone")
}
