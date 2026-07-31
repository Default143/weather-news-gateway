package news

import "testing"

func TestArticleStruct(t *testing.T) {
	article := Article{
		Title:       "Test title",
		Description: "Test description",
		URL:         "https://example.com",
	}

	if article.Title != "Test title" {
		t.Errorf("expected title Test title, got %s", article.Title)
	}

	if article.URL == "" {
		t.Error("expected URL to be set")
	}
}
