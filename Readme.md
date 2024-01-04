# This is an image server created during the Poynt build hackathon using golang.

# SAMPLE REQUESTs
## Uploading an image
### url: localhost:8000/image/upload
### request body: {"data": "base64encodedstrung"}
### response: "imagePath": "image_url"

## Retrieving an image
### url: http://localhost:8000/image/{imageurl}
### response: the image will be served.
