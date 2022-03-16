package config

import "net/url"

const SERVER_PORT = 5001

var MONGO_URL string = "mongodb+srv://rshantharaju:" + url.QueryEscape("Ravi@1999") + "@cluster0.05bio.mongodb.net/test"
