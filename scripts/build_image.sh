#!/bin/bash

AWS_DEFAULT_REGION="ap-southeast-1"
AWS_ACCOUNT_ID="461429446948"
IMAGE_NAME="commission-service-dev"

echo "AWS_DEFAULT_REGION $AWS_DEFAULT_REGION";
echo "AWS_ACCOUNT_ID $AWS_ACCOUNT_ID";
echo "IMAGE_NAME $IMAGE_NAME";
echo "BRANCH_NAME $BRANCH_NAME";

IMAGE_REGISTRY=$AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com
REPOSITORY_URI=$IMAGE_REGISTRY/$IMAGE_NAME

echo Building the Docker image...
if [ "$BRANCH_NAME" == "commission-service-dev" ]; then
    docker build -t $REPOSITORY_URI:$BRANCH_NAME-latest .
fi

echo Pushing the Docker image...
docker push $REPOSITORY_URI:$BRANCH_NAME-latest

if [ -z $1 ]; then
    echo "Don't have tag";
else
    TAG=$1
    echo "tag is set to $TAG";
    echo "Tag image with tag $TAG ..."
    docker tag $REPOSITORY_URI:$BRANCH_NAME-latest $REPOSITORY_URI:$BRANCH_NAME-$TAG

    echo "Pushing the Docker image with tag $TAG ..."
    docker push $REPOSITORY_URI:$BRANCH_NAME-$TAG
fi
