v1: working solution (non submitable)
- Approach - Brute force - Read whole file at once
- works for sample test
- fails for 1B rows obviously since reading whole file at once


v2: working and submitable solution for 1B rows
- Approach - Using bufio read 1 line at a time and process it sequentially
- time taken - 6m15.571s
- CPU consumption ~ 16%
- Memory consumtion ~ 6MB