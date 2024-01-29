# Daily Coding Problem: Problem #1608 [Medium]

## Problem Statement

This problem was asked by Microsoft.

Write a program to determine how many distinct ways
there are to create a max heap from a list of `N` given integers.

For example,
if `N = 3`,
and our integers are `[1, 2, 3]`,
there are two ways, shown below.

```
  3      3
 / \    / \
1   2  2   1
```
## Analysis

I would not get the Microsoft job.
It took me quite a lot of thinking and drawing heaps as binary trees
to figure out a way to do this, and it's probably not the best way.

### Code

I wrote a quick and dirty [heap implementation](heap) to understand things, 
I cribbed most of the code from a "heap sort" [linked list](https://github.com/bediger4000/linked_lists)
sorting algorithm, which was in turn transliterated from a C heap sort
algorithm I wrote after receiving enlightenment about heaps.
The point of this is to try inserting integers into a genuine heap
to see if particular configurations can be acheived.

Programs that use parts of `package heap`, or that I used to understand this problem:

* [example heap sort](cmd/heapdemo.go), to verify quick-and-dirty heap implementation
  - `go build cmd/heapdemo.go; ./heapdemo 5 2 1 0 3 4`
* [heap vizualization](cmd/heapviz.go) using [GraphViz](https://graphviz.org/)
  - `go build cmd/heapviz.go; ./heapviz 5 2 1 0 3 4 > h.dot; dot -Tpng -o h.png h.dot;`
  - Then view PNG file `h.png`
* [test heap construction](cmd/buildcheck.go) and show backing storage on command line
  - `go build cmd/buildcheck.go; ./buildcheck 4 0 2 1 3 5`
* [try "is heap" test](cmd/arraycheck.go) on arbitrary arrays
  - `go build cmd/arraycheck.go; ./arraycheck 3 2 1; ./arraycheck 1 2 3`
  - My implementation of the problem's solution uses Go slices as heap structure backing storage,
  this program can give negative answers, where `buildcheck` should only give positive answers.
* [try Heap's algorithm](cmd/permutations.go) on integers from the command line
  - `go build cmd/permutations.go; ./permutations 1 2 3`
  - `go build cmd/permutations.go; ./permutations 1 2 3 4 5 | sort | uniq | wc -l`
    * Should get N! (N factorial), if all integers on command line are unique
  - If you input duplicate integers, you get a few non-unique permutations

## Interview Analysis

I'm not sure what this question is after.
It could be a question to help the interviewer decide if a candidate
has a grasp of algorithms.
At the very least, the candidate would have to know of and understand heaps
(and inside that, binary trees).

My implementation of this problem required
knowing how to permute N integers,
and knowing enough about heap properties and implementations to use a Go slice
as the underlying data storage.
