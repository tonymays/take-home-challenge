package fstree

type FSData struct {
	ID       uint
	ParentID uint
	Name     string
	IsDir    bool
}

type FSDuplicateDataNode struct {
	ID         uint
	ParentID   uint
	Name       string
	IsDir      bool
	LevelFound int
}

type FSTree struct {
	ID    uint
	Name  string
	IsDir bool
	Level int
	Nodes []FSTree
}
