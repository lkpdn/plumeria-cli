package route

type Route struct {
	Address string `param:"address" long:"address" description:"address (Length:1 - 46. Default:.)"`
	Gateway string `param:"gateway" long:"gateway" description:"gateway (Length:1 - 46. Default:.)"`
}
