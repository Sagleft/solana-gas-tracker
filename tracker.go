package solanagastracker

type Tracker interface{}

type defaultTracker struct{}

func New() Tracker {
	return &defaultTracker{}
}
