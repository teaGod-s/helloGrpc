package docs

const (
	swagger = `{
  "swagger": "2.0",
  "info": {
    "title": "hello服务的api文档",
    "description": "描述待会再写",
    "version": "1.0"
  },
  "tags": [
    {
      "name": "HelloService"
    },
    {
      "name": "AuthService"
    }
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/authipb.AuthService/ListUserGroupByRoleID": {
      "post": {
        "summary": "ListUserGroupByRoleID 根据角色ID获取他所关联的的用户组的列表",
        "operationId": "AuthService_ListUserGroupByRoleID",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/authipbListUserGroupByRoleIDResp"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/authipbListUserGroupByRoleIDReq"
            }
          }
        ],
        "tags": [
          "AuthService"
        ]
      }
    },
    "/hello.HelloService/SayHello": {
      "post": {
        "summary": "打招呼",
        "operationId": "HelloService_SayHello",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/helloSayHelloResponse"
            }
          },
          "default": {
            "description": "An unexpected error response.",
            "schema": {
              "$ref": "#/definitions/googlerpcStatus"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/helloSayHelloRequest"
            }
          }
        ],
        "tags": [
          "HelloService"
        ]
      }
    }
  },
  "definitions": {
    "authipbListUserGroupByRoleIDReq": {
      "type": "object",
      "properties": {
        "CommonIn": {
          "$ref": "#/definitions/xcommonCommonIn"
        },
        "RoleId": {
          "type": "string",
          "format": "int64",
          "title": "角色ID"
        }
      },
      "title": "ListUserGroupByRoleIDReq 列出一个角色下面所有的用户组列表请求包"
    },
    "authipbListUserGroupByRoleIDResp": {
      "type": "object",
      "properties": {
        "CommonOut": {
          "$ref": "#/definitions/xcommonCommonOut"
        },
        "UserGroupInfos": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/xcommonMarkMeta"
          }
        }
      },
      "title": "ListUserGroupByRoleIDResp 列出一个角色下面所有的用户组列表回包"
    },
    "googlerpcStatus": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "details": {
          "type": "array",
          "items": {
            "type": "object",
            "$ref": "#/definitions/protobufAny"
          }
        }
      }
    },
    "helloHello": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "age": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "helloSayHelloRequest": {
      "type": "object",
      "properties": {
        "msg": {
          "$ref": "#/definitions/helloHello"
        }
      }
    },
    "helloSayHelloResponse": {
      "type": "object",
      "properties": {
        "status": {
          "$ref": "#/definitions/xcommonStatus"
        }
      }
    },
    "protobufAny": {
      "type": "object",
      "properties": {
        "@type": {
          "type": "string"
        }
      },
      "additionalProperties": {}
    },
    "xcommonCommonIn": {
      "type": "object",
      "properties": {
        "TrackId": {
          "type": "string",
          "title": "追踪Id"
        },
        "ClientId": {
          "type": "string",
          "title": "客户端Id"
        },
        "Lang": {
          "type": "string",
          "title": "语言"
        },
        "Token": {
          "type": "string",
          "title": "令牌, 需要登录的接口必传"
        },
        "UserId": {
          "type": "string",
          "format": "int64",
          "title": "用户Id, 可选字段, 客户端请求不传"
        }
      },
      "title": "CommonIn 通用请求包字段"
    },
    "xcommonCommonOut": {
      "type": "object",
      "properties": {
        "Code": {
          "type": "integer",
          "format": "int32",
          "title": "状态码"
        },
        "ErrMsg": {
          "type": "string",
          "title": "错误信息"
        }
      },
      "title": "CommonOut 通用回包字段"
    },
    "xcommonMark": {
      "type": "integer",
      "format": "int32",
      "enum": [
        0,
        1,
        2
      ],
      "default": 0,
      "description": "- 1: 是选择状态\n - 2: 不是选择状态",
      "title": "选择框是否为选择状态的枚举定义"
    },
    "xcommonMarkMeta": {
      "type": "object",
      "properties": {
        "Id": {
          "type": "string",
          "format": "int64"
        },
        "Name": {
          "type": "string"
        },
        "Mark": {
          "$ref": "#/definitions/xcommonMark"
        }
      },
      "title": "MarkMeta is a meta message with an addition mark field"
    },
    "xcommonStatus": {
      "type": "integer",
      "format": "int32",
      "enum": [
        0,
        1,
        2
      ],
      "default": 0,
      "title": "- 1: 正常\n - 2: 非正常"
    }
  }
}
`
)
