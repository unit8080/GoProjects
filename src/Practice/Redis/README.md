# Redis service


# How to build

go build  -o <redis-server> *.go

# Launch ipam server

nohup ./redis-server > redis-nohup.out  2>&1 &

# File structure

- main.go: 
    - launches server 
    
# Future enhancements

