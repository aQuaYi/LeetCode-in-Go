# [925. Long Pressed Name](https://leetcode.com/problems/long-pressed-name/)

Your friend is typing his `name` into a keyboard.  Sometimes, when typing a character `c`, the key might get long pressed, and the character will be typed 1 or more times.

You examine the `typed` characters of the keyboard.  Return `True` if it is possible that it was your friends name, with some characters (possibly none) being long pressed.

Example 1:

```text
Input: name = "alex", typed = "aaleex"
Output: true
Explanation: 'a' and 'e' in 'alex' were long pressed.
```

Example 2:

```text
Input: name = "saeed", typed = "ssaaedd"
Output: false
Explanation: 'e' must have been pressed twice, but it wasn't in the typed output.
```

Example 3:

```text
Input: name = "leelee", typed = "lleeelee"
Output: true
```

Example 4:

```text
Input: name = "laiden", typed = "laiden"
Output: true
Explanation: It's not necessary to long press any character.
```

Note:

1. `name.length <= 1000`
1. `typed.length <= 1000`
1. The characters of `name` and `typed` are lowercase letters.