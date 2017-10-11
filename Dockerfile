FROM scratch

COPY bin/linux-amd64/scrabbler /

CMD ["./scrabbler"]
