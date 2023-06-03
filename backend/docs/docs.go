// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/ride/:id": {
            "get": {
                "description": "Get Ride by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ride"
                ],
                "summary": "Get Ride by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id to filter by",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Ride"
                        }
                    },
                    "404": {
                        "description": "ride not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/ride/customer/:id": {
            "get": {
                "description": "Get all Rides from one customer.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ride"
                ],
                "summary": "Get all Rides from one customer",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id to filter by",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Ride"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "enum_models.RideStatus": {
            "type": "integer",
            "enum": [
                0,
                1
            ],
            "x-enum-varnames": [
                "Ongoing",
                "Completed"
            ]
        },
        "model.Ride": {
            "type": "object",
            "properties": {
                "customerId": {
                    "type": "integer"
                },
                "distance": {
                    "type": "number"
                },
                "driverId": {
                    "type": "integer"
                },
                "endLat": {
                    "type": "number"
                },
                "endLon": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "rating": {
                    "type": "integer"
                },
                "startLat": {
                    "type": "number"
                },
                "startLon": {
                    "type": "number"
                },
                "status": {
                    "$ref": "#/definitions/enum_models.RideStatus"
                },
                "vehicleId": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "UberClientServer API",
	Description:      "This is an API to manage Users, Vehicles and Rides.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}