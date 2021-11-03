type path struct {
    location int
    steps int
}

func treeDiameter(n int, tree [][]int) int {

    graph := map[int][]int{}
    
    for _, connection := range tree {
        a := connection[0]
        b := connection[1]
        graph[a] = append(graph[a], b)
        graph[b] = append(graph[b], a)
    }
    
    q := []path{{0, 0}}
    seen := map[int]bool{}
    var at path
    
    // find the farthest node from 0
    for len(q) > 0 {

        at = q[0]
        q = q[1:]
        
        seen[at.location] = true
        
        for _, dst := range graph[at.location] {

            if seen[dst]{
                continue
            }

            nextLocation := dst
            nextSteps := at.steps + 1
            q = append(q, path{nextLocation, nextSteps})
        }
    }

    q = []path{{at.location, 0}}
    seen = map[int]bool{}

    // find the farthest node from where we were
    for len(q) > 0 {

        at = q[0]
        q = q[1:]
        
        if seen[at.location]{
            continue
        }
        seen[at.location] = true
        
        for _, dst := range graph[at.location] {

            if seen[dst]{
                continue
            }

            nextLocation := dst
            nextSteps := at.steps + 1
            q = append(q, path{nextLocation, nextSteps})
        }
    }
    
    return at.steps
}
