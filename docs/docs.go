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
        "/csv-to-mongo": {
            "post": {
                "description": "After each time that this endpoint gets called, all of the csv data will be imported to the database.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ImportData"
                ],
                "summary": "import csv data",
                "responses": {}
            }
        },
        "/series": {
            "get": {
                "description": "Retrieve list of all the series from series collection, important note: swagger might be slow to represent large amount of data, use postman for better experience",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Series"
                ],
                "summary": "all the series (for better experience, use postman on this endpoint)",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Series"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new series using the provided fields.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Series"
                ],
                "summary": "create new series",
                "parameters": [
                    {
                        "description": "Series object",
                        "name": "series",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.SwaggerParams"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "This endpoint will wipe out the database, and deletes all of the series",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Series"
                ],
                "summary": "delete all of the series",
                "responses": {}
            }
        },
        "/series/{id}": {
            "get": {
                "description": "Get detail of the single series object",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Series"
                ],
                "summary": "single series",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Series ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Series"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a single series object with the given id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Series"
                ],
                "summary": "delete a single series object",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Series ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "patch": {
                "description": "Change any field(except id) of any series",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Series"
                ],
                "summary": "update the series fields",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Series ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Series object",
                        "name": "series",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.SwaggerParams"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "main.Series": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "data_value": {
                    "type": "number"
                },
                "group": {
                    "type": "string"
                },
                "magnitude": {
                    "type": "integer"
                },
                "period": {
                    "type": "string"
                },
                "reference": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "subject": {
                    "type": "string"
                },
                "suppressed": {
                    "type": "boolean"
                },
                "titles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "units": {
                    "type": "string"
                }
            }
        },
        "main.SwaggerParams": {
            "type": "object",
            "properties": {
                "data_value": {
                    "type": "number"
                },
                "group": {
                    "type": "string"
                },
                "magnitude": {
                    "type": "integer"
                },
                "period": {
                    "type": "string"
                },
                "reference": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "subject": {
                    "type": "string"
                },
                "suppressed": {
                    "type": "boolean"
                },
                "titles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "units": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
