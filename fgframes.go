package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const BasePath = "data/title"

type TitleMap struct {
	name         string
	characterMap map[string]Character
}

type Character struct {
	name string
	data string
}

func BuildData(title string) map[string]Character {
	dataMap := make(map[string]Character)

	titleMap := map[string]bool{
		"usf4": true,
		"sfv":  true,
	}
	_, ok := titleMap[title]

	if !ok {
		return dataMap
	}

	names := CharacterNames(title)
	for _, name := range names {
		_, data := FindData(title, name)
		dataMap[name] = Character{name, data}
	}
	return dataMap
}

func CharacterNames(title string) []string {
	path := BasePath + "/" + title + "/" + "json"
	names := []string{}
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		baseName := file.Name()
		name := strings.TrimSuffix(baseName, filepath.Ext(baseName))
		names = append(names, name)
	}
	return names
}

func FindData(title string, name string) (bool, string) {
	path := BasePath + "/" + title + "/" + "json" + "/" + name + ".json"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return false, "Character not found"
	}
	d := string(data)
	return true, d
}

func main() {
	router := gin.Default()
	usf4Map := TitleMap{name: "usf4", characterMap: BuildData("usf4")}

	router.GET("/:title/:name", func(context *gin.Context) {
		switch title := context.Param("title"); title {
		case "usf4":
			name := context.Param("name")
			if strings.EqualFold(name, "characters") {
				context.JSON(200, gin.H{
					"characters": CharacterNames(title),
				})
			} else {
				character, characterFound := usf4Map.characterMap[name]
				if !characterFound {
					context.JSON(404, gin.H{
						"message": "character not found",
					})
				}
				context.String(200, character.data)
			}
		case "sfv":
			context.JSON(501, gin.H{
				"message": "coming soon",
			})
		default:
			context.JSON(404, gin.H{
				"message": "title not found",
			})
		}

	})

	router.Run(":8080")
}
