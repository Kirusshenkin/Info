FROM nginx:alpine

ADD nginx.conf /etc/nginx/

RUN apk update \
    && apk upgrade \
    && apk add --no-cache bash \
    && adduser -D -H -u 1000 -s /bin/bash -Gwww-data www-data

# Set upstream conf and remove the default conf
# RUN echo "upstream php-upstream { server gamblingpro.php-fpm:9000; }" > /etc/nginx/conf.d/upstream.conf \
#     && rm /etc/nginx/conf.d/default.conf
RUN rm /etc/nginx/conf.d/default.conf

CMD ["nginx"]

EXPOSE 80
