# Worker
__Work in progress..__

[ ![Codeship Status for site-builder/worker](https://codeship.com/projects/b3384680-cac8-0132-98c9-22c60209e864/status?branch=master)](https://codeship.com/projects/75592)

A Go app for fetching a shallow copy of a git repository and
saving it to an Amazon S3 bucket.

### Get gom and Install Dependencies

This project uses [gom][3] to manage dependencies.

    $ go get github.com/mattn/gom
    $ gom install

### Running Specs

Grab the [Ginkgo][1] test frameworks and the [Gomega][2] matcher
library.

    $ gom exec ginkgo -r

### Running the app

    $ gom run main.go

### Running the app with logging

    $ LOG=* gom run main.go

[1]: http://onsi.github.io/ginkgo/
[2]: http://onsi.github.io/gomega/
[3]: https://github.com/mattn/gom
