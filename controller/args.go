package controller

import "flag"

type Args struct {
	Port       int
	Origin     string
	ClearCache bool
}

func ParseArgs() Args {
	port := flag.Int("port", 0, "Port for the proxy server.")
	origin := flag.String("origin", "", "Origin server URL.")
	clearCache := flag.Bool("clear", false, "Clear the cache.")
	flag.Parse()

	return Args{
		Port:       *port,
		Origin:     *origin,
		ClearCache: *clearCache,
	}
}
