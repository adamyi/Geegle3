# Geegle3

Monorepo for COMP9301 CTF.

---

## Encryption
Secret keys and configurations are encrypted using [git-crypt](https://github.com/AGWA/git-crypt)

Please download `g3.key` from https://drive.google.com/open?id=1vRF2AqRcSQQ-aYQCh3uGaTKbhUtwbZgs 
and use `git-crypt unlock PATH_TO_g3.key`

## Building docker files

### Bazel (Experimental)
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

Please do not push to GCR manually

## Running 

### Master server
```
bazel build //infra/jsonnet:cluster-master-docker-compose
docker-compose -f dist/bin/infra/jsonnet/cluster-master-docker-compose.json up -d
```

### Team server
```
bazel build //infra/jsonnet:cluster-team-docker-compose
docker-compose -f dist/bin/infra/jsonnet/cluster-team-docker-compose.json up -d
```

### Test server
```
bazel build //infra/jsonnet:all-docker-compose
docker-compose -f dist/bin/infra/jsonnet/all-docker-compose.json up -d
```

---

## Progression
https://docs.google.com/spreadsheets/d/15xOhZdRnNxNbSMNUSxPG_8K92lHa4z5SKJWPPTy5tAc/edit
