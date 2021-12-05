package bgpview

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) (*Client, *http.ServeMux) {
	t.Helper()

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)

	client := NewClient()
	client.baseURL, _ = url.Parse(server.URL)
	client.HTTPClient = server.Client()

	return client, mux
}

func testHandler(filename string) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			http.Error(rw, fmt.Sprintf("unsupported method: %s", req.Method), http.StatusMethodNotAllowed)
			return
		}

		file, err := os.Open(filepath.Join("fixtures", filename))
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		defer func() { _ = file.Close() }()

		_, err = io.Copy(rw, file)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func TestClient_GetASN(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/asn/61138", testHandler("asn.json"))

	details, err := client.GetASN(context.Background(), 61138)
	require.NoError(t, err)

	expected := &ASNInfo{
		Status:        "ok",
		StatusMessage: "Query was successful",
		Data: ASNData{
			ASN:               61138,
			Name:              "ZAPPIE-HOST-AS",
			DescriptionShort:  "Zappie Host",
			DescriptionFull:   []string{"Zappie Host"},
			CountryCode:       "US",
			Website:           "https://zappiehost.com/",
			EmailContacts:     []string{"abuse@zappiehost.com", "admin@zappiehost.com", "noc@zappiehost.com"},
			AbuseContacts:     []string{"abuse@zappiehost.com"},
			LookingGlass:      "https://lg-nz.zappiehost.com",
			TrafficEstimation: "1-5Gbps",
			TrafficRatio:      "Mostly Outbound",
			OwnerAddress:      []string{"16192 Coastal HWY", "DE 19958", "Lewes", "UNITED STATES"},
			RIRAllocation: AllocationData{
				RIRName:       "RIPE",
				CountryCode:   "US",
				DateAllocated: "2015-03-04 00:00:00",
			},
			DateUpdated: "2021-11-21 04:02:05",
		},
		Meta: Meta{TimeZone: "UTC", APIVersion: 1, ExecutionTime: "24.48 ms"},
	}

	assert.Equal(t, expected, details)
}

func TestClient_GetASNPrefixes(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/asn/61138/prefixes", testHandler("asn-prefixes.json"))

	details, err := client.GetASNPrefixes(context.Background(), 61138)
	require.NoError(t, err)

	expected := &ASNPrefixesInfo{
		Status:        "ok",
		StatusMessage: "Query was successful",
		Data: ASNPrefixesData{
			IPv4Prefixes: []ASNIPPrefixesData{
				{Prefix: "45.67.13.0/24", IP: "45.67.13.0", CIDR: 24, RoaStatus: "None", Name: "QUICKVIRT-PRA01", Description: "QUICKVIRT PRA01", CountryCode: "CZ", Parent: ASNPrefixesParent{Prefix: "45.67.12.0/22", IP: "45.67.12.0", CIDR: 22, RIRName: "RIPE"}},
				{Prefix: "45.146.105.0/24", IP: "45.146.105.0", CIDR: 24, RoaStatus: "None", Name: "", Description: "", CountryCode: "", Parent: ASNPrefixesParent{Prefix: "", IP: "", CIDR: 0, RIRName: ""}},
				{Prefix: "45.155.65.0/24", IP: "45.155.65.0", CIDR: 24, RoaStatus: "None", Name: "Heficed", Description: "UAB Xantho", CountryCode: "GB", Parent: ASNPrefixesParent{Prefix: "45.155.64.0/22", IP: "45.155.64.0", CIDR: 22, RIRName: "RIPE"}},
				{Prefix: "45.155.66.0/24", IP: "45.155.66.0", CIDR: 24, RoaStatus: "None", Name: "Heficed", Description: "UAB Xantho", CountryCode: "GB", Parent: ASNPrefixesParent{Prefix: "45.155.64.0/22", IP: "45.155.64.0", CIDR: 22, RIRName: "RIPE"}},
				{Prefix: "89.117.126.0/24", IP: "89.117.126.0", CIDR: 24, RoaStatus: "None", Name: "LT-LRTC-20060503", Description: "SC \"Lithuanian Radio and TV Center\"", CountryCode: "LT", Parent: ASNPrefixesParent{Prefix: "89.116.0.0/15", IP: "89.116.0.0", CIDR: 15, RIRName: "RIPE"}},
				{Prefix: "103.208.86.0/24", IP: "103.208.86.0", CIDR: 24, RoaStatus: "None", Name: "ZAPPIE-HOST-NZ-3", Description: "Zappie Host - Auckland, New Zealand", CountryCode: "NZ", Parent: ASNPrefixesParent{Prefix: "103.208.84.0/22", IP: "103.208.84.0", CIDR: 22, RIRName: "APNIC"}},
				{Prefix: "104.247.99.0/24", IP: "104.247.99.0", CIDR: 24, RoaStatus: "None", Name: "DNET", Description: "Dnetworks LLC", CountryCode: "US", Parent: ASNPrefixesParent{Prefix: "104.247.99.0/24", IP: "104.247.99.0", CIDR: 24, RIRName: "ARIN"}},
				{Prefix: "144.48.80.0/24", IP: "144.48.80.0", CIDR: 24, RoaStatus: "None", Name: "BITACCEL-NETWORK", Description: "BitAccel", CountryCode: "US", Parent: ASNPrefixesParent{Prefix: "144.48.80.0/22", IP: "144.48.80.0", CIDR: 22, RIRName: "APNIC"}},
				{Prefix: "169.239.128.0/23", IP: "169.239.128.0", CIDR: 23, RoaStatus: "None", Name: "ZAPPIE-HOST-ZA-1", Description: "Zappie Host - Johannesburg, South Africa", CountryCode: "ZA", Parent: ASNPrefixesParent{Prefix: "169.239.128.0/22", IP: "169.239.128.0", CIDR: 22, RIRName: "AfriNIC"}},
				{Prefix: "169.239.130.0/23", IP: "169.239.130.0", CIDR: 23, RoaStatus: "None", Name: "ZAPPIE-HOST-ZA-2", Description: "Zappie Host - Johannesburg, South Africa", CountryCode: "ZA", Parent: ASNPrefixesParent{Prefix: "169.239.128.0/22", IP: "169.239.128.0", CIDR: 22, RIRName: "AfriNIC"}},
				{Prefix: "185.99.132.0/24", IP: "185.99.132.0", CIDR: 24, RoaStatus: "None", Name: "ZAPPIE-HOST-NZ", Description: "Zappie Host - Auckland, New Zealand", CountryCode: "NZ", Parent: ASNPrefixesParent{Prefix: "185.99.132.0/22", IP: "185.99.132.0", CIDR: 22, RIRName: "RIPE"}},
				{Prefix: "185.99.133.0/24", IP: "185.99.133.0", CIDR: 24, RoaStatus: "None", Name: "ZAPPIE-HOST-NZ", Description: "Zappie Host - Auckland, New Zealand", CountryCode: "NZ", Parent: ASNPrefixesParent{Prefix: "185.99.132.0/22", IP: "185.99.132.0", CIDR: 22, RIRName: "RIPE"}},
				{Prefix: "185.121.168.0/24", IP: "185.121.168.0", CIDR: 24, RoaStatus: "None", Name: "ZAPPIE-HOST-NZ", Description: "Zappie Host - Auckland, New Zealand", CountryCode: "NZ", Parent: ASNPrefixesParent{Prefix: "185.121.168.0/22", IP: "185.121.168.0", CIDR: 22, RIRName: "RIPE"}},
				{Prefix: "185.195.239.0/24", IP: "185.195.239.0", CIDR: 24, RoaStatus: "None", Name: "ZAP-NZ-ISP-PREFIX", Description: "Zappie ISP Services - New Zealand", CountryCode: "NZ", Parent: ASNPrefixesParent{Prefix: "185.195.236.0/22", IP: "185.195.236.0", CIDR: 22, RIRName: "RIPE"}},
				{Prefix: "216.73.158.0/24", IP: "216.73.158.0", CIDR: 24, RoaStatus: "None", Name: "ZAP-CID-6006", Description: "Private Customer", CountryCode: "US", Parent: ASNPrefixesParent{Prefix: "216.73.156.0/22", IP: "216.73.156.0", CIDR: 22, RIRName: "ARIN"}},
				{Prefix: "216.73.159.0/24", IP: "216.73.159.0", CIDR: 24, RoaStatus: "None", Name: "ZAPPIE-HOST-CL-1", Description: "Zappie Host - Valdivia, Chile", CountryCode: "CL", Parent: ASNPrefixesParent{Prefix: "216.73.156.0/22", IP: "216.73.156.0", CIDR: 22, RIRName: "ARIN"}},
			},
			IPv6Prefixes: []ASNIPPrefixesData{
				{Prefix: "2404:3d80::/32", IP: "2404:3d80::", CIDR: 32, RoaStatus: "None", Name: "ZAPPIEHOST-AP-20160203", Description: "Zappie Host LLC", CountryCode: "NZ", Parent: ASNPrefixesParent{Prefix: "2404:3d80::/32", IP: "2404:3d80::", CIDR: 32, RIRName: "APNIC"}},
				{Prefix: "2407:c280:b103::/48", IP: "2407:c280:b103::", CIDR: 48, RoaStatus: "None", Name: "DavidLiu", Description: "fixmix Technologies Ltd", CountryCode: "EU", Parent: ASNPrefixesParent{Prefix: "2407:c280::/32", IP: "2407:c280::", CIDR: 32, RIRName: "APNIC"}},
				{Prefix: "2a05:dfc0::/29", IP: "2a05:dfc0::", CIDR: 29, RoaStatus: "None", Name: "US-ZAPPIE-20150303", Description: "Zappie Host LLC", CountryCode: "GB", Parent: ASNPrefixesParent{Prefix: "2a05:dfc0::/29", IP: "2a05:dfc0::", CIDR: 29, RIRName: "RIPE"}},
				{Prefix: "2a06:1280::/29", IP: "2a06:1280::", CIDR: 29, RoaStatus: "None", Name: "US-ZAPPIE-20150507", Description: "Zappie Host LLC", CountryCode: "BY", Parent: ASNPrefixesParent{Prefix: "2a06:1280::/29", IP: "2a06:1280::", CIDR: 29, RIRName: "RIPE"}},
				{Prefix: "2a06:1280::/32", IP: "2a06:1280::", CIDR: 32, RoaStatus: "None", Name: "ZAPPIE-HOST-NZ-v6", Description: "Zappie Host - Auckland, New Zealand v6", CountryCode: "NZ", Parent: ASNPrefixesParent{Prefix: "2a06:1280::/29", IP: "2a06:1280::", CIDR: 29, RIRName: "RIPE"}},
				{Prefix: "2a06:1280:ae02::/48", IP: "2a06:1280:ae02::", CIDR: 48, RoaStatus: "None", Name: "ZAPPIE-HOST-NZ-v6", Description: "Zappie Host - Auckland, New Zealand v6", CountryCode: "NZ", Parent: ASNPrefixesParent{Prefix: "2a06:1280::/29", IP: "2a06:1280::", CIDR: 29, RIRName: "RIPE"}},
				{Prefix: "2a06:1280:aee1::/48", IP: "2a06:1280:aee1::", CIDR: 48, RoaStatus: "None", Name: "", Description: "", CountryCode: "US", Parent: ASNPrefixesParent{Prefix: "2a06:1280::/29", IP: "2a06:1280::", CIDR: 29, RIRName: "RIPE"}},
				{Prefix: "2a06:1280:c0de::/48", IP: "2a06:1280:c0de::", CIDR: 48, RoaStatus: "None", Name: "THENETWORKCREW-AU", Description: "The Network Crew Pty Ltd", CountryCode: "NZ", Parent: ASNPrefixesParent{Prefix: "2a06:1280::/29", IP: "2a06:1280::", CIDR: 29, RIRName: "RIPE"}},
				{Prefix: "2a06:1280:ce01::/48", IP: "2a06:1280:ce01::", CIDR: 48, RoaStatus: "None", Name: "NICE-CO-NZ-v6", Description: "Nice.co.nz Ltd IPv6", CountryCode: "NZ", Parent: ASNPrefixesParent{Prefix: "2a06:1280::/29", IP: "2a06:1280::", CIDR: 29, RIRName: "RIPE"}},
				{Prefix: "2a06:1280:ce02::/48", IP: "2a06:1280:ce02::", CIDR: 48, RoaStatus: "None", Name: "ZAPPIE-HOST-NZ-v6", Description: "Zappie Host - Auckland, New Zealand v6", CountryCode: "NZ", Parent: ASNPrefixesParent{Prefix: "2a06:1280::/29", IP: "2a06:1280::", CIDR: 29, RIRName: "RIPE"}},
				{Prefix: "2a06:1280:ce04::/48", IP: "2a06:1280:ce04::", CIDR: 48, RoaStatus: "None", Name: "c9bdd059-ad72-4382-8d8d-da98a40d563e", Description: "Kevin Holly trading as Silent Ghost e.U.", CountryCode: "NL", Parent: ASNPrefixesParent{Prefix: "2a06:1280::/29", IP: "2a06:1280::", CIDR: 29, RIRName: "RIPE"}},
				{Prefix: "2a06:9f80::/29", IP: "2a06:9f80::", CIDR: 29, RoaStatus: "None", Name: "US-ZAPPIE-20151015", Description: "Zappie Host LLC", CountryCode: "VA", Parent: ASNPrefixesParent{Prefix: "2a06:9f80::/29", IP: "2a06:9f80::", CIDR: 29, RIRName: "RIPE"}},
				{Prefix: "2a06:a006::/32", IP: "2a06:a006::", CIDR: 32, RoaStatus: "None", Name: "ZAPPIE-HOST-CL-1", Description: "Zappie Host - Valdivia, Chile", CountryCode: "CL", Parent: ASNPrefixesParent{Prefix: "2a06:a000::/29", IP: "2a06:a000::", CIDR: 29, RIRName: "RIPE"}},
				{Prefix: "2a06:e881:6203::/48", IP: "2a06:e881:6203::", CIDR: 48, RoaStatus: "None", Name: "FR-RANXPLORER-20190213", Description: "Ranxplorer", CountryCode: "FR", Parent: ASNPrefixesParent{Prefix: "2a06:e880::/29", IP: "2a06:e880::", CIDR: 29, RIRName: "RIPE"}},
				{Prefix: "2a07:54c0::/29", IP: "2a07:54c0::", CIDR: 29, RoaStatus: "None", Name: "US-ZAPPIE-20160412", Description: "Zappie Host LLC", CountryCode: "IM", Parent: ASNPrefixesParent{Prefix: "2a07:54c0::/29", IP: "2a07:54c0::", CIDR: 29, RIRName: "RIPE"}},
				{Prefix: "2a09:54c0::/29", IP: "2a09:54c0::", CIDR: 29, RoaStatus: "None", Name: "NZ-KIWIANAHOSTING-20190208", Description: "Kiwiana Hosting Limited", CountryCode: "CY", Parent: ASNPrefixesParent{Prefix: "2a09:54c0::/29", IP: "2a09:54c0::", CIDR: 29, RIRName: "RIPE"}},
				{Prefix: "2a0a:6040::/29", IP: "2a0a:6040::", CIDR: 29, RoaStatus: "None", Name: "US-ZAP-20170323", Description: "Zappie Host LLC", CountryCode: "US", Parent: ASNPrefixesParent{Prefix: "", IP: "", CIDR: 0, RIRName: ""}},
				{Prefix: "2a0b:9e40::/29", IP: "2a0b:9e40::", CIDR: 29, RoaStatus: "None", Name: "NZ-MONTEHOSTING-20170726", Description: "MonteHosting LTD", CountryCode: "ME", Parent: ASNPrefixesParent{Prefix: "2a0b:9e40::/29", IP: "2a0b:9e40::", CIDR: 29, RIRName: "RIPE"}},
				{Prefix: "2a0c:9a40:8082::/48", IP: "2a0c:9a40:8082::", CIDR: 48, RoaStatus: "None", Name: "", Description: "", CountryCode: "", Parent: ASNPrefixesParent{Prefix: "", IP: "", CIDR: 0, RIRName: ""}},
				{Prefix: "2a0c:b642:1a0e::/48", IP: "2a0c:b642:1a0e::", CIDR: 48, RoaStatus: "None", Name: "RANEXPLORER_220219", Description: "Ranxplorer", CountryCode: "FR", Parent: ASNPrefixesParent{Prefix: "2a0c:b640::/29", IP: "2a0c:b640::", CIDR: 29, RIRName: "RIPE"}},
				{Prefix: "2a0c:e640:24::/48", IP: "2a0c:e640:24::", CIDR: 48, RoaStatus: "None", Name: "FIXMIX-NET6-AKL", Description: "fixmix GEN - Auckland", CountryCode: "NZ", Parent: ASNPrefixesParent{Prefix: "2a0c:e640::/29", IP: "2a0c:e640::", CIDR: 29, RIRName: "RIPE"}},
				{Prefix: "2a0d:d900::/29", IP: "2a0d:d900::", CIDR: 29, RoaStatus: "None", Name: "NZ-HOSTINGLOVE-20171213", Description: "HostingLove LImited", CountryCode: "FR", Parent: ASNPrefixesParent{Prefix: "2a0d:d900::/29", IP: "2a0d:d900::", CIDR: 29, RIRName: "RIPE"}},
				{Prefix: "2a0e:fd45:40fd::/48", IP: "2a0e:fd45:40fd::", CIDR: 48, RoaStatus: "None", Name: "", Description: "", CountryCode: "", Parent: ASNPrefixesParent{Prefix: "", IP: "", CIDR: 0, RIRName: ""}},
				{Prefix: "2a0f:5707:ab2f::/48", IP: "2a0f:5707:ab2f::", CIDR: 48, RoaStatus: "None", Name: "", Description: "", CountryCode: "", Parent: ASNPrefixesParent{Prefix: "", IP: "", CIDR: 0, RIRName: ""}},
				{Prefix: "2c0f:f530::/44", IP: "2c0f:f530::", CIDR: 44, RoaStatus: "None", Name: "ZAPPIE-HOST-ZA-V6", Description: "Zappie Host - Johannesburg, South Africa v6", CountryCode: "ZA", Parent: ASNPrefixesParent{Prefix: "2c0f:f530::/32", IP: "2c0f:f530::", CIDR: 32, RIRName: "AfriNIC"}},
				{Prefix: "2c0f:f530:20::/44", IP: "2c0f:f530:20::", CIDR: 44, RoaStatus: "None", Name: "ZAPPIE-HOST-ZA-V6", Description: "Zappie Host - Johannesburg, South Africa v6", CountryCode: "ZA", Parent: ASNPrefixesParent{Prefix: "2c0f:f530::/32", IP: "2c0f:f530::", CIDR: 32, RIRName: "AfriNIC"}},
			},
		},
		Meta: Meta{TimeZone: "UTC", APIVersion: 1, ExecutionTime: "50.85 ms"},
	}

	assert.Equal(t, expected, details)
}

func TestClient_GetASNPeers(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/asn/61138/peers", testHandler("asn-peers.json"))

	details, err := client.GetASNPeers(context.Background(), 61138)
	require.NoError(t, err)

	expected := &ASNPeersInfo{
		Status:        "ok",
		StatusMessage: "Query was successful",
		Data: ASNPeersData{
			IPv4Peers: []ASNIPPeersData{
				{ASN: 137409, Name: "GSLNETWORKS-AS-AP", Description: "GSL Networks Pty LTD", CountryCode: "AU"},
				{ASN: 37153, Name: "xneelo", Description: "xneelo (Pty) Ltd", CountryCode: "ZA"},
				{ASN: 9179, Name: "FIDO-ANYCAST", Description: "FidoNet Registration Services Ltd", CountryCode: "GB"},
				{ASN: 270013, Name: "J AND J SPA (INFOFRACTAL)", Description: "J AND J SPA (INFOFRACTAL)", CountryCode: "CL"},
			},
			IPv6Peers: []ASNIPPeersData{
				{ASN: 210872, Name: "LENETWORK-AS", Description: "Cui, Jiacheng", CountryCode: "DE"},
				{ASN: 211876, Name: "FIXMIX-GEN", Description: "fixmix Technologies", CountryCode: "GB"},
				{ASN: 137409, Name: "GSLNETWORKS-AS-AP", Description: "GSL Networks Pty LTD", CountryCode: "AU"},
				{ASN: 147297, Name: "FATBOY-NETWORKS", Description: "Fat Boy Networks LLC", CountryCode: "US"},
				{ASN: 147028, Name: "WANGYONGJIN-AS-JP", Description: "Wang Yongjin", CountryCode: "HK"},
				{ASN: 34927, Name: "iFog-GmbH", Description: "iFog GmbH", CountryCode: "CH"},
				{ASN: 210481, Name: "EDGENET", Description: "Personal ASN", CountryCode: "RU"},
				{ASN: 9179, Name: "FIDO-ANYCAST", Description: "FidoNet Registration Services Ltd", CountryCode: "GB"},
				{ASN: 14570, Name: "IABAL-PUBLIC-01", Description: "Ansible Networks, LLC", CountryCode: "US"},
				{ASN: 8298, Name: "IPNG", Description: "Pim van Pelt", CountryCode: "CH"},
				{ASN: 209870, Name: "OPTIX-AS", Description: "Optix Transit Ltd", CountryCode: "GB"},
				{ASN: 270013, Name: "J AND J SPA (INFOFRACTAL)", Description: "J AND J SPA (INFOFRACTAL)", CountryCode: "CL"},
				{ASN: 141694, Name: "OTAKUJAPAN-AS", Description: "Otaku Limited", CountryCode: "JP"},
				{ASN: 45177, Name: "DEVOLI-AS-AP", Description: "Devoli", CountryCode: "NZ"},
				{ASN: 212085, Name: "BRUEGGUS-AS", Description: "Alexander Bruegmann", CountryCode: "DE"},
				{ASN: 36369, Name: "LIMEWAVE", Description: "Limewave Communications", CountryCode: "CA"},
				{ASN: 35661, Name: "VIRTUA-SYSTEMS", Description: "VIRTUA SYSTEMS SAS", CountryCode: "FR"},
				{ASN: 37153, Name: "xneelo", Description: "xneelo (Pty) Ltd", CountryCode: "ZA"},
				{ASN: 140731, Name: "TOHU-OP-AP", Description: "TOHU Public Internet", CountryCode: "CN"},
				{ASN: 211013, Name: "JMJITSOLUTIONS", Description: "Jaroslaw Labiszewski trading as JMJ IT SOLUTIONS", CountryCode: "PL"},
			},
		},
		Meta: Meta{TimeZone: "UTC", APIVersion: 1, ExecutionTime: "43.38 ms"},
	}

	assert.Equal(t, expected, details)
}

func TestClient_GetASNUpstreams(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/asn/61138/upstreams", testHandler("asn-upstreams.json"))

	details, err := client.GetASNUpstreams(context.Background(), 61138)
	require.NoError(t, err)

	expected := &ASNUpstreamsInfo{
		Status:        "ok",
		StatusMessage: "Query was successful",
		Data: ASNUpstreamsData{
			IPv4Upstreams: []ASNIPUpstreamsData{
				{ASN: 137409, Name: "GSLNETWORKS-AS-AP", Description: "GSL Networks Pty LTD", CountryCode: "AU"},
				{ASN: 37153, Name: "xneelo", Description: "xneelo (Pty) Ltd", CountryCode: "ZA"},
				{ASN: 270013, Name: "J AND J SPA (INFOFRACTAL)", Description: "J AND J SPA (INFOFRACTAL)", CountryCode: "CL"},
			},
			IPv6Upstreams: []ASNIPUpstreamsData{
				{ASN: 137409, Name: "GSLNETWORKS-AS-AP", Description: "GSL Networks Pty LTD", CountryCode: "AU"},
				{ASN: 37153, Name: "xneelo", Description: "xneelo (Pty) Ltd", CountryCode: "ZA"},
				{ASN: 270013, Name: "J AND J SPA (INFOFRACTAL)", Description: "J AND J SPA (INFOFRACTAL)", CountryCode: "CL"},
				{ASN: 35661, Name: "VIRTUA-SYSTEMS", Description: "VIRTUA SYSTEMS SAS", CountryCode: "FR"},
				{ASN: 36369, Name: "LIMEWAVE", Description: "Limewave Communications", CountryCode: "CA"},
			},
		},
		Meta: Meta{TimeZone: "UTC", APIVersion: 1, ExecutionTime: "24.1 ms"},
	}

	assert.Equal(t, expected, details)
}

func TestClient_GetASNDownstreams(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/asn/61138/downstreams", testHandler("asn-downstreams.json"))

	details, err := client.GetASNDownstreams(context.Background(), 61138)
	require.NoError(t, err)

	expected := &ASNDownstreamsInfo{
		Status:        "ok",
		StatusMessage: "Query was successful",
		Data: ASNDownstreamsData{
			IPv4Downstreams: []ASNIPDownstreamsData{},
			IPv6Downstreams: []ASNIPDownstreamsData{
				{ASN: 211876, Name: "FIXMIX-GEN", Description: "fixmix Technologies", CountryCode: "GB"},
				{ASN: 147028, Name: "WANGYONGJIN-AS-JP", Description: "Wang Yongjin", CountryCode: "HK"},
				{ASN: 147297, Name: "FATBOY-NETWORKS", Description: "Fat Boy Networks LLC", CountryCode: "US"},
				{ASN: 209870, Name: "OPTIX-AS", Description: "Optix Transit Ltd", CountryCode: "GB"},
				{ASN: 212085, Name: "BRUEGGUS-AS", Description: "Alexander Bruegmann", CountryCode: "DE"},
				{ASN: 210481, Name: "EDGENET", Description: "Personal ASN", CountryCode: "RU"},
				{ASN: 14570, Name: "IABAL-PUBLIC-01", Description: "Ansible Networks, LLC", CountryCode: "US"},
			},
		},
		Meta: Meta{TimeZone: "UTC", APIVersion: 1, ExecutionTime: "20.23 ms"},
	}

	assert.Equal(t, expected, details)
}

func TestClient_GetASNIxs(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/asn/61138/ixs", testHandler("asn-ixs.json"))

	details, err := client.GetASNIxs(context.Background(), 61138)
	require.NoError(t, err)

	expected := &ASNIxsInfo{
		Status:        "ok",
		StatusMessage: "Query was successful",
		Data: []ASNIxsData{
			{IxID: 585, Name: "EVIX", NameFull: "Experimental Virtual Internet Exchange", CountryCode: "US", IPv4Address: "206.81.104.165", IPv6Address: "2602:fed2:fff:ffff::165", Speed: 100},
			{IxID: 599, Name: "LL-IX", NameFull: "LL-IX", CountryCode: "RO", IPv4Address: "5.101.92.255", IPv6Address: "2001:678:4fc::92:255", Speed: 100},
			{IxID: 780, Name: "TOHU IX", NameFull: "TOHU IX", CountryCode: "CN", IPv4Address: "", IPv6Address: "2406:840:eb8f:1:0:6:1138:1", Speed: 100},
			{IxID: 829, Name: "PyramIX", NameFull: "Pyramids Internet Exchange", CountryCode: "EG", IPv4Address: "104.167.214.111", IPv6Address: "2a0e:46c4:102::611:38:1", Speed: 100},
			{IxID: 857, Name: "HamroIX-Amsterdam", NameFull: "Hamro Internet eXchange", CountryCode: "NL", IPv4Address: "104.167.214.131", IPv6Address: "2a0e:b107:f21::113", Speed: 100},
		},
		Meta: Meta{TimeZone: "UTC", APIVersion: 1, ExecutionTime: "22.68 ms"},
	}

	assert.Equal(t, expected, details)
}

func TestClient_GetPrefix(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/prefix/192.209.63.0/24", testHandler("prefix.json"))

	details, err := client.GetPrefix(context.Background(), "192.209.63.0", 24)
	require.NoError(t, err)

	expected := &PrefixInfo{
		Status:        "ok",
		StatusMessage: "Query was successful",
		Data: PrefixData{
			Prefix: "192.209.63.0/24",
			IP:     "192.209.63.0",
			CIDR:   24,
			ASNs: []ASN{{
				ASN:         1239,
				Name:        "SPRINTLINK",
				Description: "Sprint",
				CountryCode: "US",
			}},
			Name:             "BITACCEL-NETWORK",
			DescriptionShort: "BitAccel",
			DescriptionFull:  []string{"BitAccel"},
			EmailContacts:    []string{"abuse@bitaccel.com"},
			AbuseContacts:    []string{"abuse@bitaccel.com"},
			OwnerAddress:     []string{"135 Red Head Ln.", "Gilmer", "TX", "75645", "US"},
			CountryCodes:     CountryCodeData{WhoisCountryCode: "US", RIRAllocationCountryCode: "US", MaxmindCountryCode: ""},
			RIRAllocation:    AllocationData{RIRName: "ARIN", CountryCode: "US", IP: "192.209.62.0", CIDR: 23, Prefix: "192.209.62.0/23", DateAllocated: "2015-04-28 00:00:00"},
			MaxMind:          MaxMindData{},
			DateUpdated:      "2020-12-06 03:30:13",
		},
		Meta: Meta{TimeZone: "UTC", APIVersion: 1, ExecutionTime: "442.72 ms"},
	}

	assert.Equal(t, expected, details)
}

func TestClient_GetIP(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/ip/2a05:dfc7:60::", testHandler("ip.json"))

	details, err := client.GetIP(context.Background(), "2a05:dfc7:60::")
	require.NoError(t, err)

	expected := &IPInfo{
		Status:        "ok",
		StatusMessage: "Query was successful",
		Data: IPData{
			Prefixes: []PrefixData{
				{
					Prefix: "2a05:dfc0::/29",
					IP:     "2a05:dfc0::",
					CIDR:   29,
					Name:   "US-ZAPPIE-20150303",
				},
			},
			RIRAllocation: IPAllocationData{RIRName: "RIPE", CountryCode: "US", IP: "2a05:dfc0::", CIDR: "29", Prefix: "2a05:dfc0::/29", DateAllocated: "2015-03-03 00:00:00"},
		},
		Meta: Meta{TimeZone: "UTC", APIVersion: 1, ExecutionTime: "57.3 ms"},
	}
	assert.Equal(t, expected, details)
}

func TestClient_GetIX(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/ix/492", testHandler("ix.json"))

	details, err := client.GetIX(context.Background(), 492)
	require.NoError(t, err)

	expected := &IXInfo{
		Status:        "ok",
		StatusMessage: "Query was successful",
		Data: IXData{
			Name:         "MIXP.me",
			NameFull:     "Montenegro Internet eXchange Point",
			Website:      "http://www.mixp.me/eng/",
			TechEmail:    "mixp@ac.me",
			TechPhone:    "+38220414282",
			PolicyEmail:  "mixp@ac.me",
			City:         "Podgorica",
			CountryCode:  "ME",
			URLStats:     json.RawMessage{0x6e, 0x75, 0x6c, 0x6c},
			MembersCount: 2,
			Members: []MemberData{
				{ASN: 200608, Name: "MIXP", Description: "University of Montenegro", CountryCode: "ME", IPv4Address: "185.1.44.1", IPv6Address: "2001:7f8:22::1", Speed: 1000},
				{ASN: 210762, Name: "FROOT_TGD1", Description: "Internet Systems Consortium Inc.", CountryCode: "US", IPv4Address: "185.1.44.90", IPv6Address: "2001:7f8:22::a", Speed: 10000},
			},
		},
		Meta: Meta{TimeZone: "UTC", APIVersion: 1, ExecutionTime: "16.73 ms"},
	}

	assert.Equal(t, expected, details)
}

func TestClient_GetSearch(t *testing.T) {
	client, mux := setupTest(t)

	mux.HandleFunc("/search", testHandler("search.json"))

	details, err := client.GetSearch(context.Background(), "digitalocean")
	require.NoError(t, err)

	expected := &SearchInfo{
		Status:        "ok",
		StatusMessage: "Query was successful",
		Data: SearchData{
			ASNs: []SearchASNData{
				{ASN: 133165, Name: "DIGITALOCEAN-AS-AP", Description: "Digital Ocean, Inc.", CountryCode: "SG", EmailContacts: []string{"abuse@digitalocean.com"}, AbuseContacts: []string{"abuse@digitalocean.com"}, RIRName: "APNIC"},
				{ASN: 135340, Name: "DIGITALOCEAN-AS-IN", Description: "Digital Ocean, Inc.", CountryCode: "US", EmailContacts: []string{"abuse@digitalocean.com"}, AbuseContacts: []string{"abuse@digitalocean.com"}, RIRName: "APNIC"},
				{ASN: 14061, Name: "DIGITALOCEAN-ASN", Description: "DigitalOcean, LLC", CountryCode: "US", EmailContacts: []string{"abuse@digitalocean.com", "noc@digitalocean.com"}, AbuseContacts: []string{"abuse@digitalocean.com"}, RIRName: "ARIN"},
				{ASN: 393406, Name: "DIGITALOCEAN-AS393406", Description: "DigitalOcean, LLC", CountryCode: "US", EmailContacts: []string{"noc@digitalocean.com", "abuse@digitalocean.com"}, AbuseContacts: []string{"abuse@digitalocean.com"}, RIRName: "ARIN"},
				{ASN: 394362, Name: "DIGITALOCEAN-AS394362", Description: "DigitalOcean, LLC", CountryCode: "US", EmailContacts: []string{"noc@digitalocean.com", "abuse@digitalocean.com"}, AbuseContacts: []string{"abuse@digitalocean.com"}, RIRName: "ARIN"},
				{ASN: 39690, Name: "DigitalOcean-NOC", Description: "Digital Ocean SRL", CountryCode: "MD", EmailContacts: []string{"abuse@digitalocean.eu.com", "info@digitalocean.eu.com"}, AbuseContacts: []string{"abuse@digitalocean.eu.com"}, RIRName: "RIPE"},
				{ASN: 62567, Name: "DIGITALOCEAN-AS62567", Description: "DigitalOcean, LLC", CountryCode: "US", EmailContacts: []string{"noc@digitalocean.com", "abuse@digitalocean.com"}, AbuseContacts: []string{"abuse@digitalocean.com"}, RIRName: "ARIN"},
			},
			IPv4Prefixes: []SearchIPPrefixesData{
				{Prefix: "103.253.144.0/22", IP: "103.253.144.0", CIDR: 22, Name: "DIGITALOCEAN-AP", CountryCode: "SG", Description: "Digital Ocean, Inc.", EmailContacts: []string{"abuse@digitalocean.com"}, AbuseContacts: []string{"abuse@digitalocean.com"}, RIRName: "APNIC", ParentPrefix: "103.253.144.0/22", ParentIP: "103.253.144.0", ParentCIDR: 22},
				{Prefix: "103.253.145.0/24", IP: "103.253.145.0", CIDR: 24, Name: "DIGITALOCEAN-AP", CountryCode: "SG", Description: "Digital Ocean, Inc.", EmailContacts: []string{"abuse@digitalocean.com"}, AbuseContacts: []string{"abuse@digitalocean.com"}, RIRName: "APNIC", ParentPrefix: "103.253.144.0/22", ParentIP: "103.253.144.0", ParentCIDR: 22},
				{Prefix: "103.253.147.0/24", IP: "103.253.147.0", CIDR: 24, Name: "DIGITALOCEAN-AP", CountryCode: "SG", Description: "Digital Ocean, Inc.", EmailContacts: []string{"abuse@digitalocean.com"}, AbuseContacts: []string{"abuse@digitalocean.com"}, RIRName: "APNIC", ParentPrefix: "103.253.144.0/22", ParentIP: "103.253.144.0", ParentCIDR: 22},
				{Prefix: "104.131.0.0/16", IP: "104.131.0.0", CIDR: 16, Name: "DIGITALOCEAN-104-131-0-0", CountryCode: "US", Description: "DigitalOcean, LLC", EmailContacts: []string{"noc@digitalocean.com", "abuse@digitalocean.com"}, AbuseContacts: []string{"abuse@digitalocean.com"}, RIRName: "ARIN", ParentPrefix: "104.131.0.0/16", ParentIP: "104.131.0.0", ParentCIDR: 16},
				{Prefix: "104.131.0.0/18", IP: "104.131.0.0", CIDR: 18, Name: "DIGITALOCEAN-104-131-0-0", CountryCode: "US", Description: "DigitalOcean, LLC", EmailContacts: []string{"noc@digitalocean.com", "abuse@digitalocean.com"}, AbuseContacts: []string{"abuse@digitalocean.com"}, RIRName: "ARIN", ParentPrefix: "104.131.0.0/16", ParentIP: "104.131.0.0", ParentCIDR: 16},
				{Prefix: "104.236.0.0/16", IP: "104.236.0.0", CIDR: 16, Name: "DIGITALOCEAN-104-236-0-0", CountryCode: "US", Description: "DigitalOcean, LLC", EmailContacts: []string{"noc@digitalocean.com", "abuse@digitalocean.com"}, AbuseContacts: []string{"abuse@digitalocean.com"}, RIRName: "ARIN", ParentPrefix: "104.236.0.0/16", ParentIP: "104.236.0.0", ParentCIDR: 16},
				{Prefix: "138.197.112.0/20", IP: "138.197.112.0", CIDR: 20, Name: "DIGITALOCEAN-138-197-0-0", CountryCode: "US", Description: "DigitalOcean, LLC", EmailContacts: []string{"noc@digitalocean.com", "abuse@digitalocean.com"}, AbuseContacts: []string{"abuse@digitalocean.com"}, RIRName: "ARIN", ParentPrefix: "138.197.0.0/16", ParentIP: "138.197.0.0", ParentCIDR: 16},
			},
			IPv6Prefixes: []SearchIPPrefixesData{
				{Prefix: "2400:6180:100::/40", IP: "2400:6180:100::", CIDR: 40, Name: "DIGITALOCEAN-AP-20131119", CountryCode: "SG", Description: "Digital Ocean, Inc.", EmailContacts: []string{"abuse@digitalocean.com"}, AbuseContacts: []string{"abuse@digitalocean.com"}, RIRName: "APNIC", ParentPrefix: "2400:6180::/32", ParentIP: "2400:6180::", ParentCIDR: 32},
				{Prefix: "2400:6180::/48", IP: "2400:6180::", CIDR: 48, Name: "DIGITALOCEAN-AP-20131119", CountryCode: "SG", Description: "Digital Ocean, Inc.", EmailContacts: []string{"abuse@digitalocean.com"}, AbuseContacts: []string{"abuse@digitalocean.com"}, RIRName: "APNIC", ParentPrefix: "2400:6180::/32", ParentIP: "2400:6180::", ParentCIDR: 32},
				{Prefix: "2400:6180:f000::/36", IP: "2400:6180:f000::", CIDR: 36, Name: "DIGITALOCEAN-AP", CountryCode: "US", Description: "DigitalOcean, LLC", EmailContacts: []string{"abuse@digitalocean.com"}, AbuseContacts: []string{"abuse@digitalocean.com"}, RIRName: "APNIC", ParentPrefix: "2400:6180::/32", ParentIP: "2400:6180::", ParentCIDR: 32},
				{Prefix: "2604:a880:1::/48", IP: "2604:a880:1::", CIDR: 48, Name: "DIGITALOCEAN", CountryCode: "US", Description: "DigitalOcean, LLC", EmailContacts: []string{"noc@digitalocean.com", "abuse@digitalocean.com"}, AbuseContacts: []string{"abuse@digitalocean.com"}, RIRName: "ARIN", ParentPrefix: "2604:a880::/32", ParentIP: "2604:a880::", ParentCIDR: 32},
				{Prefix: "2604:a880:2::/48", IP: "2604:a880:2::", CIDR: 48, Name: "DIGITALOCEAN", CountryCode: "US", Description: "DigitalOcean, LLC", EmailContacts: []string{"noc@digitalocean.com", "abuse@digitalocean.com"}, AbuseContacts: []string{"abuse@digitalocean.com"}, RIRName: "ARIN", ParentPrefix: "2604:a880::/32", ParentIP: "2604:a880::", ParentCIDR: 32},
				{Prefix: "2a03:b0c0:3::/48", IP: "2a03:b0c0:3::", CIDR: 48, Name: "DIGITALOCEAN", CountryCode: "DE", Description: "DIGITALOCEAN", EmailContacts: []string{"abuse@digitalocean.com", "noc@digitalocean.com"}, AbuseContacts: []string{"abuse@digitalocean.com"}, RIRName: "RIPE", ParentPrefix: "2a03:b0c0::/32", ParentIP: "2a03:b0c0::", ParentCIDR: 32},
				{Prefix: "2a03:b0c0::/48", IP: "2a03:b0c0::", CIDR: 48, Name: "DIGITALOCEAN", CountryCode: "NL", Description: "DIGITALOCEAN", EmailContacts: []string{"abuse@digitalocean.com", "noc@digitalocean.com"}, AbuseContacts: []string{"abuse@digitalocean.com"}, RIRName: "RIPE", ParentPrefix: "2a03:b0c0::/32", ParentIP: "2a03:b0c0::", ParentCIDR: 32},
				{Prefix: "2a03:b0c0::/32", IP: "2a03:b0c0::", CIDR: 32, Name: "US-DIGITALOCEANLLC-20121228", CountryCode: "NL", Description: "DigitalOcean, LLC", EmailContacts: []string{"abuse@digitalocean.com", "noc@digitalocean.com"}, AbuseContacts: []string{"abuse@digitalocean.com"}, RIRName: "RIPE", ParentPrefix: "2a03:b0c0::/32", ParentIP: "2a03:b0c0::", ParentCIDR: 32},
			},
		},
		Meta: Meta{TimeZone: "UTC", APIVersion: 1, ExecutionTime: "95.27 ms"},
	}

	assert.Equal(t, expected, details)
}
