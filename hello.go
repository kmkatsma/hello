package main

import (
	"fmt"
	"math"

	"github.com/kataras/iris"
	"github.com/kmkatsma/hello/models"
	"github.com/kmkatsma/hello/os"
	"github.com/kmkatsma/hello/products"
)

func main() {

	i := os.NewWrapper()
	i.RunsOnWindows()

	// http://localhost:5700/api/user/42
	// Method: "GET"
	iris.Get("/api/user/:id", func(ctx *iris.Context) {

		// take the :id from the path, parse to integer
		// and set it to the new userID local variable.
		userID, _ := ctx.ParamInt("id")

		user := iris.Map{"username": "Joe", "userid": userID}
		// userRepo, imaginary database service <- your only job.
		//user := userRepo.GetByID(userID)

		// send back a response to the client,
		// .JSON: content type as application/json; charset="utf-8"
		// iris.StatusOK: with 200 http status code.
		//
		// send user as it is or make use of any json valid golang type,
		// like the iris.Map{"username" : user.Username}.
		ctx.JSON(iris.StatusOK, user)

	})

	iris.Get("/products/:id", products.Get)
	iris.Post("/products", products.Post)
	//iris.Put("products/:id", editProduct)
	//iris.Delete("/products/:id", deleteProduct)

	iris.Listen(":8080")

}

func createVertex(n int) {
	v := models.Vertex{X: 1, Y: 2}
	v.X = n
}

func counting() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func arrays() {
	var a [2]string
	a[0] = "Hello"
}

func pointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
