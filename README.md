# Worker

[ ![Codeship Status for site-builder/worker][3]](https://codeship.com/projects/75592)

A Go app that fetches a shallow copy of a git repository, runs it through [Jekyll][4], then pushes it back up to another git repository.

### Getting started

#### Install dependencies

```sh

go get \
  "github.com/onsi/ginkgo/ginkgo" \
  "github.com/onsi/gomega" \
  "github.com/azer/logger" \
  "github.com/satori/go.uuid"

```

### Running Specs

This app uses [Ginko][1] and [Gomega][2] for specs.

```sh

$ ginkgo -r

```

### Running the app

```sh

$ go run main.go \
  -source-branch=master \
  -source-repo=git@github.com:site-builder/test-site.git \
  -destination-branch=gh-pages \
  -destination-repo=git@github.com:site-builder/test-site.git

```

### Running the app with logging

```sh

$ LOG=* go run main.go \
  -source-branch=master \
  -source-repo=git@github.com:site-builder/test-site.git \
  -destination-branch=gh-pages \
  -destination-repo=git@github.com:site-builder/test-site.git

```

### Installing the app locally

```sh

$ go install

```

### Running after installing

```sh

$ LOG=* worker \
  -source-branch=master \
  -source-repo=git@github.com:site-builder/test-site.git \
  -destination-branch=gh-pages \
  -destination-repo=git@github.com:site-builder/test-site.git

```

[1]: http://onsi.github.io/ginkgo/
[2]: http://onsi.github.io/gomega/
[3]: https://codeship.com/projects/b3384680-cac8-0132-98c9-22c60209e864/status?branch=master
