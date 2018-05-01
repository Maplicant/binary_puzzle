# Binary Puzzle Solver
## What is a binary puzzle?
A binary puzzle is a puzzle that consists of a square grid partially filled with 0's and 1's. The goal is to fill in the grid such that every square is filled. There are rules on where you're allowed to place 0's and 1's, however. There are two rules:
1. Every row and column must have the same amount of 0's as 1's
2. There must not be three identical numbers in a row, whether it be horizontal or vertical.

This program will take a binary puzzle and give you its solution

## Input
This program expects a puzzle from stdin in a special format. The first line is an integer `n`, the width and height of the grid. The next `n` lines are strings with the length of `n`. Each character in the string is either a `0`, a `1` or a `?`. An example:

    4
    11?0
    ??1?
    0??1
    ??0?




