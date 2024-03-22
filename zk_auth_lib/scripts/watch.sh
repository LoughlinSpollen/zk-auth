#!/usr/bin/env zsh

watchexec -w . -e py --restart 'clear && ./scripts/test.sh'