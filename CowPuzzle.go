package main

import (
	"fmt"
	"strconv"
)

const (
	 NOSE_LEFT = 1
	 SMALL_HEAD = 2
	 BIG_HEAD = 3
	 LAY_TOP = 4
	 LAY_BOTTOM = LAY_TOP * -1
	 BIG_TAIL = BIG_HEAD * -1
	 SMALL_TAIL = SMALL_HEAD * -1
	 NOSE_RIGHT = NOSE_LEFT * -1
)

type Tile struct { 
	Top int
    Right int
    Bottom int
    Left int
    Id int 
}

var bag = [9]Tile {{NOSE_LEFT, SMALL_TAIL, LAY_BOTTOM, BIG_TAIL, 0},
        {LAY_BOTTOM, LAY_BOTTOM, BIG_HEAD, SMALL_HEAD, 1},
        {SMALL_HEAD, BIG_HEAD, LAY_BOTTOM, SMALL_HEAD, 2},
        {BIG_TAIL, NOSE_LEFT, LAY_BOTTOM, NOSE_RIGHT, 3},
        {LAY_TOP, BIG_TAIL, NOSE_RIGHT, BIG_TAIL, 4},
        {LAY_BOTTOM, SMALL_TAIL, NOSE_LEFT, BIG_TAIL, 5},
        {NOSE_RIGHT, BIG_HEAD, LAY_TOP, SMALL_HEAD, 6},
        {BIG_HEAD, SMALL_TAIL, NOSE_LEFT, NOSE_RIGHT, 7},
        {LAY_TOP, BIG_TAIL, NOSE_RIGHT, SMALL_HEAD, 8}}
		
func rotate(input Tile, rotation int) Tile {
	switch rotation {
    case 0: return input
    case 1: return Tile{input.Left, input.Top, input.Right, input.Bottom, input.Id}
    case 2: return Tile{input.Bottom, input.Left, input.Top, input.Right, input.Id}
    case 3: return Tile{input.Right, input.Bottom, input.Left, input.Top, input.Id}
    }
	//TODO: this should throw an error
	return input
}

func push(currentStack []Tile, item Tile) []Tile {
	return append(currentStack, item)
}

func pop(currentStack []Tile) (Tile, []Tile) {
	item := currentStack[len(currentStack) - 1]
	return item, currentStack[0:len(currentStack) - 1]
}

func stackContainsId(currentStack []Tile, id int) bool {
	
	for _, thisTile := range currentStack {
		if thisTile.Id == id {
			return true
		}
	}
	return false
}

func stackToString(currentStack []Tile) string {
	state := ""
	for _, thisTile := range currentStack {
		state = state + strconv.Itoa(thisTile.Id) + " "
	}
	
	return state
}

func findEdgeOnLeft(potentialTile Tile, edgeToFind int) []Tile {
	potentialTileRotations := []Tile{potentialTile, rotate(potentialTile, 1), rotate(potentialTile, 2), rotate(potentialTile, 3)}
	var matchingTileRotations []Tile
	for _, currentTile := range potentialTileRotations {
		if currentTile.Left == edgeToFind {
			matchingTileRotations = append(matchingTileRotations, currentTile)
		}
	}
	return matchingTileRotations
}

func findEdgeOnTop(potentialTile Tile, edgeToFind int) []Tile {
	potentialTileRotations := []Tile{potentialTile, rotate(potentialTile, 1), rotate(potentialTile, 2), rotate(potentialTile, 3)}
	var matchingTileRotations []Tile
	for _, currentTile := range potentialTileRotations {
		if currentTile.Top == edgeToFind {
			matchingTileRotations = append(matchingTileRotations, currentTile)
		}
	}
	return matchingTileRotations
}

func findEdgeOnLeftAndTop(potentialTile Tile, leftEdgeToFind int, topEdgeToFind int) []Tile {
	potentialTileRotations := []Tile{potentialTile, rotate(potentialTile, 1), rotate(potentialTile, 2), rotate(potentialTile, 3)}
	var matchingTileRotations []Tile
	for _, currentTile := range potentialTileRotations {
		if currentTile.Left == leftEdgeToFind && currentTile.Top == topEdgeToFind {
			matchingTileRotations = append(matchingTileRotations, currentTile)
		}
	}
	return matchingTileRotations
}

func recurse (input []Tile) string {

	switch len(input) {
	case 0:
		return start(input)
	case 9:
		return stackToString(input)
	case 3, 6:
		return recurseBottom(input)
	case 1, 2:
		return recurseRight(input)
	case 4, 5, 7, 8:
		return recurseRightAndBottom(input)
	}
	return ""
}

func start (input []Tile) string {
	var solutions []string
	var allSolutions string
	
	for _, initialTile := range bag {
		solutions = append(solutions, recurse([]Tile{initialTile}))
	}
	
	if len(solutions) > 0 {
		for _, thisSolution := range solutions {
			allSolutions = allSolutions + thisSolution + "\r\n"
		}
	}
	
	return allSolutions
}

func recurseRight(input []Tile) string {
	nextEdgeValue := input[len(input)-1].Right * -1
	unusedTiles := filterBag(input)
	
	var unusedTilesWithEdgeMatchOnLeft []Tile
	for _,thisTile := range unusedTiles {
		matchingEdgesOnLeft := findEdgeOnLeft(thisTile, nextEdgeValue)
		for _,thisTileAndRotation := range matchingEdgesOnLeft {
			unusedTilesWithEdgeMatchOnLeft = append(unusedTilesWithEdgeMatchOnLeft, thisTileAndRotation)
		}
	}
	
	var solutions []string
	var allSolutions string
	
	for _, thisTileAndRotation := range unusedTilesWithEdgeMatchOnLeft {
		solutions = append(solutions, recurse(push(input, thisTileAndRotation)))
	}
	
	if len(solutions) > 0 {
		for _, thisSolution := range solutions {
			allSolutions = allSolutions + thisSolution + "\r\n"
		}
	}
	
	return allSolutions
}

func recurseBottom(input []Tile) string {
	nextEdgeValue := input[len(input)-3].Bottom * -1
	unusedTiles := filterBag(input)
	
	var unusedTilesWithEdgeMatchOnTop []Tile
	for _,thisTile := range unusedTiles {
		matchingEdgesOnTop := findEdgeOnTop(thisTile, nextEdgeValue)
		for _,thisTileAndRotation := range matchingEdgesOnTop {
			unusedTilesWithEdgeMatchOnTop = append(unusedTilesWithEdgeMatchOnTop, thisTileAndRotation)
		}
	}
	
	var solutions []string
	var allSolutions string
	
	for _, thisTileAndRotation := range unusedTilesWithEdgeMatchOnTop {
		solutions = append(solutions, recurse(push(input, thisTileAndRotation)))
	}
	
	if len(solutions) > 0 {
		for _, thisSolution := range solutions {
			allSolutions = allSolutions + thisSolution + "\r\n"
		}
	}
	
	return allSolutions
}

func recurseRightAndBottom(input []Tile) string {
	nextTopEdgeValue := input[len(input)-3].Bottom * -1
	nextLeftEdgeValue := input[len(input)-1].Right * -1
	unusedTiles := filterBag(input)
	
	var unusedTilesWithEdgeMatches []Tile
	for _,thisTile := range unusedTiles {
		matchingEdgesOnTop := findEdgeOnLeftAndTop(thisTile, nextLeftEdgeValue, nextTopEdgeValue)
		for _,thisTileAndRotation := range matchingEdgesOnTop {
			unusedTilesWithEdgeMatches = append(unusedTilesWithEdgeMatches, thisTileAndRotation)
		}
	}
	
	var solutions []string
	var allSolutions string
	
	for _, thisTileAndRotation := range unusedTilesWithEdgeMatches {
		solutions = append(solutions, recurse(push(input, thisTileAndRotation)))
	}
	
	if len(solutions) > 0 {
		for _, thisSolution := range solutions {
			allSolutions = allSolutions + thisSolution + "\r\n"
		}
	}
	
	return allSolutions
}

func filterBag(input []Tile) []Tile {
	var availableTiles []Tile
	for _,thisBagTile := range bag {
		var tileFound = false
		for _,thisStackTile := range input {
			if thisBagTile.Id == thisStackTile.Id {
				tileFound = true
			}
		}
		if !tileFound {
			availableTiles = append(availableTiles, thisBagTile)
		}
	}
	
	return availableTiles
}

func main() {
	/*fmt.Println(bag[1])
	
	var s []Tile
	s1 := push(s, bag[1])
	
	fmt.Println(stackContainsId(s1, 2))
	fmt.Println(stackContainsId(s1, 1))
	fmt.Println(stackToString(s1))
	i, s2 := pop(s1)
	
	fmt.Println(i)
	fmt.Println(s2)
	fmt.Println(stackToString(s2))
	
	fmt.Println(findEdgeOnLeft(bag[1], -4))
	*/
	
	var usedTileStack []Tile
	allSolutions := recurse(usedTileStack)
	
	fmt.Println(allSolutions)
}