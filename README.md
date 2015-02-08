Hello World Application
# Build Image
docker build -t 192.168.1.36:5000/daocloud/hello .

# Push Image to Registry
docker push 192.168.1.36:5000/daocloud/hello

# Run Container
docker run -d -p 80:80 192.168.1.36:5000/daocloud/hello
