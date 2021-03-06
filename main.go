package main

import (
	"fmt"
	"github.com/msvens/photoeditor/pkg"
)

func main() {
	//e := photos.InstaEditor(400, 400)
	e := editor.NewEditor(1200, 400, 400)
	var err error
	if err = editor.CreateDirs("test", true); err != nil {
		fmt.Println(err)
		return
	}
	err = e.GenerateAll("samples/square.jpg", "test", true)
	fmt.Println(err)
}
