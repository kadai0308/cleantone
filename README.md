# cleantone

[![Go Report Card](https://goreportcard.com/badge/github.com/etcd-io/etcd?style=flat-square)]()
[![Tests](https://github.com/etcd-io/etcd/actions/workflows/tests.yaml/badge.svg)]()
[![Docs](https://img.shields.io/badge/docs-latest-green.svg)]()
[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)]()

<img width="590" alt="Screen Shot 2023-05-25 at 8 35 35 PM" src="https://github.com/kadai0308/cleantone/assets/24975318/29bb7042-5511-4fb4-a607-48932cb318ca">
<br/>
<br/>

clonetone, which is made with Golang, offers a user-friendly and easy-to-configure key-value store system with the following features:

- It incorporates an append log-based persistence system to ensure fast write operations.
- It automatically prunes data files to minimize disk usage.
- It is easily configurable, allowing users to adjust the size of each data file and the frequency of pruning.
- It supports several data persistence formats, including CSV, JSON(WIP), and Protobuf(WIP).

### License

cleantone is under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.
