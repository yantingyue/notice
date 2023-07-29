#依赖镜像（母镜像），可以先使用docker search命令搜索
FROM notice:latest


#Docker工作目录
WORKDIR /app/job


#在Docker工作目录下执行命令
RUN go build ./app/job/main.go

#暴露端口
EXPOSE 8888

#编译后在根目录下生成而非./core目录下，最终运行Docker的初始命令
ENTRYPOINT ["./main"]
