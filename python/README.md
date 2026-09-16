# Python client for GMaps Lead Finder

```bash
pip install google-maps-scraper-sdk
export GMF_API_KEY=gmf_your_key_here
gmaps-scraper me
# or: python -m gmaps_scraper me
```

```python
from gmaps_scraper import Client

rows = Client().scrape("dentists in Austin TX")
```

PyPI: https://pypi.org/project/google-maps-scraper-sdk/  
Docs: [../docs/python.md](../docs/python.md) · API: https://gmapsleadfinder.com/docs/api
