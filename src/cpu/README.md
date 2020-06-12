# Cache benchmark information

## Machine info

- Lenovo Thinkpad T490s
- CPU Core I5 - Ram 8Gb
- Cache information (in bytes)

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

- CPU Cache line hit:
	- BenchmarkMissCacheline
	- BenchmarkHitCacheline
- False sharing: 
	- BenchmarkSimpleStruct
	- BenchmarkPaddingStruct - a simple value with padding

```
goos: linux
goarch: amd64
BenchmarkMissCacheline-8   	       2	 997796674 ns/op
BenchmarkHitCacheline-8    	      36	  32767078 ns/op
BenchmarkSimpleStruct-8    	     207	   5346920 ns/op
BenchmarkPaddingStruct-8   	     438	   2668565 ns/op
PASS
ok  	_/home/duynguyen/Workplace/learning/src/benchmark	9.908s
```
