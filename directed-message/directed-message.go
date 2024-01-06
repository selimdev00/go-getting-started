package directedmessage

import "fmt"

func DirectedMessage(name string, msg string) string {
	return fmt.Sprintf("Hey %v, here is message for u: %v", name, msg)
}