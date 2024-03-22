#! /usr/bin/env zsh

echo
echo "Installing PostgreSQL 14"
brew install postgresql@14
ln -sfv /usr/local/opt/postgresql@14/*.plist ~/Library/LaunchAgents

echo "Adding start stop aliases to .zshrc"
echo "" >> ~/.zshrc
echo "alias pg_start=\"launchctl load ~/Library/LaunchAgents/homebrew.mxcl.postgresql@14.plist\"" >> ~/.zshrc
echo "alias pg_stop=\"launchctl unload ~/Library/LaunchAgents/homebrew.mxcl.postgresql@14.plist\"" >> ~/.zshrc
echo "alias pg_restart=\"pg_stop && pg_start\"" >> ~/.zshrc

exec $SHELL -l

pg_start 



