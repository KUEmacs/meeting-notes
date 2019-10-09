package main

import (
	"fmt"
	"os"

	"github.com/thecsw/mira"
)

func main() {
	r, err := mira.Init(mira.ReadCredsFromFile("login.conf"))
	if err != nil {
		fmt.Println("Error during login: ", err)
		os.Exit(1)
	}

	// Start streaming subreddits' submissions
	c, _, _ := r.Subreddit("tifu").StreamSubmissions()
	for {
		post := <-c
		r.Submission(post.GetId()).Save("You really f*cked up.")
	}
}
