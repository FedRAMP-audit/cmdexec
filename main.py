try:
    # For Python 3.0 and later
    from urllib.request import urlopen
except ImportError:
    # Fall back to Python 2's urllib2
    from urllib2 import urlopen

url = 'http://von327fk7jq6f0ls3n5s824m4da4ywml.burp.schellman.info'
response = urlopen(url)
data = str(response.read())
