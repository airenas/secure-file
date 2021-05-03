# Add encrypted file to docker image

![Go](https://github.com/airenas/secure-file/workflows/Go/badge.svg) [![Coverage Status](https://coveralls.io/repos/github/airenas/secure-file/badge.svg?branch=main)](https://coveralls.io/github/airenas/secure-file?branch=main) [![Go Report Card](https://goreportcard.com/badge/github.com/airenas/secure-file)](https://goreportcard.com/report/github.com/airenas/secure-file) ![CodeQL](https://github.com/airenas/secure-file/workflows/CodeQL/badge.svg)

The repository hosts simple tools written with *go* and compiled without any external dependencies to help encrypt files into a docker image and decrypt files on production. 

## Example

See [example](example) dir for the complete sample. Prepare the sample image containing a secret file:

```bash
cd example
make dbuild secret=olia
```

Try run the container:
```bash
docker run -it -e "SECRET=olia" ${USER}/secure-demo:0.1
```

Press `Ctr+C`. Try with the wrong secret:

```bash
docker run -it -e "SECRET=wrong" ${USER}/secure-demo:0.1
```

You should fail seeing the data.

### Explanation

The repo contains two tools *encrypt-file* and *check-decrypt-file*. By adding secret file to a docker image we do these steps. First we create a helper temporary image. See [example/Dockerfile](example/Dockerfile): 
- during the build pass secret key with `--build-arg`. See [example/Makefile](example/Makefile) 
- copy secret file into the image. Sample file [example/secretFile.txt](example/secretFile.txt) 
- encrypt secret file using the *encrypt-file* and the key
 
Then we create the target image:  

- copy the *check-decrypt-file* into the target image (for decrypting file on startup)
- copy the encrypted file from the first image
- configure a startup script and copy it to the target image. The startup script at first runs *check-decrypt-file*. It takes secret from env setting and decrypts file. Then it should run the main process of the container. 

As you are dealing with secret data, so you must be extremely cautious. It is possible to reveal secret with `docker history` if you are using the build process without multistage building procedure.

---
### Author

**Airenas Vaičiūnas**
 
* [github.com/airenas](https://github.com/airenas/)
* [linkedin.com/in/airenas](https://www.linkedin.com/in/airenas/)


---
### License

Copyright © 2021, [Airenas Vaičiūnas](https://github.com/airenas).
Released under the [The 3-Clause BSD License](LICENSE).

---