#!/usr/bin/env node
import { App } from 'aws-cdk-lib';
import { ApiStack } from '../lib/api-stack';
import { AwsGithubOidc } from '../lib/aws-github-oidc';

const app = new App();
new ApiStack(app, 'DuckAwareness-APIStack');
new AwsGithubOidc(app, 'AwsGithubOidc')