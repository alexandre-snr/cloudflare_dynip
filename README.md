# cloudflare_dynip
A lightweight Golang program to update **Cloudflare**'s DNS record when using a **dynamic IP** address.
It is designed to be run in a Docker container and is lightweight (**13Mb** docker image).

## Installation

The first step is to get an API token from Cloudflare:

1. Go to [My Profile/API Tokens](https://dash.cloudflare.com/profile/api-tokens)
2. Click `Create Token`
3. Use the `Edit zone DNS` template
4. In the `Zone Resources` category, select the domain you want to use
5. Copy the generated token for later

You can now run cloudflare_dynip using the docker run:

```bash
docker run -d -e DOMAIN=example.com -e API_TOKEN=yourkey cloudflare_dynip
```

The check will be run every 30 minutes.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)
