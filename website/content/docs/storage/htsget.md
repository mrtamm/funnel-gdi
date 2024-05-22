---
title: Htsget Storage
menu:
  main:
    parent: Storage
---

# Htsget Storage

Funnel supports content-retrieval using [Htsget][spec]-compatible API, if the
host environment has [htsget](htsget-client) and [crypt4gh](crypt4gh)
(including `crypt4gh-keygen`) software installed.
(These programs are not part of Funnel itself.)

Htsget is a protocol that enables downloading only specific parts of genomic
data (reads/variants). The first HTTP query receives a JSON that instructs next
HTTP requests for fetching the parts. Finally the parts need to be concatenated
(in the order they were specified) into a single valid file (e.g. VCF or BAM).
Note that the htsget storage supports only retrieval and not storing the data!

The task input file URL needs to specify `htsget` as the protocol. Funnel
replaces it with the protocol specified in the configuration (default is
`https`).

If the service expects a `Bearer` token, it can be specified in the URL.
For example: `htsget://bearer:your-token-here@fakedomain.com/...`.
Here the `bearer:` part is the required syntax to active the `your-token-here`
value to be sent to the htsget-service as a header value:
`Authorization: Bearer your-token-here`.

If the htsget-service expects the client (Funnel) to send its public key
(crypt4gh), the `SendPublicKey` option must be set to `true` in the
configuration. In this scenario, Funnel will generate a local key-pair and
send its public key in the `client-public-key` header value. Htsget-service is
expected to send the content encrypted with the public key, and Funnel will
decrypt the data locally using `crypt4gh`.

```yaml
HTSGETStorage:
  Disabled: false
  Protocol: https
  SendPublicKey: false
```

### Example task
```
{
  "name": "Hello world",
  "inputs": [{
    "url": "htsget://fakedomain.com/variants/genome2341?referenceName=1&start=10000&end=20000",
    "path": "/inputs/genome.vcf.gz"
  }],
  "outputs": [{
    "url": "file:///path/to/funnel-data/output.txt",
    "path": "/outputs/out.txt"
  }],
  "executors": [{
    "image": "alpine",
    "command": [
      "sh",
      "-c",
      "zcat /inputs/genome.vcf.gz | wc -l"
    ],
    "stdout": "/outputs/out.txt",
  }]
}
```

[spec]: https://samtools.github.io/hts-specs/htsget.html
[htsget-client]: https://htsget.readthedocs.io/en/latest/
[crypt4gh]: https://crypt4gh.readthedocs.io/en/latest/
