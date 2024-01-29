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
