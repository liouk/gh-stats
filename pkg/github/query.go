package github

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
)

func RunSimpleQuery(c *githubv4.Client) error {
	var query struct {
		Viewer struct {
			Login     githubv4.String
			CreatedAt githubv4.DateTime
		}
	}

	err := c.Query(context.Background(), &query, nil)
	if err != nil {
		return err
	}

	fmt.Println("    Login:", query.Viewer.Login)
	fmt.Println("CreatedAt:", query.Viewer.CreatedAt)
	return nil
}
