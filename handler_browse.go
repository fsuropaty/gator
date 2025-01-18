package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/fsuropaty/gator-go/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) > 0 {
		convertedArgs, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			fmt.Printf("Couldn't convert the args... using default limit of %d\n", limit)
		} else {
			limit = convertedArgs
		}
	}

	getPostParams := database.GetPostForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.db.GetPostForUser(context.Background(), getPostParams)
	if err != nil {
		return fmt.Errorf("Couldn't get post for user: %w", err)
	}

	if len(posts) == 0 {
		fmt.Println("The posts is empty")
		return nil
	}

	for _, post := range posts {
		fmt.Println("Post Title: ", post.Title)
		fmt.Println("Post URL: ", post.Url)
		fmt.Println("Post Description: ", post.Description)
		fmt.Println("Post Published At: ", post.PublishedAt)
		fmt.Println()
	}

	return nil
}
