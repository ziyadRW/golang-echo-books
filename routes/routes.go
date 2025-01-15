package routes

func SetUpRoutes(g *echo.Group) {
	
	g.GET("/books", book)

}