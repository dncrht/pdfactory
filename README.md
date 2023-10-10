# PDFactory

A late jump on the microservices bandwagon (it's been dead for years, hasn't it?), PDFactory allows you to detach PDF generation from your main codebase.

The main motivation that drove me to write this service was the humongous size of the Rails app I work with. The deploy slug is 414 MB. And I realised 288 MB of those belong to `wkhtmltopdf`, a gem for PDF generation.

By moving PDF generation to a fast, small, dedicated external service I could keep my app slug size on a healthy 129 MB mark.

After considering several languages that could fit the purpose and I also could learn on a weekend, I chose Go.

## Installation

1. Clone the repo.
2. Configure an app on Heroku.
3. Add wkhtmltopdf binary on Heroku.
```bash
heroku buildpacks:add https://github.com/dscout/wkhtmltopdf-buildpack.git
heroku config:set WKHTMLTOPDF_VERSION="0.12.4"
```
4. Set ENV vars `USER` and `PASSWORD` to restrict access to your server. Do not set them to leave it open to the whole Internet.
5. Push to deploy.

Refer to [Heroku](https://devcenter.heroku.com/categories/go-support) for extra details and troubleshooting.

## Usage

Make a POST HTTP request to your newly deployed server, eg:

```
https://yourapp.herokuapps.com/pdf
```

Note the path is `/pdf`. You might also need HTTP credentials if you did set them at step 4 above.

Pass a `html` parameter with the HTML code you want to convert to PDF.

If all goes well, response code will be 200 and the server will reply with the PDF as Base64 encoded text.

In case of error, response code will be 422.

Alternatively, you may want to use this [Ruby library](https://github.com/dncrht/pdfactory-client).

## Development

Clone locally and install dependencies:

```bash
go mod vendor
brew install wkhtmltopdf
```

Then compile and run locally:

```bash
./build_n_serve
```

And open http://localhost:4000

## Contributing

Tests are insufficient, but can be run via:
`go test`

Bug reports and pull requests are welcome on GitHub at https://github.com/dncrht/pdfactory.

## License

This codebase is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
