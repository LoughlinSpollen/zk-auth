#! /usr/bin/env sh

echo "Starting python zk-client"

if [ ${DEBUG} = true ]; then
    python -O ./main.py
else
    python ./main.py
fi
