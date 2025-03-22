package fakerinput

var FI *FakerInput
var Enabled bool

func InitGlobal() *FakerInput {
	Enabled = true
	FI = NewFakerInput()
	FI.Connect()
	return FI
}
