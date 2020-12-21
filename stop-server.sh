#! /bin/sh

docker exec $(cat process.cid) kill -1 1
rm process.cid