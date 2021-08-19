// GeoIP generator
//
// Before running this file, the GeoIP database must be downloaded and present.
// To download GeoIP database: https://dev.maxmind.com/geoip/geoip2/geolite2/
// Inside you will find block files for IPv4 and IPv6 and country code mapping.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/v2fly/v2ray-core/v4/app/router"
	"github.com/v2fly/v2ray-core/v4/common"
	"github.com/v2fly/v2ray-core/v4/infra/conf/rule"
	"google.golang.org/protobuf/proto"
)

var (
	chinaIPFile     = flag.String("chinaip", "china_ip_list.txt", "Path to the IPList for China by IPIP.NET file")
	outputName      = flag.String("outputname", "geoip.dat", "Name of the generated file")
	outputDir       = flag.String("outputdir", "./", "Path to the output directory")
)

var privateIPs = []string{
	"0.0.0.0/8",
	"10.0.0.0/8",
	"100.64.0.0/10",
	"127.0.0.0/8",
	"169.254.0.0/16",
	"172.16.0.0/12",
	"192.0.0.0/24",
	"192.0.2.0/24",
	"192.88.99.0/24",
	"192.168.0.0/16",
	"198.18.0.0/15",
	"198.51.100.0/24",
	"203.0.113.0/24",
	"224.0.0.0/4",
	"240.0.0.0/4",
	"255.255.255.255/32",
	"::1/128",
	"fc00::/7",
	"fe80::/10",
}

func getPrivateIPs() *router.GeoIP {
	cidr := make([]*router.CIDR, 0, len(privateIPs))
	for _, ip := range privateIPs {
		c, err := rule.ParseIP(ip)
		common.Must(err)
		cidr = append(cidr, c)
	}
	return &router.GeoIP{
		CountryCode: "PRIVATE",
		Cidr:        cidr,
	}
}

func getChinaIPs() *router.GeoIP {
	chinaIPReader, err := os.Open(*chinaIPFile)
	common.Must(err)
	defer chinaIPReader.Close()

	chinaIPContent, err := ioutil.ReadAll(chinaIPReader)
	common.Must(err)
	chinaIPs := strings.Split(string(chinaIPContent), "\n")

	cidr := make([]*router.CIDR, 0, len(chinaIPs))
	for _, ip := range chinaIPs {
		c, err := rule.ParseIP(ip)
		common.Must(err)
		cidr = append(cidr, c)
	}
	return &router.GeoIP{
		CountryCode: "CN",
		Cidr:        cidr,
	}
}

func main() {
	flag.Parse()

	geoIPList := new(router.GeoIPList)
	geoIPList.Entry = append(geoIPList.Entry, getPrivateIPs())
	geoIPList.Entry = append(geoIPList.Entry, getChinaIPs())

	geoIPBytes, err := proto.Marshal(geoIPList)
	if err != nil {
		fmt.Println("Error marshalling geoip list:", err)
		os.Exit(1)
	}

	// Create output directory if not exist
	if _, err := os.Stat(*outputDir); os.IsNotExist(err) {
		if mkErr := os.MkdirAll(*outputDir, 0755); mkErr != nil {
			fmt.Println("Failed: ", mkErr)
			os.Exit(1)
		}
	}

	if err := ioutil.WriteFile(filepath.Join(*outputDir, *outputName), geoIPBytes, 0644); err != nil {
		fmt.Println("Error writing geoip to file:", err)
		os.Exit(1)
	} else {
		fmt.Println(*outputName, "has been generated successfully.")
	}
}
