FROM maven:3.8.5-openjdk-17 AS build
WORKDIR /project
COPY . .
RUN mvn package -B -D maven.test.skip=true

FROM openjdk:19-jdk-alpine3.16 AS production
RUN apk add --no-cache tzdata
ARG TAG_BUILD_NUMBER
ENV TAG_BUILD_NUMBER=$TAG_BUILD_NUMBER
ENV TZ=Europe/Moscow
VOLUME /tmp
COPY --from=build /project/target/gid-service-1.0-SNAPSHOT.jar app.jar
ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom","-jar","/app.jar"]

