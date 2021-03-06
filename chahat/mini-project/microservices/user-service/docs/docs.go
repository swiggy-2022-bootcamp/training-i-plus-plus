// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Chahat Bhatia",
            "email": "chahatbhatia2014@gmail.com"
        },
        "license": {
            "name": "Apache 2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/signup": {
            "post": {
                "description": "This request will register a new user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "To register a new user in the online shopping application",
                "parameters": [
                    {
                        "description": "User details",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "number"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "Get details of all Users.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get all Users list.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.User"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "/users/addtocart/:userid": {
            "post": {
                "description": "This request will add a product in user's cart.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User can add products in cart.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Product to add",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ProductUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "number"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "number"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "/users/getcart/:userid": {
            "get": {
                "description": "This request will add a product in user's cart.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User can view products in cart.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.ProductUser"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "number"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "number"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "This request will login a user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "User gets logged in.",
                "parameters": [
                    {
                        "description": "password",
                        "name": "Password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "email id",
                        "name": "Email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "number"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "number"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "View individual user details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User by ID.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "Edit details of a User.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Edit User by ID.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "User details",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "number"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "User can delete his account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete User by ID.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user token",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "number"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ProductUser": {
            "type": "object",
            "required": [
                "category",
                "price",
                "seller",
                "title"
            ],
            "properties": {
                "category": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "product_id": {
                    "type": "string"
                },
                "ratings": {
                    "type": "integer"
                },
                "seller": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name",
                "password",
                "phone",
                "usertype"
            ],
            "properties": {
                "_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_cart": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ProductUser"
                    }
                },
                "user_id": {
                    "type": "string"
                },
                "usertype": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer Token": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Online Shopping Application - User Module",
	Description:      "This microservice is for user module in the online shopping application.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
