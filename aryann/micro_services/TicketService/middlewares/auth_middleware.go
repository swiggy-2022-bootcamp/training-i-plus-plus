package middlewares

var (
	mySigningKey = []byte("secret")
)

func GetMySigingKey() []byte {
	return mySigningKey
}
