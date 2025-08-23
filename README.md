# ClamAV REST Service

Official [Cisco-Talos/clamav](https://github.com/Cisco-Talos/clamav) docker image with a REST API enhancement for file scanning.

## User Information

API Endpoints:

```
# Scan File for virus:
# returns 200 on success, 400 Bad Request if file not readable and 422 Unprocessable Entity if the file contains a virus
http://<<your-address>>:33779/scan-document
```

| Key                                | Description                                                      | Default-Value |
| ---------------------------------- | ---------------------------------------------------------------- | ------------- |
| CLAMAV_REST_CLAMAV_EXECUTION_PATH  | Path to clamscan execution binary                                | clamdscan     |
| CLAMAV_REST_REST_SCANFILE_FORM_KEY | HTTP form-file key where the file lives inside your HTTP request | file          |

## Developer information

Build a docker image:

```bash
docker build -t clamav-wrapper --platform=linux/amd64 .
```

Run docker image:

```bash
docker run --name clamav-wrapper --rm -p 33779:33779 clamav-wrapper:latest
```
