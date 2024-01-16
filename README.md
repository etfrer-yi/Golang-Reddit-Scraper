<div align="center">

# Golang-Reddit-Scraper

</div>

<div align="center">

![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/etfrer-yi/Golang-Reddit-Scraper?color=blue)
![GitHub repo file count (file type)](https://img.shields.io/github/directory-file-count/etfrer-yi/Golang-Reddit-Scraper?color=red)
![GitHub language count](https://img.shields.io/github/languages/count/etfrer-yi/Golang-Reddit-Scraper?color=purple)
![GitHub top language](https://img.shields.io/github/languages/top/etfrer-yi/Golang-Reddit-Scraper?color=orange)

</div>

<div align="center">

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

</div>
Screenshots of performance comparisons:
## Single-threaded querying first:
![image](https://github.com/etfrer-yi/Golang-Reddit-Scraper/assets/77317763/a2f5845e-dc5d-4ae3-bf4e-c3510cb17f22)
## Multi-threaded querying first
![image](https://github.com/etfrer-yi/Golang-Reddit-Scraper/assets/77317763/c66553ed-ce53-440b-b14b-74591f149cee)
## Conclusions
Whether we run single-threaded or the multi-threaded (goroutine-based) way of querying for Reddit data, it all amounts to the same result: the multi-threaded version runs faster by a bit.
There are, however, some rare cases where the multi-threaded version runs slower when run first.
