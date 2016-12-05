FROM python:alpine

RUN set -x \
    && apk --update add libxslt-dev libxml2-dev build-base \
    && pip install ezodf lxml klein openpyxl
