# go-torrent

Tiny BitTorrent client written in Go.

## Install

```sh
go get github.com/amdzy/go-torrent
```

## Usage

Try downloading [Debian](https://cdimage.debian.org/debian-cd/current/amd64/bt-cd/#indexlist)!

```sh
go-torrent debian-10.2.0-amd64-netinst.iso.torrent
```

## Limitations

- Only supports `.torrent` files (no magnet links)
- Only supports HTTP trackers
- Does not support multi-file torrents
- Strictly leeches (does not support uploading pieces)
