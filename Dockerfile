FROM bmartel/go:nonroot

ENV APP_ENV production
ENV APP_DEBUG false
ENV APP_DOMAIN http://localhost:8080
ENV APP_PORT :8080
ENV SESSION_KEY pulse_session

ENV VIEW_DIR views
ENV VIEW_EXT .jade

ENV DATABASE_URL /src/db/data.sqlite3
ENV DATABASE_TYPE sqlite3
ENV ASSET_PATH /static
ENV ASSET_DIR public

COPY . /src
WORKDIR /src
RUN build
USER nobody

EXPOSE 8080

CMD [ "/src/webapp" ]
