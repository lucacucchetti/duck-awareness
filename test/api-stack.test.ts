import { App } from 'aws-cdk-lib';
import { Template } from 'aws-cdk-lib/assertions';
import { ApiStack } from '../lib/api-stack';

test('Lambda and API Gateway created', () => {
  const app = new App();
  const stack = new ApiStack(app, 'MyTestStack');
  const template = Template.fromStack(stack);

  template.resourceCountIs('AWS::Lambda::Function', 2); // in order to set log retention, the CDK will use a custom resource lambda to create the log group with a custom retention period
  template.resourceCountIs('AWS::ApiGateway::RestApi', 1);

  template.hasResourceProperties('AWS::ApiGateway::Stage', {
    MethodSettings: [{ ThrottlingBurstLimit: 5, ThrottlingRateLimit: 10 }],
  });
});
