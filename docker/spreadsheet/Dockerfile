FROM peteraba/d5-python

COPY bin/csv /usr/bin/
COPY bin/xlsx /usr/bin/
COPY bin/ods /usr/bin/
COPY bin/spreadsheet /usr/bin/

EXPOSE 10100

CMD /usr/bin/spreadsheet --port 10100
