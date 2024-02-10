package gifenc

type Output struct {
	Name string
	Path string
}

type Config struct {
	Output Output
	Delay  int // encode delay
	Width  int // decode width
	Height int // decode height
}
