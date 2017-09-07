package main

import (
	"log"
	"time"

	"github.com/ethereum/go-ethereum/whisper/whisperv2"
	fus "github.com/sonm-io/core/fusrodah"
	"github.com/sshaman1101/p2p-discovery/common"
)

func main() {
	bootNodes := common.LoadBootnodes()
	if len(bootNodes) == 0 {
		panic("No bootnodes")
	}

	serv, err := fus.NewServer(nil, common.SeekerPort, bootNodes)
	if err != nil {
		panic("Cannot start Frd server: " + err.Error())
	}

	serv.Start()
	log.Println("Start discovery")
	cc := common.NewCounter()

	serv.AddHandling(nil, nil, func(msg *whisperv2.Message) {
		from := msg.Recover()
		payload := string(msg.Payload)

		if from == nil {
			cc.AddAnon()
			log.Printf("[>] Anon   payload: %s\r\n", payload)
			log.Printf("[>] Anon   total:   %d\r\n", cc.GetAnon())
		} else {
			cc.AddSigned()
			log.Printf("[#] Signed payload: %s\r\n", payload)
			log.Printf("[#] Signed total:   %d\r\n", cc.GetSigned())
		}
	}, common.DiscoveryTopic)

	tk := time.NewTicker(10 * time.Second)
	defer tk.Stop()

	for {
		select {
		case <-tk.C:
			log.Printf("Heartbit")
		}
	}
}
