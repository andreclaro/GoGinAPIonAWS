import { App, Duration, Stack } from "@aws-cdk/core";
import { Code, Function, Runtime } from "@aws-cdk/aws-lambda";
import { LambdaRestApi } from "@aws-cdk/aws-apigateway";
import { Asset } from "@aws-cdk/aws-s3-assets";
import * as path from "path";

export class GoAPILambdaStack extends Stack {
  constructor(app: App, id: string) {
    super(app, id);

    // See https://docs.aws.amazon.com/cdk/api/latest/docs/aws-s3-assets-readme.html
    const lambdaAsset = new Asset(
      // @ts-ignore - this expects Construct not cdk.Construct :thinking:
      this,
      "GoAPILambdaFunctionZip",
      {
        path: path.join(__dirname, "../lambda-function/"),
      }
    );

    const lambdaFn = new Function(this, "GoAPILambdaFunction", {
      code: Code.fromBucket(lambdaAsset.bucket, lambdaAsset.s3ObjectKey),
      timeout: Duration.seconds(300),
      runtime: Runtime.GO_1_X,
      handler: "main",
    });

    // API Gateway
    new LambdaRestApi(
      // @ts-ignore - this expects Construct not cdk.Construct :thinking:
      this,
      "GoAPIGatewayRestAPI",
      {
        handler: lambdaFn,
      }
    );
  }
}

const app = new App();
new GoAPILambdaStack(app, "GoAPILambdaStack");
app.synth();
