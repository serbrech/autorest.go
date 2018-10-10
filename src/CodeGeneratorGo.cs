﻿// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

using AutoRest.Core;
using AutoRest.Core.Model;
using AutoRest.Go.Model;
using AutoRest.Go.Templates;
using System;
using System.Linq;
using System.Threading.Tasks;
using System.Collections.Generic;

namespace AutoRest.Go
{
    public class CodeGeneratorGo : CodeGenerator
    {
        public string Name
        {
            get { return "Go"; }
        }

        public override string UsageInstructions
        {
            get { return string.Empty; }
        }

        public override string ImplementationFileExtension
        {
            get { return ".go"; }
        }

        /// <summary>
        /// Generates Go code for service client.
        /// </summary>
        /// <param name="serviceClient"></param>
        /// <returns></returns>
        public override async Task Generate(CodeModel cm)
        {
            // if preview-chk:true is specified verify that preview swagger is output under a preview subdirectory.
            // this is a bit of a hack until we have proper support for this in the swagger->sdk bot so it's opt-in.
            if (Settings.Instance.Host.GetValue<bool>("preview-chk").Result)
            {
                const string previewSubdir = "/preview/";
                var files = await Settings.Instance.Host.GetValue<string[]>("input-file");
                // only evaluate composite builds if all swaggers are preview as we don't have a well-defined model for mixed preview/stable swaggers
                if (files.All(file => file.IndexOf(previewSubdir) > -1) &&
                    Settings.Instance.Host.GetValue<string>("output-folder").Result.IndexOf(previewSubdir) == -1)
                {
                    throw new InvalidOperationException($"codegen for preview swagger {files[0]} must be under a preview subdirectory");
                }
            }

            var codeModel = cm as CodeModelGo;
            if (codeModel == null)
            {
                throw new Exception("Code model is not a Go Code Model");
            }

            // unfortunately there is an ordering issue here.  during model generation we might
            // flatten some types (but not all depending on type).  then during client generation
            // the validation codegen needs to know if a type was flattened so it can generate
            // the correct code, so we need to generate models before clients.

            // Models
            var modelsTemplate = new ModelsTemplate
            {
                Model = codeModel
            };
            await Write(modelsTemplate, FormatFileName("models"));

            // Service client
            var serviceClientTemplate = new ServiceClientTemplate
            {
                Model = codeModel
            };

            await Write(serviceClientTemplate, FormatFileName("client"));

            // by convention the methods in the method group with an empty
            // name go into the client template so skip them here.
            HashSet<string> ReservedFiles = new HashSet<string>(StringComparer.OrdinalIgnoreCase)
            {
                "models",
                "client",
                "version",
            };

            foreach (var methodGroup in codeModel.MethodGroups.Where(mg => !string.IsNullOrEmpty(mg.Name)))
            {
                if (ReservedFiles.Contains(methodGroup.Name.Value))
                {
                    methodGroup.Name += "group";
                }
                ReservedFiles.Add(methodGroup.Name);
                var methodGroupTemplate = new MethodGroupTemplate
                {
                    Model = methodGroup
                };
                await Write(methodGroupTemplate, FormatFileName(methodGroup.Name).ToLowerInvariant());
            }

            // Version
            var versionTemplate = new VersionTemplate { Model = codeModel };
            await Write(versionTemplate, FormatFileName("version"));
        }

        private string FormatFileName(string fileName)
        {
            return $"{fileName}{ImplementationFileExtension}";
        }
    }
}
