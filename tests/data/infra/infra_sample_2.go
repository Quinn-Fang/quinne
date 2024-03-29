package main

func main() {
	Platform := "aws"
	Authenticate()
	Ret := CreateEKSCluster("EKSCluster-1", 3)
	additionalNodeCount := lambda currentNode, expectedNode int: expectedNode-currentNode if expectedNode > currentNode else 5 if currentNode > 0  else -1 
	currentNodeCount := getCurrentNodeCount()
	expectedNodeCount := 6
	additionalNode := additionalNodeCount(currentNodeCount, expectedNodeCount)
	expandEKSCluster(additionalNode)
}
