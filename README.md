# go-dsa

[![Go Report Card](https://goreportcard.com/badge/github.com/medunes/go-dsa)](https://goreportcard.com/report/github.com/medunes/go-dsa)
[![Go Reference](https://pkg.go.dev/badge/github.com/medunes/go-dsa.svg)](https://pkg.go.dev/github.com/medunes/go-dsa)
[![Tests](https://github.com/medunes/go-dsa/actions/workflows/tests.yml/badge.svg)](https://github.com/medunes/go-dsa/actions/workflows/tests.yml)

## Data Structures

### [Binary Search Tree (BST)](./bst/bst.go)
### [LRU Cache](https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU)

## Algorithms

### [Sorting](./sort/sort.go)

1- [Selection Sort](https://en.wikipedia.org/wiki/Selection_sort)

2- [Insertion Sort](https://en.wikipedia.org/wiki/Insertion_sort)

3- [Bubble Sort](https://en.wikipedia.org/wiki/Bubble_sort)

4- [Quick Sort / Lomuto](https://en.wikipedia.org/wiki/Quicksort#Lomuto_partition_scheme)

## Running The tests

```bash
$ medunes@medunes:~/projects/go-dsa$ make test
go run gotest.tools/gotestsum@latest --format=testdox
github.com/medunes/go-dsa:

github.com/medunes/go-dsa/dsa:
 ✓ Insertion (0.00s)
 ✓ Insertion array with zeroes (0.00s)
 ✓ Insertion ascending order (0.00s)
 ✓ Insertion boundary values (0.00s)
 ✓ Insertion descending order (0.00s)
 ✓ Insertion descending order with duplicates (0.00s)
 ✓ Insertion duplicate elements (0.00s)
 ✓ Insertion empty (0.00s)
 ✓ Insertion large array (0.00s)
 ✓ Insertion mixed sign numbers (0.00s)
 ✓ Insertion more than one element (0.00s)
 ✓ Insertion more than one element 2 (0.00s)
 ✓ Insertion more than one element with repetition (0.00s)
 ✓ Insertion negative numbers (0.00s)
 ✓ Insertion one element (0.00s)
 ✓ Insertion random order (0.00s)
 ✓ Insertion same element (0.00s)
 ✓ Insertion single negative element (0.00s)
 ✓ Selection (0.00s)
 ✓ Selection array with zeroes (0.00s)
 ✓...
 ✓...
DONE 36 tests in 0.081s
```
