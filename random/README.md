## Notable statistics:

> Problem: Calculate the sum of (1^1 + 2^2 + .... + n^n) % 12345678

**Normal aggregation**: Simple loop from 1->n

**Goroutines aggregation**: Generate a goroutine for each int in 1->n, and atomically add to the sum

**Worker aggregation**: Generate X workers to receive any int from 1->n, and summarize the integers. Doing aggregation on the result set later

|                              | Normal aggregation | Goroutines aggregation | Worker aggregation(16 workers) |
|------------------------------|--------------------|------------------------|--------------------------------|
| O(n) power - n = 10000       | 210.87774ms        | 219.6336ms             | 61.128447ms                    |
| O(log n) power - n = 1000000 | 222.848129ms       | 378.557311ms           | 230.724272ms                   |