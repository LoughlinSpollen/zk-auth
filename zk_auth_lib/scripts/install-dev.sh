#! /usr/bin/env sh

ver=$(python -V 2>&1 | sed 's/.* \([0-9]\).\([0-9]\).*/\1\2/')
if [ "$ver" -lt "37" ]; then
    echo "please install python 3.7 or greater"
    exit 1
fi

pip --no-cache-dir install -r requirements.txt

go mod init zk_auth_lib
