package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"goProject3/models"
	"strconv"

	es "github.com/olivere/elastic/v7"
)

type CommentES struct {
	client *es.Client
	index  string
}

func NewCommentES(client *es.Client) *CommentES {
	return &CommentES{
		client: client,
		index:  "comments",
	}
}

func (c *CommentES) CreateComment(ctx context.Context, comment models.Comment) error {
	_, err := c.client.Index().
		Index(c.index).
		Id(strconv.Itoa(int(comment.ID))).
		BodyJson(comment).
		Do(ctx)
	if err != nil {
		return fmt.Errorf("Cannot create comment in ES")
	}
	return nil
}

func (c *CommentES) DeleteComment(ctx context.Context, commentId int) error {
	_, err := c.client.Delete().Index(c.index).Id(strconv.Itoa(commentId)).Do(ctx)
	if err != nil {
		return fmt.Errorf("Cannot delete comment in ES")
	}
	return nil
}

func (c *CommentES) SearchComment(ctx context.Context, query string) ([]models.Comment, error) {
	esQuery := es.NewMultiMatchQuery(query, "comment").
		Fuzziness("2").MinimumShouldMatch("2")
	result, err := c.client.Search().Index(c.index).Query(esQuery).Do(ctx)

	if err != nil {
		return nil, fmt.Errorf("Search Comment error")
	}

	var comments []models.Comment
	for _, hit := range result.Hits.Hits {
		var comment models.Comment
		json.Unmarshal(hit.Source, &comment)
		comments = append(comments, comment)
	}

	return comments, nil
}
