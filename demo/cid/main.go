package main

import (
	"fmt"
	"github.com/ipfs/go-cid"
	mh "github.com/multiformats/go-multihash"
)

func main() {
	// Create a cid manually by specifying the 'prefix' parameters
	pref := cid.Prefix{
		Version:  1,
		Codec:    cid.Raw,
		MhType:   mh.SHA2_256,
		MhLength: -1, // default length
	}

	// And then feed it some data
	c, err := pref.Sum([]byte("172.16.50.132"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Created CID: ", c)
}
