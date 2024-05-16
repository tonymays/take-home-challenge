package fstree

import (
	"errors"

	"github.com/kingledion/go-tools/tree"
)

// CheckDuplicateIDs examines file system data for duplicates and if any duplicates are
// found return a *FSTree node identifying the offending node along with the nested level
// within the overall tree
func CheckDuplicateIDs(fsData []FSData) (*FSTree, int, error) {
	if len(fsData) == 0 {
		return nil, 0, SentinelError("missing file system data")
	}

	// marshal to data we were given into a tree
	_, duplicates, err := MarshalFSTree(fsData)
	if err != nil {
		return nil, 0, err
	}

	// check for duplicate nodes
	shallowestLevel := 100000000000
	foundAt := -1
	for i, duplicate := range duplicates {
		if duplicate.LevelFound < shallowestLevel {
			shallowestLevel = duplicate.LevelFound
			foundAt = i
		}
	}

	// return what found if we found anything
	if foundAt > -1 {
		fsTree := &FSTree{
			ID:    duplicates[foundAt].ID,
			Name:  duplicates[foundAt].Name,
			IsDir: duplicates[foundAt].IsDir,
			Level: duplicates[foundAt].LevelFound,
			Nodes: []FSTree{},
		}
		return fsTree, shallowestLevel, nil
	}

	// let the world know we found nothing
	return nil, 0, nil
}

// MarshalFSTree marshals file system raw flat data into an FSTree node on complete success
// Note that this marshaler will NOT ALLOW DUPLICATES to be added and any are found they are
// store in a slice of FSDuplicateDataNode for later examination
func MarshalFSTree(fsData []FSData) ([]FSTree, []FSDuplicateDataNode, error) {
	if len(fsData) == 0 {
		return nil, nil, SentinelError("missing file system data")
	}

	// establish a FSDuplicateData slice
	var duplicateData []FSDuplicateDataNode

	// build out the tree
	checkTree := tree.Empty[*FSData]()
	for _, data := range fsData {
		item := data
		// the marshaler needs to ensure that that when the id and parentID are the same that the parentID is represented as 0
		// in order for the tree to build the root parent nodes correctly
		if item.ID == item.ParentID {
			item.ParentID = 0
		}

		// set the the id and parent id
		nodeID := uint(item.ID)
		parentID := uint(item.ParentID)

		// find the parent
		_, found := checkTree.Find(parentID)

		// if we do not have a root parent...
		if !found {
			// then create one
			added, exists := checkTree.Add(parentID, 0, nil)

			// even this error really make no sense at this stage ... it is prudent to check it as a just in case
			if exists {
				return nil, nil, errors.New("could not add parent node to tree, already exists somehow")
			} else if !added {
				// fail if we cannot add the parent
				return nil, nil, errors.New("could not add parent node to tree")
			}
		}

		// now lets start processing each data node
		if _, found := checkTree.Find(nodeID); found {
			// this means we have found a full duplicate node therefore ignore it here but append to the FSDuplicateData slice
			duplicateData = append(duplicateData, FSDuplicateDataNode{
				ID:         item.ID,
				ParentID:   item.ParentID,
				Name:       item.Name,
				IsDir:      item.IsDir,
				LevelFound: 0,
			})
		} else {
			added, exists := checkTree.Add(nodeID, uint(item.ParentID), &item)
			if exists {
				return nil, nil, errors.New("could not add node to tree, already exists")
			} else if !added {
				return nil, nil, errors.New("could not add node to tree")
			}
		}
	}

	// now lets build the tree that we have serialized the data
	var fsTree []FSTree
	for _, rootChild := range checkTree.Root().GetChildren() {
		rootChildData := rootChild.GetData()
		level := 0
		fsTree = append(fsTree, FSTree{
			ID:    rootChildData.ID,
			Name:  rootChildData.Name,
			IsDir: rootChildData.IsDir,
			Level: level,
			Nodes: constructChildFSTree(rootChild, duplicateData, level),
		})
	}

	return fsTree, duplicateData, nil
}

// constructChildFSTree construct child FSData into an FSTree
func constructChildFSTree(parent tree.Node[*FSData], duplicateData []FSDuplicateDataNode, level int) []FSTree {
	// set the tree and iteration level
	fsTree := []FSTree{}
	level++

	// walk any parent children found
	for _, child := range parent.GetChildren() {
		node := child.GetData()

		// set the level accordingly if we have duplicates
		for i := range duplicateData {
			if duplicateData[i].ID == node.ID {
				duplicateData[i].LevelFound = level
			}
		}

		// append then children taking note that this is a very recursive process
		fsTree = append(fsTree, FSTree{
			ID:    node.ID,
			Name:  node.Name,
			IsDir: node.IsDir,
			Level: level,
			Nodes: constructChildFSTree(child, duplicateData, level),
		})
	}

	return fsTree
}
