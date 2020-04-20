**SPOILER ALERT: THIS IS THE OFFICAL ANALYSIS GIVEN BY GOOGLE.**

## Analysis

### Test Set 1

Test Set 1 is small enough to solve by hand. We can speed this up with a couple of observations:

-   We can notice that every position with an even (X + Y) (apart from the origin) --- hereafter an "even" position --- seems to be unreachable. We can prove to ourselves that this is true: our initial X and Y coordinates of (0, 0) are both even, but only the first of our possible jumps (the 1-unit one) is of an odd length, and all jumps after that are of even lengths. So there is no way to reach any other "even" position starting from the origin, no matter how much jumping we do.
-   We can find that all "odd" positions, on the other hand, can be reached using no more than 3 moves.
-   To speed up solving the "odd" positions, we can take advantage of symmetry, as suggested in the explanation for Sample Case #2. For example, if we learn that `EEN` is a solution for (3, 4), then we also know that `WWS` is a solution for (-3, -4), and `EES` is a solution for (3, -4), and so on. Because of all the horizontal, vertical, and diagonal symmetry, there are really only six fundamentally different cases!
-   We can check that our solutions are optimally short by using an argument like the one in the explanation for Sample Case #1. Any position with a Manhattan distance (that is, |**X**| + |**Y**|) of 1 cannot be reached in fewer than one jump; positions with Manhattan distances up to 3 and 7 require at least two or three jumps, respectively. If our solution lengths match these lower bounds --- and they probably do unless we have jumped in an unusually indirect way --- then they are valid.

### Test Set 2

Based on the observations above, we may think to try a [breadth-first search](https://en.wikipedia.org/wiki/Breadth-first_search) of all possible jumping paths, and continue until every "odd" (X, Y) position (with -100 ≤ X ≤ 100 and -100 ≤ Y ≤ 100) has been reached. It turns out that each such position is reachable in no more than 8 moves. We know that these solutions are optimally short because of the breadth-first nature of the search.

### Test Set 3

Suppose that (**X**, **Y**) = (7, 10). In what direction should we make our initial 1-unit jump? As we saw above, we need our final X coordinate to be odd, but it is currently even, and we have only one chance to go from even to odd. Moving north or south will make our Y coordinate odd, but then we will never have another chance to make that even and the X coordinate odd. So we should either move west or east. For now, let's guess that we will go west; we will revisit the other possibility later.

That jump will take us to (-1, 0), and we will next need to make a 2-unit jump. Notice that we can make this look identical to the original problem setup, if we make two changes:

1.  Shift (-1, 0) to be the new (0, 0). Then the goal becomes (8, 10) rather than (7, 10).
2.  Transform the scale of the problem such that a 2-unit jump (to a "neighboring" cell) becomes the new 1-unit jump. Then the goal becomes (4, 5) instead of (8, 10).

With this in mind, let's revisit our original decision to jump to the west. If we had jumped east instead, we would have ended up at (1, 0), and if we had changed the problem in the same way we did above, our new goal would have been (3, 5). But that would be an "even" position (after rescaling), which cannot be reached! So we had no choice after all; we had to move west to be able to eventually reach the goal. It's a good thing we were so lucky!

So now the problem has "reset", and we are at (0, 0) and trying to get to (4, 5). In what direction should we make our "first" jump? Now we know we must move vertically, since 5 is odd and we will only have "one chance" to go from even to odd. If we jump north, the next rescaling will have a target of (2, 2), but if we jump south, the target will be (2, 3), which is the "odd" position that we want. From there, we should jump south to change the target to (1, 2), then east to change the target to (0, 1). At that point, we have a choice between jumping north and reaching the goal, and jumping south (which could still allow us to reach the goal after some further moves, e.g. one more to the south and then one more to the north). But the problem requires that we choose the shortest solution, so we should jump right to the goal! Therefore, the answer in this case is `WSSEN`.

Notice that this method is deterministic: we always have only one choice out of the four possible directions. We can rule out two of them because they will not make the correct coordinate odd. Of the other two, the new goal states they would leave us with must differ only in one of the coordinates and only by exactly 1 unit, and therefore one must be an "odd" position and the other must be an "even" position. It is possible that that "even" position is the goal, in which case we should jump there, but otherwise, we must choose the "odd" position.

The above analysis also shows that the only time we have a choice is when one of those options is to jump directly to the goal, in which case we obviously should. So we can be confident that our method produces the shortest possible solution. (We also know that that solution is unique, since if we were to choose to not jump directly to the goal when we had that option, we would only end up with a longer solution.)

Our method has a running time which is logarithmic in the magnitudes of the coordinates, so it will solve this test set blazingly fast!

### A Code Jam callback!

This problem is a riff on the Pogo problem from Round 1C in 2013. If you were familiar with that problem, the analysis might have helped a bit with this one... but, like a well-designed pogo stick, Expogo is not *too* difficult to get a handle on anyway.