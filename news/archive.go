package news

type Article struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	URL    string `json:"url"`
}

type Archive map[string][]Article

func NewArchive() Archive {
	return make(map[string][]Article, 0)
}

func (a Archive) collect(category string) {
	sources := getSources(category)
	articles := getArticles(sources)
	a[category] = articles
}

func (a Archive) result(category string) []Article {
	return a[category]
}
