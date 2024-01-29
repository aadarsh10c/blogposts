package blogposts

import (
	"bufio"
	"io"
	"io/fs"
)

type Post struct {
	Title       string
	Description string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
)

func NewPostFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(filesSystem fs.FS, fileName string) (Post, error) {
	//open the file
	postFile, err := filesSystem.Open(fileName)
	if err != nil {
		return Post{}, nil
	}
	defer postFile.Close()

	//read th contents of file
	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	//read title
	titleLine := readLine(scanner)[len(titleSeparator):]

	//read the description
	description := readLine(scanner)[len(descriptionSeparator):]

	post := Post{Title: titleLine, Description: description}
	return post, nil
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	text := scanner.Text()
	return text
}
