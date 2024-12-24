package main

import "fmt"

const SIZE = 5

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Lenght int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

func (c *Cache) Check(str string) {
	node := &Node{}

	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{
			Val: str,
		}
	}

	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Add(n *Node) {
	fmt.Printf("Add: %s \n", n.Val)

	temp := c.Queue.Head.Right

	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = temp
	temp.Left = n

	c.Hash[n.Val] = n
	c.Queue.Lenght += 1

	if c.Queue.Lenght > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}
	c.Display()
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right

	fmt.Printf("%d = [", q.Lenght)

	for i := 0; i < q.Lenght; i++ {
		fmt.Printf("%s", node.Val)

		if i < q.Lenght-1 {
			fmt.Printf(" <--> ")
		}
		node = node.Right
	}

	fmt.Printf("]\n")

}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("Remove: %s \n", n.Val)
	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left

	c.Queue.Lenght -= 1
	delete(c.Hash, n.Val)
	return n
}

type Hash map[string]*Node

func main() {
	fmt.Printf("Start Cache!!\n")
	cache := NewCache()

	for _, word := range []string{"parrot", "banana", "apple", "mango", "patato", "dog", "cat"} {
		cache.Check(word)
	}
}
