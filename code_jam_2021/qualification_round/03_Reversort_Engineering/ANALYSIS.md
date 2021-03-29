## Analysis

### Test Set 1

The solution to the problem is a permutation of the numbers from 1 to N. The cost of each permutation can be calculated by simulating the Reversort algorithm as described in the [analysis of the Reversort problem](https://codingcompetitions.withgoogle.com/codejam/round/000000000043580a/00000000006d0a5c#analysis) in O(N^2) time complexity. There are N! distinct permutations of size NN, containing the numbers from 1 to N exactly once each. The cost of each permutation can be calculated and the answer is any permutation that has a cost less than or equal to C. If there is no such permutation, output `IMPOSSIBLE`. The time complexity of the overall solution is O(N!⋅N^2).

### Test set 2

As N is large for Test Set 2, we cannot generate every possible permutation. The major observation here is that the range of valid costs for a given N lies between N-1 (when the cost of each reverse operation is the minimum possible, which is 1) and N⋅(N+1)/2-1 (when the cost of each reverse operation is the maximum possible, which is N-i. Cost = N-1 when the array is already sorted.

All costs in between those two limits are possible, as we shall see. Hence, if C is not in the valid range for given N, output `IMPOSSIBLE`. Otherwise, we perform the following construction by recursion, which also serves as proof that the costs in range are indeed possible. The first iteration costs between 1 and N, so we should choose a cost x for it such that C-x, fits in the possible range for a permutation of size. You can check that this is always possible, and even compute the full range of x values that work by solving the system of inequalities.

Now, recursively generate a permutation P of size N-1 and cost C-x. Then, add 1 to all integers in PP and insert 1 at its left end, getting a new permutation of integers between 1 and N. Then, reverse the prefix of P of length x as the cost of the initial iteration should be x. The non-recursive steps take O(N) to adjust P. Since we perform those for each index, the overall complexity of the solution is O(N2).
