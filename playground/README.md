## Notable statistics:

> Problem: Calculate the sum of (1^1 + 2^2 + .... + n^n) % 12345678

**Normal aggregation**: Simple loop from 1->n

**Goroutines aggregation**: Generate a goroutine for each int in 1->n, and atomically add to the sum

**Worker aggregation**: Generate X workers to receive any int from 1->n, and summarize the integers. Doing aggregation on the result set later

|                              | Normal aggregation | Goroutines aggregation | Worker aggregation(16 workers) |
|------------------------------|--------------------|------------------------|--------------------------------|
| O(n) power - n = 10000       | 210.026578ms       | 219.875791ms           | 58.084432ms                    |
| O(log n) power - n = 1000000 | 232.834191ms       | 376.371211ms           | 244.836398ms                   |