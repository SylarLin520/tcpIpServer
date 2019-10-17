# tcpIpServer
tcpIpServer

router  
post http://localhost:8080/2thrift  
post http://localhost:8080/somePost  
get http://localhost:8080/someGet  



go get github.com/gin-gonic/gin

go get git.apache.org/thrift.git/lib/go

thrift -r --gen go echo.thrift

find ./src -type d -name .git -exec rm {} \;