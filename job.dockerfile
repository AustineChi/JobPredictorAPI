##Build stage
FROM golang:1.21-alpine as builder

#run the command on the docker image we are building
RUN mkdir /imageApp

COPY . /imageApp

WORKDIR /imageApp

#Build the golang code
RUN CGO_ENABLE=0 go build -o jobApp ./main.go

#Run chmod command and add the executable flag
RUN chmod +x /imageApp/jobApp

##Runnig Stage
FROM alpine:latest

RUN mkdir /imageApp

#Coyy files from the build stage to /imageApp
COPY --from=builder /imageApp/jobApp /imageApp

EXPOSE 8080

CMD [ "/imageApp/jobApp" ]