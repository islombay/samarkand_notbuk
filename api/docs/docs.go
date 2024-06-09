// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/v1/auth/access_token": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get new access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Get new access token",
                "operationId": "GetAccessToken",
                "responses": {}
            }
        },
        "/api/v1/auth/login_admin": {
            "post": {
                "description": "Login as admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login as admin",
                "operationId": "LoginAdmin",
                "parameters": [
                    {
                        "description": "Login request",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Login"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/brand": {
            "get": {
                "description": "Get brand list",
                "tags": [
                    "brand"
                ],
                "summary": "Get brand list",
                "operationId": "BrandGetList",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "q",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of brands",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update brand. Update only given values",
                "tags": [
                    "brand"
                ],
                "summary": "Update brand",
                "operationId": "BrandUpdate",
                "parameters": [
                    {
                        "description": "Update",
                        "name": "brand",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateBrand"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "brand",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "400": {
                        "description": "Bad id",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create brand",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "brand"
                ],
                "summary": "Create brand",
                "operationId": "BrandCreate",
                "parameters": [
                    {
                        "description": "Create brand request",
                        "name": "brand",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateBrand"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Brand",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "400": {
                        "description": "Bad request / Bad id",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "500": {
                        "description": "Internel server error",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete brand",
                "tags": [
                    "brand"
                ],
                "summary": "Delete brand",
                "operationId": "BrandDelete",
                "parameters": [
                    {
                        "description": "Delete",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.DeleteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "400": {
                        "description": "Bad request / Bad id",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    }
                }
            }
        },
        "/api/v1/brand/{id}": {
            "get": {
                "description": "Get brand by id",
                "tags": [
                    "brand"
                ],
                "summary": "Get brand by id",
                "operationId": "BrandGetByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "brand id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "brand",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "400": {
                        "description": "Invalid id",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    }
                }
            }
        },
        "/api/v1/category": {
            "get": {
                "description": "Get category list",
                "tags": [
                    "category"
                ],
                "summary": "Get category list",
                "operationId": "CategoryGetList",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "q",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of categories",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update category. Update only given values",
                "tags": [
                    "category"
                ],
                "summary": "Update category",
                "operationId": "CategoryUpdate",
                "parameters": [
                    {
                        "description": "Update",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateCategory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Category",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "400": {
                        "description": "Bad id",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create category",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Create category",
                "operationId": "CreateCategory",
                "parameters": [
                    {
                        "description": "Create category request",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateCategory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Category",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "400": {
                        "description": "Bad request / Bad id",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "500": {
                        "description": "Internel server error",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete category",
                "tags": [
                    "category"
                ],
                "summary": "Delete category",
                "operationId": "CategoryDelete",
                "parameters": [
                    {
                        "description": "Delete",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.DeleteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "400": {
                        "description": "Bad request / Bad id",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    }
                }
            }
        },
        "/api/v1/category/{id}": {
            "get": {
                "description": "Get category by id",
                "tags": [
                    "category"
                ],
                "summary": "Get category by id",
                "operationId": "CategoryGetByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "category id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Category",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "400": {
                        "description": "Invalid id",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/status.Status"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CreateBrand": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "models.CreateCategory": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 30
                },
                "parent_id": {
                    "type": "string"
                }
            }
        },
        "models.DeleteRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "models.Login": {
            "type": "object",
            "required": [
                "password",
                "phone_number"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "models.UpdateBrand": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.UpdateCategory": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 30
                },
                "parent_id": {
                    "type": "string"
                }
            }
        },
        "status.Status": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "count": {
                    "type": "integer"
                },
                "data": {},
                "error": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "message": {
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
