package config

import "net/url"

const USER_SERVICE_SERVER_PORT = 5004

var MONGO_URL string = "mongodb+srv://rshantharaju:" + url.QueryEscape("Ravi@1999") + "@cluster0.05bio.mongodb.net/test"
