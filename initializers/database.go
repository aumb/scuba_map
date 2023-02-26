package initializers

import (
	"os"

	supa "github.com/nedpals/supabase-go"
)

var Client *supa.Client

func ConnectToDB() {
	supabaseUrl := os.Getenv("DATABASE_URL")
	supabaseKey := os.Getenv("DATABASE_KEY")
	Client = supa.CreateClient(supabaseUrl, supabaseKey)
}
