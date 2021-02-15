## encrypt file
FROM airenas/secure-file:0.1.2 as encryptor
WORKDIR /app
COPY secretFile.txt /app/
ARG SECRET
# use the value to set the ENV var default
ENV SECRET=$SECRET
RUN ./encrypt-file -f secretFile.txt -s $SECRET > secretFile.txt.aes


## final image
FROM iron/base
LABEL Description="This image is used to demo secure file encryption/decryption in docker container" \ 
      Maintainer="airenass@gmail.com" Vendor="Airenas Vaiciunas" Version="1.0" 

WORKDIR /app
## copy binary to decrypt file
COPY --from=encryptor /app/check-decrypt-file /app/
## copy encrypted file
COPY --from=encryptor /app/secretFile.txt.aes /app/
## copy startup script for ensuring file decryption
COPY start.sh /app/
RUN ["chmod", "+x", "/app/start.sh"]

ENTRYPOINT ["./start.sh"]
