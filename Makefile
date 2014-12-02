export GOPATH=$(PWD)

#	go get gopkg.in/qml.v1

all: light_clean find dependencies gobuild done

find:
	@which go > /dev/null || if [ $$? -ne 0 ]; then echo "Can't find go! Please 'sudo apt-get install golang'"; exit 1; fi

bin:
	mkdir bin

pkg:
	mkdir pkg

dependencies: bin pkg
	go get github.com/go-gl/gl
	go get github.com/go-gl/glu
	go get github.com/go-gl/glh
	go get github.com/rhencke/glut
	go get github.com/ianremmler/ode

gobuild: 
	go build -o project main

run:
	@./project

clean:
	@rm -f ./project
	@rm -rf pkg/
	@rm -rf src/gopkg.in
	@rm -rf src/github.com

light_clean:
	@rm -f pkg/*/*.a

done:
	@echo "BUILD COMPLETE!"

#/%:
#/	@echo "Building package $@"
#/	@go build $@
#/	@go install $@
#/
# glutil world util actor entity 
