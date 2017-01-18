package products

import (
	"github.com/labstack/echo"
	"net/http"
)

var (
	products = map[string]*Product{}
	seq      = 1
)

//Product is the representation of a product.
type Product struct {
	ProductName string `json:"productName"`
	ProductID   string `json:"productId"`
}

//Get returns a product based on id.
func Get(ctx echo.Context) error {
	// Get id from path '/products/:id'
	id := ctx.Param("id")
	if p, ok := products[id]; ok {
		return ctx.JSON(http.StatusOK, p)
	}
	return ctx.NoContent(http.StatusNotFound)

}

//Post creates a product.
func Post(ctx echo.Context) error {
	p := new(Product)
	if err := ctx.Bind(p); err != nil {
		return err
	}
	products[p.ProductID] = p
	return ctx.JSON(http.StatusCreated, p)
}
