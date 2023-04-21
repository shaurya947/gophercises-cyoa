package story

type StoryArcObj struct {
	Title   string
	Story   []string
	Options []StoryArcOption
}

type StoryArcOption struct {
	Text, Arc string
}
