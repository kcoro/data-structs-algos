// ugraph is an implementation of a simple undirected graph in Go.

package ugraph

import "sync"

// Node is a single node with a data member of type string
type Node struct {
	data string
}

// Graph is a struct containing a slice of nodes, a map with key of Node and value of slice of Node ptrs,
// and a Read Write Mutex.
type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
	lock  sync.RWMutex
}

// addNode takes a *Node and adds it to the receiving *Graph.
// returns a boolean indicating success of insertion.
func (graph *Graph) addNode(node *Node) bool {
	graph.lock.Lock()

	if graph.nodes == nil {
		graph.lock.Unlock()
		return false
	}

	graph.nodes = append(graph.nodes, node)
	graph.lock.Unlock()
	return true
}

//addEdge takes a pair of *Node {A, B}, and inserts them into a *Graph
// returns a boolean indicating success of insertion.
func (graph *Graph) addEdge(nodeA, nodeB *Node) bool {
	graph.lock.Lock()

	if graph.edges == nil {
		graph.lock.Unlock()
		return false
	}

	graph.edges[*nodeA] = append(graph.edges[*nodeA], nodeB)
	graph.edges[*nodeB] = append(graph.edges[*nodeB], nodeA)
	graph.lock.Unlock()
	return true
}

// printNodesEdges will return a string visualizing each Node in the graph and
// the Nodes they each share an edge with.
func (graph *Graph) showNodesEdges() string {
	graph.lock.RLock()
	str := ""

	if graph.nodes == nil {
		graph.lock.RUnlock()
		return str
	}

	// Each node in the map contains a slice of Node* that share edges at that Node
	for i := range graph.nodes { //
		str += graph.nodes[i].data + " -> "
		if graph.edges[*graph.nodes[i]] == nil {
			continue
		}

		// for the given Node in the map, iterate over all the Node values
		connections := graph.edges[*graph.nodes[i]]
		for j := range graph.edges[*graph.nodes[i]] {
			str += connections[j].data + ", " // => A -> B, C, ...
		}

		str += "\n"
	}

	graph.lock.RUnlock()
	return str
}
