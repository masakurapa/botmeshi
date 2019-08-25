#!/bin/sh

aws cloudformation package \
    --template-file botmeshi.yml \
    --s3-bucket $AWS_BUCKET_NAME \
    --s3-prefix package \
    --output-template-file package.yml

aws cloudformation deploy \
    --template-file package.yml
    --stack-name botmeshi-endpoint
    --s3-bucket $AWS_BUCKET_NAME \
    --s3-prefix deploy \
    --capabilities CAPABILITY_NAMED_IAM \
    --role-arn $AWS_DEPLOY_ROLE_ARN \
    --parameter-overrides \
        BotVerificationToken=$BOT_VERIFICATION_TOKEN \
        BotAccessToken=$BOT_ACCESS_TOKEN \
        SearchEngineId=$SEARCH_ENGINE_ID \
        PlaceApiKey=$PLACE_API_KEY \
        CustomSearchApiKey=$CUSTOM_SEARCH_API_KEY
