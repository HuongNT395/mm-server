FROM postgres:9.4.21
ENV PATH="/opt/mattermost/bin:${PATH}"
# Configure SQL
RUN apt-get update && apt-get install -y ca-certificates curl jq musl-dev glibc-source glibc-doc
RUN mkdir -p /opt/mattermost
COPY . /opt/mattermost/
VOLUME /opt/mattermost/data
VOLUME /opt/mattermost/plugins
VOLUME /opt/mattermost/logs
VOLUME /opt/mattermost/config
ENV POSTGRES_DB=mattermost
ENV POSTGRES_USER=mmuser
ENV POSTGRES_PASSWORD=mmpassword
EXPOSE 8065
COPY docker-entry.sh .
RUN chmod +x ./docker-entry.sh
ENTRYPOINT ["sh","./docker-entry.sh"]
HEALTHCHECK CMD curl --fail http://localhost:8065 || exit 1