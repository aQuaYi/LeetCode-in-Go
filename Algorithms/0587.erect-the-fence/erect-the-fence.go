package Problem0587

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type Point = kit.Point

func outerTrees(points []Point) []Point {
    n := len(points)
    if n < 4 {
        return points
    }
    
    // Find the leftmost point as the first point in the result
    first := 0
    for i := 1; i < n; i++ {
        if points[i].X < points[first].X {
            first = i
        }
    }
    
    results := append([]Point{}, points[first])
    
    cur := first
    for {
        next := (cur + 1) % n
        for i := 0; i < n; i++ {
            if i == cur { continue }
            cross := crossProductLength(points[cur], points[i], points[next])  
            if cross < 0 || (cross == 0 && distance(points[cur], points[i]) > distance(points[cur], points[next])) {
                next = i
            }
        }
        
        for i := 0; i < n; i++ {
            if i == cur || i == first { continue }
            if crossProductLength(points[cur], points[i], points[next]) == 0 {
                results = append(results, points[i])
            }
        }
        
        cur = next
        
        if cur == first {
            break
        }
    }
    
    return results
}

func crossProductLength(a, b, c Point) int {
    return (a.X - b.X) * (c.Y - b.Y) - (a.Y - b.Y) * (c.X - b.X)
}

func distance(a, b Point) int {
    return (a.X - b.X) * (a.X - b.X) + (a.Y - b.Y) * (a.Y - b.Y)
}
