// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "xiayoushuang",
            "email": "york-xia@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/permission/print": {
            "get": {
                "description": "登录验证",
                "tags": [
                    "登录验证"
                ],
                "summary": "登录验证"
            }
        },
        "/api/v1/permission/user": {
            "post": {
                "description": "创建用户",
                "tags": [
                    "权限管理 - 管理员"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "description": "UserReq",
                        "name": "parameters",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.UserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "创建用户成功"
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "401": {
                        "description": "当前用户登录令牌失效",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "403": {
                        "description": "当前操作无权限",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "用户登录",
                "tags": [
                    "登录"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "LoginReq",
                        "name": "parameters",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "响应成功",
                        "schema": {
                            "$ref": "#/definitions/vo.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "401": {
                        "description": "当前用户登录令牌失效",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "403": {
                        "description": "当前操作无权限",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "get": {
                "description": "用户登出",
                "tags": [
                    "登录"
                ],
                "summary": "用户登出",
                "responses": {
                    "200": {
                        "description": "响应成功"
                    },
                    "400": {
                        "description": "请求参数错误",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "401": {
                        "description": "当前用户登录令牌失效",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "403": {
                        "description": "当前操作无权限",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "$ref": "#/definitions/vo.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "vo.Error": {
            "type": "object",
            "properties": {
                "args": {
                    "description": "参数",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "code": {
                    "description": "错误码",
                    "type": "integer"
                },
                "msg": {
                    "description": "错误消息",
                    "type": "string"
                }
            }
        },
        "vo.LoginReq": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "vo.LoginResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expiry": {
                    "type": "integer"
                },
                "token_type": {
                    "type": "string"
                }
            }
        },
        "vo.UserReq": {
            "type": "object",
            "properties": {
                "is_admin": {
                    "type": "boolean"
                },
                "password": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{"http", "https"},
	Title:       "JS流量统计管理系统后台API",
	Description: "JS流量统计管理系统后台API",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
