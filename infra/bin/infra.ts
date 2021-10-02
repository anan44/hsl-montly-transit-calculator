#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from '@aws-cdk/core';
import { HSLMonthlyTransitStack } from '../lib/infra-stack';

const app = new cdk.App();
new HSLMonthlyTransitStack(app, 'HSL-Transit-Stack');
