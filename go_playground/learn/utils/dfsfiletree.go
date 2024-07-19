package utils

import (

	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// as is jump directory recursively searches through file system depth first search. not ment to be used as is. current implementation was to DFS the file tree using recursive calls. -- not for use -- practice go language 

func JumpDirectory(name string, currDir string) string {
	os.Chdir(currDir)
	cleanPaths, path := getDirs(currDir, name)

	if path != ""{
		return path
	}

	for _, absPath := range cleanPaths{
		path = JumpDirectory(name, absPath)
		if path != "" {
			break
		}
	}

	fmt.Println(path)
	return path
}


func searchVisDirs(dirs []fs.DirEntry, name string)([]string, string) {
	var visDirs []string
	for _, dir := range dirs{
		if dir.Name()[0] == '.' || !dir.IsDir() || dir.Name() == "node_modules"{
			continue
		}else if checkMatch(name, dir){
			path, _ := filepath.Abs(dir.Name())
			return nil, path
		}

		if len(dirs) != 0{
			path, err := filepath.Abs(dir.Name())
			if err != nil {
				fmt.Println(err)
			}
			visDirs = append(visDirs, path)
		}
		
	}
	return visDirs, ""
}

func checkMatch(name string, dir fs.DirEntry)bool{	
	fmt.Println(dir.Name())	
	return dir.Name() == name 
	
}

func getDirs(dir string, name string)([]string, string){
	dirs, err := os.ReadDir(dir)
	
	if err != nil {
		fmt.Println(err)
	}

	return searchVisDirs(dirs, name)
}

