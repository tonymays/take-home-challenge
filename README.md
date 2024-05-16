Challenge: Find the shallowest duplicate value within a tree.

Ok.  So I pulled a Kobayashi Maru.  Today I felt like Captain Kirk.  I hope you guys are Ok with that.

The jest of my solution is to take in raw flat data and marshal it into a tree structure using
a pre-built go package.  The marshaler will not allow duplicates to be added so they are captured
and stored into a slice of struct that can later be examined for duplicate nodes with corresponding
nesting levels.

After reviewing two other PRs, in your repo, which both provided a hand coded solution to the challenge,
I choose not to provide a third.  So instead, I built out a small fstree (file system tree) package 
that could be included within any of your platforms out of the box.

Instead of hand coding a tree traverser, I used a tree package that takes in raw file system data 
and places into a tree structure using the following package:

https://pkg.go.dev/github.com/kingledion/go-tools@v0.6.0/tree

Today's go world relies heavily on well constructed go packages for the sake of reusability with a minimum 
set of code risk or technical debt.

So I choose to alter the test a little in order to achieve another solution possibility.

The tree package that I employed does not allow duplicates to be added which allowed me
to offload the duplicates into a data duplication slice that could be evaluated after the tree
had been built by the MarshalFSTree(...) method.  This simplified the number of operations
needed in order to capture the desired output: the most shallow offending duplicate node and the level
at which it lived.

Further, I did not need to provide a main.go file since I am providing a custom package containing the 
tools along with the desired solution.  Thus, the built in tests is all you will have to test out the solution.  

You can run these tests from command line by:
1 - changing to the `tree-search` directory
2 - run the `go test ./...  -coverpkg=./...` go command

For convenience, you can also run these test from VSCode (or other editors) by loading the
test file.

Further, if a traverse method were needed for other reasons the package that I provide here will allow 
you to choose between the BFS and DFS algorithms.  So instead of forcing the tree traverse 
with one or the other; as in the other PRs I examined, you can establish a choice which provides 
flexibility, efficiency and the removal of potential technical debt.

The solution is perceived to be n(0).  A reading of the documentation does not discuss this so I am
forced to assume.  This assumption would be benchmarked, in the real world, against other tree packages 
to ensure our development needs were met with best tools available at the time of development.

I can only hope you find my poor attempt to be Captain Kirk acceptable.

I appreciate the opportunity.

