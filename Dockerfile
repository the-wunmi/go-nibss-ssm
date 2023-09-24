FROM ghcr.io/graalvm/graalvm-ce:22.3.1 AS graalvm

RUN gu install native-image

WORKDIR /app
COPY . ./
WORKDIR /app/ssm
RUN javac -cp com/bcpg-jdk14-145.jar:com/bcprov-ext-jdk14-145.jar:com/pgplib-2.5.jar:com/SSecurityModule.jar  -d build com/SSM.java
RUN jar --create --file SSM.jar --main-class SSM -C build .
RUN native-image -jar SSM.jar -cp com/bcpg-jdk14-145.jar:com/bcprov-ext-jdk14-145.jar:com/pgplib-2.5.jar:com/SSecurityModule.jar -J-Djava.security.properties=java.security.overrides --enable-all-security-services --enable-https '-H:AdditionalSecurityProviders=org.bouncycastle.jce.provider.BouncyCastleProvider' -H:ReflectionConfigurationFiles=reflect-config.json --rerun-class-initialization-at-runtime=org.bouncycastle.crypto.prng.SP800SecureRandom --rerun-class-initialization-at-runtime=org.bouncycastle.jcajce.provider.drbg.DRBG$Default --rerun-class-initialization-at-runtime=org.bouncycastle.jcajce.provider.drbg.DRBG$NonceAndIV --report-unsupported-elements-at-runtime -Dsecurity.provider.3=org.bouncycastle.jce.provider.BouncyCastleProvider --initialize-at-build-time --shared
WORKDIR /app

COPY --from=golang:1.20.6 /usr/local/go /usr/local/go
ENV GO_VERSION=1.20.6
ENV PATH=$PATH:/usr/local/go/bin

RUN go mod download

RUN CGO_ENABLED=1 GOOS=linux go build -o ./go-nibss-ssm

FROM debian:12-slim
WORKDIR /
COPY --from=graalvm /app/go-nibss-ssm .
COPY --from=graalvm /app/ssm ./ssm
COPY --from=graalvm /app/*.xml ./
EXPOSE 8080
ENTRYPOINT ["./go-nibss-ssm"]