// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-10-12 20:02:18.474818 +0300 MSK m=+51.574849200

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
        "/ping": {
            "get": {
                "description": "ping",
                "summary": "Ping",
                "responses": {
                    "200": {},
                    "400": {}
                }
            }
        },
        "/tasks/{user_id}": {
            "get": {
                "description": "Retrieve list of tasks.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get list of tasks.",
                "operationId": "get-tasks",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
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
                                "$ref": "#/definitions/model.Task"
                            }
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Retrieve a user info.",
                "produces": [
                    "application/json"
                ],
                "summary": "Get user.",
                "operationId": "get-user",
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
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {}
                }
            }
        }
    },
    "definitions": {
        "model.Achievement": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "expirationDate": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "model.Task": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "expirationDate": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "managerID": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "score": {
                    "type": "integer"
                },
                "taskPriority": {
                    "type": "object",
                    "$ref": "#/definitions/model.TaskPriority"
                },
                "taskPriorityID": {
                    "type": "integer"
                },
                "taskStatus": {
                    "type": "object",
                    "$ref": "#/definitions/model.TaskStatus"
                },
                "taskStatusID": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "model.TaskPriority": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "priority": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.TaskStatus": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "achievement": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Achievement"
                    }
                },
                "balance": {
                    "type": "number"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "department": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                },
                "skill": {
                    "type": "object",
                    "$ref": "#/definitions/postgres.Jsonb"
                },
                "tasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Task"
                    }
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "postgres.Jsonb": {
            "type": "object"
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
