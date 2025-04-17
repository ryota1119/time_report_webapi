FROM mysql:9.0

# timezone
RUN ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
ENV TZ=Asia/Tokyo

# language
RUN microdnf update -y \
    && microdnf install -y glibc-locale-source \
    && localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8

ENV LANG ja_JP.UTF-8
ENV LC_ALL ja_JP.UTF-8