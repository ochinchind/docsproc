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
        "/google_callback": {
            "get": {
                "description": "Handles Google's OAuth callback, processes authentication, and returns an access token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GoogleOAuth"
                ],
                "summary": "Google OAuth Callback",
                "operationId": "google_callback",
                "responses": {
                    "200": {
                        "description": "Successfully authenticated",
                        "schema": {
                            "$ref": "#/definitions/v1.googleCallbackResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/google_login": {
            "get": {
                "description": "Redirects the user to Google's OAuth authentication page.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GoogleOAuth"
                ],
                "summary": "Google Login",
                "operationId": "google_login",
                "responses": {
                    "303": {
                        "description": "Redirect to Google login"
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Authenticates a user and returns an access token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "User Login",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "Login request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully authenticated",
                        "schema": {
                            "$ref": "#/definitions/v1.loginResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid credentials",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "406": {
                        "description": "Invalid request payload",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Registers a new user and returns an access token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "User Registration",
                "operationId": "register",
                "parameters": [
                    {
                        "description": "Register request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully registered",
                        "schema": {
                            "$ref": "#/definitions/v1.loginResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "406": {
                        "description": "Invalid request payload",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/v1/users": {
            "get": {
                "description": "Fetch a paginated list of users.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get Users",
                "operationId": "get_users",
                "responses": {
                    "200": {
                        "description": "Successful response with user list",
                        "schema": {
                            "$ref": "#/definitions/v1.getUsersResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/v1/users/{id}": {
            "get": {
                "description": "Fetch details of a user by their ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get User",
                "operationId": "get_user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response with user details",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update details of an existing user by providing user ID and update data.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update User",
                "operationId": "update_user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User update data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateUserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User successfully updated",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.LoginDTO": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "username": {
                    "type": "string",
                    "example": "john.doe"
                }
            }
        },
        "dto.RegisterDTO": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "abc@example.com"
                },
                "name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "phone": {
                    "type": "string",
                    "example": "+1234567890"
                },
                "role": {
                    "type": "string",
                    "enum": [
                        "admin",
                        "user",
                        "methodologist"
                    ],
                    "example": "user"
                },
                "surname": {
                    "type": "string",
                    "example": "Doe"
                },
                "username": {
                    "type": "string",
                    "example": "john.doe"
                }
            }
        },
        "dto.UpdateUserDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "abc@example.com"
                },
                "name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "phone": {
                    "type": "string",
                    "example": "+1234567890"
                },
                "role": {
                    "type": "string",
                    "enum": [
                        "admin",
                        "user",
                        "methodologist"
                    ],
                    "example": "user"
                },
                "surname": {
                    "type": "string",
                    "example": "Doe"
                },
                "username": {
                    "type": "string",
                    "example": "john.doe"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "properties": {
                "api_token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
                },
                "created_at": {
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z"
                },
                "email": {
                    "type": "string",
                    "example": "abc@example.com"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "John"
                },
                "phone": {
                    "type": "string",
                    "example": "877755544434"
                },
                "picture": {
                    "type": "string",
                    "example": "https://example.com/picture.jpg"
                },
                "role": {
                    "type": "string",
                    "example": "admin"
                },
                "surname": {
                    "type": "string",
                    "example": "Doe"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z"
                },
                "username": {
                    "type": "string",
                    "example": "johndoe"
                }
            }
        },
        "v1.getUsersResponse": {
            "type": "object",
            "properties": {
                "total": {
                    "type": "integer"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.User"
                    }
                }
            }
        },
        "v1.googleCallbackResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "v1.loginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "v1.response": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "message"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Go Clean Template API",
	Description:      "Using a translation service as an example",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
