# Introduction

`tlsx` was a private library I was using to analyse TLS Client Hello messages sent by browsers.

I didn't continue to the project, but others asked about it, so I thought I'd open source it without warranty.

The library requires the TCP payload of a TLS Client Hello message, which can be provided by
(gopacket)[https://github.com/google/gopacket] (see example).

*This program is not used internally by myself anymore, and may not have an updated list of ciphers, extensions etc. But
it may work for you.*

# Example usage

Run an example program, which listens on an interface for inbound or outbound packets on port 443.

```
cd example; go build; sudo ./example -iface br0
```

Make a connection:

```
curl https://www.google.com.au
```

Review the output:

```
2017/06/24 21:30:48 Client hello from port 37066 to 443(https)
Version: TLS 1.0
Handshake Type: 1
Handshake Version: TLS 1.2
SessionID: []byte(nil)
Cipher Suites (54): [TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384 TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA
TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256 TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384
TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256 TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA
TLS_DHE_RSA_WITH_AES_256_GCM_SHA384 TLS_DHE_RSA_WITH_AES_256_CBC_SHA TLS_DHE_DSS_WITH_AES_256_CBC_SHA
TLS_DHE_RSA_WITH_AES_256_CBC_SHA256 TLS_DHE_RSA_WITH_AES_128_GCM_SHA256 TLS_DHE_RSA_WITH_AES_128_CBC_SHA
TLS_DHE_DSS_WITH_AES_128_CBC_SHA TLS_DHE_RSA_WITH_AES_128_CBC_SHA256 TLS_DHE_RSA_WITH_3DES_EDE_CBC_SHA
TLS_DHE_DSS_WITH_3DES_EDE_CBC_SHA TLS_RSA_WITH_AES_256_GCM_SHA384 TLS_RSA_WITH_AES_256_CBC_SHA
TLS_RSA_WITH_AES_256_CBC_SHA256 TLS_RSA_WITH_AES_128_GCM_SHA256 TLS_RSA_WITH_AES_128_CBC_SHA
TLS_RSA_WITH_AES_128_CBC_SHA256 TLS_RSA_WITH_3DES_EDE_CBC_SHA TLS_RSA_WITH_RC4_128_SHA TLS_RSA_WITH_RC4_128_MD5]
Compression Methods: [0]
Extensions: map[renegotiation_info:1]
SNI: "google.com.au"
Signature Algorithms: []uint16{0x403, 0x503, 0x603, 0x203, 0x401, 0x501, 0x601, 0x201, 0x402, 0x502, 0x602, 0x202}
Groups: []uint16{0x17, 0x18, 0x19}
Points: []byte{0x0}
OSCP: false
ALPNs: []
```

Chrome example:

```
2017/06/24 21:26:32 Client hello from port 61746 to 443(https)
Version: TLS 1.0
Handshake Type: 1
Handshake Version: TLS 1.2
SessionID: []byte(nil)
Cipher Suites (32): [0x8a8a (unknown) TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256 TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256 TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384 TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384 0xcca9 (unknown) 0xcca8 (unknown) 0xcc14 (unknown) 0xcc13 (unknown) TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA TLS_RSA_WITH_AES_128_GCM_SHA256 TLS_RSA_WITH_AES_256_GCM_SHA384 TLS_RSA_WITH_AES_128_CBC_SHA TLS_RSA_WITH_AES_256_CBC_SHA TLS_RSA_WITH_3DES_EDE_CBC_SHA]
Compression Methods: [0]
Extensions: map[0x7550 (unknown):0 0x7a7a (unknown):1 0xfafa (unknown):0 renegotiation_info:1 extended_master_secret:0 SessionTicket TLS:0 signed_certificate_timestamp:0]
SNI: "example.com"
Signature Algorithms: []uint16{0x403, 0x804, 0x401, 0x503, 0x805, 0x501, 0x806, 0x601, 0x201}
Groups: []uint16{0x6a6a, 0x1d, 0x17, 0x18}
Points: []byte{0x0}
OSCP: true
ALPNs: [h2 http/1.1]
```
