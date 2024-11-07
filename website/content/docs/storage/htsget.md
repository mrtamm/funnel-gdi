---
title: Htsget Storage
menu:
  main:
    parent: Storage
---

# Htsget Storage

Funnel supports content-retrieval from an [Htsget][htsget]-compatible API.
When the received content is encrypted using [Crypt4gh][crypt4gh], Funnel
automatically decrypts the received content (using internally generated
key-pair) so that the executor wouldn't have to.

Htsget is a protocol that enables downloading only specific parts of genomic
data (reads/variants). The first HTTP query receives a JSON that instructs next
HTTP requests for fetching the parts. Finally the parts need to be concatenated
(in the order they were specified) into a single valid file (e.g. VCF or BAM).
Note that the Htsget storage supports only retrieval, and not storing the data!

The task input file URL needs to specify `htsget` as the resource protocol.
Funnel will extract path information from the specified URL and append it to
the service URL specified in the Funnel configuration. The format of the input
data URL is following:

```
htsget://reads/<resource/path><?htsget_params>
htsget://variants/<resource/path><?htsget_params>
```

As valid examples:
1. `htsget://reads/DATASET_2000/synthetic-bam?class=header`
2. `htsget://variants/DATASET_2000/synthetic-vcf?referenceName=chr20`


If the service expects a `Bearer` (token) or `Basic` (username:password)
authentication, it can be specified at the end of the URL right after the
hash-sign (`#`). For example: `htsget://variants/file?class=header#user:pass`.
Note that when the task is submitted to Funnel using a valid `Bearer` token for
user authentication, the same token will be automatically appended to the
htsget URL, so the request to the HTSGET service would use the same token.
Exception is when the URL already specifies the hash-sign (`#`) – then it won't
be updated.

Funnel always sends its public key in the header (`client-public-key`) of the
request to the Htsget service. When the Htsget service supports [the content
encryption using Crypt4gh][htsget-crypt4gh], the service can generate a custom
Crypt4gh file header containing the symmetric key for decrypting the referred
content (Crypt4gh formatted data-blocks). Funnel checks the beginning of the
received content to know whether Crypt4gh decryption can be applied. Therefore,
tasks always receive the data decrypted. For sensitive data, the deployment
environment (server) should pay attention to restricting access to the Funnel's
data directories, possibly having separate Funnel instances for different
data-projects.

Htsget Storage configuration just requires a service URL to become active:

```yaml
HTSGETStorage:
  ServiceURL: https://example.org:8443/htsget/
  Timeout: 30s
```

If it's necessary to hard-code a fixed `Basic` authentication user, it can be
specified as `https://some-user:some-pass@example.org:8443/htsget/`.

If the `ServiceUrl` is undefined, `htsget` protocol will be disabled.

### Example task

```json
{
  "name": "Hello world",
  "inputs": [{
    "url": "htsget://variants/genome2341?referenceName=1&start=10000&end=20000",
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

[htsget]: https://samtools.github.io/hts-specs/htsget.html
[crypt4gh]: http://samtools.github.io/hts-specs/crypt4gh.pdf
[htsget-crypt4gh]: https://github.com/umccr/htsget-rs/blob/crypt4gh/docs/crypt4gh/ARCHITECTURE.md