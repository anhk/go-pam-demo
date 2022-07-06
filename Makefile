export GOPROXY=goproxy.cn,direct

OBJ = pam_demo.so

module:
	go build -buildmode=c-shared -mod=vendor -o ${OBJ}

prepare:
	apt install libpam0g-dev

clean:
	rm -fr ${OBJ}
