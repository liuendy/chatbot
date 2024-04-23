# 使用官方 Golang 镜像作为构建环境
FROM golang:1.18 as builder


ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载所有依赖
RUN go mod download

# 复制源代码到容器中
COPY . .

# 编译应用
RUN CGO_ENABLED=0 GOOS=linux go build -o chatbot ./cmd

# 使用 scratch 作为运行环境
FROM scratch

# 从 builder 镜像中复制编译好的应用
COPY --from=builder /app/chatbot /chatbot

ENV TZ Asia/Shanghai

# 暴露端口
EXPOSE 8080

# 运行聊天机器人应用
CMD ["/chatbot"]
