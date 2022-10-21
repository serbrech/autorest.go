/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { AutorestExtensionHost } from '@autorest/extension-base';
import { Config } from '../common/constant';
import { ImportManager } from '@autorest/go/dist/generator/imports';
import { TestCodeModel } from '@autorest/testmodeler/dist/src/core/model';
import { TestConfig } from '@autorest/testmodeler/dist/src/common/testConfig';
export class GenerateContext {
  public packageName: string;
  public importManager: ImportManager;

  public constructor(public host: AutorestExtensionHost, public codeModel: TestCodeModel, public testConfig: TestConfig, public swaggerCommit = 'main') {
    this.packageName = this.codeModel?.language?.go?.packageName;
    this.importManager = new ImportManager();
    let module = undefined;
    if (this.packageName) {
      const majorVersion = this.testConfig.getValue(Config.moduleVersion, '1.0.0').split('.')[0];
      if (majorVersion > 1) {
        module = `${this.testConfig.getValue(
          Config.module,
          `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/${this.packageName.substr(3)}/${this.packageName}`,
        )}/v${majorVersion}`;
      } else {
        module = this.testConfig.getValue(Config.module, `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/${this.packageName.substr(3)}/${this.packageName}`);
      }
    }
    if (module) {
      this.importManager.add(module);
    }
  }
}
