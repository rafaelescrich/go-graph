package main

/* The basic operations provided by a graph data structure G usually include:

adjacent(G, x, y): tests whether there is an edge from the vertex x to the vertex y;

neighbors(G, x): lists all vertices y such that there is an edge from the vertex x to the vertex y;

add_vertex(G, x): adds the vertex x, if it is not there;

remove_vertex(G, x): removes the vertex x, if it is there;

add_edge(G, x, y): adds the edge from the vertex x to the vertex y, if it is not there;

remove_edge(G, x, y): removes the edge from the vertex x to the vertex y, if it is there;

get_vertex_value(G, x): returns the value associated with the vertex x;

set_vertex_value(G, x, v): sets the value associated with the vertex x to v.

Structures that associate values to the edges usually also provide:

get_edge_value(G, x, y): returns the value associated with the edge (x, y);

set_edge_value(G, x, y, v): sets the value associated with the edge (x, y) to v. */

/*
Depth-first search

A depth-first search (DFS) is an algorithm for traversing a finite graph.
DFS visits the child vertices before visiting the sibling vertices; that is,
it traverses the depth of any particular path before exploring its breadth.
A stack (often the program's call stack via recursion) is generally used when
implementing the algorithm.

Breadth-first search

This section needs expansion. You can help by adding to it. (October 2012)
A breadth-first search (BFS) is another technique for traversing a finite graph.
BFS visits the sibling vertices before visiting the child vertices, and a queue
is used in the search process. This algorithm is often used to find the shortest
path from one vertex to another.

*/

// Graph is a representation of a Graph Data Structure
type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

// Node of a graph
type Node struct {
	value interface{}
}

type Edge struct {
	link map[Node][]*Node
}

// Adjacent tests whether there is an edge from the vertex x to the vertex y;
func (g *Graph) Adjacent(x Node, y Node) (Edge, error) {
	return nil, nil
}

//neighbors(G, x): lists all vertices y such that there is an edge from the vertex x to the vertex y;

//add_vertex(G, x): adds the vertex x, if it is not there;

//remove_vertex(G, x): removes the vertex x, if it is there;

//add_edge(G, x, y): adds the edge from the vertex x to the vertex y, if it is not there;

//remove_edge(G, x, y): removes the edge from the vertex x to the vertex y, if it is there;

//get_vertex_value(G, x): returns the value associated with the vertex x;

//set_vertex_value(G, x, v): sets the value associated with the vertex x to v.

func main() {

}
