local challenges = import 'chals/challenges.libsonnet';
local utils = import 'infra/jsonnet/utils.libsonnet';

local services = std.flattenArrays([utils.extractServices(chal) for chal in challenges]);

local tservices = {
  [service.name]: {
    image: "gcr.io/geegle/chals/%s/%s:latest" % [service.category, service.name],
    networks: {
      ["beyondcorp_" + service.name]: {
        aliases: [
          service.name + ".corp.geegle.org"
        ]
      }
    },
    dns_search: [
      "corp.geegle.org",
      "geegle.org",
    ]
  }
  for service in services
};

local networks = {
  ["beyondcorp_" + services[i].name]: {
    ipam: {
      driver: "default",
      config: [
        {
          "subnet": "100.88.66.%d/29" % [i * 8],
        }
      ],
    }
  } for i in std.range(0, std.length(services) - 1)
};

{
  version: "2",
  services: {
    uberproxy: {
      image: "gcr.io/geegle/infra/uberproxy:latest",
      networks: {
        ["beyondcorp_" + service.name]: {}
        for service in services
      }
    },
  } + tservices,
  networks: networks,
}
