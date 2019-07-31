#!/bin/sh
grep -Po '[A-Z0-9]+\{.{5,30}\}' output.log
