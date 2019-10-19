local challenges = import 'chals/challenges.libsonnet';
local utils = import 'infra/jsonnet/utils.libsonnet';

std.flattenArrays([utils.extractFlags(chal) for chal in challenges])
