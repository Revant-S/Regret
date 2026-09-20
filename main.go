package main

import "Regret/simulator/server"

const (
	SimulatorPort = ":8080"
)

func main() {
	// starting the simulator Server
	server.Server(SimulatorPort)
}
