# cleantone

<img width="516" alt="image" src="https://github.com/kadai0308/cleantone/assets/24975318/c1b1168a-7164-44b8-9ac9-c33bbb7310c2">

<img width="590" alt="Screen Shot 2023-05-25 at 8 35 35 PM" src="https://github.com/kadai0308/cleantone/assets/24975318/29bb7042-5511-4fb4-a607-48932cb318ca">
<br/>
<br/>

clonetone, which is made with Golang, offers a user-friendly and easy-to-configure key-value store system with the following features:

- It incorporates an append log-based persistence system to ensure fast write operations.
- It automatically prunes data files to minimize disk usage.
- It is easily configurable, allowing users to adjust the size of each data file and the frequency of pruning.
- It supports several data persistence formats, including CSV, JSON(WIP), and Protobuf(WIP).

## Performance Report

<img width="601" alt="image" src="https://github.com/kadai0308/cleantone/assets/24975318/545335b5-c359-4479-9278-4d8bc0925dc4">

Read will be the same because both are from memory (use map to do index)

## Getting started

### Installation

### Example

## License

cleantone is under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.
