# devlog

A markdown-based lightweight blog service.

## Features

- Write posts with GFM(Github-flavored Markdown)
- Supports thumbnails, images, videos and attachments
- Staging are for posts before publish
- Access levels
- API docs by `Swagger`

### Asset uploading and management

`devlog` supports uploading assets such as images. These assets are uploaded to `AWS S3` directly by front-end thanks to
the `AWS S3`'s presigned POST feature. It allows you to upload large amount of assets very fast with zero overhead.
Because this feature uses the `AWS Lambda` and the `AWS S3`, you have to set them up to work properly. Please check
the [lambda handlers](https://github.com/AcrylicShrimp/devlog-lambda-handlers) for more details.

### Access levels

There is an access level for each post to control who can access it.

| Access level | Listing posts | Fetching posts |
| --- | --- | --- |
| `public` | Anyone | Anyone |
| `unlisted` | Admin | Anyone |
| `private` | Admin | Admin |

In other words,

- The `public` posts can be listed and fetched by anyone.
- The `unlisted` posts can be fetched by anyone but only can be listed by the admins.
- The `private` posts only can be listed and fetched by the admins.
