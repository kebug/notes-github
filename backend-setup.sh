mkdir note-sync-backend
cd note-sync-backend
go mod init note-sync
# 安装必要依赖
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/google/go-github/v45/github
go get -u golang.org/x/oauth2 