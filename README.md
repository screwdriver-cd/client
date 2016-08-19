# Screwdriver Client

> Screwdriver CLI  
## How to get a swagger.json  
1.  In your client directory run `$ wget <your swagger url here>/swagger.json` in order to get your swagger.json

## Generating the client and models directories  
1.  (If go-swagger is not installed, install it)
2.  In the root directory of client run `$ swagger generate client swagger.json`

## Building from source  
1. Clone down the go-swagger library into github.com/go-swagger/go-swagger  
<<<<<<< HEAD
2. `$ git checkout tags/0.5.0 in the go-swagger library`
3. Ensure that the client and models directories have been generated according to the swagger standard
4. Navigate into screwdriver client run `$ go get` to install dependencies
5. Run `$ go build -o goclient` to create the binary
=======
2. `$ git checkout tags/0.5.0` in the go-swagger directory
3. Navigate into screwdriver client run `$ go get` to install dependencies
4. Run `$ go build` to create the binary
>>>>>>> 1269ef296fd9986c0a81150adcb16b3b1b230e05

## Usage
`$ ./goclient <command-name>`

## License

Code licensed under the BSD 3-Clause license. See LICENSE file for terms.

[downloads-image]: https://img.shields.io/npm/dt/screwdriver-client.svg
[license-image]: https://img.shields.io/npm/l/screwdriver-client.svg
[issues-image]: https://img.shields.io/github/issues/screwdriver-cd/client.svg
[issues-url]: https://github.com/screwdriver-cd/client/issues
[wercker-image]: https://app.wercker.com/status/4f8829ae9447c940abcf8bf283a69889
[wercker-url]: https://app.wercker.com/project/bykey/4f8829ae9447c940abcf8bf283a69889
[daviddm-image]: https://david-dm.org/screwdriver-cd/client.svg?theme=shields.io
[daviddm-url]: https://david-dm.org/screwdriver-cd/client
