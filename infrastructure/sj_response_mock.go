package infrastructure

func SearchMock() []byte {
	mockResponse := []byte(`{
  "ERROR_CODE": "EC:0000",
  "ERROR_MESSAGE": "Success.",
  "DATA": [
    [
      {
        "CityFrom": "CGK",
        "CityTo": "DPS",
        "Std": "11-DEC-19 06.00.00 AM",
        "Sta": "11-DEC-19 08.50.00 AM",
        "FlightTime": "110",
        "Segments": [
          {
            "CarrierCode": "SJ",
            "NoFlight": "272",
            "DepartureStation": "CGK",
            "ArrivalStation": "DPS",
            "Std": "11-DEC-19 06.00.00 AM",
            "Sta": "11-DEC-19 08.50.00 AM",
            "Legs": [
              {
                "DepartureStation": "CGK",
                "ArrivalStation": "DPS",
                "Std": "11-DEC-19 06.00.00 AM",
                "Sta": "11-DEC-19 08.50.00 AM",
                "DepartureStationName": "Jakarta",
                "ArrivalStationName": "Denpasar",
                "Std_": "11-Dec-2019 06:00",
                "Sta_": "11-Dec-2019 08:50"
              }
            ],
            "DepartureStationName": "Jakarta",
            "ArrivalStationName": "Denpasar",
            "RouteStatus": "D",
            "Std_": "11-Dec-2019 06:00",
            "Sta_": "11-Dec-2019 08:50"
          }
        ],
        "CityFromName": "Jakarta",
        "CityToName": "Denpasar",
        "Std_": "11-Dec-2019 06:00",
        "Sta_": "11-Dec-2019 08:50",
        "Keterangan": "",
        "ClassesAvailable_B2C": {
          "ECONOMY": [
            {
              "Key": "19837633:Y:S",
              "Availability": "9",
              "Class": "Y",
              "SeatAvail": "OPEN",
              "Price": "10784900",
              "PriceDetail": [
                {
                  "PaxCategory": "ADT",
                  "Total_1": "1540700",
                  "Nta_1": "1507960",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1287000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "85000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "128700",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-32740",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                },
                {
                  "PaxCategory": "CHD",
                  "Total_1": "1540700",
                  "Nta_1": "1507960",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1287000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "85000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "128700",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-32740",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                }
              ],
              "Currency": "IDR",
              "StatusAvail": "Normal",
              "PriceDetail_Y": "",
              "ClassType": "ECONOMY"
            }
          ]
        }
      },
      {
        "CityFrom": "CGK",
        "CityTo": "DPS",
        "Std": "11-DEC-19 09.50.00 AM",
        "Sta": "11-DEC-19 12.40.00 PM",
        "FlightTime": "110",
        "Segments": [
          {
            "CarrierCode": "SJ",
            "NoFlight": "274",
            "DepartureStation": "CGK",
            "ArrivalStation": "DPS",
            "Std": "11-DEC-19 09.50.00 AM",
            "Sta": "11-DEC-19 12.40.00 PM",
            "Legs": [
              {
                "DepartureStation": "CGK",
                "ArrivalStation": "DPS",
                "Std": "11-DEC-19 09.50.00 AM",
                "Sta": "11-DEC-19 12.40.00 PM",
                "DepartureStationName": "Jakarta",
                "ArrivalStationName": "Denpasar",
                "Std_": "11-Dec-2019 09:50",
                "Sta_": "11-Dec-2019 12:40"
              }
            ],
            "DepartureStationName": "Jakarta",
            "ArrivalStationName": "Denpasar",
            "RouteStatus": "D",
            "Std_": "11-Dec-2019 09:50",
            "Sta_": "11-Dec-2019 12:40"
          }
        ],
        "CityFromName": "Jakarta",
        "CityToName": "Denpasar",
        "Std_": "11-Dec-2019 09:50",
        "Sta_": "11-Dec-2019 12:40",
        "Keterangan": "",
        "ClassesAvailable_B2C": {
          "ECONOMY": [
            {
              "Key": "24797752:V:S",
              "Availability": "9",
              "Class": "V",
              "SeatAvail": "OPEN",
              "Price": "8806000",
              "PriceDetail": [
                {
                  "PaxCategory": "ADT",
                  "Total_1": "1258000",
                  "Nta_1": "1230400",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1030000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "85000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "103000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-27600",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                },
                {
                  "PaxCategory": "CHD",
                  "Total_1": "1258000",
                  "Nta_1": "1230400",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1030000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "85000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "103000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-27600",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                }
              ],
              "Currency": "IDR",
              "StatusAvail": "Normal",
              "PriceDetail_Y": {
                "Key": "24797755:Y:S",
                "Availability": "9",
                "Class": "Y",
                "SeatAvail": "OPEN",
                "Price": "10784900",
                "PriceDetail": [
                  {
                    "PaxCategory": "ADT",
                    "Total_1": "1540700",
                    "Nta_1": "1507960",
                    "FareComponent": [
                      {
                        "FareChargeTypeCode": "BF",
                        "FareChargeTypeDesc": "Basic Fare",
                        "Amount": "1287000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "PSC",
                        "FareChargeTypeDesc": "PSC Fee",
                        "Amount": "85000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "VAT",
                        "FareChargeTypeDesc": "PPN - TAX",
                        "Amount": "128700",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "IWJR",
                        "FareChargeTypeDesc": "Government Insurance",
                        "Amount": "5000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "SC",
                        "FareChargeTypeDesc": "Surcharge",
                        "Amount": "0",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "STI",
                        "FareChargeTypeDesc": "Extra Cover",
                        "Amount": "35000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "AC",
                        "FareChargeTypeDesc": "Agent Discount",
                        "Amount": "-32740",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "INCV",
                        "FareChargeTypeDesc": "Incentive",
                        "Amount": "-0",
                        "CurrencyCode": "IDR"
                      }
                    ]
                  },
                  {
                    "PaxCategory": "CHD",
                    "Total_1": "1540700",
                    "Nta_1": "1507960",
                    "FareComponent": [
                      {
                        "FareChargeTypeCode": "BF",
                        "FareChargeTypeDesc": "Basic Fare",
                        "Amount": "1287000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "PSC",
                        "FareChargeTypeDesc": "PSC Fee",
                        "Amount": "85000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "VAT",
                        "FareChargeTypeDesc": "PPN - TAX",
                        "Amount": "128700",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "IWJR",
                        "FareChargeTypeDesc": "Government Insurance",
                        "Amount": "5000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "SC",
                        "FareChargeTypeDesc": "Surcharge",
                        "Amount": "0",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "STI",
                        "FareChargeTypeDesc": "Extra Cover",
                        "Amount": "35000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "AC",
                        "FareChargeTypeDesc": "Agent Discount",
                        "Amount": "-32740",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "INCV",
                        "FareChargeTypeDesc": "Incentive",
                        "Amount": "-0",
                        "CurrencyCode": "IDR"
                      }
                    ]
                  }
                ],
                "Currency": "IDR",
                "StatusAvail": "Normal"
              },
              "ClassType": "ECONOMY"
            }
          ]
        }
      }
    ],
    [
      {
        "CityFrom": "DPS",
        "CityTo": "CGK",
        "Std": "14-DEC-19 01.25.00 PM",
        "Sta": "14-DEC-19 02.20.00 PM",
        "FlightTime": "115",
        "Segments": [
          {
            "CarrierCode": "SJ",
            "NoFlight": "275",
            "DepartureStation": "DPS",
            "ArrivalStation": "CGK",
            "Std": "14-DEC-19 01.25.00 PM",
            "Sta": "14-DEC-19 02.20.00 PM",
            "Legs": [
              {
                "DepartureStation": "DPS",
                "ArrivalStation": "CGK",
                "Std": "14-DEC-19 01.25.00 PM",
                "Sta": "14-DEC-19 02.20.00 PM",
                "DepartureStationName": "Denpasar",
                "ArrivalStationName": "Jakarta",
                "Std_": "14-Dec-2019 13:25",
                "Sta_": "14-Dec-2019 14:20"
              }
            ],
            "DepartureStationName": "Denpasar",
            "ArrivalStationName": "Jakarta",
            "RouteStatus": "D",
            "Std_": "14-Dec-2019 13:25",
            "Sta_": "14-Dec-2019 14:20"
          }
        ],
        "CityFromName": "Denpasar",
        "CityToName": "Jakarta",
        "Std_": "14-Dec-2019 13:25",
        "Sta_": "14-Dec-2019 14:20",
        "Keterangan": "",
        "ClassesAvailable_B2C": {
          "ECONOMY": [
            {
              "Key": "24818765:V:S",
              "Availability": "9",
              "Class": "V",
              "SeatAvail": "OPEN",
              "Price": "8911000",
              "PriceDetail": [
                {
                  "PaxCategory": "ADT",
                  "Total_1": "1273000",
                  "Nta_1": "1245400",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1030000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "100000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "103000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-27600",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                },
                {
                  "PaxCategory": "CHD",
                  "Total_1": "1273000",
                  "Nta_1": "1245400",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1030000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "100000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "103000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-27600",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                }
              ],
              "Currency": "IDR",
              "StatusAvail": "Normal",
              "PriceDetail_Y": {
                "Key": "24818768:Y:S",
                "Availability": "9",
                "Class": "Y",
                "SeatAvail": "OPEN",
                "Price": "10889900",
                "PriceDetail": [
                  {
                    "PaxCategory": "ADT",
                    "Total_1": "1555700",
                    "Nta_1": "1522960",
                    "FareComponent": [
                      {
                        "FareChargeTypeCode": "BF",
                        "FareChargeTypeDesc": "Basic Fare",
                        "Amount": "1287000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "PSC",
                        "FareChargeTypeDesc": "PSC Fee",
                        "Amount": "100000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "VAT",
                        "FareChargeTypeDesc": "PPN - TAX",
                        "Amount": "128700",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "IWJR",
                        "FareChargeTypeDesc": "Government Insurance",
                        "Amount": "5000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "SC",
                        "FareChargeTypeDesc": "Surcharge",
                        "Amount": "0",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "STI",
                        "FareChargeTypeDesc": "Extra Cover",
                        "Amount": "35000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "AC",
                        "FareChargeTypeDesc": "Agent Discount",
                        "Amount": "-32740",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "INCV",
                        "FareChargeTypeDesc": "Incentive",
                        "Amount": "-0",
                        "CurrencyCode": "IDR"
                      }
                    ]
                  },
                  {
                    "PaxCategory": "CHD",
                    "Total_1": "1555700",
                    "Nta_1": "1522960",
                    "FareComponent": [
                      {
                        "FareChargeTypeCode": "BF",
                        "FareChargeTypeDesc": "Basic Fare",
                        "Amount": "1287000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "PSC",
                        "FareChargeTypeDesc": "PSC Fee",
                        "Amount": "100000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "VAT",
                        "FareChargeTypeDesc": "PPN - TAX",
                        "Amount": "128700",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "IWJR",
                        "FareChargeTypeDesc": "Government Insurance",
                        "Amount": "5000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "SC",
                        "FareChargeTypeDesc": "Surcharge",
                        "Amount": "0",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "STI",
                        "FareChargeTypeDesc": "Extra Cover",
                        "Amount": "35000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "AC",
                        "FareChargeTypeDesc": "Agent Discount",
                        "Amount": "-32740",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "INCV",
                        "FareChargeTypeDesc": "Incentive",
                        "Amount": "-0",
                        "CurrencyCode": "IDR"
                      }
                    ]
                  }
                ],
                "Currency": "IDR",
                "StatusAvail": "Normal"
              },
              "ClassType": "ECONOMY"
            }
          ]
        }
      },
      {
        "CityFrom": "DPS",
        "CityTo": "CGK",
        "Std": "14-DEC-19 05.55.00 PM",
        "Sta": "14-DEC-19 06.45.00 PM",
        "FlightTime": "110",
        "Segments": [
          {
            "CarrierCode": "SJ",
            "NoFlight": "273",
            "DepartureStation": "DPS",
            "ArrivalStation": "CGK",
            "Std": "14-DEC-19 05.55.00 PM",
            "Sta": "14-DEC-19 06.45.00 PM",
            "Legs": [
              {
                "DepartureStation": "DPS",
                "ArrivalStation": "CGK",
                "Std": "14-DEC-19 05.55.00 PM",
                "Sta": "14-DEC-19 06.45.00 PM",
                "DepartureStationName": "Denpasar",
                "ArrivalStationName": "Jakarta",
                "Std_": "14-Dec-2019 17:55",
                "Sta_": "14-Dec-2019 18:45"
              }
            ],
            "DepartureStationName": "Denpasar",
            "ArrivalStationName": "Jakarta",
            "RouteStatus": "D",
            "Std_": "14-Dec-2019 17:55",
            "Sta_": "14-Dec-2019 18:45"
          }
        ],
        "CityFromName": "Denpasar",
        "CityToName": "Jakarta",
        "Std_": "14-Dec-2019 17:55",
        "Sta_": "14-Dec-2019 18:45",
        "Keterangan": "",
        "ClassesAvailable_B2C": {
          "ECONOMY": [
            {
              "Key": "19853778:V:S",
              "Availability": "9",
              "Class": "V",
              "SeatAvail": "OPEN",
              "Price": "8911000",
              "PriceDetail": [
                {
                  "PaxCategory": "ADT",
                  "Total_1": "1273000",
                  "Nta_1": "1245400",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1030000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "100000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "103000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-27600",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                },
                {
                  "PaxCategory": "CHD",
                  "Total_1": "1273000",
                  "Nta_1": "1245400",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1030000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "100000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "103000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-27600",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                }
              ],
              "Currency": "IDR",
              "StatusAvail": "Normal",
              "PriceDetail_Y": {
                "Key": "19853781:Y:S",
                "Availability": "9",
                "Class": "Y",
                "SeatAvail": "OPEN",
                "Price": "10889900",
                "PriceDetail": [
                  {
                    "PaxCategory": "ADT",
                    "Total_1": "1555700",
                    "Nta_1": "1522960",
                    "FareComponent": [
                      {
                        "FareChargeTypeCode": "BF",
                        "FareChargeTypeDesc": "Basic Fare",
                        "Amount": "1287000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "PSC",
                        "FareChargeTypeDesc": "PSC Fee",
                        "Amount": "100000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "VAT",
                        "FareChargeTypeDesc": "PPN - TAX",
                        "Amount": "128700",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "IWJR",
                        "FareChargeTypeDesc": "Government Insurance",
                        "Amount": "5000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "SC",
                        "FareChargeTypeDesc": "Surcharge",
                        "Amount": "0",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "STI",
                        "FareChargeTypeDesc": "Extra Cover",
                        "Amount": "35000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "AC",
                        "FareChargeTypeDesc": "Agent Discount",
                        "Amount": "-32740",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "INCV",
                        "FareChargeTypeDesc": "Incentive",
                        "Amount": "-0",
                        "CurrencyCode": "IDR"
                      }
                    ]
                  },
                  {
                    "PaxCategory": "CHD",
                    "Total_1": "1555700",
                    "Nta_1": "1522960",
                    "FareComponent": [
                      {
                        "FareChargeTypeCode": "BF",
                        "FareChargeTypeDesc": "Basic Fare",
                        "Amount": "1287000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "PSC",
                        "FareChargeTypeDesc": "PSC Fee",
                        "Amount": "100000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "VAT",
                        "FareChargeTypeDesc": "PPN - TAX",
                        "Amount": "128700",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "IWJR",
                        "FareChargeTypeDesc": "Government Insurance",
                        "Amount": "5000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "SC",
                        "FareChargeTypeDesc": "Surcharge",
                        "Amount": "0",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "STI",
                        "FareChargeTypeDesc": "Extra Cover",
                        "Amount": "35000",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "AC",
                        "FareChargeTypeDesc": "Agent Discount",
                        "Amount": "-32740",
                        "CurrencyCode": "IDR"
                      },
                      {
                        "FareChargeTypeCode": "INCV",
                        "FareChargeTypeDesc": "Incentive",
                        "Amount": "-0",
                        "CurrencyCode": "IDR"
                      }
                    ]
                  }
                ],
                "Currency": "IDR",
                "StatusAvail": "Normal"
              },
              "ClassType": "ECONOMY"
            }
          ]
        }
      },
      {
        "CityFrom": "DPS",
        "CityTo": "CGK",
        "Std": "14-DEC-19 01.10.00 PM",
        "Sta": "14-DEC-19 03.25.00 PM",
        "FlightTime": "195",
        "Segments": [
          {
            "CarrierCode": "SJ",
            "NoFlight": "107",
            "DepartureStation": "DPS",
            "ArrivalStation": "SUB",
            "Std": "14-DEC-19 01.10.00 PM",
            "Sta": "14-DEC-19 01.10.00 PM",
            "Legs": [
              {
                "DepartureStation": "DPS",
                "ArrivalStation": "SUB",
                "Std": "14-DEC-19 01.10.00 PM",
                "Sta": "14-DEC-19 01.10.00 PM",
                "DepartureStationName": "Denpasar",
                "ArrivalStationName": "Surabaya",
                "Std_": "14-Dec-2019 13:10",
                "Sta_": "14-Dec-2019 13:10"
              }
            ],
            "DepartureStationName": "Denpasar",
            "ArrivalStationName": "Surabaya",
            "RouteStatus": "D",
            "Std_": "14-Dec-2019 13:10",
            "Sta_": "14-Dec-2019 13:10"
          },
          {
            "CarrierCode": "SJ",
            "NoFlight": "269",
            "DepartureStation": "SUB",
            "ArrivalStation": "CGK",
            "Std": "14-DEC-19 01.55.00 PM",
            "Sta": "14-DEC-19 03.25.00 PM",
            "Legs": [
              {
                "DepartureStation": "SUB",
                "ArrivalStation": "CGK",
                "Std": "14-DEC-19 01.55.00 PM",
                "Sta": "14-DEC-19 03.25.00 PM",
                "DepartureStationName": "Surabaya",
                "ArrivalStationName": "Jakarta",
                "Std_": "14-Dec-2019 13:55",
                "Sta_": "14-Dec-2019 15:25"
              }
            ],
            "DepartureStationName": "Surabaya",
            "ArrivalStationName": "Jakarta",
            "RouteStatus": "D",
            "Std_": "14-Dec-2019 13:55",
            "Sta_": "14-Dec-2019 15:25"
          }
        ],
        "CityFromName": "Denpasar",
        "CityToName": "Jakarta",
        "Std_": "14-Dec-2019 13:10",
        "Sta_": "14-Dec-2019 15:25",
        "Keterangan": "",
        "ClassesAvailable_B2C": {
          "ECONOMY": [
            {
              "Key": "28710391:V:S",
              "Availability": "9",
              "Class": "V",
              "SeatAvail": "OPEN",
              "Price": "4514300",
              "PriceDetail": [
                {
                  "PaxCategory": "ADT",
                  "Total_1": "644900",
                  "Nta_1": "628720",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "459000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "100000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "45900",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-16180",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                },
                {
                  "PaxCategory": "CHD",
                  "Total_1": "644900",
                  "Nta_1": "628720",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "459000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "100000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "45900",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-16180",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                }
              ],
              "Currency": "IDR",
              "StatusAvail": "Normal",
              "PriceDetail_Y": "",
              "ClassType": "ECONOMY"
            },
            {
              "Key": "20979851:Y:S",
              "Availability": "9",
              "Class": "Y",
              "SeatAvail": "OPEN",
              "Price": "8365000",
              "PriceDetail": [
                {
                  "PaxCategory": "ADT",
                  "Total_1": "1195000",
                  "Nta_1": "1167000",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1050000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "105000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-28000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                },
                {
                  "PaxCategory": "CHD",
                  "Total_1": "1195000",
                  "Nta_1": "1167000",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1050000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "105000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-28000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                }
              ],
              "Currency": "IDR",
              "StatusAvail": "Normal",
              "PriceDetail_Y": "",
              "ClassType": "ECONOMY"
            }
          ]
        }
      },
      {
        "CityFrom": "DPS",
        "CityTo": "CGK",
        "Std": "14-DEC-19 01.10.00 PM",
        "Sta": "14-DEC-19 08.45.00 PM",
        "FlightTime": "515",
        "Segments": [
          {
            "CarrierCode": "SJ",
            "NoFlight": "107",
            "DepartureStation": "DPS",
            "ArrivalStation": "SUB",
            "Std": "14-DEC-19 01.10.00 PM",
            "Sta": "14-DEC-19 01.10.00 PM",
            "Legs": [
              {
                "DepartureStation": "DPS",
                "ArrivalStation": "SUB",
                "Std": "14-DEC-19 01.10.00 PM",
                "Sta": "14-DEC-19 01.10.00 PM",
                "DepartureStationName": "Denpasar",
                "ArrivalStationName": "Surabaya",
                "Std_": "14-Dec-2019 13:10",
                "Sta_": "14-Dec-2019 13:10"
              }
            ],
            "DepartureStationName": "Denpasar",
            "ArrivalStationName": "Surabaya",
            "RouteStatus": "D",
            "Std_": "14-Dec-2019 13:10",
            "Sta_": "14-Dec-2019 13:10"
          },
          {
            "CarrierCode": "SJ",
            "NoFlight": "267",
            "DepartureStation": "SUB",
            "ArrivalStation": "CGK",
            "Std": "14-DEC-19 07.20.00 PM",
            "Sta": "14-DEC-19 08.45.00 PM",
            "Legs": [
              {
                "DepartureStation": "SUB",
                "ArrivalStation": "CGK",
                "Std": "14-DEC-19 07.20.00 PM",
                "Sta": "14-DEC-19 08.45.00 PM",
                "DepartureStationName": "Surabaya",
                "ArrivalStationName": "Jakarta",
                "Std_": "14-Dec-2019 19:20",
                "Sta_": "14-Dec-2019 20:45"
              }
            ],
            "DepartureStationName": "Surabaya",
            "ArrivalStationName": "Jakarta",
            "RouteStatus": "D",
            "Std_": "14-Dec-2019 19:20",
            "Sta_": "14-Dec-2019 20:45"
          }
        ],
        "CityFromName": "Denpasar",
        "CityToName": "Jakarta",
        "Std_": "14-Dec-2019 13:10",
        "Sta_": "14-Dec-2019 20:45",
        "Keterangan": "",
        "ClassesAvailable_B2C": {
          "ECONOMY": [
            {
              "Key": "28710391:V:S",
              "Availability": "9",
              "Class": "V",
              "SeatAvail": "OPEN",
              "Price": "4514300",
              "PriceDetail": [
                {
                  "PaxCategory": "ADT",
                  "Total_1": "644900",
                  "Nta_1": "628720",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "459000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "100000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "45900",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-16180",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                },
                {
                  "PaxCategory": "CHD",
                  "Total_1": "644900",
                  "Nta_1": "628720",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "459000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "100000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "45900",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-16180",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                }
              ],
              "Currency": "IDR",
              "StatusAvail": "Normal",
              "PriceDetail_Y": "",
              "ClassType": "ECONOMY"
            },
            {
              "Key": "20894909:Y:S",
              "Availability": "9",
              "Class": "Y",
              "SeatAvail": "OPEN",
              "Price": "8365000",
              "PriceDetail": [
                {
                  "PaxCategory": "ADT",
                  "Total_1": "1195000",
                  "Nta_1": "1167000",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1050000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "105000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-28000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                },
                {
                  "PaxCategory": "CHD",
                  "Total_1": "1195000",
                  "Nta_1": "1167000",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1050000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "105000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-28000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                }
              ],
              "Currency": "IDR",
              "StatusAvail": "Normal",
              "PriceDetail_Y": "",
              "ClassType": "ECONOMY"
            }
          ]
        }
      },
      {
        "CityFrom": "DPS",
        "CityTo": "CGK",
        "Std": "14-DEC-19 01.10.00 PM",
        "Sta": "14-DEC-19 09.40.00 PM",
        "FlightTime": "570",
        "Segments": [
          {
            "CarrierCode": "SJ",
            "NoFlight": "107",
            "DepartureStation": "DPS",
            "ArrivalStation": "SUB",
            "Std": "14-DEC-19 01.10.00 PM",
            "Sta": "14-DEC-19 01.10.00 PM",
            "Legs": [
              {
                "DepartureStation": "DPS",
                "ArrivalStation": "SUB",
                "Std": "14-DEC-19 01.10.00 PM",
                "Sta": "14-DEC-19 01.10.00 PM",
                "DepartureStationName": "Denpasar",
                "ArrivalStationName": "Surabaya",
                "Std_": "14-Dec-2019 13:10",
                "Sta_": "14-Dec-2019 13:10"
              }
            ],
            "DepartureStationName": "Denpasar",
            "ArrivalStationName": "Surabaya",
            "RouteStatus": "D",
            "Std_": "14-Dec-2019 13:10",
            "Sta_": "14-Dec-2019 13:10"
          },
          {
            "CarrierCode": "SJ",
            "NoFlight": "255",
            "DepartureStation": "SUB",
            "ArrivalStation": "CGK",
            "Std": "14-DEC-19 08.20.00 PM",
            "Sta": "14-DEC-19 09.40.00 PM",
            "Legs": [
              {
                "DepartureStation": "SUB",
                "ArrivalStation": "CGK",
                "Std": "14-DEC-19 08.20.00 PM",
                "Sta": "14-DEC-19 09.40.00 PM",
                "DepartureStationName": "Surabaya",
                "ArrivalStationName": "Jakarta",
                "Std_": "14-Dec-2019 20:20",
                "Sta_": "14-Dec-2019 21:40"
              }
            ],
            "DepartureStationName": "Surabaya",
            "ArrivalStationName": "Jakarta",
            "RouteStatus": "D",
            "Std_": "14-Dec-2019 20:20",
            "Sta_": "14-Dec-2019 21:40"
          }
        ],
        "CityFromName": "Denpasar",
        "CityToName": "Jakarta",
        "Std_": "14-Dec-2019 13:10",
        "Sta_": "14-Dec-2019 21:40",
        "Keterangan": "",
        "ClassesAvailable_B2C": {
          "ECONOMY": [
            {
              "Key": "28710391:V:S",
              "Availability": "9",
              "Class": "V",
              "SeatAvail": "OPEN",
              "Price": "4514300",
              "PriceDetail": [
                {
                  "PaxCategory": "ADT",
                  "Total_1": "644900",
                  "Nta_1": "628720",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "459000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "100000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "45900",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-16180",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                },
                {
                  "PaxCategory": "CHD",
                  "Total_1": "644900",
                  "Nta_1": "628720",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "459000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "100000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "45900",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-16180",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                }
              ],
              "Currency": "IDR",
              "StatusAvail": "Normal",
              "PriceDetail_Y": "",
              "ClassType": "ECONOMY"
            },
            {
              "Key": "21059469:Y:S",
              "Availability": "9",
              "Class": "Y",
              "SeatAvail": "OPEN",
              "Price": "8365000",
              "PriceDetail": [
                {
                  "PaxCategory": "ADT",
                  "Total_1": "1195000",
                  "Nta_1": "1167000",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1050000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "105000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-28000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                },
                {
                  "PaxCategory": "CHD",
                  "Total_1": "1195000",
                  "Nta_1": "1167000",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1050000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "105000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-28000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                }
              ],
              "Currency": "IDR",
              "StatusAvail": "Normal",
              "PriceDetail_Y": "",
              "ClassType": "ECONOMY"
            }
          ]
        }
      },
      {
        "CityFrom": "DPS",
        "CityTo": "CGK",
        "Std": "14-DEC-19 01.10.00 PM",
        "Sta": "15-DEC-19 06.25.00 AM",
        "FlightTime": "1095",
        "Segments": [
          {
            "CarrierCode": "SJ",
            "NoFlight": "107",
            "DepartureStation": "DPS",
            "ArrivalStation": "SUB",
            "Std": "14-DEC-19 01.10.00 PM",
            "Sta": "14-DEC-19 01.10.00 PM",
            "Legs": [
              {
                "DepartureStation": "DPS",
                "ArrivalStation": "SUB",
                "Std": "14-DEC-19 01.10.00 PM",
                "Sta": "14-DEC-19 01.10.00 PM",
                "DepartureStationName": "Denpasar",
                "ArrivalStationName": "Surabaya",
                "Std_": "14-Dec-2019 13:10",
                "Sta_": "14-Dec-2019 13:10"
              }
            ],
            "DepartureStationName": "Denpasar",
            "ArrivalStationName": "Surabaya",
            "RouteStatus": "D",
            "Std_": "14-Dec-2019 13:10",
            "Sta_": "14-Dec-2019 13:10"
          },
          {
            "CarrierCode": "SJ",
            "NoFlight": "257",
            "DepartureStation": "SUB",
            "ArrivalStation": "CGK",
            "Std": "15-DEC-19 05.00.00 AM",
            "Sta": "15-DEC-19 06.25.00 AM",
            "Legs": [
              {
                "DepartureStation": "SUB",
                "ArrivalStation": "CGK",
                "Std": "15-DEC-19 05.00.00 AM",
                "Sta": "15-DEC-19 06.25.00 AM",
                "DepartureStationName": "Surabaya",
                "ArrivalStationName": "Jakarta",
                "Std_": "15-Dec-2019 05:00",
                "Sta_": "15-Dec-2019 06:25"
              }
            ],
            "DepartureStationName": "Surabaya",
            "ArrivalStationName": "Jakarta",
            "RouteStatus": "D",
            "Std_": "15-Dec-2019 05:00",
            "Sta_": "15-Dec-2019 06:25"
          }
        ],
        "CityFromName": "Denpasar",
        "CityToName": "Jakarta",
        "Std_": "14-Dec-2019 13:10",
        "Sta_": "15-Dec-2019 06:25",
        "Keterangan": "",
        "ClassesAvailable_B2C": {
          "ECONOMY": [
            {
              "Key": "28710391:V:S",
              "Availability": "9",
              "Class": "V",
              "SeatAvail": "OPEN",
              "Price": "4514300",
              "PriceDetail": [
                {
                  "PaxCategory": "ADT",
                  "Total_1": "644900",
                  "Nta_1": "628720",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "459000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "100000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "45900",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-16180",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                },
                {
                  "PaxCategory": "CHD",
                  "Total_1": "644900",
                  "Nta_1": "628720",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "459000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "100000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "45900",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-16180",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                }
              ],
              "Currency": "IDR",
              "StatusAvail": "Normal",
              "PriceDetail_Y": "",
              "ClassType": "ECONOMY"
            },
            {
              "Key": "21106579:Y:S",
              "Availability": "9",
              "Class": "Y",
              "SeatAvail": "OPEN",
              "Price": "8365000",
              "PriceDetail": [
                {
                  "PaxCategory": "ADT",
                  "Total_1": "1195000",
                  "Nta_1": "1167000",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1050000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "105000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-28000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                },
                {
                  "PaxCategory": "CHD",
                  "Total_1": "1195000",
                  "Nta_1": "1167000",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1050000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "105000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-28000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                }
              ],
              "Currency": "IDR",
              "StatusAvail": "Normal",
              "PriceDetail_Y": "",
              "ClassType": "ECONOMY"
            }
          ]
        }
      },
      {
        "CityFrom": "DPS",
        "CityTo": "CGK",
        "Std": "14-DEC-19 04.35.00 PM",
        "Sta": "15-DEC-19 09.50.00 AM",
        "FlightTime": "1095",
        "Segments": [
          {
            "CarrierCode": "IN",
            "NoFlight": "255",
            "DepartureStation": "DPS",
            "ArrivalStation": "SRG",
            "Std": "14-DEC-19 04.35.00 PM",
            "Sta": "14-DEC-19 05.40.00 PM",
            "Legs": [
              {
                "DepartureStation": "DPS",
                "ArrivalStation": "SRG",
                "Std": "14-DEC-19 04.35.00 PM",
                "Sta": "14-DEC-19 05.40.00 PM",
                "DepartureStationName": "Denpasar",
                "ArrivalStationName": "Semarang",
                "Std_": "14-Dec-2019 16:35",
                "Sta_": "14-Dec-2019 17:40"
              }
            ],
            "DepartureStationName": "Denpasar",
            "ArrivalStationName": "Semarang",
            "RouteStatus": "D",
            "Std_": "14-Dec-2019 16:35",
            "Sta_": "14-Dec-2019 17:40"
          },
          {
            "CarrierCode": "SJ",
            "NoFlight": "225",
            "DepartureStation": "SRG",
            "ArrivalStation": "CGK",
            "Std": "15-DEC-19 08.45.00 AM",
            "Sta": "15-DEC-19 09.50.00 AM",
            "Legs": [
              {
                "DepartureStation": "SRG",
                "ArrivalStation": "CGK",
                "Std": "15-DEC-19 08.45.00 AM",
                "Sta": "15-DEC-19 09.50.00 AM",
                "DepartureStationName": "Semarang",
                "ArrivalStationName": "Jakarta",
                "Std_": "15-Dec-2019 08:45",
                "Sta_": "15-Dec-2019 09:50"
              }
            ],
            "DepartureStationName": "Semarang",
            "ArrivalStationName": "Jakarta",
            "RouteStatus": "D",
            "Std_": "15-Dec-2019 08:45",
            "Sta_": "15-Dec-2019 09:50"
          }
        ],
        "CityFromName": "Denpasar",
        "CityToName": "Jakarta",
        "Std_": "14-Dec-2019 16:35",
        "Sta_": "15-Dec-2019 09:50",
        "Keterangan": "",
        "ClassesAvailable_B2C": {
          "ECONOMY": [
            {
              "Key": "24512998:Y:S",
              "Availability": "9",
              "Class": "Y",
              "SeatAvail": "OPEN",
              "Price": "12753300",
              "PriceDetail": [
                {
                  "PaxCategory": "ADT",
                  "Total_1": "1821900",
                  "Nta_1": "1784320",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1529000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "100000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "152900",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-37580",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                },
                {
                  "PaxCategory": "CHD",
                  "Total_1": "1821900",
                  "Nta_1": "1784320",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1529000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "100000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "152900",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-37580",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                }
              ],
              "Currency": "IDR",
              "StatusAvail": "Normal",
              "PriceDetail_Y": "",
              "ClassType": "ECONOMY"
            },
            {
              "Key": "20511402:V:S",
              "Availability": "9",
              "Class": "V",
              "SeatAvail": "OPEN",
              "Price": "4692100",
              "PriceDetail": [
                {
                  "PaxCategory": "ADT",
                  "Total_1": "670300",
                  "Nta_1": "651840",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "573000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "57300",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-18460",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                },
                {
                  "PaxCategory": "CHD",
                  "Total_1": "670300",
                  "Nta_1": "651840",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "573000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "57300",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-18460",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                }
              ],
              "Currency": "IDR",
              "StatusAvail": "Normal",
              "PriceDetail_Y": "",
              "ClassType": "ECONOMY"
            }
          ]
        }
      },
      {
        "CityFrom": "DPS",
        "CityTo": "CGK",
        "Std": "14-DEC-19 05.30.00 PM",
        "Sta": "14-DEC-19 08.45.00 PM",
        "FlightTime": "255",
        "Segments": [
          {
            "CarrierCode": "IN",
            "NoFlight": "663",
            "DepartureStation": "DPS",
            "ArrivalStation": "SUB",
            "Std": "14-DEC-19 05.30.00 PM",
            "Sta": "14-DEC-19 05.30.00 PM",
            "Legs": [
              {
                "DepartureStation": "DPS",
                "ArrivalStation": "SUB",
                "Std": "14-DEC-19 05.30.00 PM",
                "Sta": "14-DEC-19 05.30.00 PM",
                "DepartureStationName": "Denpasar",
                "ArrivalStationName": "Surabaya",
                "Std_": "14-Dec-2019 17:30",
                "Sta_": "14-Dec-2019 17:30"
              }
            ],
            "DepartureStationName": "Denpasar",
            "ArrivalStationName": "Surabaya",
            "RouteStatus": "D",
            "Std_": "14-Dec-2019 17:30",
            "Sta_": "14-Dec-2019 17:30"
          },
          {
            "CarrierCode": "SJ",
            "NoFlight": "267",
            "DepartureStation": "SUB",
            "ArrivalStation": "CGK",
            "Std": "14-DEC-19 07.20.00 PM",
            "Sta": "14-DEC-19 08.45.00 PM",
            "Legs": [
              {
                "DepartureStation": "SUB",
                "ArrivalStation": "CGK",
                "Std": "14-DEC-19 07.20.00 PM",
                "Sta": "14-DEC-19 08.45.00 PM",
                "DepartureStationName": "Surabaya",
                "ArrivalStationName": "Jakarta",
                "Std_": "14-Dec-2019 19:20",
                "Sta_": "14-Dec-2019 20:45"
              }
            ],
            "DepartureStationName": "Surabaya",
            "ArrivalStationName": "Jakarta",
            "RouteStatus": "D",
            "Std_": "14-Dec-2019 19:20",
            "Sta_": "14-Dec-2019 20:45"
          }
        ],
        "CityFromName": "Denpasar",
        "CityToName": "Jakarta",
        "Std_": "14-Dec-2019 17:30",
        "Sta_": "14-Dec-2019 20:45",
        "Keterangan": "",
        "ClassesAvailable_B2C": {
          "ECONOMY": [
            {
              "Key": "28311398:Y:S",
              "Availability": "9",
              "Class": "Y",
              "SeatAvail": "OPEN",
              "Price": "5399800",
              "PriceDetail": [
                {
                  "PaxCategory": "ADT",
                  "Total_1": "771400",
                  "Nta_1": "752920",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "574000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "100000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "57400",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-18480",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                },
                {
                  "PaxCategory": "CHD",
                  "Total_1": "771400",
                  "Nta_1": "752920",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "574000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "100000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "57400",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-18480",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                }
              ],
              "Currency": "IDR",
              "StatusAvail": "Normal",
              "PriceDetail_Y": "",
              "ClassType": "ECONOMY"
            },
            {
              "Key": "20894909:Y:S",
              "Availability": "9",
              "Class": "Y",
              "SeatAvail": "OPEN",
              "Price": "8365000",
              "PriceDetail": [
                {
                  "PaxCategory": "ADT",
                  "Total_1": "1195000",
                  "Nta_1": "1167000",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1050000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "105000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-28000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                },
                {
                  "PaxCategory": "CHD",
                  "Total_1": "1195000",
                  "Nta_1": "1167000",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1050000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "105000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-28000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                }
              ],
              "Currency": "IDR",
              "StatusAvail": "Normal",
              "PriceDetail_Y": "",
              "ClassType": "ECONOMY"
            }
          ]
        }
      },
      {
        "CityFrom": "DPS",
        "CityTo": "CGK",
        "Std": "14-DEC-19 05.30.00 PM",
        "Sta": "14-DEC-19 09.40.00 PM",
        "FlightTime": "310",
        "Segments": [
          {
            "CarrierCode": "IN",
            "NoFlight": "663",
            "DepartureStation": "DPS",
            "ArrivalStation": "SUB",
            "Std": "14-DEC-19 05.30.00 PM",
            "Sta": "14-DEC-19 05.30.00 PM",
            "Legs": [
              {
                "DepartureStation": "DPS",
                "ArrivalStation": "SUB",
                "Std": "14-DEC-19 05.30.00 PM",
                "Sta": "14-DEC-19 05.30.00 PM",
                "DepartureStationName": "Denpasar",
                "ArrivalStationName": "Surabaya",
                "Std_": "14-Dec-2019 17:30",
                "Sta_": "14-Dec-2019 17:30"
              }
            ],
            "DepartureStationName": "Denpasar",
            "ArrivalStationName": "Surabaya",
            "RouteStatus": "D",
            "Std_": "14-Dec-2019 17:30",
            "Sta_": "14-Dec-2019 17:30"
          },
          {
            "CarrierCode": "SJ",
            "NoFlight": "255",
            "DepartureStation": "SUB",
            "ArrivalStation": "CGK",
            "Std": "14-DEC-19 08.20.00 PM",
            "Sta": "14-DEC-19 09.40.00 PM",
            "Legs": [
              {
                "DepartureStation": "SUB",
                "ArrivalStation": "CGK",
                "Std": "14-DEC-19 08.20.00 PM",
                "Sta": "14-DEC-19 09.40.00 PM",
                "DepartureStationName": "Surabaya",
                "ArrivalStationName": "Jakarta",
                "Std_": "14-Dec-2019 20:20",
                "Sta_": "14-Dec-2019 21:40"
              }
            ],
            "DepartureStationName": "Surabaya",
            "ArrivalStationName": "Jakarta",
            "RouteStatus": "D",
            "Std_": "14-Dec-2019 20:20",
            "Sta_": "14-Dec-2019 21:40"
          }
        ],
        "CityFromName": "Denpasar",
        "CityToName": "Jakarta",
        "Std_": "14-Dec-2019 17:30",
        "Sta_": "14-Dec-2019 21:40",
        "Keterangan": "",
        "ClassesAvailable_B2C": {
          "ECONOMY": [
            {
              "Key": "28311398:Y:S",
              "Availability": "9",
              "Class": "Y",
              "SeatAvail": "OPEN",
              "Price": "5399800",
              "PriceDetail": [
                {
                  "PaxCategory": "ADT",
                  "Total_1": "771400",
                  "Nta_1": "752920",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "574000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "100000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "57400",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-18480",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                },
                {
                  "PaxCategory": "CHD",
                  "Total_1": "771400",
                  "Nta_1": "752920",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "574000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "100000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "57400",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-18480",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                }
              ],
              "Currency": "IDR",
              "StatusAvail": "Normal",
              "PriceDetail_Y": "",
              "ClassType": "ECONOMY"
            },
            {
              "Key": "21059469:Y:S",
              "Availability": "9",
              "Class": "Y",
              "SeatAvail": "OPEN",
              "Price": "8365000",
              "PriceDetail": [
                {
                  "PaxCategory": "ADT",
                  "Total_1": "1195000",
                  "Nta_1": "1167000",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1050000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "105000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-28000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                },
                {
                  "PaxCategory": "CHD",
                  "Total_1": "1195000",
                  "Nta_1": "1167000",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1050000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "105000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-28000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                }
              ],
              "Currency": "IDR",
              "StatusAvail": "Normal",
              "PriceDetail_Y": "",
              "ClassType": "ECONOMY"
            }
          ]
        }
      },
      {
        "CityFrom": "DPS",
        "CityTo": "CGK",
        "Std": "14-DEC-19 05.30.00 PM",
        "Sta": "15-DEC-19 06.25.00 AM",
        "FlightTime": "835",
        "Segments": [
          {
            "CarrierCode": "IN",
            "NoFlight": "663",
            "DepartureStation": "DPS",
            "ArrivalStation": "SUB",
            "Std": "14-DEC-19 05.30.00 PM",
            "Sta": "14-DEC-19 05.30.00 PM",
            "Legs": [
              {
                "DepartureStation": "DPS",
                "ArrivalStation": "SUB",
                "Std": "14-DEC-19 05.30.00 PM",
                "Sta": "14-DEC-19 05.30.00 PM",
                "DepartureStationName": "Denpasar",
                "ArrivalStationName": "Surabaya",
                "Std_": "14-Dec-2019 17:30",
                "Sta_": "14-Dec-2019 17:30"
              }
            ],
            "DepartureStationName": "Denpasar",
            "ArrivalStationName": "Surabaya",
            "RouteStatus": "D",
            "Std_": "14-Dec-2019 17:30",
            "Sta_": "14-Dec-2019 17:30"
          },
          {
            "CarrierCode": "SJ",
            "NoFlight": "257",
            "DepartureStation": "SUB",
            "ArrivalStation": "CGK",
            "Std": "15-DEC-19 05.00.00 AM",
            "Sta": "15-DEC-19 06.25.00 AM",
            "Legs": [
              {
                "DepartureStation": "SUB",
                "ArrivalStation": "CGK",
                "Std": "15-DEC-19 05.00.00 AM",
                "Sta": "15-DEC-19 06.25.00 AM",
                "DepartureStationName": "Surabaya",
                "ArrivalStationName": "Jakarta",
                "Std_": "15-Dec-2019 05:00",
                "Sta_": "15-Dec-2019 06:25"
              }
            ],
            "DepartureStationName": "Surabaya",
            "ArrivalStationName": "Jakarta",
            "RouteStatus": "D",
            "Std_": "15-Dec-2019 05:00",
            "Sta_": "15-Dec-2019 06:25"
          }
        ],
        "CityFromName": "Denpasar",
        "CityToName": "Jakarta",
        "Std_": "14-Dec-2019 17:30",
        "Sta_": "15-Dec-2019 06:25",
        "Keterangan": "",
        "ClassesAvailable_B2C": {
          "ECONOMY": [
            {
              "Key": "28311398:Y:S",
              "Availability": "9",
              "Class": "Y",
              "SeatAvail": "OPEN",
              "Price": "5399800",
              "PriceDetail": [
                {
                  "PaxCategory": "ADT",
                  "Total_1": "771400",
                  "Nta_1": "752920",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "574000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "100000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "57400",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-18480",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                },
                {
                  "PaxCategory": "CHD",
                  "Total_1": "771400",
                  "Nta_1": "752920",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "574000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "PSC",
                      "FareChargeTypeDesc": "PSC Fee",
                      "Amount": "100000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "57400",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-18480",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                }
              ],
              "Currency": "IDR",
              "StatusAvail": "Normal",
              "PriceDetail_Y": "",
              "ClassType": "ECONOMY"
            },
            {
              "Key": "21106579:Y:S",
              "Availability": "9",
              "Class": "Y",
              "SeatAvail": "OPEN",
              "Price": "8365000",
              "PriceDetail": [
                {
                  "PaxCategory": "ADT",
                  "Total_1": "1195000",
                  "Nta_1": "1167000",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1050000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "105000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-28000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                },
                {
                  "PaxCategory": "CHD",
                  "Total_1": "1195000",
                  "Nta_1": "1167000",
                  "FareComponent": [
                    {
                      "FareChargeTypeCode": "BF",
                      "FareChargeTypeDesc": "Basic Fare",
                      "Amount": "1050000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "VAT",
                      "FareChargeTypeDesc": "PPN - TAX",
                      "Amount": "105000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "IWJR",
                      "FareChargeTypeDesc": "Government Insurance",
                      "Amount": "5000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "SC",
                      "FareChargeTypeDesc": "Surcharge",
                      "Amount": "0",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "STI",
                      "FareChargeTypeDesc": "Extra Cover",
                      "Amount": "35000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "AC",
                      "FareChargeTypeDesc": "Agent Discount",
                      "Amount": "-28000",
                      "CurrencyCode": "IDR"
                    },
                    {
                      "FareChargeTypeCode": "INCV",
                      "FareChargeTypeDesc": "Incentive",
                      "Amount": "-0",
                      "CurrencyCode": "IDR"
                    }
                  ]
                }
              ],
              "Currency": "IDR",
              "StatusAvail": "Normal",
              "PriceDetail_Y": "",
              "ClassType": "ECONOMY"
            }
          ]
        }
      }
    ]
  ],
  "SearchKey": "020800137085906585866676678869829067807765697785",
  "STI_STATUS": "YES",
  "TIME_INT": "900"
}`)
	return mockResponse
}
func BookInfoMock() []byte {
	mockResponse := []byte(`{
	"ERROR_CODE": "EC:0000",
    "ERROR_MESSAGE": "Success.",
    "DATA": {
      "BookingCode": "QRHEYH",
      "CHECKIN_DATE": "",
      "CHECKIN_FLAG": "",
      "ErrorCode": "RETRIEVE0000",
      "ErrorMessage": "Success.",
      "NUMERIC_BOOKING_CODE": "7779434187632",
      "PAYMENT_METHOD": null,
      "PromoCode": null,
      "Username": "AMBOOKING",
      "YourItineraryDetails": {
        "AdditionalInformation": null,
        "AgentDetails": {
          "BookedBy": "HMNAG006896",
          "IssuedBy": "HMNAG006896"
        },
        "BookingRemarks": null,
        "ContactList": [
          {
            "Description": "Main",
            "Type": "Phone",
            "Value": "08116256665"
          },
          {
            "Description": "Main",
            "Type": "Phone",
            "Value": "02125565555 EXT 33129"
          },
          {
            "Description": "Work",
            "Type": "Email",
            "Value": "PJTIMONICA@GMAIL.COM"
          }
        ],
        "ItineraryDetails": {
          "Journey": [
            {
              "Segment": [
                {
                  "CheckInStatus": "NO",
                  "CityFrom": "PDG",
                  "CityFromName": "Padang",
                  "CityTo": "KNO",
                  "CityToName": "Medan",
                  "Class": "K",
                  "FlightNo": "SJ020",
                  "FlownDate": "23-OCT-19",
                  "ReservationStatus": "RR",
                  "StaLT": "10:50 LT",
                  "StdLT": "09:40 LT"
                }
              ]
            },
            {
              "Segment": [
                {
                  "CheckInStatus": "NO",
                  "CityFrom": "KNO",
                  "CityFromName": "Medan",
                  "CityTo": "PDG",
                  "CityToName": "Padang",
                  "Class": "K",
                  "FlightNo": "SJ021",
                  "FlownDate": "27-OCT-19",
                  "ReservationStatus": "RR",
                  "StaLT": "12:45 LT",
                  "StdLT": "11:35 LT"
                }
              ]
            }
          ]
        },
        "PassengerDetails": [
          {
            "FirstName": "DANIL",
            "FrequentFlyerNumber": null,
            "LastName": "ILHAM",
            "No": "1",
            "SeatNo": "-/-",
            "SeatQty": "1",
            "SpecialRequest": "N/A",
            "Suffix": "Mr",
            "TicketNumber": "9771076125527C1, C2"
          }
        ],
        "PaymentDetails": {
          "AddOn": "0",
          "BasicFare": "1696000",
          "CurrencyCode": "IDR",
          "DirectVoucher": "0",
          "Nta": "1981680",
          "Others": "319600",
          "Sti": null,
          "Total": "2015600"
        },
        "ReservationDetails": {
          "BalanceDue": "0",
          "BalanceDueRemarks": "*Extra Cover Insurance (STI) not include in balance due.",
          "BookingCode": "QRHEYH",
          "BookingDate": "19 Oct 2019 16:03 (GMT+7)",
          "CurrencyCode": "IDR",
          "Status": "Confirm",
          "Time": "19 Oct 2019 16:48 (GMT+7)",
          "TimeDescription": "DateOfIssue"
        }
      },
      "status_show": "Confirmed"
    },
    "ERROR_CODE": "EC:0000",
    "ERROR_MESSAGE": "Success."
  }`)
	return mockResponse
}
func BookMock() []byte {
	mockResponse := []byte(`{
    "ERROR_CODE": "EC:0000",
    "ERROR_MESSAGE": "Success.",
    "DATA": {
        "Username": "AMBOOKING",
        "BookingCode": "FCBNUG",
        "PromoCode": null,
        "YourItineraryDetails": {
            "ReservationDetails": {
                "BookingCode": "FCBNUG",
                "BookingDate": "06 Jul 2019 23:57 (GMT+7)",
                "BalanceDue": "40044600",
                "BalanceDueRemarks": "*Extra Cover Insurance (STI) not include in balance due.",
                "CurrencyCode": "IDR",
                "Time": "07 Jul 2019 00:57 (GMT+7)",
                "TimeDescription": "TimeLimit",
                "Status": "Hold"
            },
            "PassengerDetails": [
                {
                    "No": "1",
                    "Suffix": "Mr",
                    "FirstName": "AGUNG",
                    "LastName": "BUDI",
                    "SeatQty": "1",
                    "TicketNumber": "N/A",
                    "FrequentFlyerNumber": null,
                    "SpecialRequest": "N/A"
                },
                {
                    "No": "2",
                    "Suffix": "Mrs",
                    "FirstName": "JINGGA",
                    "LastName": "RATNA",
                    "SeatQty": "1",
                    "TicketNumber": "N/A",
                    "FrequentFlyerNumber": null,
                    "SpecialRequest": "N/A"
                },
                {
                    "No": "3",
                    "Suffix": "Mr",
                    "FirstName": "RUDI",
                    "LastName": "SANTOSA",
                    "SeatQty": "1",
                    "TicketNumber": "N/A",
                    "FrequentFlyerNumber": null,
                    "SpecialRequest": "N/A"
                },
                {
                    "No": "4",
                    "Suffix": "Mr",
                    "FirstName": "IBRAHIM",
                    "LastName": "SAMAD",
                    "SeatQty": "1",
                    "TicketNumber": "N/A",
                    "FrequentFlyerNumber": null,
                    "SpecialRequest": "N/A"
                },
                {
                    "No": "5",
                    "Suffix": "Mrs",
                    "FirstName": "VIA",
                    "LastName": "RAHMA",
                    "SeatQty": "1",
                    "TicketNumber": "N/A",
                    "FrequentFlyerNumber": null,
                    "SpecialRequest": "N/A"
                },
                {
                    "No": "6",
                    "Suffix": "Mstr",
                    "FirstName": "YUDHA",
                    "LastName": "MUBAROK",
                    "SeatQty": "1",
                    "TicketNumber": "N/A",
                    "FrequentFlyerNumber": null,
                    "SpecialRequest": "N/A"
                },
                {
                    "No": "7",
                    "Suffix": "Miss",
                    "FirstName": "ARMIN",
                    "LastName": "VAHIRA",
                    "SeatQty": "1",
                    "TicketNumber": "N/A",
                    "FrequentFlyerNumber": null,
                    "SpecialRequest": "N/A"
                },
                {
                    "No": "8",
                    "Suffix": "Miss",
                    "FirstName": "CITRA",
                    "LastName": "PUTRI",
                    "SeatQty": "1",
                    "TicketNumber": "N/A",
                    "FrequentFlyerNumber": null,
                    "SpecialRequest": "N/A"
                },
                {
                    "No": "9",
                    "Suffix": "Miss",
                    "FirstName": "INDAH",
                    "LastName": "PERMATA",
                    "SeatQty": "1",
                    "TicketNumber": "N/A",
                    "FrequentFlyerNumber": null,
                    "SpecialRequest": "N/A"
                },
                {
                    "No": "10",
                    "Suffix": "Inf",
                    "FirstName": "MALIKA",
                    "LastName": "PUTRI",
                    "SeatQty": "0",
                    "TicketNumber": "N/A",
                    "FrequentFlyerNumber": null,
                    "SpecialRequest": "N/A"
                },
                {
                    "No": "11",
                    "Suffix": "Inf",
                    "FirstName": "YULAIKHA",
                    "LastName": "PUTRI",
                    "SeatQty": "0",
                    "TicketNumber": "N/A",
                    "FrequentFlyerNumber": null,
                    "SpecialRequest": "N/A"
                },
                {
                    "No": "12",
                    "Suffix": "Inf",
                    "FirstName": "AISHA",
                    "LastName": "PUTRI",
                    "SeatQty": "0",
                    "TicketNumber": "N/A",
                    "FrequentFlyerNumber": null,
                    "SpecialRequest": "N/A"
                }
            ],
            "ItineraryDetails": {
                "Journey": [
                    {
                        "Segment": [
                            {
                                "FlownDate": "20-NOV-19",
                                "FlightNo": "SJ219",
                                "CityFrom": "JOG",
                                "CityTo": "CGK",
                                "CityFromName": "Yogyakarta",
                                "CityToName": "Jakarta",
                                "StdLT": "05:00 LT",
                                "StaLT": "06:00 LT",
                                "ReservationStatus": "HK",
                                "Class": "Y",
                                "CheckInStatus": "NO"
                            },
                            {
                                "FlownDate": "20-NOV-19",
                                "FlightNo": "SJ032",
                                "CityFrom": "CGK",
                                "CityTo": "BTH",
                                "CityFromName": "Jakarta",
                                "CityToName": "Batam",
                                "StdLT": "07:00 LT",
                                "StaLT": "08:35 LT",
                                "ReservationStatus": "HK",
                                "Class": "K",
                                "CheckInStatus": "NO"
                            }
                        ]
                    },
                    {
                        "Segment": [
                            {
                                "FlownDate": "23-NOV-19",
                                "FlightNo": "SJ033",
                                "CityFrom": "BTH",
                                "CityTo": "CGK",
                                "CityFromName": "Batam",
                                "CityToName": "Jakarta",
                                "StdLT": "13:05 LT",
                                "StaLT": "14:45 LT",
                                "ReservationStatus": "HK",
                                "Class": "K",
                                "CheckInStatus": "NO"
                            },
                            {
                                "FlownDate": "23-NOV-19",
                                "FlightNo": "SJ234",
                                "CityFrom": "CGK",
                                "CityTo": "JOG",
                                "CityFromName": "Jakarta",
                                "CityToName": "Yogyakarta",
                                "StdLT": "17:45 LT",
                                "StaLT": "18:55 LT",
                                "ReservationStatus": "HK",
                                "Class": "Y",
                                "CheckInStatus": "NO"
                            }
                        ]
                    }
                ]
            },
            "PaymentDetails": {
                "BasicFare": "35286000",
                "Others": "4758600",
                "Sti": null,
                "Total": "40044600",
                "DirectVoucher": "0",
                "AddOn": "0",
                "Nta": "39361320",
                "CurrencyCode": "IDR"
            },
            "ContactList": [
                {
                    "Type": "Phone",
                    "Description": "Main",
                    "Value": "090897779"
                },
                {
                    "Type": "Email",
                    "Description": "Work",
                    "Value": "AGUNG.UI09@GMAIL.COM"
                }
            ],
            "AgentDetails": {
                "BookedBy": "AMBOOKING",
                "IssuedBy": "-"
            },
            "BookingRemarks": null,
            "AdditionalInformation": null
        },
        "ErrorCode": "GENERATE0000",
        "ErrorMessage": "Success."
    }
}`)
	return mockResponse
}
func SetPaymentMock() []byte {
	mockResponse := []byte(`{
    "ERROR_CODE": "EC:0000",
    "ERROR_MESSAGE": "Success",
    "DATA": [
        {
            "PAYMENT_CODE": "7779133040527",
            "PAYMENT_METHOD": "ATM-BCA",
            "PAYMENT_METHOD_DESCRIPTION": "ATM",
            "AMOUNT": 40044600,
            "FF_DETAIL": null,
            "SPECIAL_REQUEST": "",
            "BNIWOW": "NO"
        }
    ]
}`)
	return mockResponse
}
