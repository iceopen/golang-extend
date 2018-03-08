package main

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/consul/api"
)

func main() {
	client, err := api.NewClient(&api.Config{Address: "192.168.1.8:8500"})
	if err != nil {
		log.Fatal(err.Error())
	}

	opts := &api.LockOptions{
		Key:        "webhook_receiver/1",
		Value:      []byte("set by sender 1"),
		SessionTTL: "10s",
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

	cancelCtx, cancelRequest := context.WithCancel(context.Background())
	req, _ := http.NewRequest("GET", "https://www.baidu.com", nil)
	req = req.WithContext(cancelCtx)
	go func() {
		http.DefaultClient.Do(req)
		select {
		case <-cancelCtx.Done():
			log.Println("request cancelled")
		default:
			log.Println("request done")

			err = lock.Unlock()
			if err != nil {
				log.Println("lock already unlocked")
			}
		}
	}()
	go func() {
		<-lockCh
		cancelRequest()
	}()

}
