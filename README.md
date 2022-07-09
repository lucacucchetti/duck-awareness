# duck-awareness

Collaborative stuff.

## Dev Instructions

This project uses AWS CDK so it needs NodeJS, I recommend using [nvm](https://nvm.sh/) to manage your node version.

We will be using Node version 16 (latest LTS at this time).

```
nvm use 16

node --version

# v16.16.0
```

We will use AWS CDK to manage infrastructure easily.

```
npm install -g aws-cdk

cdk --version

# 2.31.1 (build 42432c6)
```

Since the backed (API Handler lambda) will be written in Go, we need to [install it](https://go.dev/doc/install).

```
go version

# go version go1.18.3 darwin/amd64
```


## AWS Configuration

We wille be using the `eu-west-2` region (London) so make sure to configure your `~/.aws/config` [accordingly](https://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/setup-credentials.html#setup-credentials-setting-region) and your [credentials](https://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/setup-credentials.html#setup-credentials-setting).