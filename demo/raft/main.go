package main

func main() {
	a := NewNode("A")
	b := NewNode("B")
	c := NewNode("C")

	write := make(chan map[string]any, 5)
	r1 := make(chan map[string]any, 1)
	r2 := make(chan map[string]any, 1)
	r3 := make(chan map[string]any, 1)
	go func() {
		for {
			data := <-write
			r1 <- data
			r2 <- data
			r3 <- data
		}
	}()

	go a.Start(write, r1)
	go b.Start(write, r2)
	go c.Start(write, r3)
	select {}
}

const (
	StateFollower = iota + 1
	StateCandidate
	StateLeader
)

type Node struct {
	Name  string
	Num   int
	State int
}

func NewNode(name string) *Node {
	return &Node{
		Name:  name,
		Num:   1,
		State: StateFollower,
	}
}

func (n *Node) Start(write chan<- map[string]any, reade <-chan map[string]any) {
	for {
		select {
		case v := <-reade:
			_, ok := v["heartbeat"]
			if ok {
				continue
			}
		default:
			if n.State == StateLeader {
				write <- map[string]any{"name": n.Name, "num": n.Num}
			}
		}
	}
}
