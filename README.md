# eywa
Generate UPI Payment request for pending balances in splitwise

### How to use
1. Generate Splitwise API key by registering [here](https://secure.splitwise.com/apps)
2. Use the prebuilt docker image to start the server
```
docker run -e SPLIT_KEY=<splitwise_API_key> akhilerm/eywa:latest eywa start <your_VPA_address> \
    --name "Your Name"
```
3. The URL shortening server is started by default at port `:8080`. The server can be disabled by setting
`--server=false` in the command line.

### Build
Multistage docker build is used to build the image
```
docker build -t akhilerm/eywa:latest .
```