// Code generated from Pkl module `o47monad.sercon.LogConfig`. DO NOT EDIT.
package level

import (
	"encoding"
	"fmt"
)

type Level string

const (
	Fatal Level = "fatal"
	Error Level = "error"
	Warn  Level = "warn"
	Info  Level = "info"
	Debug Level = "debug"
)

// String returns the string representation of Level
func (rcv Level) String() string {
	return string(rcv)
}

var _ encoding.BinaryUnmarshaler = new(Level)

// UnmarshalBinary implements encoding.BinaryUnmarshaler for Level.
func (rcv *Level) UnmarshalBinary(data []byte) error {
	switch str := string(data); str {
	case "fatal":
		*rcv = Fatal
	case "error":
		*rcv = Error
	case "warn":
		*rcv = Warn
	case "info":
		*rcv = Info
	case "debug":
		*rcv = Debug
	default:
		return fmt.Errorf(`illegal: "%s" is not a valid Level`, str)
	}
	return nil
}
