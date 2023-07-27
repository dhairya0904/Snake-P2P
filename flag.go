package main

import (
	"flag"
)

func parseFlags() *Config {
	c := &Config{}

	flag.StringVar(&c.RendezvousString, "rendezvous", "meetme", "Unique string to identify group of nodes. Share this with your friends to let them connect with you")
	flag.StringVar(&c.ListenHost, "host", "0.0.0.0", "The bootstrap node host listen address\n")
	flag.StringVar(&c.ProtocolID, "pid", "/chat/1.1.0", "Sets a protocol id for stream headers")
	flag.StringVar(&c.NodeType, "node", "master", "Sets node type")
	flag.IntVar(&c.ListenPort, "port", 4001, "node listen port")
	flag.StringVar(&c.peerAddress, "peer", "", "Sets peer address")
	flag.StringVar(&c.logLevel, "logLevel", "info", "Sets lob level for debugging")

	flag.Parse()
	return c
}
