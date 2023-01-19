package table

type State struct {
	Name            string
	GuideCharacters []string
	Shift           bool
	Error           bool
	NextStateID     int
	PushToStack     bool
	End             bool
}
