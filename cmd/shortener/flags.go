package main

import "flag"



var serverFlags struct{
	address string
	shortenBaseAddress string
}

func readFlags() {
	flag.StringVar(&serverFlags.address, "a", "localhost:8080", "address of server")
	flag.StringVar(&serverFlags.shortenBaseAddress, "b", "localhost:8080", "shorten address base")
}