package blogposts

import (
	"io/fs"
)

type Posts struct {
}

func NewPostFromFS(fileSystem fs.FS) ([]Posts, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Posts
	for range dir {
		posts = append(posts, Posts{})
	}
	return posts, nil
}
