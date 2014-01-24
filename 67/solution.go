package main

import (
  "io/ioutil"
  "fmt"
  "strings"
  "strconv"
)

//////
// A* implementation
//
// Finding the most costly route
//
// Average values for triangle values: 50.121362
// Heuristic value per node: 51
//

// A point in the triangle
type Point struct {
  X, Y int
  // triangle[p.Y][p.X]
}

// Calculate an estimated cost for traveling down the triangle (must be larger 
// than the actual cost)
func heuristic_cost_estimate(point, goal Point) int64 {
  cost := int64(goal.Y - point.Y) * int64(100)
  return cost
}

// Reconstruct the path the A* has created
func reconstruct_path(came_from map[Point]Point, current_node Point) []Point {
  var origin Point
  var ok bool
  origin, ok = came_from[current_node]

  if ok {
    path := reconstruct_path(came_from, origin)
    path = append(path, current_node)
    return path
  } else {
    return []Point{current_node}
  }

  return []Point{}
}

// Find the neighbouring nodes, always down
func neighbour_nodes(current Point) []Point {
  neighbours := make([]Point, 2)
  neighbours[0] = Point{current.X, current.Y+1}
  neighbours[1] = Point{current.X+1, current.Y+1}

  return neighbours
}

// Find the index of a point in a slice of Points
func point_in_slice(p1 Point, slice []Point) int {
  for i, p2 := range slice {
    if p1 == p2 {
      return i
    }
  }

  return -1
}

func a_star(start, goal Point, triangle [][]int64) int64 {
  var closed_set []Point
  var open_set []Point
  open_set = append(open_set, start)

  came_from := make(map[Point]Point)

  g_score := make(map[Point]int64)
  f_score := make(map[Point]int64)

  g_score[start] = 0
  f_score[start] = g_score[start] + heuristic_cost_estimate(start, goal)

  for len(open_set) > 0 {
    var current Point
    current_index := -1
    // Find the node in open_set with highest f_score
    for i, open := range open_set {
      if current_index == -1 || f_score[open] > f_score[current] {
        current = open
        current_index = i
      }
    }
    // We've reached our goal
    if current == goal {
      var score int64
      path := reconstruct_path(came_from, goal)
      for _, point := range path {
        score += triangle[point.Y][point.X]
      }
      return score
    }

    // Remove current from open_set
    open_set = open_set[:current_index+copy(
      open_set[current_index:], open_set[current_index+1:])]
    // Add current to closed_set
    closed_set = append(closed_set, current)

    if current.Y >= len(triangle)-1 {
      continue
    }

    // for each neighbour node
    for _, neighbour := range neighbour_nodes(current) {
      // Already visited
      if point_in_slice(neighbour, closed_set) >= 0 {
        continue
      }

      tentative_g_score := g_score[current] + triangle[neighbour.Y][neighbour.X]

      // If not in open_set or tentative_g_score is greater than before
      if point_in_slice(neighbour, open_set) == -1 || tentative_g_score > g_score[neighbour] {
        came_from[neighbour] = current
        g_score[neighbour] = tentative_g_score
        f_score[neighbour] = tentative_g_score + heuristic_cost_estimate(neighbour, goal)

        if point_in_slice(neighbour, open_set) == -1 {
          open_set = append(open_set, neighbour)
        }
      }
    }
  }

  return 0
}


//////
// Reading and building triangle
//

func ReadFile(filename string) ([]string, error) {
  file, err := ioutil.ReadFile(filename)

  if err != nil {
    return nil, err
  }

  lines := strings.Split(string(file), "\n")

  return lines[:len(lines) - 1], nil
}

func MakeTriangle(lines []string) [][]int64 {
  var triangle [][]int64

  for i := 0; i < len(lines); i += 1 {
    // split line by spaces
    words := strings.Split(string(lines[i]), " ")
    // create triangle layer
    var triangle_layer []int64

    // go through each "word" in line
    for j := 0; j < len(words); j += 1 {
      // parse word as integer
      value, _ := strconv.ParseInt(words[j], 10, 0)
      // append to triangle layer
      triangle_layer = append(triangle_layer, value)
    }

    // append triangle layer to triangle
    triangle = append(triangle, triangle_layer)
  }

  return triangle
}

func main() {
  lines, err := ReadFile("triangle.txt")

  if err != nil {
    fmt.Println("reading failed")
    return
  }

  /*
  for _, line := range lines {
    fmt.Println(line)
  }
  /**/

  triangle := MakeTriangle(lines)

  /*
  for _, layer := range triangle {
    fmt.Println(layer)
  }
  /**/

  // Reverse A* through the triangle

  last_layer := len(triangle)-1
  var max_sum int64
  for i, _ := range triangle[last_layer] {
    sum := a_star(Point{0,0}, Point{i, last_layer}, triangle)
    if sum > max_sum {
      max_sum = sum
    }
  }

  fmt.Println(max_sum)
}
