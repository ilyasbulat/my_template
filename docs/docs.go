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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/translation/do-translate": {
            "post": {
                "description": "Translate a text",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "translation"
                ],
                "summary": "Translate",
                "operationId": "do-translate",
                "parameters": [
                    {
                        "description": "Set up translation",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.doTranslateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Translation"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/apperror.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apperror.AppError"
                        }
                    }
                }
            }
        },
        "/translation/history": {
            "get": {
                "description": "Show all translation history",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "translation"
                ],
                "summary": "Show history",
                "operationId": "history",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.historyResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/apperror.AppError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "apperror.AppError": {
            "type": "object",
            "properties": {
                "developer_message": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "entity.Translation": {
            "type": "object",
            "properties": {
                "destination": {
                    "type": "string",
                    "example": "en"
                },
                "original": {
                    "type": "string",
                    "example": "текст для перевода"
                },
                "source": {
                    "type": "string",
                    "example": "auto"
                },
                "translation": {
                    "type": "string",
                    "example": "text for translation"
                }
            }
        },
        "v1.doTranslateRequest": {
            "type": "object",
            "required": [
                "destination",
                "original",
                "source"
            ],
            "properties": {
                "destination": {
                    "type": "string",
                    "example": "en"
                },
                "original": {
                    "type": "string",
                    "example": "текст для перевода"
                },
                "source": {
                    "type": "string",
                    "example": "auto"
                }
            }
        },
        "v1.historyResponse": {
            "type": "object",
            "properties": {
                "history": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Translation"
                    }
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
