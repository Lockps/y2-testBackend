package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	list := [6]string{"อาการสำคัญ 1", "อาการสำคัญ 2", "อาการสำคัญ 3", "อาการสำคัญ 4", "อาการรอง 1", "อาการรอง 2"}
	port := ":3030"
	samp := 0
	isacid := true
	app := fiber.New()
	app.Post("/", func(c *fiber.Ctx) error {
		if string(c.BodyRaw()) == "clear" {
			isacid = true
			samp = 0
			return c.SendString("adssa") 
		}
		if samp < 4 {
			if string(c.BodyRaw()) == "0" {
				isacid = false
			}
			samp++
			return c.SendString(list[samp-1] + "asd")
		}
		if samp >= 4 && samp < 6 {
			if isacid {
				samp++
				return c.SendString(list[samp-1])
			} else {
				return c.SendString("No")
			}
		}
		return c.SendString("yes")
	})

	app.Listen(port)
}
