# API

Authentication token will be passed by url query _(`...?token=123`)_. The token
won't be required for public link generating.

## Public Create Link

url: `/a/links`

body:

| name | description | value structure | optional |
| - | - | - | - |
| link | The url the short link should point to | An url starting with `http` or `https` | no |

response:

| code | message | data | meaning |
|-|-|-|-|
| 200 | `"link inserted"` | the slug | Link has been accepted and inserted into the database |
| 400 | `"invalid link"` | _none_ | Link has no valid format |
| 403 | `"invalid token"` | _none_ | The given token is not valid |
| 500 | `"try again later"` | _none_ | The server has an current issue |

## Create link

url: `/u/a/links`

body:

| name | description | value structure | optional |
|-|-|-|-|
| link | The url the short link should point to | An url starting with `http` or `https` | no |
| slug | Customize the slug the url will use | Any value | yes |

response:

| code | message | data | meaning |
|-|-|-|-|
| 200 | `"link inserted"` | the slug | Link has been accepted and inserted into the database |
| 400 | `"invalid link"` | _none_ | Link has no valid format |
| 400 | `"slug in use"` | _none_ | The slug you request is already used |
| 403 | `"invalid token"` | _none_ | The given token is not valid |
| 500 | `"try again later"` | _none_ | The server has an current issue |

## Delete link

url: `/u/a/del/:slug:`

body: _none_

response:

| code | message | data | meaning |
|-|-|-|-|
| 200 | `"link deleted"` | _none_ | The link had been deleted |
| 400 | `"link does not exist"` | _none_ | The link could not be found |
| 403 | `"invalid token"` | _none_ | The given token is not valid |
