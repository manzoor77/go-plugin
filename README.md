# go-plugin
build a math plugin that performs multiple mathematical operations. It start by creating a package main that performs basic mathematical operations viz, addition and subtraction. 

build the Go plugin using the command:

go build -buildmode=plugin -o math.so

Above command uses the buildmode=plugin that suggests the Go compiler to compile the code in plugin mode and not as a package or a command binary.
math.so is the shared library or the plugin generated by Go compiler. You could choose to store the built plugin anywhere on your file system (generally recommended to have a plugins folder).

