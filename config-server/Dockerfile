FROM bellsoft/liberica-openjre-centos:17.0.8 as builder
ARG BUILD_ID
LABEL stage=build
LABEL build=$BUILD_ID
WORKDIR /app
COPY build/libs/*.jar /app/service.jar
RUN java -Djarmode=layertools -jar service.jar extract

FROM bellsoft/liberica-openjre-centos:17.0.8
WORKDIR /app
ENV TZ=Europe/Moscow
ENV JAVA_OPTS=""
COPY --from=builder app/spring-boot-loader ./
COPY --from=builder app/dependencies/ ./
COPY --from=builder app/k8s/ ./
COPY --from=builder app/application/ ./
ENTRYPOINT exec java $JAVA_OPTS org.springframework.boot.loader.JarLauncher
