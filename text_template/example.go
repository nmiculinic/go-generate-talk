package text_template

//go:generate bash -c "go run ./bin/* -- $GOFILE"
type FInt interface {
	Filter(map[string]int) map[string]int
}

type FString interface {
	Filter(map[string]string) map[string]string
}
