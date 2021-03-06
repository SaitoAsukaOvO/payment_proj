/*
 * EAP Jupyter Backend API
 *
 * This file contains REST API specification for Kubeflow notebooks. The file is autogenerated from the swagger definition.
 *
 * API version: 0.1.38
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment_proj/controllers"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name        string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method      string
	// Pattern is the pattern of the URI.
	Pattern     string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

var routes = Routes{
	{
		"PublicKey",
		http.MethodGet,
		"/api/publicKey",
		controllers.PublicKey,
	},
	{
		"CreateUser",
		http.MethodPost,
		"/api/user",
		controllers.CreateUser,
	},
	{
		"GetUser",
		http.MethodGet,
		"/api/user",
		controllers.GetUser,
	},
	{
		"CreateWallet",
		http.MethodPost,
		"/api/wallet",
		controllers.CreateWallet,
	},
	{
		"CreateTransaction",
		http.MethodPost,
		"/api/transaction",
		controllers.CreateTransaction,
	},
}
