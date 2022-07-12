import { Stack, StackProps, CfnOutput} from 'aws-cdk-lib';
import * as iam from 'aws-cdk-lib/aws-iam';
import { PrincipalWithConditions } from 'aws-cdk-lib/aws-iam';
import { Construct } from 'constructs';

export class AwsGithubOidc extends Stack {
    constructor(scope: Construct, id: string, props?: StackProps) {
        super(scope, id, props);

        const ghOidcProvider = new iam.OpenIdConnectProvider(this, 'github-oidc-provider', {            
            url: 'https://token.actions.githubusercontent.com',
            thumbprints: ['6938fd4d98bab03faadb97b34396831e3780aea1'],
            clientIds: ['https://github.com/lucacucchetti','sts.amazonaws.com']
        })

        const ghFederatedPrincipal = new iam.FederatedPrincipal(ghOidcProvider.openIdConnectProviderArn, {            
            'StringLike': {
                'token.actions.githubusercontent.com:sub': 'repo:lucacucchetti/duck-awareness:*'
            }
        }, 'sts:AssumeRoleWithWebIdentity')

        const ghRole = new iam.Role(this, 'github-actions', {
            description: 'used by Github Actions to deploy the application stack',
            roleName: 'github-actions',
            assumedBy: ghFederatedPrincipal,
            managedPolicies: [
                iam.ManagedPolicy.fromManagedPolicyArn(this, 'github-actions-read-only-access' ,'arn:aws:iam::aws:policy/ReadOnlyAccess')
            ],
        })

        const h = new CfnOutput(this, 'githubActionsRoleArn', {
            description: 'Github Actions role ARN',
            value: ghRole.roleArn
        })

    }
}
