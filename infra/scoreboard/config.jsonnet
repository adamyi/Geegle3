local emails = import 'emails.libsonnet';
local flags = import 'flags.libsonnet';
local config = import 'config.json';

config + {Challenges: emails, Flags: flags}
