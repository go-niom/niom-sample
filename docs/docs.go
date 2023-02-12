// Code generated by swaggo/swag. DO NOT EDIT
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
        "/v1/app": {
            "get": {
                "security": [
                    {
                        "ApiKeyApp": []
                    }
                ],
                "description": "get all app",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "App"
                ],
                "summary": "Get all app",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.ResMessage"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyApp": []
                    }
                ],
                "description": "create app",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "App"
                ],
                "summary": "create a app",
                "parameters": [
                    {
                        "description": "Update App Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.CreateAppDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.ResMessage"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "ApiKeyApp": []
                    }
                ],
                "description": "update app",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "App"
                ],
                "summary": "Update App by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "App ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Create App Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.UpdateAppDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.ResMessage"
                        }
                    }
                }
            }
        },
        "/v1/app/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyApp": []
                    }
                ],
                "description": "get app",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "App"
                ],
                "summary": "get a app",
                "parameters": [
                    {
                        "type": "string",
                        "description": "App ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helpers.ResMessage"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyApp": []
                    }
                ],
                "description": "delete app",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "App"
                ],
                "summary": "delete a app",
                "parameters": [
                    {
                        "type": "string",
                        "description": "App ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "app.CreateAppDto": {
            "type": "object",
            "properties": {
                "sample": {
                    "type": "string"
                }
            }
        },
        "app.UpdateAppDto": {
            "type": "object",
            "properties": {
                "sample": {
                    "type": "string"
                }
            }
        },
        "helpers.ResMessage": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error": {},
                "message": {
                    "type": "string"
                },
                "result": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:7000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Niom-Sample",
	Description:      "Niom-Sample Backend REST API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
