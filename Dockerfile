# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev
WORKDIR /views/
# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
ENV SRC_DIR=/go/src/github.com/jjarrett21/go_airplanes
# Add the source code:
ADD . $SRC_DIR
# Build it:
RUN cd $SRC_DIR; go build -o myapp; cp myapp /views/
ENTRYPOINT ["./myapp"]