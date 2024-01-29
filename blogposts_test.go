package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/aadarsh10c/blogposts"
)

type subFailFS struct {
}

func (s subFailFS) Open(name string) (fs.File, error) {
	return nil, errors.New("Always failing test")
}
func TestNewBlogPosts(t *testing.T) {
	t.Run("Read 2 files from the FS", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("hi")},
			"hello-world2.md": {Data: []byte("hola")},
		}

		posts, err := blogposts.NewPostFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("Got %d posts, but wanted %d posts", len(posts), len(fs))
		}
	})
	t.Run("Failed test", func(t *testing.T) {
		fs := subFailFS{}

		_, err := blogposts.NewPostFromFS(fs)
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Read title form the post", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("Title: Post 1")},
			"hello-world2.md": {Data: []byte("Title: post 2")},
		}
		posts, _ := blogposts.NewPostFromFS(fs)
		got := posts[0]
		want := blogposts.Post{Title: "Post 1"}

		assertPosts(got, want, t)

	})
	t.Run("Read file title and description", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1`
			secondBody = `Title: Post 2
Description: Description 2`
		)

		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, _ := blogposts.NewPostFromFS(fs)
		got := posts[0]

		assertPosts(got, blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
		}, t)
	})
}

func assertPosts(got blogposts.Post, want blogposts.Post, t *testing.T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %+v , wanted %+v", got, want)
	}
}
