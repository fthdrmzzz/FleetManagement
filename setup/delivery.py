import requests
import json

url = "http://localhost:8080/makedelivery"

payload = json.dumps({
  "plate": "34 TL 34",
  "route": [
    {
      "deliveryPoint": 1,
      "deliveries": [
        {
          "barcode": "P7988000121"
        },
        {
          "barcode": "P7988000122"
        },
        {
          "barcode": "P7988000123"
        },
        {
          "barcode": "P8988000121"
        },
        {
          "barcode": "C725799"
        }
      ]
    },
    {
      "deliveryPoint": 2,
      "deliveries": [
        {
          "barcode": "P8988000123"
        },
        {
          "barcode": "P8988000124"
        },
        {
          "barcode": "P8988000125"
        },
        {
          "barcode": "C725799"
        }
      ]
    },
    {
      "deliveryPoint": 3,
      "deliveries": [
        {
          "barcode": "P9988000126"
        },
        {
          "barcode": "P9988000127"
        },
        {
          "barcode": "P9988000128"
        },
        {
          "barcode": "P9988000129"
        },
        {
          "barcode": "P9988000130"
        }
      ]
    }
  ]
})
headers = {
  'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(json.dumps(response.json(), indent=4, sort_keys=True))
