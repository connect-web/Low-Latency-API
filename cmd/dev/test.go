package main

import cache "github.com/connect-web/low-latency-cache-controller/wrapper"

func main() {
	//globalstats.GetGlobalStats()
	cache.StartUp("http://127.0.0.1:4050")
}
