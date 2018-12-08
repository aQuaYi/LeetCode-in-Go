# [920. Number of Music Playlists](https://leetcode.com/problems/number-of-music-playlists/)

Your music player contains `N` different songs and she wants to listen to `L` (not necessarily different) songs during your trip.  You create a playlist so that:

- Every song is played at least once
- A song can only be played again only if `K` other songs have been played

Return the number of possible playlists. **As the answer can be very large, return it modulo** `10^9 + 7`.

Example 1:

```text
Input: N = 3, L = 3, K = 1
Output: 6
Explanation: There are 6 possible playlists. [1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1].
```

Example 2:

```text
Input: N = 2, L = 3, K = 0
Output: 6
Explanation: There are 6 possible playlists. [1, 1, 2], [1, 2, 1], [2, 1, 1], [2, 2, 1], [2, 1, 2], [1, 2, 2]
```

Example 3:

```text
Input: N = 2, L = 3, K = 1
Output: 2
Explanation: There are 2 possible playlists. [1, 2, 1], [2, 1, 2]
```

Note:

1. 0 <= K < N <= L <= 100
