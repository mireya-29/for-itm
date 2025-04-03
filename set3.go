// Placeholder
package main

import (
	"strings"
)

func main() {
	// Placeholder
}

// Relationship status
//
// Let's pretend that you are building a new app with social media functionality.
// Users can have relationships with other users.
//
// The two guidelines for describing relationships are:
// 1. Any user can follow any other user.
// 2. If two users follow each other, they are considered friends.
//
// This function describes the relationship that two users have with each other.
//
// Please see the sample data for examples of `socialGraph`.
//
// Params:
// - fromMember, the subject member
// - toMember, the object member
// - socialGraph, the relationship data
//
// Returns:
// - "follower" if fromMember follows toMember; "followed by" if fromMember is followed by toMember; "friends" if fromMember and toMember follow each other; "no relationship otherwise."

func relationshipStatus(fromMember string, toMember string, socialGraph map[string]map[string]string) string {
	contains := func(slice []string, item string) bool {
		for _, v := range slice {
			if v == item {
				return true
			}
		}
		return false
	}

	getFollowing := func(member string) []string {
		if memberCheck, exists := socialGraph[member]; exists {
			followingList := strings.Split(memberCheck["following"], ",")
			if followingList[0] == "" {
				return []string{}
			}
			return followingList
		}
		return []string{}
	}

	fromFollowing := getFollowing(fromMember)
	toFollowing := getFollowing(toMember)

	if contains(fromFollowing, toMember) {
		if contains(toFollowing, fromMember) {
			return "friends"
		}
		return "follower"
	}

	if contains(toFollowing, fromMember) {
		return "followed by"
	}
	
	return "no relationship"
}

// Tic tac toe
//
// Tic Tac Toe is a common paper-and-pencil game.
// Players must attempt to draw a line of their symbol across a grid.
// The player that does this first is considered the winner.
//
// This function evaluates a Tic Tac Toe game board and returns the winner.
//
// Please see the sample data for examples of `board`.
//
// Params:
// - board, the representation of the Tic Tac Toe board as a square slice of slices of strings. The size of the slice will range between 3x3 to 6x6. The board will never have more than 1 winner. There will only ever be 2 unique symbols at the same time.
//
// Returns:
// - the symbol of the winner, or "NO WINNER" if there is no winner.

func ticTacToe(board [][]string) string {
	size := len(board)

    check := func(line []string) string {
		if len(line) > 0 && line[0] != "" {
            win := true
            for _, cell := range line {
                if cell != line[0] {
                    win = false
                    break
                }
            }
            if win {
                return line[0]
            }
        }
        return ""
    }
  
    for _, row := range board {
        if winner := check(row); winner != "" {
            return winner
        }
    }
  
    for col := 0; col < size; col++ {
        column := make([]string, size)
        for row := 0; row < size; row++ {
            column[row] = board[row][col]
        }
        if winner := check(column); winner != "" {
            return winner
        }
    }

    diag1 := make([]string, size)
    diag2 := make([]string, size)
    for i := 0; i < size; i++ {
        diag1[i] = board[i][i]
        diag2[i] = board[i][size-1-i]
    }
    if winner := check(diag1); winner != "" {
        return winner
    }
    if winner := check(diag2); winner != "" {
        return winner
    }

    return "NO WINNER"
}

// ETA
//
// A shuttle van service is tasked to travel one way along a predefined circular route.
// The route is divided into several legs between stops.
// The route is fully connected to itself.
//
// This function returns how long it will take the shuttle to arrive at a stop after leaving anothe rstop.
//
// Please see the sample data for examples of `routeMap`.
//
// Params:
// - firstStop, the stop that the shuttle will leave
// - secondStop, the stop that the shuttle will arrive at
// - routeMap, the data describing the routes
//
// Returns:
// - the time that it will take the shuttle to travel from firstStop to secondStop

func eta(firstStop string, secondStop string, routeMap map[string]map[string]int) int {
	if firstStop == secondStop {
		return 0
	}
	
	visited := make(map[string]bool)
	time := 0
	current := firstStop
	
	for {
		visited[current] = true
		found := false
	
		for route, data := range routeMap {
			stops := strings.Split(route, ",")
			if stops[0] == current {
				time += data["travel_time_mins"]
				current = stops[1]
				found = true
				if current == secondStop {
					return time
				}
				break
				}
			}
	
			if !found || visited[current] {
				return -1
			}
		}
	}	
