import sys
import requests


# !!!PUT YOUR API KEY HERE!!!
api_key="YOUR_API_KEY"


if len(sys.argv) != 2:
    print("usage: ", sys.argv[0], " <email>")
    sys.exit(1)

email = sys.argv[1]

j = requests.get("https://api.hubuco.com/api/v3/?api="+api_key+"&"+"email="+email).json()

c = j['resultcode']

if c == 1:
    print("Ok")
elif c == 2:
    print("Catch All")
elif c == 3:
    print("Unknown")
elif c == 4:
    print("Error: " + j['error'])
elif c == 5:
    print("Disposable")
elif c == 6:
    print("Invlaid")

