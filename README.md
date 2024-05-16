Challenge: Find the shallowest duplicate value within a tree.

I did notice the PRs on the other solutions and I took a moment to review them so as not
to waste your time with another hand coded tree traverser looking for a duplicate.

Today's go world relies heavily on well constructed go packages for the sake of 
complete reusability with a minimum set of code risk or technical debt.

So I am going to implement the solution using a go package:
https://pkg.go.dev/github.com/kingledion/go-tools@v0.6.0/tree

Further, I am not going to provide a main.go file which will require you to 
rely on the go test file and the tools provided by go to run individual tests.  If you
use VSCode or another editor that allows you to run tests, testing this solution would
be as simple as opening the test file and running the test.  I choose this method since
any tree solution integration with your platform would literally take this path.

The package I choose allows you to choose the traverse method whether BFS or DFS.  So instead
of forcing the tree traverse with one or the other, you can establish a choice which provides 
flexibility, efficiency and the removal of potential technical debt.

I did not notice a requirement to determine a solution result where two nodes were found duplicated 
at the same level of shallowness, so I am altering that requirement as a suggestion for 
consideration pending new requirements.

The solution is perceived to be n(0).  A reading of the documentation does not discuss this so I am
forced to assume.  The assumption would be benchmarked against other tree packages in the real world
to ensure our development needs would be meet from a usage model of evaluation.

I will leave the choice of solution to your opinion.

