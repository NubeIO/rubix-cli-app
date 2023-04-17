# CHANGELOG

## [v0.6.5](https://github.com/NubeIO/rubix-edge/tree/v0.6.5) (2022-04-17)

- Upgrade lib-systemctl-go to v0.3.0 for the monotonic timestamps

## [v0.6.4](https://github.com/NubeIO/rubix-edge/tree/v0.6.4) (2022-03-24)

- Updates to get system info

## [v0.6.3](https://github.com/NubeIO/rubix-edge/tree/v0.6.3) (2022-03-23)

- Add cron for restarting services
- Add new API to get system info

## [v0.6.2](https://github.com/NubeIO/rubix-edge/tree/v0.6.2) (2022-02-21)

- Exclude socat files & improvement on zip file name
- Exclude nubeio-rubix-edge-bios.service
- Don't run CopyDir in parallel (#65)
- Upgrade ubuntu-20.04 as runner (ubuntu-18.04 is discontinued)

## [v0.6.1](https://github.com/NubeIO/rubix-edge/tree/v0.6.1) (2022-02-16)

- Exclude rubix-assist from snapping
- Remove apps before adding in (#62)
- Attach arch in snapshot filename
- Add validation on restore

## [v0.6.0](https://github.com/NubeIO/rubix-edge/tree/v0.6.0) (2022-02-15)

- Add create and restore snapshots API endpoints

## [v0.5.9](https://github.com/NubeIO/rubix-edge/tree/v0.5.9) (2022-02-14)

- Create public device-info getter API

## [v0.5.8](https://github.com/NubeIO/rubix-edge/tree/v0.5.8) (2022-02-13)

- delete auth headers in chirpstack proxy

## [v0.5.7](https://github.com/NubeIO/rubix-edge/tree/v0.5.7) (2022-02-13)

- added chirpstack proxy

## [v0.5.6](https://github.com/NubeIO/rubix-edge/tree/v0.5.6) (2022-01-27)

- updates to get stream logs

## [v0.5.5](https://github.com/NubeIO/rubix-edge/tree/v0.5.5) (2022-01-25)

- Protect `/api/system/device` API by Auth
    - we use this API to ping and render the device status
- Fix: logs APIs

## [v0.5.4](https://github.com/NubeIO/rubix-edge/tree/v0.5.4) (2022-01-14)

- Fix: sudo: executable file not found in $PATH

## [v0.5.3](https://github.com/NubeIO/rubix-edge/tree/v0.5.3) (2022-01-11)

- added enable/disable npt

## [v0.5.2](https://github.com/NubeIO/rubix-edge/tree/v0.5.2) (2022-01-10)

- Update to networking

## [v0.5.1](https://github.com/NubeIO/rubix-edge/tree/v0.5.1) (2022-12-12)

- Remove suffix slash (/) from APIs for to support reverse proxy
- Set ubuntu-18.04 as the runner OS & update packages

## [v0.5.0](https://github.com/NubeIO/rubix-edge/tree/v0.5.0) (2022-11-24)

- Upgrade files, dirs, zip, systemctl, syscall APIs
- Remove lib-rubix-installer and upgrade lib-files
- Upgrade rubix-registry-go for device_type field
- Remove unused APIs

## [v0.4.0](https://github.com/NubeIO/rubix-edge/tree/v0.4.0) (2022-11-13)

- Add auth handler back

## [v0.3.2](https://github.com/NubeIO/rubix-edge/tree/v0.3.2) (2022-10-26)

- Added edge-proxy

## [v0.3.1](https://github.com/NubeIO/rubix-edge/tree/v0.3.1) (2022-10-16)

- Upgrade lib-rubix-installer to version v0.3.1 to fix wires installation

## [v0.3.0](https://github.com/NubeIO/rubix-edge/tree/v0.3.0) (2022-09-22)

- Lots of improvements

## [v0.2.0](https://github.com/NubeIO/rubix-edge/tree/v0.2.0) (2022-08-29)

- Added networking apis

## [v0.1.7](https://github.com/NubeIO/rubix-edge/tree/v0.1.7) (2022-08-24)

- Fix install file permission

## [v0.1.6](https://github.com/NubeIO/rubix-edge/tree/v0.1.6) (2022-08-22)

- added all the system api, for eth, date and timezones

## [v0.1.3](https://github.com/NubeIO/rubix-edge/tree/v0.1.3) (2022-08-14)

- added apis for plugins

## [v0.1.2](https://github.com/NubeIO/rubix-edge/tree/v0.1.2) (2022-08-12)

- added api for delete an app

## [v0.1.1](https://github.com/NubeIO/rubix-edge/tree/v0.1.1) (2022-08-11)

- stop service on install app

## [v0.1.0](https://github.com/NubeIO/rubix-edge/tree/v0.1.0) (2022-08-11)

- Got install of apps working from assist

## [v0.0.9](https://github.com/NubeIO/rubix-edge/tree/v0.0.9) (2022-08-09)

- First initial release for rubix-service installable
