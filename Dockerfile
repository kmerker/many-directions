## BUILDING
##   (from project root directory)
##   $ docker build -t kmerker-many-directions .
##
## RUNNING
##   $ docker run kmerker-many-directions

FROM gcr.io/stacksmith-images/debian:wheezy-r07

MAINTAINER Bitnami <containers@bitnami.com>

ENV STACKSMITH_STACK_ID="muexek2" \
    STACKSMITH_STACK_NAME="kmerker/many-directions" \
    STACKSMITH_STACK_PRIVATE="1"

RUN bitnami-pkg install go-1.6.2-0 --checksum 79b828a23b7582197799e775d65501b082d71d589ba6eed7aa3d68cf75b94a19

ENV PATH=/opt/bitnami/go/bin:$PATH

## STACKSMITH-END: Modifications below this line will be unchanged when regenerating

# Go base template
COPY . /app
WORKDIR /app
