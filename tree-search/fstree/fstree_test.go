package fstree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_MarshalFSTree tests marshaling raw file system data
func Test_MarshalFSTree(t *testing.T) {
	for name, test := range map[string]struct {
		data           []FSData
		expectedResult []FSTree
		expectedError  error
	}{
		"data items are empty": {
			data:           []FSData{},
			expectedResult: nil,
			expectedError:  SentinelError("missing file system data"),
		},
		"data items are nil": {
			data:           nil,
			expectedResult: nil,
			expectedError:  SentinelError("missing file system data"),
		},
		"data items with multiple parents which have no children": {
			data: []FSData{
				{
					ID:       1,
					ParentID: 1,
					Name:     "home",
					IsDir:    true,
				},
				{
					ID:       2,
					ParentID: 2,
					Name:     "usr",
					IsDir:    true,
				},
				{
					ID:       3,
					ParentID: 3,
					Name:     "tmp",
					IsDir:    true,
				},
			},
			expectedResult: []FSTree{
				{
					ID:    1,
					Name:  "home",
					IsDir: true,
					Level: 0,
					Nodes: []FSTree{},
				},
				{
					ID:    2,
					Name:  "usr",
					IsDir: true,
					Level: 0,
					Nodes: []FSTree{},
				},
				{
					ID:    3,
					Name:  "tmp",
					IsDir: true,
					Level: 0,
					Nodes: []FSTree{},
				},
			},
			expectedError: nil,
		},
		"data items with multiple parents which have multiple children": {
			data: []FSData{
				{
					ID:       1,
					ParentID: 1,
					Name:     "home",
					IsDir:    true,
				},
				{
					ID:       2,
					ParentID: 1,
					Name:     "anthony_mays",
					IsDir:    true,
				},
				{
					ID:       3,
					ParentID: 2,
					Name:     "personal",
					IsDir:    true,
				},
				{
					ID:       4,
					ParentID: 3,
					Name:     "resume.doc",
					IsDir:    false,
				},
				{
					ID:       5,
					ParentID: 5,
					Name:     "usr",
					IsDir:    true,
				},
				{
					ID:       6,
					ParentID: 5,
					Name:     "bin",
					IsDir:    true,
				},
			},
			expectedResult: []FSTree{
				{
					ID:    1,
					Name:  "home",
					IsDir: true,
					Level: 0,
					Nodes: []FSTree{
						{
							ID:    2,
							Name:  "anthony_mays",
							IsDir: true,
							Level: 1,
							Nodes: []FSTree{
								{
									ID:    3,
									Name:  "personal",
									IsDir: true,
									Level: 2,
									Nodes: []FSTree{
										{
											ID:    4,
											Name:  "resume.doc",
											IsDir: false,
											Level: 3,
											Nodes: []FSTree{},
										},
									},
								},
							},
						},
					},
				},
				{
					ID:    5,
					Name:  "usr",
					IsDir: true,
					Level: 0,
					Nodes: []FSTree{
						{
							ID:    6,
							Name:  "bin",
							IsDir: true,
							Level: 1,
							Nodes: []FSTree{},
						},
					},
				},
			},
			expectedError: nil,
		},
	} {
		t.Run(name, func(t *testing.T) {
			actualResult, duplicateData, actualError := MarshalFSTree(test.data)
			assert.ErrorIs(t, actualError, test.expectedError, "unexpected error")
			assert.Equal(t, actualResult, test.expectedResult, "unexpected result(s)")
			assert.Nil(t, duplicateData, "no duplicate data is expected from this marshaling test")
		})
	}
}

// Test_CheckDuplicateIDs tests file system data to identify duplicate fs data nodes
func Test_CheckDuplicateIDs(t *testing.T) {
	for name, test := range map[string]struct {
		data           []FSData
		expectedResult *FSTree
		expectedLevel  int
		expectedError  error
	}{
		"data items are empty": {
			data:           []FSData{},
			expectedResult: nil,
			expectedLevel:  0,
			expectedError:  SentinelError("missing file system data"),
		},
		"data items are nil": {
			data:           nil,
			expectedResult: nil,
			expectedLevel:  0,
			expectedError:  SentinelError("missing file system data"),
		},
		"data items with multiple parents which have no children and no duplicates": {
			data: []FSData{
				{
					ID:       1,
					ParentID: 1,
					Name:     "home",
					IsDir:    true,
				},
				{
					ID:       2,
					ParentID: 2,
					Name:     "usr",
					IsDir:    true,
				},
				{
					ID:       3,
					ParentID: 3,
					Name:     "tmp",
					IsDir:    true,
				},
			},
			expectedResult: nil,
			expectedLevel:  0,
			expectedError:  nil,
		},
		"data items with multiple parents which have no children and 1 duplicate": {
			data: []FSData{
				{
					ID:       1,
					ParentID: 1,
					Name:     "home",
					IsDir:    true,
				},
				{
					ID:       2,
					ParentID: 2,
					Name:     "usr",
					IsDir:    true,
				},
				{
					ID:       3,
					ParentID: 3,
					Name:     "tmp",
					IsDir:    true,
				},
				{
					ID:       3,
					ParentID: 3,
					Name:     "tmp",
					IsDir:    true,
				},
			},
			expectedResult: &FSTree{
				ID:    3,
				Name:  "tmp",
				IsDir: true,
				Level: 0,
				Nodes: []FSTree{},
			},
			expectedLevel: 0,
			expectedError: nil,
		},
		"data items with multiple parents which have multiple children and no duplicates": {
			data: []FSData{
				{
					ID:       1,
					ParentID: 1,
					Name:     "home",
					IsDir:    true,
				},
				{
					ID:       2,
					ParentID: 1,
					Name:     "anthony_mays",
					IsDir:    true,
				},
				{
					ID:       3,
					ParentID: 2,
					Name:     "personal",
					IsDir:    true,
				},
				{
					ID:       4,
					ParentID: 3,
					Name:     "resume.doc",
					IsDir:    false,
				},
				{
					ID:       5,
					ParentID: 5,
					Name:     "usr",
					IsDir:    true,
				},
				{
					ID:       6,
					ParentID: 5,
					Name:     "bin",
					IsDir:    true,
				},
			},
			expectedResult: nil,
			expectedLevel:  0,
			expectedError:  nil,
		},
		"data items with multiple parents which have multiple children and has 1 duplicate": {
			data: []FSData{
				{
					ID:       1,
					ParentID: 1,
					Name:     "home",
					IsDir:    true,
				},
				{
					ID:       2,
					ParentID: 1,
					Name:     "anthony_mays",
					IsDir:    true,
				},
				{
					ID:       3,
					ParentID: 2,
					Name:     "personal",
					IsDir:    true,
				},
				{
					ID:       4,
					ParentID: 3,
					Name:     "resume.doc",
					IsDir:    false,
				},
				{
					ID:       4,
					ParentID: 3,
					Name:     "resume.doc",
					IsDir:    false,
				},
				{
					ID:       5,
					ParentID: 5,
					Name:     "usr",
					IsDir:    true,
				},
				{
					ID:       6,
					ParentID: 5,
					Name:     "bin",
					IsDir:    true,
				},
			},
			expectedResult: &FSTree{
				ID:    4,
				Name:  "resume.doc",
				IsDir: false,
				Level: 3,
				Nodes: []FSTree{},
			},
			expectedLevel: 3,
			expectedError: nil,
		},
		"data items with multiple parents which have multiple children and has 2 duplicate": {
			data: []FSData{
				{
					ID:       1,
					ParentID: 1,
					Name:     "home",
					IsDir:    true,
				},
				{
					ID:       2,
					ParentID: 1,
					Name:     "anthony_mays",
					IsDir:    true,
				},
				{
					ID:       2,
					ParentID: 1,
					Name:     "anthony_mays",
					IsDir:    true,
				},
				{
					ID:       3,
					ParentID: 2,
					Name:     "personal",
					IsDir:    true,
				},
				{
					ID:       4,
					ParentID: 3,
					Name:     "resume.doc",
					IsDir:    false,
				},
				{
					ID:       4,
					ParentID: 3,
					Name:     "resume.doc",
					IsDir:    false,
				},
				{
					ID:       5,
					ParentID: 5,
					Name:     "usr",
					IsDir:    true,
				},
				{
					ID:       6,
					ParentID: 5,
					Name:     "bin",
					IsDir:    true,
				},
			},
			expectedResult: &FSTree{
				ID:    2,
				Name:  "anthony_mays",
				IsDir: true,
				Level: 1,
				Nodes: []FSTree{},
			},
			expectedLevel: 1,
			expectedError: nil,
		},
	} {
		t.Run(name, func(t *testing.T) {
			actualResult, actualLevel, actualError := CheckDuplicateIDs(test.data)
			assert.ErrorIs(t, actualError, test.expectedError, "unexpected error")
			assert.Equal(t, test.expectedResult, actualResult, "unexpected results failure")
			assert.Equal(t, test.expectedLevel, actualLevel, "unexpected level failure")
		})
	}
}
