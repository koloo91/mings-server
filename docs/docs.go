// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-05-11 14:39:36.674142 +0200 CEST m=+0.022519763

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/documents": {
            "get": {
                "description": "Get all stored documents",
                "produces": [
                    "application/json"
                ],
                "summary": "Get all stored documents",
                "operationId": "get-documents",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Documents"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ApiError"
                        }
                    }
                }
            },
            "post": {
                "description": "Uploads a single document",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Uploads a single document",
                "operationId": "upload-document",
                "parameters": [
                    {
                        "type": "file",
                        "description": "document file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Document"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ApiError"
                        }
                    }
                }
            }
        },
        "/documents/{id}": {
            "get": {
                "description": "Get all stored documents",
                "produces": [
                    "application/json"
                ],
                "summary": "Get all stored documents",
                "operationId": "get-document-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "document id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Document"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ApiError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ApiError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "model.DependsOn": {
            "type": "object",
            "properties": {
                "external": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.DependsOnService"
                    }
                },
                "internal": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.DependsOnService"
                    }
                }
            }
        },
        "model.DependsOnService": {
            "type": "object",
            "properties": {
                "serviceName": {
                    "type": "string"
                },
                "why": {
                    "type": "string"
                }
            }
        },
        "model.Document": {
            "type": "object",
            "properties": {
                "contact": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "links": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "owner": {
                    "type": "string"
                },
                "service": {
                    "type": "object",
                    "$ref": "#/definitions/model.Service"
                },
                "shortName": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.Documents": {
            "type": "object",
            "properties": {
                "documents": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Document"
                    }
                }
            }
        },
        "model.Provide": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                },
                "protocol": {
                    "type": "string"
                },
                "serviceName": {
                    "type": "string"
                },
                "transportProtocol": {
                    "type": "string"
                }
            }
        },
        "model.Service": {
            "type": "object",
            "properties": {
                "dependsOn": {
                    "type": "object",
                    "$ref": "#/definitions/model.DependsOn"
                },
                "provides": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Provide"
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
