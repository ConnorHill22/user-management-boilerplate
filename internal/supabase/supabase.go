package sup

import supa "github.com/nedpals/supabase-go"

var Client *supa.Client

func Init(url string, key string) {
	Client = supa.CreateClient(url, key)
}
