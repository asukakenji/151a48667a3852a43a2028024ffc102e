package lib

import "googlemaps.github.io/maps"

var (
	// Input: []string{"22.372081,114.107877", "22.284419,114.159510", "22.326442,114.167811"}
	resp1 = &maps.DistanceMatrixResponse{
		OriginAddresses: []string{
			"11 Hoi Shing Rd, Tsuen Wan, Hong Kong",
			"Laguna City, Central, Hong Kong",
			"789 Nathan Rd, Mong Kok, Hong Kong",
		},
		DestinationAddresses: []string{
			"11 Hoi Shing Rd, Tsuen Wan, Hong Kong",
			"Laguna City, Central, Hong Kong",
			"789 Nathan Rd, Mong Kok, Hong Kong",
		},
		Rows: []maps.DistanceMatrixElementsRow{
			maps.DistanceMatrixElementsRow{
				Elements: []*maps.DistanceMatrixElement{
					&maps.DistanceMatrixElement{
						Status:            "OK",
						Duration:          0,
						DurationInTraffic: 38000000000,
						Distance: maps.Distance{
							HumanReadable: "1 m",
							Meters:        0,
						},
					},
					&maps.DistanceMatrixElement{
						Status:            "OK",
						Duration:          995000000000,
						DurationInTraffic: 1040000000000,
						Distance: maps.Distance{
							HumanReadable: "15.5 km",
							Meters:        15518,
						},
					},
					&maps.DistanceMatrixElement{
						Status:            "OK",
						Duration:          845000000000,
						DurationInTraffic: 925000000000,
						Distance: maps.Distance{
							HumanReadable: "9.7 km",
							Meters:        9667,
						},
					},
				},
			},
			maps.DistanceMatrixElementsRow{
				Elements: []*maps.DistanceMatrixElement{
					&maps.DistanceMatrixElement{
						Status:            "OK",
						Duration:          878000000000,
						DurationInTraffic: 914000000000,
						Distance: maps.Distance{
							HumanReadable: "15.2 km",
							Meters:        15223,
						},
					},
					&maps.DistanceMatrixElement{
						Status:            "OK",
						Duration:          0,
						DurationInTraffic: 3000000000,
						Distance: maps.Distance{
							HumanReadable: "1 m",
							Meters:        0,
						},
					},
					&maps.DistanceMatrixElement{
						Status:            "OK",
						Duration:          883000000000,
						DurationInTraffic: 932000000000,
						Distance: maps.Distance{
							HumanReadable: "8.3 km",
							Meters:        8333,
						},
					},
				},
			},
			maps.DistanceMatrixElementsRow{
				Elements: []*maps.DistanceMatrixElement{
					&maps.DistanceMatrixElement{
						Status:            "OK",
						Duration:          816000000000,
						DurationInTraffic: 788000000000,
						Distance: maps.Distance{
							HumanReadable: "10.3 km",
							Meters:        10329,
						},
					},
					&maps.DistanceMatrixElement{
						Status:            "OK",
						Duration:          908000000000,
						DurationInTraffic: 938000000000,
						Distance: maps.Distance{
							HumanReadable: "8.5 km",
							Meters:        8464,
						},
					},
					&maps.DistanceMatrixElement{
						Status:            "OK",
						Duration:          0,
						DurationInTraffic: 1000000000,
						Distance: maps.Distance{
							HumanReadable: "1 m",
							Meters:        0,
						},
					},
				},
			},
		},
	}

	// Input: []string{"22.324339,114.169027", "22.147034,113.559835"}
	resp2 = &maps.DistanceMatrixResponse{
		OriginAddresses: []string{
			"142 Prince Edward Rd W, Mong Kok, Hong Kong",
			"22.147034,113.559835",
		},
		DestinationAddresses: []string{
			"142 Prince Edward Rd W, Mong Kok, Hong Kong",
			"22.147034,113.559835",
		},
		Rows: []maps.DistanceMatrixElementsRow{
			maps.DistanceMatrixElementsRow{
				Elements: []*maps.DistanceMatrixElement{
					&maps.DistanceMatrixElement{
						Status:            "OK",
						Duration:          0,
						DurationInTraffic: 67000000000,
						Distance: maps.Distance{
							HumanReadable: "1 m",
							Meters:        0,
						},
					},
					&maps.DistanceMatrixElement{
						Status:            "ZERO_RESULTS",
						Duration:          0,
						DurationInTraffic: 0,
						Distance: maps.Distance{
							HumanReadable: "",
							Meters:        0,
						},
					},
				},
			},
			maps.DistanceMatrixElementsRow{
				Elements: []*maps.DistanceMatrixElement{
					&maps.DistanceMatrixElement{
						Status:            "ZERO_RESULTS",
						Duration:          0,
						DurationInTraffic: 0,
						Distance: maps.Distance{
							HumanReadable: "",
							Meters:        0,
						},
					},
					&maps.DistanceMatrixElement{
						Status:            "OK",
						Duration:          0,
						DurationInTraffic: 0,
						Distance: maps.Distance{
							HumanReadable: "1 m",
							Meters:        0,
						},
					},
				},
			},
		},
	}

	// Input: []string{"90.100000,-60.528112", "80.262833,-60.528112"}
	resp3 = &maps.DistanceMatrixResponse{
		OriginAddresses: []string{
			"",
			"",
		},
		DestinationAddresses: []string{
			"",
			"",
		},
		Rows: []maps.DistanceMatrixElementsRow{
			maps.DistanceMatrixElementsRow{
				Elements: []*maps.DistanceMatrixElement{
					&maps.DistanceMatrixElement{
						Status:            "NOT_FOUND",
						Duration:          0,
						DurationInTraffic: 0,
						Distance: maps.Distance{
							HumanReadable: "",
							Meters:        0,
						},
					},
					&maps.DistanceMatrixElement{
						Status:            "NOT_FOUND",
						Duration:          0,
						DurationInTraffic: 0,
						Distance: maps.Distance{
							HumanReadable: "",
							Meters:        0,
						},
					},
				},
			},
			maps.DistanceMatrixElementsRow{
				Elements: []*maps.DistanceMatrixElement{
					&maps.DistanceMatrixElement{
						Status:            "NOT_FOUND",
						Duration:          0,
						DurationInTraffic: 0,
						Distance: maps.Distance{
							HumanReadable: "",
							Meters:        0,
						},
					},
					&maps.DistanceMatrixElement{
						Status:            "NOT_FOUND",
						Duration:          0,
						DurationInTraffic: 0,
						Distance: maps.Distance{
							HumanReadable: "",
							Meters:        0,
						},
					},
				},
			},
		},
	}

	// Fake
	respX = &maps.DistanceMatrixResponse{
		OriginAddresses: []string{
			"",
			"",
		},
		DestinationAddresses: []string{
			"",
			"",
		},
		Rows: []maps.DistanceMatrixElementsRow{
			maps.DistanceMatrixElementsRow{
				Elements: []*maps.DistanceMatrixElement{
					&maps.DistanceMatrixElement{
						Status:            "XXX",
						Duration:          0,
						DurationInTraffic: 0,
						Distance: maps.Distance{
							HumanReadable: "",
							Meters:        0,
						},
					},
					&maps.DistanceMatrixElement{
						Status:            "NOT_FOUND",
						Duration:          0,
						DurationInTraffic: 0,
						Distance: maps.Distance{
							HumanReadable: "",
							Meters:        0,
						},
					},
				},
			},
			maps.DistanceMatrixElementsRow{
				Elements: []*maps.DistanceMatrixElement{
					&maps.DistanceMatrixElement{
						Status:            "NOT_FOUND",
						Duration:          0,
						DurationInTraffic: 0,
						Distance: maps.Distance{
							HumanReadable: "",
							Meters:        0,
						},
					},
					&maps.DistanceMatrixElement{
						Status:            "NOT_FOUND",
						Duration:          0,
						DurationInTraffic: 0,
						Distance: maps.Distance{
							HumanReadable: "",
							Meters:        0,
						},
					},
				},
			},
		},
	}
)
