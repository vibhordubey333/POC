package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"encoding/json"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gofiber/fiber/v2"
)

var cache = memcache.New("localhost:11211")

func verifyCache(c *fiber.Ctx) error {
	id := c.Params("id")
	val, err := cache.Get(id)
	if err != nil {
		return c.Next()
	}

	data := toJson(val.Value)
	return c.JSON(fiber.Map{"Cached": data})
}

func main() {
	app := fiber.New()

	app.Get("/:id", verifyCache, func(c *fiber.Ctx) error {
		id := c.Params("id")
		res, err := http.Get("https://jsonplaceholder.typicode.com/photos/" + id)
		if err != nil {
			panic(err)
		}

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		cacheErr := cache.Set(&memcache.Item{Key: id, Value: body, Expiration: 10})
		if cacheErr != nil {
			panic(cacheErr)
		}

		data := toJson(body)
		return c.JSON(fiber.Map{"Data": data})
	})
	log.Println("Started application:")
	app.Listen(":3000")
}
func toJson(val []byte) Photo {
	photo := Photo{}
	err := json.Unmarshal(val, &photo)
	if err != nil {
		panic(err)
	}
	return photo
}

type Photo struct {
	AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}
