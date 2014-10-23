export GOPATH=$(PWD)

all: clean find dependencies gobuild done

find:
	@which go > /dev/null || if [ $$? -ne 0 ]; then echo "Can't find go! Please 'sudo apt-get install golang'"; exit 1; fi

bin:
	mkdir bin

pkg:
	mkdir pkg

dependencies: bin pkg
	go get github.com/go-gl/gl
	go get github.com/go-gl/glu
	go get github.com/rhencke/glut
	go get github.com/ianremmler/ode
	go get gopkg.in/qml.v1

gobuild: 
	go build -o hw6 main

run:
	@./hw6

clean:
	@rm -f ./hw6
	@rm -rf pkg/
	@rm -rf src/gopkg.in
	@rm -rf src/github.com

done:
	@echo "BUILD COMPLETE!"

#/%:
#/	@echo "Building package $@"
#/	@go build $@
#/	@go install $@
#/
# glutil world util actor entity 
