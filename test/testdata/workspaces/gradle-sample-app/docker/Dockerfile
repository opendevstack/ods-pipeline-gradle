FROM registry.access.redhat.com/ubi8/openjdk-17-runtime:1.13

COPY app.jar app.jar

EXPOSE 8080

CMD ["java","-Xmx512m", "-jar", "app.jar"]
