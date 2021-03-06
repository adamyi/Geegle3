local challenges = import 'chals/challenges.libsonnet';
local infras = import 'infra/challenges.libsonnet';
local utils = import 'infra/jsonnet/utils.libsonnet';

local combined = challenges + infras;

local image(service) = if "image" in service then
  service.image
  else if service.category == "infra" then
    "gcr.io/geegle/infra/%s:latest" % service.name
  else
    "gcr.io/geegle/chals/%s/%s:latest" % [service.category, service.name];

local tmpservices = std.flattenArrays([utils.extractServices(chal) for chal in combined]);

local services = [service for service in tmpservices if ((std.extVar('cluster') == 'all') || ('clustertype' in service && service['clustertype'] == std.extVar('cluster')))];

// local searchdomains = ["corp.geegle.org", "geegle.org"]; // NOTES(adamyi@): they mess up with uberproxy... disable it for now
local searchdomains = [];

local tservices = {
  [services[i].name]: {
    local service = services[i],
    image: image(service),
    networks: {
      ["beyondcorp_" + service.name]: {
        aliases: [
          service.name + if "domain" in service then service.domain else ".corp.geegle.org"
        ],
        ipv4_address: "100.88.66.%d" % [i * 8 + 3],
      }
    },
    dns: "100.88.66.%d" % [i * 8 + 4],
    dns_search: searchdomains,
    ports: if "ports" in service then service.ports else [],
  }
  for i in std.range(0, std.length(services) - 1)
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
    dns: {
      image: "gcr.io/geegle/infra/dns:latest",
      networks: {
        ["beyondcorp_" + services[i].name]: {
          ipv4_address: "100.88.66.%d" % [i * 8 + 4],
        }
        for i in std.range(0, std.length(services) - 1)
      },
    },
    uberproxy: {
      image: "gcr.io/geegle/infra/uberproxy:latest",
      networks: {
        ["beyondcorp_" + services[i].name]: {
          ipv4_address: "100.88.66.%d" % [i * 8 + 2],
        }
        for i in std.range(0, std.length(services) - 1)
      },
      ports: [
        "80:80",
        "443:443",
      ],
      dns_search: searchdomains,
      environment: [
        "UBERPROXY_CLUSTER=" + std.extVar('cluster')
      ],
    },
  } + tservices,
  networks: networks,
}
