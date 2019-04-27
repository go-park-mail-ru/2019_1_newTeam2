package chatroulette

import (
	"fmt"
	"sync"

	"golang.org/x/net/context"
)

type ChatManager struct {
	mu       sync.RWMutex
}

func NewChatManager() *ChatManager {
	return &ChatManager{
		mu:       sync.RWMutex{},
	}
}

func (sm *ChatManager) Check(ctx context.Context, in *Username) (*Status, error) {
	fmt.Println("call Check", in)
	fmt.Println("eeeeeeeeeeeeeeee booooooooooooooooooooooooooooooooooooooooooooooooooooooy")
	return &Status{CheckStatus: true}, nil
}