export GOPROXY=goproxy.cn,direct

OBJ = pam_demo.so

module:
	go build -buildmode=c-shared -o ${OBJ}

clean:
	rm -fr ${OBJ}
