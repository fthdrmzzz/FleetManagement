import requests
import json

url = "http://localhost:8080/vehicles"

payload = json.dumps({
  "plate": "34 TL 34"
})
headers = {
  'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)
print("POST VEHICLE")
print(json.dumps(response.json(), indent=4, sort_keys=True))

print()

url = "http://localhost:8080/deliverypoints"

payload = json.dumps({
  "delivery_points": [
    {
      "id": 1,
      "name": "Branch"
    },
    {
      "id": 2,
      "name": "Distribution Center"
    },
    {
      "id": 3,
      "name": "Transfer Center"
    }
  ]
})
headers = {
  'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)
print("POST DELIVERY POINTS")
print(json.dumps(response.json(), indent=4, sort_keys=True))

print()

url = "http://localhost:8080/bags"

payload = json.dumps({
  "bags": [
    {
      "barcode": "C725799",
      "delivery_id": 2
    },
    {
      "barcode": "C725800",
      "delivery_id": 3
    }
  ]
})
headers = {
  'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)
print("POST BAGS")
print(json.dumps(response.json(), indent=4, sort_keys=True))

print()


url = "http://localhost:8080/packages"

payload = json.dumps({
  "packages": [
    {
      "barcode": "P7988000121",
      "delivery_id": 1,
      "package_weight": 5
    },
    {
      "barcode": "P7988000122",
      "delivery_id": 1,
      "package_weight": 5
    },
    {
      "barcode": "P7988000123",
      "delivery_id": 1,
      "package_weight": 9
    },
    {
      "barcode": "P8988000120",
      "delivery_id": 2,
      "package_weight": 33
    },
    {
      "barcode": "P8988000121",
      "delivery_id": 2,
      "package_weight": 17
    },
    {
      "barcode": "P8988000122",
      "delivery_id": 2,
      "package_weight": 26
    },
    {
      "barcode": "P8988000123",
      "delivery_id": 2,
      "package_weight": 35
    },
    {
      "barcode": "P8988000124",
      "delivery_id": 2,
      "package_weight": 1
    },
    {
      "barcode": "P8988000125",
      "delivery_id": 2,
      "package_weight": 200
    },
    {
      "barcode": "P8988000126",
      "delivery_id": 2,
      "package_weight": 50
    },
    {
      "barcode": "P9988000126",
      "delivery_id": 3,
      "package_weight": 15
    },
    {
      "barcode": "P9988000127",
      "delivery_id": 3,
      "package_weight": 16
    },
    {
      "barcode": "P9988000128",
      "delivery_id": 3,
      "package_weight": 55
    },
    {
      "barcode": "P9988000129",
      "delivery_id": 3,
      "package_weight": 28
    },
    {
      "barcode": "P9988000130",
      "delivery_id": 3,
      "package_weight": 17
    }
  ]
})
headers = {
  'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)
print("POST PACKAGES")
print(json.dumps(response.json(), indent=4, sort_keys=True))

print()

url = "http://localhost:8080/assignpackages"

payload = json.dumps({
  "assignments": [
    {
      "package_barcode": "P8988000122",
      "bag_barcode": "C725799"
    },
    {
      "package_barcode": "P8988000126",
      "bag_barcode": "C725799"
    },
    {
      "package_barcode": "P9988000128",
      "bag_barcode": "C725800"
    },
    {
      "package_barcode": "P9988000129",
      "bag_barcode": "C725800"
    }
  ]
})
headers = {
  'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)

print("POST ASSIGNMENTS")
print(json.dumps(response.json(), indent=4, sort_keys=True))

print()
