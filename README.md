# GoShareFiles
Small file server written in Go. Mainly for temporary file sharing over local networks 

# Usage
```
go build -o GoShareFiles ./main.go
```

Then run compiled exe from the directory with files you want to share.
Files are now accessible from any browser on you local network like that:
`%host_ip%:65535/`
You can check your local ip in `ip addr`
