package greetings

import (
	"errors"
	"fmt"
)

func YaaYaa(name string) (string, error) {
	if name == "" {
		return "", errors.New("名無しですよ")
	}
	message := fmt.Sprintf("やあやあ, %v. さん。こんにちは", name)
	return message, nil
}
