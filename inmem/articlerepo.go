package inmem

import (
	"context"
	"fmt"
	"sync"

	"github.com/DagmarC/simple-gokit-example/article"
)

type articlesRepository struct {
	mtx      sync.RWMutex
	articles map[string]article.Article
}

func NewArticlesRepository() *articlesRepository {
	return &articlesRepository{
		articles: map[string]article.Article{},
	}
}

// not safe
func (r *articlesRepository) GetArticle(ctx context.Context, id string) (article.Article, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	a, ok := r.articles[id]
	if !ok {
		return article.Article{}, fmt.Errorf("article wasn't found")
	}

	return a, nil
}

// not safe
func (r *articlesRepository) InsertArticle(ctx context.Context, article article.Article) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.articles[article.ID] = article
	return nil
}
