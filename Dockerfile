FROM golang:1.18

# Add Maintainer Info
LABEL maintainer="Poornima Y N <rg86poornimayn@gmail.com>"
RUN mkdir src1
COPY src ./src1/
WORKDIR ./src1
RUN make all
CMD ["make all"]

