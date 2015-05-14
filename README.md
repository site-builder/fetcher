# Fetcher 
__Work in progress..__

[ ![Codeship Status for site-builder/fetcher](https://codeship.com/projects/b3384680-cac8-0132-98c9-22c60209e864/status?branch=master)](https://codeship.com/projects/75592)

A Go app for fetching a shallow copy of a git repository and
saving it to an Amazon S3 bucket.

### Running Specs

Grab the [Ginkgo][1] test frameworks and the [Gomega][2] matcher
library.

    $ go get github.com/onsi/ginkgo/ginkgo
    $ go get github.com/onsi/gomega

Run the specs

    $ ginkgo -r

### Running the app

    $ go run main.go

[1]: http://onsi.github.io/ginkgo/
[2]: http://onsi.github.io/gomega/
