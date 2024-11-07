---
title: SDA Storage
menu:
  main:
    parent: Storage
---

# Sensitive Data Archive (SDA) Storage

Funnel supports content-retrieval from an [SDA-Download][sda]-compatible API.
This is a service with an HTTP-based REST-API with some additions:

1. The request must be authenticated using a Bearer token, a JSON Web Token to
   be validated by SDA, checking it would contain a valid GA4GH Visa permitting
   access to the targeted dataset.
2. If the targeted file is encrypted using [Crypt4gh](crypt4gh) (it has ".c4gh"
   extension), the client (e.g. Funnel) needs to send its public key so that
   SDA would reencrypt the file header, which would enable to obtain the cipher
   key from the header using its private key. Funnel makes the file accessible
   to the computation task without encryption.

The task input file URL needs to specify `sda` as the resource protocol.
Funnel will extract path information from the specified URL and append it to
the service URL specified in the Funnel configuration. The format of the input
data URL is following:

```
sda://<dataset-id>/<resource/path>
```

For example: `sda://DATASET_2000/synthetic/sample.bam`


If the service expects a `Bearer` (token) or `Basic` (username:password)
authentication, it can be specified at the end of the URL right after the
hash-sign (`#`). For example: `sda://dataset/file#jwt-token-here`.
Note that when the task is submitted to Funnel using a valid `Bearer` token for
user authentication, the same token will be automatically appended to the
SDA URL, so the request to the SDA service would use the same token.
Exception is when the URL already specifies the hash-sign (`#`) – then it won't
be updated.

Funnel sends its Crypt4gh public key in the header (`client-public-key`) of the
request to the SDA service, when the requested file has ".c4gh" extension.

For sensitive data, the deployment environment (server) should pay attention to
restricting access to the Funnel's data directories, possibly having separate
Funnel instances for different data-projects.

SDA Storage configuration just requires a service URL to become active:

```yaml
SDAStorage:
  ServiceURL: https://example.org:8443/sda/
  Timeout: 30s
```

If the `ServiceUrl` is undefined, `sda` protocol will be disabled.

### Example task

```json
{
  "name": "Hello world",
  "inputs": [{
    "url": "sda://DATASET-2024-012345/variants/genome2341.vcf.gz",
    "path": "/inputs/genome.vcf.gz"
  }],
  "outputs": [{
    "url": "file:///results/line_count.txt",
    "path": "/outputs/line_count.txt"
  }],
  "executors": [{
    "image": "alpine",
    "command": [
      "sh",
      "-c",
      "zcat /inputs/genome.vcf.gz | wc -l"
    ],
    "stdout": "/outputs/line_count.txt"
  }]
}
```

[sda]: https://github.com/neicnordic/sensitive-data-archive/blob/main/sda-download/api/api.md
[crypt4gh]: http://samtools.github.io/hts-specs/crypt4gh.pdf
