#!/bin/bash

echo "echo Logging in to Docker ..."

AWS_DEFAULT_REGION="ap-southeast-1"
AWS_ACCOUNT_ID="461429446948"
IMAGE_NAME="commission-service-dev"
BRANCH_NAME="$APPLICATION_NAME"
AWS_ACCESS_KEY_ID="AKIAWW32WTESP7FWZU5M"
AWS_SECRET_ACCESS_KEY="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"

aws configure set aws_access_key_id $AWS_ACCESS_KEY_ID
aws configure set aws_secret_access_key $AWS_SECRET_ACCESS_KEY
aws configure set default.region $AWS_DEFAULT_REGION

echo "AWS_DEFAULT_REGION $AWS_DEFAULT_REGION";
echo "AWS_ACCOUNT_ID $AWS_ACCOUNT_ID";
echo "IMAGE_NAME $IMAGE_NAME";
echo "BRANCH_NAME $BRANCH_NAME";

aws --version
aws ecr get-login-password --region $AWS_DEFAULT_REGION | docker login --username AWS --password-stdin $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com
aws ecr describe-repositories

docker pull $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/$IMAGE_NAME:$BRANCH_NAME-latest

if [ "$BRANCH_NAME" == "commission-service-dev" ]; then
    docker run -p 5009:8010 -d --name $BRANCH_NAME $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/$IMAGE_NAME:$BRANCH_NAME-latest
fi

if [ "$BRANCH_NAME" == "commission-service-stage" ]; then
    docker run -p 5669:8010 -d --name $BRANCH_NAME $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/$IMAGE_NAME:$BRANCH_NAME-latest
fi

if [ "$BRANCH_NAME" == "commission-service-preprod" ]; then
    docker run -p 5669:8010 -d --name $BRANCH_NAME $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/$IMAGE_NAME:$BRANCH_NAME-latest
fi

# docker exec $BRANCH_NAME bash -c "composer dump-autoload"

docker images --filter "dangling=true" -q --no-trunc | xargs -r docker rmi -f
echo "Success true"
