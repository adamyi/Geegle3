local challenges = import 'chals/challenges.libsonnet';
local utils = import 'infra/jsonnet/utils.libsonnet';

local addServiceName(files, name) = [file + {service: name} for file in files];

local services = std.flattenArrays([utils.extractServices(chal) for chal in challenges]);

local serviceFiles = std.flattenArrays([addServiceName(utils.extractFiles(service), service.name + "@services.geegle.org" ) for service in services]);

serviceFiles + std.flattenArrays([utils.extractFiles(chal) for chal in challenges])
