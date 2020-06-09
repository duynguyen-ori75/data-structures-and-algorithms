# Cache benchmark information

## Machine info
Lenovo Thinkpad T490s
CPU Core I5 - Ram 8Gb
Cache information (in bytes)

```
LEVEL1_ICACHE_SIZE                 32768
LEVEL1_ICACHE_LINESIZE             64
LEVEL1_DCACHE_SIZE                 32768
LEVEL1_DCACHE_LINESIZE             64
LEVEL2_CACHE_SIZE                  262144
LEVEL2_CACHE_LINESIZE              64
LEVEL3_CACHE_SIZE                  6291456
LEVEL3_CACHE_LINESIZE              64
```

## Benchmark result

```
goos: linux
goarch: amd64
BenchmarkMissCacheline-8   	       1	1350396934 ns/op
BenchmarkHitCacheline-8    	      15	  67326674 ns/op
PASS
```