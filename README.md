# zzz
* Correct request: http://localhost:7999/zzz/?ms=1000
  * "ok" will return after 1000 milli seconds later.
* Uncorrect request: http://localhost:7999/zzz/?ms=20000
  * "Bad Request" will return with status code 200(ok).

## Usage
* git clone https://github.com/shogonir/zzz.git
* cd zzz
* go build
* ./zzz

## Docker Build
* git clone https://github.com/shogonir/zzz.git
* cd zzz
* go build
* docker build -t zzz:latest .
* docker run -d -p 7999:7999 zzz:latest

