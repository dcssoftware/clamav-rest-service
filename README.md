# ClamAV REST Service

Official [Cisco-Talos/clamav](https://github.com/Cisco-Talos/clamav) docker image with a REST API enhancement for file scanning. ClamAV is contained.

## User Information

API Endpoints:

```
# Scan File for virus:
# returns 200 on success, 400 Bad Request if file not readable and 422 Unprocessable Entity if the file contains a virus
POST http://<<your-address>>:33779/scan-document
```

Example Request (use testfiles -> [Testfiles](https://www.eicar.org/download-anti-malware-testfile/)):

```bash
curl -X POST http://127.0.0.1:33779/scan-document -F "file=@/path/to/your/file.zip" -v
```

Custom ENV configuration options: ("CLAMAV*REST*" is the application prefix)
| Key | Default-Value | Description |
| ----------------------------------- | ------------- | -------------------------------------------- |
| `CLAMAV_REST_CLAMAV_EXECUTION_PATH` | `clamdscan` | Path to clamscan execution binary |
| `CLAMAV_REST_REST_SCANFILE_FORM_KEY` | `file` | HTTP form-file key where the file lives inside your HTTP request |

Beside that options you can pass all other clamav configuration options

### Healthcheck

Do not rely on the rest api for health check purposes because it starts async from claimav. claimav itsself always takes longer

Use [this official clamav healthcheck](https://github.com/Cisco-Talos/clamav/blob/main/README.Docker.md#container-clamd-health-check) instead

## Developer information

Build a docker image:

```bash
docker build -t clamav-wrapper --platform=linux/amd64 .
```

Run docker image:

```bash
docker run --name clamav-wrapper --rm -p 33779:33779 clamav-wrapper:latest
```

## General Information

This application is still in beta and developed in one day, if you find errors, please report them!
