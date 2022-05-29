FROM golang:alpine AS builder

# 为我们设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    ProRunMode=prod \
    GOPROXY=https://proxy.golang.com.cn,direct

# 移动到工作目录: /build
WORKDIR /build

# 复制项目中的 go.mod和go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件beegoBackstage
RUN go build -o beegoBackstage .

###################
# 接下来创建一个小镜像
###################
FROM scratch

COPY ./conf /conf
COPY ./static /static

# 从builder镜像中把/dist/beegoBackstage 拷贝到当前目录
COPY --from=builder /build/beegoBackstage /

# 声明服务端口
EXPOSE 8098

# 需要的命令
CMD ["conf/app.conf","/beegoBackstage"]