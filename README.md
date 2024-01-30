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

I observed that common max-heap implementations use a variable-length array
to keep the elements or nodes or values.
If an element is at index `n` in the array,
its child elements are at indexes `2n+1` and `2n+2`.
Its parent node is at index `(n-1)/2`, which is conveniently rounded down
by ones-complement integer division.

Once you've put elements or values in such a max heap,
you've got an array of integers.
Such an array can be examined to see if it's got the [heap data structure](https://en.wikipedia.org/wiki/Heap_(data_structure))
partial ordering.
Filling in the array backing storage guarantees the "shape" property of
a heap data structure is preserved.

The partial ordering property can be satisfied by some unintuitive binary trees:

![example 5-value max heap](example_max_heap.png)

The left-hand subtree has all elements greater than
the elements of the right-hand subtree.
I could not figure out a way to count these without algorithmically checking them,
nor could I figure out some way to "flip" subtrees to get an arbitrary
heap value arrangement.

I decided to look at all permutations of the list of `N` integers,
running the "is this a heap?" test on each permutation.
Convenient implementations of [Heap's algorithm](https://en.wikipedia.org/wiki/Heap's_algorithm)
for permuting a list of items are recursive.
I took advantage of the [Go]() programming language's easy casual concurrency.
I ran Heap's algorithm in one goroutine,
having code in that goroutine write each permutation to a channel.
The main goroutine reads permutations from the channel,
counting unique permutations.
The problem statement doesn't specify `N` unique integers,
so it's possible that an input would have duplicates.
Two or more "permutations" of the `N` integer list would then be duplicates.

This is a [design pattern](https://bruceediger.com/posts/golang-enabled-pattern/)
that can be useful quite often.
In this case, it divides the work, putting the permutation into one thread,
and unique-detection, counting and output in a second thread.

Here are the 8 ways to have legit heaps when the integer values in the
heap are `0, 1, 2, 3, 4`:

![first 4 heaps](big1.png)
![second 4 heaps](big2.png)

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
My implementation of this problem (which may not be the best!) required
knowing how to permute N integers,
and knowing enough about heap properties and implementations to use a Go slice
as the underlying data storage.
