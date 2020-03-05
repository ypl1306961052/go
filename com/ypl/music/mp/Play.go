package mp

type Play struct {
	Stat     int
	Progress int
}
type Player interface {
	Player(source string)
}

func (m *Play) Player(source string) {

}

type Mp3Player struct {
	*Play
}
type AvmPlayer struct {
	*Play
}
