package bgpview

import "encoding/json"

type Meta struct {
	TimeZone      string `json:"time_zone,omitempty"`
	APIVersion    int    `json:"api_version,omitempty"`
	ExecutionTime string `json:"execution_time,omitempty"`
}

type ASNInfo struct {
	Status        string  `json:"status,omitempty"`
	StatusMessage string  `json:"status_message,omitempty"`
	Data          ASNData `json:"data,omitempty"`
	Meta          Meta    `json:"@meta,omitempty"`
}

type ASNData struct {
	ASN               int            `json:"asn,omitempty"`
	Name              string         `json:"name,omitempty"`
	DescriptionShort  string         `json:"description_short,omitempty"`
	DescriptionFull   []string       `json:"description_full,omitempty"`
	CountryCode       string         `json:"country_code,omitempty"`
	Website           string         `json:"website,omitempty"`
	EmailContacts     []string       `json:"email_contacts,omitempty"`
	AbuseContacts     []string       `json:"abuse_contacts,omitempty"`
	LookingGlass      string         `json:"looking_glass,omitempty"`
	TrafficEstimation string         `json:"traffic_estimation,omitempty"`
	TrafficRatio      string         `json:"traffic_ratio,omitempty"`
	OwnerAddress      []string       `json:"owner_address,omitempty"`
	RIRAllocation     AllocationData `json:"rir_allocation,omitempty"`
	DateUpdated       string         `json:"date_updated,omitempty"`
}

type ASNPrefixesInfo struct {
	Status        string          `json:"status,omitempty"`
	StatusMessage string          `json:"status_message,omitempty"`
	Data          ASNPrefixesData `json:"data,omitempty"`
	Meta          Meta            `json:"@meta,omitempty"`
}

type ASNPrefixesData struct {
	IPv4Prefixes []ASNIPPrefixesData `json:"ipv4_prefixes,omitempty"`
	IPv6Prefixes []ASNIPPrefixesData `json:"ipv6_prefixes,omitempty"`
}

type ASNIPPrefixesData struct {
	Prefix      string            `json:"prefix,omitempty"`
	IP          string            `json:"ip,omitempty"`
	CIDR        int               `json:"cidr,omitempty"`
	RoaStatus   string            `json:"roa_status,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	CountryCode string            `json:"country_code,omitempty"`
	Parent      ASNPrefixesParent `json:"parent,omitempty"`
}

type ASNPrefixesParent struct {
	Prefix  string `json:"prefix,omitempty"`
	IP      string `json:"ip,omitempty"`
	CIDR    int    `json:"cidr,omitempty"`
	RIRName string `json:"rir_name,omitempty"`
}

type ASNPeersInfo struct {
	Status        string       `json:"status,omitempty"`
	StatusMessage string       `json:"status_message,omitempty"`
	Data          ASNPeersData `json:"data,omitempty"`
	Meta          Meta         `json:"@meta,omitempty"`
}

type ASNPeersData struct {
	IPv4Peers []ASNIPPeersData `json:"ipv4_peers,omitempty"`
	IPv6Peers []ASNIPPeersData `json:"ipv6_peers,omitempty"`
}

type ASNIPPeersData struct {
	ASN         int    `json:"asn,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

type ASNUpstreamsInfo struct {
	Status        string           `json:"status,omitempty"`
	StatusMessage string           `json:"status_message,omitempty"`
	Data          ASNUpstreamsData `json:"data,omitempty"`
	Meta          Meta             `json:"@meta,omitempty"`
}

type ASNUpstreamsData struct {
	IPv4Upstreams []ASNIPUpstreamsData `json:"ipv4_upstreams,omitempty"`
	IPv6Upstreams []ASNIPUpstreamsData `json:"ipv6_upstreams,omitempty"`
}

type ASNIPUpstreamsData struct {
	ASN         int      `json:"asn,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	CountryCode string   `json:"country_code,omitempty"`
	BgpPaths    []string `json:"bgp_paths,omitempty"`
}

type ASNDownstreamsInfo struct {
	Status        string             `json:"status,omitempty"`
	StatusMessage string             `json:"status_message,omitempty"`
	Data          ASNDownstreamsData `json:"data,omitempty"`
	Meta          Meta               `json:"@meta,omitempty"`
}

type ASNIPDownstreamsData struct {
	ASN         int      `json:"asn,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	CountryCode string   `json:"country_code,omitempty"`
	BgpPaths    []string `json:"bgp_paths,omitempty"`
}

type ASNDownstreamsData struct {
	IPv4Downstreams []ASNIPDownstreamsData `json:"ipv4_downstreams,omitempty"`
	IPv6Downstreams []ASNIPDownstreamsData `json:"ipv6_downstreams,omitempty"`
}

type ASNIxsInfo struct {
	Status        string       `json:"status,omitempty"`
	StatusMessage string       `json:"status_message,omitempty"`
	Data          []ASNIxsData `json:"data,omitempty"`
	Meta          Meta         `json:"@meta,omitempty"`
}

type ASNIxsData struct {
	IxID        int    `json:"ix_id,omitempty"`
	Name        string `json:"name,omitempty"`
	NameFull    string `json:"name_full,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	IPv4Address string `json:"ipv4_address,omitempty"`
	IPv6Address string `json:"ipv6_address,omitempty"`
	Speed       int    `json:"speed,omitempty"`
}

type PrefixInfo struct {
	Status        string     `json:"status,omitempty"`
	StatusMessage string     `json:"status_message,omitempty"`
	Data          PrefixData `json:"data,omitempty"`
	Meta          Meta       `json:"@meta,omitempty"`
}

type PrefixData struct {
	Prefix           string          `json:"prefix,omitempty"`
	IP               string          `json:"ip,omitempty"`
	CIDR             int             `json:"cidr,omitempty"`
	ASNs             []ASN           `json:"asns,omitempty"`
	Name             string          `json:"name,omitempty"`
	DescriptionShort string          `json:"description_short,omitempty"`
	DescriptionFull  []string        `json:"description_full,omitempty"`
	EmailContacts    []string        `json:"email_contacts,omitempty"`
	AbuseContacts    []string        `json:"abuse_contacts,omitempty"`
	OwnerAddress     []string        `json:"owner_address,omitempty"`
	CountryCodes     CountryCodeData `json:"country_codes,omitempty"`
	RIRAllocation    AllocationData  `json:"rir_allocation,omitempty"`
	MaxMind          MaxMindData     `json:"maxmind,omitempty"`
	DateUpdated      string          `json:"date_updated,omitempty"`
}

type MaxMindData struct {
	CountryCode string      `json:"country_code,omitempty"`
	City        interface{} `json:"city,omitempty"`
}

type AllocationData struct {
	RIRName       string `json:"rir_name,omitempty"`
	CountryCode   string `json:"country_code,omitempty"`
	IP            string `json:"ip,omitempty"`
	CIDR          int    `json:"cidr,omitempty"`
	Prefix        string `json:"prefix,omitempty"`
	DateAllocated string `json:"date_allocated,omitempty"`
}

type CountryCodeData struct {
	WhoisCountryCode         string `json:"whois_country_code,omitempty"`
	RIRAllocationCountryCode string `json:"rir_allocation_country_code,omitempty"`
	MaxmindCountryCode       string `json:"maxmind_country_code,omitempty"`
}

type ASN struct {
	ASN         int    `json:"asn,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

type IPInfo struct {
	Status        string `json:"status,omitempty"`
	StatusMessage string `json:"status_message,omitempty"`
	Data          IPData `json:"data,omitempty"`
	Meta          Meta   `json:"@meta,omitempty"`
}

type IPData struct {
	Prefixes        []PrefixData     `json:"prefixes,omitempty"`
	RIRAllocation   IPAllocationData `json:"rir_allocation,omitempty"`
	MaxMind         MaxMindData      `json:"maxmind,omitempty"`
	RelatedPrefixes []PrefixData     `json:"related_prefixes,omitempty"`
}

type IPAllocationData struct {
	RIRName       string `json:"rir_name,omitempty"`
	CountryCode   string `json:"country_code,omitempty"`
	IP            string `json:"ip,omitempty"`
	CIDR          string `json:"cidr,omitempty"`
	Prefix        string `json:"prefix,omitempty"`
	DateAllocated string `json:"date_allocated,omitempty"`
}

type IXInfo struct {
	Status        string `json:"status,omitempty"`
	StatusMessage string `json:"status_message,omitempty"`
	Data          IXData `json:"data,omitempty"`
	Meta          Meta   `json:"@meta,omitempty"`
}

type IXData struct {
	Name         string          `json:"name,omitempty"`
	NameFull     string          `json:"name_full,omitempty"`
	Website      string          `json:"website,omitempty"`
	TechEmail    string          `json:"tech_email,omitempty"`
	TechPhone    string          `json:"tech_phone,omitempty"`
	PolicyEmail  string          `json:"policy_email,omitempty"`
	PolicyPhone  string          `json:"policy_phone,omitempty"`
	City         string          `json:"city,omitempty"`
	CountryCode  string          `json:"country_code,omitempty"`
	URLStats     json.RawMessage `json:"url_stats,omitempty"`
	MembersCount int             `json:"members_count,omitempty"`
	Members      []MemberData    `json:"members,omitempty"`
}

type MemberData struct {
	ASN         int    `json:"asn,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	IPv4Address string `json:"ipv4_address,omitempty"`
	IPv6Address string `json:"ipv6_address,omitempty"`
	Speed       int    `json:"speed,omitempty"`
}

type SearchInfo struct {
	Status        string     `json:"status,omitempty"`
	StatusMessage string     `json:"status_message,omitempty"`
	Data          SearchData `json:"data,omitempty"`
	Meta          Meta       `json:"@meta,omitempty"`
}

type SearchData struct {
	ASNs         []SearchASNData        `json:"asns,omitempty"`
	IPv4Prefixes []SearchIPPrefixesData `json:"ipv4_prefixes,omitempty"`
	IPv6Prefixes []SearchIPPrefixesData `json:"ipv6_prefixes,omitempty"`
}

type SearchASNData struct {
	ASN           int      `json:"asn,omitempty"`
	Name          string   `json:"name,omitempty"`
	Description   string   `json:"description,omitempty"`
	CountryCode   string   `json:"country_code,omitempty"`
	EmailContacts []string `json:"email_contacts,omitempty"`
	AbuseContacts []string `json:"abuse_contacts,omitempty"`
	RIRName       string   `json:"rir_name,omitempty"`
}

type SearchIPPrefixesData struct {
	Prefix        string   `json:"prefix,omitempty"`
	IP            string   `json:"ip,omitempty"`
	CIDR          int      `json:"cidr,omitempty"`
	Name          string   `json:"name,omitempty"`
	CountryCode   string   `json:"country_code,omitempty"`
	Description   string   `json:"description,omitempty"`
	EmailContacts []string `json:"email_contacts,omitempty"`
	AbuseContacts []string `json:"abuse_contacts,omitempty"`
	RIRName       string   `json:"rir_name,omitempty"`
	ParentPrefix  string   `json:"parent_prefix,omitempty"`
	ParentIP      string   `json:"parent_ip,omitempty"`
	ParentCIDR    int      `json:"parent_cidr,omitempty"`
}
