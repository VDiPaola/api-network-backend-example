package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/VDiPaola/api-network-server/helpers"
	"github.com/VDiPaola/api-network-server/models"
	"github.com/VDiPaola/api-network-server/nodes"
	"github.com/gofiber/fiber/v2"
)

func NodeTest(c *fiber.Ctx) error {
	nodes.Request("http://127.0.0.1:4000/ping", helpers.RequestMethod.GET, nil, func(response *http.Response, err error) {
		//make sure response exists
		if err != nil {
			log.Fatalf("error from request: %s", err)
			return
		}

		if response == nil {
			log.Fatalf("Response nil")
			return
		}

		//check status code
		if response.StatusCode != 200 {
			log.Fatalf("Invalid Response")
			return
		}

		//get response data
		defer response.Body.Close()
		// read body
		resBody, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatalf("Error reading body of response: %s", err)
			return
		}

		//unmarshall
		var responseData models.NodeResponseType
		if err := json.Unmarshal(resBody, &responseData); err != nil {
			log.Fatalf("Error unmarshalling response: %s", err)
			return
		}

		log.Printf("res body: %s", string(resBody))

		if responseData.JSON != nil {
			c.Status(200).JSON(responseData.JSON)
			return
		} else {
			if responseData.Text != "" {
				c.Status(200).SendString(responseData.Text)
				return
			} else {
				log.Fatalf("No valid data in request")
				return
			}
		}

		//send data to client
	})
	return nil
}

func Ping(c *fiber.Ctx) error {
	test := struct{ Something string }{
		Something: "hi",
	}
	return c.Status(200).JSON(&test)
}
