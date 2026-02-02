package generator

type ReadmeDoc struct {
	Title       string
	Description string
	Sections    []Section
}

type Section struct {
	Name    string
	Content string
}

type CommentDoc struct {
	FilePath string
	Target   string
	Text     string
}

type DocBundle struct {
	Readme   *ReadmeDoc
	Comments []CommentDoc
}
