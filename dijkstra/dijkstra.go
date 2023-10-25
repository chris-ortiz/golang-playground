package main

import (
	"fmt"
)

type Origin struct {
	weight uint16
	from   *Vertex
}

type Vertex struct {
	name      string
	adjacents []WeightedEdge
}

type WeightedEdge struct {
	vertex *Vertex
	weight uint16
}

func main() {
	hannover := &Vertex{"Hannover", []WeightedEdge{}}
	dortmund := &Vertex{"Dortmund", []WeightedEdge{
		{hannover, 50},
	}}

	ohara := &Vertex{"ohara", []WeightedEdge{
		{hannover, 10},
	}}
	neustrelitz := &Vertex{"Neustrelitz", []WeightedEdge{
		{ohara, 10},
	}}
	siegmaringen := &Vertex{"Siegmaringen", []WeightedEdge{
		{neustrelitz, 10},
	}}

	stuttgart := &Vertex{"Stuttgart", []WeightedEdge{
		{hannover, 600},
		{dortmund, 400},
		{siegmaringen, 10},
	}}

	distances := map[*Vertex]Origin{
		stuttgart: {0, stuttgart},
	}

	visit(stuttgart, distances)

	for k, v := range distances {
		fmt.Printf("%s: %d over %s\n", k.name, v.weight, v.from.name)
	}

}

func visit(v *Vertex, distances map[*Vertex]Origin) {
	for _, edge := range v.adjacents {
		currentShortestDistance, ok := distances[edge.vertex]

		startingDistance := distances[v].weight
		distance := edge.weight + startingDistance

		if !ok || distance < currentShortestDistance.weight {
			distances[edge.vertex] = Origin{distance, v}
		}

		visit(edge.vertex, distances)
	}
}
