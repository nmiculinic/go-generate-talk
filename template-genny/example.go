package template_text

type FInt interface {
	Filter(map[string]int) map[string]int
}

type FString interface {
	Filter(map[string]string) map[string]string
}

