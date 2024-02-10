package gifenc

type Output struct {
	Name string `default:"img"`
	Path string
}

type Config struct {
	Output Output
	Delay  int // encode delay
	Width  int // decode width
	Height int // decode height
}
