#!/usr/bin/env sh

echo "running chum pedersen test in python"
cd py
pytest --log-cli-level=DEBUG chaum_pedersen_test.py
cd ..
echo "done"

echo "running chum pedersen test in go"
cd go
go test -v
cd ..
echo "done"


