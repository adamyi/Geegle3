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

### Dockerfile (Deprecated)
To build docker files
./docker-build.sh    

## Running 

To run
docker-compose up

## Scenario 1: Work Experience
Work experience time! You've just started a work experience placement at one of Australia's biggest tech companies, Geegle. Famous for their products, Geegle have a solution for everything. Welcome to day 1 of the security experience program - your training starts now.


## Scenario 2: New Hire
Welcome to the Security Response Division, newbie. We're responsible for overseeing all other security teams in Geegle, so you better have a way of getting your head across everything. Keep an eye out too - we've had rumours that one of your new cohort may be compromised.

However you got here - work experience program, grad hire, career step - welcome to the team.

---

## Progression
https://docs.google.com/spreadsheets/d/15xOhZdRnNxNbSMNUSxPG_8K92lHa4z5SKJWPPTy5tAc/edit
