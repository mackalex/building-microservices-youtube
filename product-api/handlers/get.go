package handlers

import (
	"github.com/gofiber/fiber"
	"github.com/mackalex/building-microservices-youtube/product-api/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) ListAll(c *fiber.Ctx) {
	p.l.Println("[DEBUG] get all records")

	prods := data.GetProducts()

	err := c.JSON(prods)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}

// swagger:route GET /products/{id} products listSingle
// Return a list of products from the database
// responses:
//	200: productResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *Products) ListSingle(c *fiber.Ctx) {
	id := getProductID(c)

	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		fiber.NewError(fiber.StatusNotFound)
		c.JSON(&GenericError{Message: err.Error()})
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		fiber.NewError(fiber.StatusInternalServerError)
		c.JSON(&GenericError{Message: err.Error()})
		return
	}

	err = c.JSON(prod)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}
