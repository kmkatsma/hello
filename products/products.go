package products

import "github.com/kataras/iris"

//Product is the representation of a product.
type Product struct {
	ProductName string `json:"productName"`
	ProductID   string `json:"productId"`
}

//Controller is an interface to represent a controller that handles requests.
type Controller interface {
	Get(*iris.Context)
	Post(*iris.Context)
}

//Get returns a product based on id.
func Get(ctx *iris.Context) {
	// Get id from path '/products/:id'
	id := ctx.Param("id")
	p := Product{"test", id}
	ctx.JSON(iris.StatusOK, p)
}

//Post creates a product.
func Post(ctx *iris.Context) {
	p := new(Product)
	if err := ctx.ReadJSON(p); err != nil {
		ctx.EmitError(iris.StatusBadRequest)
		return
	}
	ctx.JSON(iris.StatusCreated, p)
}
