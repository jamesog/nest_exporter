# Nest Exporter

This is a Prometheus exporter for Nest products, using the Starling Home Hub API.

The Works with Nest and Google Device Access APIs are not implemented, but should be fairly straightforward to add.

## Usage

1. Follow the [Hub Setup instructions](https://sidewinder.starlinghome.io/sdc/) from Starling to enable the API and create an API key on your Home Hub.
2. Set the API key in your environment as `STARLING_API_KEY`.
3. Run the exporter:

       nest_exporter --starling.api http://192.168.4.10:3080/api/connect/v1

   Where `192.168.4.10` is the IP address of your Home Hub. Note that you need to give the full base URL, including the base path.
   
   The API URL can be set in the environment as `STARLING_API_URL` instead of using the flag. If both are set, the flag overrides the environment variable.

The exporter binds to `:3081` by default. This can be changed with the `--listen` flag.

## Implemented Devices

- [x] Nest Thermostat
- [ ] Nest Temperature Sensor
- [ ] Nest Protect
- [ ] Nest Camera
- [ ] Nest Guard
- [ ] Nest x Yale Lock
- [ ] Nest Weather Service
