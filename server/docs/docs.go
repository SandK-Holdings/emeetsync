// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "/analytics/scanned-poster": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "analytics"
                ],
                "summary": "Notifies us when poster QR code has been scanned",
                "parameters": [
                    {
                        "description": "Object containing the location that poster was scanned from and the url that was scanned",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "allOf": [
                                {
                                    "type": "object"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "location": {
                                            "$ref": "#/definitions/models.Location"
                                        },
                                        "url": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "Signs user in and sets the access token session variable",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Signs user in",
                "parameters": [
                    {
                        "description": "Object containing the Google authorization code and the user's timezone offset",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "allOf": [
                                {
                                    "type": "object"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "string"
                                        },
                                        "timezoneOffset": {
                                            "type": "integer"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/auth/sign-in-mobile": {
            "post": {
                "description": "Signs user in and sets the access token session variable",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Signs user in from mobile",
                "parameters": [
                    {
                        "description": "Object containing the Google authorization code and the user's timezone offset",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "allOf": [
                                {
                                    "type": "object"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "accessToken": {
                                            "type": "string"
                                        },
                                        "expiresIn": {
                                            "type": "integer"
                                        },
                                        "idToken": {
                                            "type": "string"
                                        },
                                        "refreshToken": {
                                            "type": "string"
                                        },
                                        "scope": {
                                            "type": "string"
                                        },
                                        "timezoneOffset": {
                                            "type": "integer"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/auth/sign-out": {
            "post": {
                "description": "Signs user out and deletes the session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Signs user out",
                "responses": {
                    "200": {}
                }
            }
        },
        "/auth/status": {
            "get": {
                "description": "Returns a 401 error if user is not signed in, 200 if they are",
                "tags": [
                    "auth"
                ],
                "summary": "Gets whether the user is signed in or not",
                "responses": {
                    "200": {},
                    "401": {
                        "description": "Error object",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            }
        },
        "/events": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Creates a new event",
                "parameters": [
                    {
                        "description": "Object containing info about the event to create",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "allOf": [
                                {
                                    "type": "object"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "dates": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
                                        },
                                        "duration": {
                                            "type": "number"
                                        },
                                        "name": {
                                            "type": "string"
                                        },
                                        "notificationsEnabled": {
                                            "type": "boolean"
                                        },
                                        "type": {
                                            "$ref": "#/definitions/models.EventType"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "type": "object"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "eventId": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/events/{eventId}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Gets an event based on its id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "eventId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Event"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Edits an event based on its id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "eventId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Object containing info about the event to update",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "allOf": [
                                {
                                    "type": "object"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "dates": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
                                        },
                                        "duration": {
                                            "type": "number"
                                        },
                                        "name": {
                                            "type": "string"
                                        },
                                        "notificationsEnabled": {
                                            "type": "boolean"
                                        },
                                        "type": {
                                            "$ref": "#/definitions/models.EventType"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Deletes an event based on its id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "eventId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/events/{eventId}/response": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Updates the current user's availability",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "eventId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Object containing info about the event response to update",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "allOf": [
                                {
                                    "type": "object"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "availability": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
                                        },
                                        "guest": {
                                            "type": "boolean"
                                        },
                                        "name": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/user": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Deletes the currently signed in user",
                "responses": {
                    "200": {}
                }
            }
        },
        "/user/add-calendar-account": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Adds a new calendar account",
                "parameters": [
                    {
                        "description": "Object containing the Google authorization code",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "allOf": [
                                {
                                    "type": "object"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/user/calendars": {
            "get": {
                "description": "Gets the user's calendar events between \"timeMin\" and \"timeMax\"",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Gets the user's calendar events",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Lower bound for event's start time to filter by",
                        "name": "timeMin",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Upper bound for event's end time to filter by",
                        "name": "timeMax",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Comma separated list of accounts to fetch calendar events from",
                        "name": "accounts",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "$ref": "#/definitions/calendar.CalendarEventsWithError"
                            }
                        }
                    }
                }
            }
        },
        "/user/events": {
            "get": {
                "description": "Returns an array containing all the user's events",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Gets all the user's events",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "type": "object"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "events": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Event"
                                            }
                                        },
                                        "joinedEvents": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Event"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/profile": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Gets the user's profile",
                "responses": {
                    "200": {
                        "description": "A user profile object",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/user/remove-calendar-account": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Removes an existing calendar account",
                "parameters": [
                    {
                        "description": "Object containing the email of the calendar account to remove",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "allOf": [
                                {
                                    "type": "object"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "email": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/user/searchContacts": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Searches the user's contacts based on the given query",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Query to search for",
                        "name": "query",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    }
                }
            }
        },
        "/user/toggle-calendar": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Toggles whether the specified calendar is enabled or disabled for the user",
                "parameters": [
                    {
                        "description": "Email of calendar account and whether to enable it",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "allOf": [
                                {
                                    "type": "object"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "email": {
                                            "type": "string"
                                        },
                                        "enabled": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/users": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Returns users that match the search query",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search query matching users' names/emails",
                        "name": "query",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "An array of user profile objects",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "calendar.CalendarEventsWithError": {
            "type": "object",
            "properties": {
                "calendarEvents": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CalendarEvent"
                    }
                },
                "error": {
                    "type": "object",
                    "$ref": "#/definitions/errs.GoogleAPIError"
                }
            }
        },
        "errs.GoogleAPIError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "details": {
                    "type": "object"
                },
                "errors": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "models.CalendarAccount": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "enabled": {
                    "type": "boolean"
                },
                "picture": {
                    "type": "string"
                }
            }
        },
        "models.CalendarEvent": {
            "type": "object",
            "properties": {
                "endDate": {
                    "type": "string"
                },
                "startDate": {
                    "type": "string"
                },
                "summary": {
                    "type": "string"
                }
            }
        },
        "models.Event": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "calendarEventId": {
                    "type": "string"
                },
                "dates": {
                    "type": "string"
                },
                "duration": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "notificationsEnabled": {
                    "type": "boolean"
                },
                "ownerId": {
                    "type": "string"
                },
                "responses": {
                    "description": "Availability responses",
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/models.Response"
                    }
                },
                "scheduledEvent": {
                    "description": "Scheduled event",
                    "type": "object",
                    "$ref": "#/definitions/models.CalendarEvent"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.EventType": {
            "type": "string"
        },
        "models.Location": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "country_code": {
                    "type": "string"
                },
                "country_name": {
                    "type": "string"
                },
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "postal": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                }
            }
        },
        "models.Response": {
            "type": "object",
            "properties": {
                "availability": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "user": {
                    "type": "object",
                    "$ref": "#/definitions/models.User"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "_id": {
                    "description": "Profile info",
                    "type": "string"
                },
                "calendarAccounts": {
                    "description": "CalendarAccounts contains all the additional accounts the user wants to see google calendar events for",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CalendarAccount"
                    }
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "timezoneOffset": {
                    "type": "integer"
                }
            }
        },
        "responses.Error": {
            "type": "object",
            "required": [
                "error"
            ],
            "properties": {
                "error": {
                    "type": "object"
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
	Host:        "localhost:3002",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Schej.it API",
	Description: "This is the API for Schej.it!",
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
