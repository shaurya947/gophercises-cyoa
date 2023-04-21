package story

type Story map[string]StoryArc

type StoryArc struct {
	Title   string
	Story   []string
	Options []StoryArcOption
}

type StoryArcOption struct {
	Text, Arc string
}
