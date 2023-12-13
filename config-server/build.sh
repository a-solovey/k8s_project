#!/bin/sh

gradle --build-cache --parallel -Dorg.gradle.workers.max="$(getconf _NPROCESSORS_ONLN)" clean build
