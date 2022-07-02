module github.com/crackeer/go-api-svc

go 1.16

require (
	github.com/crackeer/gopkg v0.0.0-20220626084626-5001afd3bc0e
	github.com/gin-gonic/gin v1.8.1
	github.com/mattn/go-sqlite3 v1.14.9
	gorm.io/driver/mysql v1.3.4 // indirect
	gorm.io/driver/sqlite v1.2.6 // indirect
	gorm.io/gorm v1.23.4
)

// replace github.com/crackeer/gopkg => /Users/liuhu016/github/my-gopkg

replace github.com/crackeer/gopkg => D:/github/gopkg
