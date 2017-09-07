package main

import (
	"flag"
	"log"
	"time"

	"github.com/satori/uuid"
	fus "github.com/sonm-io/core/fusrodah"
	"github.com/sshaman1101/p2p-discovery/common"
)

var nodeID string

func init() {
	flag.StringVar(&nodeID, "id", "", "uniq node id for discovery")
	flag.Parse()
}

func main() {
	bootNodes := common.LoadBootnodes()
	if len(bootNodes) == 0 {
		panic("No bootnodes")
	}

	serv, err := fus.NewServer(nil, common.CasterPort, bootNodes)
	if err != nil {
		panic("Cannot start Frd server: " + err.Error())
	}

	if nodeID == "" {
		nodeID = uuid.NewV4().String()
	}

	serv.Start()
	log.Printf("Start broadcasting from nodeID = %s\r\n", nodeID)

	anTk := time.NewTicker(10 * time.Second)
	defer anTk.Stop()

	sigTk := time.NewTicker(10 * time.Second)
	defer sigTk.Stop()

	for {
		select {
		case <-anTk.C:
			log.Printf("Send anon topic")
			serv.Send(nodeID, true, common.DiscoveryTopic)

		case <-sigTk.C:
			log.Printf("Send signed topic")
			serv.Send(nodeID, false, common.DiscoveryTopic)
		}
	}
}
