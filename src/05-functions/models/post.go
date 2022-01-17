package models

import (
	"fmt"
)

type Post struct {
	Text     string
	Hasthags []Hashtag
	Posts    []Post
}

func (p Post) findHashtags() (hashtags []Hashtag) {
	hashtagFound := false
	hashtag := ""

	for i, c := range p.Text {
		if c == ' ' {
			hashtagFound = false
		} else if c == '#' {
			hashtagFound = true
		}
		if hashtagFound {
			hashtag += string(c)
		}
		if (c == ' ' || i == len(p.Text)-1) && hashtag != "" {
			hashtags = append(hashtags, Hashtag{Text: hashtag})
			hashtag = ""
		}
	}

	return // return variable/value already defined in function signature
}

func (p Post) ProcessHashtags() {
	fmt.Println(p.findHashtags())

	for _, post := range p.Posts {
		post.ProcessHashtags()
	}
}
