package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"goProject3/models"

	es "github.com/olivere/elastic/v7"
)

type UserES struct {
	client *es.Client
	index  string
}

func NewUserES(client *es.Client) *UserES {
	return &UserES{
		client: client,
		index:  "users",
	}
}

func (u *UserES) CreateUser(ctx context.Context, user models.User2) error {
	_, err := u.client.Index().
		Index(u.index).
		Id(strconv.Itoa(int(user.ID))).
		BodyJson(user).
		Do(ctx)
	if err != nil {
		return fmt.Errorf("Cannot create user in ES")
	}
	return nil
}

func (u *UserES) DeleteUser(ctx context.Context, userid int) error {
	_, err := u.client.Delete().Index(u.index).Id(strconv.Itoa(userid)).Do(ctx)
	if err != nil {
		return fmt.Errorf("Cannot delete user in ES")
	}
	return nil
}

func (u *UserES) SearchUser(ctx context.Context, query string) ([]models.User2, error) {
	esQuery := es.NewMultiMatchQuery(query, "name", "email").
		Fuzziness("2").MinimumShouldMatch("2")
	result, err := u.client.Search().Index(u.index).Query(esQuery).Do(ctx)

	if err != nil {
		return nil, fmt.Errorf("Search user error")
	}

	var users []models.User2
	for _, hit := range result.Hits.Hits {
		var user models.User2
		json.Unmarshal(hit.Source, &user)
		users = append(users, user)
	}

	return users, nil
}
