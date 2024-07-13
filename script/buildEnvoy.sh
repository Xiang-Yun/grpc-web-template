#docker compose up -d --build --remove-orphans

cd grpc-web-template/envoy
sudo docker rm htmx_envoy
sudo docker run --name htmx_envoy --network host -v /home/yian/yian/lpr-service/lpr-service/envoy/envoy.yaml:/etc/envoy/envoy.yaml --restart always -d thegrandpkizzle/envoy:1.26.1