FROM alpine

LABEL maintainer="Yosia <yosiaagustadewa@gmail.com>"

WORKDIR /app

COPY user .

RUN chmod +x /app/user

EXPOSE 8001

CMD ["./user"]