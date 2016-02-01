# test_web_service

Run command:<br>
  go get gopkg.in/mgo.v2<br><br>

Run server code:<br>
  go run server.go<br><br>
Run client code:<br>
go run client.go --request_url=http://127.0.0.1:8000/forum/boardsfromcrawler <br>
go run client.go --request_url=http://127.0.0.1:8000/forum/boardsfromdb<br>

go run client.go --request_url=http://127.0.0.1:8000/forum/topicsfromcrawler/18.0<br>
go run client.go --request_url=http://127.0.0.1:8000/forum/topicsfromdb/18.0<br>

go run client.go --request_url=http://127.0.0.1:8000/forum/contentfromcrawler/4757.0<br>
go run client.go --request_url=http://127.0.0.1:8000/forum/contentfromdb/4757.0<br>

