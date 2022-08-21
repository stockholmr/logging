package logging

import "log"

func Error(msg string) {
	log.Print("ERROR ", msg)
}

func Errorf(pattern string, args ...interface{}) {
	log.Printf("ERROR "+pattern, args...)
}

func Fatal(msg string) {
	log.Fatal("FATAL ", msg)
}

func Fatalf(pattern string, args ...interface{}) {
	log.Fatalf("FATAL "+pattern, args...)
}

func Info(msg string) {
	log.Print("INFO ", msg)
}

func Infof(pattern string, args ...interface{}) {
	log.Printf("INFO "+pattern, args...)
}
