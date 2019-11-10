# Geegle3

Monorepo for infrastructure and challenges of SECedu CTF 2019.

---

## Infra
Everything is behind the settings of a fictional company, Geegle. Geegle has its own BeyondCorp-like
zero-trust network via [UberProxy](infra/uberproxy). We have a working email server that supports
company internal emails as well as inbound and outbound emails at [geemail-backend](infra/geemail-backend),
[geemail-frontend](infra/geemail-frontend), [geemail-client](infra/geemail-client), [gsmtpd](infra/gsmtpd).
All challenge descriptions are sent to team through emails. Challenge emails are unlocked as players progress
in the CTF. Every challlenge has its own unified configuration file called `challenge.libsonnet`.See
[chals/pwn/geelang/challenge.libsonnet](chals/pwn/geelang/challenge.libsonnet) for an example. Its emails,
container services, static files, flags, etc. are all in that single file. We have a shared server that every
player connects to, as well as separate team servers for each of the team. The `clustertype` in configuration
determines whether a specific service should run on the shared server or in a separate team server. This makes
it possible that some services are shared to facilitate inter-team communication while some services offer
isolation between teams.

Players send their flag to flag@geegle.org to claim points. They can also interact with [xssbot](infra/xssbot)
through company internal emails.

Binary challenges are also tunneled through UberProxy with websocket. See [cli-relay](infra/cli-relay), 
[cli-static](infra/cli-static), and [uberproxy/websocket.go](infra/uberproxy/websocket.go). Static files are served
using shared infra [sffe](infra/sffe), a general-purpose static file front-end on top of SSTable (leveldb).

Other infra services we have include: [scoreboard](infra/scoreboard), [dns](infra/dns) (internal DNS service used by
all containers to help connect to uberproxy), [gaia](infra/gaia) (internal authentication service), [gae](infra/gae)
(a service like Google App Engine and Amazon Lambda), [requestz](infra/requestz) (a simple network debugging service),
[mss](infra/mss) (internal KV databse service integrated with Geegle services authentication).

Everything (Golang, Python, C, TypeScript, Bash, JSONNet, Java, PHP) are built with [bazel](https://bazel.build). Containers
images are pushed to [GCR](https://cloud.google.com/container-registry/), while docker-compose files are auto-generated as well.

## Challenges
See https://docs.google.com/spreadsheets/d/15xOhZdRnNxNbSMNUSxPG_8K92lHa4z5SKJWPPTy5tAc/edit

---

## Running Your Own CTF
If you want to use the same Geegle infrastructure to host your own CTF, we are more than happy to support you. Simply remove
all challenges from [chals](chals) directory and put in your own challenges, and change the root [BUILD](BUILD) file accordingly.

Please do let us know if you use Geegle infra to host your own CTF. We can't wait to hear about the amazing work you have done :)

### SSL Certificates
Please put your HTTPS certificates and keys to [infra/uberproxy/certs/](infra/uberproxy/certs/) and change
[infra/uberproxy/ssl.go](infra/uberproxy/ssl.go) accordingly.

### Building Container Images
Please build using Linux AMD64. Cuz it's hard to set up cross-compiling for C programs on mac, ceebs.

Build only:
```
bazel build //:all_containers
```

Build and tag locally (so that you can use docker-compose to boot them up):
```
bazel run //:all_containers
```

Commits submitted to master branch will be automatically pushed to gcr.io/geegle, our container repo

If you are deploying your own CTF using this infra, please change BUILD file to push to a different
container registry, since gcr.io/geegle is not public.

### Deploying

#### Master Server (Shared Server)
```
bazel build //infra/jsonnet:cluster-master-docker-compose
docker-compose -f dist/bin/infra/jsonnet/cluster-master-docker-compose.json up -d
```

#### Team Server (Separate Isolated Server)
```
bazel build //infra/jsonnet:cluster-team-docker-compose
docker-compose -f dist/bin/infra/jsonnet/cluster-team-docker-compose.json up -d
```

#### Test Server (All-in-one Server)
```
bazel build //infra/jsonnet:all-docker-compose
docker-compose -f dist/bin/infra/jsonnet/all-docker-compose.json up -d
```

---

## LICENSE

Copyright (c) 2019 [Adam Yi](mailto:i@adamyi.com), [Adam Tanana](mailto:adam@tanana.io), [Lachlan Jones](mailto:contact@lachjones.com)

To check the author for an individual challenge/infra service, check [CODEOWNERS](CODEOWNERS).

Open-sourced with love, under [Apache 2.0 License](LICENSE).
