// Code generated by bin from example.go. DO NOT EDIT.

package template_text

type EmptyInt struct{}

func (f *EmptyInt) Filter(in map[string]int) map[string]int {
	return in
}

type ComposeInt []FInt

func (f *ComposeInt) Filter(in map[string]int) map[string]int {
	out := in
	for _, filter := range f {
		out = filter(out)
	}
	return out
}


type EmptyString struct{}

func (f *EmptyString) Filter(in map[string]string) map[string]string {
	return in
}

type ComposeString []FString

func (f *ComposeString) Filter(in map[string]string) map[string]string {
	out := in
	for _, filter := range f {
		out = filter(out)
	}
	return out
}


