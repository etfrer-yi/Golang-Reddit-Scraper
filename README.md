# Golang-Reddit-Scraper

Screenshots of performance comparisons:
# Single-threaded querying first:
![image](https://github.com/etfrer-yi/Golang-Reddit-Scraper/assets/77317763/a2f5845e-dc5d-4ae3-bf4e-c3510cb17f22)
![image](https://github.com/etfrer-yi/Golang-Reddit-Scraper/assets/77317763/c66553ed-ce53-440b-b14b-74591f149cee)
# Multi-threaded querying first
Whether we run single-threaded or the multi-threaded (goroutine-based) way of querying for Reddit data, it all amounts to the same result: the multi-threaded version runs faster by a bit.
There are, however, some rare cases where the multi-threaded version runs slower when run first.
