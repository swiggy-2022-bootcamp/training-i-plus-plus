// Package docs GENERATED BY SWAG; DO NOT EDIT
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
            "name": "Swastik Sahoo",
            "email": "swastiksahoo22@gmail.com"
        },
        "license": {
            "name": "Apache 2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ticket/book": {
            "post": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "To book ticket.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ticket"
                ],
                "summary": "Book Ticket",
                "parameters": [
                    {
                        "description": "Ticket structure",
                        "name": "ticket",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Ticket"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
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
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "/ticket/delete/{pnr_number}": {
            "delete": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "To remove a particular ticket.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ticket"
                ],
                "summary": "Delete Ticket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "PNR Number",
                        "name": "pnr_number",
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
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "/ticket/get/{pnr_number}": {
            "get": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "To get ticket details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ticket"
                ],
                "summary": "Get Ticket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "PNR Number",
                        "name": "pnr_number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Ticket"
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "/ticket/getall": {
            "get": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "To get every ticket detail.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ticket"
                ],
                "summary": "Get all Ticket details",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Ticket"
                            }
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "/ticket/update/{pnr_number}": {
            "patch": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "To update ticket details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ticket"
                ],
                "summary": "Update Ticket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "PNR Number",
                        "name": "pnr_number",
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
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Ticket": {
            "type": "object",
            "properties": {
                "date_time": {
                    "type": "string"
                },
                "destination": {
                    "type": "string"
                },
                "passenger_name": {
                    "type": "string"
                },
                "pnr_number": {
                    "type": "integer"
                },
                "seat_number": {
                    "type": "integer"
                },
                "source": {
                    "type": "string"
                },
                "train_number": {
                    "type": "integer"
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
	Host:             "localhost:8082",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Ticket Module",
	Description:      "This microservice is for ticket module.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
