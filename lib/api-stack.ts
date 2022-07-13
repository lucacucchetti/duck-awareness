import { GoFunction } from '@aws-cdk/aws-lambda-go-alpha';
import { Stack, StackProps } from 'aws-cdk-lib';
import { LambdaRestApi } from 'aws-cdk-lib/aws-apigateway';
import { Construct } from 'constructs';

export class ApiStack extends Stack {
  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    // https://docs.aws.amazon.com/cdk/api/v2/docs/aws-lambda-go-alpha-readme.html
    const handler = new GoFunction(this, 'DuckAwareness-APIHandler', {
      entry: `${__dirname}/app/api-handler/main.go`,
      logRetention: 7,
      // If we need environment variables (like a database name/arn)
      // bundling: {
      //   environment: {
      //     HELLO: 'WORLD',
      //   },
      // },
    });

    // Default endpoint type is EndpointType.EDGE which suits us fine: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-endpoint-types.html
    new LambdaRestApi(this, 'DuckAwareness-API', {
      endpointExportName: 'DuckAwareness-APIUrl',
      proxy: true, // route all requests to the Lambda Function, in the future if we want to use the same API for the UI we might want to add a proxy resouce instead
      handler,
      integrationOptions: { allowTestInvoke: false },
      deployOptions: {
        throttlingBurstLimit: 5, // concurrently
        throttlingRateLimit: 10, // per second
      },
    });
  }
}
