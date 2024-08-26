package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// io.Reader and io.Writer are interfaces that are used for reading and writing data.
func main() {
	str := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam suscipit lacinia risus, vel efficitur mi lobortis et. Maecenas tempus nulla eu viverra aliquam. In hendrerit tempus massa ut laoreet. Nam pretium posuere aliquam. Fusce eros enim, egestas id sollicitudin quis, vehicula eget tellus. Nam eu odio ac nisl facilisis semper. Quisque scelerisque sit amet diam ac aliquet. Aenean semper ante quis ipsum eleifend, eget tempor sapien tempor. Praesent scelerisque, turpis vel sodales condimentum, tortor diam semper lorem, accumsan mattis arcu dui ut mauris. Interdum et malesuada fames ac ante ipsum primis in faucibus. Integer vestibulum rhoncus egestas. Morbi velit massa, pulvinar at rutrum vitae, rhoncus id diam.\n\nMorbi finibus finibus tempus. Aenean non nisi eros. Donec fermentum dui neque, id mattis justo faucibus luctus. Duis eros magna, tempor non ultrices vestibulum, rhoncus in orci. Nunc eu bibendum orci. Praesent in velit justo. Cras mollis et est sed tristique. Sed finibus risus ac metus commodo, non tincidunt quam tincidunt. Sed quis sapien quam. Duis nisi urna, laoreet non nunc sit amet, hendrerit blandit nibh. Aenean rutrum, metus viverra rutrum rutrum, sem tortor aliquam sapien, sed sagittis est augue eu nisl. Sed tempor nisl et dictum rhoncus. Duis tempor ex arcu, a vehicula augue lacinia id. Donec et rutrum nisl, at pharetra ex."
	reader := strings.NewReader(str)
	writer := MyWriterTransformer{w: os.Stdout}

	buffer := make([]byte, 264)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}

		_, err = writer.Write(buffer[:n])
		if err != nil {
			panic(err)
		}
	}
}

type MyWriter struct{}

func (m MyWriter) Write(bytes []byte) (n int, err error) {
	fmt.Println("Read:", len(bytes), "bytes")
	fmt.Println(string(bytes))
	fmt.Println("")

	return len(bytes), nil
}

type MyWriterTransformer struct {
	w io.Writer
}

func (mw MyWriterTransformer) Write(bytes []byte) (n int, err error) {
	for i, b := range bytes {
		bytes[i] = b + 10
	}

	return mw.w.Write(bytes)
}
