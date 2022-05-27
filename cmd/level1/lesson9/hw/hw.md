#Homework 9


`
//Q: 1->2->3->4->5
//A: 5->4->3->2->1
type Node struct {
    Val int
    Next *Node
}

func reverse(head *Node) *Node {
}

//Q: 1<->2<->3<->4<->5
//A: 5<->4<->3<->2<->1

type Node struct {
    Val int
    Prev *Node
    Next *Node
}
`


