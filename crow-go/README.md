# crow-go

```Go
import (
  "github.com/crowci/crow/v3/crow-go/crow"
  "golang.org/x/oauth2"
)

const (
  token = "dummyToken"
  host  = "http://crow.company.tld"
)

func main() {
  // create an http client with oauth authentication.
  config := new(oauth2.Config)
  authenticator := config.Client(
    oauth2.NoContext,
    &oauth2.Token{
      AccessToken: token,
    },
  )

  // create the crow client with authenticator
  client := crow.NewClient(host, authenticator)

  // gets the current user
  user, err := client.Self()
  fmt.Println(user, err)

  // gets the named repository information
  repo, err := client.RepoLookup("crow-ci/crow")
  fmt.Println(repo, err)
}
```
