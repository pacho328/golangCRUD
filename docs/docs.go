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
            "name": "Francisco Anacona",
            "url": "http://artemisa.unicauca.edu.co/~javieranacona/index.html",
            "email": "pacho328@gmail.com"
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
        "/medidor": {
            "post": {
                "description": "Takes a Medidor JSON and store in DB postgres.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Medidor"
                ],
                "summary": "Create a new medidor",
                "parameters": [
                    {
                        "description": "Medidor JSON",
                        "name": "Medidor",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.MedidorSwCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.MedidorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Updates and returns a single Medidor whose Id value matches the id. New data must be passed in the body.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Medidor"
                ],
                "summary": "Update single Medidor by id",
                "parameters": [
                    {
                        "description": "update Medidor by id",
                        "name": "Medidor",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.MedidorSwUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.MedidorSwUpdateResponse"
                        }
                    }
                }
            }
        },
        "/medidor/{id}": {
            "get": {
                "description": "Returns the Medidor whose Id value matches the Id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Medidor"
                ],
                "summary": "Get single Medidor by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search Medidor by Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Medidor"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a single entry from the database based on id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Medidor"
                ],
                "summary": "Remove single Medidor by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "delete Medidor by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/medidores": {
            "get": {
                "description": "Responds with the list of all Medidores as JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Medidores"
                ],
                "summary": "Get Medidores array",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Medidor"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Medidor": {
            "type": "object",
            "required": [
                "address",
                "lines"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "brand": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "installationdate": {
                    "type": "string"
                },
                "isactive": {
                    "type": "boolean"
                },
                "lines": {
                    "type": "integer"
                },
                "retirementdate": {
                    "type": "string"
                },
                "serial": {
                    "type": "string"
                }
            }
        },
        "models.MedidorResponse": {
            "type": "object",
            "properties": {
                "Id": {
                    "type": "string",
                    "example": "6cd9f3c8-7bc8-40e7-8a4b-b575e63f0..."
                },
                "Message": {
                    "type": "string",
                    "example": "Medidor ..."
                }
            }
        },
        "models.MedidorSwCreate": {
            "type": "object",
            "required": [
                "lines"
            ],
            "properties": {
                "address": {
                    "type": "string",
                    "example": "Calle-street x"
                },
                "brand": {
                    "type": "string",
                    "example": "Marca x"
                },
                "installationdate": {
                    "type": "string",
                    "example": "2022-05-25T00:53:16.535668Z"
                },
                "isactive": {
                    "type": "boolean"
                },
                "lines": {
                    "description": "min: 1\nexample: 1",
                    "type": "integer"
                },
                "retirementdate": {
                    "type": "string",
                    "example": "2022-05-25T00:53:16.535668Z"
                },
                "serial": {
                    "type": "string",
                    "example": "Serial x"
                }
            }
        },
        "models.MedidorSwUpdate": {
            "type": "object",
            "required": [
                "lines"
            ],
            "properties": {
                "address": {
                    "type": "string",
                    "example": "Calle-street x"
                },
                "id": {
                    "type": "string",
                    "example": "6cd9f3c8-7bc8-40e7-8a4b-b575e63f0..."
                },
                "isactive": {
                    "type": "boolean"
                },
                "lines": {
                    "description": "min: 1\nexample: 1",
                    "type": "integer"
                },
                "retirementdate": {
                    "type": "string",
                    "example": "2022-05-25T00:53:16.535668Z"
                }
            }
        },
        "models.MedidorSwUpdateResponse": {
            "type": "object",
            "properties": {
                "Message": {
                    "type": "string",
                    "example": "Medidor ..."
                },
                "result": {
                    "$ref": "#/definitions/models.Medidor"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:5000",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Medidor Service",
	Description:      "A medidor management service API in Go using GORM.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}