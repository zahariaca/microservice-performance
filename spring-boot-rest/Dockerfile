FROM openjdk:21

RUN mkdir /app

COPY ./build/lib/spring-boot-rest/build/libs/spring-boot-rest-0.0.1-SNAPSHOT.jar   /app/app.jar

CMD [ "java", "-jar", "/app/app.jar"]