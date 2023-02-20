#!/bin/bash
docker run \
   -p 9000:9000 \
   -p 9090:9090 \
   --name minio \
   -v ~/minio/data:/data \
   -e "MINIO_ROOT_USER=minioadmin" \
   -e "MINIO_ROOT_PASSWORD=minioadmin" \
   quay.io/minio/minio server /data --console-address ":9090"


# docker run \
#    -p 9000:9000 \
#    -p 9090:9090 \
#    --name minio \
#    -v /tmp/data:/data \
#    -e "MINIO_ROOT_USER=admin" \
#    -e "MINIO_ROOT_PASSWORD=nonprodpasswd" \
#    minio/minio server /data --console-address ":9090"