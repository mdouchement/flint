FROM alpine:latest

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh-client ca-certificates

ADD dist/flint /bin/flint

RUN adduser -D -g '' astrocorp

USER astrocorp

WORKDIR /flint

CMD ["flint"]
