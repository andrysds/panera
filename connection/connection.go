package connection

// Init initiates application connections
func Init() {
	InitMongoDB()
	InitRedis()
}
