import * as cdk from "@aws-cdk/core";
import { Code, Function, Runtime } from "@aws-cdk/aws-lambda";
import { HttpApi, HttpMethod } from "@aws-cdk/aws-apigatewayv2";
import { join } from "path";
import { LambdaProxyIntegration } from "@aws-cdk/aws-apigatewayv2-integrations";
import {Duration} from "@aws-cdk/core";

export class HSLMonthlyTransitStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    // The code that defines your stack goes here
    const transitCalculatorFn = new Function(this, "TransitCalculator", {
      runtime: Runtime.GO_1_X,
      handler: "main",
      timeout: Duration.seconds(10),
      code: Code.fromAsset(join(__dirname, "../../function.zip")),
    });

    const TransitCalculatorInt = new LambdaProxyIntegration({
      handler: transitCalculatorFn,
    });

    const httpApi = new HttpApi(this, "CommuteAPI");
    httpApi.addRoutes({
      path: "/total-transit",
      methods: [HttpMethod.POST],
      integration: TransitCalculatorInt,
    });
  }
}
