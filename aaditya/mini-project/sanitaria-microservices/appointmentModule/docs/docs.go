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
            "name": "Aaditya Khetan",
            "email": "aadityakhetan123@gmail.com"
        },
        "license": {
            "name": "Apache 2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/appointments": {
            "get": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "This request will fetch all the available appointments(open slots).",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Appointment"
                ],
                "summary": "To display all the available appointments.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Appointment"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
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
        "/book-appointment/{userId}": {
            "post": {
                "description": "This request will book and an appointment for the user with given userId.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Appointment"
                ],
                "summary": "To book a doctor's appointment.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "GeneralUser id",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Appointment"
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
        "models.Appointment": {
            "type": "object",
            "properties": {
                "doctorID": {
                    "type": "string"
                },
                "fees": {
                    "type": "integer"
                },
                "generalUserID": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "occupied": {
                    "type": "boolean"
                },
                "slot": {
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
	Host:             "localhost:8082",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Sanitaria - Doctor Module",
	Description:      "This microservice is for doctor module in the sanitaria application.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}