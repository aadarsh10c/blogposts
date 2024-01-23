package blogposts_test

import (
	"errors"
	"io/fs"
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
		if err != nil{
			t.Fatal(err)
		}
	})
}
