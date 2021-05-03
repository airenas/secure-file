## encrypt file
FROM airenas/secure-file:0.2.9 as encryptor
WORKDIR /app
COPY secretFile.txt /app/
COPY secretFile2.txt /app/
COPY secretFileList.txt /app/
ARG SECRET
# use the value to set the ENV var default
ENV SECRET=$SECRET
RUN ./encrypt-file -fl secretFileList.txt -s $SECRET


## final image
FROM iron/base
LABEL Description="This image is used to demo secure file encryption/decryption in docker container" \ 
      Maintainer="airenass@gmail.com" Vendor="Airenas Vaiciunas" Version="1.0" 

WORKDIR /app
## copy binary to decrypt file
COPY --from=encryptor /app/check-decrypt-file /app/
## copy encrypted files
COPY --from=encryptor /app/sec*.txt.aes /app/
## copy file list
COPY --from=encryptor /app/secretFileList.txt /app/
## copy startup script for ensuring file decryption
COPY start.sh /app/

RUN ["chmod", "+x", "/app/start.sh"]

ENTRYPOINT ["./start.sh"]
