local challenges = import 'chals/challenges.libsonnet';
local infras = import 'infra/challenges.libsonnet';
local utils = import 'infra/jsonnet/utils.libsonnet';

local combined = challenges + infras;

std.flattenArrays([utils.extractFlags(chal) for chal in combined])
