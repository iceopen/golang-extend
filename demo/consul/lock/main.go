package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func main() {
	client, err := api.NewClient(&api.Config{Address: "127.0.0.1:8500"})
	if err != nil {
		fmt.Errorf(err.Error())
	}

	opts := &api.LockOptions{
		Key:        "webhook_receiver/1",
		Value:      []byte("set by sender 1"),
		SessionTTL: "5s",
		SessionOpts: &api.SessionEntry{
			Checks:   []string{"serfHealth"},
			Behavior: "release",
		},
	}

	lock, err := client.LockOpts(opts)
	stopCh := make(chan struct{})

	lockCh, err := lock.Lock(stopCh)

	if err != nil {
		panic(err)
	}
	fmt.Println("lockCh")
	<-lockCh
	err = lock.Unlock()
	if err != nil {
		fmt.Println("lock already unlocked")
	}
}
