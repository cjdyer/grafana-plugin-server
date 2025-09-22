# Grafana Plugin Server

Grafana plugin repository to allow private plugins to be stored and installed to a Grafana instance.

## Getting Started

To run the project use the following `Makefile` targets:

```bash
# Install dependencies
make deps
# Build frontend and backend, then run
make run
```

This will start the Go server and build the assets needed to serve the frontend. Once built, the
server binary is placed into the `./build` directory and the frontend assets are placed into the
`./dist` directory.

### Grafana Config

The following config changes are needed to make the plugin server visible to Grafana:

| `.ini` Config     | Docker Config        | Change To               | Reason                                                                                                            |
| ----------------- | -------------------- | ----------------------- | ----------------------------------------------------------------------------------------------------------------- |
| `grafana_net.url` | `GF_GRAFANA_NET_URL` | `http://localhost:3838` | Points The instance to request from the plugin server instead of Grafana's inbuilt server (`https://grafana.com`) |

Where `localhost:3838` is the location of the Grafana Plugin Server being hosted. This should be changed if the server location is changed.

## Additional Changes (if Grafana is open sourced)

Some changes can be made to the Grafana source code to improve the functionality. These changes are strictly not required to get the server working, but will allow for an increased feature set to be available.

### Update Checking

Prior to [Grafana v11.1.0](https://github.com/grafana/grafana/tree/v11.1.0) the endpoint used to query plugin versions is hard-coded. The code for this can be found [here](https://github.com/grafana/grafana/blob/v11.0.0/pkg/services/updatechecker/plugins.go#L118), and is requesting at `https://grafana.com/api/plugins/versioncheck`.

This means the version check query will not be able to be completed, resulting in the update functionality not being available. If this change is not patched the plugin can be updated by uninstalling and installing again, as the install will use the configured `grafana_net.url` endpoint.

To fix this the endpoint can be hard-coded to the new server location (default: `http://localhost:3838/api/plugins/versioncheck`) or can be updated dynamically using the same fix used by the Grafana team in v11.1.0, the code for this fix can be found [here](https://github.com/grafana/grafana/blob/main/pkg/services/updatemanager/plugins.go#L43).

### Plugin Icons

Prior to [Grafana v11.1.0](https://github.com/grafana/grafana/tree/v11.1.0) the endpoint used to query remote plugin icons is hard-coded to be `https://grafana.com/api/plugins`. The code for this can be found [here](https://github.com/grafana/grafana/blob/v11.0.0/public/app/features/plugins/admin/helpers.ts#L125-L126) and [here](https://github.com/grafana/grafana/blob/v11.0.0/public/app/features/plugins/admin/helpers.ts#L213-L214).

This results in the plugin icons in the Grafana Plugin Server not being displayed when installing. This does not affect plugin functionality and is purely cosmetic for the plugin list page.

To fix this the endpoint can be hard-coded to the new server location (default: `http://localhost:3838/api/plugins`) or can be updated dynamically using the same fix used by the Grafana team in v11.1.0, the code for this fix can be found [here](https://github.com/grafana/grafana/blob/v11.1.0/public/app/features/plugins/admin/helpers.ts#L129C16-L130) and [here](https://github.com/grafana/grafana/blob/v11.1.0/public/app/features/plugins/admin/helpers.ts#L217-L218).
