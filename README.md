# go-commerce
An e-commerce microservice written with golang. It provides APIs for the following:


- User creation and Mangement
- Store creation and Management
- Producrts creation and Mangagements
- Carts and payments managements


{
    "swagger": "2.0",
    "info": {
        "description": "A golang microservice for ecommerce",
        "title": "An ecommerce Api",
        "contact": {},
        "version": "1.0"
    },
    "host": "localhost:8000",
    "basePath": "/",
    "paths": {
        "/products/get-all": {
            "get": {
                "description": "Retrieve a list of all products",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get all Products",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/products.ProductSerializer"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/shops/create": {
            "post": {
                "description": "New shop using json request body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shops"
                ],
                "summary": "Create a new shop using",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Shops Name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/shops.ShopsSerializer"
                        }
                    }
                }
            }
        },
        "/shops/get-all": {
            "get": {
                "description": "Retrieve a list of all shops",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shops"
                ],
                "summary": "Get all shops",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/shops.ShopsSerializer"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Categories": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.Shops": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "owner": {
                    "$ref": "#/definitions/models.User"
                },
                "shop_name": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "last_name",
                "password",
                "phone"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string",
                    "maxLength": 40,
                    "minLength": 2
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string",
                    "maxLength": 40,
                    "minLength": 2
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "products.ProductSerializer": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "main_image": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "product_category": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Categories"
                    }
                },
                "shop": {
                    "$ref": "#/definitions/models.Shops"
                },
                "sub_images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "shops.ShopsSerializer": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "shop_name": {
                    "type": "string"
                }
            }
        }
    }
}