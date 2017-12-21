package tree

import (
	"errors"
	"fmt"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type Mismatch struct{}

func (m Mismatch) Error() string {
	return "c"
}

// Build a tree
// We'll take a two-pass approach for efficiency. The first pass will put together a map
// of record-IDs to empty nodes, and the second will populate those nodes
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	// Sort the records by ID: We're going through them sequentially
	sort.SliceStable(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodeMap := make(map[int]*Node, len(records))

	// Build a map of empty nodes and verify records are continuous
	for i, r := range records {
		if i != r.ID {
			return nil, fmt.Errorf("Missing record %d", i)
		}

		nodeMap[r.ID] = &Node{r.ID, nil}
	}

	// Rearrange the nodes into a tree
	for _, r := range records {

		if r.ID == 0 {
			if r.Parent != 0 {
				return nil, errors.New("Root node must not have a parent")
			}

			continue
		}

		if r.ID == r.Parent {
			return nil, fmt.Errorf("Record %d cannot be its own parent", r.ID)
		}

		if r.ID < r.Parent {
			return nil, fmt.Errorf("Record %d must have a parent with a smaller ID", r.ID)
		}

		parentNode, ok := nodeMap[r.Parent]
		if !ok {
			return nil, errors.New("No record for parent")
		}

		parentNode.Children = append(parentNode.Children, nodeMap[r.ID])
	}

	// And then serve the root
	return nodeMap[0], nil
}
