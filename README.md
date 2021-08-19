# Light GeoIP List for V2Ray

Light version of `geoip.dat` for V2Ray.

Contains `cn` and `private` ONLY.

## Download links

- **geoip.dat**：[https://github.com/Mukou-Aoi/geoip_cn_private/releases/latest/download/geoip.dat](https://github.com/Mukou-Aoi/geoip_cn_private/releases/latest/download/geoip.dat)
- **geoip.dat.sha256sum**：[https://github.com/Mukou-Aoi/geoip_cn_private/releases/latest/download/geoip.dat.sha256sum](https://github.com/Mukou-Aoi/geoip_cn_private/releases/latest/download/geoip.dat.sha256sum)

## Usage example

```json
"routing": {
  "rules": [
    {
      "type": "field",
      "outboundTag": "Direct",
      "ip": [
        "geoip:cn",
        "geoip:private"
      ]
    }
  ]
}
```

## Notice

This project is a fork of [GeoIP List for V2Ray](https://github.com/v2fly/geoip)

This project uses [IPList for China by IPIP.NET](https://github.com/17mon/china_ip_list)

## License

[CC-BY-NC-SA-4.0](https://creativecommons.org/licenses/by-nc-sa/4.0/)
