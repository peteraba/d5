FROM python:alpine

RUN set -x \
    && apk --update add libxslt-dev libxml2-dev build-base \
    && pip install ezodf lxml klein openpyxl

COPY csv /usr/bin/
COPY xlsx /usr/bin/
COPY ods /usr/bin/
COPY spreadsheet /usr/bin/

EXPOSE 10100

CMD /usr/bin/spreadsheet --port 10100
