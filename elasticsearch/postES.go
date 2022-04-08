package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"goProject3/models"
	"strconv"

	es "github.com/olivere/elastic/v7"
)

type PostES struct {
	client *es.Client
	index  string
}

func NewPostES(client *es.Client) *PostES {
	return &PostES{
		client: client,
		index:  "posts",
	}
}

func (p *PostES) CreatePost(ctx context.Context, post models.Post) error {
	_, err := p.client.Index().
		Index(p.index).
		Id(strconv.Itoa(int(post.ID))).
		BodyJson(post).
		Do(ctx)
	if err != nil {
		return fmt.Errorf("Cannot create post in ES")
	}
	return nil
}

func (p *PostES) DeletePost(ctx context.Context, postId int) error {
	_, err := p.client.Delete().Index(p.index).Id(strconv.Itoa(postId)).Do(ctx)
	if err != nil {
		return fmt.Errorf("Cannot delete post in ES")
	}
	return nil
}

func (p *PostES) SearchPost(ctx context.Context, query string) ([]models.Post, error) {
	esQuery := es.NewMultiMatchQuery(query, "caption", "image").
		Fuzziness("2").MinimumShouldMatch("2")
	result, err := p.client.Search().Index(p.index).Query(esQuery).Do(ctx)

	if err != nil {
		return nil, fmt.Errorf("Search post error")
	}

	var posts []models.Post
	for _, hit := range result.Hits.Hits {
		var post models.Post
		json.Unmarshal(hit.Source, &post)
		posts = append(posts, post)
	}

	return posts, nil
}
