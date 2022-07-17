import { Stack, StackProps, CfnOutput } from 'aws-cdk-lib';
import {
    Effect,
    FederatedPrincipal,
    ManagedPolicy,
    OpenIdConnectProvider,
    PolicyDocument,
    PolicyStatement,
    Role
} from 'aws-cdk-lib/aws-iam';
import { Construct } from 'constructs';

export class AwsGithubOidc extends Stack {
    constructor(scope: Construct, id: string, props?: StackProps) {
        super(scope, id, props);

        const ghOidcProvider = new OpenIdConnectProvider(this, 'github-oidc-provider', {
            url: 'https://token.actions.githubusercontent.com',
            thumbprints: ['6938fd4d98bab03faadb97b34396831e3780aea1'],
            clientIds: ['sts.amazonaws.com']
        })

        const ghFederatedPrincipal = new FederatedPrincipal(ghOidcProvider.openIdConnectProviderArn, {
            'StringLike': {
                'token.actions.githubusercontent.com:sub': 'repo:lucacucchetti/duck-awareness:*'
            }
        }, 'sts:AssumeRoleWithWebIdentity')

        const ghRole = new Role(this, 'github-actions', {
            description: 'used by Github Actions to deploy the application stack',
            roleName: 'github-actions',
            assumedBy: ghFederatedPrincipal,
            managedPolicies: [
                ManagedPolicy.fromManagedPolicyArn(this, 'github-actions-read-only-access', 'arn:aws:iam::aws:policy/ReadOnlyAccess')
            ],
            inlinePolicies: {
                AllowAssumeCdkRole: new PolicyDocument({
                    statements: [new PolicyStatement({
                        actions: ['sts:AssumeRole'],
                        effect: Effect.ALLOW,
                        resources: [`arn:aws:iam::${this.account}:role/cdk*`]
                    })]
                })
            }
        })

        const h = new CfnOutput(this, 'githubActionsRoleArn', {
            description: 'Github Actions role ARN',
            value: ghRole.roleArn
        })

    }
}
