package models

import (
	"errors"
	"fmt"
)

type News struct {
	Author    string   `json:"author"`
	Title     string   `json:"title"`
	Summary   string   `json:"summary"`
	CreatedAt string   `json:"createdAt"`
	Content   string   `json:"content"`
	Source    string   `json:"source"`
	Tags      []string `json:"tags"`
}

func (n News) Validate() (err error) {
	if n.Author == "" {
		err = errors.Join(err, fmt.Errorf("The author is empty: %s", n.Author))
	}
	if n.Title == "" {
		err = errors.Join(err, fmt.Errorf("The title is empty: %s", n.Title))
	}
	if n.Summary == "" {
		err = errors.Join(err, fmt.Errorf("The summary is empty: %s", n.Summary))
	}
	if n.Content == "" {
		err = errors.Join(err, fmt.Errorf("The content is empty: %s", n.Content))
	}

	return nil
}
