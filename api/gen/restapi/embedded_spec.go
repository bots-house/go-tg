// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Birzzha",
    "version": "v1"
  },
  "host": "localhost:8000",
  "basePath": "/v1",
  "paths": {
    "/bot": {
      "get": {
        "security": [],
        "description": "Получение информации о боте\n",
        "tags": [
          "bot"
        ],
        "summary": "Get Bot Info",
        "operationId": "getBotInfo",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/BotInfo"
            }
          }
        }
      },
      "post": {
        "security": [],
        "description": "Обработка события от Telegram. S2S метод.\n",
        "tags": [
          "bot"
        ],
        "summary": "Handle Update",
        "operationId": "handleUpdate",
        "parameters": [
          {
            "name": "payload",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/TelegramUpdate"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "401": {
            "description": "Unauthorized"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    },
    "/token": {
      "post": {
        "security": [],
        "description": "Получение JWT-токена на основе данных авторизации от Telegram,\nполученных c [Telegram Login Widget](https://core.telegram.org/widgets/login) и [LoginUrl](https://core.telegram.org/bots/api#loginurl).\nТокен действителен в течении 24 часов, с момента создания.\n\nВозможные ошибки:\n  - ` + "`" + `telegram_widget_info_expired` + "`" + ` - если с момента получение данных с виджета прошло более 1 минуты;\n  - ` + "`" + `telegram_widget_info_invalid` + "`" + ` - если подпись данных (` + "`" + `hash` + "`" + `) невалидна;\n",
        "tags": [
          "auth"
        ],
        "summary": "Create Token",
        "operationId": "createToken",
        "parameters": [
          {
            "$ref": "#/parameters/RequestID"
          },
          {
            "name": "payload",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/InputAuth"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/Token"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/user": {
      "get": {
        "description": "Возвращает информацию о текущем пользователе.",
        "tags": [
          "auth"
        ],
        "summary": "Get Current User",
        "operationId": "getUser",
        "parameters": [
          {
            "$ref": "#/parameters/RequestID"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "BotInfo": {
      "description": "Информация о боте",
      "required": [
        "name",
        "username",
        "auth_deep_link"
      ],
      "properties": {
        "auth_deep_link": {
          "description": "Значение параметра ` + "`" + `?start=` + "`" + ` который нужно передать при формировании URL для авторизации через ` + "`" + `LoginUrl` + "`" + `.",
          "type": "string",
          "x-order": 2
        },
        "name": {
          "description": "Имя бота",
          "type": "string",
          "x-order": 0
        },
        "username": {
          "description": "Юзернейм бота",
          "type": "string",
          "x-order": 1
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "code",
        "description"
      ],
      "properties": {
        "code": {
          "description": "Код ошибки",
          "type": "string",
          "enum": [
            "test"
          ],
          "x-order": 0
        },
        "description": {
          "description": "Описание ошибки",
          "type": "string",
          "x-order": 1
        }
      }
    },
    "Identity": {
      "x-go-type": {
        "import": {
          "alias": "authz",
          "package": "github.com/bots-house/birzzha/api/authz"
        },
        "type": "Identity"
      }
    },
    "InputAuth": {
      "description": "Данный полученные с [Telegram Login Widget](https://core.telegram.org/widgets/login) и [LoginUrl](https://core.telegram.org/bots/api#loginurl).",
      "required": [
        "id",
        "first_name",
        "auth_date",
        "hash"
      ],
      "properties": {
        "auth_date": {
          "description": "Дата авторизации пользователя (UNIX)",
          "type": "integer",
          "format": "int64",
          "x-order": 5
        },
        "first_name": {
          "description": "Имя пользователя в Telegram",
          "type": "string",
          "x-order": 1,
          "example": "Sasha"
        },
        "hash": {
          "description": "Подписи к данным.",
          "type": "string",
          "x-order": 6
        },
        "id": {
          "description": "ID пользователя в Telegram",
          "type": "integer",
          "x-order": 0,
          "example": 492933123
        },
        "last_name": {
          "description": "Фамилия пользователя в Telegram",
          "type": "string",
          "x-order": 2,
          "example": "Savchuk"
        },
        "photo_url": {
          "description": "Ссылка на авартку пользователя.",
          "type": "string",
          "format": "url",
          "x-order": 4
        },
        "username": {
          "description": "Username пользователя в Telegram",
          "type": "string",
          "x-order": 3
        }
      }
    },
    "TelegramUpdate": {
      "x-go-type": {
        "import": {
          "alias": "tgbotapi",
          "package": "github.com/go-telegram-bot-api/telegram-bot-api"
        },
        "type": "Update"
      }
    },
    "Token": {
      "description": "Токен для доступа к API.",
      "required": [
        "token",
        "user"
      ],
      "properties": {
        "token": {
          "description": "JWT-токен для доступа к API.",
          "type": "string",
          "x-order": 0
        },
        "user": {
          "x-order": 1,
          "$ref": "#/definitions/User"
        }
      }
    },
    "User": {
      "description": "Объект пользователя",
      "type": "object",
      "required": [
        "id",
        "telegram",
        "first_name",
        "last_name",
        "avatar",
        "is_admin",
        "joined_at"
      ],
      "properties": {
        "avatar": {
          "description": "Path to avatar",
          "type": "string",
          "x-order": 4
        },
        "first_name": {
          "description": "Имя пользователя в Telegram",
          "type": "string",
          "x-order": 2
        },
        "id": {
          "description": "Уникальный ID пользователя в Birzzha.",
          "type": "integer",
          "x-order": 0
        },
        "is_admin": {
          "description": "True, если пользователь админ Birzzha.",
          "type": "boolean",
          "x-order": 5
        },
        "joined_at": {
          "description": "Дата и время регистрации на бирже, в Unix-time.",
          "type": "integer",
          "x-order": 6
        },
        "last_name": {
          "description": "Фамилия пользователя в Telegram (может быть ` + "`" + `null` + "`" + `)",
          "type": "string",
          "x-order": 3
        },
        "telegram": {
          "description": "Информация о пользователе из Telegram",
          "type": "object",
          "required": [
            "id",
            "username"
          ],
          "properties": {
            "id": {
              "description": "ID пользователя в Telegram",
              "type": "integer",
              "x-order": 0
            },
            "username": {
              "description": "Username пользователя в Telegram",
              "type": "string",
              "x-order": 1
            }
          },
          "x-order": 1
        }
      }
    }
  },
  "parameters": {
    "RequestID": {
      "type": "string",
      "format": "uuid",
      "description": "Уникальный ID запроса. Используется для трейсинга.",
      "name": "X-Request-Id",
      "in": "header"
    }
  },
  "securityDefinitions": {
    "TokenHeader": {
      "type": "apiKey",
      "name": "X-Token",
      "in": "header"
    },
    "TokenQuery": {
      "type": "apiKey",
      "name": "token",
      "in": "query"
    }
  },
  "security": [
    {
      "TokenHeader": []
    },
    {
      "TokenQuery": []
    }
  ],
  "tags": [
    {
      "description": "Авторизация пользователя.\nТокен может быть передан через заголовок (` + "`" + `X-Token` + "`" + `) или query параметр (` + "`" + `token` + "`" + `).\nЕсли токен невалидный API вернет ответ  **403 Forbidden** или **401 Unauthorized**.\nТокены живут в течении 24 часов.\nВремя на передачу данных с виджета 1 минута.\n",
      "name": "auth"
    },
    {
      "name": "bot"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Birzzha",
    "version": "v1"
  },
  "host": "localhost:8000",
  "basePath": "/v1",
  "paths": {
    "/bot": {
      "get": {
        "security": [],
        "description": "Получение информации о боте\n",
        "tags": [
          "bot"
        ],
        "summary": "Get Bot Info",
        "operationId": "getBotInfo",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/BotInfo"
            }
          }
        }
      },
      "post": {
        "security": [],
        "description": "Обработка события от Telegram. S2S метод.\n",
        "tags": [
          "bot"
        ],
        "summary": "Handle Update",
        "operationId": "handleUpdate",
        "parameters": [
          {
            "name": "payload",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/TelegramUpdate"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "401": {
            "description": "Unauthorized"
          },
          "500": {
            "description": "Internal Server Error"
          }
        }
      }
    },
    "/token": {
      "post": {
        "security": [],
        "description": "Получение JWT-токена на основе данных авторизации от Telegram,\nполученных c [Telegram Login Widget](https://core.telegram.org/widgets/login) и [LoginUrl](https://core.telegram.org/bots/api#loginurl).\nТокен действителен в течении 24 часов, с момента создания.\n\nВозможные ошибки:\n  - ` + "`" + `telegram_widget_info_expired` + "`" + ` - если с момента получение данных с виджета прошло более 1 минуты;\n  - ` + "`" + `telegram_widget_info_invalid` + "`" + ` - если подпись данных (` + "`" + `hash` + "`" + `) невалидна;\n",
        "tags": [
          "auth"
        ],
        "summary": "Create Token",
        "operationId": "createToken",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Уникальный ID запроса. Используется для трейсинга.",
            "name": "X-Request-Id",
            "in": "header"
          },
          {
            "name": "payload",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/InputAuth"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/Token"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/user": {
      "get": {
        "description": "Возвращает информацию о текущем пользователе.",
        "tags": [
          "auth"
        ],
        "summary": "Get Current User",
        "operationId": "getUser",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "Уникальный ID запроса. Используется для трейсинга.",
            "name": "X-Request-Id",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "BotInfo": {
      "description": "Информация о боте",
      "required": [
        "name",
        "username",
        "auth_deep_link"
      ],
      "properties": {
        "auth_deep_link": {
          "description": "Значение параметра ` + "`" + `?start=` + "`" + ` который нужно передать при формировании URL для авторизации через ` + "`" + `LoginUrl` + "`" + `.",
          "type": "string",
          "x-order": 2
        },
        "name": {
          "description": "Имя бота",
          "type": "string",
          "x-order": 0
        },
        "username": {
          "description": "Юзернейм бота",
          "type": "string",
          "x-order": 1
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "code",
        "description"
      ],
      "properties": {
        "code": {
          "description": "Код ошибки",
          "type": "string",
          "enum": [
            "test"
          ],
          "x-order": 0
        },
        "description": {
          "description": "Описание ошибки",
          "type": "string",
          "x-order": 1
        }
      }
    },
    "Identity": {
      "x-go-type": {
        "import": {
          "alias": "authz",
          "package": "github.com/bots-house/birzzha/api/authz"
        },
        "type": "Identity"
      }
    },
    "InputAuth": {
      "description": "Данный полученные с [Telegram Login Widget](https://core.telegram.org/widgets/login) и [LoginUrl](https://core.telegram.org/bots/api#loginurl).",
      "required": [
        "id",
        "first_name",
        "auth_date",
        "hash"
      ],
      "properties": {
        "auth_date": {
          "description": "Дата авторизации пользователя (UNIX)",
          "type": "integer",
          "format": "int64",
          "x-order": 5
        },
        "first_name": {
          "description": "Имя пользователя в Telegram",
          "type": "string",
          "x-order": 1,
          "example": "Sasha"
        },
        "hash": {
          "description": "Подписи к данным.",
          "type": "string",
          "x-order": 6
        },
        "id": {
          "description": "ID пользователя в Telegram",
          "type": "integer",
          "x-order": 0,
          "example": 492933123
        },
        "last_name": {
          "description": "Фамилия пользователя в Telegram",
          "type": "string",
          "x-order": 2,
          "example": "Savchuk"
        },
        "photo_url": {
          "description": "Ссылка на авартку пользователя.",
          "type": "string",
          "format": "url",
          "x-order": 4
        },
        "username": {
          "description": "Username пользователя в Telegram",
          "type": "string",
          "x-order": 3
        }
      }
    },
    "TelegramUpdate": {
      "x-go-type": {
        "import": {
          "alias": "tgbotapi",
          "package": "github.com/go-telegram-bot-api/telegram-bot-api"
        },
        "type": "Update"
      }
    },
    "Token": {
      "description": "Токен для доступа к API.",
      "required": [
        "token",
        "user"
      ],
      "properties": {
        "token": {
          "description": "JWT-токен для доступа к API.",
          "type": "string",
          "x-order": 0
        },
        "user": {
          "x-order": 1,
          "$ref": "#/definitions/User"
        }
      }
    },
    "User": {
      "description": "Объект пользователя",
      "type": "object",
      "required": [
        "id",
        "telegram",
        "first_name",
        "last_name",
        "avatar",
        "is_admin",
        "joined_at"
      ],
      "properties": {
        "avatar": {
          "description": "Path to avatar",
          "type": "string",
          "x-order": 4
        },
        "first_name": {
          "description": "Имя пользователя в Telegram",
          "type": "string",
          "x-order": 2
        },
        "id": {
          "description": "Уникальный ID пользователя в Birzzha.",
          "type": "integer",
          "x-order": 0
        },
        "is_admin": {
          "description": "True, если пользователь админ Birzzha.",
          "type": "boolean",
          "x-order": 5
        },
        "joined_at": {
          "description": "Дата и время регистрации на бирже, в Unix-time.",
          "type": "integer",
          "x-order": 6
        },
        "last_name": {
          "description": "Фамилия пользователя в Telegram (может быть ` + "`" + `null` + "`" + `)",
          "type": "string",
          "x-order": 3
        },
        "telegram": {
          "description": "Информация о пользователе из Telegram",
          "type": "object",
          "required": [
            "id",
            "username"
          ],
          "properties": {
            "id": {
              "description": "ID пользователя в Telegram",
              "type": "integer",
              "x-order": 0
            },
            "username": {
              "description": "Username пользователя в Telegram",
              "type": "string",
              "x-order": 1
            }
          },
          "x-order": 1
        }
      }
    },
    "UserTelegram": {
      "description": "Информация о пользователе из Telegram",
      "type": "object",
      "required": [
        "id",
        "username"
      ],
      "properties": {
        "id": {
          "description": "ID пользователя в Telegram",
          "type": "integer",
          "x-order": 0
        },
        "username": {
          "description": "Username пользователя в Telegram",
          "type": "string",
          "x-order": 1
        }
      },
      "x-order": 1
    }
  },
  "parameters": {
    "RequestID": {
      "type": "string",
      "format": "uuid",
      "description": "Уникальный ID запроса. Используется для трейсинга.",
      "name": "X-Request-Id",
      "in": "header"
    }
  },
  "securityDefinitions": {
    "TokenHeader": {
      "type": "apiKey",
      "name": "X-Token",
      "in": "header"
    },
    "TokenQuery": {
      "type": "apiKey",
      "name": "token",
      "in": "query"
    }
  },
  "security": [
    {
      "TokenHeader": []
    },
    {
      "TokenQuery": []
    }
  ],
  "tags": [
    {
      "description": "Авторизация пользователя.\nТокен может быть передан через заголовок (` + "`" + `X-Token` + "`" + `) или query параметр (` + "`" + `token` + "`" + `).\nЕсли токен невалидный API вернет ответ  **403 Forbidden** или **401 Unauthorized**.\nТокены живут в течении 24 часов.\nВремя на передачу данных с виджета 1 минута.\n",
      "name": "auth"
    },
    {
      "name": "bot"
    }
  ]
}`))
}
