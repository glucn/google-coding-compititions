**SPOILER ALERT: THIS IS THE OFFICAL ANALYSIS FROM GOOGLE. STOP READING WHEN TRYING TO SOLVE THE PROBLEM ON YOUR OWN.**

## Analysis

### Test Set 1

In Test Set 1, we can try every possible path we can take. During any given minute we have 5 options: walk a block in one of the 4 directions, or stay still. Since Peppurr's tour is at most 8 minutes long, we can only walk for at most 8 minutes. That is at most 5^8^ possibilities, which is a pretty small number for a computer. For each combination, we simulate our own path and Peppurr's path, recording any encounters. After trying all possibilities, the solution is `IMPOSSIBLE` if we did not record any encounters, or the earliest time of those if we did.

### Test Set 2

In Test Set 2, just as in Test Set 1, Peppurr's tour will remain within a single north-south street. Notice that if we want to meet Peppurr at an intersection (a, b), the order in which we walk the blocks doesn't matter as long as we walk east a times more than west and north b times more than south. So we may assume that we can finish all of our eastward walking, for example, before walking in other direction. This means an optimal strategy can begin with **X** blocks of walking east. After that, we are in the same north-south street as Peppurr, so we can walk towards the tour until we meet or we are just 1 block away, in which case we need to stand still for 1 minute to avoid crossing paths with the tour in the middle of a block.

### Test Set 3

For Test Set 3, we can simulate Peppurr's tour. If after R minutes it is X<sub>R</sub> blocks to the east of us and Y<sub>R</sub> blocks to the north (X<sub>R</sub> and Y<sub>R</sub> can be negative to represent being west or south), we only need to check whether we can reach that intersection in R minutes. Fortunately, this is easy to check: the intersection is reachable in R minutes or less if and only if |X<sub>R</sub>| + |Y<sub>R</sub>| ≤ R. That is, the intersection must be within an [L1 distance](https://en.wikipedia.org/wiki/Taxicab_geometry) (also known as Manhattan Distance) of R.

Therefore, we can solve the problem by simulating Peppurr's path, and for the i-th intersection visited, check if it's reachable in i minutes. If it is, then i is the answer; otherwise, we need to keep looking. If none of the intersections is reachable within the required time, we answer `IMPOSSIBLE`.
