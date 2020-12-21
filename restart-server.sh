#! /bin/sh

if [ -f "process.cid" ]; then
    bash stop-server.sh
fi

bash start-server.sh&
