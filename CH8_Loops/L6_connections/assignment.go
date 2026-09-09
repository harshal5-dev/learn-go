package main

func countConnections(groupSize int) int {
	totalConn := 0
	for conn := 0; conn < groupSize; conn++ {
		totalConn += conn
	}
	return totalConn
}
