# What is this?
I have a lot of Philps hue lights, and I want to be able to control my lights
from my terminal and keyboard shortcuts in my window manager.

# Install
```bash
git clone git@github.com:skykosiner/control-hue-lights-golang
cd ./control-hue-lights-golang
go build -o lights ./cmd/main.go && mv lights ~/.local/bin/lights
```
