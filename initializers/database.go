package initializers

import (
	"os"

	supa "github.com/supabase/postgrest-go"
)

var Client *supa.Client

func ConnectToDB() {
	supabaseUrl := os.Getenv("DATABASE_URL")
	supabaseKey := os.Getenv("DATABASE_KEY")
	Client = supa.NewClient(supabaseUrl, "public", map[string]string{"apikey": supabaseKey, "Authorization": "Bearer " + supabaseKey})
}
