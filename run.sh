
protoc -I=. --go_out=./cmd/web --go_opt=paths=source_relative --go-grpc_out=./cmd/web --go-grpc_opt=paths=source_relative api_proto/web.proto

protoc -I=. api_proto/web.proto --js_out=import_style=commonjs:./static/js --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./static/js

browserify static/js/api_proto/client.js static/js/api_proto/app.js | uglifyjs -cm > static/js/main.js
