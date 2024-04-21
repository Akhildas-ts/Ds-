package main

import (
	"fmt"
)

// the hole vertex store int the vertces that is graph
type Graph struct {
	vertces []*vertex
}

// each vertex we have value and where the current vertex is connected to where ? , that also there in the vertex ..
type vertex struct {
	val      int
	adjacent []*vertex
}

// where if the value is not  there then we insert into the graph ..

func (g *Graph) InsertVertex(val int) {

	if !contains(g.vertces, val) {

		g.vertces = append(g.vertces, &vertex{val: val})

	}

}

// add edges to each to vertex

// make connect to each vertex to another vertex (edges )

func (g *Graph) AddEdges(from, to int) {

	//getting from wher  to where we need to make connection
	fromVertex := GetVertex(g.vertces, from)
	toVertex := GetVertex(g.vertces, to)

	if fromVertex == nil || toVertex == nil {

		fmt.Println(" there is no more vertex can't make edges")

		return
	}

	for _, val := range fromVertex.adjacent {

		if val.val == to {

			fmt.Println("duplicate spoted")
			return
		}
	}

	// if the value is there then make connection each other..

	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	toVertex.adjacent = append(toVertex.adjacent, fromVertex)

}

// GetVertex if there
func GetVertex(vertces []*vertex, value int) *vertex {

	for _, val := range vertces {

		if val.val == value {
			return val
		}
	}

	return nil
}

// check if vertex is there..
func contains(vertex []*vertex, value int) bool {

	for _, val := range vertex {

		if val.val == value {
			return true
		}
	}

	return false

}

// display the graph

func (g *Graph) Display() {

	// check each vertex , how many connection have there..
	for _, val := range g.vertces {

		fmt.Print("the vertex ", val.val, ": ")
		for _, nighbor := range val.adjacent {
			fmt.Print(" ", nighbor.val, " ")

		}

		fmt.Println()

	}
}

// BFS - BREAST FRIST SEARCH

// its used queue DataStructure

type queue struct {
	arr []*vertex
}

// bst what we are doing ..
// frist of all we got the value, we need to start with the vertex
// check the vertex is there
// if there then add the element as  seen [MAP]
//
//	we need add into queue
//
// then start the processes :
// store the value
// remove the vertex from the queue
// because verfication almost over
// then check the neibours of the vertex and print .
func (g *Graph) bst(key int) {

	q := &queue{}

	seen := make(map[int]bool)
	start := GetVertex(g.vertces, key)

	if start == nil {
		fmt.Println("there is no more vertics is there for start Bst :)")
		return
	}

	seen[start.val] = true

	q.arr = append(q.arr, start)

	for len(q.arr) != 0 {

		vertex := q.arr[0]

		q.arr = q.arr[1:]
		fmt.Print("the vertex ", vertex.val, " :")

		for _, neigbours := range vertex.adjacent {

			if !seen[neigbours.val] {
				fmt.Print(neigbours.val)
				fmt.Println()

				seen[neigbours.val] = true
				q.arr = append(q.arr, neigbours)
			}

		}

	}
	fmt.Println()

}

// DFS  DEAPTH FRIST SEARCH
type stack struct {
	arr []*vertex
}

func (g *Graph) dfs(key int) {

	s := &stack{}

	seen := make(map[int]bool)
	start := GetVertex(g.vertces, key)

	if start == nil {
		fmt.Println("there is no more elments are there ")
		return
	}

	seen[start.val] = true
	s.arr = append(s.arr, start)

	for len(s.arr) > 0 {

		vertex := s.arr[len(s.arr)-1]
		s.arr = s.arr[:len(s.arr)-1]
         fmt.Println()
		fmt.Print("the vertex :",vertex.val, ":")

		for _, neibours := range vertex.adjacent {
			if !seen[neibours.val] {

				seen[neibours.val] = true

				fmt.Print(" ",neibours.val)

				s.arr = append(s.arr, neibours)

			}
		}
	}

}

func main() {

	g := &Graph{}

	g.InsertVertex(10)
	g.InsertVertex(20)
	g.InsertVertex(20)
	g.InsertVertex(30)
	g.InsertVertex(40)
	g.InsertVertex(50)
	g.InsertVertex(60)
	g.AddEdges(30, 20)
	g.AddEdges(20, 10)
	g.AddEdges(10, 50)
	g.AddEdges(20, 60)
	g.dfs(30)
	// g.Display()

}
