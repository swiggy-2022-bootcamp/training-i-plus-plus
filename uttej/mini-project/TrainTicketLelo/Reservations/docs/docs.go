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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Uttej Immadi",
            "url": "http://www.swagger.io/support",
            "email": "immadiuttej@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/:userId/reservations": {
            "get": {
                "description": "Get All Tickets for a user by providing the id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reservation"
                ],
                "summary": "Fetch All Tickets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Reservation"
                            }
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
        "/reservation": {
            "post": {
                "description": "Buy a Ticket by providing the details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reservation"
                ],
                "summary": "Buy a Ticket",
                "parameters": [
                    {
                        "description": "id, Departure Date, Purchase Date \u0026 Status Will Be Populated Automatically",
                        "name": "Ticket",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Reservation"
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
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "/reservation/:ticketId/cancel": {
            "post": {
                "description": "Cancel a Ticket that you've reserved",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reservation"
                ],
                "summary": "Cancel a Ticket",
                "parameters": [
                    {
                        "description": "unique ticket id",
                        "name": "TicketId",
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
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "/reservation/:ticketId/payment": {
            "post": {
                "description": "Pay For a ticket that you've reserved",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reservation"
                ],
                "summary": "Pay For a Ticket",
                "parameters": [
                    {
                        "description": "unique ticket id",
                        "name": "TicketId",
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
        "models.Reservation": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "departureDate": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "purchaseDate": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "trainIDs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "userId": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8002",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger TrainTicketLelo Reservations Service",
	Description:      "Swagger TrainTicketLelo Reservations Service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
