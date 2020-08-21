// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
const exec = require('child_process').exec;

// limit to 8 concurrent builds
const sem = require('./semaphore')(8);

const swaggerDir = 'src/node_modules/@microsoft.azure/autorest.testserver/swagger/';

const goMappings = {
    'additionalpropsgroup': 'additionalProperties.json',
    'arraygroup': 'body-array.json',
    'azurereportgroup': 'azure-report.json',
    'azurespecialsgroup': 'azure-special-properties.json',
    'booleangroup': 'body-boolean.json',
    'bytegroup': 'body-byte.json',
    'complexgroup': 'body-complex.json',
    'complexmodelgroup': 'complex-model.json',
    'custombaseurlgroup': 'custom-baseUrl.json',
    //'dategroup': 'body-date.json',
    'datetimegroup': 'body-datetime.json',
    'datetimerfc1123group': 'body-datetime-rfc1123.json',
    'dictionarygroup': 'body-dictionary.json',
    'durationgroup': 'body-duration.json',
    'errorsgroup': 'xms-error-responses.json',
    'extenumsgroup': 'extensible-enums-swagger.json',
    'filegroup': 'body-file.json',
    //'formdatagroup': 'body-formdata.json',
    'headergroup': 'header.json',
    'httpinfrastructuregroup': 'httpInfrastructure.json',
    'integergroup': 'body-integer.json',
    'lrogroup': 'lro.json',
    'migroup': 'multiple-inheritance.json',
    //'modelflatteninggroup': 'model-flattening.json',
    'morecustombaseurigroup': 'custom-baseUrl-more-options.json',
    'nonstringenumgroup': 'non-string-enum.json',
    'numbergroup': 'body-number.json',
    'optionalgroup': 'required-optional.json',
    'paginggroup': 'paging.json',
    'paramgroupinggroup': 'azure-parameter-grouping.json',
    'reportgroup': 'report.json',
    'stringgroup': 'body-string.json',
    'urlgroup': 'url.json',
    'urlmultigroup': 'url-multi-collectionFormat.json',
    'validationgroup': 'validation.json',
    'xmlgroup': 'xml-service.json',
};

// loop through all of the namespaces in goMappings
for (namespace in goMappings) {
    // for each swagger run the autorest command to generate code based on the swagger for the relevant namespace and output to the /generated directory
    const inputFile = swaggerDir + goMappings[namespace];
    generate(inputFile, 'test/autorest/generated/' + namespace);
}

const blobStorage = 'https://raw.githubusercontent.com/Azure/azure-rest-api-specs/storage-dataplane-preview/specification/storage/data-plane/Microsoft.BlobStorage/preview/2019-07-07/blob.json';
generate(blobStorage, 'test/storage/2019-07-07/azblob', '--credential-scope="https://storage.azure.com/.default" --module="azstorage" --export-client="false" --file-prefix="zz_generated_" --openapi-type="data-plane"');

const network = 'https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/network/resource-manager/readme.md';
generateFromReadme(network, 'package-2020-03', 'test/network/2020-03-01/armnetwork', '--credential-scope="https://management.azure.com//.default" --module=armnetwork');

// const synapseArtifacts = 'https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/synapse/data-plane/Microsoft.Synapse/preview/2019-06-01-preview/artifacts.json';
// generate(synapseArtifacts, 'test/synapse/2019-06-01/azartifacts', '--module="azartifacts" --openapi-type="data-plane"');

// helper to log the package being generated before invocation
function generate(inputFile, outputDir, additionalArgs) {
    console.log('generating ' + inputFile);
    sem.take(function() {
        exec('autorest --use=. --clear-output-folder --license-header=MICROSOFT_MIT_NO_VERSION --input-file=' + inputFile + ' --output-folder=' + outputDir + ' ' + additionalArgs, autorestCallback(outputDir, inputFile));
    });
}

function generateFromReadme(readme, tag, outputDir, additionalArgs) {
    console.log('generating ' + readme);
    sem.take(function() {
        exec('autorest --use=. ' + readme + ' --tag=' + tag + ' --clear-output-folder --license-header=MICROSOFT_MIT_NO_VERSION --output-folder=' + outputDir + ' ' + additionalArgs, autorestCallback(outputDir, readme));
    });
}

// use a function factory to create the closure so that the values of namespace and inputFile are captured on each iteration
function autorestCallback(outputDir, inputFile) {
    return function (error, stdout, stderr) {
        // print any output or error from the autorest command
        logResult(error, stdout, stderr);
        console.log('done generating ' + inputFile);
        // format the output on success
        // print any output or error from go fmt
        if (stderr === '' && error === null) {
            exec('go fmt ./' + outputDir,
            function (error, stdout, stderr) {
                console.log('formatting ' + outputDir);
                logResult(error, stdout, stderr);
            });
        }
        sem.leave();
    };
}

function logResult(error, stdout, stderr) {
    if (stdout !== '') {
        console.log('stdout: ' + stdout);
    }
    if (stderr !== '') {
        console.error('\x1b[91m%s\x1b[0m', 'stderr: ' + stderr);
    }
    if (error !== null) {
        console.error('\x1b[91m%s\x1b[0m', 'exec error: ' + error);
    }
}